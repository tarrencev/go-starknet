package contracts

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"
	"unicode"

	"github.com/tarrencev/go-starknet/contracts/abi"
)

func Bind(types []string, abis []string) (string, error) {
	// contracts is the map of each individual contract requested binding
	var contracts = make(map[string]*tmplContract)

	for i := 0; i < len(types); i++ {
		// Parse the actual ABI to generate the binding for
		parsedABI, err := abi.JSON(strings.NewReader(abis[i]))
		if err != nil {
			return "", err
		}

		// Strip any whitespace from the JSON ABI
		strippedABI := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		}, abis[i])

		// Extract the call and transact methods; events, struct definitions; and sort them alphabetically
		var (
			calls     = make(map[string]*tmplMethod)
			transacts = make(map[string]*tmplMethod)
			events    = make(map[string]*tmplEvent)
			fallback  *tmplMethod
			receive   *tmplMethod

			// identifiers are used to detect duplicated identifiers of functions
			// and events. For all calls, transacts and events, abigen will generate
			// corresponding bindings. However we have to ensure there is no
			// identifier collisions in the bindings of these categories.
			callIdentifiers     = make(map[string]bool)
			transactIdentifiers = make(map[string]bool)
		)
		for _, original := range parsedABI.Methods {
			normalized := original
			// Normalize the method for capital cases and non-anonymous inputs/outputs
			name := abi.ToCamelCase(original.Name)
			// Ensure there is no duplicated identifier
			var identifiers = callIdentifiers
			if !original.IsConstant() {
				identifiers = transactIdentifiers
			}
			if identifiers[name] {
				return "", fmt.Errorf("duplicated identifier \"%s\"(normalized \"%s\"), use --alias for renaming", original.Name, name)
			}
			identifiers[name] = true
			normalized.Name = name
			normalized.Inputs = make([]abi.Argument, len(original.Inputs))
			copy(normalized.Inputs, original.Inputs)
			for j, input := range normalized.Inputs {
				if input.Name == "" {
					normalized.Inputs[j].Name = fmt.Sprintf("arg%d", j)
				}
			}
			normalized.Outputs = make([]abi.Argument, len(original.Outputs))
			copy(normalized.Outputs, original.Outputs)
			for j, output := range normalized.Outputs {
				if output.Name != "" {
					normalized.Outputs[j].Name = abi.ToCamelCase(output.Name)
				}
			}
			// Append the methods to the call or transact lists
			if original.IsConstant() {
				calls[original.Name] = &tmplMethod{Original: original, Normalized: normalized, Structured: structured(original.Outputs)}
			} else {
				transacts[original.Name] = &tmplMethod{Original: original, Normalized: normalized, Structured: structured(original.Outputs)}
			}
		}
		// for _, original := range parsedABI.Events {
		// 	// Skip anonymous events as they don't support explicit filtering
		// 	if original.Anonymous {
		// 		continue
		// 	}
		// 	// Normalize the event for capital cases and non-anonymous outputs
		// 	normalized := original

		// 	// Ensure there is no duplicated identifier
		// 	name := methodNormalizer[lang](alias(aliases, original.Name))
		// 	if eventIdentifiers[name] {
		// 		return "", fmt.Errorf("duplicated identifier \"%s\"(normalized \"%s\"), use --alias for renaming", original.Name, name)
		// 	}
		// 	eventIdentifiers[name] = true
		// 	normalized.Name = name

		// 	normalized.Inputs = make([]abi.Argument, len(original.Inputs))
		// 	copy(normalized.Inputs, original.Inputs)
		// 	for j, input := range normalized.Inputs {
		// 		if input.Name == "" {
		// 			normalized.Inputs[j].Name = fmt.Sprintf("arg%d", j)
		// 		}
		// 		if hasStruct(input.Type) {
		// 			bindStructType[lang](input.Type, structs)
		// 		}
		// 	}
		// 	// Append the event to the accumulator list
		// 	events[original.Name] = &tmplEvent{Original: original, Normalized: normalized}
		// }

		contracts[types[i]] = &tmplContract{
			Type:     abi.ToCamelCase(types[i]),
			InputABI: strings.Replace(strippedABI, "\"", "\\\"", -1),
			// InputBin:    strings.TrimPrefix(strings.TrimSpace(bytecodes[i]), "0x"),
			Constructor: parsedABI.Constructor,
			Calls:       calls,
			Transacts:   transacts,
			Fallback:    fallback,
			Receive:     receive,
			Events:      events,
		}
	}

	// Generate the contract template data content and render it
	data := &tmplData{
		Package:   "starknet",
		Contracts: contracts,
	}
	buffer := new(bytes.Buffer)

	funcs := map[string]interface{}{
		"bindtype":      bindBasicType,
		"bindtopictype": bindTopicType,
		"capitalise":    abi.ToCamelCase,
		"decapitalise":  decapitalise,
	}

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(tmplSourceGo))
	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}

	code, err := format.Source(buffer.Bytes())
	if err != nil {
		return "", fmt.Errorf("%v\n%s", err, buffer)
	}
	return string(code), nil
}

// structured checks whether a list of ABI data types has enough information to
// operate through a proper Go struct or if flat returns are needed.
func structured(args abi.Arguments) bool {
	if len(args) < 2 {
		return false
	}
	exists := make(map[string]bool)
	for _, out := range args {
		// If the name is anonymous, we can't organize into a struct
		if out.Name == "" {
			return false
		}
		// If the field name is empty when normalized or collides (var, Var, _var, _Var),
		// we can't organize into a struct
		field := abi.ToCamelCase(out.Name)
		if field == "" || exists[field] {
			return false
		}
		exists[field] = true
	}
	return true
}

// bindBasicType converts basic solidity types(except array, slice and tuple) to Go ones.
func bindBasicType(kind abi.Type) string {
	switch kind.T {
	case abi.FeltTy:
		return "*big.Int"
	case abi.FunctionTy:
		return "[24]byte"
	default:
		// string, bool types
		return kind.String()
	}
}

// bindTopicType converts a Solidity topic type to a Go one. It is almost the same
// functionality as for simple types, but dynamic types get converted to hashes.
func bindTopicType(kind abi.Type) string {
	bound := bindBasicType(kind)

	// todo(rjl493456442) according solidity documentation, indexed event
	// parameters that are not value types i.e. arrays and structs are not
	// stored directly but instead a keccak256-hash of an encoding is stored.
	//
	// We only convert stringS and bytes to hash, still need to deal with
	// array(both fixed-size and dynamic-size) and struct.
	if bound == "string" || bound == "[]byte" {
		bound = "common.Hash"
	}
	return bound
}

// decapitalise makes a camel-case string which starts with a lower case character.
func decapitalise(input string) string {
	if len(input) == 0 {
		return input
	}

	goForm := abi.ToCamelCase(input)
	return strings.ToLower(goForm[:1]) + goForm[1:]
}

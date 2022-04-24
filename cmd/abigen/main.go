package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/urfave/cli/v2"
)

var bindCmd = &cli.Command{
	Name:  "bind",
	Usage: "create ethereum contract bindings from a abis",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "outdir", Usage: "output directory"},
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
	},
	Action: func(ctx *cli.Context) error {
		path, err := filepath.Abs(ctx.Args().First())
		if err != nil {
			return err
		}

		outdir, err := filepath.Abs(ctx.String("outdir"))
		if err != nil {
			return err
		}

		if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
			return err
		}

		if err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				if matched, err := filepath.Match("*.abi", filepath.Base(path)); err != nil {
					return err
				} else if matched {
					dir, fn := filepath.Split(path)
					typ := strings.TrimSuffix(fn, filepath.Ext(fn))
					name := strings.ToLower(typ)
					pkg := strings.ToLower(filepath.Base(outdir))

					if outdir == "" {
						pkg = filepath.Base(dir)
						outdir = dir
					}

					abi, err := ioutil.ReadFile(path)
					if err != nil {
						return fmt.Errorf("reading abi: %w", err)
					}

					bin, err := ioutil.ReadFile(filepath.Join(dir, typ+".bin"))
					if errors.Is(err, os.ErrNotExist) {
						bin = []byte{}
					} else if err != nil {
						return fmt.Errorf("reading abi: %w", err)
					}

					abis := []string{string(abi)}
					types := []string{typ}
					bins := []string{string(bin)}

					// Generate the contract binding
					src, err := bind.Bind(types, abis, bins, nil, pkg, bind.LangGo, nil, nil)
					if err != nil {
						return fmt.Errorf("binding abi: %w", err)
					}

					if err := ioutil.WriteFile(filepath.Join(outdir, name+".go"), []byte(src), 0600); err != nil {
						return fmt.Errorf("writing ABI binding: %w", err)
					}
				}
				return nil
			},
		); err != nil {
			return err
		}
		return nil
	},
}

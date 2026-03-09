package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"code"
)


// test

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return cli.Exit("Please provide two file paths", 1)
			}

			filepath1 := c.Args().Get(0)
			filepath2 := c.Args().Get(1)
			format := c.String("format")

			result, err := code.GenDiff(filepath1, filepath2, format)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

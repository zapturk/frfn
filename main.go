/*
Copyright © 2024 Scott McKeefer scottmckeefer@gmail.com
*/
package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/zapturk/frfn/actions"
)

func main() {
	app := &cli.Command{
		Name:  "frfn",
		Usage: "Find and Replace, Prepend, and Append to file names.",
		Commands: []*cli.Command{
			{
				Name:                  "replace",
				Usage:                 "Replace old text of a file name with new text for all files in the current directory.",
				EnableShellCompletion: true,
				Aliases:               []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "OldText",
						Aliases:  []string{"o"},
						Usage:    "The old text you want to repalce in the file name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "NewText",
						Aliases:  []string{"n"},
						Usage:    "The new text you want to have in the file name",
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "Force",
						Aliases: []string{"f"},
						Usage:   "Skips the dry run step and changes the file names",
						Value:   false,
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					return actions.ReplaceFileName(c.String("OldText"), c.String("NewText"), c.Bool("Force"))
				},
			},
			{
				Name:    "prepend",
				Usage:   "Add text to the front of all files name in the current directory.",
				Aliases: []string{"p"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "Text",
						Aliases:  []string{"t"},
						Usage:    "The text you want to add to the front of the file name",
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "Force",
						Aliases: []string{"f"},
						Usage:   "Skips the dry run step and changes the file names",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					return actions.PrependFileName(c.String("Text"), c.Bool("Force"))
				},
			},
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

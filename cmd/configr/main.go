package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	inventory   *Inventory
	configrPath string
)

func main() {
	// TODO: allow user to specify via environment variable the path
	//       to their inventory file.
	configrPath, _ = os.UserHomeDir()
	configrPath += "/.config/configr/configr.json"

	// If configrPath file doesn't exist, exit with failure.
	// TODO: Create a template file if the file doesnt exist with default
	//       configuration.
	if _, err := os.Stat(configrPath); os.IsNotExist(err) {
		fmt.Printf("Err: could not find file at %s\n", configrPath)
		os.Exit(1)
	}

	inventory = NewInventory(configrPath)

	var app = &cli.App{
		Name:    "configr",
		Usage:   "Configuration file management utility",
		Version: "0.0.1",
	}

	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "file",
			Usage:   "Interact with files in the inventory",
			Aliases: []string{"f"},
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:  "edit",
					Usage: "Open a file in vim",
					Action: func(c *cli.Context) error {
						return cmdEdit(c)
					},
				},
				&cli.Command{
					Name:  "add",
					Usage: "Add a new file to the inventory",
					Action: func(c *cli.Context) error {
						return cmdAddFile(c)
					},
				},
				&cli.Command{
					Name:  "rm",
					Usage: "Remove a file from the inventory",
					Action: func(c *cli.Context) error {
						return cmdRemoveFile(c)
					},
				},
				&cli.Command{
					Name:  "ls",
					Usage: "Print list of files",
					Action: func(c *cli.Context) error {
						return cmdListFiles(c)
					},
				},
				&cli.Command{
					Name:  "inv",
					Usage: "Edit the inventory file",
					Action: func(c *cli.Context) error {
						return cmdEditInventory(c)
					},
				},
			},
		},
		&cli.Command{
			Name:  "edit",
			Usage: "Open a file in vim",
			Action: func(c *cli.Context) error {
				return cmdEdit(c)
			},
		},
	}

	app.Run(os.Args)

}

func cmdListFiles(c *cli.Context) error {
	fmt.Printf("ALIAS\t\tPATH\n")
	for _, file := range inventory.Files {
		fmt.Printf("%s\t\t%s\n", file.Alias, file.Path)
	}
	return nil

}

func cmdEdit(c *cli.Context) error {
	var input = c.Args().Get(0)
	_, file := inventory.FindFile(input)

	file.Edit()
	return nil
}

func cmdEditInventory(c *cli.Context) error {
	var invFile = NewFile(configrPath)
	invFile.Edit()
	return nil
}

func cmdAddFile(c *cli.Context) error {
	var input = c.Args().Get(0)
	inventory.AddFile(input)
	return nil
}

func cmdRemoveFile(c *cli.Context) error {
	var input = c.Args().Get(0)
	inventory.RemoveFile(input)
	return nil
}

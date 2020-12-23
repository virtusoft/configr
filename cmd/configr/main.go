package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	inventory   *Inventory
	collection  *Collection
	configrPath string
)

func main() {
	// TODO: allow user to specify via environment variable the path
	//       to their inventory file.
	var homeDir, _ = os.UserHomeDir()

	configrPath = fmt.Sprintf("%s/%s", homeDir,
		"/.config/configr/configr.json")

	// If configrPath file doesn't exist, exit with failure.
	// TODO: Create a template file if the file doesnt exist with default
	//       configuration.
	if _, err := os.Stat(configrPath); os.IsNotExist(err) {
		fmt.Printf("Err: could not find file at %s\n", configrPath)
		os.Exit(1)
	}

	inventory = NewInventory(configrPath)

	// TODO: Collection path should be configurable by the user.
	collection = NewCollection(fmt.Sprintf("%s/%s", homeDir,
		".cache/configr"))

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
			Name:    "collection",
			Usage:   "interact with the collection",
			Aliases: []string{"c"},
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:  "gather",
					Usage: "Gather files from inventory in collection",
					Action: func(c *cli.Context) error {
						return cmdCollectionGather(c)
					},
				},
				{
					Name:  "deliver",
					Usage: "Deliver files from collection",
					Action: func(c *cli.Context) error {
						return cmdCollectionDeliver(c)
					},
				},
				{
					Name:  "clear",
					Usage: "Delete files in collection",
					Action: func(c *cli.Context) error {
						return cmdCollectionClear(c)
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
		&cli.Command{
			Name:  "inv",
			Usage: "Edit the inventory file",
			Action: func(c *cli.Context) error {
				return cmdEditInventory(c)
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

	// If the file doesn't exist in the inventory,
	// exit with failure.
	if file.Path != "" {
		file.Edit()
	} else {
		fmt.Printf("Could not find \"%s\" in the inventory.\n"+
			"  You may add it with \"configr file add\".\n", input)
		os.Exit(1)
	}

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

func cmdCollectionGather(c *cli.Context) error {
	collection.Gather()
	return nil
}

func cmdCollectionDeliver(c *cli.Context) error {
	collection.Deliver()
	return nil
}

func cmdCollectionClear(c *cli.Context) error {
	collection.Clear()
	return nil
}

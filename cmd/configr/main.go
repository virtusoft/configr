package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/darrienkennedy/configfile"
	"github.com/urfave/cli/v2"
)

var (
	configrPath   string
	configMap     map[string]string
	configFile    *configfile.ConfigFile
	inventory     *Inventory
	collection    *Collection
	inventoryFile *File
)

func main() {
	var homeDir, _ = os.UserHomeDir()
	configrPath = filepath.Join(homeDir, ".config/configr")

	// Create the configrPath directory if it doesn't already exist
	if _, err := os.Stat(configrPath); os.IsNotExist(err) {
		os.MkdirAll(configrPath, os.ModePerm)
	}

	// Initialize configfile based on configrPath
	configFile = configfile.NewConfigFile(filepath.Join(configrPath, "configr.conf"))
	configFile.ConfigData = []*configfile.ConfigData{
		configfile.NewConfigData("collectionPath"),
	}
	configFile.CreateDefault = true

	configFile.Read()
	configMap = configFile.MapConfigs()

	// Initialize inventoryFile variable based on configrPath
	inventoryFile = NewFile(filepath.Join(configrPath, "inventory.json"))

	// If inventory file does not exist on the system,
	// create an empty valid inventory file.
	if _, err := os.Stat(inventoryFile.Path); os.IsNotExist(err) {
		inventory = &Inventory{}
		inventory.WriteToFile()
	} else {
		inventory = NewInventory(inventoryFile.Path)
	}

	// Initialize the collection based on the provided value in the config file
	// for `collectionPath`
	collection = NewCollection(configMap["collectionPath"])

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
				{
					Name:  "path",
					Usage: "Print the path to the collection",
					Action: func(c *cli.Context) error {
						return cmdCollectionPath(c)
					},
				},
			},
		},
		{
			Name:    "config",
			Usage:   "Edit configr configuration settings",
			Aliases: []string{"conf"},
			Action: func(c *cli.Context) error {
				return cmdConfig(c)
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
	var invFile = NewFile(inventoryFile.Path)
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

func cmdCollectionPath(c *cli.Context) error {
	fmt.Println(collection.Path)
	return nil
}

func cmdConfig(c *cli.Context) error {
	var configrConfigFile = NewFile(configFile.Path)
	configrConfigFile.Edit()
	return nil
}

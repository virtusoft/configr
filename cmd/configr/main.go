package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	// TODO: allow user to specify via environment variable the path
	//       to their inventory file.
	var configrPath, _ = os.UserHomeDir()
	configrPath += "/.config/configr/configr.json"

	// If configrPath file doesn't exist, exit with failure.
	// TODO: Create a template file if the file doesnt exist with default
	//       configuration.
	if _, err := os.Stat(configrPath); os.IsNotExist(err) {
		fmt.Printf("Err: could not find file at %s\n", configrPath)
		os.Exit(1)
	}

	var inventory = NewInventory(configrPath)

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
						var input = c.Args().Get(0)
						var found bool = false
						var selectedFile File

						// First check for a path that matches user input.
						for _, file := range inventory.Files {
							if input == file.Path {
								found = true
								selectedFile = *file
								break
							}
						}

						// If the input doesn't match a path, then see if there is an alias which matches.
						if !found {
							for _, file := range inventory.Files {
								if input == file.Alias {
									found = true
									selectedFile = *file
									break
								}
							}
						}

						if found {
							editFile(selectedFile.Path)
						} else {
							fmt.Printf("No matches found for input: `%s`\n", input)
						}
						return nil
					},
				},
				&cli.Command{
					Name:  "ls",
					Usage: "Print list of files",
					Action: func(c *cli.Context) error {
						fmt.Printf("ALIAS\t\tPATH\n")
						for _, file := range inventory.Files {
							fmt.Printf("%s\t\t%s\n", file.Alias, file.Path)
						}
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)

}

func editFile(filePath string) {
	var cmd = exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}

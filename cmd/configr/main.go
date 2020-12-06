package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// TODO: allow user to specify via environment variable the path
	//       to their inventory file.
	var configrPath, _ = os.UserHomeDir()
	configrPath += "/.configr.json"

	// If configrPath file doesn't exist, exit with failure.
	// TODO: Create a template file if the file doesnt exist with default
	//       configuration.
	if _, err := os.Stat(configrPath); os.IsNotExist(err) {
		fmt.Printf("Err: could not find file at %s\n", configrPath)
		os.Exit(1)
	}

	var inventory = NewInventory(configrPath)
	var selectedFile = inventory.Files[0]
	var fd = NewFileData(selectedFile)

	editFile(fd.Path)
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

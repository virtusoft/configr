package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var inventory = NewInventory("file.json")
	var selectedFile = inventory.Files[0]
	var fd = NewFileData(selectedFile)

	editFile(*fd)
}

func editFile(fd FileData) {
	var cmd = exec.Command("vim", fd.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}

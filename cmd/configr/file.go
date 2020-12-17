package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type File struct {
	Path  string
	Alias string
}

func NewFile(path string) *File {
	var file = File{}
	file.Path = path
	file.Alias = ""
	return &file
}

func (f *File) Edit() error {
	var cmd = exec.Command("vim", f.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func (f *File) Copy(toPath string) error {
	// Open file to copy from.
	from, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer from.Close()

	// Open destination file.
	to, err := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	// Copy contents from one file to the other.
	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) GetName() string {
	// Check for matches of just the end of the file path.
	var splitPath = strings.Split(f.Path, "/")
	var fileName = splitPath[len(splitPath)-1]

	// If the file name is just `config` also include the directory
	// in the file name.
	if fileName == "config" {
		fileName = fmt.Sprintf("%s/%s",
			splitPath[(len(splitPath)-2)],
			fileName)
	}

	return fileName
}

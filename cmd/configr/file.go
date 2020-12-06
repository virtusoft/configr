package main

import (
	"os"
	"os/exec"
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

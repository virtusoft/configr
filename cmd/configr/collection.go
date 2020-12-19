package main

import (
	"fmt"
	"os"
)

type Collection struct {
	Path  string
	Files []*File
}

func NewCollection(path string) *Collection {
	var collection = Collection{}
	collection.Path = path
	collection.Files = inventory.GetFiles()
	return &collection
}

func (c *Collection) Gather() {
	var fileName, copyDest string

	// Copy all files from the inventory to a directory
	for _, f := range c.Files {
		fileName = f.GetName()
		copyDest = fmt.Sprintf("%s/%s", c.Path, fileName)
		f.Copy(copyDest)
		// TODO: fix issue with files named `config` not being copied.
	}
}

func (c *Collection) Clear() {
	var fileName, removeFileName string

	for _, f := range c.Files {
		fileName = f.GetName()
		removeFileName = fmt.Sprintf("%s/%s", c.Path, fileName)
		os.Remove(removeFileName)
	}
}

package main

import (
	"fmt"
)

func Gather() {
	var files = inventory.GetFiles()
	var fileName, copyDest string

	// Copy all files from the inventory to a directory
	for _, f := range files {
		fileName = f.GetName()
		copyDest = fmt.Sprintf("%s/%s", collectionPath, fileName)
		fmt.Printf("attempting %s\n", fileName)
		f.Copy(copyDest)
		// TODO: fix issue with files named `config` not being copied.
	}
}

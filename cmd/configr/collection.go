package main

import (
	"fmt"
	"os"
)

func CollectionGather() {
	var files = inventory.GetFiles()
	var fileName, copyDest string

	// Copy all files from the inventory to a directory
	for _, f := range files {
		fileName = f.GetName()
		copyDest = fmt.Sprintf("%s/%s", collectionPath, fileName)
		f.Copy(copyDest)
		// TODO: fix issue with files named `config` not being copied.
	}
}

func CollectionClear() {
	var files = inventory.GetFiles()
	var fileName, removeFileName string

	for _, f := range files {
		fileName = f.GetName()
		removeFileName = fmt.Sprintf("%s/%s", collectionPath, fileName)
		os.Remove(removeFileName)
	}
}

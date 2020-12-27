package main

import (
	"os"
	"path/filepath"
)

type Collection struct {
	Path string
	Data []*CollectionData
}

type CollectionData struct {
	SourceFile *File
	CopyFile   *File
}

func NewCollection(path string) *Collection {
	var collection = Collection{}
	var files = inventory.GetFiles()
	var colData *CollectionData

	collection.Path = path
	for _, f := range files {
		colData = NewCollectionData(f.Path, path)
		collection.Data = append(collection.Data, colData)
	}

	return &collection
}

func NewCollectionData(sourceFilePath, collectionPath string) *CollectionData {
	var collectionData = CollectionData{}
	var fileName, copyFilePath string

	collectionData.SourceFile = NewFile(sourceFilePath)
	fileName = collectionData.SourceFile.GetName()
	copyFilePath = filepath.Join(collectionPath, fileName)

	collectionData.CopyFile = NewFile(copyFilePath)
	return &collectionData
}

func (c *Collection) Gather() {
	// Copy all files from the inventory to a directory
	for _, cd := range c.Data {
		cd.SourceFile.Copy(cd.CopyFile.Path)
		// TODO: fix issue with files named `config` not being copied.
	}
}

func (c *Collection) Deliver() {
	// Copy all files from collection to inventory locations
	for _, cd := range c.Data {
		cd.CopyFile.Copy(cd.SourceFile.Path)
		// TODO: fix issue with files named `config` not being copied.
	}
}

func (c *Collection) Clear() {
	// Remove all files from collection
	for _, cd := range c.Data {
		os.Remove(cd.CopyFile.Path)
	}
}

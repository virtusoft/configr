package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Inventory struct {
	Files []*File
}

type InventoryData struct {
	Files []InventoryFileData
}

type InventoryFileData struct {
	Path  string
	Alias string
}

func NewInventory(inventoryFile string) *Inventory {
	var inventory = Inventory{}
	var inventoryData = InventoryData{}

	var data, err = ioutil.ReadFile(fmt.Sprintf("%s", inventoryFile))
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &inventoryData)
	if err != nil {
		fmt.Println(err)
	}

	for _, strFile := range inventoryData.Files {
		var newFile = NewFile(strFile.Path)
		newFile.Alias = strFile.Alias
		inventory.Files = append(inventory.Files, newFile)
	}

	return &inventory
}

func (i *Inventory) AddFile(input string) error {
	index, _ := i.FindFile(input)

	if index == -1 {
		var file = NewFile(input)
		i.Files = append(i.Files, file)
		i.WriteToFile()
	} else {
		fmt.Println("Err: File already exists in inventory.")
	}

	return nil
}

func (i *Inventory) RemoveFile(input string) error {
	index, _ := i.FindFile(input)

	// If there exists an index
	if index != -1 {
		var arr = i.Files

		// TODO: Validate no index out of bound exception.
		arr = append(arr[:index], arr[index+1:]...)
		i.Files = arr
		i.WriteToFile()
	}

	return nil
}

func (i *Inventory) FindFile(target string) (index int, file *File) {
	// First check for a path that matches user input.
	for index, file := range i.Files {
		// Check for exact matches.
		if target == file.Path {
			return index, file
		}

		// Check for matches of just the end of the file path.
		var splitPath = strings.Split(file.Path, "/")
		var fileName = splitPath[len(splitPath)-1]

		// If the file name is just `config` also include the directory
		// in the file name.
		if fileName == "config" {
			fileName = fmt.Sprintf("%s/%s", splitPath[(len(splitPath)-2)], fileName)
		}

		if target == fileName {
			return index, file
		}
	}

	// If the input doesn't match a path, then see if there is an alias which matches.
	for index, file := range inventory.Files {
		if target == file.Alias {
			return index, file
		}
	}

	// If unable to locate file in inventory, return -1 index
	// and reference to an empty file object.
	return -1, &File{}
}

func (i *Inventory) WriteToFile() error {
	newJsonData, err := json.MarshalIndent(i, "", "\t")
	ioutil.WriteFile(configrPath, newJsonData, 0644)
	return err
}

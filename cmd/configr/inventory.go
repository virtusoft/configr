package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (i *Inventory) GetFiles() []*File {
	return i.Files
}

func (i *Inventory) AddFile(input string) error {
	index, _ := i.FindFile(input)

	if index == -1 {
		var file = NewFile(input)
		if file.Exists() {
			i.Files = append(i.Files, file)
			i.WriteToFile()
		} else {
			fmt.Printf("Err: File '%s' doesn't exist on system.\n", input)
		}
	} else {
		fmt.Printf("Err: File '%s' already exists in inventory.\n", input)
	}

	return nil
}

func (i *Inventory) RemoveFile(input string) error {
	index, _ := i.FindFile(input)

	// If there exists an index
	if index != -1 {
		var arr = i.Files

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

		var fileName = file.GetName()
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
	ioutil.WriteFile(inventoryFile.Path, newJsonData, 0644)
	return err
}

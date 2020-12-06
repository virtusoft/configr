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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Inventory struct {
	Files []string
}

func NewInventory(inventoryFile string) *Inventory {
	var inventory = Inventory{}

	var data, err = ioutil.ReadFile(fmt.Sprintf("%s", inventoryFile))
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &inventory)
	if err != nil {
		fmt.Println(err)
	}

	return &inventory
}

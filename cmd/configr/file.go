package main

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

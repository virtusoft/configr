package main

type FileData struct {
	Path string
}

func NewFileData(path string) *FileData {
	var filedata = FileData{}
	filedata.Path = path
	return &filedata
}

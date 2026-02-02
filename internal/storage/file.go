package storage

const DefaultFilePath = "expenses.json"

type File struct {
	Path string
}

func NewFile(path string) *File {
	return &File{Path: path}
}

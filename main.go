package main

import (
	"os"

	"github.com/udaykumar-dhokia/Gopherledger/cmd"
	"github.com/udaykumar-dhokia/Gopherledger/internal/storage"
)

func main() {
	file := storage.NewFile(storage.DefaultFilePath)

	_, err := os.Stat(file.Path)
	if err != nil {
		os.Create(storage.DefaultFilePath)
	}

	cmd.Execute()
}

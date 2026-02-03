package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/udaykumar-dhokia/Gopherledger/internal/expense"
	"github.com/udaykumar-dhokia/Gopherledger/internal/storage"
)

func OpenAndReadFile(expenses *[]expense.Expense) (status bool, path string) {
	file := storage.NewFile(storage.DefaultFilePath)

	_, err := os.Stat(file.Path)
	if err != nil {
		fmt.Println("No file is connected")
		fmt.Println("Creating new file...")
		os.Create(storage.DefaultFilePath)
		return false, ""
	}

	content, err := os.ReadFile(file.Path)
	if err != nil {
		fmt.Println("Error reading file")
		return false, ""
	}

	if len(content) > 0 {
		json.Unmarshal(content, &expenses)
	}

	return true, file.Path
}

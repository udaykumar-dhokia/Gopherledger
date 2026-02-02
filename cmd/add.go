package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/internal/expense"
	"github.com/udaykumar-dhokia/Gopherledger/internal/storage"
	"github.com/udaykumar-dhokia/Gopherledger/pkg/utils"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Long:  `Add a new expense to your ledger. You need to provide the amount and an optional note.`,
	Run:   add,
	Args:  cobra.ArbitraryArgs,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Add expense")
	addCmd.Flags().StringVarP(&note, "note", "n", "", "Add note")

	addCmd.MarkFlagRequired("amount")
}

func add(cmd *cobra.Command, args []string) {
	file := storage.NewFile(storage.DefaultFilePath)

	_, err := os.Stat(file.Path)
	if err != nil {
		fmt.Println("No file is connected")
		fmt.Println("Creating new file...")
		os.Create(storage.DefaultFilePath)
		return
	}

	content, err := os.ReadFile(file.Path)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	if len(content) > 0 {
		json.Unmarshal(content, &expenses)
	}

	newExpense := expense.Expense{
		ID:        utils.GenerateID(),
		Amount:    amount,
		Note:      note,
		CreatedAt: time.Now().Format("2006-01-02 Mon 15:04"),
	}

	expenses = append(expenses, newExpense)

	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	os.WriteFile(file.Path, data, 0644)
	fmt.Printf("Added transaction worth %.2f on %s with ID: %d\n", amount, newExpense.CreatedAt, newExpense.ID)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/pkg/utils"
)

var (
	id int
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long:  `Delete an existing expense from your ledger using its unique ID.`,
	Args:  cobra.ArbitraryArgs,
	Run:   delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Transaction ID")

	deleteCmd.MarkFlagRequired("id")
}

func delete(cmd *cobra.Command, args []string) {
	status, path := utils.OpenAndReadFile(&expenses)

	if !status {
		return
	}

	idx, exists := utils.FindExpenseByID(expenses, id)

	if !exists {
		fmt.Println("Expense not found")
		return
	}

	expenses = append(expenses[:idx], expenses[idx+1:]...)

	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	os.WriteFile(path, data, 0644)
	fmt.Printf("Deleted transaction worth %.2f with ID: %d\n", expenses[idx].Amount, expenses[idx].ID)
}

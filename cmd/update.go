package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/pkg/utils"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update your expenses",
	Long:  `Update your expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		amount, _ := cmd.Flags().GetFloat64("amount")
		note, _ := cmd.Flags().GetString("note")

		updateExpense(cmd, id, amount, note)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntP("id", "i", 0, "Id of the expense")
	updateCmd.Flags().Float64P("amount", "a", 0, "New amount")
	updateCmd.Flags().StringP("note", "n", "", "Updated note")

	updateCmd.MarkFlagRequired("id")
}

func updateExpense(cmd *cobra.Command, id int, amount float64, note string) {
	status, path := utils.OpenAndReadFile(&expenses)
	if !status {
		return
	}

	idx, exists := utils.FindExpenseByID(expenses, id)
	if !exists {
		fmt.Println("Expense not found")
		return
	}

	if cmd.Flags().Changed("amount") {
		expenses[idx].Amount = amount
	}

	if cmd.Flags().Changed("note") {
		expenses[idx].Note = note
	}

	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	os.WriteFile(path, data, 0644)
	fmt.Printf(
		"Updated transaction worth %.2f with ID: %d\n",
		expenses[idx].Amount,
		expenses[idx].ID,
	)
}

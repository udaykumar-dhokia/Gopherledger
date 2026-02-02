package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/internal/expense"
	"github.com/udaykumar-dhokia/Gopherledger/internal/storage"
	"github.com/udaykumar-dhokia/Gopherledger/pkg/utils"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View all expenses",
	Long:  `View all expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		all, _ := cmd.Flags().GetBool("all")
		year, _ := cmd.Flags().GetString("year")
		if id != 0 {
			viewExpense(id)
		} else if all {
			viewAllExpenses()
		} else if year != "" {
			viewYearExpenses(year)
		} else {
			fmt.Println("Please provide an ID or use the --all flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	viewCmd.Flags().IntP("id", "i", 0, "ID of the expense to view")
	viewCmd.Flags().BoolP("all", "a", false, "View all expenses")
	viewCmd.Flags().StringP("year", "y", "", "View expenses of particular year")
}

func viewExpense(id int) {
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
	idx, exists := utils.FindExpenseByID(expenses, id)

	if !exists {
		fmt.Println("Expense not found")
		return
	}

	expense := expenses[idx]
	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	fmt.Printf("%-5d %-10.2f %-15s %-10s\n", expense.ID, expense.Amount, expense.Note, expense.CreatedAt)

}

func viewAllExpenses() {

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

	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	for _, expense := range expenses {
		fmt.Printf("%-5d %-20.2f %-15s %-10s\n", expense.ID, expense.Amount, expense.Note, expense.CreatedAt)
	}
}

func viewYearExpenses(year string) {
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

	var sortedExpenses []expense.Expense

	for _, exp := range expenses {
		if strings.Contains(exp.CreatedAt, year) {
			sortedExpenses = append(sortedExpenses, exp)
		}
	}

	if len(sortedExpenses) == 0 {
		fmt.Printf("No expenses found for year %s", year)
		return
	}

	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	for _, exp := range sortedExpenses {
		fmt.Printf("%-5d %-20.2f %-15s %-10s\n", exp.ID, exp.Amount, exp.Note, exp.CreatedAt)
	}
}

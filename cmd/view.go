package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/udaykumar-dhokia/Gopherledger/internal/expense"
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
		greater, _ := cmd.Flags().GetFloat64("greater")
		less, _ := cmd.Flags().GetFloat64("less")
		if id != 0 {
			viewExpense(id)
		} else if all {
			viewAllExpenses()
		} else if year != "" {
			viewYearExpenses(year)
		} else if greater != 0 {
			viewGreaterThanExpenses(greater)
		} else if less != 0 {
			viewLessThanExpenses(less)
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
	viewCmd.Flags().Float64P("greater", "g", 0, "View expenses greater than X")
	viewCmd.Flags().Float64P("less", "l", 0, "View expenses less than X")
}

func viewExpense(id int) {

	status, _ := utils.OpenAndReadFile(&expenses)

	if !status {
		return
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

	status, _ := utils.OpenAndReadFile(&expenses)

	if !status {
		return
	}

	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	for _, expense := range expenses {
		fmt.Printf("%-5d %-20.2f %-15s %-10s\n", expense.ID, expense.Amount, expense.Note, expense.CreatedAt)
	}
}

func viewYearExpenses(year string) {

	status, _ := utils.OpenAndReadFile(&expenses)

	if !status {
		return
	}

	var sortedExpenses []expense.Expense

	for _, exp := range expenses {
		if exp.Year == year {
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

func viewGreaterThanExpenses(greater float64) {
	status, _ := utils.OpenAndReadFile(&expenses)

	if !status {
		return
	}

	var sortedExpenses []expense.Expense

	for _, exp := range expenses {
		if exp.Amount >= greater {
			sortedExpenses = append(sortedExpenses, exp)
		}
	}

	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	for _, exp := range sortedExpenses {
		fmt.Printf("%-5d %-20.2f %-15s %-10s\n", exp.ID, exp.Amount, exp.Note, exp.CreatedAt)
	}
}

func viewLessThanExpenses(less float64) {
	status, _ := utils.OpenAndReadFile(&expenses)

	if !status {
		return
	}

	var sortedExpenses []expense.Expense

	for _, exp := range expenses {
		if exp.Amount < less {
			sortedExpenses = append(sortedExpenses, exp)
		}
	}

	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Amount", "Note", "Date")
	for _, exp := range sortedExpenses {
		fmt.Printf("%-5d %-20.2f %-15s %-10s\n", exp.ID, exp.Amount, exp.Note, exp.CreatedAt)
	}
}

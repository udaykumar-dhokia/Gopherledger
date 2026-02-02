package utils

import "github.com/udaykumar-dhokia/Gopherledger/internal/expense"

func FindExpenseByID(expenses []expense.Expense, id int) (int, bool) {
	for i, e := range expenses {
		if e.ID == id {
			return i, true
		}
	}
	return 0, false
}

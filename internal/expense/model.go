package expense

type Expense struct {
	ID        int     `json:"id"`
	Amount    float64 `json:"amount"`
	Note      string  `json:"note"`
	CreatedAt string  `json:"created_at"`
}

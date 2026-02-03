package expense

type Expense struct {
	ID        int     `json:"id"`
	Amount    float64 `json:"amount"`
	Note      string  `json:"note"`
	CreatedAt string  `json:"created_at"`
	Date      string  `json:"date"`
	Month     string  `json:"month"`
	Year      string  `json:"year"`
}

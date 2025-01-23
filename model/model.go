package model

type Expense struct {
	TimeStamp string
	Category  string
	Amount    float64
}

type ExpenseJSON struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

type PageData struct {
	LoggedIn        bool
	Expenses        []Expense
	CurrentMonthSum float64
}

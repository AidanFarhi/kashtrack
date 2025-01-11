package model

type Expense struct {
	TimeStamp string
	Category  string
	Amount    float64
}

type PageData struct {
	Expenses []Expense
}

package responses

import "time"

type OrdersCustom struct {
	OrderId      int64
	OrderNumber  int64
	Quantity     int
	Cost         float32
	Currency     int64
	OrderDate    time.Time
	DateCreated  time.Time
	DateModified time.Time
}

type TransactionsCustom struct {
	TransactionId       int64
	Order               *OrdersCustom
	Amount              float32
	TransactingCurrency int64
	StatusId            int64
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
}

type OrderOriResponseDTO struct {
	StatusCode  int
	Transaction *TransactionsCustom
	StatusDesc  string
}

type OrderResponseDTO struct {
	Success    bool
	Result     *TransactionsCustom
	StatusDesc string
}

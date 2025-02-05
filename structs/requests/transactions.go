package requests

type PostTransactionRequest struct {
	Items                []Item
	CurrencyId           int
	RequestType          string
	PaymentProofImageUrl string
	PaymentMethodId      int64
	Comment              string
	OrderLocation        string
	OrderBy              int64
	OrderStartDate       string
	OrderEndDate         string
	CustomerId           int64
}

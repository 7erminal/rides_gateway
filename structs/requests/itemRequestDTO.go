package requests

type AddItemRequestDTO struct {
	ProductName  string
	Quantity     int
	CostPrice    float64
	SellingPrice float64
	BranchId     int64
	ImagePath    string
	ReorderLevel int
}

type AddSalesItemRequestDTO struct {
	ProductName  string
	Quantity     int
	CostPrice    float64
	SellingPrice float64
	BranchId     int64
	ImagePath    string
}

type AddRentalItemRequestDTO struct {
	ProductName  string
	Quantity     int
	ReorderLevel int
	RentalPrice  float64
	BranchId     int64
	ImagePath    string
}

type UpdateItemRequestDTO struct {
	ProductName   string
	Quantity      int
	CostPrice     float64
	SellingPrice  float64
	BranchId      int64
	ImagePath     string
	ProductTypeId int64
}

type Product struct {
	ProductId int64
	Quantity  int64
}

type Item struct {
	ItemId   int64
	Quantity int64
}

type RentalRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      int64
	OrderLocation        string
	CustomerId           int64
	OrderStartDate       string
	OrderEndDate         string
}

type SalesRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      int64
	CustomerId           int64
	OrderDate            string
}

package requests

type AddCustomer struct {
	Email       string
	Name        string
	PhoneNumber string
	Location    string
	IdType      string
	IdNumber    string
	ImagePath   string
}

type UpdateCustomer struct {
	Email       string
	Name        string
	PhoneNumber string
	Location    string
	IdType      string
	IdNumber    string
	ImagePath   string
}

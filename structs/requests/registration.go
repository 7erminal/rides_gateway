package requests

type Registration struct {
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	Token       string
}

type RegisterUser struct {
	Email       string
	Name        string
	Gender      string
	PhoneNumber string
	Password    string
	Dob         string
	RoleId      string
}

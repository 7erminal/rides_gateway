package requests

type StringRequestDTO struct {
	Value string
}

type UsernameRequestDTO struct {
	Username string
}

type InviteRequestDTO struct {
	Email string
	Role  int64
}

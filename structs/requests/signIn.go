package requests

type SignIn struct {
	Email    string
	Password string
}

type ChangePassword struct {
	OldPassword string
	NewPassword string
}

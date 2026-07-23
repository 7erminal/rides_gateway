package requests

type SignIn struct {
	Email    string
	Password string
}

type ChangePassword struct {
	OldPassword string
	NewPassword string
}

type SendOTP struct {
	MobileNumber string
}

type VerifyOTP struct {
	MobileNumber string
	Code         string
}

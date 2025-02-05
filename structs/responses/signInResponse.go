package responses

type UserOriResponseDTO struct {
	StatusCode int
	User       *UsersOri
	StatusDesc string
}

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}

type UserGatewayResponseDTO struct {
	Success    bool
	Result     *UserGateway
	StatusDesc string
}

type UsersGatewayResponseDTO struct {
	Success    bool
	Result     *[]UserGateway
	StatusDesc string
}

type UsersOriResponseDTO struct {
	StatusCode int
	Users      *[]UsersOri
	StatusDesc string
}

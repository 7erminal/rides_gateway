package responses

type TokenDestructureResponseDTO struct {
	Email  string
	RoleId string
}

type InviteDecodeResponseDTO struct {
	StatusCode int
	Value      *TokenDestructureResponseDTO
	StatusDesc string
}

type StringResponseDTO struct {
	Success    bool
	Result     *string
	StatusDesc string
}

type StringOriResponseDTO struct {
	StatusCode int
	Value      string
	StatusDesc string
}

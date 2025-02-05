package responses

type IDTypeResponse struct {
	IdentificationTypeId int64
	Name                 string
	Code                 string
}

type IDTypeResponseDTO struct {
	StatusCode int
	IdTypes    *[]IDTypeResponse
	StatusDesc string
}

type IDTypesGatewayResponseDTO struct {
	Success    bool
	Result     *[]IDTypeResponse
	StatusDesc string
}

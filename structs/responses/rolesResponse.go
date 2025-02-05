package responses

// type RolesResponseDTO struct {
// 	StatusCode int
// 	Roles      *[]models.Roles
// 	StatusDesc string
// }

type Roles struct {
	RoleId      int64  `orm:"auto"`
	Role        string `orm:"size(100)"`
	Description string `orm:"size(500)"`
	// Active      int
}

type RoleResponseDTO struct {
	StatusCode int
	Role       *Roles
	StatusDesc string
}

type RolesAllResponseDTO struct {
	StatusCode int
	Roles      *[]Roles
	StatusDesc string
}

type RolesAllGatewayResponseDTO struct {
	Success    bool
	Result     *[]Roles
	StatusDesc string
}

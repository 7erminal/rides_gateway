package responses

type PaymentMethods struct {
	PaymentMethodId int64  `orm:"auto"`
	PaymentMethod   string `orm:"size(128)"`
	// DateCreated     time.Time `orm:"type(datetime)"`
	// DateModified    time.Time `orm:"type(datetime)"`
	// CreatedBy       int
	// ModifiedBy      int
	// Active          int
}

type PaymentMethodsOriResponseDTO struct {
	StatusCode     int
	PaymentMethods *[]PaymentMethods
	StatusDesc     string
}

type PaymentMethodsResponseDTO struct {
	Success    bool
	Result     *[]PaymentMethods
	StatusDesc string
}

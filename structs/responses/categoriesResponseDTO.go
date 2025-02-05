package responses

type Categories struct {
	CategoryId   int64  `orm:"auto;omitempty"`
	CategoryName string `orm:"size(40)"`
	ImagePath    string `orm:"size(250)"`
	Icon         string `orm:"size(250)"`
}

type CategoriesResponseDTO struct {
	Success    bool
	Result     *[]Categories
	StatusDesc string
}

type CategoriesOriResponseDTO struct {
	StatusCode int
	Categories *[]Categories
	StatusDesc string
}

type CategoryOriResponseDTO struct {
	StatusCode int
	Category   *Categories
	StatusDesc string
}

type CategoryResponseDTO struct {
	Success    bool
	Result     *Categories
	StatusDesc string
}

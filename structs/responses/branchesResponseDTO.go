package responses

import (
	"time"
)

type Currencies struct {
	CurrencyId   int64     `orm:"auto;omitempty"`
	Symbol       string    `orm:"size(20)"`
	Currency     string    `orm:"size(50)"`
	Active       int       `orm:"omitempty"`
	DateCreated  time.Time `orm:"type(datetime);omitempty"`
	DateModified time.Time `orm:"type(datetime);omitempty"`
	CreatedBy    int       `orm:"omitempty"`
	ModifiedBy   int       `orm:"omitempty"`
}

type CurrencyResp struct {
	Symbol   string
	Currency string
}

type Countries struct {
	CountryId       int64  `orm:"auto"`
	Country         string `orm:"size(255)"`
	Description     string `orm:"size(500)"`
	CountryCode     string `orm:"size(20)"`
	DefaultCurrency *Currencies
	DateCreated     time.Time `orm:"type(datetime)"`
	DateModified    time.Time `orm:"type(datetime)"`
	CreatedBy       int
	ModifiedBy      int
}

type CountryRespOri struct {
	Country         string
	CountryCode     string
	DefaultCurrency *CurrencyResp
}

type CountryResp struct {
	Country     string
	CountryCode string
	Currency    *CurrencyResp
}

type CountriesOriResponseDTO struct {
	StatusCode int
	Countries  *[]Countries
	StatusDesc string
}

type CountriesResponseDTO struct {
	Success    bool
	Result     *[]CountryResp
	StatusDesc string
}

type Branches struct {
	BranchId      int64      `orm:"auto"`
	Branch        string     `orm:"size(80)"`
	Country       *Countries `orm:"rel(fk);column(country)"`
	Location      string
	PhoneNumber   string
	Active        int       `orm:"omitempty"`
	DateCreated   time.Time `orm:"type(datetime);omitempty"`
	DateModified  time.Time `orm:"type(datetime);omitempty"`
	CreatedBy     int       `orm:"omitempty"`
	ModifiedBy    int       `orm:"omitempty"`
	BranchManager *Users
}

type BranchRespOri struct {
	BranchId    int64
	Branch      string
	Country     *CountryRespOri
	Location    string
	PhoneNumber string
	DateCreated time.Time
}

type BranchResp struct {
	BranchId int64
	Branch   string
	// Country     *CountryResp
	Location      string
	PhoneNumber   string
	BranchManager *UserGateway
	DateCreated   time.Time
}

type BranchesOriResponseDTO struct {
	StatusCode int
	Branches   *[]Branches
	StatusDesc string
}

type BranchesData struct {
	Data  *[]BranchResp
	Count int
}

type BranchesResponseDTO struct {
	Success    bool
	Result     *BranchesData
	StatusDesc string
}

type BranchOriResponseDTO struct {
	StatusCode int
	Branch     *Branches
	StatusDesc string
}

type BranchResponseDTO struct {
	Success    bool
	Result     *BranchResp
	StatusDesc string
}

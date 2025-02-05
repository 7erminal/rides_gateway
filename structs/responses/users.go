package responses

import (
	"time"
)

type Role struct {
	Role string
}

type UserExtraDetails struct {
	// CustomerId int64
	// User       int64
	Branch *BranchResp
	// Shop             *Shops
	// CustomerCategory *Customer_categories
	// Nickname         string
	// DateCreated      time.Time
	// DateModified     time.Time
	// CreatedBy        int
	// ModifiedBy       int
	// Active           int
}

type UsersOri struct {
	UserId        int64
	UserType      int
	ImagePath     string
	UserDetails   *UserExtraDetails
	FullName      string
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	Gender        string
	Dob           time.Time
	Address       string
	IdType        string
	IdNumber      string
	Role          *Role
	MaritalStatus string
	Active        int
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
}

type Users struct {
	UserId        int64 `orm:"auto"`
	UserType      int
	ImagePath     string
	Customer      *UserExtraDetails `orm:"rel(fk);column(customer_id)"`
	FullName      string            `orm:"size(255)"`
	Username      string            `orm:"size(255)"`
	Password      string            `orm:"size(255)"`
	Email         string            `orm:"size(255)"`
	PhoneNumber   string            `orm:"size(255)"`
	Gender        string            `orm:"size(10)"`
	Dob           time.Time         `orm:"type(datetime)"`
	Address       string            `orm:"size(255)"`
	IdType        string            `orm:"size(5)"`
	IdNumber      string            `orm:"size(100)"`
	Role          *Role
	MaritalStatus string `orm:"size(255);omitempty"`
	Active        int
	IsVerified    bool
	DateCreated   time.Time `orm:"type(datetime)"`
	DateModified  time.Time `orm:"type(datetime)"`
	CreatedBy     int
	ModifiedBy    int
	Branch        *BranchResp
}

type UserGateway struct {
	UserId int64 `orm:"auto"`
	// UserType    int
	FirstName   string `orm:"size(255)"`
	LastName    string `orm:"size(255)"`
	Username    string `orm:"size(255)"`
	Email       string `orm:"size(255)"`
	PhoneNumber string `orm:"size(255)"`
	ImagePath   string
	Customer    *UserExtraDetails
	// Gender         string    `orm:"size(10)"`
	// Dob            time.Time `orm:"type(datetime)"`
	// Address        string    `orm:"size(255)"`
	// IdType         string    `orm:"size(5)"`
	// IdNumber       string    `orm:"size(100)"`
	Status         string
	IsVerified     bool
	Role           *Role
	DateRegistered time.Time `orm:"type(datetime)"`
	// DateModified time.Time `orm:"type(datetime)"`
	// CreatedBy    int
	// ModifiedBy   int
}

type UserTokens struct {
	Token      string
	ExpiryDate time.Time
}

type UserInvitesOri struct {
	UserInviteId    int64
	InvitedBy       *UsersOri
	InvitationToken *UserTokens
	Email           string
	Role            *Role
	Status          string
	DateCreated     time.Time
}

type UserInvites struct {
	UserInviteId int64
	// InvitedBy    *UserGateway
	// InvitationToken *UserTokens
	Email       string
	Role        string
	Status      string
	DateCreated time.Time
}

type UserInvitesResponseDTO struct {
	StatusCode  int
	UserInvites *[]UserInvitesOri
	StatusDesc  string
}

type UserInvitesResponse struct {
	Success    bool
	Result     *[]UserInvites
	StatusDesc string
}

type UserInviteResponseDTO struct {
	StatusCode int
	UserInvite *UserInvitesOri
	StatusDesc string
}

type UserInviteResponse struct {
	Success    bool
	Result     *UserInvites
	StatusDesc string
}

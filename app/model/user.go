package model

import (
	"time"

	"github.com/erwinwahyura/go-boilerplate/utils"
)

type IdentityType string

const (
	KTP      IdentityType = "ktp"
	SIM      IdentityType = "sim"
	PASSPORT IdentityType = "passport"
	KITAS    IdentityType = "kitas"
)

func (s *User) ToUserRequest() UserRequest {
	return UserRequest{
		Email:              s.Email,
		FirstName:          utils.PtrToValue(s.FirstName),
		LastName:           utils.PtrToValue(s.LastName),
		PhoneNumber:        utils.PtrToValue(s.PhoneNumber),
		Username:           utils.PtrToValue(s.Username),
		Password:           utils.PtrToValue(s.Password),
		LastLogin:          utils.PtrToValue(s.LastLogin),
		IsSuperUser:        s.IsSuperUser,
		IsStaff:            s.IsStaff,
		IsActive:           s.IsActive,
		IsVerified:         s.IsVerified,
		Properties:         s.Properties,
		CorporateAccountID: s.CorporateAccountID,
		AuthorID:           s.AuthorID,
	}
}

// NOTE: maybe we can remove the pointer and use COALESCE on the query instead?
type User struct {
	ID          int64   `db:"id"`
	Email       string  `db:"email"`
	FirstName   *string `db:"first_name"`
	LastName    *string `db:"last_name"`
	PhoneNumber *string `db:"phone_number"`
	Username    *string `db:"username"`
	// sample password
	// scoop$ffff2ab42c2a9423c27f86487b65d92babe4e0a15e49b4a42e4bc50de2621692$4a4a578741bf263d334c2b68a85556a4796460b32a50af6685d45de6a201fb1201bdefc1a37b570ab0f839b4cc1c5de5c9a81859e6deb98d2df734617dddcdd9
	// 2917b0dc1a2b32ce0c4c1910815a0c6add1e08be86fe4ecf95be7848e46f9d3a
	Password    *string    `db:"password"`
	LastLogin   *time.Time `db:"last_login"`
	IsSuperUser bool       `db:"is_superuser"`
	IsStaff     bool       `db:"is_staff"`
	IsActive    bool       `db:"is_active"`
	IsVerified  bool       `db:"verified"` // this column didn't exists it the current staging db (gb_staging2)

	// i think this one can be removed since its not an effective business
	IsGuest bool `db:"is_guest"`

	// can be removed *?
	IsDeleted bool `db:"is_deleted"` // this column didn't exists it the current staging db (gb_staging2)

	// if we change to the new db then should be change the name into created_at instead of date_joined
	CreatedAt time.Time `db:"date_joined"`

	// not quite sure properties in legacy db is data type array, should be string and has a separator (; or ,)
	Properties string `db:"properties"`

	// corporate_account_id is a FK from table corporate_partner
	CorporateAccountID int64 `db:"corporate_account_id"`

	// new added, a FK from table author
	AuthorID int64 `db:"author_id"`

	// user_profile's data
	BirthPlace *string    `db:"birth_place"`
	BirthDate  *time.Time `db:"birth_date"`
	Gender     *string    `db:"gender"`
	HomePhone  *string    `db:"home_phone_number"`
	Job        *string    `db:"occupation"`

	// the legacy DB is using array as data type, can use string and separate it with (, or ;)
	Hobby *string `db:"hobby"`

	// fk from table user
	// UserID         int64  `db:"user_id"`
	IdentityImage  *string       `db:"identity_image"`
	IdentityNumber *string       `db:"identity_number"`
	IdentityType   *IdentityType `db:"identity_type"`
}

// sample data of user
// 	[
// 	{
// 		"id": 13,
// 		"password": "",
// 		"last_login": null,
// 		"is_superuser": false,
// 		"is_staff": false,
// 		"is_active": true,
// 		"date_joined": "2017-07-27 11:23:20.189689+00",
// 		"email": "author-yoshiki-naka@email.com",
// 		"username": "author-yoshiki-naka",
// 		"is_guest": false,
// 		"first_name": "Yoshiki",
// 		"last_name": "Naka",
// 		"phone_number": null,
// 		"properties": null,
// 		"verified": false,
// 		"corporate_account_id": null,
// 		"is_delete": false
// 	}
// ]

// Table User and UserProfile can be merged into one table
type ModifiedUser struct {
	ID                 int64     `db:"id"`
	Email              string    `db:"email"`
	FirstName          string    `db:"first_name"`
	LastName           string    `db:"last_name"`
	Username           string    `db:"username"`
	Password           string    `db:"password"`
	LastLogin          time.Time `db:"last_login"`
	IsSuperUser        bool      `db:"is_superuser"`
	IsStaff            bool      `db:"is_staff"`
	IsActive           bool      `db:"is_active"`
	IsVerified         bool      `db:"verified"`
	Properties         string    `db:"properties"`
	CorporateAccountID int64     `db:"corporate_account_id"`
	AuthorID           int64     `db:"author_id"`
	BirthPlace         string    `db:"birth_place"`
	BirthDate          time.Time `db:"birth_date"`
	Gender             string    `db:"gender"`
	HomePhone          string    `db:"home_phone_number"`
	PhoneNumber        string    `db:"phone_number"`
	Job                string    `db:"occupation"`
	Hobby              string    `db:"hobby"`
	IdentityImage      string    `db:"identity_image"`
	IdentityNumber     string    `db:"identity_number"`
	IdentityType       string    `db:"identity_type"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Request and Response struct User below

type UserRequest struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	IsSuperUser bool      `json:"is_superuser"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	IsVerified  bool      `json:"verified"`
	// not quite sure properties in legacy db is data type array, should be string and has a separator (; or ,)
	Properties string `db:"properties"`

	// corporate_account_id is a FK from table corporate_partner
	CorporateAccountID int64 `db:"corporate_account_id"`

	// new added, a FK from table author
	AuthorID int64 `db:"author_id"`
}

type ErrorMessageCode string

const (
	INVALID_USERNAME    ErrorMessageCode = "INVALID_USERNAME"
	INVALID_PASSWORD    ErrorMessageCode = "INVALID_PASSWORD"
	INVALID_PHONENUMBER ErrorMessageCode = "INVALID_PHONENUMBER"
	INVALID_EMAIL       ErrorMessageCode = "INVALID_EMAIL"
)

// List of Messages
var errorMessages = map[ErrorMessageCode]string{
	INVALID_USERNAME:    "username is invalid",
	INVALID_PASSWORD:    "password is invalid",
	INVALID_PHONENUMBER: "phone is invalid",
	INVALID_EMAIL:       "email is invalid",
}

// ErrorMessage converts to its string representation
func (code ErrorMessageCode) String() string {
	return errorMessages[code]
}

// String converts to its string representation
func (code ErrorMessageCode) Code() string {
	return string(code)
}

type ProfileResponse struct {
	Email        string `json:"email"`
	Fullname     string `json:"fullname"`
	MyValuePoint int    `json:"myvalue_point"`
	Avatar       string `json:"avatar"`
}

package repository

import (
	"context"
	"fmt"

	"github.com/erwinwahyura/go-boilerplate/app/database"
	"github.com/erwinwahyura/go-boilerplate/app/model"

	"github.com/erwinwahyura/go-boilerplate/utils"
)

var (
	TableUser = fmt.Sprintf("%v.%v", "public", "user")
)

type UserRepositoryFilter struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	// This field is added for registration via MyValue to check if user with specified email already exists
	// If ValueID from MyValueAPI's UserInfo was supposed to be saved as Username then this field can be deleted.
	Email *string `json:"email"`
}

type (

	// Repository Inteface
	UserRepository interface {
		Create(ctx context.Context, user string) (*model.User, error)
	}

	// Implementation
	UserRepositoryImpl struct {
		postgresCollection database.PostgresCollection
	}
)

// New Repository User
func NewUserRepository(postgresCollection database.PostgresCollection) UserRepository {
	return UserRepositoryImpl{
		postgresCollection: postgresCollection,
	}
}

func (r UserRepositoryImpl) Create(ctx context.Context, user string) (*model.User, error) {
	var responseUser model.User
	// TODO: add get user
	responseUser.AuthorID = 1
	responseUser.FirstName = &user

	// call db

	return nil, utils.ErrorNotFound
}

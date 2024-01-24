package user

import (
	"context"
	"fmt"
	"time"

	"github.com/erwinwahyura/go-boilerplate/app/database"
	"github.com/erwinwahyura/go-boilerplate/app/model"

	"github.com/erwinwahyura/go-boilerplate/app/repository"
	"github.com/opentracing/opentracing-go"
)

type (
	// UserService service
	UserService interface {
		CreateUser(ctx context.Context, userReq model.UserRequest) (int64, error)
	}

	// UserServiceImpl implementation
	UserServiceImpl struct {
		config          model.Config
		mongoCollection database.MongoCollection
		userRepo        repository.UserRepository
	}
)

// New
func NewService(
	config model.Config,
	mongoCollection database.MongoCollection,
	userRepository repository.UserRepository,
) UserService {
	return UserServiceImpl{
		config:          config,
		mongoCollection: mongoCollection,
		userRepo:        userRepository,
	}
}

// Create user
func (s UserServiceImpl) CreateUser(ctx context.Context, userReq model.UserRequest) (int64, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserServiceImpl.CreateUser")
	defer span.Finish()

	var err error
	defer func(start time.Time, err error) {
		if err != nil {
			span.SetTag("Error", true)
			span.LogKV("ErrorMsg", err.Error())
		}
	}(time.Now(), err)

	var response int64
	// bussiness logic
	// user, err := s.HandleCreateUser(userReq)
	// if err != nil {
	// 	return 0, err
	// }

	// call save user repository
	res, err := s.userRepo.Create(ctx, "user admin")
	if err != nil {
		fmt.Println("error on creating user", err)
		return 0, err
	}

	if res != nil {
		response = res.AuthorID
	}

	return response, nil
}

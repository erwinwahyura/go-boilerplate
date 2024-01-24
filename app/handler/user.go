package handler

import (
	"net/http"

	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/app/service/user"
	"github.com/erwinwahyura/go-boilerplate/utils"

	// "github.com/erwinwahyura/go-boilerplate/utils/jaegerutil"
	"github.com/rs/zerolog/log"
)

type (
	// UserHandler controller
	UserHandler interface {
		CreateUser(w http.ResponseWriter, r *http.Request)
	}

	// UserHandlerImpl health controller
	UserHandlerImpl struct {
		userService user.UserService
	}
)

// NewUserHandler initialize health controller
func NewUserHandler(h user.UserService) UserHandler {
	return &UserHandlerImpl{userService: h}
}

// Check godoc
// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} model.BaseResponse
// @Router /healthcheck [get]
func (h *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	// span, _ := jaegerutil.StartSpan(r.Context(), utils.GetCurrentFunctionName())
	// defer span.Finish()

	var err error
	// defer jaegerutil.SetErrorSpan(span, time.Now(), err)

	data, err := h.userService.CreateUser(r.Context(), model.UserRequest{})
	if err != nil {
		log.Error().Msgf("error when healthService.Check(), err: %v", err)
		model.MapBaseResponse(w, r, err.Error(), data, nil, err)
		return
	}
	model.MapBaseResponse(w, r, utils.Success, data, nil, nil)
}

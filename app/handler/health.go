package handler

import (
	"net/http"

	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/app/service/healthcheck"
	"github.com/erwinwahyura/go-boilerplate/utils"

	// "github.com/erwinwahyura/go-boilerplate/utils/jaegerutil"
	"github.com/rs/zerolog/log"
)

type (
	// HealthHandler controller
	HealthHandler interface {
		Check(w http.ResponseWriter, r *http.Request)
	}

	// HealthHandlerImpl health controller
	HealthHandlerImpl struct {
		healthService healthcheck.HealthService
	}
)

// NewHealthHandler initialize health controller
func NewHealthHandler(h healthcheck.HealthService) HealthHandler {
	return &HealthHandlerImpl{healthService: h}
}

// Check godoc
// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} model.BaseResponse
// @Router /healthcheck [get]
func (h *HealthHandlerImpl) Check(w http.ResponseWriter, r *http.Request) {
	// span, _ := jaegerutil.StartSpan(r.Context(), utils.GetCurrentFunctionName())
	// defer span.Finish()

	var err error
	// defer jaegerutil.SetErrorSpan(span, time.Now(), err)

	data, err := h.healthService.Check(r.Context())
	if err != nil {
		log.Error().Msgf("error when healthService.Check(), err: %v", err)
		model.MapBaseResponse(w, r, err.Error(), data, nil, err)
		return
	}
	model.MapBaseResponse(w, r, utils.Success, data, nil, nil)
}

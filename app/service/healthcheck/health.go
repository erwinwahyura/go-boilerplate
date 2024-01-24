package healthcheck

import (
	"context"
	"errors"
	"time"

	"github.com/erwinwahyura/go-boilerplate/app/database"
	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/hellofresh/health-go/v4"
	"github.com/opentracing/opentracing-go"
)

type (
	// HealthService health service
	HealthService interface {
		Check(ctx context.Context) (model.HealthCheckResponse, error)
	}

	// HealthServiceImpl implementation
	HealthServiceImpl struct {
		config             model.Config
		mongoCollection    database.MongoCollection
		postgresCollection database.PostgresCollection
	}
)

// NewService initialize health usecase
func NewService(
	config model.Config,
	mongoCollection database.MongoCollection,
	postgresCollection database.PostgresCollection,
) HealthService {
	return HealthServiceImpl{
		config:             config,
		mongoCollection:    mongoCollection,
		postgresCollection: postgresCollection,
	}
}

// Check health check
func (uc HealthServiceImpl) Check(ctx context.Context) (model.HealthCheckResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "HealthServiceImpl.Check")
	defer span.Finish()

	var err error
	defer func(start time.Time, err error) {
		if err != nil {
			span.SetTag("Error", true)
			span.LogKV("ErrorMsg", err.Error())
		}
	}(time.Now(), err)

	var response model.HealthCheckResponse
	h, err := health.New()
	if err != nil {
		response.Errors = append(response.Errors, "failed to init healthcheck")
	}
	// custom health check example (fail)
	healthTest := h.Register(health.Config{
		Name:      "service-check",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: func(context.Context) error {
			return errors.New("failed during health check")
		},
	})
	if healthTest != nil {
		response.Errors = append(response.Errors, "error ping service")
		return response, err
	}
	response.Success = append(response.Success, "success ping service")

	// TODO: check db connection

	return response, err
}

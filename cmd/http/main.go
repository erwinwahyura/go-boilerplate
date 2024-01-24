package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	c "github.com/erwinwahyura/go-boilerplate/app/config"
	"github.com/erwinwahyura/go-boilerplate/app/database"
	"github.com/erwinwahyura/go-boilerplate/app/handler"
	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/app/repository"
	"github.com/erwinwahyura/go-boilerplate/app/route"
	"github.com/erwinwahyura/go-boilerplate/app/service/healthcheck"
	"github.com/erwinwahyura/go-boilerplate/app/service/user"
	"github.com/erwinwahyura/go-boilerplate/docs"
	"github.com/labstack/gommon/color"
	"github.com/spf13/viper"
)

// Init initialize config to viper
func LoadConfig(path string) (config model.Config, err error) {
	viper.AutomaticEnv()

	// Check if the .env file exists
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return config, err
		}
	}

	// do viper bind
	c.ViperBind()

	err = viper.Unmarshal(&config)
	return config, err
}

// SetSwaggerInfo swagger
func setSwaggerInfo(config model.Config) {
	docs.SwaggerInfo.Title = "Api"
	docs.SwaggerInfo.Description = "Api server"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
}

func main() {

	// Config
	log.Println("[INFO] Loading environment")
	cfg, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	// reload secret
	c.Reload()

	// DB
	log.Println("[INFO] Loading database")
	mongoCollection := database.NewMongoCollection(cfg)
	postgresCollection := database.NewPostgresCollection(cfg)

	// Repository
	log.Println("[INFO] Loading repository")
	userRepo := repository.NewUserRepository(postgresCollection)

	// Outbound
	log.Println("[INFO] Loading outbound")

	// NSQ Producer
	log.Println("[INFO] Loading nsq producer")

	// Shared Service
	log.Println("[INFO] Loading Shared Service")

	// Service
	log.Println("[INFO] Loading service")
	healthService := healthcheck.NewService(cfg, mongoCollection, postgresCollection)
	userService := user.NewService(cfg, mongoCollection, userRepo)

	// Handler
	log.Println("[INFO] Loading handler")
	healthHandler := handler.NewHealthHandler(healthService)
	userHandler := handler.NewUserHandler(userService)

	// NSQ Consumer
	log.Println("[INFO] Loading nsq consumer")

	// Swagger
	setSwaggerInfo(cfg)

	// Server & Router
	log.Println("[INFO] Loading router")
	router := route.NewRoutes(cfg, healthHandler, userHandler)

	// Server Runner
	log.Println("[INFO] Loading server")
	serverRunner(cfg, router)
}

// var tracer trace.Tracer

// ServerRunner run server
func serverRunner(
	cfg model.Config,
	handler http.Handler,
) {
	// Tracer
	// tracer, closer := jaegerutil.NewTracerJaeger("api-starter", cfg.Jaeger.URL, cfg.Jaeger.Disable)
	// // Set the singleton opentracing.Tracer with the Jaeger tracer.
	// opentracing.SetGlobalTracer(tracer)
	// defer closer.Close()

	// The HTTP Server
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host.Address, cfg.Host.Port),
		WriteTimeout: time.Second * time.Duration(cfg.Host.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.Host.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.Host.IdleTimeout),
		Handler:      handler,
	}

	// Run Server
	go func() {
		color.Printf("⇨ http server started on %s\n", color.Green(server.Addr))
		server.ListenAndServe()
	}()

	// NSQ Consumer

	// Wait to shooting down
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("received shutdown signal. Trying to shutdown gracefully", sig)

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stop Server
	color.Println(color.Red("Stopping HTTP Server"))
	server.SetKeepAlivesEnabled(false)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Failure while shutting down gracefully, errApp: ", err)
	}

	// Stop NSQ Consumer

	log.Println("Shutdown gracefully completed")
}

// func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
// 	// Ensure default SDK resources and the required service name are set.
// 	r, err := resource.Merge(
// 		resource.Default(),
// 		resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceName("APIStarter"),
// 		),
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return sdktrace.NewTracerProvider(
// 		sdktrace.WithBatcher(exp),
// 		sdktrace.WithResource(r),
// 	)
// }

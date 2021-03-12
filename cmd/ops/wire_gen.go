// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"go-template/internal/app/ops"
	"go-template/internal/app/ops/controllers"
	"go-template/internal/app/ops/job"
	"go-template/internal/app/ops/repositories"
	"go-template/internal/app/ops/services"
	"go-template/internal/pkg/app"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/http"
	"go-template/internal/pkg/log"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	opsOptions, err := ops.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	dbOptions, err := db.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	dbDB, err := db.New(dbOptions, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	demoRepository := repositories.NewDemoRepository(dbDB)
	demoService := services.NewDemoService(demoRepository)
	demoController := controllers.NewDemoController(demoService)
	initControllers := controllers.InitControllersFn(demoController)
	engine := http.NewRouter(httpOptions, logger, initControllers)
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	scheduler := job.NewScheduler(logger)
	application, err := ops.NewApp(opsOptions, logger, dbDB, server, scheduler)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, db.ProviderSet, repositories.ProviderSet, services.ProviderSet, controllers.ProviderSet, job.ProviderSet, http.ProviderSet, ops.ProviderSet)

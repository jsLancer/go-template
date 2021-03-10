//+build wireinject

package main

import (
	"github.com/google/wire"
	"go-template/internal/app/ops"
	"go-template/internal/app/ops/controllers"
	schedule "go-template/internal/app/ops/job"
	"go-template/internal/app/ops/repositories"
	"go-template/internal/app/ops/services"
	"go-template/internal/pkg/app"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/http"
	"go-template/internal/pkg/log"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	db.ProviderSet,
	repositories.ProviderSet,
	services.ProviderSet,
	controllers.ProviderSet,
	schedule.ProviderSet,
	http.ProviderSet,
	ops.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}

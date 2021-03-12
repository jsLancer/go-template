package ops

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-template/internal/app/ops/job"
	"go-template/internal/pkg/app"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/http"
	"go.uber.org/zap"
)

// Options define options of agent app
type Options struct {
	Name string
}

// NewOptions define constructor of agent app options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, db *db.DB, http *http.Server,
	scheduler *job.Scheduler) (*app.Application, error) {
	a, err := app.New(o.Name, logger, db, http, scheduler)

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

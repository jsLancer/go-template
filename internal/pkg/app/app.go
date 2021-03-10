package app

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type LifeStyle interface {
	Start() error
	Reload() error
	Stop() error
}

type Application struct {
	name    string
	logger  *zap.Logger
	servers []LifeStyle
}

func New(name string, logger *zap.Logger, servers ...LifeStyle) (*Application, error) {

	return &Application{
		name:    name,
		logger:  logger,
		servers: servers,
	}, nil
}

func (a *Application) Start() error {
	for _, server := range a.servers {
		if err := server.Start(); err != nil {
			a.logger.Warn("start server error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (a *Application) Stop() error {
	for _, server := range a.servers {
		if err := server.Stop(); err != nil {
			a.logger.Warn("stop server error", zap.Error(err))
			return err
		}
	}
	return nil
}

// AwaitSignal graceful shutdown when receive terminal signal
// kill -15
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if err := a.Stop(); err != nil {
			a.logger.Warn("stop server error", zap.Error(err))
		}
		os.Exit(0)
	}
}

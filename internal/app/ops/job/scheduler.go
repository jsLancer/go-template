package schedule

import (
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Job interface {
	cron.Job
	Cron() string
	Name() string
}

type Scheduler struct {
	jobs   map[string]Job
	logger *zap.Logger
	cron   *cron.Cron
}

func NewScheduler(logger *zap.Logger) *Scheduler {
	s := &Scheduler{
		jobs:   map[string]Job{},
		logger: logger,
		cron:   newWithSeconds(),
	}

	//s.Register()

	return s
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func (s *Scheduler) Register(job Job) {
	s.jobs[job.Name()] = job
	_, err := s.cron.AddFunc(job.Cron(), job.Run)
	if err != nil {
		s.logger.Error("register job error "+job.Name(), zap.Error(err))
	} else {
		s.logger.Info("register job success " + job.Name())
	}
}

func (s *Scheduler) Start() error {
	s.cron.Start()
	return nil
}

func (s *Scheduler) Stop() error {
	s.cron.Stop()
	return nil
}

var ProviderSet = wire.NewSet(
	NewScheduler,
)

package mq

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	NameSrv []string `yaml:"namesrv"`
}

type RocketMQ struct {
	o         *Options
	logger    *zap.Logger
	producers map[string]rocketmq.Producer
	consumers map[string]rocketmq.PushConsumer
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("rocketmq", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal rocketmq option error")
	}
	return o, nil
}

func NewRocketMQ(o *Options, logger *zap.Logger) (*RocketMQ, error) {
	return &RocketMQ{
		o:         o,
		logger:    logger,
		producers: make(map[string]rocketmq.Producer, 0),
		consumers: make(map[string]rocketmq.PushConsumer, 0),
	}, nil
}

func (r *RocketMQ) Start() error {
	return nil
}

func (r *RocketMQ) Reload() error {
	return nil
}

func (r *RocketMQ) Stop() error {
	for group, p := range r.producers {
		err := p.Shutdown()
		if err != nil {
			r.logger.Error(fmt.Sprintf("Producer[%s] shutdown error", group), zap.Error(err))
		}
	}
	for group, c := range r.consumers {
		err := c.Shutdown()
		if err != nil {
			r.logger.Error(fmt.Sprintf("Consumer[%s] shutdown error", group), zap.Error(err))
		}
	}
	return nil
}

func (r *RocketMQ) NewPushConsumer(groupName string, opts ...consumer.Option) (rocketmq.PushConsumer, error) {
	if r.consumers[groupName] != nil {
		return r.consumers[groupName], nil
	}

	opts = append(opts, consumer.WithNameServer(r.o.NameSrv))
	opts = append(opts, consumer.WithGroupName(groupName))
	c, err := rocketmq.NewPushConsumer(opts...)
	if err != nil {
		return nil, err
	}
	r.consumers[groupName] = c
	return c, nil
}

func (r *RocketMQ) NewProducer(groupName string, opts ...producer.Option) (rocketmq.Producer, error) {
	if r.producers[groupName] != nil {
		return r.producers[groupName], nil
	}

	opts = append(opts, producer.WithNameServer(r.o.NameSrv))
	opts = append(opts, producer.WithGroupName(groupName))
	p, err := rocketmq.NewProducer(opts...)
	if err != nil {
		return nil, err
	}
	r.producers[groupName] = p
	err = p.Start()
	if err != nil {
		return nil, err
	}
	return p, nil
}

var ProviderSet = wire.NewSet(NewRocketMQ, NewOptions)

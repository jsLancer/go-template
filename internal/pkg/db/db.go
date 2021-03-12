package db

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-template/internal/app/ops/models"
	"go.uber.org/zap"
)

type Options struct {
	URL   string `yaml:"url"`
	Debug bool   `yaml:"debug"`
}

type DB struct {
	o      *Options
	logger *zap.Logger
	DB     *gorm.DB
}

// NewOptions build database option from viper
func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	return o, err
}

func New(o *Options, logger *zap.Logger) (*DB, error) {
	var err error
	db, err := gorm.Open("mysql", o.URL)
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}

	if o.Debug {
		db = db.Debug()
	}

	logger.Info("initialize database success", zap.String("url", o.URL))

	return &DB{
		o:      o,
		logger: logger,
		DB:     db,
	}, nil
}

func (db *DB) Start() error {
	db.DB.AutoMigrate(&models.Demo{})
	return nil
}

func (db *DB) Reload() error {
	return nil
}

func (db *DB) Stop() error {
	err := db.DB.Close()
	if err != nil {
		db.logger.Error("close db error", zap.Error(err))
	}
	return nil
}

// ProviderSet define provider set of db
var ProviderSet = wire.NewSet(New, NewOptions)

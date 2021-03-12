package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-template/internal/app/ops/models"
	"go-template/internal/pkg/db"
)

type DemoRepository interface {
	Create(demo *models.Demo) error
	Update(demo *models.Demo) error
	FindAll() ([]*models.Demo, error)
	DeleteByID(ID uint64) error
	GetByID(ID uint64) (*models.Demo, error)
}

type demoRepository struct {
	db *gorm.DB
}

func NewDemoRepository(db *db.DB) DemoRepository {
	return &demoRepository{
		db: db.DB,
	}
}

func (d *demoRepository) Create(demo *models.Demo) error {
	if err := d.db.Table(models.DemoTableName).Create(demo).Error; err != nil {
		return errors.Wrap(err, "create demo error")
	}
	return nil
}

func (d *demoRepository) Update(demo *models.Demo) error {
	if err := d.db.Table(models.DemoTableName).Save(demo).Error; err != nil {
		return errors.Wrap(err, "update demo error")
	}
	return nil
}

func (d *demoRepository) FindAll() ([]*models.Demo, error) {
	var results []*models.Demo
	if err := d.db.Table(models.DemoTableName).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "findAll demo error")
	}
	return results, nil
}

func (d *demoRepository) GetByID(ID uint64) (*models.Demo, error) {
	demo := new(models.Demo)
	if err := d.db.Table(models.DemoTableName).Where("id = ?", ID).First(&demo).Error; err != nil {
		return nil, errors.Wrap(err, "get by id error")
	}
	return demo, nil
}

func (d *demoRepository) DeleteByID(ID uint64) error {
	if err := d.db.Table(models.DemoTableName).Delete(models.Demo{}, "id = ?", ID).Error; err != nil {
		return errors.Wrap(err, "delete demo error")
	}
	return nil
}

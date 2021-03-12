package services

import (
	"go-template/internal/app/ops/models"
	"go-template/internal/app/ops/repositories"
)

type DemoService interface {
	CreateDemo(demo *models.Demo) error
	GetByID(ID uint64) (*models.Demo, error)
	FindAll() ([]*models.Demo, error)
	Update(demo *models.Demo) error
	DeleteByID(ID uint64) error
}

type demoService struct {
	demoRepo repositories.DemoRepository
}

func NewDemoService(demoRepo repositories.DemoRepository) DemoService {
	return &demoService{
		demoRepo: demoRepo,
	}
}

func (d *demoService) CreateDemo(demo *models.Demo) error {
	return d.demoRepo.Create(demo)
}

func (d *demoService) FindAll() ([]*models.Demo, error) {
	return d.demoRepo.FindAll()
}

func (d *demoService) GetByID(ID uint64) (*models.Demo, error) {
	return d.demoRepo.GetByID(ID)
}

func (d *demoService) Update(demo *models.Demo) error {
	return d.demoRepo.Update(demo)
}

func (d *demoService) DeleteByID(ID uint64) error {
	return d.demoRepo.DeleteByID(ID)
}

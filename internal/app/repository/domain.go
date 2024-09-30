package repository

import (
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type DomainRepo struct {
}

type IDomainRepo interface {
	GetDomainByID(id uint) (*model.Domain, error)
	GetListDomain() ([]model.Domain, error)
	CreateDomain(domain *model.Domain) error
	UpdateDomain(domain *model.Domain) error
	DeleteDomain(id uint) error
	GetDomainByDomain(domain string) (*model.Domain, error)
}

func NewIDomainRepo() IDomainRepo {
	return &DomainRepo{}
}

func (d *DomainRepo) GetDomainByID(id uint) (*model.Domain, error) {
	var domain model.Domain
	db := global.DB
	err := db.Where("id = ?", id).First(&domain).Error
	return &domain, err
}

func (d *DomainRepo) GetListDomain() ([]model.Domain, error) {
	var domain []model.Domain
	db := global.DB
	err := db.Find(&domain).Error
	return domain, err
}

func (d *DomainRepo) CreateDomain(domain *model.Domain) error {
	db := global.DB
	err := db.Create(domain).Error
	return err
}

func (d *DomainRepo) UpdateDomain(domain *model.Domain) error {
	db := global.DB
	err := db.Model(domain).Updates(domain).Error
	return err
}

func (d *DomainRepo) DeleteDomain(id uint) error {
	db := global.DB
	err := db.Delete(&model.Domain{}, id).Error
	return err
}

func (d *DomainRepo) GetDomainByDomain(domain string) (*model.Domain, error) {
	var domainModel model.Domain
	db := global.DB
	err := db.Where("domain = ?", domain).First(&domainModel).Error
	return &domainModel, err
}

package service

import (
	"first-api/models"
	"first-api/store"
)

type Service interface {
	UpdateMobile(id string) (*models.Mobile, error)
	GetAllMobiles() ([]models.Mobile, error)
	GetMobileByID(id string) (*models.Mobile, error)
	DeleteMobile(id string) (bool, error)
	CreateMobile(mobile *models.Mobile) error
}

type service struct {
	store *store.Store
}

func New() *service {
	return &service{
		store: store.GetStore(),
	}
}

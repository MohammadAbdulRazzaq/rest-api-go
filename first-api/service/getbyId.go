package service

import "first-api/models"

func (s *service) GetMobileByID(id string) (*models.Mobile, error) {
	mobile, err := s.store.GetMobileByID(id)
	if err != nil {
		return &models.Mobile{}, err
	}
	return mobile, err
}

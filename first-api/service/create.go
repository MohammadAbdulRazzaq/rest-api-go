package service

import "first-api/models"

func (s *service) CreateMobile(mobile *models.Mobile) error {
	err := s.store.CreateMobile(mobile)

	if err != nil {
		return err
	}
	return nil
}

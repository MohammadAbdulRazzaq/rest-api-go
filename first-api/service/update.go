package service

import "first-api/models"

func (s *service) UpdateMobile(id string) (*models.Mobile, error) {
	mobile, err := s.store.UpdateMobile(id)
	if err != nil {
		return mobile, err
	}
	return mobile, nil
}

package service

import "first-api/models"

func (s *service) GetAllMobiles() ([]models.Mobile, error) {
	mobile, err := s.store.GetAllMobiles()
	if err != nil {
		return nil, err
	}
	return mobile, nil
}

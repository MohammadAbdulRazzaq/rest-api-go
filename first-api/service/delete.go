package service

import "fmt"

func (s *service) DeleteMobile(id string) (bool, error) {
	success, err := s.store.DeleteMobile(id)
	if err != nil {
		return false, err
	}

	if !success {
		return false, fmt.Errorf("failed to delete mobile with ID %s", id)
	}

	return true, nil
}

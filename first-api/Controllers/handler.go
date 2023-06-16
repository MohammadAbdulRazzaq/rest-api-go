package Controllers

import (
	"first-api/service"
)

type Handler struct {
	Service service.Service
}

func New() *Handler {
	return &Handler{
		Service: service.New(),
	}
}

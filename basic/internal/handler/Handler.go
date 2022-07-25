package handler

import "github.com/mrbelka12000/netfix/basic/internal/service"

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{srv}
}

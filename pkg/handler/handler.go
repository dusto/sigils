package handler

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/pkg/repository"
)

type Handler struct {
	api      huma.API
	configDB *repository.Queries
}

func New(api huma.API, configDB *repository.Queries) *Handler {
	return &Handler{
		api:      api,
		configDB: configDB,
	}
}

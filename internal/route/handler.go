package route

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/repository"
)

type Handler struct {
	api      huma.API
	configDB *repository.Queries
}

func NewHandler(api huma.API, configDB *repository.Queries) *Handler {
	return &Handler{
		api:      api,
		configDB: configDB,
	}
}

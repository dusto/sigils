package route

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/repository"
	"github.com/go-chi/httplog/v2"
)

type Handler struct {
	api    huma.API
	rawDB  *repository.MultiSqliteDB
	query  *repository.Queries
	logger *httplog.Logger
}

func NewHandler(api huma.API, rawDB *repository.MultiSqliteDB, query *repository.Queries, logger *httplog.Logger) *Handler {
	return &Handler{
		api:    api,
		rawDB:  rawDB,
		query:  query,
		logger: logger,
	}
}

package handler

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func (h *Handler) Register() {

	// config endpoint
	huma.Register(h.api, huma.Operation{
		OperationID: "list-configs",
		Method:      http.MethodGet,
		Path:        "/clusters",
		Summary:     "List Machine Clusters",
	}, h.ClusterGetAnyOf)
	huma.Register(h.api, huma.Operation{
		OperationID: "get-configs",
		Method:      http.MethodGet,
		Path:        "/cluster/{id}",
		Summary:     "Get Machine Cluster",
	}, h.ClusterGetOneOf)
	huma.Register(h.api, huma.Operation{
		OperationID:   "post-configs",
		Method:        http.MethodPost,
		Path:          "/cluster/add",
		Summary:       "Add new Machine Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterPost)
	huma.Register(h.api, huma.Operation{
		OperationID: "delete-configs",
		Method:      http.MethodDelete,
		Path:        "/cluster/{id}",
		Summary:     "Delete Machine Cluster",
	}, h.ClusterDelete)
	huma.Register(h.api, huma.Operation{
		OperationID:   "gen-configs",
		Method:        http.MethodPost,
		Path:          "/cluster/gen",
		Summary:       "Generate new Machine Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterGen)
}

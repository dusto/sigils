package route

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
		Summary:     "List Clusters",
	}, h.ClusterGetAnyOf)
	huma.Register(h.api, huma.Operation{
		OperationID: "get-configs",
		Method:      http.MethodGet,
		Path:        "/clusters/{id}",
		Summary:     "Get Cluster",
	}, h.ClusterGetOneOf)
	huma.Register(h.api, huma.Operation{
		OperationID:   "post-configs",
		Method:        http.MethodPost,
		Path:          "/clusters",
		Summary:       "Import/Manually Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterPost)
	huma.Register(h.api, huma.Operation{
		OperationID: "delete-configs",
		Method:      http.MethodDelete,
		Path:        "/clusters/{id}",
		Summary:     "Delete Cluster",
	}, h.ClusterDelete)
	huma.Register(h.api, huma.Operation{
		OperationID:   "gen-configs",
		Method:        http.MethodPost,
		Path:          "/clusters/generate",
		Summary:       "Automatically generate new Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterGen)
}

package route

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/autopatch"
)

func (h *Handler) Register() {

	// cluster endpoint
	huma.Register(h.api, huma.Operation{
		OperationID: "list-clusters",
		Method:      http.MethodGet,
		Path:        "/clusters",
		Summary:     "List Clusters",
	}, h.ClusterGetAnyOf)
	huma.Register(h.api, huma.Operation{
		OperationID: "get-cluster",
		Method:      http.MethodGet,
		Path:        "/clusters/{uuid}",
		Summary:     "Get Cluster",
	}, h.ClusterGetOneOf)
	huma.Register(h.api, huma.Operation{
		OperationID:   "post-cluster",
		Method:        http.MethodPost,
		Path:          "/clusters",
		Summary:       "Import/Manually Add Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterPost)
	huma.Register(h.api, huma.Operation{
		OperationID: "delete-cluster",
		Method:      http.MethodDelete,
		Path:        "/clusters/{uuid}",
		Summary:     "Delete Cluster",
	}, h.ClusterDelete)
	huma.Register(h.api, huma.Operation{
		OperationID:   "gen-cluster",
		Method:        http.MethodPost,
		Path:          "/clusters/generate",
		Summary:       "Automatically generate new Cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterGen)
	huma.Register(h.api, huma.Operation{
		OperationID:   "attach-host",
		Method:        http.MethodPost,
		Path:          "/clusters/{cluster_uuid}/attach/{host_uuid}",
		Summary:       "Attach/Add host to cluster",
		DefaultStatus: http.StatusCreated,
	}, h.ClusterAttachHost)

	// hosts endpoint
	huma.Register(h.api, huma.Operation{
		OperationID: "list-hosts",
		Method:      http.MethodGet,
		Path:        "/hosts",
		Summary:     "List hosts",
	}, h.HostGetAnyOf)
	huma.Register(h.api, huma.Operation{
		OperationID: "get-host",
		Method:      http.MethodGet,
		Path:        "/hosts/{uuid}",
		Summary:     "Get Host",
	}, h.HostGetOneOf)
	huma.Register(h.api, huma.Operation{
		OperationID:   "post-hosts",
		Method:        http.MethodPost,
		Path:          "/hosts",
		Summary:       "Add Hosts",
		DefaultStatus: http.StatusCreated,
	}, h.HostPost)
	huma.Register(h.api, huma.Operation{
		OperationID: "delete-host",
		Method:      http.MethodDelete,
		Path:        "/hosts/{uuid}",
		Summary:     "Delete Host",
	}, h.HostDelete)
	huma.Register(h.api, huma.Operation{
		OperationID:   "attach-profile",
		Method:        http.MethodPost,
		Path:          "/hosts/{host_uuid}/attach/{profile_name}",
		Summary:       "Attach/Add profile to host",
		DefaultStatus: http.StatusCreated,
	}, h.HostAttachProfile)

	// profile endpoint
	huma.Register(h.api, huma.Operation{
		OperationID: "list-profiles",
		Method:      http.MethodGet,
		Path:        "/profiles",
		Summary:     "List profiles",
	}, h.ProfileGetAnyOf)
	huma.Register(h.api, huma.Operation{
		OperationID: "get-profile",
		Method:      http.MethodGet,
		Path:        "/profiles/{id}",
		Summary:     "Get Profile",
	}, h.ProfileGetOneOf)
	huma.Register(h.api, huma.Operation{
		OperationID:   "post-profiles",
		Method:        http.MethodPost,
		Path:          "/profiles",
		Summary:       "Add Profiles",
		DefaultStatus: http.StatusCreated,
	}, h.ProfilePost)
	huma.Register(h.api, huma.Operation{
		OperationID: "delete-profile",
		Method:      http.MethodDelete,
		Path:        "/profiles/{id}",
		Summary:     "Delete Profile",
	}, h.ProfileDelete)

	// machineconfig endpoint
	huma.Register(h.api, huma.Operation{
		OperationID:   "machineconfig",
		Method:        http.MethodGet,
		Path:          "/machineconfig",
		Summary:       "Get a patched machineconfig for a specific host",
		DefaultStatus: http.StatusOK,
	}, h.GetMachineConfig)

	autopatch.AutoPatch(h.api)
}

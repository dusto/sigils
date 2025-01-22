package route

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/pkg/repository"
)

type ClusterOutput struct {
	Body []ClusterBody
}

type ClusterListInput struct {
}

type ClusterAnyOfInput struct {
	Search string `query:"search" required:"false"`
}

type ClusterOneOfInput struct {
	Id int64 `path:"id" required:"false"`
}

type ClusterBody struct {
	Id       int64  `json:"id" doc:"Cluster ID" required:"true"`
	Name     string `json:"name" doc:"Cluster Name" required:"true"`
	Endpoint string `json:"endpoint" doc:"Cluster Endpoint" required:"true"`
}

type ClusterPostInput struct {
	Body []ClusterBody
}

type ClusterGen struct {
	Name     string `json:"name" doc:"Cluster Name" required:"true"`
	Endpoint string `json:"endpoint" doc:"Cluster Endpoint" required:"true"`
}

type ClusterGenPostInput struct {
	Body []ClusterGen
}

func (h *Handler) ClusterGetOneOf(ctx context.Context, input *ClusterOneOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	cluster, err := h.configDB.GetClusterById(ctx, input.Id)
	if err != nil {
		return resp, nil
	}

	cBody := ClusterBody{
		Id:       cluster.ID,
		Name:     cluster.Name,
		Endpoint: cluster.Endpoint,
	}

	resp.Body = append(resp.Body, cBody)

	return resp, nil
}

func (h *Handler) ClusterGetAnyOf(ctx context.Context, input *ClusterAnyOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	var clusters []repository.Cluster
	var err error

	if input.Search != "" {
		clusters, err = h.configDB.ListClustersSearch(ctx, input.Search)
	} else {
		clusters, err = h.configDB.ListClusters(ctx)
	}

	if err != nil {
		return resp, nil
	}
	for _, cluster := range clusters {
		resp.Body = append(resp.Body, ClusterBody{
			Id:       cluster.ID,
			Name:     cluster.Name,
			Endpoint: cluster.Endpoint,
		})

	}

	return resp, nil
}

func (h *Handler) ClusterPost(ctx context.Context, input *ClusterPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	for _, cIn := range input.Body {
		_, err := h.configDB.InsertCluster(ctx, repository.InsertClusterParams{
			Name:     cIn.Name,
			Endpoint: cIn.Endpoint,
		})
		if err != nil {
			return resp, huma.Error500InternalServerError("Could not add cluster", err)
		}
	}

	return resp, nil
}

func (h *Handler) ClusterDelete(ctx context.Context, input *ClusterOneOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	err := h.configDB.DeleteCluster(ctx, input.Id)

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not remove cluster", err)
	}

	return resp, nil
}

func (h *Handler) ClusterGen(ctx context.Context, input *ClusterGenPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	return resp, nil
}

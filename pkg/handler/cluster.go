package handler

import (
	"context"
)

type ClusterOutput struct {
	Body []struct{}
}

type ClusterListInput struct {
}

type ClusterAnyOfInput struct {
	Name string `query:"name" required:"false"`
	Id   string `query:"id" required:"false"`
}

type ClusterOneOfInput struct {
	Id string `path:"id"`
}

type ClusterBody struct {
	Name   string `json:"name" doc:"Base MachineConfig Name" required:"true"`
	Config string `json:"config" doc:"Base MachineConfig" required:"true"`
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
	return resp, nil
}

func (h *Handler) ClusterGetAnyOf(ctx context.Context, input *ClusterAnyOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	return resp, nil
}

func (h *Handler) ClusterPost(ctx context.Context, input *ClusterPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	return resp, nil
}

func (h *Handler) ClusterDelete(ctx context.Context, input *ClusterAnyOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	return resp, nil
}

func (h *Handler) ClusterGen(ctx context.Context, input *ClusterGenPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	return resp, nil
}

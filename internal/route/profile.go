package route

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/model"
)

type ProfileOneOfInput struct {
	Id int64 `path:"id" doc:"Profile ID" required:"true"`
}

type ProfileAnyOfInput struct {
}

type ProfileOutput struct {
	Body []model.Profile
}

type ProfilePostInput struct {
}

func (h *Handler) ProfileGetOneOf(ctx context.Context, input *ProfileOneOfInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}

	host, err := h.configDB.GetProfile(ctx, input.Id)

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not find host", err)
	}

	resp.Body = append(resp.Body, host)

	return resp, nil
}

func (h *Handler) ProfileGetAnyOf(ctx context.Context, input *ProfileAnyOfInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}
	return resp, nil
}

func (h *Handler) ProfilePost(ctx context.Context, input *ProfilePostInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}
	return resp, nil
}

func (h *Handler) ProfileDelete(ctx context.Context, input *ProfilePostInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}
	return resp, nil
}

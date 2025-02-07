package route

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/model"
	"github.com/dusto/sigils/internal/repository"
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
	Body []model.Profile
}

type ProfilePatchDelete struct {
	ProfileOneOfInput
	PatchId int64 `query:"patch_id" doc:"Patch ID to remove a patch from a profile" required:"false"`
}

func (h *Handler) ProfileGetOneOf(ctx context.Context, input *ProfileOneOfInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}

	host, err := h.query.GetProfile(ctx, input.Id)

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not find host", err)
	}

	resp.Body = append(resp.Body, host)

	return resp, nil
}

func (h *Handler) ProfileGetAnyOf(ctx context.Context, input *ProfileAnyOfInput) (*ProfileOutput, error) {
	resp := &ProfileOutput{}
	var err error

	resp.Body, err = h.query.GetProfiles(ctx)

	if err != nil {
		return nil, huma.Error500InternalServerError("Could not get list of hosts", err)
	}

	return resp, nil
}

func (h *Handler) ProfilePost(ctx context.Context, input *ProfilePostInput) (*struct{}, error) {
	for _, profile := range input.Body {
		tx, err := h.rawDB.BeginWriteTx(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		queryTx := h.query.WithTx(tx)

		pid, err := queryTx.InsertProfile(ctx, profile.Name)
		if err != nil {
			return nil, huma.Error500InternalServerError("Could not insert profile", err)
		}
		for _, patch := range profile.Patches {
			err := queryTx.InsertPatch(ctx, repository.InsertPatchParams{
				ProfileID: pid,
				NodeType:  patch.NodeType,
				Fqdn:      patch.Fqdn,
				Patch:     patch.Patch,
			})
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not save patch", err)
			}
		}
		err = tx.Commit()
		if err != nil {
			return nil, huma.Error500InternalServerError("Could not commit profile", err)
		}
	}

	return nil, nil
}

func (h *Handler) ProfileDelete(ctx context.Context, input *ProfilePatchDelete) (*struct{}, error) {
	if input.PatchId > 0 {
		return nil, huma.Error501NotImplemented("Deleting of specific patches is not implemented")
	}

	if err := h.query.DeleteProfile(ctx, input.Id); err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("Could not delete profile %i", input.Id), err)
	}

	return nil, nil
}

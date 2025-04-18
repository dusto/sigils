package route

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/model"
	"github.com/dusto/sigils/internal/repository"
	"github.com/google/uuid"
)

type HostOneOfInput struct {
	Uuid uuid.UUID `path:"uuid" doc:"Host ID" format:"uuid" required:"true"`
}

type HostAttachProfileInput struct {
	HostUUID    uuid.UUID `path:"host_uuid" doc:"Host ID" format:"uuid" required:"true"`
	ProfileName string    `path:"profile_name" doc:"Profile Name" required:"false"`
}

type HostAnyOfInput struct {
	Search string `query:"search" required:"false"`
}

type HostOutput struct {
	Body []model.Host
}

type HostInput struct {
	Uuid     string   `json:"uuid" format:"uuid" doc:"Host SMBIOS UUID"`
	Mac      string   `json:"mac,omitempty" doc:"Host Mac Address" required:"false"`
	Fqdn     string   `json:"fqdn" format:"hostname" doc:"Host FQDN"`
	NodeType string   `json:"nodetype" doc:"Host Node Type" enum:"controlplane,worker"`
	Profiles []string `json:"profiles,omitempty" doc:"List of Profile Names to associate with Host" default:""`
}

type HostPostInput struct {
	Body []HostInput
}

func (h *Handler) HostGetOneOf(ctx context.Context, input *HostOneOfInput) (*HostOutput, error) {
	resp := &HostOutput{}

	host, err := h.query.GetHost(ctx, input.Uuid)

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not find host", err)
	}

	resp.Body = append(resp.Body, host)

	return resp, nil
}

func (h *Handler) HostGetAnyOf(ctx context.Context, input *HostAnyOfInput) (*HostOutput, error) {
	resp := &HostOutput{}
	hosts, err := h.query.GetHosts(ctx)
	if err != nil {
		return resp, huma.Error500InternalServerError("Could get hosts", err)
	}
	resp.Body = hosts
	return resp, nil
}

func (h *Handler) HostPost(ctx context.Context, input *HostPostInput) (*struct{}, error) {
	for _, inHost := range input.Body {
		hInsert := repository.InsertHostParams{
			Uuid:     uuid.MustParse(inHost.Uuid),
			Mac:      []byte(inHost.Mac),
			Fqdn:     inHost.Fqdn,
			Nodetype: inHost.NodeType,
		}

		tx, err := h.rawDB.BeginWriteTx(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		queryTx := h.query.WithTx(tx)

		err = queryTx.InsertHost(ctx, hInsert)
		if err != nil {
			return nil, huma.Error500InternalServerError("Could not add host", err)
		}

		if len(inHost.Profiles) != 0 {
			for _, pName := range inHost.Profiles {
				pid, err := h.query.GetProfileId(ctx, pName)
				if err != nil {
					return nil, huma.Error422UnprocessableEntity("Profile not found", err)
				}
				err = queryTx.AttachHostProfile(ctx, repository.AttachHostProfileParams{
					HostUuid:  uuid.MustParse(inHost.Uuid),
					ProfileID: pid,
				})
				if err != nil {
					return nil, huma.Error500InternalServerError("Could not save profile host relation", err)
				}
			}

		}

		err = tx.Commit()
		if err != nil {
			return nil, huma.Error500InternalServerError("Could not commit host to database", err)
		}

	}
	return nil, nil
}

func (h *Handler) HostDelete(ctx context.Context, input *HostOneOfInput) (*struct{}, error) {

	if err := h.query.DeleteHost(ctx, input.Uuid); err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("Could not delete Host %s", input.Uuid), err)
	}

	return nil, nil
}

func (h *Handler) HostAttachProfile(ctx context.Context, input *HostAttachProfileInput) (*struct{}, error) {

	var pId int64
	var err error

	if input.ProfileName == "" {
		return nil, huma.Error422UnprocessableEntity("Profile not passed")
	}
	pId, err = h.query.GetProfileId(ctx, input.ProfileName)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Profile not found", err)
	}
	err = h.query.AttachHostProfile(ctx, repository.AttachHostProfileParams{
		HostUuid:  input.HostUUID,
		ProfileID: pId,
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save host profile relation", err)
	}
	return nil, nil
}

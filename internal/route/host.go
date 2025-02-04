package route

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/model"
	"github.com/dusto/sigils/internal/repository"
	"github.com/google/uuid"
)

type HostOneOfInput struct {
	Uuid string `path:"uuid" doc:"Host ID" format:"uuid" required:"true"`
}

type HostAnyOfInput struct {
	Search string `query:"search" required:"false"`
}

type HostOutput struct {
	Body []model.Host
}

type HostInput struct {
	Uuid       string  `json:"uuid" format:"uuid" doc:"Host SMBIOS UUID"`
	Fqdn       string  `json:"fqdn" format:"uri" doc:"Host FQDN"`
	NodeType   string  `json:"nodetype" doc:"Host Node Type" enum:"controlplane,worker"`
	ProfileIds []int64 `json:"profileids,omitempty" doc:"List of Profile Ids to associate with Host" default:""`
}
type HostPostInput struct {
	Body []HostInput
}

func (h *Handler) HostGetOneOf(ctx context.Context, input *HostOneOfInput) (*HostOutput, error) {
	resp := &HostOutput{}

	host, err := h.configDB.GetHost(ctx, uuid.MustParse(input.Uuid))

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not find host", err)
	}

	resp.Body = append(resp.Body, host)

	return resp, nil
}

func (h *Handler) HostGetAnyOf(ctx context.Context, input *HostAnyOfInput) (*HostOutput, error) {
	resp := &HostOutput{}
	hosts, err := h.configDB.GetHosts(ctx)
	if err != nil {
		return resp, huma.Error500InternalServerError("Could get hosts", err)
	}
	resp.Body = hosts
	return resp, nil
}

func (h *Handler) HostPost(ctx context.Context, input *HostPostInput) (*HostOutput, error) {
	resp := &HostOutput{}
	for _, inHost := range input.Body {
		hInsert := repository.InsertHostParams{
			Uuid:     uuid.MustParse(inHost.Uuid),
			Fqdn:     inHost.Fqdn,
			NodeType: inHost.NodeType,
		}
		pInsert := []repository.AttachHostProfileParams{}

		if len(inHost.ProfileIds) > 0 {
			for _, pid := range inHost.ProfileIds {
				if _, err := h.configDB.GetProfile(ctx, pid); err != nil {
					return resp, huma.Error422UnprocessableEntity("Profile ID not found", err)
				}
				pInsert = append(pInsert, repository.AttachHostProfileParams{
					HostUuid:  uuid.MustParse(inHost.Uuid),
					ProfileID: pid,
				})
			}

		}
		h.configDB.InsertHost(ctx, hInsert)
		if len(pInsert) > 0 {
			for _, pIn := range pInsert {
				h.configDB.AttachHostProfile(ctx, pIn)
			}

		}
	}
	return resp, nil
}

func (h *Handler) HostDelete(ctx context.Context, input *HostPostInput) (*HostOutput, error) {
	resp := &HostOutput{}
	return resp, nil
}

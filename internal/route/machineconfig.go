package route

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/repository"
	"github.com/dusto/sigils/internal/talosconfig"
	"github.com/google/uuid"
	"github.com/siderolabs/talos/pkg/machinery/config/configloader"
	"github.com/siderolabs/talos/pkg/machinery/config/configpatcher"
)

type MachineConfigInput struct {
	UUID string `query:"uuid,omitempty" format:"uuid"`
	MAC  string `query:"mac,omitempty" format:"mac"`
	FQDN string `query:"fqdn,omitempty" format:"hostname"`
}

type MachineConfigOutput struct {
	ContentType string `header:"Content-Type"`
	Body        []byte
}

func (h *Handler) GetMachineConfig(ctx context.Context, input *MachineConfigInput) (*MachineConfigOutput, error) {
	mcOut := &MachineConfigOutput{}
	mcOut.ContentType = "application/json"
	if (MachineConfigInput{}) == *input {
		return nil, huma.Error400BadRequest("No query parameters passed")
	}

	mc, err := h.query.GetMachineConfig(ctx, uuid.MustParse(input.UUID), input.MAC, input.FQDN)
	if err != nil {
		if h.opts.AutoAdd {
			err := h.query.InsertHost(ctx, repository.InsertHostParams{
				Uuid:     uuid.MustParse(input.UUID),
				Mac:      []byte(input.MAC),
				Fqdn:     input.FQDN,
				Nodetype: string(talosconfig.NodeTypeNoDef),
			})
			if err != nil {
				h.logger.Error("Host exists or failed to auto add", "request", input, "error", err)
			}
		}
		return mcOut, huma.Error500InternalServerError("Failed to lookup host", err)
	}

	config, err := configloader.NewFromBytes([]byte(mc.MachineConfig))

	if err != nil {
		return mcOut, huma.Error500InternalServerError("Could not parse Base Machine config", err)
	}

	configInput := configpatcher.WithConfig(config)
	var patchList []string
	for _, profile := range mc.Profiles {
		for _, patch := range profile.Patches {
			patchList = append(patchList, patch.Patch)
		}
	}
	listPatches, err := configpatcher.LoadPatches(patchList)

	patchOut, err := configpatcher.Apply(configInput, listPatches)
	if err != nil {
		return mcOut, huma.Error500InternalServerError("Could not patch configs", err)
	}
	rawOut, err := patchOut.Bytes()
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not get raw machine config", err)
	}
	mcOut.Body = rawOut

	mcOut.ContentType = "application/yaml"
	return mcOut, nil
}

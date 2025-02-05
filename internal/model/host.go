package model

import (
	"net"

	"github.com/danielgtaylor/huma/v2"
)

type Host struct {
	Uuid        string    `json:"uuid" format:"uuid" doc:"Host SMBIOS UUID"`
	Mac         string    `json:"mac,omitempty" doc:"Host Mac Address" required:"false"`
	Fqdn        string    `json:"fqdn" doc:"Host FQDN" format:"hostname"`
	NodeType    string    `json:"nodetype" doc:"Host Node Type" enum:"controlplane,worker"`
	ClusterName string    `json:"clustername,omitempty" doc:"Name of cluster associated with host"`
	Profiles    []Profile `json:"profiles,omitempty" doc:"List of Patches for host"`
}

func (h *Host) Resolve(ctx huma.Context, prefix *huma.PathBuffer) []error {

	if _, err := net.ParseMAC(h.Mac); err != nil {

		return []error{&huma.ErrorDetail{
			Message:  "Invalid Mac",
			Location: prefix.With("Mac"),
			Value:    err.Error(),
		}}
	}
	return nil
}

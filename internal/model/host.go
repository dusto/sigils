package model

type Host struct {
	Uuid     string    `json:"uuid" format:"uuid" doc:"Host SMBIOS UUID"`
	Fqdn     string    `json:"fqdn" doc:"Host FQDN" format:"hostname"`
	NodeType string    `json:"nodetype" doc:"Host Node Type" enum:"controlplane,worker"`
	Profiles []Profile `json:"profiles" doc:"List of Patches for host"`
}

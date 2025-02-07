package model

type MachineConfig struct {
	Host
	MachineConfig string `json:"machineconfig" doc:"Base config for node type"`
}

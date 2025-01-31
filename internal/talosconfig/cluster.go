package talosconfig

import (
	clientconfig "github.com/siderolabs/talos/pkg/machinery/client/config"
	"github.com/siderolabs/talos/pkg/machinery/config"
	"github.com/siderolabs/talos/pkg/machinery/config/generate"
	"github.com/siderolabs/talos/pkg/machinery/config/machine"
)

type Cluster struct {
	ControlPlaneConfig config.Provider
	WorkerConfig       config.Provider
	TalosCtlConfig     *clientconfig.Config
}

// Generate New Cluster config
func NewCluster(name string, endpoint string, kubeversion string, talosversion string) (*Cluster, error) {
	cluster := &Cluster{}
	versioncontract, err := config.ParseContractFromVersion(talosversion)
	if err != nil {
		return nil, err
	}
	input, err := generate.NewInput(
		name,
		endpoint,
		kubeversion,
		generate.WithKubePrismPort(0),
		generate.WithHostDNSForwardKubeDNSToHost(true),
		generate.WithVersionContract(versioncontract))

	if err != nil {
		return nil, err
	}

	cluster.ControlPlaneConfig, err = input.Config(machine.TypeControlPlane)

	if err != nil {
		return nil, err
	}

	cluster.WorkerConfig, err = input.Config(machine.TypeWorker)

	if err != nil {
		return nil, err
	}

	cluster.TalosCtlConfig, err = input.Talosconfig()

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

package talosconfig

type ConfigType string

const (
	ConfigTypeControlPlane ConfigType = "controlplane"
	ConfigTypeWorker       ConfigType = "worker"
	ConfigTypeTalosctl     ConfigType = "talosctl"
)

type NodeType string

const (
	NodeTypeAll          NodeType = "all"
	NodeTypeControlPlane NodeType = "controlplane"
	NodeTypeWorker       NodeType = "worker"
	NodeTypeNoDef        NodeType = "nodef"
)

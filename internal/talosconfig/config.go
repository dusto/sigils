package talosconfig

const (
	_ = iota
	ConfigTypeControlPlane
	ConfigTypeWorker
	ConfigTypeTalosctl
)

const (
	_ = iota
	NodeTypeControlPlane
	NodeTypeWorker
)

package talosconfig

import (
	"errors"
	"fmt"
)

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

func ConfigTypeToInt(in string) (int64, error) {
	// Returns 0 if not matching type
	var ctype int64
	switch in {
	case "controlplane":
		ctype = ConfigTypeControlPlane
	case "worker":
		ctype = ConfigTypeWorker
	case "talosctl":
		ctype = ConfigTypeTalosctl
	default:
		return 0, errors.New(fmt.Sprintf("Invalid Config Type %s", in))

	}
	return ctype, nil
}

func ConfigTypeToString(in int64) (string, error) {

	var ctype string
	switch in {
	case ConfigTypeControlPlane:
		ctype = "controlplane"
	case ConfigTypeWorker:
		ctype = "worker"
	case ConfigTypeTalosctl:
		ctype = "talosctl"
	default:
		return "", errors.New(fmt.Sprintf("Invalid Conifg Type int %i", in))
	}

	return ctype, nil
}

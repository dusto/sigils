package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/dusto/sigils/internal/talosconfig"
)

// Shared models between API and Repo

type Cluster struct {
	Uuid     string          `json:"uuid,omitempty" format:"uuid" doc:"Cluster ID" required:"false"`
	Name     string          `json:"name" doc:"Cluster Name" required:"true"`
	Endpoint string          `json:"endpoint" doc:"Cluster Endpoint" required:"true"`
	Configs  []ClusterConfig `json:"configs,omitempty" doc:"Cluster Configs" require:"false"`
}

type ClusterConfig struct {
	ID          int64                  `json:"id,omitempty" doc:"ID of Config" required:"false"`
	ClusterUuid string                 `json:"uuid,omitempty" format:"uuid" doc:"Cluster ID" required:"false"`
	ConfigType  talosconfig.ConfigType `json:"configtype" doc:"Config type controlplane,worker,talosctl" enum:"controlplane,worker,talosctl"`
	Config      string                 `json:"config" doc:"Yaml representation of the config"`
}

// Intermediate type to parse json from sql result
type CConfigType struct {
	Configs []ClusterConfig
}

func (cc *CConfigType) Scan(value interface{}) error {
	// When getting from the Database we need to parse a json string

	var valBytes []byte

	switch v := value.(type) {
	case []byte:
		valBytes = v
	case string:
		valBytes = []byte(v)

	}

	err := json.Unmarshal(valBytes, &cc.Configs)
	if err != nil {
		return err
	}

	return nil
}

func (cc *ClusterConfig) Value() (driver.Value, error) {
	return json.Marshal(cc)
}

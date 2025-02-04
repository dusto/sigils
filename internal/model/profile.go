package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
	"github.com/siderolabs/talos/pkg/machinery/config/configpatcher"
)

type Profile struct {
	Id      int64   `json:"id,omitempty" doc:"Profile ID"`
	Name    string  `json:"name" doc:"Profile Name"`
	Patches []Patch `json:"patches" doc:"Collection of patches associated with profile"`
}

type Patch struct {
	Id       int64  `json:"id,omitempty" doc:"Patch ID"`
	NodeType string `json:"nodetype" doc:"Type of node for patch to apply to" enum:"all,controlplane,worker"`
	Host     string `json:"host" doc:"Host FQDN/UUID of specific host for patch to apply" default:""`
	Patch    string `json:"patch" doc:"JSON6902 patch or Strategic Merge patch"`
}

func (p *Patch) Resolve(ctx huma.Context, prefix *huma.PathBuffer) []error {

	if _, err := configpatcher.LoadPatch([]byte(p.Patch)); err != nil {

		return []error{&huma.ErrorDetail{
			Message:  "Invalid Patch",
			Location: prefix.With("Patch"),
			Value:    err,
		}}
	}
	return nil
}

// Intermediate type to parse json from sql result
type CProfileType struct {
	Profiles []Profile
}

func (cp *CProfileType) Scan(value interface{}) error {
	// When getting from the Database we need to parse a json string

	var valBytes []byte

	switch v := value.(type) {
	case []byte:
		valBytes = v
	case string:
		valBytes = []byte(v)

	}

	err := json.Unmarshal(valBytes, &cp.Profiles)
	if err != nil {
		return err
	}

	return nil
}

func (p *Profile) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Intermediate type to parse json from sql result
type CPatchType struct {
	Patches []Patch
}

func (cp *CPatchType) Scan(value interface{}) error {
	// When getting from the Database we need to parse a json string

	var valBytes []byte

	switch v := value.(type) {
	case []byte:
		valBytes = v
	case string:
		valBytes = []byte(v)

	}

	err := json.Unmarshal(valBytes, &cp.Patches)
	if err != nil {
		return err
	}

	return nil
}

func (p *Patch) Value() (driver.Value, error) {
	return json.Marshal(p)
}

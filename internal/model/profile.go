package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
	"github.com/siderolabs/talos/pkg/machinery/config/configpatcher"
)

type Profile struct {
	Id      int64   `json:"id,omitempty" doc:"Profile ID"`
	Name    string  `json:"name,omitempty" doc:"Profile Name"`
	Patches []Patch `json:"patches,omitempty" doc:"Collection of patches associated with profile"`
}

type Patch struct {
	Id       int64  `json:"id,omitempty" doc:"Patch ID"`
	NodeType string `json:"nodetype,omitempty" doc:"Type of node for patch to apply to" enum:"all,controlplane,worker" default:"all"`
	Fqdn     string `json:"fqdn,omitempty" doc:"Host FQDN/UUID of specific host for patch to apply" format:"hostname"`
	Patch    string `json:"patch,omitempty" doc:"JSON6902 patch or Strategic Merge patch"`
}

func (p *Patch) Resolve(ctx huma.Context, prefix *huma.PathBuffer) []error {

	if _, err := configpatcher.LoadPatch([]byte(p.Patch)); err != nil {

		return []error{&huma.ErrorDetail{
			Message:  "Invalid Patch",
			Location: prefix.With("Patch"),
			Value:    err.Error(),
		}}
	}
	return nil
}

// Intermediate type to parse json from sql result
type CProfileType struct {
	Profiles []Profile
	Valid    bool
}

func (cp *CProfileType) Scan(value interface{}) error {
	// When getting from the Database we need to parse a json string

	if value == nil {
		cp.Profiles, cp.Valid = []Profile{}, false
		return nil
	}

	var valBytes []byte

	switch v := value.(type) {
	case []byte:
		valBytes = v
	case string:
		valBytes = []byte(v)
	}

	var profiles []Profile
	err := json.Unmarshal(valBytes, &profiles)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		if profile.Id != 0 {
			cp.Profiles = append(cp.Profiles, profile)
		}
	}

	return nil
}

func (cp *CProfileType) Value() (driver.Value, error) {
	if !cp.Valid {
		return []Profile{}, nil
	}
	return json.Marshal(cp)
}

func (p *Profile) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Intermediate type to parse json from sql result
type CPatchType struct {
	Patches []Patch
	Valid   bool
}

func (cp *CPatchType) Scan(value interface{}) error {
	// When getting from the Database we need to parse a json string

	if value == nil {
		cp.Patches, cp.Valid = []Patch{}, false
		return nil
	}

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

func (cp *CPatchType) Value() (driver.Value, error) {
	if !cp.Valid {
		return []Patch{}, nil
	}
	return json.Marshal(cp)
}

func (p *Patch) Value() (driver.Value, error) {
	return json.Marshal(p)
}

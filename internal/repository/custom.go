package repository

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type Config struct {
	ConfigType int64  `json:"configtype"`
	Config     string `json:"config"`
}

type ConfigList struct {
	Uuid     uuid.UUID
	Name     string
	Endpoint string
	Configs  []Config
}

// Pulled from generated query.sql.go to customize
const getFullClusterConfigs = `-- name: GetFullClusterConfigs :many
WITH cview AS (
  SELECT c.uuid, c.name, c.endpoint , json_group_array(json_object('configtype', cc.config_type, 'config', cc.config)) configs
FROM clusters c
INNER JOIN cluster_configs cc on c.uuid = cc.cluster_uuid
GROUP BY c.uuid
)
SELECT uuid, name, endpoint, configs FROM cview
`

type GetFullClusterConfigsRow struct {
	Uuid     uuid.UUID
	Name     string
	Endpoint string
	Configs  string
}

func (q *Queries) GetFullClusterConfigs(ctx context.Context) ([]ConfigList, error) {
	rows, err := q.db.QueryContext(ctx, getFullClusterConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ConfigList
	for rows.Next() {
		var i GetFullClusterConfigsRow
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Endpoint,
			&i.Configs,
		); err != nil {
			return nil, err
		}
		ii := ConfigList{
			Uuid:     i.Uuid,
			Name:     i.Name,
			Endpoint: i.Endpoint,
		}
		if err = json.Unmarshal([]byte(i.Configs), &ii.Configs); err != nil {
			return nil, err
		}

		items = append(items, ii)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

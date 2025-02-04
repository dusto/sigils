// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const attachHostCluster = `-- name: AttachHostCluster :exec
INSERT INTO host_clusters (host_uuid, cluster_uuid) VALUES (?, ?)
`

type AttachHostClusterParams struct {
	HostUuid    uuid.UUID
	ClusterUuid uuid.UUID
}

func (q *Queries) AttachHostCluster(ctx context.Context, arg AttachHostClusterParams) error {
	_, err := q.db.ExecContext(ctx, attachHostCluster, arg.HostUuid, arg.ClusterUuid)
	return err
}

const attachHostProfile = `-- name: AttachHostProfile :exec
INSERT INTO host_profiles (host_uuid, profile_id) VALUES (?, ?)
`

type AttachHostProfileParams struct {
	HostUuid  uuid.UUID
	ProfileID int64
}

func (q *Queries) AttachHostProfile(ctx context.Context, arg AttachHostProfileParams) error {
	_, err := q.db.ExecContext(ctx, attachHostProfile, arg.HostUuid, arg.ProfileID)
	return err
}

const deleteCluster = `-- name: DeleteCluster :exec
DELETE FROM clusters where uuid = ?
`

func (q *Queries) DeleteCluster(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCluster, argUuid)
	return err
}

const deleteHost = `-- name: DeleteHost :exec
DELETE FROM hosts where uuid = ?
`

func (q *Queries) DeleteHost(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteHost, argUuid)
	return err
}

const deletePatch = `-- name: DeletePatch :exec
DELETE FROM patches where id = ?
`

func (q *Queries) DeletePatch(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePatch, id)
	return err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles where id = ?
`

func (q *Queries) DeleteProfile(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, id)
	return err
}

const insertCluster = `-- name: InsertCluster :exec
INSERT INTO clusters (uuid, name, endpoint) VALUES (?, ?, ?)
`

type InsertClusterParams struct {
	Uuid     uuid.UUID
	Name     string
	Endpoint string
}

func (q *Queries) InsertCluster(ctx context.Context, arg InsertClusterParams) error {
	_, err := q.db.ExecContext(ctx, insertCluster, arg.Uuid, arg.Name, arg.Endpoint)
	return err
}

const insertClusterConfig = `-- name: InsertClusterConfig :exec
INSERT INTO cluster_configs (cluster_uuid, config_type, config) VALUES (?, ?, ?)
`

type InsertClusterConfigParams struct {
	ClusterUuid uuid.UUID
	ConfigType  string
	Config      string
}

func (q *Queries) InsertClusterConfig(ctx context.Context, arg InsertClusterConfigParams) error {
	_, err := q.db.ExecContext(ctx, insertClusterConfig, arg.ClusterUuid, arg.ConfigType, arg.Config)
	return err
}

const insertHost = `-- name: InsertHost :one
INSERT INTO hosts ( uuid, fqdn, node_type ) VALUES ( ?, ?, ? )
RETURNING uuid
`

type InsertHostParams struct {
	Uuid     uuid.UUID
	Fqdn     string
	NodeType string
}

func (q *Queries) InsertHost(ctx context.Context, arg InsertHostParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, insertHost, arg.Uuid, arg.Fqdn, arg.NodeType)
	var uuid uuid.UUID
	err := row.Scan(&uuid)
	return uuid, err
}

const insertPatch = `-- name: InsertPatch :one
INSERT INTO patches ( profile_id, node_type, fqdn, patch ) VALUES ( ?, ?, ?, ? )
RETURNING id
`

type InsertPatchParams struct {
	ProfileID int64
	NodeType  string
	Fqdn      string
	Patch     string
}

func (q *Queries) InsertPatch(ctx context.Context, arg InsertPatchParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertPatch,
		arg.ProfileID,
		arg.NodeType,
		arg.Fqdn,
		arg.Patch,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertProfile = `-- name: InsertProfile :one
INSERT INTO profiles ( name ) VALUES ( ? )
RETURNING id
`

func (q *Queries) InsertProfile(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertProfile, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateHost = `-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, node_type = ?, uuid = ? where uuid = ?
`

type UpdateHostParams struct {
	Fqdn     string
	NodeType string
	Uuid     uuid.UUID
	Uuid_2   uuid.UUID
}

func (q *Queries) UpdateHost(ctx context.Context, arg UpdateHostParams) error {
	_, err := q.db.ExecContext(ctx, updateHost,
		arg.Fqdn,
		arg.NodeType,
		arg.Uuid,
		arg.Uuid_2,
	)
	return err
}

const updateProfile = `-- name: UpdateProfile :exec
UPDATE profiles set name = ? WHERE id = ?
`

type UpdateProfileParams struct {
	Name string
	ID   int64
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateProfile, arg.Name, arg.ID)
	return err
}

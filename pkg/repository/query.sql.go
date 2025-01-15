// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"
	"database/sql"
)

const attachHostProfile = `-- name: AttachHostProfile :exec
INSERT INTO host_profiles (host_id, profile_id) VALUES ( ?, ? )
`

type AttachHostProfileParams struct {
	HostID    int64
	ProfileID int64
}

func (q *Queries) AttachHostProfile(ctx context.Context, arg AttachHostProfileParams) error {
	_, err := q.db.ExecContext(ctx, attachHostProfile, arg.HostID, arg.ProfileID)
	return err
}

const deleteConfig = `-- name: DeleteConfig :exec
DELETE FROM base where id = ?
`

func (q *Queries) DeleteConfig(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteConfig, id)
	return err
}

const deleteHost = `-- name: DeleteHost :exec
DELETE FROM hosts where id = ?
`

func (q *Queries) DeleteHost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteHost, id)
	return err
}

const deleteHostProfile = `-- name: DeleteHostProfile :exec
DELETE FROM host_profiles where host_id = ? and profile_id = ?
`

type DeleteHostProfileParams struct {
	HostID    int64
	ProfileID int64
}

func (q *Queries) DeleteHostProfile(ctx context.Context, arg DeleteHostProfileParams) error {
	_, err := q.db.ExecContext(ctx, deleteHostProfile, arg.HostID, arg.ProfileID)
	return err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles where id = ?
`

func (q *Queries) DeleteProfile(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, id)
	return err
}

const getConfigById = `-- name: GetConfigById :one
SELECT id, name, config FROM base WHERE id = ? LIMIT 1
`

func (q *Queries) GetConfigById(ctx context.Context, id int64) (Base, error) {
	row := q.db.QueryRowContext(ctx, getConfigById, id)
	var i Base
	err := row.Scan(&i.ID, &i.Name, &i.Config)
	return i, err
}

const getConfigByName = `-- name: GetConfigByName :one
SELECT id, name, config FROM base WHERE name = ? LIMIT 1
`

func (q *Queries) GetConfigByName(ctx context.Context, name string) (Base, error) {
	row := q.db.QueryRowContext(ctx, getConfigByName, name)
	var i Base
	err := row.Scan(&i.ID, &i.Name, &i.Config)
	return i, err
}

const getHostByFqdn = `-- name: GetHostByFqdn :one
SELECT id, fqdn, base_type, network FROM hosts where fqdn = ? LIMIT 1
`

func (q *Queries) GetHostByFqdn(ctx context.Context, fqdn string) (Host, error) {
	row := q.db.QueryRowContext(ctx, getHostByFqdn, fqdn)
	var i Host
	err := row.Scan(
		&i.ID,
		&i.Fqdn,
		&i.BaseType,
		&i.Network,
	)
	return i, err
}

const getHostById = `-- name: GetHostById :one
SELECT id, fqdn, base_type, network FROM hosts where id = ? LIMIT 1
`

func (q *Queries) GetHostById(ctx context.Context, id int64) (Host, error) {
	row := q.db.QueryRowContext(ctx, getHostById, id)
	var i Host
	err := row.Scan(
		&i.ID,
		&i.Fqdn,
		&i.BaseType,
		&i.Network,
	)
	return i, err
}

const getHostProfiles = `-- name: GetHostProfiles :many
SELECT host_profiles.profile_id, profiles.name, patches.patch FROM host_profiles
JOIN profiles ON profiles.id = host_profiles.profile_id
JOIN patches ON  patches.profile_id = host_profiles.profile_id
WHERE host_id = ?
`

type GetHostProfilesRow struct {
	ProfileID int64
	Name      string
	Patch     string
}

func (q *Queries) GetHostProfiles(ctx context.Context, hostID int64) ([]GetHostProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getHostProfiles, hostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetHostProfilesRow
	for rows.Next() {
		var i GetHostProfilesRow
		if err := rows.Scan(&i.ProfileID, &i.Name, &i.Patch); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfileById = `-- name: GetProfileById :one
SELECT id, name FROM profiles WHERE id = ? LIMIT 1
`

func (q *Queries) GetProfileById(ctx context.Context, id int64) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileById, id)
	var i Profile
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getProfileByName = `-- name: GetProfileByName :one
SELECT id, name FROM profiles WHERE name = ? LIMIT 1
`

func (q *Queries) GetProfileByName(ctx context.Context, name string) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileByName, name)
	var i Profile
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getProfilePatches = `-- name: GetProfilePatches :many
SELECT id, profile_id, patch FROM patches where profile_id = ?
`

func (q *Queries) GetProfilePatches(ctx context.Context, profileID sql.NullInt64) ([]Patch, error) {
	rows, err := q.db.QueryContext(ctx, getProfilePatches, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Patch
	for rows.Next() {
		var i Patch
		if err := rows.Scan(&i.ID, &i.ProfileID, &i.Patch); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertConfig = `-- name: InsertConfig :one
INSERT INTO base (name, config) VALUES (?, ?)
RETURNING id
`

type InsertConfigParams struct {
	Name   string
	Config string
}

func (q *Queries) InsertConfig(ctx context.Context, arg InsertConfigParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertConfig, arg.Name, arg.Config)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertHost = `-- name: InsertHost :one
INSERT INTO hosts ( fqdn, base_type, network ) VALUES ( ?, ?, ? )
RETURNING id
`

type InsertHostParams struct {
	Fqdn     string
	BaseType int64
	Network  string
}

func (q *Queries) InsertHost(ctx context.Context, arg InsertHostParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertHost, arg.Fqdn, arg.BaseType, arg.Network)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertPatch = `-- name: InsertPatch :one
INSERT INTO patches ( profile_id, patch ) VALUES ( ?, ? )
RETURNING id
`

type InsertPatchParams struct {
	ProfileID sql.NullInt64
	Patch     string
}

func (q *Queries) InsertPatch(ctx context.Context, arg InsertPatchParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertPatch, arg.ProfileID, arg.Patch)
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

const listConfigs = `-- name: ListConfigs :many
SELECT id, name, config FROM base
`

func (q *Queries) ListConfigs(ctx context.Context) ([]Base, error) {
	rows, err := q.db.QueryContext(ctx, listConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Base
	for rows.Next() {
		var i Base
		if err := rows.Scan(&i.ID, &i.Name, &i.Config); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProfiles = `-- name: ListProfiles :many
SELECT id, name FROM profiles
`

func (q *Queries) ListProfiles(ctx context.Context) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, listProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateConfig = `-- name: UpdateConfig :exec
UPDATE base SET name = ?, config = ? WHERE id = ?
`

type UpdateConfigParams struct {
	Name   string
	Config string
	ID     int64
}

func (q *Queries) UpdateConfig(ctx context.Context, arg UpdateConfigParams) error {
	_, err := q.db.ExecContext(ctx, updateConfig, arg.Name, arg.Config, arg.ID)
	return err
}

const updateHost = `-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, base_type = ?, network = ? where id = ?
`

type UpdateHostParams struct {
	Fqdn     string
	BaseType int64
	Network  string
	ID       int64
}

func (q *Queries) UpdateHost(ctx context.Context, arg UpdateHostParams) error {
	_, err := q.db.ExecContext(ctx, updateHost,
		arg.Fqdn,
		arg.BaseType,
		arg.Network,
		arg.ID,
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

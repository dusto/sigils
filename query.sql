-- name: InsertConfig :one
INSERT INTO base (name, config) VALUES (?, ?)
RETURNING id;

-- name: UpdateConfig :exec
UPDATE base SET name = ?, config = ? WHERE id = ?;

-- name: DeleteConfig :exec
DELETE FROM base where id = ?;

-- name: GetConfigById :one
SELECT * FROM base WHERE id = ? LIMIT 1;

-- name: GetConfigByName :one
SELECT * FROM base WHERE name = ? LIMIT 1;

-- name: ListConfigs :many
SELECT * FROM base;

-- name: InsertProfile :one
INSERT INTO profiles ( name ) VALUES ( ? )
RETURNING id;

-- name: UpdateProfile :exec
UPDATE profiles set name = ? WHERE id = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles where id = ?;

-- name: GetProfileById :one
SELECT * FROM profiles WHERE id = ? LIMIT 1;

-- name: GetProfileByName :one
SELECT * FROM profiles WHERE name = ? LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM profiles;

-- name: InsertPatch :one
INSERT INTO patches ( profile_id, patch ) VALUES ( ?, ? )
RETURNING id;

-- name: GetProfilePatches :many
SELECT * FROM patches where profile_id = ?; 

-- name: InsertHost :one
INSERT INTO hosts ( fqdn, base_type, network ) VALUES ( ?, ?, ? )
RETURNING id;

-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, base_type = ?, network = ? where id = ?;

-- name: DeleteHost :exec
DELETE FROM hosts where id = ?;

-- name: GetHostById :one
SELECT * FROM hosts where id = ? LIMIT 1;

-- name: GetHostByFqdn :one
SELECT * FROM hosts where fqdn = ? LIMIT 1;

-- name: AttachHostProfile :exec
INSERT INTO host_profiles (host_id, profile_id) VALUES ( ?, ? );

-- name: DeleteHostProfile :exec
DELETE FROM host_profiles where host_id = ? and profile_id = ?;

-- name: GetHostProfiles :many
SELECT host_profiles.profile_id, profiles.name, patches.patch FROM host_profiles
JOIN profiles ON profiles.id = host_profiles.profile_id
JOIN patches ON  patches.profile_id = host_profiles.profile_id
WHERE host_id = ?;

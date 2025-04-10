-- name: InsertCluster :exec
INSERT INTO clusters (uuid, name, endpoint) VALUES (?, ?, ?);

-- name: InsertClusterConfig :exec
INSERT INTO cluster_configs (cluster_uuid, config_type, config) VALUES (?, ?, ?);

-- name: DeleteCluster :exec
DELETE FROM clusters where uuid = ?;

-- name: InsertProfile :one
INSERT INTO profiles ( name ) VALUES ( ? )
RETURNING id;

-- name: UpdateProfile :exec
UPDATE profiles set name = ? WHERE id = ?;

-- name: GetProfileId :one
SELECT id FROM profiles where name = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles where id = ?;

-- name: InsertPatch :exec
INSERT INTO patches ( profile_id, nodetype, fqdn, patch ) VALUES ( ?, ?, ?, ? );

-- name: DeletePatch :exec
DELETE FROM patches where id = ?;

-- name: InsertHost :exec
INSERT INTO hosts ( uuid, mac, fqdn, nodetype ) VALUES ( ?, ?, ?, ? );

-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, nodetype = ?, uuid = ? where uuid = ?;

-- name: DeleteHost :exec
DELETE FROM hosts where uuid = ?;

-- name: AttachHostProfile :exec
INSERT INTO host_profiles (host_uuid, profile_id) VALUES (?, ?);

-- name: AttachHostCluster :exec
INSERT INTO host_clusters (host_uuid, cluster_uuid) VALUES (?, ?);

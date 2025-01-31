-- name: InsertCluster :exec
INSERT INTO clusters (uuid, name, endpoint) VALUES (?, ?, ?);

-- name: InsertClusterConfig :exec
INSERT INTO cluster_configs (cluster_uuid, config_type, config) VALUES (?, ?, ?);

-- name: DeleteCluster :exec
DELETE FROM clusters where uuid = ?;

-- name: GetClusterByUUID :one
SELECT clusters.uuid, clusters.name, clusters.endpoint, cluster_configs.config_type, cluster_configs.config FROM clusters
JOIN cluster_configs ON cluster_configs.cluster_uuid = clusters.uuid
WHERE uuid = ?;

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
INSERT INTO patches ( profile_id, controlplane, worker, host, patch ) VALUES ( ?, ?, ?, ?, ? )
RETURNING id;

-- name: GetProfilePatches :many
SELECT * FROM patches where profile_id = ?; 

-- name: InsertHost :one
INSERT INTO hosts ( uuid, fqdn, node_type ) VALUES ( ?, ?, ? )
RETURNING uuid;

-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, node_type = ?, uuid = ? where uuid = ?;

-- name: DeleteHost :exec
DELETE FROM hosts where uuid = ?;

-- name: GetHostByUuid :one
SELECT * FROM hosts where uuid = ? LIMIT 1;

-- name: GetHostByFqdn :one
SELECT * FROM hosts where fqdn = ? LIMIT 1;

-- name: AttachHostProfile :exec
INSERT INTO host_profiles (host_uuid, profile_id) VALUES ( ?, ? );

-- name: DeleteHostProfile :exec
DELETE FROM host_profiles where host_uuid = ? and profile_id = ?;

-- name: GetHostProfiles :many
SELECT host_profiles.profile_id, profiles.name, patches.patch FROM host_profiles
JOIN profiles ON profiles.id = host_profiles.profile_id
JOIN patches ON  patches.profile_id = host_profiles.profile_id
WHERE host_uuid = ?;

-- name: AttachHostCluster :exec
INSERT INTO host_clusters (host_uuid, cluster_uuid) VALUES ( ?, ? );

-- name: DeleteHostCluster :exec
DELETE FROM host_clusters where host_uuid = ? and cluster_uuid = ?;

-- name: GetHostClusters :many
SELECT hosts.uuid, hosts.fqdn, clusters.uuid, clusters.name  FROM host_clusters
JOIN clusters ON clusters.uuid = host_clusters.cluster_uuid
JOIN hosts ON hosts.id = host_clusters.host_uuid
WHERE host_clusters.host_uuid = ?;

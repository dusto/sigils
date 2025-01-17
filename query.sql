-- name: InsertCluster :one
INSERT INTO clusters (name, endpoint) VALUES (?, ?)
RETURNING id;

-- name: InsertClusterConfig :exec
INSERT INTO cluster_configs (clus_id, node_type, config) VALUES (?, ?, ?);

-- name: DeleteCluster :exec
DELETE FROM clusters where id = ?;

-- name: GetClusterById :one
SELECT * FROM clusters WHERE id = ? LIMIT 1;

-- name: GetClusterByName :one
SELECT * FROM clusters WHERE name = ? LIMIT 1;

-- name: ListClusters :many
SELECT * FROM clusters;

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
INSERT INTO hosts ( fqdn, node_type, network ) VALUES ( ?, ?, ? )
RETURNING id;

-- name: UpdateHost :exec
UPDATE hosts set fqdn = ?, node_type = ?, network = ? where id = ?;

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

-- name: AttachHostCluster :exec
INSERT INTO host_clusters (host_id, clus_id) VALUES ( ?, ? );

-- name: DeleteHostCluster :exec
DELETE FROM host_clusters where host_id = ? and clus_id = ?;

-- name: GetHostClusters :many
SELECT hosts.id, hosts.fqdn, clusters.id, clusters.name  FROM host_clusters
JOIN clusters ON clusters.id = host_clusters.clus_id
JOIN hosts ON hosts.id = host_clusters.host_id
WHERE host_clusters.host_id = ?;

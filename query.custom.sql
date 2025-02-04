-- name: GetClusterByUUID :one
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
WITH cview AS (
  SELECT c.uuid, c.name, c.endpoint , json_group_array(json_object('configtype', cc.config_type, 'config', cc.config)) configs
FROM clusters c
INNER JOIN cluster_configs cc on c.uuid = cc.cluster_uuid
GROUP BY c.uuid
)
SELECT uuid, name, endpoint, configs FROM cview
WHERE uuid = ?;

-- name: GetFullClusterConfigs :many
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
WITH cview AS (
  SELECT c.uuid, c.name, c.endpoint , json_group_array(json_object('configtype', cc.config_type, 'config', cc.config)) configs
FROM clusters c
INNER JOIN cluster_configs cc on c.uuid = cc.cluster_uuid
GROUP BY c.uuid
)
SELECT uuid, name, endpoint, configs FROM cview;

-- name: GetProfile :one
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
SELECT 
  p.id,
  p.name,
  (
    SELECT
    json_group_array(json_object(
                'id', pa.id, 
                'node_type', pa.node_type,
                'fqdn', pa.fqdn,
                'patch', pa.patch))
    FROM patches pa
    WHERE  pa.profile_id = p.id 
    GROUP BY pa.profile_id
  ) patches
FROM profiles p
WHERE p.id = ?
GROUP BY p.id;

-- name: GetProfiles :many
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
SELECT 
  p.id,
  p.name,
  (
    SELECT
    json_group_array(json_object(
                'id', pa.id, 
                'node_type', pa.node_type,
                'fqdn', pa.fqdn,
                'patch', pa.patch))
    FROM patches pa
    WHERE  pa.profile_id = p.id 
    GROUP BY pa.profile_id
  ) patches
FROM profiles p
GROUP BY p.id;

-- name: GetHosts :many
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
SELECT h.uuid, h.fqdn, h.node_type,
json_group_array(
  (SELECT json_object(
    'id', p.id,
    'name', p.name,
    'patches', (
    SELECT
    json_group_array(json_object(
                'id', pa.id, 
                'profile_id', pa.profile_id,
                'node_type', pa.node_type,
                'fqdn', pa.fqdn,
                'patch', pa.patch))
    FROM patches pa
    WHERE ((pa.node_type IN (0,h.node_type) AND pa.fqdn = '') OR pa.fqdn = h.fqdn) AND pa.profile_id = p.id 
    GROUP BY pa.profile_id
    )
  )
  FROM profiles p
  INNER JOIN host_profiles hp ON hp.profile_id = p.id
  WHERE hp.host_uuid = h.uuid
  GROUP BY p.id
)) profiles
FROM hosts h
GROUP BY h.uuid;

-- name: GetHost :one
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
SELECT h.uuid, h.fqdn, h.node_type,
json_group_array(
  (SELECT json_object(
    'id', p.id,
    'name', p.name,
    'patches', (
    SELECT
    json_group_array(json_object(
                'id', pa.id, 
                'profile_id', pa.profile_id,
                'node_type', pa.node_type,
                'fqdn', pa.fqdn,
                'patch', pa.patch))
    FROM patches pa
    WHERE ((pa.node_type IN (0,h.node_type) AND pa.fqdn = '') OR pa.fqdn = h.fqdn) AND pa.profile_id = p.id 
    GROUP BY pa.profile_id
    )
  )
  FROM profiles p
  INNER JOIN host_profiles hp ON hp.profile_id = p.id
  WHERE hp.host_uuid = h.uuid
  GROUP BY p.id
)) profiles
FROM hosts h
WHERE h.uuid = ?
GROUP BY h.uuid;

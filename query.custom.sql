-- name: GetFullClusterConfigs :many
-- IT IS HANDLED VIA A CUSTOM FUNCTION DO NOT AUTO GENERATE
WITH cview AS (
  SELECT c.uuid, c.name, c.endpoint , json_group_array(json_object('configtype', cc.config_type, 'config', cc.config) configs
FROM clusters c
INNER JOIN cluster_configs cc on c.uuid = cc.cluster_uuid
GROUP BY c.uuid
)
SELECT uuid, name, endpoint, configs FROM cview;

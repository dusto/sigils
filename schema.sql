-- Store Cluster configs
CREATE TABLE IF NOT EXISTS clusters (
  uuid BLOB NOT NULL,
  name TEXT NOT NULL,
  endpoint TEXT NOT NULL,
  UNIQUE(name) ON CONFLICT FAIL,
  UNIQUE(endpoint) ON CONFLICT FAIL,
  UNIQUE(uuid) ON CONFLICT FAIL
);
CREATE INDEX IF NOT EXISTS clusters_name_idx on clusters (name);
CREATE INDEX IF NOT EXISTS clusters_uuid_idx on clusters (uuid);

CREATE TABLE IF NOT EXISTS cluster_configs (
  id INTEGER PRIMARY KEY,
  cluster_uuid BLOB NOT NULL,
  config_type TEXT NOT NULL,
  config TEXT NOT NULL,
  FOREIGN KEY (cluster_uuid) REFERENCES clusters(uuid) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS cluster_configs_idx on cluster_configs (cluster_uuid);
CREATE INDEX IF NOT EXISTS cluster_configs_type_idx on cluster_configs (config_type);

-- Store profile info
CREATE TABLE IF NOT EXISTS profiles (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

-- Store profile patch sets
CREATE TABLE IF NOT EXISTS patches (
  id INTEGER PRIMARY KEY,
  profile_id INTEGER NOT NULL,
  node_type TEXT NOT NULL DEFAULT 0,
  fqdn TEXT NOT NULL,
  patch TEXT NOT NULL,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS profile_patches_idx on patches ( profile_id );

-- Store hosts
CREATE TABLE IF NOT EXISTS hosts (
  uuid BLOB NOT NULL,
  fqdn TEXT NOT NULL,
  node_type TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS host_uuid_idx on hosts ( uuid );
CREATE INDEX IF NOT EXISTS host_fqdn_idx on hosts ( fqdn );
CREATE INDEX IF NOT EXISTS host_node_type_idx on hosts ( node_type );

-- Store hosts profile associations
CREATE TABLE IF NOT EXISTS host_profiles (
  host_uuid BLOB NOT NULL,
  profile_id INTEGER NOT NULL,
  FOREIGN KEY (host_uuid) REFERENCES hosts(uuid) ON DELETE CASCADE,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE,
  primary key (host_uuid, profile_id)
 );
CREATE INDEX IF NOT EXISTS host_profiles_idx on host_profiles (profile_id);
CREATE INDEX IF NOT EXISTS host_hosts_idx on host_profiles (host_uuid);

-- Store hosts cluster associations
CREATE TABLE IF NOT EXISTS host_clusters (
  host_uuid BLOB NOT NULL,
  cluster_uuid INTEGER NOT NULL,
  FOREIGN KEY (host_uuid) REFERENCES hosts(uuid) ON DELETE CASCADE,
  FOREIGN KEY (cluster_uuid) REFERENCES clusters(uuid) ON DELETE CASCADE,
  primary key (host_uuid, cluster_uuid)
 );
CREATE INDEX IF NOT EXISTS host_clusters_cluster_uuidx on host_clusters (cluster_uuid);
CREATE INDEX IF NOT EXISTS host_cluster_host_idx on host_clusters (host_uuid);

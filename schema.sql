-- Store Cluster configs
CREATE TABLE IF NOT EXISTS clusters (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  endpoint TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS clusters_name_idx on clusters (name);

CREATE TABLE IF NOT EXISTS cluster_configs (
  id INTEGER PRIMARY KEY,
  clus_id INTEGER NOT NULL,
  node_type INTEGER NOT NULL,
  config TEXT NOT NULL,
  FOREIGN KEY (clus_id) REFERENCES clusters(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS cluster_configs_idx on cluster_configs (clus_id);
CREATE INDEX IF NOT EXISTS cluster_configs_node_idx on cluster_configs (node_type);

CREATE TABLE IF NOT EXISTS node_types (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

-- Store profile info
CREATE TABLE IF NOT EXISTS profiles (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

-- Store profile patch sets
CREATE TABLE IF NOT EXISTS patches (
  id INTEGER PRIMARY KEY,
  profile_id INTEGER,
  patch TEXT NOT NULL,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS profile_patches_idx on patches ( profile_id );

-- Store hosts
CREATE TABLE IF NOT EXISTS hosts (
  id INTEGER PRIMARY KEY,
  fqdn TEXT NOT NULL,
  node_type INTEGER NOT NULL,
  network TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS host_fqdn_idx on hosts ( fqdn );
CREATE INDEX IF NOT EXISTS host_node_type_idx on hosts ( node_type );


-- Store hosts profile associations
CREATE TABLE IF NOT EXISTS host_profiles (
  host_id INTEGER NOT NULL,
  profile_id INTEGER NOT NULL,
  FOREIGN KEY (host_id) REFERENCES hosts(id) ON DELETE CASCADE,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE,
  primary key (host_id, profile_id)
 );
CREATE INDEX IF NOT EXISTS host_profiles_idx on host_profiles (profile_id);
CREATE INDEX IF NOT EXISTS host_hosts_idx on host_profiles (host_id);

-- Store hosts cluster associations
CREATE TABLE IF NOT EXISTS host_clusters (
  host_id INTEGER NOT NULL,
  clus_id INTEGER NOT NULL,
  FOREIGN KEY (host_id) REFERENCES hosts(id) ON DELETE CASCADE,
  FOREIGN KEY (clus_id) REFERENCES clusters(id) ON DELETE CASCADE,
  primary key (host_id, clus_id)
 );
CREATE INDEX IF NOT EXISTS host_clusters_clus_idx on host_clusters (clus_id);
CREATE INDEX IF NOT EXISTS host_cluster_host_idx on host_clusters (host_id);

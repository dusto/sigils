-- Store Base configs
CREATE TABLE base (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  config text NOT NULL
);

-- Store profile info
CREATE TABLE profiles (
  id INTEGER PRIMARY KEY,
  name text NOT NULL
);

-- Store profile patch sets
CREATE TABLE patches (
  id INTEGER PRIMARY KEY,
  profile_id INTEGER,
  patch text NOT NULL,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);
CREATE INDEX profile_patches_idx on patches ( profile_id );

-- Store hosts
CREATE TABLE hosts (
  id INTEGER PRIMARY KEY,
  fqdn text NOT NULL,
  base_type INTEGER NOT NULL,
  network text NOT NULL
);
CREATE INDEX host_fqdn_idx on hosts ( fqdn );

-- Store hosts profile associations
CREATE TABLE host_profiles (
  host_id INTEGER NOT NULL,
  profile_id INTEGER NOT NULL,
  FOREIGN KEY (host_id) REFERENCES hosts(id) ON DELETE CASCADE,
  FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE,
  primary key (host_id, profile_id)
 );
CREATE INDEX host_profiles_idx on hosts_profiles (profile_id);
CREATE INDEX host_hosts_idx on hosts_profiles (host_id);


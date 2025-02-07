# Sigils

> Version 0.0.1

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/clusters](#getclusters) | List Clusters |
| POST | [/clusters](#postclusters) | Import/Manually Add Cluster |
| POST | [/clusters/generate](#postclustersgenerate) | Automatically generate new Cluster |
| POST | [/clusters/{cluster_uuid}/attach/{host_uuid}](#postclusterscluster_uuidattachhost_uuid) | Attach/Add host to cluster |
| DELETE | [/clusters/{uuid}](#deleteclustersuuid) | Delete Cluster |
| GET | [/clusters/{uuid}](#getclustersuuid) | Get Cluster |
| GET | [/hosts](#gethosts) | List hosts |
| POST | [/hosts](#posthosts) | Add Hosts |
| POST | [/hosts/{host_uuid}/attach/{profile_id}](#posthostshost_uuidattachprofile_id) | Attach/Add profile to host |
| DELETE | [/hosts/{uuid}](#deletehostsuuid) | Delete Host |
| GET | [/hosts/{uuid}](#gethostsuuid) | Get Host |
| GET | [/machineconfig](#getmachineconfig) | Get a patched machineconfig for a specific host |
| GET | [/profiles](#getprofiles) | List profiles |
| POST | [/profiles](#postprofiles) | Add Profiles |
| DELETE | [/profiles/{id}](#deleteprofilesid) | Delete Profile |
| GET | [/profiles/{id}](#getprofilesid) | Get Profile |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Cluster | [#/components/schemas/Cluster](#componentsschemascluster) |  |
| ClusterConfig | [#/components/schemas/ClusterConfig](#componentsschemasclusterconfig) |  |
| ClusterGen | [#/components/schemas/ClusterGen](#componentsschemasclustergen) |  |
| ErrorDetail | [#/components/schemas/ErrorDetail](#componentsschemaserrordetail) |  |
| ErrorModel | [#/components/schemas/ErrorModel](#componentsschemaserrormodel) |  |
| Host | [#/components/schemas/Host](#componentsschemashost) |  |
| HostInput | [#/components/schemas/HostInput](#componentsschemashostinput) |  |
| Patch | [#/components/schemas/Patch](#componentsschemaspatch) |  |
| Profile | [#/components/schemas/Profile](#componentsschemasprofile) |  |

## Path Details

***

### [GET]/clusters

- Summary  
List Clusters

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Cluster"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/clusters

- Summary  
Import/Manually Add Cluster

#### RequestBody

- application/json

```ts
{
  "items": {
    "$ref": "#/components/schemas/Cluster"
  },
  "type": [
    "array",
    "null"
  ]
}
```

#### Responses

- 201 Created

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Cluster"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/clusters/generate

- Summary  
Automatically generate new Cluster

#### RequestBody

- application/json

```ts
{
  "items": {
    "$ref": "#/components/schemas/ClusterGen"
  },
  "type": [
    "array",
    "null"
  ]
}
```

#### Responses

- 201 Created

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Cluster"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/clusters/{cluster_uuid}/attach/{host_uuid}

- Summary  
Attach/Add host to cluster

#### Responses

- 201 Created

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [DELETE]/clusters/{uuid}

- Summary  
Delete Cluster

#### Responses

- 204 No Content

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/clusters/{uuid}

- Summary  
Get Cluster

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Cluster"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/hosts

- Summary  
List hosts

#### Parameters(Query)

```ts
search?: string
```

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Host"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/hosts

- Summary  
Add Hosts

#### RequestBody

- application/json

```ts
{
  "items": {
    "$ref": "#/components/schemas/HostInput"
  },
  "type": [
    "array",
    "null"
  ]
}
```

#### Responses

- 201 Created

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/hosts/{host_uuid}/attach/{profile_id}

- Summary  
Attach/Add profile to host

#### Responses

- 201 Created

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [DELETE]/hosts/{uuid}

- Summary  
Delete Host

#### Responses

- 204 No Content

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/hosts/{uuid}

- Summary  
Get Host

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Host"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/machineconfig

- Summary  
Get a patched machineconfig for a specific host

#### Parameters(Query)

```ts
uuid?: string
```

```ts
mac?: string
```

```ts
fqdn?: string
```

#### Responses

- 201 Created

`application/json`

```ts
{
  "type": "string"
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/profiles

- Summary  
List profiles

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Profile"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [POST]/profiles

- Summary  
Add Profiles

#### RequestBody

- application/json

```ts
{
  "items": {
    "$ref": "#/components/schemas/Profile"
  },
  "type": [
    "array",
    "null"
  ]
}
```

#### Responses

- 201 Created

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [DELETE]/profiles/{id}

- Summary  
Delete Profile

#### Parameters(Query)

```ts
// Patch ID to remove a patch from a profile
patch_id?: integer
```

#### Responses

- 204 No Content

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

***

### [GET]/profiles/{id}

- Summary  
Get Profile

#### Responses

- 200 OK

`application/json`

```ts
{
  "items": {
    "$ref": "#/components/schemas/Profile"
  },
  "type": [
    "array",
    "null"
  ]
}
```

- default Error

`application/problem+json`

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

## References

### #/components/schemas/Cluster

```ts
{
  // Cluster Configs
  configs?: array | null
  // Cluster Endpoint
  endpoint: string
  // Cluster Name
  name: string
  // Cluster ID
  uuid?: string
}
```

### #/components/schemas/ClusterConfig

```ts
{
  // Yaml representation of the config
  config: string
  // Config type controlplane,worker,talosctl
  configtype: enum[controlplane, worker, talosctl]
  // ID of Config
  id?: integer
  // Cluster ID
  uuid?: string
}
```

### #/components/schemas/ClusterGen

```ts
{
  // Cluster Endpoint
  endpoint: string
  // Kubernetes Version https://www.talos.dev/v1.9/introduction/support-matrix/ for supported versions
  kubernetesversion: string
  // Cluster Name
  name: string
  // TalosOS version for config contract
  talosversion: string //default: 1.9
}
```

### #/components/schemas/ErrorDetail

```ts
{
  // Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id'
  location?: string
  // Error message text
  message?: string
}
```

### #/components/schemas/ErrorModel

```ts
{
  // A URL to the JSON Schema for this object.
  $schema?: string
  // A human-readable explanation specific to this occurrence of the problem.
  detail?: string
  // Optional list of individual error details
  errors?: array | null
  // A URI reference that identifies the specific occurrence of the problem.
  instance?: string
  // HTTP status code
  status?: integer
  // A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
  title?: string
  // A URI reference to human-readable documentation for the error.
  type?: string //default: about:blank
}
```

### #/components/schemas/Host

```ts
{
  // Name of cluster associated with host
  clustername?: string
  // Host FQDN
  fqdn: string
  // Host Mac Address
  mac?: string
  // Host Node Type
  nodetype: enum[controlplane, worker]
  // List of Patches for host
  profiles?: array | null
  // Host SMBIOS UUID
  uuid: string
}
```

### #/components/schemas/HostInput

```ts
{
  // Host FQDN
  fqdn: string
  // Host Mac Address
  mac?: string
  // Host Node Type
  nodetype: enum[controlplane, worker]
  // List of Profile Ids to associate with Host
  profileids?: array | null
  // Host SMBIOS UUID
  uuid: string
}
```

### #/components/schemas/Patch

```ts
{
  // Host FQDN/UUID of specific host for patch to apply
  fqdn?: string
  // Patch ID
  id?: integer
  // Type of node for patch to apply to
  nodetype?: enum[all, controlplane, worker] //default: all
  // JSON6902 patch or Strategic Merge patch
  patch?: string
}
```

### #/components/schemas/Profile

```ts
{
  // Profile ID
  id?: integer
  // Profile Name
  name?: string
  // Collection of patches associated with profile
  patches?: array | null
}
```

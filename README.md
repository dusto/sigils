# Sigils

Rest API for generating, patch, managing machine configs for Talos Linux.

## Why

To fill a need for creating/managing home lab environment and constantly creating new clusters to test different configurations.
Also, this was a good project to refresh my Go skills after not using Go for a couple years.

## Features

* Uses the packages from `github.com/siderolabs/talos` to generate and patch all configs so it acts the same as `talosctl`.
* Generate base controlplane, worker, talosctl configs and save as a cluster.
* Import existing base controlplane, worker, talosctl configs.
* Allows for multiple cluster definitions.
* Profiles that are a collection of Patches. Support both JSON6902 and Strategic patches that talosctl supports.
* Hosts are stored with UUID, Mac, FQDN (hostname) so they can be fetched via PXE booting based on parameters available to `talos.config` kernel parameters.
* Associate profiles to multiple hosts for reusing patches
* Associate Host to a cluster
* Removing the main components Cluster, Host, Profile does not remove any other only associations
* Based on association the Host has with a Cluster and Profiles generates a patched machine config of the correct type i.e. controlplane or worker


## TODO

* Improve CRUD operations:
    * Missing some operations for disassociation clusters/hosts/profiles
    * Add ability to manage update clusters/hosts/profiles/patch
    * Add abilty to manage patches in a profile, currently have to delete a profile and re-add with changes
* Auth:
    * No intentions exposing to the internet
    * No expectations for this to be used for anything beyond homelab and my own musings
* GRPC:
    * Move all CRUD operations to be done via GRPC and allow for a single HTTP endpoint for PXE

## API docs

Sigils is using `https://huma.rocks` which gives live `docs` endpoint to provide the openapi spec along with UI for working with the API. Go to `http://localhost:8888/docs` if you start the binary on your local host.

The openapi spec is in the `docs/` directory in the repo

Also, see the markdown of the openapi spec [here](API.md)

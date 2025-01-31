package route

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/repository"
	"github.com/dusto/sigils/internal/talosconfig"
	"github.com/google/uuid"
)

type ClusterOutput struct {
	Body []ClusterBody
}

type ClusterListInput struct {
}

type ClusterAnyOfInput struct {
}

type ClusterOneOfInput struct {
	Uuid uuid.UUID `path:"uuid" required:"false"`
}

type ClusterBody struct {
	Uuid     string          `json:"uuid,omitempty" format:"uuid" doc:"Cluster ID"`
	Name     string          `json:"name" doc:"Cluster Name" required:"true"`
	Endpoint string          `json:"endpoint" doc:"Cluster Endpoint" required:"true"`
	Configs  []ClusterConfig `json:"configs,omitempty" doc:"Cluster Configs" require:"false"`
}

type ClusterConfig struct {
	ConfigType string `json:"configtype" doc:"Config type controlplane,worker,talosctl"`
	Config     string `json:"config" doc:"Yaml representation of the config"`
}

type ClusterPostInput struct {
	Body []ClusterBody
}

type ClusterGen struct {
	Name              string `json:"name" doc:"Cluster Name" required:"true"`
	Endpoint          string `json:"endpoint" doc:"Cluster Endpoint" required:"true"`
	KubernetesVersion string `json:"kubernetesversion" doc:"Kubernetes Version https://www.talos.dev/v1.9/introduction/support-matrix/ for supported versions"`
	Talosversion      string `json:"talosversion" doc:"TalosOS version for config contract" default:"1.9"`
}

type ClusterGenPostInput struct {
	Body []ClusterGen
}

func (h *Handler) ClusterGetOneOf(ctx context.Context, input *ClusterOneOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	cluster, err := h.configDB.GetClusterByUUID(ctx, input.Uuid)
	if err != nil {
		return resp, nil
	}

	cBody := ClusterBody{
		Uuid:     cluster.Uuid.String(),
		Name:     cluster.Name,
		Endpoint: cluster.Endpoint,
	}
	resp.Body = append(resp.Body, cBody)

	return resp, nil
}

func (h *Handler) ClusterGetAnyOf(ctx context.Context, input *ClusterAnyOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	clusters, err := h.configDB.GetFullClusterConfigs(ctx)

	if err != nil {
		return nil, huma.Error500InternalServerError("Could not fetch clusters", err)
	}

	for _, cluster := range clusters {
		clus := ClusterBody{
			Uuid:     cluster.Uuid.String(),
			Name:     cluster.Name,
			Endpoint: cluster.Endpoint,
		}

		for _, config := range cluster.Configs {
			ctype := ""
			if config.ConfigType == talosconfig.ConfigTypeControlPlane {
				ctype = "controlplane"
			} else if config.ConfigType == talosconfig.ConfigTypeWorker {
				ctype = "worker"
			} else {
				ctype = "talosctl"
			}
			clus.Configs = append(clus.Configs, ClusterConfig{
				ConfigType: ctype,
				Config:     config.Config,
			})
		}
		resp.Body = append(resp.Body, clus)
	}
	return resp, nil
}

func (h *Handler) ClusterPost(ctx context.Context, input *ClusterPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	for _, cIn := range input.Body {
		if cIn.Uuid == "" {
			cIn.Uuid = uuid.New().String()
			fmt.Printf("New Uuid: %s", cIn.Uuid)
		}
		err := h.configDB.InsertCluster(ctx, repository.InsertClusterParams{
			Uuid:     uuid.MustParse(cIn.Uuid),
			Name:     cIn.Name,
			Endpoint: cIn.Endpoint,
		})
		if err != nil {
			return resp, huma.Error500InternalServerError("Could not add cluster", err)
		}
	}

	return resp, nil
}

func (h *Handler) ClusterDelete(ctx context.Context, input *ClusterOneOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	err := h.configDB.DeleteCluster(ctx, input.Uuid)

	if err != nil {
		return resp, huma.Error500InternalServerError("Could not remove cluster", err)
	}

	return resp, nil
}

func (h *Handler) ClusterGen(ctx context.Context, input *ClusterGenPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	cluster := ClusterBody{}
	cluster.Uuid = uuid.New().String()
	cluster.Name = input.Body[0].Name
	cluster.Endpoint = input.Body[0].Endpoint

	err := h.configDB.InsertCluster(ctx, repository.InsertClusterParams{
		Uuid:     uuid.MustParse(cluster.Uuid),
		Name:     cluster.Name,
		Endpoint: cluster.Endpoint,
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save cluster", err)
	}

	clusterConfigs, err := talosconfig.NewCluster(cluster.Name, cluster.Endpoint, input.Body[0].KubernetesVersion, input.Body[0].Talosversion)
	if err != nil {
		return resp, huma.Error500InternalServerError("Could not generate cluster config", err)
	}

	controlPlane, _ := clusterConfigs.ControlPlaneConfig.Bytes()
	cluster.Configs = append(cluster.Configs, ClusterConfig{
		ConfigType: "controlplane",
		Config:     string(controlPlane),
	})
	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  talosconfig.ConfigTypeControlPlane,
		Config:      string(controlPlane),
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save control plane config", err)
	}

	worker, _ := clusterConfigs.WorkerConfig.Bytes()
	cluster.Configs = append(cluster.Configs, ClusterConfig{
		ConfigType: "worker",
		Config:     string(worker),
	})
	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  talosconfig.ConfigTypeWorker,
		Config:      string(worker),
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save worker config", err)
	}

	talosctl, _ := clusterConfigs.TalosCtlConfig.Bytes()
	cluster.Configs = append(cluster.Configs, ClusterConfig{
		ConfigType: "talosctl",
		Config:     string(talosctl),
	})

	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  talosconfig.ConfigTypeTalosctl,
		Config:      string(talosctl),
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save talosctl config", err)
	}

	resp.Body = append(resp.Body, cluster)

	return resp, nil
}

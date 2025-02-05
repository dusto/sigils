package route

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/dusto/sigils/internal/model"
	"github.com/dusto/sigils/internal/repository"
	"github.com/dusto/sigils/internal/talosconfig"
	"github.com/google/uuid"

	clientconfig "github.com/siderolabs/talos/pkg/machinery/client/config"
	"github.com/siderolabs/talos/pkg/machinery/config/configloader"
)

type ClusterOutput struct {
	Body []model.Cluster
}

type ClusterListInput struct {
}

type ClusterAnyOfInput struct {
}

type ClusterOneOfInput struct {
	Uuid uuid.UUID `path:"uuid" required:"false"`
}

type ClusterAttachHostInput struct {
	ClusterUUID uuid.UUID `path:"cluster_uuid" required:"true"`
	HostUUID    uuid.UUID `path:"host_uuid" required:"true"`
}

type ClusterPostInput struct {
	Body []model.Cluster
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

	resp.Body = append(resp.Body, cluster)

	return resp, nil
}

func (h *Handler) ClusterGetAnyOf(ctx context.Context, input *ClusterAnyOfInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	clusters, err := h.configDB.GetFullClusterConfigs(ctx)

	if err != nil {
		return nil, huma.Error500InternalServerError("Could not fetch clusters", err)
	}

	for _, cluster := range clusters {
		resp.Body = append(resp.Body, cluster)
	}
	return resp, nil
}

func (h *Handler) ClusterPost(ctx context.Context, input *ClusterPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}

	for _, cIn := range input.Body {
		if cIn.Uuid != "" {
			cIn.Uuid = uuid.New().String()
		}

		// TODO: Add transaction

		err := h.configDB.InsertCluster(ctx, repository.InsertClusterParams{
			Uuid:     uuid.MustParse(cIn.Uuid),
			Name:     cIn.Name,
			Endpoint: cIn.Endpoint,
		})

		if err != nil {
			return resp, huma.Error500InternalServerError("Could not add cluster", err)
		}

		if len(cIn.Configs) > 0 {
			for _, config := range cIn.Configs {
				switch config.ConfigType {
				case talosconfig.ConfigTypeTalosctl:
					_, err := clientconfig.FromString(config.Config)

					if err != nil {
						return nil, huma.Error422UnprocessableEntity("Problem validating Talosctl config", err)
					}

				case talosconfig.ConfigTypeControlPlane, talosconfig.ConfigTypeWorker:
					_, err := configloader.NewFromBytes([]byte(config.Config))

					if err != nil {
						return nil, huma.Error422UnprocessableEntity("Problem validating Controlplane/Worker config", err)
					}

				}

				err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
					ClusterUuid: uuid.MustParse(cIn.Uuid),
					ConfigType:  string(config.ConfigType),
					Config:      string(config.Config),
				})
				if err != nil {
					return nil, huma.Error500InternalServerError("Could not save config", err)
				}
			}
		}
	}

	return resp, nil
}

func (h *Handler) ClusterDelete(ctx context.Context, input *ClusterOneOfInput) (*struct{}, error) {

	if err := h.configDB.DeleteCluster(ctx, input.Uuid); err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("Could not remove cluster %s", input.Uuid), err)
	}

	return nil, nil
}

func (h *Handler) ClusterGen(ctx context.Context, input *ClusterGenPostInput) (*ClusterOutput, error) {
	resp := &ClusterOutput{}
	cluster := model.Cluster{}
	cluster.Uuid = uuid.New().String()
	cluster.Name = input.Body[0].Name
	cluster.Endpoint = input.Body[0].Endpoint

	// TODO: Add transaction

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
	cluster.Configs = append(cluster.Configs, model.ClusterConfig{
		ConfigType: talosconfig.ConfigTypeControlPlane,
		Config:     string(controlPlane),
	})
	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  string(talosconfig.ConfigTypeControlPlane),
		Config:      string(controlPlane),
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save control plane config", err)
	}

	worker, _ := clusterConfigs.WorkerConfig.Bytes()
	cluster.Configs = append(cluster.Configs, model.ClusterConfig{
		ConfigType: talosconfig.ConfigTypeWorker,
		Config:     string(worker),
	})
	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  string(talosconfig.ConfigTypeWorker),
		Config:      string(worker),
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save worker config", err)
	}

	talosctl, _ := clusterConfigs.TalosCtlConfig.Bytes()
	cluster.Configs = append(cluster.Configs, model.ClusterConfig{
		ConfigType: talosconfig.ConfigTypeTalosctl,
		Config:     string(talosctl),
	})

	err = h.configDB.InsertClusterConfig(ctx, repository.InsertClusterConfigParams{
		ClusterUuid: uuid.MustParse(cluster.Uuid),
		ConfigType:  string(talosconfig.ConfigTypeTalosctl),
		Config:      string(talosctl),
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("Could not save talosctl config", err)
	}

	resp.Body = append(resp.Body, cluster)

	return resp, nil
}

func (h *Handler) ClusterAttachHost(ctx context.Context, input *ClusterAttachHostInput) (*struct{}, error) {

	err := h.configDB.AttachHostCluster(ctx, repository.AttachHostClusterParams{
		ClusterUuid: input.ClusterUUID,
		HostUuid:    input.HostUUID,
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Could not add Cluster Host relation", err)
	}
	return nil, nil
}

package hostedcp

import (
	configv1 "github.com/openshift/api/config/v1"
	configv1alpha1 "github.com/openshift/api/config/v1alpha1"
	mcfgv1 "github.com/openshift/api/machineconfiguration/v1"
	operatorsv1alpha1 "github.com/openshift/api/operator/v1alpha1"
)

type IgnitionPayloadRequest struct {
	PullSecret  []byte `json:"pullSecret"`
	CloudConfig []byte `json:"cloudConfig"`

	AdditionalTrustBundle []byte `json:"additionalTrustBundle"`
	CloudProviderCABundle []byte `json:"cloudProviderCABundle"`
	MCSCABundle           []byte `json:"mcsCABundle"`
	KubeAPIServerCABundle []byte `json:"kubeAPIServerCABundle"`

	Cluster  ClusterConfiguration  `json:"cluster"`
	NodePool NodePoolConfiguration `json:"nodePool"`
}

type NodePoolConfiguration struct {
	KubeletConfigs          []*mcfgv1.KubeletConfig          `json:"kubeletConfigs"`
	MachineConfigs          []*mcfgv1.MachineConfig          `json:"machineConfigs"`
	ContainerRuntimeConfigs []*mcfgv1.ContainerRuntimeConfig `json:"containerRuntimeConfigs"`
}

type ClusterConfiguration struct {
	Proxy              *configv1.Proxy                    `json:"proxy"`
	Infrastructure     *configv1.Infrastructure           `json:"infrastructure"`
	Network            *configv1.Network                  `json:"network"`
	DNS                *configv1.DNS                      `json:"dns"`
	ClusterImagePolicy *configv1alpha1.ClusterImagePolicy `json:"clusterImagePolicy"`
	FeatureGate        *configv1.FeatureGate              `json:"featureGate"`
	Image              *configv1.Image                    `json:"image"`

	ImageContentSourcePolicies []*operatorsv1alpha1.ImageContentSourcePolicy `json:"imageContentSourcePolicies"`
	ImageDigestMirrorSets      []*configv1.ImageDigestMirrorSet              `json:"imageDigestMirrorSets"`
	ImageTagMirrorSets         []*configv1.ImageTagMirrorSet                 `json:"imageTagMirrorSets"`
}

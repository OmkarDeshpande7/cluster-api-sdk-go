package infrastructure

import (
	"context"
	"fmt"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
)

type CreateAWSClusterInput struct {
	Name, Namespace string

	// AWS region to host machines
	Region string `json:"region"`

	// Available zones to be considered for hosting machines
	AZs []string `json:"azs"`

	// AWS VPC ID
	VPCID string `json:"vpcid,omitempty"`

	// AWS VPC CIDR used to created new VPC for hosting machines
	// Default 10.0.0.0/16
	VPCCIDR string `json:"cidr,omitempty" default:"10.0.0.0/16"`

	// List of subnets spec to be consider to host machines.
	Subnets []awsv2.SubnetSpec `json:"subnets"`

	// Deploy spot barmetal nodes
	EnableSpot bool `json:"enableSpot,omitempty"`

	// Bidding price of spot machines
	SpotBidPrice string `json:"spotBidPrice,omitempty"`

	// ContainerCIDR configures CIDR for the BareMetal cluster containers
	ContainerCIDR string `json:"containerCIDR,omitempty"`

	// ServiceCIDR configures CIDR for the BareMetal cluster services
	ServiceCIDR string `json:"serviceCIDR,omitempty"`

	// +kubebuilder:default:="10.0.2.0/24"
	// +optional
	// VmNetworkCIDR specifies CIDR for internal VM network
	VmNetworkCIDR string `json:"vmNetworkCIDR,omitempty"`
}

type RootDiskConfig struct {
	// Device name
	// +optional
	DeviceName string `json:"deviceName,omitempty"`

	// Size specifies size (in Gi) of the storage device.
	// Must be greater than the image snapshot size or 8 (whichever is greater).
	// +kubebuilder:validation:Minimum=8
	Size int64 `json:"size"`

	// Type is the type of the volume (e.g. gp2, io1, etc...).
	// +optional
	Type string `json:"type,omitempty"`

	// IOPS is the number of IOPS requested for the disk. Not applicable to all types.
	// +optional
	IOPS int64 `json:"iops,omitempty"`

	// Throughput to provision in MiB/s supported for the volume type. Not applicable to all types.
	// +optional
	Throughput *int64 `json:"throughput,omitempty"`

	// Encrypted is whether the volume should be encrypted or not.
	// +optional
	Encrypted *bool `json:"encrypted,omitempty"`

	// EncryptionKey is the KMS key to use to encrypt the volume. Can be either a KMS key ID or ARN.
	// If Encrypted is set and this is omitted, the default AWS key will be used.
	// The key must already exist and be accessible by the controller.
	// +optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
}
type SubnetSpec struct {
	// AWS Subnet ID associated to a VPC
	ID string `json:"id,omitempty"`

	// AWS Subnet is private or not.
	IsPrivate bool `json:"isPrivate"`

	// AZ this subnet belongs to
	AZ string `json:"az"`

	// CIDR of subnet, required for create VPC usecase.
	CIDR string `json:"cidr,omitempty"`
}

func (c CreateAWSClusterInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraCluster(ctx context.Context, input infrastructure.CreateInfraClusterInput) error {
	awsInput, ok := input.(CreateAWSClusterInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateInfraCluster, input is not type '%s'", TypeCreateAWSClusterInput)
	}
	awsCluster := &awsv2.AWSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.Name,
			Namespace: awsInput.Namespace,
		},
		Spec: awsv2.AWSClusterSpec{
			NetworkSpec: awsv2.NetworkSpec{
				VPC: awsv2.VPCSpec{
					ID: awsInput.VPCID,
				},
				Subnets: awsInput.Subnets,
				CNI: &awsv2.CNISpec{
					CNIIngressRules: awsv2.CNIIngressRules{},
				},
				SecurityGroupOverrides:             map[awsv2.SecurityGroupRole]string{},
				AdditionalControlPlaneIngressRules: []awsv2.IngressRule{},
			},
			Region: awsInput.Region,
		},
	}
	err := a.Client.Create(ctx, awsCluster)
	if err != nil {
		return err
	}
	return nil
}

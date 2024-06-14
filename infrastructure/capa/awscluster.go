package infrastructure

import (
	"context"
	"fmt"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
)

type GetAWSClusterInput struct {
	Name, Namespace string
}
type CreateAWSClusterInput struct {
	Name, Namespace string

	// AWS region to host machines
	Region string `json:"region"`

	// Available zones to be considered for hosting machines
	AZs []string `json:"azs,omitempty"`

	// AWS VPC ID
	VPCID string `json:"vpcid,omitempty"`

	// AWS VPC CIDR used to created new VPC for hosting machines
	// Default 10.0.0.0/16
	VPCCIDR string `json:"cidr,omitempty" default:"10.0.0.0/16"`

	// List of subnets spec to be consider to host machines.
	Subnets []awsv2.SubnetSpec `json:"subnets"`

	SecurityGroupOverrides map[awsv2.SecurityGroupRole]string

	AdditionalControlPlaneIngressRules []awsv2.IngressRule
	// SSH key to exec into machine
	SSHKey string `json:"sshKey"`

	AdditionalLabels map[string]string

	ControlPlaneLoadBalancer *awsv2.AWSLoadBalancerSpec

	IdentityRef *awsv2.AWSIdentityReference
}

type DeleteAWSClusterInput struct {
	Name, Namespace string
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

func (c GetAWSClusterInput) GetName() string {
	return c.Name
}

func (c DeleteAWSClusterInput) GetName() string {
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
			Region:     awsInput.Region,
			SSHKeyName: &awsInput.SSHKey,
		},
	}

	if awsInput.VPCID != "" {
		awsCluster.Spec.NetworkSpec.VPC = awsv2.VPCSpec{
			ID: awsInput.VPCID,
		}
	}

	awsCluster.Spec.IdentityRef = awsInput.IdentityRef

	awsCluster.Spec.ControlPlaneLoadBalancer = awsInput.ControlPlaneLoadBalancer

	if len(awsInput.Subnets) > 0 {
		awsCluster.Spec.NetworkSpec.Subnets = awsInput.Subnets
	}

	if len(awsInput.SecurityGroupOverrides) > 0 {
		awsCluster.Spec.NetworkSpec.SecurityGroupOverrides = awsInput.SecurityGroupOverrides
	}

	if len(awsInput.AdditionalControlPlaneIngressRules) > 0 {
		awsCluster.Spec.NetworkSpec.AdditionalControlPlaneIngressRules = awsInput.AdditionalControlPlaneIngressRules
	}

	if len(awsInput.AdditionalLabels) > 0 {
		awsCluster.ObjectMeta.Labels = make(map[string]string)
		for key, value := range awsInput.AdditionalLabels {
			awsCluster.ObjectMeta.Labels[key] = value
		}
	}

	err := a.Client.Create(ctx, awsCluster)
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSProvider) GetInfraCluster(ctx context.Context, input GetAWSClusterInput) (*awsv2.AWSCluster, error) {
	awsCluster := &awsv2.AWSCluster{}
	err := c.Client.Get(ctx, types.NamespacedName{
		Name:      input.Name,
		Namespace: input.Namespace,
	}, awsCluster)
	if err != nil {
		return nil, err
	}
	return awsCluster, nil
}

func (c *AWSProvider) DeleteInfraCluster(ctx context.Context, input DeleteAWSClusterInput) error {
	awsCluster := &awsv2.AWSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: input.Namespace,
		},
	}

	err := c.Client.Delete(ctx, awsCluster)
	if err != nil {
		return err
	}
	return nil
}

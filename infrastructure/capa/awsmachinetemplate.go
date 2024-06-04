package infrastructure

import (
	"context"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
	"k8s.io/apimachinery/pkg/util/sets"
)

type CreateAWSMachineTemplateInput struct {
	Name, Namespace string
	// AWS EC2 Instance type flavor/type ex: m5zn.metal
	InstanceType string `json:"instanceType"`

	AMIID string `json:"id"`
	// IAM role for instance/machine.
	IAMInstanceProfile string `json:"iamInstanceProfile"`

	// To enable/disable swap while bootstrapping instance to pf9.
	SwapEnabled bool `json:"swapEnabled"`

	// Root disk type and size to attach to instances.
	RootDisk RootDiskConfig `json:"rootDisk"`

	// +optional
	NonRootDisk []RootDiskConfig `json:"nonRootVolumes,omitempty"`

	// SSH key to exec into machine
	SSHKey string `json:"sshKey"`
}

// VolumeType describes the EBS volume type.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html
type VolumeType string

var (
	// VolumeTypeIO1 is the string representing a provisioned iops ssd io1 volume.
	VolumeTypeIO1 = VolumeType("io1")

	// VolumeTypeIO2 is the string representing a provisioned iops ssd io2 volume.
	VolumeTypeIO2 = VolumeType("io2")

	// VolumeTypeGP2 is the string representing a general purpose ssd gp2 volume.
	VolumeTypeGP2 = VolumeType("gp2")

	// VolumeTypeGP3 is the string representing a general purpose ssd gp3 volume.
	VolumeTypeGP3 = VolumeType("gp3")

	// VolumeTypesGP are volume types provisioned for general purpose io.
	VolumeTypesGP = sets.NewString(
		string(VolumeTypeIO1),
		string(VolumeTypeIO2),
	)

	// VolumeTypesProvisioned are volume types provisioned for high performance io.
	VolumeTypesProvisioned = sets.NewString(
		string(VolumeTypeIO1),
		string(VolumeTypeIO2),
	)
)

func (c *CreateAWSMachineTemplateInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraMachineTemplate(ctx context.Context, input *infrastructure.CreateInfraMachineTemplateInput) error {
	return nil
}

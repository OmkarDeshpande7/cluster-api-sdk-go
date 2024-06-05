package infrastructure

import (
	"context"
	"fmt"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/ptr"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
)

type CreateAWSMachineTemplateInput struct {
	Name, Namespace string
	// AWS EC2 Instance type flavor/type ex: m5zn.metal
	InstanceType string `json:"instanceType"`

	AMIID *string `json:"id"`
	// IAM role for instance/machine.
	IAMInstanceProfile string `json:"iamInstanceProfile"`

	// To enable/disable swap while bootstrapping instance to pf9.
	SwapEnabled bool `json:"swapEnabled"`

	// Root disk type and size to attach to instances.
	RootDisk RootDiskConfig `json:"rootDisk"`

	// +optional
	NonRootDisk []awsv2.Volume `json:"nonRootVolumes,omitempty"`

	// SSH key to exec into machine
	SSHKey string `json:"sshKey"`

	// Subnet
	Subnet awsv2.SubnetSpec `json:"subnet"`
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

func (c CreateAWSMachineTemplateInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraMachineTemplate(ctx context.Context, input infrastructure.CreateInfraMachineTemplateInput) error {
	awsInput, ok := input.(CreateAWSMachineTemplateInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateInfraMachineTemplate, input is not type '%s'", TypeCreateAWSMachineTemplateInput)
	}
	awsMachineTemplate := awsv2.AWSMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.Name,
			Namespace: awsInput.Namespace,
		},
		Spec: awsv2.AWSMachineTemplateSpec{
			Template: awsv2.AWSMachineTemplateResource{
				Spec: awsv2.AWSMachineSpec{
					AMI: awsv2.AMIReference{
						ID: awsInput.AMIID,
					},
					IAMInstanceProfile: awsInput.IAMInstanceProfile,
					InstanceType:       awsInput.InstanceType,
					SSHKeyName:         &awsInput.SSHKey,
					Subnet: &awsv2.AWSResourceReference{
						ID: ptr.To[string](awsInput.Subnet.ID),
					},
					RootVolume: &awsv2.Volume{
						Size: awsInput.RootDisk.Size,
						Type: awsv2.VolumeType(awsInput.RootDisk.Type),
						IOPS: awsInput.RootDisk.IOPS,
					},
					NonRootVolumes: awsInput.NonRootDisk,
				},
			},
		},
	}
	err := a.Client.Create(ctx, &awsMachineTemplate)
	if err != nil {
		return err
	}
	return nil
}

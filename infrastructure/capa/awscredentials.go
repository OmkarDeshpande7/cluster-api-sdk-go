package infrastructure

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSClusterStaticIdentityInput struct {
	Name, Namespace, SecretRef, AccessKeyID, SecretAccessKey string
}

func (c CreateAWSClusterStaticIdentityInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraStaticIdentity(ctx context.Context, input infrastructure.CreateInfraClusterStaticIdentityInput) error {
	awsInput, ok := input.(CreateAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateInfraStaticIdentity, input is not type '%s'", TypeCreateAWSClusterStaticIdentity)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.SecretRef,
			Namespace: NamespaceCAPASystem,
		},
		Data: map[string][]byte{
			KeyAccessKeyID:     []byte(awsInput.AccessKeyID),
			KeySecretAccessKey: []byte(awsInput.SecretAccessKey),
		},
	}
	err := a.Client.Create(ctx, secret)
	if err != nil {
		return err
	}

	awssi := &awsv2.AWSClusterStaticIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.Name,
			Namespace: awsInput.Namespace,
		},
		Spec: awsv2.AWSClusterStaticIdentitySpec{
			AWSClusterIdentitySpec: awsv2.AWSClusterIdentitySpec{},
			SecretRef:              awsInput.SecretRef,
		},
	}

	err = a.Client.Create(ctx, awssi)
	if err != nil {
		return err
	}
	return nil
}

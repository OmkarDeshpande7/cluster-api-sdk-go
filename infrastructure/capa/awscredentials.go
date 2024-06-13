package infrastructure

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSClusterStaticIdentityInput struct {
	Name, Namespace, SecretRef, AccessKeyID, SecretAccessKey string
	MatchLabels                                              map[string]string
}

type DeleteAWSClusterStaticIdentityInput struct {
	Name, Namespace string
}

func (c CreateAWSClusterStaticIdentityInput) GetName() string {
	return c.Name
}

func (c DeleteAWSClusterStaticIdentityInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraStaticIdentity(ctx context.Context, input infrastructure.CreateInfraClusterStaticIdentityInput) error {
	awsInput, ok := input.(CreateAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateInfraStaticIdentity, input is not type '%s'", TypeCreateAWSClusterStaticIdentity)
	}
	awssi := &awsv2.AWSClusterStaticIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.Name,
			Namespace: awsInput.Namespace,
		},
		Spec: awsv2.AWSClusterStaticIdentitySpec{
			AWSClusterIdentitySpec: awsv2.AWSClusterIdentitySpec{
				AllowedNamespaces: &awsv2.AllowedNamespaces{
					Selector: metav1.LabelSelector{
						MatchLabels: awsInput.MatchLabels,
					},
				},
			},
			SecretRef: awsInput.SecretRef,
		},
	}

	err := a.Client.Create(ctx, awssi)
	if err != nil {
		return err
	}
	return nil
}

func (a *AWSProvider) DeleteInfraStaticIdentity(ctx context.Context, input infrastructure.DeleteInfraClusterStaticIdentityInput) error {
	awsInput, ok := input.(DeleteAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to DeleteInfraStaticIdentity, input is not type '%s'", TypeDeleteAWSClusterStaticIdentity)
	}

	awssi := &awsv2.AWSClusterStaticIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.Name,
			Namespace: awsInput.Namespace,
		},
	}

	err := a.Client.Delete(ctx, awssi)
	if err != nil {
		return err
	}
	return nil
}

func (a *AWSProvider) CreateSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterStaticIdentityInput) error {
	awsInput, ok := input.(CreateAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateSecretForAWSSI, input is not type '%s'", TypeCreateAWSClusterStaticIdentity)
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
	return nil
}

func (a *AWSProvider) DeleteSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterStaticIdentityInput) error {
	awsInput, ok := input.(DeleteAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to DeleteSecretForAWSSI, input is not type '%s'", TypeDeleteAWSClusterStaticIdentity)
	}
	awssi := &awsv2.AWSClusterStaticIdentity{}

	err := a.Client.Get(ctx, types.NamespacedName{
		Name:      awsInput.Name,
		Namespace: awsInput.Namespace,
	}, awssi)
	if err != nil {
		return err
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awssi.Spec.SecretRef,
			Namespace: NamespaceCAPASystem,
		},
	}

	err = a.Client.Delete(ctx, secret)
	if err != nil {
		return err
	}
	return nil
}

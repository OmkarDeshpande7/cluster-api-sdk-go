package infrastructure

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"

	"github.com/platform9/cluster-api-sdk-go/infrastructure"
)

type CreateAWSClusterStaticIdentityInput struct {
	Name, Namespace, SecretName, AccessKeyID, SecretAccessKey string
	MatchLabels                                               map[string]string
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

type CreateAWSClusterRoleIdentityInput struct {
	Name, Namespace, ExternalID, RoleARN, InlinePolicy, SessionName string
	DurationSeconds                                                 int32
	AllowedNamespaces, PolicyARNs                                   []string
	LabelSelectors                                                  metav1.LabelSelector
	SourceIdentityRef                                               *awsv2.AWSIdentityReference
}

func (c CreateAWSClusterRoleIdentityInput) GetName() string {
	return c.Name
}

type DeleteAWSClusterRoleIdentityInput struct {
	Name, Namespace string
}

func (d DeleteAWSClusterRoleIdentityInput) GetName() string {
	return d.Name
}

type GetAWSClusterRoleIdentityInput struct {
	Name, Namespace string
}

func (d GetAWSClusterRoleIdentityInput) GetName() string {
	return d.Name
}

func (a *AWSProviderImpl) CreateInfraStaticIdentity(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error {
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
			SecretRef: awsInput.SecretName,
		},
	}

	err := a.Client.Create(ctx, awssi)
	if err != nil {
		return err
	}
	return nil
}

func (a *AWSProviderImpl) DeleteInfraStaticIdentity(ctx context.Context, input infrastructure.DeleteInfraClusterIdentityInput) error {
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

func (a *AWSProviderImpl) CreateSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error {
	awsInput, ok := input.(CreateAWSClusterStaticIdentityInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateSecretForAWSSI, input is not type '%s'", TypeCreateAWSClusterStaticIdentity)
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      awsInput.SecretName,
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

func (a *AWSProviderImpl) DeleteSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error {
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

func (a *AWSProviderImpl) CreateClusterRoleIdentity(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error {
	roleInput, ok := input.(CreateAWSClusterRoleIdentityInput)
	if !ok {
		return fmt.Errorf("cannot convert to CreateAWSClusterRoleIdentityInput: %w", infrastructure.ErrInvalidParameterType)
	}
	awsRi := awsv2.AWSClusterRoleIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleInput.Name,
			Namespace: roleInput.Namespace,
		},
		Spec: awsv2.AWSClusterRoleIdentitySpec{
			AWSClusterIdentitySpec: awsv2.AWSClusterIdentitySpec{
				AllowedNamespaces: &awsv2.AllowedNamespaces{
					NamespaceList: roleInput.AllowedNamespaces,
					Selector:      roleInput.LabelSelectors,
				},
			},
			AWSRoleSpec: awsv2.AWSRoleSpec{
				RoleArn:         roleInput.RoleARN,
				SessionName:     roleInput.SessionName,
				DurationSeconds: roleInput.DurationSeconds,
				InlinePolicy:    roleInput.InlinePolicy,
				PolicyARNs:      roleInput.PolicyARNs,
			},
			ExternalID:        roleInput.ExternalID,
			SourceIdentityRef: roleInput.SourceIdentityRef,
		},
	}
	if err := a.Client.Create(ctx, &awsRi); err != nil {
		return err
	}
	return nil
}

func (a *AWSProviderImpl) DeleteClusterRoleIdentity(ctx context.Context, input infrastructure.DeleteInfraClusterIdentityInput) error {
	deleteRoleInput, ok := input.(DeleteAWSClusterRoleIdentityInput)
	if !ok {
		return fmt.Errorf("cannot convert to DeleteAWSClusterRoleIdentityInput: %w", infrastructure.ErrInvalidParameterType)
	}
	return a.Client.Delete(ctx, &awsv2.AWSClusterRoleIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deleteRoleInput.Name,
			Namespace: deleteRoleInput.Namespace,
		},
	})
}

func (a *AWSProviderImpl) GetClusterRoleIdentity(ctx context.Context, input infrastructure.GetInfraClusterIdentityInput) (*awsv2.AWSClusterRoleIdentity, error) {
	getRoleInput, ok := input.(GetAWSClusterRoleIdentityInput)
	if !ok {
		return nil, fmt.Errorf("cannot convert to GetAWSClusterRoleIdentityInput: %w", infrastructure.ErrInvalidParameterType)
	}
	obj := &awsv2.AWSClusterRoleIdentity{}
	if err := a.Client.Get(ctx, types.NamespacedName{
		Namespace: getRoleInput.Namespace,
		Name:      getRoleInput.Name,
	}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

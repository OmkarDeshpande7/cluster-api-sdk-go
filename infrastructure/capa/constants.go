package infrastructure

import (
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
)

const (
	TypeCreateAWSClusterInput          = "CreateAWSClusterInput"
	TypeCreateAWSClusterStaticIdentity = "CreateAWSClusterStaticIdentity"
	TypeCreateAWSMachineTemplateInput  = "CreateAWSMachineTemplateInput"

	NamespaceCAPASystem = "capa-system"
	KeyAccessKeyID      = "AccessKeyID"
	KeySecretAccessKey  = "SecretAccessKey"

	TypeDeleteAWSClusterInput          = "DeleteAWSClusterInput"
	TypeDeleteAWSClusterStaticIdentity = "DeleteAWSClusterStaticIdentity"
	TypeDeleteAWSMachineTemplateInput  = "DeleteAWSMachineTemplateInput"
)

var (
	KindAWSClusterStaticIdentity = string(awsv2.ClusterStaticIdentityKind)
)

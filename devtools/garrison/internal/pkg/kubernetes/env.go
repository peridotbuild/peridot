package kubernetes

import (
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvironmentVariables struct {
	// Stage is the environment the current deployment will run in
	Stage starlark.Stage

	// Kubernetes-specific

	// Namespace is the Kubernetes namespace to deploy to
	Namespace string

	// ServiceAccountName is the Kubernetes service account to use
	ServiceAccountName string

	// Cloud-specific

	// AwsRegion is the AWS region to use as default
	AwsRegion string

	// Dev-specific

	// LocalstackEndpoint is the endpoint to use for localstack
	LocalstackEndpoint string
}

func GetEnvironmentVariables(state *starlark.State, md *metav1.ObjectMeta) *EnvironmentVariables {
	defaultAwsRegion := "us-east-2"
	if state.Default.AwsRegion != nil {
		defaultAwsRegion = *state.Default.AwsRegion
	}

	localstackEndpoint := "http://localstack.default.svc.cluster.local:4566"
	if state.Stage != starlark.StageDevelopment {
		localstackEndpoint = ""
	}

	return &EnvironmentVariables{
		Stage:              state.Stage,
		Namespace:          md.Namespace,
		ServiceAccountName: md.Name,
		AwsRegion:          defaultAwsRegion,
		LocalstackEndpoint: localstackEndpoint,
	}
}

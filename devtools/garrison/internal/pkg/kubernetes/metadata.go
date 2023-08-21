package kubernetes

import (
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *State) getMetadata(service *starlark.Service) *metav1.ObjectMeta {
	ns := ""
	// If the defaults have a namespace, that is the first preference.
	if s.Namespace != nil {
		ns = *s.Namespace
	}
	// If the service has a namespace, that is the second preference.
	if service.Namespace != nil {
		ns = *service.Namespace
	}
	// If neither, use the service name.
	if ns == "" {
		ns = *service.Name
	}

	// Append stage to the namespace.
	ns = ns + "-" + string(s.Stage)

	return &metav1.ObjectMeta{
		Name:      *service.Name,
		Namespace: ns,
	}
}

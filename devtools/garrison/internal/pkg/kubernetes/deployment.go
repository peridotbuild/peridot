package kubernetes

import (
	"bytes"
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (s *State) NewDeployment(service *starlark.Service) error {
	md := s.getMetadata(service)

	podLabels := map[string]string{
		"app":   *service.Name,
		"stage": string(s.Stage),
	}

	// Copy the metadata and add the labels.
	podMd := *md
	podMd.Labels = podLabels

	imagePullPolicy := v1.PullIfNotPresent

	affinity := &v1.Affinity{
		PodAntiAffinity: &v1.PodAntiAffinity{
			PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
				{
					Weight: int32(99),
					PodAffinityTerm: v1.PodAffinityTerm{
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      "app",
									Operator: metav1.LabelSelectorOpIn,
									Values:   []string{*service.Name},
								},
							},
						},
						TopologyKey: "kubernetes.io/hostname",
					},
				},
				{
					Weight: int32(100),
					PodAffinityTerm: v1.PodAffinityTerm{
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      "app",
									Operator: metav1.LabelSelectorOpIn,
									Values:   []string{*service.Name},
								},
							},
						},
						TopologyKey: "topology.kubernetes.io/zone",
					},
				},
			},
		},
	}
	if service.NoAntiAffinity != nil && *service.NoAntiAffinity {
		affinity = nil
	}

	app := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: *md,
		Spec: appsv1.DeploymentSpec{
			RevisionHistoryLimit: base.Pointer(int32(15)),
			Selector: &metav1.LabelSelector{
				MatchLabels: podLabels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: podMd,
				Spec: v1.PodSpec{
					// For most use cases this is acceptable for RESF
					AutomountServiceAccountToken: base.Pointer(true),
					SecurityContext: &v1.PodSecurityContext{
						// Volume ownership is set to 1000:1000 to match the default
						FSGroup: base.Pointer(int64(1000)),
					},
					Affinity: affinity,
					Containers: []v1.Container{
						{
							Name:            *service.Name,
							ImagePullPolicy: imagePullPolicy,
							Image:           *service.Image,
							Command:         service.Command,
							Args:            service.Args,
						},
					},
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					// We don't want any unavailable pods during the update.
					MaxUnavailable: base.Pointer(intstr.FromInt32(0)),
					// Allow surging up to 300% of the desired replicas.
					MaxSurge: base.Pointer(intstr.FromString("300%")),
				},
			},
		},
	}

	// Add the ports
	if service.Ports != nil {
		for _, port := range service.Ports {
			nPort := v1.ContainerPort{
				Name:          *port.Name,
				ContainerPort: int32(*port.ContainerPort),
			}
			app.Spec.Template.Spec.Containers[0].Ports = append(app.Spec.Template.Spec.Containers[0].Ports, nPort)
		}
	}

	// Add the environment variables
	if service.Env != nil {
		for k, v := range service.Env {
			nEnv := v1.EnvVar{
				Name:  k,
				Value: v,
			}
			app.Spec.Template.Spec.Containers[0].Env = append(app.Spec.Template.Spec.Containers[0].Env, nEnv)
		}
	}

	// Add readiness/liveness probes
	if service.Health != nil {
		if service.Health.Http != nil {
			probe := v1.Probe{
				ProbeHandler: v1.ProbeHandler{
					HTTPGet: &v1.HTTPGetAction{
						Path: service.Health.Http.Path,
						Port: intstr.FromInt32(int32(service.Health.Http.Port)),
					},
				},
				InitialDelaySeconds: int32(*service.Health.InitialDelaySeconds),
				PeriodSeconds:       int32(*service.Health.PeriodSeconds),
				TimeoutSeconds:      int32(*service.Health.TimeoutSeconds),
				SuccessThreshold:    int32(*service.Health.SuccessThreshold),
				FailureThreshold:    int32(*service.Health.FailureThreshold),
			}
			app.Spec.Template.Spec.Containers[0].ReadinessProbe = &probe
			app.Spec.Template.Spec.Containers[0].LivenessProbe = &probe
		}

		if service.Health.Grpc != nil {
			probe := v1.Probe{
				ProbeHandler: v1.ProbeHandler{
					Exec: &v1.ExecAction{
						Command: []string{
							"grpc_health_probe",
							"-connect-timeout=4s",
							"-v",
							fmt.Sprintf("-addr=:%d", service.Health.Grpc.Port),
						},
					},
				},
			}

			app.Spec.Template.Spec.Containers[0].ReadinessProbe = &probe
			app.Spec.Template.Spec.Containers[0].LivenessProbe = &probe
		}
	}

	// Create yaml buffer
	var yamlBuffer bytes.Buffer
	err := s.printer.PrintObj(app, &yamlBuffer)
	if err != nil {
		return err
	}

	// Write yaml buffer to manifests
	s.manifests = append(s.manifests, yamlBuffer.String())

	return nil
}

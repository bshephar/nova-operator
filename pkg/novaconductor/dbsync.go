/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package novaconductor

import (
	novav1 "github.com/openstack-k8s-operators/nova-operator/api/v1beta1"
	"github.com/openstack-k8s-operators/nova-operator/pkg/nova"

	env "github.com/openstack-k8s-operators/lib-common/modules/common/env"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

// CellDBSyncJob - define a batchv1.Job to be run to apply the cel DB schema
func CellDBSyncJob(
	instance *novav1.NovaConductor,
	labels map[string]string,
	annotations map[string]string,
) *batchv1.Job {
	args := []string{"-c", nova.KollaServiceCommand}

	envVars := map[string]env.Setter{}
	envVars["KOLLA_CONFIG_STRATEGY"] = env.SetValue("COPY_ALWAYS")
	envVars["KOLLA_BOOTSTRAP"] = env.SetValue("true")

	envVars["CELL_NAME"] = env.SetValue(instance.Spec.CellName)

	env := env.MergeEnvs([]corev1.EnvVar{}, envVars)

	// create Volume and VolumeMounts
	volumes := []corev1.Volume{
		nova.GetConfigVolume(nova.GetServiceConfigSecretName(instance.Name)),
		nova.GetScriptVolume(nova.GetScriptSecretName(instance.Name)),
	}
	volumeMounts := []corev1.VolumeMount{
		nova.GetConfigVolumeMount(),
		nova.GetScriptVolumeMount(),
		nova.GetKollaConfigVolumeMount("nova-conductor-dbsync"),
	}

	// add CA cert if defined
	if instance.Spec.TLS.CaBundleSecretName != "" {
		volumes = append(volumes, instance.Spec.TLS.CreateVolume())
		volumeMounts = append(volumeMounts, instance.Spec.TLS.CreateVolumeMounts(nil)...)
	}

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name + "-db-sync",
			Namespace:   instance.Namespace,
			Annotations: annotations,
			Labels:      labels,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy:      corev1.RestartPolicyOnFailure,
					ServiceAccountName: instance.Spec.ServiceAccount,
					Volumes:            volumes,
					Containers: []corev1.Container{
						{
							Name: instance.Name + "-db-sync",
							Command: []string{
								"/bin/bash",
							},
							Args:  args,
							Image: instance.Spec.ContainerImage,
							SecurityContext: &corev1.SecurityContext{
								RunAsUser: ptr.To(nova.NovaUserID),
							},
							Env:          env,
							VolumeMounts: volumeMounts,
						},
					},
				},
			},
		},
	}

	if instance.Spec.NodeSelector != nil {
		job.Spec.Template.Spec.NodeSelector = *instance.Spec.NodeSelector
	}

	return job
}

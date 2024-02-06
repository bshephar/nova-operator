/*
Copyright 2023.

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

//
// Generated by:
//
// operator-sdk create webhook --group nova --version v1beta1 --kind NovaScheduler --programmatic-validation --defaulting
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// NovaSchedulerDefaults -
type NovaSchedulerDefaults struct {
	ContainerImageURL string
}

var novaSchedulerDefaults NovaSchedulerDefaults

// log is for logging in this package.
var novaschedulerlog = logf.Log.WithName("novascheduler-resource")

// SetupNovaSchedulerDefaults - initialize NovaScheduler spec defaults for use with either internal or external webhooks
func SetupNovaSchedulerDefaults(defaults NovaSchedulerDefaults) {
	novaSchedulerDefaults = defaults
	novaschedulerlog.Info("NovaScheduler defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *NovaScheduler) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-nova-openstack-org-v1beta1-novascheduler,mutating=true,failurePolicy=fail,sideEffects=None,groups=nova.openstack.org,resources=novaschedulers,verbs=create;update,versions=v1beta1,name=mnovascheduler.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &NovaScheduler{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *NovaScheduler) Default() {
	novaschedulerlog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this NovaScheduler spec
func (spec *NovaSchedulerSpec) Default() {
	if spec.ContainerImage == "" {
		spec.ContainerImage = novaSchedulerDefaults.ContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-nova-openstack-org-v1beta1-novascheduler,mutating=false,failurePolicy=fail,sideEffects=None,groups=nova.openstack.org,resources=novaschedulers,verbs=create;update,versions=v1beta1,name=vnovascheduler.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &NovaScheduler{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *NovaScheduler) ValidateCreate() (admission.Warnings, error) {
	novaschedulerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *NovaScheduler) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	novaschedulerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *NovaScheduler) ValidateDelete() (admission.Warnings, error) {
	novaschedulerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

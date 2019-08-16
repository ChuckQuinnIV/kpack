/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	"github.com/knative/pkg/apis"
	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	"github.com/knative/pkg/kmeta"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildSpec   `json:"spec"`
	Status BuildStatus `json:"status"`
}

var (
	_ apis.Validatable   = (*Build)(nil)
	_ apis.Defaultable   = (*Build)(nil)
	_ kmeta.OwnerRefable = (*Build)(nil)
)

type BuildSpec struct {
	Tag                  string                      `json:"tag"`
	BuilderRef           string                      `json:"builderRef"`
	ServiceAccount       string                      `json:"serviceAccount"`
	Source               SourceConfig                `json:"source"`
	CacheName            string                      `json:"cacheName"`
	AdditionalImageNames []string                    `json:"additionalImageNames"`
	Env                  []corev1.EnvVar             `json:"env"`
	Resources            corev1.ResourceRequirements `json:"resources"`
}

type BuildStatus struct {
	duckv1alpha1.Status `json:",inline"`
	BuildMetadata       BuildpackMetadataList   `json:"buildMetadata"`
	LatestImage         string                  `json:"latestImage"`
	PodName             string                  `json:"podName"`
	StepStates          []corev1.ContainerState `json:"stepStates,omitempty"`
	StepsCompleted      []string                `json:"stepsCompleted",omitempty`
	Builder             string                  `json:"builder"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Build `json:"items"`
}

func (*Build) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Build")
}

func (b *Build) ServiceAccount() string {
	return b.Spec.ServiceAccount
}

func (b *Build) Tag() string {
	return b.Spec.Tag
}

func (b *Build) HasSecret() bool {
	return true
}

func (b *Build) Namespace() string {
	return b.ObjectMeta.Namespace
}

func (b *Build) IsRunning() bool {
	if b == nil {
		return false
	}

	return b.Status.GetCondition(duckv1alpha1.ConditionSucceeded).IsUnknown()
}

func (b *Build) BuildRef() string {
	if b == nil {
		return ""
	}

	return b.GetName()
}

func (b *Build) BuiltImage() string {
	if b == nil {
		return ""
	}
	if !b.IsSuccess() {
		return ""
	}

	return b.Status.LatestImage
}

func (b *Build) IsSuccess() bool {
	if b == nil {
		return false
	}
	return b.Status.GetCondition(duckv1alpha1.ConditionSucceeded).IsTrue()
}

func (b *Build) IsFailure() bool {
	if b == nil {
		return false
	}
	return b.Status.GetCondition(duckv1alpha1.ConditionSucceeded).IsFalse()
}

func (b *Build) PodName() string {
	return kmeta.ChildName(b.Name, "-build-pod")
}

func (b *Build) MetadataReady(pod *corev1.Pod) bool {
	return !b.Status.GetCondition(duckv1alpha1.ConditionSucceeded).IsTrue() &&
		pod.Status.Phase == "Succeeded"
}

func (b *Build) Finished() bool {
	return !b.Status.GetCondition(duckv1alpha1.ConditionSucceeded).IsUnknown()
}

func (b *Build) BuildEnvVars() []corev1.EnvVar {
	return b.Spec.Source.Source().BuildEnvVars()
}

func (b *Build) ImagePullSecretsVolume() corev1.Volume {
	return b.Spec.Source.Source().ImagePullSecretsVolume()
}

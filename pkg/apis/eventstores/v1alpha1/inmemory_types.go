/*
Copyright (c) 2020 TriggerMesh Inc.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InMemoryStore is an event store that keeps data in memory.
type InMemoryStore struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InMemoryStoreSpec   `json:"spec"`
	Status InMemoryStoreStatus `json:"status,omitempty"`
}

// Check the interfaces InMemoryStore should be implementing.
var (
	_ runtime.Object     = (*InMemoryStore)(nil)
	_ kmeta.OwnerRefable = (*InMemoryStore)(nil)
)

// InMemoryStoreSpec holds the desired state of the InMemoryStore.
type InMemoryStoreSpec struct {
	// DefaultGlobalTTL is the ammount of seconds that the globaly stored data is retrievable.
	// +optional
	DefaultGlobalTTL *int `json:"defaultGlobalTTL,omitempty"`

	// DefaultBridgeTTL is the ammount of seconds that the bridge stored data is retrievable.
	// +optional
	DefaultBridgeTTL *int `json:"defaultBridgeTTL,omitempty"`

	// DefaultInstanceTTL is the ammount of seconds that the bridge instance stored data is retrievable.
	// +optional
	DefaultInstanceTTL *int `json:"defaultInstanceTTL,omitempty"`

	// DefaultExpiredGCPeriod is the frequency in secods that the garbage collector will clean expired items.
	// +optional
	DefaultExpiredGCPeriod *int `json:"defaultExpiredGCPeriod,omitempty"`
}

// InMemoryStoreStatus communicates the observed state of the InMemoryStore (from the controller).
type InMemoryStoreStatus struct {
	duckv1.Status `json:",inline"`
	// Address for the InMemoryStore
	Address *Addressable `json:"address,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InMemoryStoreList is a list of InMemoryStore resources
type InMemoryStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []InMemoryStore `json:"items"`
}

// Addressable host an URI location indicating the destination
// for storage requests.
type Addressable struct {
	URI *string `json:"url,omitempty"`
}

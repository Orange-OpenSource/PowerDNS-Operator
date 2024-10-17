/*
 * Software Name : PowerDNS-Operator
 *
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Apache-2.0
 *
 * This software is distributed under the Apache 2.0 License,
 * see the "LICENSE" file for more details
 */

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RRsetSpec defines the desired state of RRset
type RRsetSpec struct {
	// Type of the record (e.g. "A", "PTR", "MX").
	Type string `json:"type"`
	// DNS TTL of the records, in seconds.
	TTL uint32 `json:"ttl"`
	// All records in this Resource Record Set.
	Records []string `json:"records"`
	// Comment on RRSet.
	// +optional
	Comment *string `json:"comment,omitempty"`
	// ZoneRef reference the zone the RRSet depends on.
	ZoneRef ZoneRef `json:"zoneRef"`
}

type ZoneRef struct {
	// Name of the zone.
	Name string `json:"name"`
}

// RRsetStatus defines the observed state of RRset
type RRsetStatus struct {
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Zone",type="string",JSONPath=".spec.zoneRef.name"
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="TTL",type="integer",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="Records",type="string",JSONPath=".spec.records"
// RRset is the Schema for the rrsets API
type RRset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RRsetSpec   `json:"spec,omitempty"`
	Status RRsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RRsetList contains a list of RRset
type RRsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RRset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RRset{}, &RRsetList{})
}

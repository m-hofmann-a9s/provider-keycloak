/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RoleMapperObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleMapperParameters struct {

	// The destination client of the role. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/corewire/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The destination client scope of the role. Cannot be used at the same time as client_id.
	// +kubebuilder:validation:Optional
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// The realm id where the associated client or client scope exists.
	// +kubebuilder:validation:Required
	RealmID *string `json:"realmId" tf:"realm_id,omitempty"`

	// Id of the role to assign
	// +crossplane:generate:reference:type=github.com/corewire/provider-keycloak/apis/role/v1alpha1.Role
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

// RoleMapperSpec defines the desired state of RoleMapper
type RoleMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleMapperParameters `json:"forProvider"`
}

// RoleMapperStatus defines the observed state of RoleMapper.
type RoleMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleMapper is the Schema for the RoleMappers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type RoleMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleMapperSpec   `json:"spec"`
	Status            RoleMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleMapperList contains a list of RoleMappers
type RoleMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleMapper `json:"items"`
}

// Repository type metadata.
var (
	RoleMapper_Kind             = "RoleMapper"
	RoleMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleMapper_Kind}.String()
	RoleMapper_KindAPIVersion   = RoleMapper_Kind + "." + CRDGroupVersion.String()
	RoleMapper_GroupVersionKind = CRDGroupVersion.WithKind(RoleMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleMapper{}, &RoleMapperList{})
}

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

type HardcodedRoleMapperInitParameters struct {

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The name of the role which should be assigned to the users. Client roles should use the format {{client_id}}.{{client_role_name}}.
	// Role to grant to user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name", false)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in role to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

type HardcodedRoleMapperObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The name of the role which should be assigned to the users. Client roles should use the format {{client_id}}.{{client_role_name}}.
	// Role to grant to user.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type HardcodedRoleMapperParameters struct {

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	// +kubebuilder:validation:Optional
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The name of the role which should be assigned to the users. Client roles should use the format {{client_id}}.{{client_role_name}}.
	// Role to grant to user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name", false)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in role to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

// HardcodedRoleMapperSpec defines the desired state of HardcodedRoleMapper
type HardcodedRoleMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HardcodedRoleMapperParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HardcodedRoleMapperInitParameters `json:"initProvider,omitempty"`
}

// HardcodedRoleMapperStatus defines the observed state of HardcodedRoleMapper.
type HardcodedRoleMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HardcodedRoleMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HardcodedRoleMapper is the Schema for the HardcodedRoleMappers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type HardcodedRoleMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HardcodedRoleMapperSpec   `json:"spec"`
	Status HardcodedRoleMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HardcodedRoleMapperList contains a list of HardcodedRoleMappers
type HardcodedRoleMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HardcodedRoleMapper `json:"items"`
}

// Repository type metadata.
var (
	HardcodedRoleMapper_Kind             = "HardcodedRoleMapper"
	HardcodedRoleMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HardcodedRoleMapper_Kind}.String()
	HardcodedRoleMapper_KindAPIVersion   = HardcodedRoleMapper_Kind + "." + CRDGroupVersion.String()
	HardcodedRoleMapper_GroupVersionKind = CRDGroupVersion.WithKind(HardcodedRoleMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&HardcodedRoleMapper{}, &HardcodedRoleMapperList{})
}

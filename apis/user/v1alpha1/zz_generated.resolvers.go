/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	common "github.com/crossplane-contrib/provider-keycloak/config/common"
	apisresolver "github.com/crossplane-contrib/provider-keycloak/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Groups.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *Groups) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("group.keycloak.crossplane.io", "v1alpha1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.GroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.GroupIdsRefs,
			Selector:      mg.Spec.ForProvider.GroupIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupIds")
	}
	mg.Spec.ForProvider.GroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("user.keycloak.crossplane.io", "v1alpha1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserIDRef,
			Selector:     mg.Spec.ForProvider.UserIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserID")
	}
	mg.Spec.ForProvider.UserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("group.keycloak.crossplane.io", "v1alpha1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.GroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.GroupIdsRefs,
			Selector:      mg.Spec.InitProvider.GroupIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupIds")
	}
	mg.Spec.InitProvider.GroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.GroupIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("user.keycloak.crossplane.io", "v1alpha1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserIDRef,
			Selector:     mg.Spec.InitProvider.UserIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.UserID")
	}
	mg.Spec.InitProvider.UserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Permissions.
func (mg *Permissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Roles.
func (mg *Roles) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("role.keycloak.crossplane.io", "v1alpha1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RoleIds),
			Extract:       common.UUIDExtractor(),
			References:    mg.Spec.ForProvider.RoleIdsRefs,
			Selector:      mg.Spec.ForProvider.RoleIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleIds")
	}
	mg.Spec.ForProvider.RoleIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("user.keycloak.crossplane.io", "v1alpha1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserIDRef,
			Selector:     mg.Spec.ForProvider.UserIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserID")
	}
	mg.Spec.ForProvider.UserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("role.keycloak.crossplane.io", "v1alpha1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.RoleIds),
			Extract:       common.UUIDExtractor(),
			References:    mg.Spec.InitProvider.RoleIdsRefs,
			Selector:      mg.Spec.InitProvider.RoleIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleIds")
	}
	mg.Spec.InitProvider.RoleIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.RoleIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("user.keycloak.crossplane.io", "v1alpha1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserIDRef,
			Selector:     mg.Spec.InitProvider.UserIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.UserID")
	}
	mg.Spec.InitProvider.UserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

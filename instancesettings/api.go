// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
// Last generated at 2024-02-10 16:07:59.699276773 +0000 UTC
package instancesettings

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Update updates the instance's settings.
func Update(ctx context.Context, params *UpdateParams) error {
	return getClient().Update(ctx, params)
}

// UpdateRestrictions updates the restriction settings of the instance.
func UpdateRestrictions(ctx context.Context, params *UpdateRestrictionsParams) (*clerk.InstanceRestrictions, error) {
	return getClient().UpdateRestrictions(ctx, params)
}

// UpdateOrganizationSettings updates the organization settings of the instance.
func UpdateOrganizationSettings(ctx context.Context, params *UpdateOrganizationSettingsParams) (*clerk.OrganizationSettings, error) {
	return getClient().UpdateOrganizationSettings(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
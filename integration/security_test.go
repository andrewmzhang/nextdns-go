//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestSecurityLifecycle(t *testing.T) {
	ctx := context.Background()

	// Check that there are no lingering integration test profiles
	name := generateUniqueName()
	basicProfile := models.Profile{Name: name}
	create, err := client.Profiles().Create(ctx, basicProfile)
	defer func() {
		// Delete test
		err := client.Profiles().Bind(create.ID).Delete(ctx)
		require.NoError(t, err, "Error deleting profile id %s: %v", create.ID, err)
	}()

	// Read test
	security, err := client.Security(create.ID).Get(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, security)
	require.False(t, security.AiThreatDetection, "By default, AI Threat Detection is disabled")

	// Check that we can update values
	security.AiThreatDetection = true
	err = client.Security(create.ID).Update(ctx, security)
	require.NoError(t, err)
	security, err = client.Security(create.ID).Get(ctx)
	require.NoError(t, err)
	require.True(t, security.AiThreatDetection, "AI Threat Detection is enabled")

	// Check setting block level tlds
	blockableTlds, err := client.SecurityTlds().List(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, blockableTlds, "Blockable TLDs list is empty")
	security.BlockedTLDs = []*models.SecurityTlds{
		&models.SecurityTlds{ID: blockableTlds[0].ID},
	}
	err = client.Security(create.ID).Update(ctx, security)
	require.NoError(t, err)
	security, err = client.Security(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Equal(t, security.BlockedTLDs[0].ID, blockableTlds[0].ID, "Update failed to add TLD to blockedTlds")

	// Check proper behavior when list is nil
	security.BlockedTLDs = nil
	err = client.Security(create.ID).Update(ctx, security)
	require.NoError(t, err)
	security, err = client.Security(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Equal(t, len(security.BlockedTLDs), 1, "A nil slice should produce no update")

	// Check deleting block level tlds
	security.BlockedTLDs = []*models.SecurityTlds{}
	err = client.Security(create.ID).Update(ctx, security)
	require.NoError(t, err)
	security, err = client.Security(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Empty(t, security.BlockedTLDs, "Blockable TLDs list should be empty")
}

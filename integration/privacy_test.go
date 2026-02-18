//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestPrivacyLifecycle(t *testing.T) {

	ctx := context.Background()
	name := generateUniqueName()
	basicProfile := models.Profile{Name: name}
	create, err := client.Profiles().Create(ctx, basicProfile)
	defer func() {
		// Delete test
		err := client.Profiles().Bind(create.ID).Delete(ctx)
		require.NoError(t, err, "Error deleting profile id %s: %v", create.ID, err)
	}()
	require.NoError(t, err)

	// Check getting privacy
	privacy, err := client.Privacy(create.ID).Get(ctx)
	require.NoError(t, err)
	fmt.Println("privacy", privacy)
	require.Empty(t, privacy.Blocklists)
	require.Empty(t, privacy.Natives)
	require.False(t, *privacy.DisguisedTrackers)
	require.False(t, *privacy.AllowAffiliate)

	// Check getting all blocklist Information
	privacyList, err := client.PrivacyBlocklists().List(ctx)
	require.NoError(t, err)

	// Check setting blocklist information
	privacy.Blocklists = []*models.PrivacyBlocklists{
		&models.PrivacyBlocklists{
			ID:        privacyList[0].ID,
			Name:      "",
			Website:   "",
			Entries:   0,
			UpdatedOn: nil,
		},
		&privacyList[1],
	}
	err = client.Privacy(create.ID).Update(ctx, privacy)
	require.NoError(t, err)
	privacy, err = client.Privacy(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Equal(t, len(privacy.Blocklists), 2)

	// Check correct nil list response
	privacy.Blocklists = nil
	err = client.Privacy(create.ID).Update(ctx, privacy)
	require.NoError(t, err)
	privacy, err = client.Privacy(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Equal(t, len(privacy.Blocklists), 2)

	// Check erasing blocklists
	privacy.Blocklists = []*models.PrivacyBlocklists{}
	err = client.Privacy(create.ID).Update(ctx, privacy)
	require.NoError(t, err)
	privacy, err = client.Privacy(create.ID).Get(ctx)
	require.NoError(t, err)
	require.Empty(t, privacy.Blocklists)

	// Check privacy natives
	privacyNativesList, err := client.PrivacyNatives().List(ctx)
	require.NoError(t, err)
	require.Equal(t, len(privacyNativesList), 8)
}
func TestPrivacyErrors(t *testing.T) {
	// TODO move to unit tests
	ctx := context.Background()

	clientService := client.Privacy

	for _, badProfileId := range []string{"", "not-a-profile-id"} {
		// Read test
		security, err := clientService(badProfileId).Get(ctx)
		require.Error(t, err)
		require.Empty(t, security)

		// Update test
		err = clientService(badProfileId).Update(ctx, security)
		require.Error(t, err)
		require.Empty(t, security)
	}
}

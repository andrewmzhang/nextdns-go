//go:build integration

package integration

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func generateUniqueName() string {
	return fmt.Sprintf("it-%d", time.Now().UnixNano())
}

func isCreatedByIntegrationTest(summary models.Profile) bool {
	return strings.Contains(summary.Name, "it-")
}
func TestProfileLifecycle(t *testing.T) {
	ctx := context.Background()

	// Test creating a profile
	name := generateUniqueName()
	basicProfile := models.Profile{Name: name}
	create, err := client.Profiles().Create(ctx, basicProfile)
	require.NoError(t, err, "Failed to create profile")
	require.Equal(t, basicProfile.Name, create.Name, "Newly created profile name doesn't match basic_profile name")
	require.NotEmpty(t, create.ID, "Profile id should not be empty")
	defer func() {
		// Delete test
		err := client.Profiles().Bind(create.ID).Delete(ctx)
		require.NoError(t, err, "Error deleting profile id %s: %v", create.ID, err)
	}()

	// List test
	profiles, err := client.Profiles().List(ctx)
	require.NotEmpty(t, profiles, "Profiles should not be empty")

	// Read test
	require.NotNil(t, client.Profiles().Bind(create.ID))
	profile, err := client.Profiles().Bind(create.ID).Get(ctx)
	require.NoError(t, err, "Could not find profile id %s", create.ID)
	require.NotEmpty(t, profile)
	require.Equal(t, profile.ID, create.ID)

	// Update test
	name = generateUniqueName()
	profile.Name = name
	err = client.Profiles().Bind(profile.ID).Update(ctx, profile)
	require.NoError(t, err)

	// No-op update test
	emptyProfile := models.Profile{}
	err = client.Profiles().Bind(profile.ID).Update(ctx, emptyProfile)
	profile, err = client.Profiles().Bind(create.ID).Get(ctx)
	require.NoError(t, err, "Could not find profile id %s", create.ID)
	require.NotEmpty(t, profile.Name)
	require.NoError(t, err)
}

func TestProfileErrors(t *testing.T) {
	// TODO move to unit test
	ctx := context.Background()

	// Read test
	_, err := client.Profiles().Bind("").Get(ctx)
	require.Error(t, err, "Empty profileId should fail the Get")
}

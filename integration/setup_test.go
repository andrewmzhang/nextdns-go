//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestSetupLifecycle(t *testing.T) {
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
	setup, err := client.Setup(create.ID).Get(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, setup)
	require.Equal(t, len(setup.Ipv6), 2, "As of time of writing, there are always 2 ipv6 addresses")
	require.Equal(t, len(setup.Ipv4), 0, "As of time of writing, there are always 0 ipv4 address")
}

//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestParentalControlLifecycle(t *testing.T) {

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

	services, err := client.ParentalControlServices().List(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, services)

}
func TestParentalControlErrors(t *testing.T) {
	// TODO move to unit tests
	ctx := context.Background()

	clientService := client.ParentalControl

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

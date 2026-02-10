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

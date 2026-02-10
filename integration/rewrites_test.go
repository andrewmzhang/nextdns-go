//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestRewriteLifecycle(t *testing.T) {
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
	rewrites, err := client.Rewrites(create.ID).List(ctx)
	require.NoError(t, err)
	require.Empty(t, rewrites)

	// Create
	rewrite, err := client.Rewrites(create.ID).Create(ctx, models.Rewrite{
		Name:    "a.example.com",
		Content: "b.example.com",
	})
	require.NoError(t, err)
	require.NotEmpty(t, rewrite.ID)

	// Delete
	err = client.Rewrites(create.ID).Bind(rewrite.ID).Delete(ctx)
	require.NoError(t, err)

}

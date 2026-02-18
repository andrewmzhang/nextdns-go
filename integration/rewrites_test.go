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

	// List-get test
	rewrite, err = client.Rewrites(create.ID).Bind(rewrite.ID).Get(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, rewrite.ID)

	// Delete
	err = client.Rewrites(create.ID).Bind(rewrite.ID).Delete(ctx)
	require.NoError(t, err)

}

func TestRewriteErrors(t *testing.T) {
	// TODO move to unit tests
	ctx := context.Background()

	clientService := client.Rewrites

	for _, badProfileId := range []string{"", "not-a-profile-id"} {
		// Read test
		rewrites, err := clientService(badProfileId).List(ctx)
		require.Error(t, err)
		require.Empty(t, rewrites)

		// Create
		rewrite, err := clientService(badProfileId).Create(ctx, models.Rewrite{
			Name:    "a.example.com",
			Content: "b.example.com",
		})
		require.Error(t, err)
		require.Nil(t, rewrite)

		for _, badRewriteId := range []string{"", "not-a-rewrite-id"} {
			// Bind
			boundRewrite := clientService(badProfileId).Bind(badRewriteId)

			// List-get test
			rewrite, err = boundRewrite.Get(ctx)
			require.Error(t, err)
			require.Nil(t, rewrite)

			// Delete
			err = boundRewrite.Delete(ctx)
			require.Error(t, err)

		}
	}
}

//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/stretchr/testify/require"
)

func TestDenylistLifecycle(t *testing.T) {
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

	denylist, err := client.Denylists(create.ID).List(ctx)
	fmt.Println(denylist)
	require.NoError(t, err)
	require.NotNil(t, denylist)
	require.Equal(t, len(denylist), 0)

	denycreate, err := client.Denylists(create.ID).Create(ctx, models.Denylist{
		ID:     "example3.com",
		Active: true,
	})
	fmt.Println(denycreate)
	require.NoError(t, err)
	// Get will not work with denylist
	denyItem, err := client.Denylists(create.ID).Bind("example3.com").Get(ctx)
	require.NoError(t, err)
	require.True(t, denyItem.Active)
	fmt.Println("denyItem:", denyItem)
	// fmt.Println(denyItem.Update(ctx, models.Denylist{
	// 	Active: true,
	// }))
	// denylist, err = client.Denylist(create.ID).List(ctx)
	// require.NoError(t, err)
	// require.True(t, denylist[0].Active)
}

func TestDenylistErrors(t *testing.T) {
	// TODO move to unit tests
	ctx := context.Background()

	clientService := client.Denylists

	for _, badProfileId := range []string{"", "not-a-profile-id"} {
		// Read test
		denylist, err := clientService(badProfileId).List(ctx)
		require.Error(t, err)
		require.Empty(t, denylist)

		// Create
		denylistitem, err := clientService(badProfileId).Create(ctx, models.Denylist{})
		require.Error(t, err)
		require.Nil(t, denylistitem)

		for _, badDenyId := range []string{"", "not-a-denylist-id"} {
			// Bind
			boundDenyItem := clientService(badProfileId).Bind(badDenyId)

			// List-get test
			denylistitem, err = boundDenyItem.Get(ctx)
			require.Error(t, err)
			require.Nil(t, denylistitem)

			// update test
			err = boundDenyItem.Update(ctx, nil)
			require.Error(t, err)
			require.Nil(t, denylistitem)

			// Delete
			err = boundDenyItem.Delete(ctx)
			require.Error(t, err)

		}
	}
}

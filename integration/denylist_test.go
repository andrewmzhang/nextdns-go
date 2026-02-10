//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/andrewmzhang/nextdns/models"
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
	// fmt.Println(denyItem.Update(ctx, models.Denylist{
	// 	Active: true,
	// }))
	// denylist, err = client.Denylist(create.ID).List(ctx)
	// require.NoError(t, err)
	// require.True(t, denylist[0].Active)
}

//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/andrewmzhang/nextdns/models"
	"github.com/stretchr/testify/require"
)

func TestAllowlistLifecycle(t *testing.T) {
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

	allowlist, err := client.Allowlists(create.ID).List(ctx)
	fmt.Println(allowlist)
	require.NoError(t, err)
	require.NotNil(t, allowlist)
	require.Equal(t, len(allowlist), 0)

	allowcreate, err := client.Allowlists(create.ID).Create(ctx, models.Allowlist{
		ID:     "example3.com",
		Active: true,
	})
	fmt.Println(allowcreate)
	require.NoError(t, err)
	// Get will not work with allowlist
	allowItem, err := client.Allowlists(create.ID).Bind("example3.com").Get(ctx)
	require.NoError(t, err)
	require.True(t, allowItem.Active)
	// fmt.Println(allowItem.Update(ctx, models.Allowlist{
	// 	Active: true,
	// }))
	// allowlist, err = client.Allowlist(create.ID).List(ctx)
	// require.NoError(t, err)
	// require.True(t, allowlist[0].Active)
}

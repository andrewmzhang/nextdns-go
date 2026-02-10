package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/andrewmzhang/nextdns/models"
	"github.com/andrewmzhang/nextdns/services"
	"github.com/stretchr/testify/require"
)

func TestAnalyticStatus(t *testing.T) {
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

	config := services.AnalyticsDevicesQueryConfig{}
	config.From = "-3M"
	analyticStatuses, err := client.AnalyticDevices("8d2963").ListAll(ctx, config, false)
	require.NoError(t, err)
	fmt.Println(analyticStatuses)

}

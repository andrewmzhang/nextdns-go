//go:build integration

package integration

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/andrewmzhang/nextdns/services"
	"github.com/joho/godotenv"
)

var (
	nextdnsApiToken string
	client          *services.NextDNSClient
)

// TestMain runs once for the entire package
func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	nextdnsApiToken = os.Getenv("NEXTDNS_API_TOKEN")
	if nextdnsApiToken == "" {
		fmt.Println("Skipping integration tests: API_BASE_URL or API_AUTH_TOKEN not set")
		os.Exit(0)
	}

	client, err = services.NewClient(services.WithAPIKey(nextdnsApiToken))
	if err != nil {
		panic(err)
	}
	// Check that there are no lingering integration test profiles
	// profiles, err := client.Profiles().List(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// for _, profileSummary := range profiles {
	// 	if isCreatedByIntegrationTest(profileSummary) {
	// 		panic("Not expecting profile " + profileSummary.Name)
	// 	}
	// }
	os.Exit(m.Run())
}

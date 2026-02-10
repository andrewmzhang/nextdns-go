//go:build integration

package integration

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/andrewmzhang/nextdns-go/services"
	"github.com/joho/godotenv"
)

var (
	nextdnsApiToken string
	client          *services.NextDNSClient
)

// TestMain runs once for the entire package
func TestMain(m *testing.M) {
	// Load .env if it exists; ignore error if file is missing
	if err := godotenv.Load("../.env"); err != nil {
		if !os.IsNotExist(err) {
			// Only log errors other than file-not-found
			log.Println("Warning: error loading .env:", err)
		}
	}
	nextdnsApiToken = os.Getenv("NEXTDNS_API_TOKEN")
	if nextdnsApiToken == "" {
		fmt.Println("Failing integration tests: NEXTDNS_API_TOKEN not set")
		os.Exit(1)
	}
	var err error
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

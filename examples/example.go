package main

import (
	"context"
	"fmt"
	"github.com/amalucelli/nextdns-go/nextdns"
	"os"
)

func main() {
	// get the api key from the environment
	key := os.Getenv("NEXTDNS_API_KEY")
	fmt.Println("API key is", key)

	// client with a custom API key
	ctx := context.Background()
	client, err := nextdns.New(
		nextdns.WithAPIKey(key),
	)
	if err != nil {
		panic(err)
	}

	// set a few settings like the name and some other attributes
	create := &nextdns.CreateProfileRequest{
		Name: "nextdns-go",
		Denylist: []*nextdns.Denylist{
			{
				ID:     "google.com",
				Active: true,
			},
			{
				ID:     "bing.com",
				Active: true,
			},
		},
		Allowlist: []*nextdns.Allowlist{
			{
				ID:     "duckduckgo.com",
				Active: true,
			},
			{
				ID:     "search.brave.com",
				Active: false,
			},
		},
		ParentalControl: &nextdns.ParentalControl{
			Categories: []*nextdns.ParentalControlCategories{
				{
					ID:     "gambling",
					Active: true,
				},
			},
		},
		Security: &nextdns.Security{
			AiThreatDetection: true,
		},
		Settings: &nextdns.Settings{
			Logs: &nextdns.SettingsLogs{
				Enabled: true,
			},
			Web3: true,
		},
	}

	// create a new profile
	id, _ := client.Profiles.Create(ctx, create)

	// set a few settings like the name and some other attributes
	update := &nextdns.UpdateProfileRequest{
		ProfileID: id,
		Profile: &nextdns.Profile{
			Name: "nextdns-go-updated",
			Settings: &nextdns.Settings{
				Logs: &nextdns.SettingsLogs{
					Enabled: false,
				},
			},
		},
	}

	// update the profile
	_ = client.Profiles.Update(ctx, update)

	// get the profile details to check the settings
	profile, _ := client.Profiles.Get(ctx, &nextdns.GetProfileRequest{
		ProfileID: id,
	})
	fmt.Printf("%q profile name: %s\n", id, profile.Name)
	fmt.Printf("%q logs status: %t\n", id, profile.Settings.Logs.Enabled)

	// list all the profiles
	profiles, _ := client.Profiles.List(ctx, &nextdns.ListProfileRequest{})
	fmt.Printf("Found %d profiles\n", len(profiles))
	for _, p := range profiles {
		fmt.Printf("ID: %q\n", p.ID)
		fmt.Printf("Name: %q\n", p.Name)
	}

	// delete the profile
	_ = client.Profiles.Delete(ctx, &nextdns.DeleteProfileRequest{
		ProfileID: id,
	})
}

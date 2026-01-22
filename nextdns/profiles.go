package nextdns

import (
	"context"
	"fmt"
	"net/http"
)

// profilesService is the HTTP path for the profiles API.
const profilesAPIPath = "profiles"

// CreateProfileRequest encapsulates the request for creating a new profile.
type CreateProfileRequest struct {
	Name            string           `json:"name,omitempty"`
	Security        *Security        `json:"security,omitempty"`
	Privacy         *Privacy         `json:"privacy,omitempty"`
	ParentalControl *ParentalControl `json:"parentalControl,omitempty"`
	Denylist        []*Denylist      `json:"denylist,omitempty"`
	Allowlist       []*Allowlist     `json:"allowlist,omitempty"`
	Settings        *Settings        `json:"settings,omitempty"`
	Rewrites        []*Rewrite       `json:"rewrites,omitempty"`
}

// UpdateProfileRequest encapsulates the request for setting custom profile settings.
type UpdateProfileRequest struct {
	ProfileID string
	Profile   *Profile
}

// GetProfileRequest encapsulates the request for getting a profile.
type GetProfileRequest struct {
	ProfileID string
}

// ListProfileRequest encapsulates the request for listing all the profiles.
type ListProfileRequest struct{}

// DeleteProfileRequest encapsulates the request for deleting a profile.
type DeleteProfileRequest struct {
	ProfileID string
}

// ProfilesService is an interface for communicating with the NextDNS API.
type ProfilesService interface {
	Create(context.Context, *CreateProfileRequest) (string, error)
	Get(context.Context, *GetProfileRequest) (*Profile, error)
	Update(context.Context, *UpdateProfileRequest) error
	List(context.Context, *ListProfileRequest) ([]*ProfileSummary, error)
	Delete(context.Context, *DeleteProfileRequest) error
}

// Profile represents a NextDNS profile. Object structure closely resembles how nextdns.io page is organized.
type Profile struct {
	// ID of the NextDNS Profile. This can be found in the profile's URL string and on the "Setup" tab -> "Endpoints"
	// -> "ID"
	ID string
	// Fingerprint of the NextDNS Profile TODO: Make this more clear
	Fingerprint string

	// Name of the NextDNS profile. Found on the top of the `nextdns.io` page. Value is unique on a per-account basis
	Name string `json:"name,omitempty"`
	// Setup contains information found in the "Setup" tab on `nextdns.io`
	Setup *Setup `json:"setup,omitempty"`
	// Security contains information found in the "Security" tab on `nextdns.io`
	Security *Security `json:"security,omitempty"`
	// Privacy contains information found in the "Privacy" tab on `nextdns.io`
	Privacy *Privacy `json:"privacy,omitempty"`
	// ParentalControl contains information found in the "Parental Control" tab on `nextdns.io`
	ParentalControl *ParentalControl `json:"parentalControl,omitempty"`
	// Denylist contains a list of denied domains found on the "Denylist" tab on `nextdns.io`
	Denylist []*Denylist `json:"denylist,omitempty"`
	// Allowlist contains a list of allowed domains found on the "Allowlist" tab on `nextdns.io`
	Allowlist []*Allowlist `json:"allowlist,omitempty"`

	// TODO implement Analytics and Logs

	// Settings contains information found in the "Settings" tab on `nextdns.io`, excluding the "Rewrites" section
	Settings *Settings `json:"settings,omitempty"`
	// TODO: Consider folding Rewrites under settings
	// Rewrites contains a list of dns domain rewrites found in the "Rewrites" section located near bottom of the
	// "Settings" tab on `nextdns.io`
	Rewrites []*Rewrite `json:"rewrites,omitempty"`
}

// newProfileRequest represents the response from a new profile request.
type newProfileResponse struct {
	Profile struct {
		ID string `json:"id"`
	} `json:"data"`
}

// ProfileSummary represents basic information for a NextDNS Profile.
type ProfileSummary struct {
	ID          string `json:"id"`
	Fingerprint string `json:"fingerprint"`
	Name        string `json:"name"`
}

// profileResponse represents the response for getting a profile using the NextDNS API.
type profileResponse struct {
	Profile *Profile `json:"data"`
}

// profilesResponse represents the response for listing all profiles using the NextDNS API.
type profilesResponse struct {
	ProfileSummaries []*ProfileSummary `json:"data"`
	Metadata         struct {
		Pagination struct {
			Cursor string `json:"cursor"`
		} `json:"pagination"`
	} `json:"meta,omitempty"`
	Errors ErrorResponse `json:"errors,omitempty"`
}

// profilesService is a concrete implementation of NextDNS profiles service interface.
type profilesService struct {
	client *Client
}

var _ ProfilesService = &profilesService{}

// NewProfilesService returns a new NextDNS profiles service.
// nolint: revive
func NewProfilesService(client *Client) ProfilesService {
	return &profilesService{
		client: client,
	}
}

// List returns a list of all profiles under a NextDNS account. Note that the 2nd argument is ignored and only exists
// for function signature consistency.
func (s *profilesService) List(ctx context.Context, _ *ListProfileRequest) ([]*ProfileSummary, error) {
	req, err := s.client.newRequest(http.MethodGet, profilesAPIPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to list the profiles: %w", err)
	}

	response := profilesResponse{}
	err = s.client.do(ctx, req, &response)
	if err != nil {
		return nil, fmt.Errorf("error making a request to list the profiles: %w", err)
	}

	return response.ProfileSummaries, nil
}

// Create creates a profile on NextDNS and returns the newly-created profile's ID.
func (s *profilesService) Create(ctx context.Context, request *CreateProfileRequest) (string, error) {
	req, err := s.client.newRequest(http.MethodPost, profilesAPIPath, request)
	if err != nil {
		return "", fmt.Errorf("error creating request to create a profile: %w", err)
	}

	response := &newProfileResponse{}
	err = s.client.do(ctx, req, &response)
	if err != nil {
		return "", fmt.Errorf("error making a request to create a profile: %w", err)
	}

	return response.Profile.ID, nil
}

// Update updates the settings of an existing profile on NextDNS.
func (s *profilesService) Update(ctx context.Context, request *UpdateProfileRequest) error {
	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
	req, err := s.client.newRequest(http.MethodPatch, path, request.Profile)
	if err != nil {
		return fmt.Errorf("error creating request to update the profile: %w", err)
	}

	response := profileResponse{}
	err = s.client.do(ctx, req, &response)
	if err != nil {
		return fmt.Errorf("error making a request to update the profile: %w", err)
	}

	return nil
}

// Get returns a profile on NextDNS.
func (s *profilesService) Get(ctx context.Context, request *GetProfileRequest) (*Profile, error) {
	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to get the profile: %w", err)
	}

	response := profileResponse{}
	err = s.client.do(ctx, req, &response)
	if err != nil {
		return nil, fmt.Errorf("error making a request to get the profile: %w", err)
	}

	return response.Profile, nil
}

// Delete deletes a profile from NextDNS.
func (s *profilesService) Delete(ctx context.Context, request *DeleteProfileRequest) error {
	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("error creating request to delete the profile: %w", err)
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error making a request to delete the profile: %w", err)
	}

	return err
}

// profileAPIPath returns the profile API path.
func profileAPIPath(profile string) string {
	return fmt.Sprintf("%s/%s", profilesAPIPath, profile)
}

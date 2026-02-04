package nextdns

import (
	"github.com/andrewmzhang/nextdns/models"
)

// profilesService is the HTTP path for the profiles API.
const profilesAPIPath = "profiles"

// CreateProfileRequest encapsulates the request for creating a new profile.
type CreateProfileRequest struct {
	Name            string                  `json:"name,omitempty"`
	Security        *models.Security        `json:"security,omitempty"`
	Privacy         *models.Privacy         `json:"privacy,omitempty"`
	ParentalControl *models.ParentalControl `json:"parentalControl,omitempty"`
	Denylist        []*models.Denylist      `json:"denylist,omitempty"`
	Allowlist       []*models.Allowlist     `json:"allowlist,omitempty"`
	Settings        *models.Settings        `json:"settings,omitempty"`
	Rewrites        []*models.Rewrite       `json:"rewrites,omitempty"`
}

// DeleteProfileRequest encapsulates the request for deleting a profile.
type DeleteProfileRequest struct {
	ProfileID string
}

// UpdateProfileRequest encapsulates the request for setting custom profile settings.
type UpdateProfileRequest struct {
	ProfileID string
	Profile   *models.Profile
}

// GetProfileRequest encapsulates the request for getting a profile.
type GetProfileRequest struct {
	ProfileID string
}

// ListProfileRequest encapsulates the request for listing all the profiles.
type ListProfileRequest struct{}

// // ProfilesService is an interface for communicating with the NextDNS API.
// type ProfilesService interface {
// 	List(context.Context, *ListProfileRequest) ([]models.Profile, error)
// 	Create(context.Context, *CreateProfileRequest) (string, error)
// 	Get(context.Context, *GetProfileRequest) (*models.Profile, error)
// 	Update(context.Context, *UpdateProfileRequest) (*models.Profile, error)
// 	Delete(context.Context, *DeleteProfileRequest) error
// }
//
// // profilesService is a concrete implementation of NextDNS profiles service interface.
// type profilesService struct {
// 	nextDNSClient *NextDNSClient
// }
//
// // NewProfilesService returns a new NextDNS profiles service.
// // nolint: revive
// func NewProfilesService(client *NextDNSClient) ProfilesService {
// 	return &profilesService{
// 		nextDNSClient: client,
// 	}
// }
//
// // List returns a list of all profiles under a NextDNS account. Note that the 2nd argument is ignored and only exists
// // for function signature consistency.
// func (s *profilesService) List(ctx context.Context, _ *ListProfileRequest) ([]models.Profile, error) {
// 	listProfilesResponse := struct {
// 		Data []models.Profile `json:"data"`
// 	}{}
// 	resp, err := s.nextDNSClient.client.R().SetResult(&listProfilesResponse).Get(profilesAPIPath)
// 	return handleResponse(listProfilesResponse.Data, resp, err)
// }
//
// // Create creates a profile on NextDNS and returns the newly-created profile's ID.
// func (s *profilesService) Create(ctx context.Context, request *CreateProfileRequest) (string, error) {
// 	// newProfileRequest represents the response from a new profile request.
// 	newProfileResponse := struct {
// 		Profile struct {
// 			ID string `json:"id"`
// 		} `json:"data"`
// 	}{}
// 	path := fmt.Sprintf("%s/", profilesAPIPath)
// 	resp, err := s.nextDNSClient.client.R().SetBody(request).SetResult(&newProfileResponse).Post(path)
// 	return handleResponse(newProfileResponse.Profile.ID, resp, err)
// }
//
// // Update updates the settings of an existing profile on NextDNS.
// func (s *profilesService) Update(ctx context.Context, request *UpdateProfileRequest) (*models.Profile, error) {
// 	response := struct {
// 		Profile *models.Profile `json:"data"`
// 	}{}
// 	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
// 	resp, err := s.nextDNSClient.client.R().SetBody(request).SetResult(&response).Patch(path)
// 	return handleResponse(response.Profile, resp, err)
// }
//
// // Get returns a profile on NextDNS.
// func (s *profilesService) Get(ctx context.Context, request *GetProfileRequest) (*models.Profile, error) {
// 	response := struct {
// 		Profile *models.Profile `json:"data"`
// 	}{}
// 	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
// 	resp, err := s.nextDNSClient.client.R().SetResult(&response).Get(path)
// 	return handleResponse(response.Profile, resp, err)
// }
//
// // Delete deletes a profile from NextDNS.
// func (s *profilesService) Delete(ctx context.Context, request *DeleteProfileRequest) error {
// 	path := fmt.Sprintf("%s/%s", profilesAPIPath, request.ProfileID)
// 	resp, err := s.nextDNSClient.client.R().Delete(path)
// 	_, respErr := handleResponse(any(nil), resp, err)
// 	return respErr
// }
//
// var _ ProfilesService = &profilesService{}
//
// func profileAPIPath(profileID string) string {
// 	return fmt.Sprintf("%s/%s", profilesAPIPath, profileID)
// }

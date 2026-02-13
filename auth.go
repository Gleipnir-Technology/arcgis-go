package arcgis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Authenticator interface {
	init(ctx context.Context, client http.Client) error
	//addAuthentication(req *http.Request) (*http.Request, error)
	addAuthHeaders(context.Context, map[string]string) map[string]string
}

type AuthenticatorUsernamePassword struct {
	AccessToken        string
	AccessTokenExpires time.Time
	Password           string
	Username           string
}

// Authentication by adding in a 'token=?' query argument to requests
type AuthenticatorToken struct {
	Token string
}

type AuthenticatorOAuth struct {
	AccessToken         string
	AccessTokenExpires  time.Time
	RefreshToken        string
	RefreshTokenExpires time.Time
}

func (a *AuthenticatorUsernamePassword) addAuthentication(req *http.Request) (*http.Request, error) {
	req.Header.Add("X-ESRI-Authorization", "Bearer "+a.AccessToken)
	return req, nil
}
func (a *AuthenticatorUsernamePassword) addAuthHeaders(ctx context.Context, h map[string]string) map[string]string {
	//logger := zerolog.Ctx(ctx)
	val := "Bearer " + a.AccessToken
	h["X-ESRI-Authorization"] = val
	//logger.Debug().Str("val", val).Msg("Added auth header")
	return h
}
func (a *AuthenticatorUsernamePassword) init(ctx context.Context, client http.Client) error {
	logger := zerolog.Ctx(ctx)
	// GET https://www.arcgis.com/sharing/rest/oauth2/authorize
	host := "https://www.arcgis.com"
	path := "/sharing/rest/oauth2/authorize"
	state, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("make uuid: %w", err)
	}
	queryargs := map[string]string{
		"response_type": "token",
		// Using 'pythonapi' is stupid, but required since it only accepts recognized clients
		"client_id":          "pythonapi",
		"redirect_uri":       "https://www.arcgis.com",
		"state":              state.String(),
		"expiration":         "600",
		"allow_verification": "false",
		"style":              "dark",
		"locale":             "en-US",
	}
	headers := map[string]string{
		"User-Agent": "Geosaurus/2.4.2",
		// Handled automatically by the transport layer
		// "Accept-Encoding": "gzip, deflate, br",
		"Accept": "*/*",
	}
	content, err := doGetParamsHeaders(ctx, client, host, path, queryargs, headers)
	if err != nil {
		return fmt.Errorf("get oauth state: %w", err)
	}

	// Extract the oauth info
	info, err := extractOAuthInfoFromHTML(logger, string(content))
	if err != nil {
		return fmt.Errorf("extract oauth: %w", err)
	}

	// POST https://www.arcgis.com/sharing/oauth2/signin
	path = "/sharing/oauth2/signin"
	queryargs = map[string]string{
		"expiration":  "600",
		"oauth_state": info.OAuthState,
		"username":    a.Username,
		"password":    a.Password,
	}
	// prevent default redirect behavior or we'll lose the location header we need
	prev_redirect := client.CheckRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	resp_headers, _, err := doPostFormParams(ctx, client, host, path, queryargs)
	if err != nil {
		return fmt.Errorf("do auth post: %w", err)
	}
	location := resp_headers.Get("location")
	if location == "" {
		return fmt.Errorf("no location header")
	}
	oauth_resp, err := parseLocationHeader(logger, location)
	if err != nil {
		return fmt.Errorf("parse location: %w", err)
	}
	returned_state, err := uuid.Parse(oauth_resp.State)
	if err != nil {
		return fmt.Errorf("parse state: %w", err)
	}
	if state != returned_state {
		return fmt.Errorf("state mismatch %s != %s", state.String(), oauth_resp.State)
	}
	// Parse out location headers
	a.AccessToken = oauth_resp.AccessToken
	a.AccessTokenExpires = oauth_resp.ExpiresAt
	// reset the redirect behavior
	client.CheckRedirect = prev_redirect

	logger.Debug().Str("access_token", oauth_resp.AccessToken).Msg("switched access token")

	return nil
}
func (a AuthenticatorToken) addAuthentication(req *http.Request) (*http.Request, error) {
	panic("Not implemented")
	//return nil, nil
}
func (a AuthenticatorToken) addAuthHeaders(ctx context.Context, h map[string]string) map[string]string {
	h["X-ESRI-Authorization"] = "Bearer " + a.Token
	return h
}
func (a AuthenticatorToken) init(ctx context.Context, client http.Client) error {
	return nil
}
func (a AuthenticatorOAuth) addAuthentication(req *http.Request) (*http.Request, error) {
	req.Header.Add("X-ESRI-Authorization", "Bearer "+a.AccessToken)
	return req, nil
}
func (a AuthenticatorOAuth) addAuthHeaders(ctx context.Context, h map[string]string) map[string]string {
	h["X-ESRI-Authorization"] = "Bearer " + a.AccessToken
	return h
}
func (a AuthenticatorOAuth) init(ctx context.Context, client http.Client) error {
	return nil
}

// OAuthInfo represents the structure of the oAuthInfo object
type oAuthInfo struct {
	OAuthState                string     `json:"oauth_state"`
	ClientID                  string     `json:"client_id"`
	AppTitle                  string     `json:"appTitle"`
	Locale                    string     `json:"locale"`
	PersistOption             bool       `json:"persistOption"`
	ShowSocialLogins          bool       `json:"showSocialLogins"`
	ContextPath               string     `json:"contextPath"`
	AppOrgInfo                appOrgInfo `json:"appOrgInfo"`
	OrgUrlBase                string     `json:"orgUrlBase"`
	HelpBase                  string     `json:"helpBase"`
	Style                     string     `json:"style"`
	SocialProviders           []string   `json:"socialProviders"`
	OriginSignin              bool       `json:"originSignin"`
	IsMfaRecoveryCodesEnabled bool       `json:"isMfaRecoveryCodesEnabled"`
	CustomOrgLoginEnabled     bool       `json:"customOrgLoginEnabled"`
}

// AppOrgInfo represents the organization information
type appOrgInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

// OAuthResponse represents the parsed OAuth response parameters
type oAuthResponse struct {
	AccessToken string
	ExpiresAt   time.Time
	Username    string
	SSL         string
	State       string
	// Add other fields as needed
}

// ExtractOAuthInfoFromHTML extracts and parses the oAuthInfo object from the HTML response
func extractOAuthInfoFromHTML(logger *zerolog.Logger, html string) (*oAuthInfo, error) {
	logger.Debug().Str("html", html).Msg("searching html")
	// Regular expression to find the oAuthInfo assignment
	//re := regexp.MustCompile(`var oAuthInfo = (.*?);`)
	re := regexp.MustCompile(`var oAuthInfo = (.*)\n`)
	matches := re.FindStringSubmatch(html)
	if matches == nil {
		return nil, fmt.Errorf("oAuthInfo not found in the HTML")
	}

	// Extract the JSON string
	json_str := matches[1]
	logger.Debug().RawJSON("json", []byte(json_str)).Msg("parsing json")

	// Parse the JSON into our struct
	var oauthInfo oAuthInfo
	if err := json.Unmarshal([]byte(json_str), &oauthInfo); err != nil {
		return nil, fmt.Errorf("failed to parse oAuthInfo: %v", err)
	}

	return &oauthInfo, nil
}

// ParseLocationHeader parses the OAuth response parameters from a location header URL
func parseLocationHeader(logger *zerolog.Logger, location string) (*oAuthResponse, error) {
	// Find the fragment (part after #)
	logger.Debug().Str("location", location).Msg("Searching for location")
	fragmentIndex := strings.IndexByte(location, '#')
	if fragmentIndex == -1 {
		return nil, fmt.Errorf("no fragment found in location header")
	}

	// Extract the fragment portion (excluding the # character)
	fragment := location[fragmentIndex+1:]

	// Split the fragment into key-value pairs
	params := make(map[string]string)
	pairs := strings.Split(fragment, "&")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			params[key] = value
		}
	}

	// Create response struct
	response := &oAuthResponse{
		AccessToken: params["access_token"],
		Username:    params["username"],
		SSL:         params["ssl"],
		State:       params["state"],
	}

	// Convert expires_in to a proper expiration time
	if expiresInStr, ok := params["expires_in"]; ok {
		expiresInSeconds, err := strconv.ParseInt(expiresInStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse expires_in value: %v", err)
		}
		// Calculate expiration time by adding seconds to current time
		response.ExpiresAt = time.Now().Add(time.Duration(expiresInSeconds) * time.Second)
	}

	return response, nil
}

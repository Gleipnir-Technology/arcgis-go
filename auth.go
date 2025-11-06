package arcgis

import (
	"net/http"
	"time"
)

type Authenticator interface {
	addAuthentication(req *http.Request) (*http.Request, error)
}

// Authentication by adding in a 'token=?' query argument to requests
type AuthenticatorToken struct {
	Token string
}

type AuthenticatorOAuth struct {
	AccessToken string
	Expires time.Time
	RefreshToken string
}

func (a AuthenticatorToken) addAuthentication(req *http.Request) (*http.Request, error) {
	return nil, nil
}
func (a AuthenticatorOAuth) addAuthentication(req *http.Request) (*http.Request, error) {
 	req.Header.Add("X-ESRI-Authorization", "Bearer "+a.AccessToken)
	return req, nil
}

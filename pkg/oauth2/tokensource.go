package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	user "github.com/kraney/stackpath/pkg/accounts_and_users"
	"golang.org/x/oauth2"
)

// TokenSource is an oauth2 token source for stackpath
type TokenSource struct {
	req    user.IdentityGetAccessTokenRequest
	Client *http.Client
}

// NewTokenSource returns a new token source for stackpath
func NewTokenSource(clientid string, clientsecret string) *TokenSource {
	creds := "client_credentials"
	return &TokenSource{
		req: user.IdentityGetAccessTokenRequest{
			ClientId:     &clientid,
			ClientSecret: &clientsecret,
			GrantType:    &creds,
		},
	}
}

// Token returns a token
func (sps *TokenSource) Token() (*oauth2.Token, error) {
	usercfg := user.NewConfiguration()
	usercfg.HTTPClient = sps.Client
	userclient := user.NewAPIClient(usercfg)
	idresp, _, err := userclient.AuthenticationApi.GetAccessToken(context.Background()).IdentityGetAccessTokenRequest(sps.req).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error obtaining oauth2 token: %v", err)
	}
	return &oauth2.Token{
		AccessToken: *idresp.AccessToken,
		TokenType:   *idresp.TokenType,
		Expiry:      time.Now().Add(time.Second * time.Duration(*idresp.ExpiresIn-30)),
	}, nil
}

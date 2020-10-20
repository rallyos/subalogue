package auth

import (
	"context"
	"log"

	"golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://subalogue.eu.auth0.com/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "4uNQLwCkcWPYKoTmSLvot6L2u39VOhWi",
		ClientSecret: "LwjwZEUxKpgWnvDi4may7-wyNBKEWdDOiaX4vE7FjhX6nwipWyNNRM12dRCC_oS7",
		RedirectURL:  "http://localhost:8000/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}

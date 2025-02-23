package auth

import "errors"

var (
	ErrCredentialNotFound = errors.New("credential not found")
)

type AuthenticatedUser struct {
	AuthenticatedByRealm string     `json:"-"` // Name of realm that authenticated this user
	Credential           Credential `json:"credential"`
	Role                 Role       `json:"role"`
}

type Credential struct {
	Type     CredentialType `json:"type"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	APIKey   string         `json:"api_key"`
}

func (c *Credential) String() string {
	return c.Username
}

type CredentialType string

const (
	CredentialTypeBasic  = CredentialType("BASIC")
	CredentialTypeAPIKey = CredentialType("BEARER")
)

func (c CredentialType) String() string {
	return string(c)
}

package coc

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client     *resty.Client
	ipAddress  string
	accounts   []APIAccount
	index      Index
	StatusCode int
	sync.Mutex
	ready bool
}

// Keep track of what key we're at and what credential we're at.
type Index struct {
	KeyAccountIndex int // Index of the credential we're at.
	KeyIndex        int // Index of the key we're at.
}

// Each API account used to log in, with a unique credential.
type APIAccount struct {
	Credential Credential    // The credential used to log in.
	Response   LoginResponse // The response from the login.
	Keys       Keys          // The current account's keys.
}

type Credential struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Auth                    Auth      `json:"auth"`
	Developer               Developer `json:"developer"`
	TemporaryAPIToken       string    `json:"temporaryAPIToken"`
	SwaggerURL              string    `json:"swaggerUrl"`
	Status                  Status    `json:"status"`
	SessionExpiresInSeconds int       `json:"sessionExpiresInSeconds"`
}

type Auth struct {
	Ua    any    `json:"ua"`
	IP    any    `json:"ip"`
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

type Developer struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Game          string `json:"game"`
	Email         string `json:"email"`
	Tier          string `json:"tier"`
	AllowedScopes any    `json:"allowedScopes"`
	MaxCidrs      any    `json:"maxCidrs"`
	PrevLoginTs   string `json:"prevLoginTs"`
	PrevLoginIP   string `json:"prevLoginIp"`
	PrevLoginUa   string `json:"prevLoginUa"`
}

type KeyResponse struct {
	Key                     Key    `json:"key,omitempty"`
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
}

type Key struct {
	Origins     any      `json:"origins"`
	ValidUntil  any      `json:"validUntil"`
	ID          string   `json:"id"`
	Developerid string   `json:"developerId"`
	Tier        string   `json:"tier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Key         string   `json:"key"`
	Scopes      []string `json:"scopes"`
	Cidrranges  []string `json:"cidrRanges"`
}

type Keys struct {
	Status        Status `json:"status"`
	Keys          []Key  `json:"keys"`
	SessionExpire int    `json:"sessionExpiresInSeconds"`
}

type Status struct {
	Detail  any    `json:"detail"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

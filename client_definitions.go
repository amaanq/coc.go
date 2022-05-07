package coc

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	sync.Mutex
	waitGroup sync.WaitGroup

	client *resty.Client
	ready  bool

	accounts []APIAccount

	index      Index
	StatusCode int
	ipAddress  string
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
	Status                  Status    `json:"status"`
	SessionExpiresInSeconds int       `json:"sessionExpiresInSeconds"`
	Auth                    Auth      `json:"auth"`
	Developer               Developer `json:"developer"`
	TemporaryAPIToken       string    `json:"temporaryAPIToken"`
	SwaggerURL              string    `json:"swaggerUrl"`
}

type Auth struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Ua    any    `json:"ua"`
	IP    any    `json:"ip"`
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
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	Key                     Key    `json:"key,omitempty"`
}

type Key struct {
	ID          string   `json:"id"`
	Developerid string   `json:"developerId"`
	Tier        string   `json:"tier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Origins     any      `json:"origins"`
	Scopes      []string `json:"scopes"`
	Cidrranges  []string `json:"cidrRanges"`
	ValidUntil  any      `json:"validUntil"`
	Key         string   `json:"key"`
}

type Keys struct {
	Keys          []Key  `json:"keys"`
	Status        Status `json:"status"`
	SessionExpire int    `json:"sessionExpiresInSeconds"`
}

type Status struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  any    `json:"detail"`
}

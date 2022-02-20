package client

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

type HTTPSessionManager struct {
	sync.RWMutex
	*resty.Client
	sync.WaitGroup

	cache *cache.Cache
	ready bool

	Credentials   []LoginCredential
	KeyNames      string
	KeyCount      int
	CacheMaxSize  int
	LoginResponse LoginResponse
	KeysList      KeysList
	RawKeysList   []Key
	KeyIndex      int
	IP            string
}

type LoginCredential struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Status                  Status    `json:"status"`
	SessionExpiresInSeconds int64     `json:"sessionExpiresInSeconds"`
	Auth                    Auth      `json:"auth"`
	Developer               Developer `json:"developer"`
	TemporaryAPIToken       string    `json:"temporaryAPIToken"`
	SwaggerURL              string    `json:"swaggerUrl"`
}

type Auth struct {
	Uid   string      `json:"uid"`
	Token string      `json:"token"`
	Ua    interface{} `json:"ua"`
	IP    interface{} `json:"ip"`
}

type Developer struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Game          string      `json:"game"`
	Email         string      `json:"email"`
	Tier          string      `json:"tier"`
	AllowedScopes interface{} `json:"allowedScopes"`
	MaxCidrs      interface{} `json:"maxCidrs"`
	PrevLoginTs   string      `json:"prevLoginTs"`
	PrevLoginIP   string      `json:"prevLoginIp"`
	PrevLoginUa   string      `json:"prevLoginUa"`
}

type KeyCreationResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int64  `json:"sessionExpiresInSeconds"`
	Key                     Key    `json:"key"`
}

type KeyDeletionResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int64  `json:"sessionExpiresInSeconds"`
}

type Key struct {
	ID          string      `json:"id"`
	Developerid string      `json:"developerId"`
	Tier        string      `json:"tier"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Origins     interface{} `json:"origins"`
	Scopes      []string    `json:"scopes"`
	Cidrranges  []string    `json:"cidrRanges"`
	ValidUntil  interface{} `json:"validUntil"`
	Key         string      `json:"key"`
}

type KeysList struct {
	Status        Status `json:"status"`
	SessionExpire int    `json:"sessionExpiresInSeconds"`
	Keys          []Key  `json:"keys"`
}

type Status struct {
	Code    int64       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  interface{} `json:"detail"`
}

type ClientError struct {
	error
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Detail  interface{} `json:"detail"`
}

func (c *ClientError) SetErr(err error) {
	c.error = err
}

func (c *ClientError) Err() error {
	return c.error
}

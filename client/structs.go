package client

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

type HTTPSessionManager struct {
	mutex     sync.Mutex
	waitGroup sync.WaitGroup

	client *resty.Client
	cache  *cache.Cache
	ready  bool

	credentials      []LoginCredential
	logins           []LoginResponse
	currentLoginKeys Keys
	allKeys          Keys

	keyIndex   int
	StatusCode int
	iP         string
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

type KeyResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int64  `json:"sessionExpiresInSeconds"`
	Key                     Key    `json:"key,omitempty"`
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

type Keys struct {
	Keys          []Key  `json:"keys"`
	Status        Status `json:"status"`
	SessionExpire int    `json:"sessionExpiresInSeconds"`
}

type Status struct {
	Code    int64       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  interface{} `json:"detail"`
}

type ClientError struct {
	Error   error
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Detail  interface{} `json:"detail"`
}

func (c *ClientError) SetErr(err error) {
	c.Error = err
}

func (c *ClientError) Err() error {
	return c.Error
}

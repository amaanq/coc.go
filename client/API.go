package client

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

type HTTPSessionManager struct {
	Credentials   []LoginCredential
	Client        *resty.Client
	KeyNames      string
	KeyCount      int
	CacheMaxSize  int
	LoginResponse LoginResponse
	KeysList      KeysList
	RawKeysList   []Key
	KeyIndex      int
	IP            string
	wg            sync.WaitGroup
	mutex         sync.RWMutex
	cache         *cache.Cache
}

type LoginCredential struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Status                  Status `json:"status,omitempty"`
	SessionExpiresInSeconds *int64 `json:"sessionExpiresInSeconds,omitempty"`
	Auth                    struct {
		Uid   *string     `json:"uid,omitempty"`
		Token *string     `json:"token,omitempty"`
		Ua    interface{} `json:"ua"`
		IP    interface{} `json:"ip"`
	} `json:"auth,omitempty"`
	Developer struct {
		ID            *string     `json:"id,omitempty"`
		Name          *string     `json:"name,omitempty"`
		Game          *string     `json:"game,omitempty"`
		Email         *string     `json:"email,omitempty"`
		Tier          *string     `json:"tier,omitempty"`
		AllowedScopes interface{} `json:"allowedScopes"`
		MaxCidrs      interface{} `json:"maxCidrs"`
		PrevLoginTs   *string     `json:"prevLoginTs,omitempty"`
		PrevLoginIP   *string     `json:"prevLoginIp,omitempty"`
		PrevLoginUa   *string     `json:"prevLoginUa,omitempty"`
	} `json:"developer,omitempty"`
	TemporaryAPIToken string `json:"temporaryAPIToken,omitempty"`
	SwaggerURL        string `json:"swaggerUrl,omitempty"`
}

type KeyCreationResponse struct {
	Status                  Status `json:"status,omitempty"`
	SessionExpiresInSeconds int64  `json:"sessionExpiresInSeconds,omitempty"`
	Key                     Key    `json:"key,omitempty"`
}

type KeyDeletionResponse struct {
	Status                  Status `json:"status,omitempty"`
	SessionExpiresInSeconds int64  `json:"sessionExpiresInSeconds,omitempty"`
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

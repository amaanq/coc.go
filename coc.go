package coc

import (
	"github.com/amaanq/coc.go/client"
)

func New(credentials map[string]string) (*client.HTTPSessionManager, error) {
	return client.New(credentials)
}

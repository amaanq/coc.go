package coc

import (
	"github.com/amaanq/coc.go/client"
	"github.com/amaanq/coc.go/tag"
)

// Pass in a map which maps a username to a password
func New(credentials map[string]string) (*client.HTTPSessionManager, error) {
	return client.New(credentials)
}

// this function is inside /client/ but this makes it easier to use outside of the client package.
func CorrectTag(_tag string) string {
	return string(tag.ToPlayerTag(_tag))
}

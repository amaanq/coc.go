package coc

// Pass in a map which maps a username to a password
func New(credentials map[string]string) (*Client, error) {
	return newClient(credentials)
}

// this function is inside /client/ but this makes it easier to use outside of the client package.
func CorrectTag(_tag string) string {
	return string(_tag)
}

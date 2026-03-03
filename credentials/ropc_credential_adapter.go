package credentials

import "context"

// GetAuthentication implements the core.Credential interface.
func (c *ROPCCredential) GetAuthentication() (string, error) {
	token, err := c.GetAuthorizationToken(context.Background(), nil, nil)
	if err != nil {
		return "", err
	}
	return "Bearer " + token, nil
}

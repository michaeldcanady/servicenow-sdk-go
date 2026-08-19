package mockdata

var TokenResponse = `{
  "access_token": "mock_access_token_abc123",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "mock_refresh_token_xyz789",
  "scope": "user_role"
}`

var RefreshTokenResponse = `{
  "access_token": "mock_refreshed_access_token_def456",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "mock_refreshed_refresh_token_uvw321",
  "scope": "user_role"
}`

var TokenErrorResponse = `{
  "error": "invalid_grant",
  "error_description": "The provided credentials are invalid."
}`

var RevocationResponse = ``

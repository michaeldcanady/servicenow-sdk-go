package auth

type StaticAuthTypeProvider struct {
	authType string
}

func (s *StaticAuthTypeProvider) GetAuthType() string {
	return s.authType
}

func NewStaticAuthTypeProvider(authType string) *StaticAuthTypeProvider {
	return &StaticAuthTypeProvider{
		authType: authType,
	}
}

package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

type redirectServer struct {
	*oauth2.Server
}

func (s *redirectServer) Result(ctx context.Context) (string, string, error) {
	res := s.Server.Result(ctx)
	return res.Code, res.State, res.Err
}

func defaultServerFactory(state string, port int) (AuthorizationCodeServer, error) {
	s, err := oauth2.NewServer(state, port)
	if err != nil {
		return nil, err
	}
	return &redirectServer{s}, nil
}

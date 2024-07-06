package auth

import (
	"context"
	"github.com/ahtamoff/protos/gen/go/auth"
	"google.golang.org/grpc"
)

type serverApi struct {
	auth.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	auth.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	panic("implement me")
}

func (s *serverApi) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	panic("implement me")
}

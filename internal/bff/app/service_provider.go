package app

import (
	"log"

	"github.com/InNOcentos/Doggo-track/backend/internal/bff/config"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	"google.golang.org/grpc"
)

type serviceProvider struct {
	grpcConfig          config.GrpcConfig
	httpConfig          config.HttpConfig
	grpcClientConn      *grpc.ClientConn
	userPbServiceClient userPb.UserServiceClient
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcConfig() config.GrpcConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HttpConfig() config.HttpConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GrpcConnection() *grpc.ClientConn {
	if s.grpcClientConn == nil {
		connStr := s.GrpcConfig().UserClientAddress()
		conn, err := grpc.Dial(connStr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed to create grpc conn: %s", err.Error())
		}

		s.grpcClientConn = conn
	}

	return s.grpcClientConn
}

func (s *serviceProvider) UserPbServiceClient() userPb.UserServiceClient {
	if s.userPbServiceClient == nil {
		s.userPbServiceClient = userPb.NewUserServiceClient(s.GrpcConnection())
	}

	return s.userPbServiceClient
}

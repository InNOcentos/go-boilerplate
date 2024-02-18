package app

import (
	"context"
	"log"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/api/user"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/client/db"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/client/db/mongo"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/config"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/repository"
	userRepository "github.com/InNOcentos/Doggo-track/backend/internal/account/repository/user"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/service"
	userService "github.com/InNOcentos/Doggo-track/backend/internal/account/service/user"
)

type serviceProvider struct {
	grpcConfig     config.GrpcConfig
	userImpl       *user.Implementation
	dbClient       db.Client
	userService    service.UserService
	userRepository repository.UserRepository
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

func (s *serviceProvider) DBClient(ctx context.Context, connectionUrl string) db.Client {
	if s.dbClient == nil {
		cl, err := mongo.New(ctx, connectionUrl)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		cl.Ping(ctx)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		cfg, err := config.NewMongoConfig()
		if err != nil {
			log.Fatalf("failed to get mongo config: %s", err.Error())
		}

		db := s.DBClient(ctx, cfg.ConnectionUrl()).Database(cfg.DatabaseName())
		s.userRepository = userRepository.NewRepository(db)
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

package app

import (
	"context"
	"flag"
	"net"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/config"
	userPb "github.com/InNOcentos/Doggo-track/backend/pkg/contracts/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer      *grpc.Server
	serviceProvider *serviceProvider
}

type AppFlags struct {
	cfgPath string
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initDeps(ctx context.Context) error {
	appFlags := &AppFlags{}
	err := a.initConfig(appFlags)
	if err != nil {
		return err
	}
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	userPb.RegisterUserServiceServer(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) initConfig(flags *AppFlags) error {
	flag.StringVar(&flags.cfgPath, "cfgPath", "./internal/account/.env.local", "Env config path")
	flag.Parse()
	err := config.Load(flags.cfgPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Run() error {
	list, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().GrpcServerAddress())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	provider := newServiceProvider()
	a.serviceProvider = provider
	return nil
}

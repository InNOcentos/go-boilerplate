package app

import (
	"context"
	"flag"
	"net/http"
	"time"

	_ "github.com/InNOcentos/Doggo-track/backend/dev/swagger"
	"github.com/InNOcentos/Doggo-track/backend/internal/bff/api"
	"github.com/InNOcentos/Doggo-track/backend/internal/bff/config"
	middleware "github.com/InNOcentos/Doggo-track/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	httpServer      *http.Server
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
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	g := gin.New()
	a.initSwagger(g)
	a.initMiddlewares(g)
	a.initRoutes(g)

	a.httpServer = &http.Server{
		Addr:           a.serviceProvider.HttpConfig().ServerAddress(),
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (a *App) initConfig(flags *AppFlags) error {
	flag.StringVar(&flags.cfgPath, "cfgPath", "./internal/bff/.env.local", "Env config path")
	flag.Parse()
	err := config.Load(flags.cfgPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Run() error {
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	provider := newServiceProvider()
	a.serviceProvider = provider
	return nil
}

func (a *App) initMiddlewares(g *gin.Engine) {
	g.Use(middleware.HttpRequestLogger())
	g.Use(gin.Logger())
}

func (a *App) initRoutes(g *gin.Engine) {
	r := g.Group("/api")
	api.RegisterUserHttpHandlers(r, a.serviceProvider.UserPbServiceClient())
}

func (a *App) initSwagger(g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

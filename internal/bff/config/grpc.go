package config

import (
	"errors"
	"net"
	"os"
)

const GrpcEnvUserClientHost = "GRPC_USER_CLIENT_HOST"
const GrpcEnvUserClientPort = "GRPC_USER_CLIENT_PORT"

const HttpEnvServerHost = "HTTP_SERVER_HOST"
const HttpEnvServerPort = "HTTP_SERVER_PORT"

type GrpcConfig interface {
	UserClientAddress() string
}

type HttpConfig interface {
	ServerAddress() string
}

type grpcConfig struct {
	userGrpcClientConfig
}

type httpConfig struct {
	httpServerConfig
}

type httpServerConfig struct {
	host string
	port string
}

type userGrpcClientConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GrpcConfig, error) {
	host := os.Getenv(GrpcEnvUserClientHost)
	if len(host) == 0 {
		return nil, errors.New("grpc user host not found")
	}

	port := os.Getenv(GrpcEnvUserClientPort)
	if len(port) == 0 {
		return nil, errors.New("grpc user port not found")
	}

	return &grpcConfig{
		userGrpcClientConfig: userGrpcClientConfig{
			host,
			port,
		},
	}, nil
}

func NewHttpConfig() (HttpConfig, error) {
	host := os.Getenv(HttpEnvServerHost)
	if len(host) == 0 {
		return nil, errors.New("http server host not found")
	}

	port := os.Getenv(HttpEnvServerPort)
	if len(port) == 0 {
		return nil, errors.New("http server port not found")
	}

	return &httpConfig{
		httpServerConfig: httpServerConfig{
			host,
			port,
		},
	}, nil
}

func (c *grpcConfig) UserClientAddress() string {
	return net.JoinHostPort(c.userGrpcClientConfig.host, c.userGrpcClientConfig.port)
}

func (c *httpConfig) ServerAddress() string {
	return net.JoinHostPort(c.httpServerConfig.host, c.httpServerConfig.port)
}

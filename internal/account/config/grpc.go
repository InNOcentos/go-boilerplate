package config

import (
	"errors"
	"net"
	"os"
)

const GrpcServerHost = "GRPC_SERVER_HOST"
const GrpcServerPort = "GRPC_SERVER_PORT"

type GrpcConfig interface {
	GrpcServerAddress() string
}

type grpcConfig struct {
	grpcServerConfig
}

type grpcServerConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GrpcConfig, error) {
	host := os.Getenv(GrpcServerHost)
	if len(host) == 0 {
		return nil, errors.New("grpc server host not found")
	}

	port := os.Getenv(GrpcServerPort)
	if len(port) == 0 {
		return nil, errors.New("grpc server port not found")
	}

	return &grpcConfig{
		grpcServerConfig: grpcServerConfig{
			host,
			port,
		},
	}, nil
}

func (c *grpcConfig) GrpcServerAddress() string {
	return net.JoinHostPort(c.grpcServerConfig.host, c.grpcServerConfig.port)
}

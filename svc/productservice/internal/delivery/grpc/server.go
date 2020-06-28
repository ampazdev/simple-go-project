package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	GRPCServer *grpc.Server
	Config     Config
}

type Config struct {
	Address           string
	ServerOptions     []grpc.ServerOption
	UnaryInterceptors []grpc.UnaryServerInterceptor
}

func NewServer(address string) *Server {
	cfg := Config{
		Address: address,
	}

	// Unary interceptors
	cfg.UnaryInterceptors = append(cfg.UnaryInterceptors)

	// Server options
	cfg.ServerOptions = append(cfg.ServerOptions)

	s := grpc.NewServer(cfg.ServerOptions...)
	reflection.Register(s)

	return &Server{
		GRPCServer: s,
		Config:     cfg,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.Config.Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := s.GRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}

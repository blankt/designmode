package functionaloptions

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxCons  int
	Tls      *tls.Config
}

type Option func(s *Server)

func MaxCons(maxCon int) Option {
	return func(s *Server) {
		s.MaxCons = maxCon
	}
}

func Protocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func TimeOut(timeOut time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeOut
	}
}

func Tls(tls *tls.Config) Option {
	return func(s *Server) {
		s.Tls = tls
	}
}

func NewServer(addr string, port int, options ...Option) Server {
	server := Server{
		Port: port,
		Addr: addr,
	}
	for _, option := range options {
		option(&server)
	}

	return server
}

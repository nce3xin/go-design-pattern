package builder

import "time"

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

type Option func(*Server)

func Protocol(proto string) Option {
	return func(s *Server) {
		s.Protocol = proto
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	s := &Server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(s)
	}
	return s, nil
}

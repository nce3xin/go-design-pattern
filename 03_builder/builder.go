package builder

import (
	"k8s.io/klog/v2"
	"time"
)

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

// usage
func example() {
	var s1, s2, s3 *Server
	s1, _ = NewServer("127.0.0.1", 1024)
	s2, _ = NewServer("127.0.0.1", 1024, Protocol("tcp"))
	s3, _ = NewServer("127.0.0.1", 1024, Protocol("tcp"),
		Timeout(60*time.Second), MaxConns(1000))
	klog.Infof("s1:%v\ns2:%v\ns3:%v", s1, s2, s3)
}

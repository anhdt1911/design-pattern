package main

import "fmt"

func main() {
	srv := New(
		WithHost("localhost"),
		WithPort(80),
	)
	srv.Start()
}

type Server struct {
	host string
	port int
}

// new variable new func
func New(options ...func(*Server)) *Server {
	srv := &Server{}
	for _, o := range options {
		o(srv)
	}
	return srv
}

func (s *Server) Start() {
	fmt.Printf("server started with host %d and port %s", s.port, s.host)
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}
func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

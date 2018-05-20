package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	starting = "starting"
	running  = "running"
	stopped  = "stoped"
)

type IServer interface {
	Start()
	Running() bool
	Stop()
}

type HTTPServer struct {
	server *http.Server
	status string
}

func NewServer(addr string) *HTTPServer {
	srv := &http.Server{
		Addr:           addr,
		Handler:        GetHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		server: srv,
		status: starting,
	}
}

func (s *HTTPServer) Start() {
	s.status = running
	err := s.server.ListenAndServe()
	if err != nil {
		s.status = stopped
		os.Exit(1)
	}
}

func (s *HTTPServer) Running() bool {
	return s.status == running
}

func (s *HTTPServer) Stop() {
	if s.server != nil {
		fmt.Println("Server is closed")
	}
}

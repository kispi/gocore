package gocore

import (
	"fmt"
	"net/http"
)

type Settings struct {
	Port int
}

type Server struct {
	Settings *Settings
}

func CreateServer(settings *Settings) *Server {
	return &Server{
		Settings: settings,
	}
}

func (s *Server) Run() {
	fmt.Printf(`Server starts on port :%d`, s.Settings.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Settings.Port), nil)
}

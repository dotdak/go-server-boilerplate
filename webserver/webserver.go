package webserver

import (
	"log"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	config *Config
	logger *log.Logger
	e      *echo.Echo
	Bolt   *Bolter
	secret []byte
}

type option func(*WebServer)

func NewWebServer(options ...option) *WebServer {

	s := &WebServer{
		logger: log.New(log.Writer(), "", log.LstdFlags),
	}

	SetConfig("config.conf")(s) // default config
	AddOptions(s, options...)

	s.e = LoadAPI(s)

	return s
}

func AddOptions(s *WebServer, options ...option) {
	for _, opt := range options {
		opt(s)
	}
}

func (s *WebServer) StartServer() error {
	return s.e.Start(s.config.Web.Addr)
}

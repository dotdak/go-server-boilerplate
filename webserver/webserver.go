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

func NewWebServer(options ...func(*WebServer)) *WebServer {

	s := &WebServer{
		logger: log.New(log.Writer(), "", log.LstdFlags),
	}

	SetConfig("config.conf")(s) // default config
	AddOptions(s, options...)

	s.e = LoadAPI(s)

	return s
}

func AddOptions(s *WebServer, options ...func(*WebServer)) {
	for _, option := range options {
		option(s)
	}
}

func (s *WebServer) StartServer() error {
	return s.e.Start(s.config.Web.Addr)
}

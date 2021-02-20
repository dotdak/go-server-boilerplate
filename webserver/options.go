package webserver

import (
	"log"
)

func SetConfig(path string) func(s *WebServer) {
	return func(s *WebServer) {
		config, err := LoadConfigFromFile(path)
		if err != nil {
			panic(err)
		}

		s.config = config
	}
}

func SetSecret(secret []byte) func(s *WebServer) {
	return func(s *WebServer) {
		s.secret = secret
	}
}

func SetLogger(logger *log.Logger) func(s *WebServer) {
	return func(s *WebServer) {
		s.logger = logger
	}
}

func LoadBoltDB() func(s *WebServer) {
	return func(s *WebServer) {
		s.Bolt = NewBolter(s.config.BoltDB.Path, s.logger)
	}
}

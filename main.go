package main

import (
	"webserver-boilerplate/webserver"
)

func main() {

	args := LoadArgs()

	s := webserver.NewWebServer(
		webserver.LoadConfig(args.ConfigFile),
		webserver.DefaultLogger(),
	)

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}

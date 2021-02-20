package main

import (
	"webserver-boilerplate/webserver"
)

func main() {

	args := LoadArgs()

	s := webserver.NewWebServer(
		webserver.SetConfig(args.ConfigFile),
		webserver.SetSecret(nil),
	)

	if args.Bolt {
		webserver.AddOptions(s, webserver.LoadBoltDB())
	}

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}

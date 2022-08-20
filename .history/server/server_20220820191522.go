package server

import "github.com/gorilla/mux"

type Config struct {
	Port string
}

type Broker struct {
	config *config
	route  mux.Route
}

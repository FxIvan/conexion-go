package server

import "github.com/gorilla/mux"

type Config struct {
	Port string
}

type Broker struct {
	config *Config
	route  *mux.Route
}

func Conexion(s *Config) (*Broker, error) {

	broker := &Broker{
		config: s,
		route:  mux.NewRouter(),
	}

	return broker, nil
}

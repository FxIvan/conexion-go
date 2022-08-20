package server

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

type Server interface {
	Conexion() *Config
}

type Broker struct {
	config *Config
	route  *mux.Router
}

func Conexion(s *Config) (*Broker, error) {

	broker := &Broker{
		config: s,
		route:  mux.NewRouter(),
	}

	fmt.Println(broker)

	return broker, nil
}

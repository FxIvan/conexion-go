package server

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string `:5050`
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

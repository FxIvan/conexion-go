package server

type Config struct {
	Port string
}

type server interface {
	Config() *Config
}

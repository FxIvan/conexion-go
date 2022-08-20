package server

type Config struct {
	Port string
}

type server interface {
	Conexion() *Config
}

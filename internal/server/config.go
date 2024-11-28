package server

type Config struct {
	ListenAddr string `yaml:"listenAddr" env:"LISTEN_ADDR" envDefault:"localhost:9080"`
}

package debug

type Config struct {
	ListenAddr string `yaml:"listenAddr" env:"LISTEN_ADDR"`
}

package config

type HTTPConfig struct {
	Address string `env:"HTTP_ADDRESS" default:""`
	Port    int    `env:"HTTP_PORT" default:"9000"`
}

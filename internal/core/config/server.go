package config

import "time"

type ServerConfig struct {
	Port            string
	Environment     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func (c *ServerConfig) IsDevelopment() bool {
	return c.Environment == "development"
}

func (c *ServerConfig) IsProduction() bool {
	return c.Environment == "production"
}

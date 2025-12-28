package interfaces

import "cleanandclean/internal/core/config"

// IConfig provides access to typed application configuration.
// The implementation can use Viper, env vars, or any other source.
type IConfig interface {
	Database() *config.DatabaseConfig
	Server() *config.ServerConfig
}

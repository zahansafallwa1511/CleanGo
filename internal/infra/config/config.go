package config

import (
	"fmt"
	"strings"
	"time"

	coreConfig "cleanandclean/internal/core/config"

	"github.com/spf13/viper"
)

// viperConfig holds the raw structure for Viper unmarshaling.
// Uses mapstructure tags to map YAML keys to struct fields.
type viperConfig struct {
	Database viperDatabaseConfig `mapstructure:"database"`
	Server   viperServerConfig   `mapstructure:"server"`
}

type viperDatabaseConfig struct {
	URL             string `mapstructure:"url"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
}

type viperServerConfig struct {
	Port            string `mapstructure:"port"`
	Environment     string `mapstructure:"environment"`
	ReadTimeout     string `mapstructure:"read_timeout"`
	WriteTimeout    string `mapstructure:"write_timeout"`
	ShutdownTimeout string `mapstructure:"shutdown_timeout"`
}

// Config implements IConfig using Viper as the backend.
type Config struct {
	database *coreConfig.DatabaseConfig
	server   *coreConfig.ServerConfig
}

// NewConfig loads configuration from files and environment variables.
func NewConfig() (*Config, error) {
	v := viper.New()

	// Config file settings
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	// Environment variables override config file
	// APP_DATABASE_URL -> database.url
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Set defaults
	setDefaults(v)

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	// Unmarshal into intermediate struct
	var raw viperConfig
	if err := v.Unmarshal(&raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Convert to core config structs
	cfg, err := toConfig(&raw)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func setDefaults(v *viper.Viper) {
	// Database defaults
	v.SetDefault("database.max_open_conns", 25)
	v.SetDefault("database.max_idle_conns", 5)
	v.SetDefault("database.conn_max_lifetime", "5m")

	// Server defaults
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.environment", "development")
	v.SetDefault("server.read_timeout", "10s")
	v.SetDefault("server.write_timeout", "10s")
	v.SetDefault("server.shutdown_timeout", "30s")
}

func toConfig(raw *viperConfig) (*Config, error) {
	connMaxLifetime, err := time.ParseDuration(raw.Database.ConnMaxLifetime)
	if err != nil {
		return nil, fmt.Errorf("invalid database.conn_max_lifetime: %w", err)
	}

	readTimeout, err := time.ParseDuration(raw.Server.ReadTimeout)
	if err != nil {
		return nil, fmt.Errorf("invalid server.read_timeout: %w", err)
	}

	writeTimeout, err := time.ParseDuration(raw.Server.WriteTimeout)
	if err != nil {
		return nil, fmt.Errorf("invalid server.write_timeout: %w", err)
	}

	shutdownTimeout, err := time.ParseDuration(raw.Server.ShutdownTimeout)
	if err != nil {
		return nil, fmt.Errorf("invalid server.shutdown_timeout: %w", err)
	}

	return &Config{
		database: &coreConfig.DatabaseConfig{
			URL:             raw.Database.URL,
			MaxOpenConns:    raw.Database.MaxOpenConns,
			MaxIdleConns:    raw.Database.MaxIdleConns,
			ConnMaxLifetime: connMaxLifetime,
		},
		server: &coreConfig.ServerConfig{
			Port:            raw.Server.Port,
			Environment:     raw.Server.Environment,
			ReadTimeout:     readTimeout,
			WriteTimeout:    writeTimeout,
			ShutdownTimeout: shutdownTimeout,
		},
	}, nil
}

func (c *Config) Database() *coreConfig.DatabaseConfig {
	return c.database
}

func (c *Config) Server() *coreConfig.ServerConfig {
	return c.server
}

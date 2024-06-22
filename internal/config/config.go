package config

import "log"

// Config holds the application configuration
type Config struct {
    LogLevel  string
    Port      int
    GlobalCSS []string
}

// InitConfig initializes the application configuration with hard-coded values
func InitConfig() *Config {
    config := &Config{
        LogLevel:  "DEBUG",
        Port:      8080,
        GlobalCSS: []string{"reset.min.css", "main.min.css"},
    }

    log.Printf("Configuration initialized: %+v\n", config)
    return config
}


package config

import (
	"os"
)

// Config holds application configuration values
type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
	Headless         bool
	MaxDailyActions  int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
		Headless:         false, // keep false for demo & stealth
		MaxDailyActions:  20,    // safe demo limit
	}
}

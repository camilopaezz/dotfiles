package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// PackageConfig represents the package configuration
type PackageConfig struct {
	Official []string `json:"official"`
	AUR      []string `json:"aur"`
}

// DotfilesConfig represents the structure of dotfiles.json
type DotfilesConfig map[string]string

// Config represents the complete configuration structure
type Config struct {
	Dotfiles map[string]string `json:"dotfiles"`
	Packages PackageConfig     `json:"packages"`
}

// LoadConfig reads and parses the configuration file
func LoadConfig(configPath string) (DotfilesConfig, error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Try to parse as new format first
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		// If that fails, try to parse as old format for backward compatibility
		var dotfiles DotfilesConfig
		if err := json.Unmarshal(configData, &dotfiles); err != nil {
			return nil, fmt.Errorf("error parsing config file: %w", err)
		}
		return dotfiles, nil
	}

	return config.Dotfiles, nil
}

// LoadCompleteConfig reads and parses the complete configuration file
func LoadCompleteConfig(configPath string) (Config, error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

// GetAllPackages returns a combined list of all packages (official + AUR)
func (pc PackageConfig) GetAllPackages() []string {
	var all []string
	all = append(all, pc.Official...)
	all = append(all, pc.AUR...)
	return all
}
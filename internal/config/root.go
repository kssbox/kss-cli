package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var GlobalConfig Config

func InitConfig() {
	// path := filepath.Join(filepath.Dir("."), "config/config.toml")
	// path := filepath.Join(filepath.Dir("."), "../../", "config/config.toml")
	path := filepath.Join(filepath.Dir("."), "../../../", "config/config.toml")
	if _, err := toml.DecodeFile(path, &GlobalConfig); err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
}

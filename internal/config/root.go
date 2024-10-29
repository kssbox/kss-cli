package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var GlobalConfig Config

func InitConfig() {
	// 加载配置文件
	if _, err := toml.DecodeFile(filepath.Join("config", "config.toml"), &GlobalConfig); err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
}

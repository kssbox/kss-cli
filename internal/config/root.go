package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var GlobalConfig Config

func InitConfig() {
	// 从当前目录开始判断一直到 / 目录
	for dir := "."; dir != "/"; dir = filepath.Dir(dir) {
		path := filepath.Join(dir, ".kssrc")
		if _, err := os.Stat(path); err == nil {
			// 找到 .kssrc 文件，加载配置
			if _, err := toml.DecodeFile(path, &GlobalConfig); err != nil {
				log.Fatalf("Failed to load config file: %v", err)
			}
			return
		}
	}
	// 如果都没找到，则在用户目录查找 .kssrc 文件
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	userConfigPath := filepath.Join(homeDir, ".kssrc")
	if _, err := os.Stat(userConfigPath); err == nil {
		if _, err := toml.DecodeFile(userConfigPath, &GlobalConfig); err != nil {
			log.Fatalf("Failed to load user config file: %v", err)
		}
		return
	}

	// 如果都没找到，则在用户目录生成 .kssrc 文件
	defaultConfig := `[github]
	api_url = "https://api.github.com/user/repos"
	token = "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"
	owner = "Kevinbrother"
	repos = "git@github.com"

	[github.apis]
	add_repos = "https://api.github.com/user/repos"
	delete_repos = "https://api.github.com/repos"
	`

	// 创建 .kssrc 文件并写入默认值
	err = os.WriteFile(filepath.Join(homeDir, ".kssrc"), []byte(defaultConfig), 0644)
	if err != nil {
		log.Fatalf("Failed to create default config file: %v", err)
	}
}

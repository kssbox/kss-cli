package github_api

import (
	"kssbox/kss-cli/kss-cli/internal/config"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestCreateRepo(t *testing.T) {
	CreateRepo("test", "test", false)
}

func TestDeleteRepo(t *testing.T) {
	DeleteRepo("test")
}

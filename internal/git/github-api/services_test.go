package github_api

import (
	"kssbox/kss-cli/kss-cli/internal/config"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestCreateRepo(t *testing.T) {
	CreateRepo("kss-cli", "test", true)
}

func TestDeleteRepo(t *testing.T) {
	DeleteRepo("Kevinbrother", "local_test")
}

func TestGetOrganizations(t *testing.T) {
	GetOrganizations()
}

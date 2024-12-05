package github_api

import (
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"net/http"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestCreateRepo(t *testing.T) {
	resp, err := CreateRepo("local_test", "test", true)
	if err != nil {
		t.Errorf("failed to create repository: %v", err)
	}

	fmt.Printf("rsp status: %v\n", resp.Status)
	// 检查响应状态
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("failed to create repository, status: %s", resp.Status)
	}
}

func TestDeleteRepo(t *testing.T) {
	resp, err := DeleteRepo("Kevinbrother", "local_test")
	if err != nil {
		t.Errorf("failed to delete repository: %v", err)
	}

	fmt.Printf("rsp status: %v\n", resp.Status)
	// 检查响应状态
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("failed to delete repository, status: %s", resp.Status)
	}
}

func TestGetOrganizations(t *testing.T) {
	resp, err := GetOrganizations()
	if err != nil {
		t.Errorf("failed to get organizations: %v", err)
	}

	fmt.Printf("rsp status: %v\n", resp.Status)
	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		t.Errorf("failed to get organizations, status: %s", resp.Status)
	}
}

package github_api

import (
	"kssbox/kss-cli/kss-cli/internal/config"
	"net/http"
)

// 创建仓库
func CreateRepo(name, description string, private bool) (*http.Response, error) {
	repo := RepoConfig{
		Name:        name,
		Description: description,
		Private:     private,
	}

	return GitHubAPI("POST", config.GlobalConfig.GitHub.GitHubAPIs.AddRepos, repo)

}

// 删除仓库
func DeleteRepo(OWNER, REPO string) (*http.Response, error) {
	return GitHubAPI("DELETE", config.GlobalConfig.GitHub.GitHubAPIs.DeleteRepos+"/"+OWNER+"/"+REPO, nil)
}

// 获取 organizations 仓库
func GetOrganizations() (*http.Response, error) {
	return GitHubAPI("GET", config.GlobalConfig.GitHub.GitHubAPIs.GetOrganizations, nil)
}

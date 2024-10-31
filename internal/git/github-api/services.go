package github_api

import "kssbox/kss-cli/kss-cli/internal/config"

// 创建仓库
func CreateRepo(name, description string, private bool) error {
	repo := RepoConfig{
		Name:        name,
		Description: description,
		Private:     private,
	}

	return GitHubAPI("POST", config.GlobalConfig.GitHub.GitHubAPIs.AddRepos, repo)

}

// 删除仓库
func DeleteRepo(OWNER, REPO string) error {
	return GitHubAPI("DELETE", config.GlobalConfig.GitHub.GitHubAPIs.DeleteRepos+OWNER+"/"+REPO, nil)
}

package github_api

// 创建仓库
func CreateRepo(name, description string, private bool) error {
	repo := RepoConfig{
		Name:        name,
		Description: description,
		Private:     private,
	}

	return GitHubAPI("POST", "https://api.github.com/user/repos", repo)

}

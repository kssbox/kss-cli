package config

// Config 用于映射配置文件中的数据
type Config struct {
	GitHub struct {
		APIURL string `toml:"api_url"`
		Token  string `toml:"token"`
		Owner  string `toml:"owner"`
		Repos  string `toml:"repos"`

		GitHubAPIs struct {
			AddRepos         string `toml:"add_repos"`
			DeleteRepos      string `toml:"delete_repos"`
			GetOrganizations string `toml:"get_organizations"`
		} `toml:"apis"` // 修改为 "apis" 以对应 config.toml 中的定义
	} `toml:"github"`
}

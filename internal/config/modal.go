package config

// Config 用于映射配置文件中的数据
type Config struct {
	GitHub struct {
		APIURL string `toml:"api_url"`
		Token  string `toml:"token"`
		Owner  string `toml:"owner"`
		Repos  string `toml:"repos"`

		GitHubAPIs struct {
			AddRepos    string `toml:"add_repos"`
			DeleteRepos string `toml:"delete_repos"`
		} `toml:"github.apis"`
	} `toml:"github"`
}

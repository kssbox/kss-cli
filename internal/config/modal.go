package config

// Config 用于映射配置文件中的数据
type Config struct {
	GitHub struct {
		APIURL string `toml:"api_url"`
		Token  string `toml:"token"`
	} `toml:"github"`
}

/**
@author: kssbox
@date: 2024-10-29
@desc: 处理 GitHub API 请求
*/

package github_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"net/http"
)

// 接受 方法、url、请求体
// data 符合 json 的结构体
func GitHubAPI(method, url string, data interface{}) (*http.Response, error) {

	fmt.Printf("GitHubAPI: %v, %v, %v\n", method, url, data)

	// 将请求数据转换为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	fmt.Printf("token: %v\n", config.GlobalConfig.GitHub.Token)
	// 设置请求头，包括认证信息和请求类型
	req.Header.Set("Authorization", "Bearer "+config.GlobalConfig.GitHub.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}

	// 打印请求信息
	fmt.Printf("Request: %v\n", req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	return resp, nil
}

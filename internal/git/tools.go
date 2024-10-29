package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// 初始化本地 Git 仓库
func initLocalRepo(localPath string) error {
	fmt.Printf("Initializing local repository at %s...\n", localPath)

	if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = localPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// 增加必要的文件 如 README.md、.gitignore 等
func addFiles(localPath string) error {
	// README.md
	readmeContent := "# Your README.md content here\n"
	readmePath := filepath.Join(localPath, "README.md")

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("could not create README.md: %v", err)
	}

	//.gitignore
	gitignoreContent := "node_modules\n"
	gitignorePath := filepath.Join(localPath, ".gitignore")

	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("could not create.gitignore: %v", err)
	}

	return nil
}

func createGitHubRepo(name, description string, private bool) error {
	// 创建请求体
	repo := RepoConfig{
		Name:        name,
		Description: description,
		Private:     private,
	}

	// 将请求数据转换为 JSON
	jsonData, err := json.Marshal(repo)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", config.GlobalConfig.GitHub.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// 设置请求头，包括认证信息和请求类型
	req.Header.Set("Authorization", "Bearer "+config.GlobalConfig.GitHub.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create repository, status: %s", resp.Status)
	}

	return nil
}

// 添加远程仓库，并推送 main 分支
func addRemoteRepo(localPath, remoteURL string) error {
	fmt.Printf("Adding remote repository %s...\n", remoteURL)

	cmd := exec.Command("git", "remote", "add", "origin", remoteURL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add remote repository: %v", err)
	}

	cmd = exec.Command("git", "push", "-u", "origin", "main")
	cmd.Dir = localPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push to remote repository: %v", err)
	}

	return nil

}

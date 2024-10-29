package git

import (
	"kssbox/kss-cli/kss-cli/internal/config"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var tempDir = "./local_test"
var remoteURL = "https://github.com/kevinbrother/example.git"

// 测试之前统一运行的代码
func init() {
	config.InitConfig()
}

func TestInitLocalRepo(t *testing.T) {
	// tempDir := t.TempDir() // 创建一个临时目录
	// 在当前目录下创建一个临时目录

	// 测试 Git 仓库初始化
	if err := initLocalRepo(tempDir); err != nil {
		t.Fatalf("Failed to initialize local repository: %v", err)
	}

	// 检查 .git 文件夹是否存在，判断仓库是否成功创建
	gitDir := filepath.Join(tempDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Fatalf("Expected .git directory to exist, but it doesn't")
	}
}

func TestAddFile(t *testing.T) {
	// tempDir := t.TempDir()

	// 先初始化本地仓库
	if err := initLocalRepo(tempDir); err != nil {
		t.Fatalf("Failed to initialize local repository: %v", err)
	}
}

func TestAddRemoteRepo(t *testing.T) {
	// tempDir := t.TempDir()

	// 先初始化本地仓库
	if err := initLocalRepo(tempDir); err != nil {
		t.Fatalf("Failed to initialize local repository: %v", err)
	}

	// 测试添加远程仓库

	if err := addRemoteRepo(tempDir, remoteURL); err != nil {
		t.Fatalf("Failed to add remote repository: %v", err)
	}

	// 验证远程仓库 URL 是否已正确添加
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = tempDir
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get remote URL: %v", err)
	}

	if string(output) != remoteURL+"\n" { // 检查输出是否与期望的 URL 匹配
		t.Fatalf("Expected remote URL %s, but got %s", remoteURL, string(output))
	}
}

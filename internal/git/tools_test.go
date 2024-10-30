package git

import (
	"kssbox/kss-cli/kss-cli/internal/config"
	"os"
	"path/filepath"
	"testing"
)

var tempDir = "../../local_test"
var remoteURL = "git@github.com:KevinBrother/local_test.git"

var name = "local_test"
var description = "local_test"
var private = true

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
	if err := addFiles(tempDir); err != nil {
		t.Fatalf("Failed to initialize local repository: %v", err)
	}
}

func TestCreateGitHubRepo(t *testing.T) {
	if err := createGitHubRepo(name, description, private); err != nil {
		t.Fatalf("Failed to create remote repository: %v", err)
	}
}

func TestAddRemoteRepo(t *testing.T) {
	if err := addRemoteRepo(tempDir, remoteURL); err != nil {
		t.Fatalf("Failed to add remote repository: %v", err)
	}
}

func TestPushMainBranch(t *testing.T) {
	if err := pushMainBranch(tempDir, ""); err != nil {
		t.Fatalf("Failed to push main branch: %v", err)
	}
}

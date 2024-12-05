package git

import (
	"fmt"
	github_api "kssbox/kss-cli/kss-cli/internal/git/github-api"
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
	_, err := github_api.CreateRepo(name, description, private)
	return err
}

// 进入到 localPath 目录下添加远程仓库
func addRemoteRepo(localPath, remoteURL string) error {
	fmt.Println(localPath)
	fmt.Printf("Adding remote repository %s...\n", remoteURL)

	cmd := exec.Command("git", "remote", "add", "origin", remoteURL)
	cmd.Dir = localPath
	fmt.Println(cmd.Args)
	return cmd.Run()
}

// git remote add origin git@github.com:KevinBrother/local_test.git
// git remote add origin git@github.com:KevinBrother/local_test.git

func initCommit(localPath string) error {

	cmd := exec.Command("git", "add", ".")
	cmd.Dir = localPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add files: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", "initial commit")
	cmd.Dir = localPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to commit: %v", err)
	}

	return nil
}

// 推送 main 分支
func pushMainBranch(localPath, branch string) error {
	if branch == "" {
		branch = "main"
	}
	cmd := exec.Command("git", "push", "-u", "origin", "main")
	cmd.Dir = localPath
	fmt.Println(cmd.Args)
	return cmd.Run()
}

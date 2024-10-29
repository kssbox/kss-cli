package internal

import (
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"kssbox/kss-cli/kss-cli/internal/git"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <local-path> <remote-url>")
		return
	}

	config.InitConfig()

	// localPath := os.Args[1]
	// remoteURL := os.Args[2]

	git.Run()

	fmt.Println("Git repository initialized and remote added successfully!")
}

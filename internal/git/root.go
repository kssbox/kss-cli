package git

import (
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"log"
	"os"
)

func Run() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <local-path> <remote-url>")
		return
	}

	localPath := os.Args[1]
	remoteURL := os.Args[2]

	err := initLocalRepo(localPath)
	if err != nil {
		log.Fatalf("Failed to initialize local repository: %v\n", err)
	}

	err = addFiles(localPath)
	if err != nil {
		log.Fatalf("Failed to create remote repository: %v\n", err)
	}

	err = addRemoteRepo(localPath, remoteURL)
	if err != nil {
		log.Fatalf("Failed to add remote repository: %v\n", err)
	}

	fmt.Println(config.GlobalConfig.GitHub.APIURL)

	fmt.Println("Git repository initialized and remote added successfully!")
}

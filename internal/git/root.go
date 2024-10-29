package git

import (
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"log"
)

func Run() {
	// 改成问答形式

	var localPath, remoteURL, name, description string
	var private bool

	fmt.Print("Please enter the local path: ")
	fmt.Scanln(&localPath)

	fmt.Print("Please enter the remote URL: ")
	fmt.Scanln(&remoteURL)

	fmt.Print("Please enter the repository name: ")
	fmt.Scanln(&name)

	fmt.Print("Please enter the repository description: ")
	fmt.Scanln(&description)

	fmt.Print("Please enter the repository private (true/false): ")
	fmt.Scanln(&private)

	err := initLocalRepo(localPath)
	if err != nil {
		log.Fatalf("Failed to initialize local repository: %v\n", err)
	}

	err = addFiles(localPath)
	if err != nil {
		log.Fatalf("Failed to create remote repository: %v\n", err)
	}

	err = createGitHubRepo(name, description, private)
	if err != nil {
		log.Fatalf("Failed to initialize remote repository: %v\n", err)
	}

	err = addRemoteRepo(localPath, remoteURL)
	if err != nil {
		log.Fatalf("Failed to add remote repository: %v\n", err)
	}

	fmt.Println(config.GlobalConfig.GitHub.APIURL)

	fmt.Println("Git repository initialized and remote added successfully!")
}
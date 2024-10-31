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

	err := initLocalRepo(localPath)
	if err != nil {
		log.Fatalf("Failed to initialize local repository: %v\n", err)
	}

	err = addFiles(localPath)
	if err != nil {
		log.Fatalf("Failed to create remote repository: %v\n", err)
	}

	fmt.Print("Please enter the repository name: ")
	fmt.Scanln(&name)

	fmt.Print("Please enter the repository description: ")
	fmt.Scanln(&description)

	fmt.Print("Please enter the repository private (true/false): ")
	fmt.Scanln(&private)

	err = createGitHubRepo(name, description, private)
	if err != nil {
		log.Fatalf("Failed to initialize remote repository: %v\n", err)
	}

	remoteURL = config.GlobalConfig.GitHub.Repos + ":" + config.GlobalConfig.GitHub.Owner + "/" + name + ".git"

	err = addRemoteRepo(localPath, remoteURL)
	if err != nil {
		log.Fatalf("Failed to add remote repository: %v\n", err)
	}

	err = pushMainBranch(localPath, remoteURL)
	if err != nil {
		log.Fatalf("Failed to push main branch: %v\n", err)
	}

	fmt.Println("Git repository initialized and remote added successfully!")
}

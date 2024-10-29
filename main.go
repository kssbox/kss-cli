package main

import (
	"fmt"
	"kssbox/kss-cli/kss-cli/internal/config"
	"kssbox/kss-cli/kss-cli/internal/git"
)

func main() {

	config.InitConfig()

	git.Run()

	fmt.Println("Git repository initialized and remote added successfully!")
}

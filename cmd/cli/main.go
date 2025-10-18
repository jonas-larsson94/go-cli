package main

import (
	"fmt"
	"log"

	"github.com/jonas-larsson94/go-cli/internal/github"
)

func main() {
	fmt.Println("Starting myapp...")
	branches, err := github.GetBranches()
	if err != nil {
		log.Fatal("Cant read remote branches")
	}

	fmt.Println(branches)
}

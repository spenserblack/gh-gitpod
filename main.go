package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
)

func main() {
	repo, err := gh.CurrentRepository()
	onError(err)

	fullRepo := fmt.Sprintf("%s/%s/%s", repo.Host(), repo.Owner(), repo.Name())

	fmt.Println(fullRepo)

	b := browser.New("", os.Stdout, os.Stderr)

	err = b.Browse(fmt.Sprintf("https://gitpod.io/#%s", fullRepo))
}

func onError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

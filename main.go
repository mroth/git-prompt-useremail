package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// TODO: add/delete functions to add new shortcuts
// e.g. git-usermail add mroth@foo.com SYMBOL

func main() {
	flag.Parse()

	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		// TODO: die loudly option
		os.Exit(0)
	}

	lconfig, err := repo.ConfigScoped(config.LocalScope)
	if err != nil {
		// TODO: die loudly option
		os.Exit(0)
	}

	// TODO: consider memoize retrieving global config, in case not needed (unlikley, but why not)
	gconfig, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		// TODO: die loudly option
		os.Exit(0)
	}

	email := os.Getenv("GIT_AUTHOR_EMAIL")
	if email == "" {
		email = lconfig.User.Email
	}
	if email == "" {
		email = gconfig.User.Email
	}

	var prompt string
	prompt = checkForPrompt(lconfig, email)
	if prompt == "" {
		prompt = checkForPrompt(gconfig, email)
	}
	if prompt == "" {
		prompt = "❔"
	}

	fmt.Println(prompt)
}

func checkForPrompt(conf *config.Config, email string) string {
	if conf.Raw.HasSection("emailprompt") {
		sec := conf.Raw.Section("emailprompt")
		for _, ss := range sec.Subsections {
			if ss.Name == email {
				if val := ss.Option("prompt"); val != "" {
					return val
				}
			}
		}
	}
	return ""
}

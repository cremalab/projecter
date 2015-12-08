package main

import (
	"fmt"
	"os"
	"os/user"
)

func help() {
	fmt.Println(helpMessage)
}

func printError(msg string) {
	fmt.Println(msg)
}

func main() {
	address := os.Getenv("SCOREBOT_ADDR")

	providers := []provider{
		githubSourceProvider{},
	}
	if len(address) == 0 {
		//printError(configErrorMessage)
		os.Exit(1)
	}

	defaultConfig := config{}
	u, err := user.Current()
	if err != nil {
		fmt.Println("Unable to determin current user. Exiting.")
		os.Exit(1)
	}
	home := u.HomeDir
	start, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to determine current working directory. Exiting.")
	}

	c, err := synthesizeConfig(defaultConfig, home, start)

	if err != nil {
		fmt.Println("Error reading config: " + err.Error())
		os.Exit(1)
	}

	// Strip name of this binary from args
	_, args := os.Args[0], os.Args[1:]

	// Allowed commands effectively defined here.
	// Note: if the providers don't check to make sure the command they want exists
	// before trying to add it (and the providers I've written *do not*) and it's
	// absent here, you'll see a runtime error.
	r := map[string][]route{
		"status": []route{},
		"init":   []route{},
	}

	//test, _ := yaml.Marshal(c)
	//fmt.Println(string(test))
	//os.Exit(0)

	for _, p := range providers {
		r = p.AddRoutes(r, c)
	}

	applyRoute(r, args)
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var client SlackAPI

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  slackapi api.test")
		fmt.Println("  slackapi auth.test")
		fmt.Println("  slackapi users.list")
		fmt.Println("  slackapi users.search [query]")
		flag.PrintDefaults()
	}

	flag.Parse()
	client.AutoConfigure()

	switch flag.Arg(0) {
	case "api.test":
		client.Test()
	case "auth.test":
		client.AuthTest()
	case "users.list":
		client.UsersList()
	case "users.search":
		client.UsersSearch(flag.Arg(1))
	case "-help":
		flag.Usage()
	}

	os.Exit(0)
}

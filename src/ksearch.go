package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("possible arguments")
	fmt.Println("short  full       name       description")
	fmt.Println("h      help       HELP       this here")
	fmt.Println("n      namespace  NAMESPACE  search for a namespace")
}

func printToStatusLine(message string) {
	fmt.Print("\033[G\033[K")
	fmt.Print(message)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments given")
		printHelp()
		return
	}

	if args[0] == "h" || args[0] == "help" {
		printHelp()
		return
	}

	results := []string{}

	if args[0] == "n" || args[0] == "namespace" {
		if len(args) == 1 {
			fmt.Println("No namespace search expression given")
			return
		}
		namespace := Namespace{}
		results = namespace.find(args[1])
	}

	if args[0] == "i" || args[0] == "ingress" {

		if len(args) < 3 {
			fmt.Println("No ingress search expression given")
			fmt.Println("search 'ksearch ingress {pattern} {searchExpression} {(optional) clusterFilter}")
			return
		}

		searchPattern := args[1]
		searchText := args[2]
		clusterFilter := ""

		if len(args) == 4 {
			clusterFilter = args[3]
		}

		if searchPattern == "url" {
			ingress := Ingress{}
			results = ingress.findByUrlV2(searchText, clusterFilter)
		} else {
			fmt.Println("no search pattern found, use 'url'")
		}
	}

	// print result
	if len(results) > 0 {
		printToStatusLine("done")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Found:")
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		fmt.Println("No results found")
	}
}

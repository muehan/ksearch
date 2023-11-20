package main

import (
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("possible arguments")
	fmt.Println("short  full       name       description")
	fmt.Println("h      help       HELP       this here")
	fmt.Println("n      namespace  NAMESPACE  search for a namespace")
}

func printToStatusLine(message string) {
	width := 100
	fill := " "
	padding := strings.Repeat(fill, width-len(message))
	result := message + padding
	fmt.Print("\r" + result)
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

	// print result
	if len(results) > 0 {
		printToStatusLine("done")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Found:")
		for _, result := range results {
			fmt.Println(result)
		}
	}
}

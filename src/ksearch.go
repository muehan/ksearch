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

	var parser QueryParserer = QueryParser{}
	parser.Print(args)
}

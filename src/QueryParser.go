package main

import "fmt"

type QueryParserer interface {
	Print(args []string)
}

type QueryParser struct{}

func (q QueryParser) Print(args []string) {
	if args[0] == "n" || args[0] == "namespace" {
		if len(args) == 1 {
			fmt.Println("No namespace search expression given")
			return
		}
		namespace := Namespace{}
		results := namespace.find(args[1])
		printResult(results)
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

		if searchPattern == "url" {
			if len(args) == 4 {
				clusterFilter = args[3]
			}
			ingress := Ingress{}
			results := ingress.findByUrlV2(searchText, clusterFilter)
			printResult(results)
		} else if searchPattern == "serversnippet" {
			if len(args) == 3 {
				clusterFilter = args[2]
			}
			ingress := Ingress{}
			results := ingress.findServerSnippets(clusterFilter)
			printResult(results)
		} else {
			fmt.Println("no search pattern found, use 'url'")
		}
	}
}

func printResult(results []string) {
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

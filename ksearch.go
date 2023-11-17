package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("possible arguments")
	fmt.Println("-h HELP this here")
	fmt.Println("-n NAMESPACE search for a namespace")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments given")
		printHelp()
		return
	}

	if args[0] == "-h" {
		printHelp()
		return
	}

	kubectl := KubectlImpl{}

	if args[0] == "-n" {
		fmt.Println("searching for namespace")
		clusters := kubectl.GetClusters()
		for _, cluster := range clusters {
			// fmt.Print("checking cluster '" + cluster + "'")
			kubectl.SetCluster(cluster)
			namespaces := kubectl.GetNamespaces(args[1])
			if len(namespaces) > 0 {
				for _, namespace := range namespaces {
					fmt.Println("Namespace '" + namespace + "' found in cluster '" + cluster + "'")
				}
			}
		}
	}

}

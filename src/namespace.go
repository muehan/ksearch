package main

import "fmt"

type Namespace struct{}

func (namespace Namespace) find(searchString string) []string {
	kubectl := KubectlImpl{}
	results := []string{}

	fmt.Println("searching for namespace")

	clusters := kubectl.GetClusters()
	for _, cluster := range clusters {
		printToStatusLine("search in Cluster: " + cluster)
		kubectl.SetCluster(cluster)
		namespaces := kubectl.GetNamespaces(searchString)
		if len(namespaces) > 0 {
			for _, namespace := range namespaces {
				message := "Namespace '" + namespace + "' found in cluster '" + cluster + "'"
				results = append(results, message)
			}
		}
	}

	return results
}

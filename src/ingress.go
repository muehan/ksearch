package main

import (
	"fmt"

	"github.com/ahmetb/go-linq"
)

type Ingress struct{}

func (ingress Ingress) find(searchString string, clusterFilter string) []string {
	kubectl := KubectlImpl{}
	results := []string{}

	fmt.Println("searching for ingress hosts")

	clusters := kubectl.GetClusters()

	if clusterFilter != "" {
		filteredClusters := []string{}
		linq.From(clusters).WhereT(func(c string) bool {
			return c == clusterFilter
		}).ToSlice(&filteredClusters)
		clusters = filteredClusters
	}

	for _, cluster := range clusters {
		kubectl.SetCluster(cluster)
		namespaces := kubectl.GetNamespaces()
		if len(namespaces) > 0 {
			for _, namespace := range namespaces {
				printToStatusLine("search in Cluster: " + cluster + " namespace: " + namespace)
				ingresses := kubectl.GetIngressesByUrl(namespace, searchString)
				if ingresses != nil {
					for _, ingress := range ingresses {
						if ingress != "" {
							ingress = ingress + " found in " + namespace + " from cluster: " + cluster
							results = append(results, ingress)
						}
					}
				}
			}
		}
	}

	return results
}

func (ingress Ingress) findByUrlV2(searchString string, clusterFilter string) []string {
	kubectl := KubectlImpl{}
	results := []string{}

	fmt.Println("searching for ingress hosts")

	clusters := kubectl.GetClusters()

	if clusterFilter != "" {
		filteredClusters := []string{}
		linq.From(clusters).WhereT(func(c string) bool {
			return c == clusterFilter
		}).ToSlice(&filteredClusters)
		clusters = filteredClusters
	}

	for _, cluster := range clusters {
		kubectl.SetCluster(cluster)
		ingresses := kubectl.GetIngressesByUrlV2(searchString)
		for _, ingress := range ingresses {
			if ingress != "" {
				ingress = ingress + " in cluster: " + cluster
				results = append(results, ingress)
			}
		}
	}

	return results
}

func (ingress Ingress) findServerSnippets(clusterFilter string) []string {
	kubectl := KubectlImpl{}
	results := []string{}

	fmt.Println("searching for ingress hosts")

	clusters := kubectl.GetClusters()

	if clusterFilter != "" {
		filteredClusters := []string{}
		linq.From(clusters).WhereT(func(c string) bool {
			return c == clusterFilter
		}).ToSlice(&filteredClusters)
		clusters = filteredClusters
	}

	for _, cluster := range clusters {
		kubectl.SetCluster(cluster)
		ingresses := kubectl.GetIngressesServerSnippets()
		for _, ingress := range ingresses {
			if ingress != "" {
				ingress = ingress + " in cluster: " + cluster
				results = append(results, ingress)
			}
		}
	}

	return results
}

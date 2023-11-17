package main

import (
	"os"
	"os/exec"
	"strings"
)

type Kubectl interface {
	GetClusters() []string
	GetNamespaces() []string

	SetCluster(contextname string)
}

type KubectlImpl struct{}

func (kubectl KubectlImpl) GetClusters() []string {
	clusterString := runCommand("kubectl config get-contexts -o name")
	clusters := strings.Split(clusterString, "\n")
	retval := []string{}
	for _, cluster := range clusters {
		if cluster != "" || !strings.Contains(cluster, "rancher") {
			retval = append(retval, cluster)
		}
	}
	return retval
}

func (kubectl KubectlImpl) GetNamespaces(filter string) []string {
	command := "kubectl get namespaces --output=jsonpath='{range .items[*]}{.metadata.name}{\"\\n\"}{end}'	"
	if filter != "" {
		command += " | grep " + filter
	}
	namespacesString := runCommand(command)
	if namespacesString == "" {
		return []string{}
	}
	namespaces := strings.Split(namespacesString, "\n")
	retval := []string{}
	for _, namespace := range namespaces {
		if namespace != "" {
			retval = append(retval, namespace)
		}
	}
	return retval
}

func (kubectl KubectlImpl) SetCluster(contextname string) {
	runCommand("kubectl config use-context " + contextname)
}

func runCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	output, _ := cmd.Output()

	if len(output) == 0 {
		return ""
	}

	outputStr := string(output)

	return outputStr
}

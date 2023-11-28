package main

import (
	"os"
	"os/exec"
	"strings"
)

type Kubectl interface {
	GetClusters() []string
	GetNamespaces() []string
	GetIngresses(namespace string) []string

	SetCluster(contextname string)
}

type KubectlImpl struct{}

func (kubectl KubectlImpl) GetClusters() []string {
	clusterString := runCommand("kubectl config get-contexts -o name")
	clusters := strings.Split(clusterString, "\n")
	retval := []string{}
	for _, cluster := range clusters {
		if cluster != "" && !strings.Contains(cluster, "rancher") {
			retval = append(retval, cluster)
		}
	}
	return retval
}

func (kubectl KubectlImpl) GetNamespaces(filter ...string) []string {
	command := "kubectl get namespaces --output=jsonpath='{range .items[*]}{.metadata.name}{\"\\n\"}{end}'"
	if len(filter) > 0 {
		command += " | grep " + filter[0]
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

func (kubectl KubectlImpl) GetIngresses(namespace string, search string) []string {
	command := "kubectl get ingresses --output=jsonpath='{range .items[*].spec.rules[*]}{.host}{\"\\n\"}{end}' -n " + namespace + " | grep " + search
	retval := runCommand(command)
	if retval == "" {
		return nil
	}
	return strings.Split(retval, "\n")
}

func (kubectl KubectlImpl) GetIngressesV2(search string) []string {
	command := "kubectl get ingresses --all-namespaces --output=jsonpath='{range .items[*]}{\" in namespace: \"}{.metadata.namespace}{\" found: \"}{.spec.rules[0].host}{\"\\n\"}{end}' | grep " + search
	//         "kubectl get ingresses --all-namespaces --output=jsonpath='{range .items[*]}{.metadata.namespace}{"\t"}{.spec.rules[0].host}{"\n"}{end}'"
	retval := runCommand(command)
	if retval == "" {
		return nil
	}
	return strings.Split(retval, "\n")
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

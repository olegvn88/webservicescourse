package main

import (
	"encoding/json"
	"fmt"
	"github.com/openshift/hive/test/functional/utils"
)

func main() {
	getClusterOperatorStatuses()
}

func getClusterOperatorStatuses() ClusterOperator {//<-chan ClusterOperator {
	var clusterOperator ClusterOperator
	cmd, err := utils.RunShellCmd("oc get clusteroperator -o json")
	printError(err)
	err = json.Unmarshal([]byte(cmd), &clusterOperator)
	printError(err)
	for i := 0; i < len(clusterOperator.Items); i++ {
		for _, condition := range clusterOperator.Items[i].Status.ConditionsOperator {
			fmt.Println(clusterOperator.Items[i].Metadata.Name, condition.Type, condition.Status)
		}
	}
	return clusterOperator
}

type ClusterOperator struct {
	ApiVersion string  `json:"apiVersion"`
	Items      []Items `json:"items"`
}

type Items struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   MetadataOperator `json:"metadata"`
	Spec       SpecOperator     `json:"spec"`
	Status     StatusOperator   `json:"status"`
}

type SpecOperator struct {
}

type MetadataOperator struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Generation        int32  `json:"generation"`
	Name              string `json:"name"`
	ResourceVersion   string `json:"resourceVersion"`
	SelfLink          string `json:"selfLink"`
	Uid               string `json:"uid"`
}

type RelatedObjects struct {
	Group    string `json:"group"`
	Name     string `json:"name"`
	Resource string `json:"resource"`
}

type Versions struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type StatusOperator struct {
	ConditionsOperator []ConditionsOperator `json:"conditions"`
	Extension          string               `json:"extension"`
	RelatedObjects     []RelatedObjects     `json:"relatedObjects"`
	Versions           []Versions           `json:"versions"`
}

type ConditionsOperator struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

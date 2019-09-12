package main

import (
	"encoding/json"
	"fmt"
	"github.com/openshift/hive/test/functional/utils"
)

func main() {
	GetClusterState()
}

func GetClusterState() ClusterState {
	var clusterState ClusterState
	cmd, err := utils.RunShellCmd("oc get clusterstate olegtest -o json")
	printError(err)
	err = json.Unmarshal([]byte(cmd), &clusterState)
	printError(err)
	for i := 0; i < len(clusterState.Status.ClusterOperators); i++ {
		for _, condition := range clusterState.Status.ClusterOperators[i].Conditions {
			if condition.Type == "Available" || condition.Type == "Degraded" || condition.Type == "Progressing" || condition.Type == "Upgradeable" {
				fmt.Println(clusterState.Status.ClusterOperators[i].Name, condition.Type, condition.Status, condition.LastTransitionTime)
			}
		}
	}
	return clusterState
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type ClusterState struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"MetadataOperator"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

type Metadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Generation        int32             `json:"generation"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	OwnerReferences   []OwnerReferences `json:"ownerReferences"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	Uid               string            `json:"uid"`
}

type OwnerReferences struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

type Spec struct {
}

type Status struct {
	ClusterOperators []ClusterOperators `json:"clusterOperators"`
	lastUpdated      string             `json:"lastUpdated"`
}

type ClusterOperators struct {
	Conditions []Conditions `json:"conditions"`
	Name       string       `json:"name"`
}

type Conditions struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/openshift/hive/test/functional/utils"
)

func main() {
	getClusterOperatorStatuses()
}

func getClusterOperatorStatuses() ClusterOperator {
	var clusterOperator ClusterOperator
	cmd, err := utils.RunShellCmd("oc get clusteroperator -o json")
	printError(err)
	err = json.Unmarshal([]byte(cmd), &clusterOperator)
	printError(err)
	for i := 0; i < len(clusterOperator.Items); i++ {
		for _, condition := range clusterOperator.Items[i].Status.ConditionsOperator {
			//fmt.Println(clusterOperator.ItemsOperator[i].MetadataState.Name, condition.Type, condition.StatusState)
			result := GetClusterState(clusterOperator.Items[i].Metadata.Name, condition.Type, condition.Status)
			if result == false {
			fmt.Println(clusterOperator.Items[i].Metadata.Name, condition.Type, result)
			}
		}
	}
	return clusterOperator
}

func GetClusterState(operatorName, conditionType string, conditionStatus string) bool {
	var result bool
	var clusterState ClusterState
	cmd, err := utils.RunShellCmd("oc get clusterstate olegtest -o json")
	printError(err)
	err = json.Unmarshal([]byte(cmd), &clusterState)
	printError(err)
	for i := 0; i < len(clusterState.Status.ClusterOperators); i++ {
		for _, condition := range clusterState.Status.ClusterOperators[i].Conditions {
			if clusterState.Status.ClusterOperators[i].Name == operatorName &&  condition.Type == conditionType && condition.Status == conditionStatus {
				result = true
			}
		}
	}
	return result
}

type ClusterOperator struct {
	ApiVersion string          `json:"apiVersion"`
	Items      []ItemsOperator `json:"items"`
}

type ItemsOperator struct {
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

type VersionsOperator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type StatusOperator struct {
	ConditionsOperator []ConditionsOperator `json:"conditions"`
	Extension          string               `json:"extension"`
	RelatedObjects     []RelatedObjects     `json:"relatedObjects"`
	Versions           []VersionsOperator   `json:"versions"`
}

type ConditionsOperator struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}


type ClusterState struct {
	ApiVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Metadata   MetadataState `json:"metadata"`
	Spec       Spec          `json:"spec"`
	Status     StatusState   `json:"status"`
}

type MetadataState struct {
	CreationTimestamp string                 `json:"creationTimestamp"`
	Generation        int32                  `json:"generation"`
	Name              string                 `json:"name"`
	Namespace         string                 `json:"namespace"`
	OwnerReferences   []OwnerReferencesState `json:"ownerReferences"`
	ResourceVersion   string                 `json:"resourceVersion"`
	SelfLink          string                 `json:"selfLink"`
	Uid               string                 `json:"uid"`
}

type OwnerReferencesState struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

type Spec struct {
}

type StatusState struct {
	ClusterOperators []ClusterOperatorsState `json:"clusterOperators"`
	lastUpdated      string                  `json:"lastUpdated"`
}

type ClusterOperatorsState struct {
	Conditions []ConditionsState `json:"conditions"`
	Name       string            `json:"name"`
}

type ConditionsState struct {
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
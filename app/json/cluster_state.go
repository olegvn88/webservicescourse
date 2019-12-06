package main

//func GetClusterState(operatorName, conditionType string, conditionStatus string) bool {
//	var result bool
//	var clusterState ClusterState
//	cmd, err := utils.RunShellCmd("oc get clusterstate olegtest -o json")
//	printError(err)
//	err = json.Unmarshal([]byte(cmd), &clusterState)
//	printError(err)
//	for i := 0; i < len(clusterState.Status.ClusterOperators); i++ {
//		for _, condition := range clusterState.Status.ClusterOperators[i].Conditions {
//			//if condition.Type == "Available" || condition.Type == "Degraded" || condition.Type == "Progressing" || condition.Type == "Upgradeable" {
//			if clusterState.Status.ClusterOperators[i].Name == operatorName &&  condition.Type == conditionType && condition.Status == conditionStatus {
//				fmt.Println(clusterState.Status.ClusterOperators[i].Name, condition.Type, condition.Status, condition.LastTransitionTime)
//				result = true
//			}
//		}
//	}
//	return result
//}

//func printError(err error) {
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//type ClusterState struct {
//	ApiVersion string        `json:"apiVersion"`
//	Kind       string        `json:"kind"`
//	Metadata   MetadataState `json:"metadata"`
//	Spec       Spec          `json:"spec"`
//	Status     StatusState   `json:"status"`
//}
//
//type MetadataState struct {
//	CreationTimestamp string                 `json:"creationTimestamp"`
//	Generation        int32                  `json:"generation"`
//	Name              string                 `json:"name"`
//	Namespace         string                 `json:"namespace"`
//	OwnerReferences   []OwnerReferencesState `json:"ownerReferences"`
//	ResourceVersion   string                 `json:"resourceVersion"`
//	SelfLink          string                 `json:"selfLink"`
//	Uid               string                 `json:"uid"`
//}
//
//type OwnerReferencesState struct {
//	ApiVersion         string `json:"apiVersion"`
//	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
//	Controller         bool   `json:"controller"`
//	Kind               string `json:"kind"`
//	Name               string `json:"name"`
//	Uid                string `json:"uid"`
//}
//
//type Spec struct {
//}
//
//type StatusState struct {
//	ClusterOperators []ClusterOperatorsState `json:"clusterOperators"`
//	lastUpdated      string                  `json:"lastUpdated"`
//}
//
//type ClusterOperatorsState struct {
//	Conditions []ConditionsState `json:"conditions"`
//	Name       string            `json:"name"`
//}
//
//type ConditionsState struct {
//	LastTransitionTime string `json:"lastTransitionTime"`
//	Reason             string `json:"reason"`
//	Status             string `json:"status"`
//	Type               string `json:"type"`
//}

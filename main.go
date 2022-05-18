package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "strings"

	_ "k8s.io/apimachinery/pkg/labels"
)

func main() {

	type Spec struct {
		// Apiservicedefinitions apiservicedefinitions `json:"apiservicedefinitions"`
		// Cleanup cleanup `json:"cleanup"`
		// Customresourcedefinitions customresourcedefinitions `json:"customresourcedefinitions"`
		// Description string `json:"description"`
		// DisplayName string `json:"displayName"`
		// Icon icon `json:"icon"`
		// Install install `json:"install"`
		// InstallModes []installModes `json:"installModes"`
		// Keywords []keywords `json:"keywords"`
		// Links []links `json:"links"`
		// Maintainers []maintainers `json:"maintainers"`
		// Maturity string `json:"maturity`
		// MinkubeVersion string `json:"minikubeVersion"`
		// Provider provider `json:"provider"`
		// RelatedImages []relatedImages `json:"relatedImages"`
		// Version string `json:"version"`
	}

	type Status struct {
		// Cleanup cleanup `json:"cleanup"`
		// Conditions []conditions `json:"conditions"`
		// LastTransitionTime string `json:"lastTransitionTime"`
		// LastUpdateTime string `json:"lastUpdateTime"`
		// Message string `json:"message"`
		// Phase string `json:"phase"`
		// Reason string `json:"reason"`
		// RequirementStatus []requirementStatus `json:"requirementStatus"`
	}

	type Annotations struct {
		Examples          string `json:"alm-examples"` //alm
		Capabilities      string `json:"capabilities"`
		Categories        string `json:"categories"`
		Certified         string `json:"certified"`
		ContainerImage    string `json:"containerImage"`
		CreatedAt         string `json:"createdAt"`
		Description       string `json:"description"`
		OperatorGroup     string `json:"olm.operatorGroup"`
		OperatorNamespace string `json:"olm.operatorNamespace"`
		SkipRange         string `json:"olm.skipRange"`
		// peratorframework.io/cluster-monitoring string `json:"operatorframework.io/cluster-monitoring"`
		// operatorframework.io/properties string  `json:"operatorframework.io/properties"`
		// operatorframework.io/suggested-namespace string `json:"operatorframework.io/suggested-namespace"`
		// operators.openshift.io/infrastructure-features string `json:"operators.openshift.io/infrastructure-features"`
		// operators.operatorframework.io/builder string `json:"operators.operatorframework.io/builder"`
		// operators.operatorframework.io/project_layout string `json:"operators.operatorframework.io/project_layout"`
		Support string `json:"support"`
	}

	type Metadatas struct {
	}

	type Metadata struct {
		Annotations Annotations `json:"annotations"`
		// CreationTimestamp string `json:"metadata"`
		// Generation string `json:"generation"`
		// Labels labels `json:"labels"`
		// ManagedFields []managedfiels `json:"managedFields"`
		// Name string `"json:"name"`
		// Namespace string `"json:"namespace"`
		// ResourceVersion string `json:"esourceVersion"`
		// Uid string `json:"uid"`
	}

	type Item struct {
		ApiVersion string   `json:"apiVersion"`
		Kind       string   `json:"kind"`
		Metadata   Metadata `json:"metadata"`
		Spec       Spec     `json:"spec"`
		Status     Status   `json:"status"`
	}

	type Sample struct {
		ApiVersion string    `json:"apiVersion"`
		Items      []Item    `json:"Items"`
		Kind       string    `json:"kind"`
		Metadata   Metadatas `json:"metadata"`
	}

	data, err := ioutil.ReadFile("csv.json")
	if err != nil {

		fmt.Println(err)
	}

	// var results map[string]interface{}pwd

	// json.Unmarshal([]byte(data), &results)
	var sample Sample

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(data, &sample)
	// for _, result := range results {

	// 	// fmt.Println("Reading Value for Key :", key)
	// 	//Reading each value by its key
	// 	fmt.Println(
	// 		"- Name :", result["name"],
	// 		"- Description :", result["description"])
	// }
	//fmt.Println(sample.Items.)

	for i := 0; i < len(sample.Items); i++ {
		fmt.Println(sample.Items[i].Metadata.Annotations.Support)
	}

}

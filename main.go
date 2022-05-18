package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "k8s.io/kubernetes/pkg/api"
	// // _ "k8s.io/kubernetes/pkg/api/install"
	// // _ "k8s.io/kubernetes/pkg/apis/extensions/install"
	// // "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
)

func main() {

	type Annotations struct {
		Deployment_kubernetes_io_revision string `json:"deployment.kubernetes.io/revision"`
	}

	type FieldsV1 struct {
	}
	type ManagedFields struct {
		Manager    string   `json:"manager"`
		Operation  string   `json:"operation"`
		ApiVersion string   `json:"apiVersion"`
		Time       string   `json:"time"`
		FieldsType string   `json:"fieldsType"`
		FieldsV1   FieldsV1 `json:"FieldsV1"`
	}

	type MetadataItem struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		Uid             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Generation      string `json:"generation"`

		CreationTimestamp string          `json:"creationTimestamp"`
		Annotations       Annotations     `json:"annotations"`
		ManagedFields     []ManagedFields `json:"managedFields"`
	}

	type Conditions struct {
	}

	type Status struct {
		ObservedGeneration string       `json:"observedGeneration"`
		Replicas           string       `json:"replicas"`
		UpdatedReplicas    string       `json:"updatedReplicas"`
		ReadyReplicas      string       `json:"readyReplicas"`
		AvailableReplicas  string       `json:"availableReplicas"`
		Conditions         []Conditions `json:"conditions"`
	}
	type MatchLabels struct {
		App string `json:"app"`
	}

	type Selector struct {
		MatchLabels MatchLabels `json:"matchLabels"`
	}
	type Labels struct {
		App string `json:"app"`
	}
	type TemplateMetadata struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Labels            Labels `json:"labels"`
	}
	type Ports struct {
	}

	type Containers struct {
		Name                     string  `json:"name"`
		Image                    string  `json:"image"`
		Ports                    []Ports `json:"ports"`
		Resources                string  `json:"resources"`
		TerminationMessagePath   string  `json:"terminationMessagePath"`
		TerminationMessagePolicy string  `json:"terminationMessagePolicy"`
		ImagePullPolicy          string  `json:"imagePullPolicy"`
	}
	type SecurityContext struct {
		//empty
	}

	type TSpec struct {
		Containers                    []Containers    `json:"containers"`
		RestartPolicy                 string          `json:"restartPolicy"`
		TerminationGracePeriodSeconds string          `json:"terminationGracePeriodSeconds"`
		DnsPolicy                     string          `json:"dnsPolicy"`
		SecurityContext               SecurityContext `json:"securityContext"`
		SchedulerName                 string          `json:"schedulerName"`
	}
	type Template struct {
		TemplateMetadata TemplateMetadata `json:"metadata"`
		TSpec            TSpec            `json:"Spec"`
	}
	type Strategy struct {
	}
	type Spec struct {
		Replicas                string   `json:"replicas"`
		Selector                Selector `json:"selector"`
		Template                Template `json:"template"`
		Strategy                Strategy `json:"strategy"`
		RevisionHistoryLimit    string   `json:"revisionHistoryLimit"`
		ProgressDeadlineSeconds string   `json:"progressDeadlineSeconds"`
	}

	type Item struct {
		Metadata MetadataItem `json:"metadata"`
		Spec     Spec         `json:"spec"`
		Status   Status       `json:"status"`
	}

	type Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	}

	type Deplo struct {
		ApiVersion string   `json:"apiVersion"`
		Items      []Item   `json:"Items"`
		Kind       string   `json:"kind"`
		Metadata   Metadata `json:"metadata"`
	}

	url := "http://localhost:8001/apis/apps/v1/namespaces/default/deployments"

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//fmt.Println(string(body))

	//data=string(body)
	// var results map[string]interface{}
	// json.Unmarshal([]byte(data), &results)
	var depl Deplo

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(body, &depl)
	// for _, result := range results {

	// 	// fmt.Println("Reading Value for Key :", key)
	// 	//Reading each value by its key
	// 	fmt.Println(
	// 		"- Name :", result["name"],
	// 		"- Description :", result["description"])
	// }
	//fmt.Println(sample.Items.)
	fmt.Println(len(depl.Items))

	for i := 0; i < len(depl.Items); i++ {
		fmt.Println(depl.Items[i].Spec.Template.TSpec.Containers[0].Image) //image name
		fmt.Println(depl.Items[i].Metadata.Name)
	}

}

//fmt.Println(depl.Items[0].Spec.Template.TSpec.Containers[i].Image)

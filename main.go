package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func GetPod() {
	const Podurl = "http://localhost:8001/api/v1/namespaces/default/pods"
	response, err := http.Get(Podurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func GetDeployment() {
	const Depurl = "http://localhost:8001/apis/apps/v1/namespaces/default/deployments"
	response, err := http.Get(Depurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code:", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	var depl Deplo

	json.Unmarshal(content, &depl)

	fmt.Println(len(depl.Items))

	for i := 0; i < len(depl.Items); i++ {
		fmt.Println(depl.Items[i].Spec.Template.TSpec.Containers[0].Image) //image naem
		fmt.Println(depl.Items[i].Metadata.Name)
	}

}
func GetService() {
	const Svcurl = "http://localhost:8001/api/v1/namespaces/default/endpoints"
	response, err := http.Get(Svcurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code:", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func GetConfigMap() {
	const Cmcurl = "http://localhost:8001/api/v1/configmaps"
	response, err := http.Get(Cmcurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code:", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PostDeployment() {

	const Depurl = "http://localhost:8001/apis/apps/v1/namespaces/default/deployments"
	// response, err := http.Post()
	// if err != nil {
	// 	panic(err)
	// }
	requestBody := strings.NewReader(`
	{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
		   "name": "simple1-deployment",
		   "labels": {
			  "app": "simple"
		   }
		},
		"spec": {
		   "replicas": 3,
		   "selector": {
			  "matchLabels": {
				 "app": "simple"
			  }
		   },
		   "template": {
			  "metadata": {
				 "labels": {
					"app": "simple"
				 }
			  },
			  "spec": {
				 "containers": [
					{
					   "name": "simpleweb",
					   "image": "sakshim21/simple-app:latest",
					   "ports": [
						  {
							 "containerPort": 8000
						  }
					   ]
					}
				 ]
			  }
		   }
		}
	 }
	
	
	`)

	response, err := http.Post(Depurl, "application/json", requestBody)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()
	c, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(c))

}

func PostPods() {

	const Podurl = "http://localhost:8001/api/v1/namespaces/default/pods"

	// response, err := http.Post()
	// if err != nil {
	// 	panic(err)
	// }
	requestBody := strings.NewReader(`
	{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
		   "name": "nginxnew",
		   "labels": {
			  "name": "nginx"
		   }
		},
		"spec": {
		   "containers": [
			  {
				 "name": "nginx",
				 "image": "nginx",
				 "ports": [
					{
					   "containerPort": 80
					}
				 ]
			  }
		   ]
		}
	 }
	
	
	
	`)

	response, err := http.Post(Podurl, "application/json", requestBody)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()
	c, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(c))

}

func PostService() {

	const Svcurl = "http://localhost:8001/api/v1/namespaces/default/endpoints"

	// response, err := http.Post()
	// if err != nil {
	// 	panic(err)
	// }
	requestBody := strings.NewReader(`
	
		{
			"apiVersion": "v1",
			"kind": "Service",
			"metadata": {
			   "name": "simple-service1"
			},
			"spec": {
			   "selector": {
				  "app": "nginx"
			   },
			   "type": "LoadBalancer",
			   "ports": [
				  {
					 "protocol": "TCP",
					 "port": 5000,
					 "targetPort": 8000,
					 "nodePort": 31110
				  }
			   ]
			}
		 }
	
	
	
	`)

	response, err := http.Post(Svcurl, "application/json", requestBody)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()
	c, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(c))

}

func PostConfiMap() {

	const Cmcurl = "http://localhost:8001/api/v1/configmaps"

	// response, err := http.Post()
	// if err != nil {
	// 	panic(err)
	// }
	requestBody := strings.NewReader(`
	
	{
		"kind": "ConfigMap",
		"apiVersion": "v1",
		"metadata": {
		   "name": "example-configmap"
		},
		"data": {
		   "database": "mongodb",
		   "database_uri": "mongodb://localhost:27017",
		   "keys": "image.public.key=771 \nrsa.public.key=42\n"
		}
	 }
	
	
	
	`)

	response, err := http.Post(Cmcurl, "application/json", requestBody)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()
	c, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(c))

}

func main() {
	//http.HandleFunc("/api",GetDeployment)
	fmt.Println("hello")
	//PostDeployment()
	GetDeployment()
	//PostPods()
	GetPod()
	//GetService()
	//PostService()
	//GetConfigMap()
	//PostConfiMap()

}

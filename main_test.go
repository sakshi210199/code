package main

import (
	"fmt"
	"testing"

	//client "sak/client"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
	//core "k8s.io/api/core/v1"
)

func TestCreateObject(t *testing.T) {
	var client Client
	client.Clientset = testclient.NewSimpleClientset()

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod5",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nginx",
					Image:           "nginx",
					ImagePullPolicy: "Always",
				},
			},
		},
	}

	_, err := client.CreatePod(pod)
	if err != nil {
		fmt.Print(err.Error())
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment2",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: nil,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// }

	_, err = client.CreateDeployment(deployment)
	if err != nil {
		fmt.Print(err.Error())
	}

	service := &v1.Service{

		ObjectMeta: metav1.ObjectMeta{
			Name:      "myservice",
			Namespace: "default",
			Labels: map[string]string{
				"app": "app",
			},
		},
		Spec: v1.ServiceSpec{
			Ports:     nil,
			Selector:  nil,
			ClusterIP: "",
		},

		//fmt.Println(x)
		//client.Getthepods(client.Clientset)
		//fmt.Println(deployment)

	}

	_, err = client.CreateService(service)
	if err != nil {
		fmt.Print(err.Error())
	}

}

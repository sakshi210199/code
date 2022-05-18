package main

import (
//"fmt"

//	client "github.com/kubernetes-sdk-for-go-101/pkg/client"
//v1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//"k8s.io/client-go/tools/clientcmd"
//"k8s.io/klog/v2"
)

func main() {

	// kubeconfig := "<path to kubeconfig file>"

	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// client := Client{
	// 	Clientset: clientset,
	// }

	// pod := &v1.Pod{
	// 	TypeMeta: metav1.TypeMeta{
	// 		Kind:       "Pod",
	// 		APIVersion: "v1",
	// 	},
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      "test-pod",
	// 		Namespace: "default",
	// 	},
	// 	Spec: v1.PodSpec{
	// 		Containers: []v1.Container{
	// 			{
	// 				Name:            "nginx",
	// 				Image:           "nginx",
	// 				ImagePullPolicy: "Always",
	// 			},
	// 		},
	// 	},
	// }

	// pod, err = client.CreatePod(pod)
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	// klog.Infof("Pod %s has been successfully created", pod.Name)

}

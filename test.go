package main

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type Client struct {
	Clientset kubernetes.Interface //clientset contain client of groups.//it mock all function that regular client
}

func (c Client) CreatePod(pod *v1.Pod) (*v1.Pod, error) {
	pod, err := c.Clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Error occured while creating pod %s: %s", pod.Name, err.Error())
		return nil, err
	}

	klog.Infof("Pod %s is succesfully created", pod.Name)
	return pod, nil
}

// func (c Client) CreateDeployment(pod *v1.Pod) (*v1.Pod, error) {
// 	pod, err := c.Clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
// 	if err != nil {
// 		klog.Errorf("Error occured while creating pod %s: %s", pod.Name, err.Error())
// 		return nil, err
// 	}

// 	klog.Infof("Pod %s is succesfully created", pod.Name)
// 	return pod, nil
// }

func (c Client) CreateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error) {

	deployment, err := c.Clientset.AppsV1().Deployments(deployment.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Error occured while creating pod %s: %s", deployment.Name, err.Error())
		return nil, err
	}

	klog.Infof("Deployment %s is succesfully created", deployment.Name)
	return deployment, nil
}
func Int32ptr(i int32) *int32 {
	return &i
}

// func Getthedeplyment(clientset *kubernetes.Clientset, err error) {
// 	fmt.Println("")
// 	fmt.Println("")
// 	fmt.Println("Deployments are")
// 	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
// 	if err != nil {
// 		fmt.Printf("listing deployments %s\n", err.Error())
// 	}
// 	for _, d := range deployments.Items {

// 		fmt.Printf("%s", d.Name)
// 		//fmt.Printf("%+v", d)
// 		fmt.Println("")
// 	}
// }

func (c Client) Getthepods(pod *v1.Pod) {
	//fmt.Println("Pods from default namespace")

	pods, err := c.Clientset.CoreV1().Pods(pod.Namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error occured while listing pod %s: %s", pod.Name, err.Error())
		//return nil, err
	}
	for _, pod := range pods.Items {
		fmt.Println("")
		fmt.Printf("%s", pod.Name)
	}
}

func (c Client) CreateService(service *v1.Service) (*v1.Service, error) {
	service, err := c.Clientset.CoreV1().Services(service.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Error occured while creating pod %s: %s", service.Name, err.Error())
		return nil, err
	}

	klog.Infof("Pod %s is succesfully created", service.Name)
	return service, nil
}

package main

import (
	"context"
	"flag"
	"fmt"
	_ "io/ioutil"
	_ "log"
	"os"

	_ "net/http"

	"path/filepath"
	"sync"

	//corev1 "k8s.io/api/core/v1"
	//core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//"k8s.io/api/apps/v1"
	//appsv1 "k8s.io/api/apps/v1"
	// "k8s.io/client-go/pkg/apis"
	// _ "k8s.io/client-go/pkg/api/install"
	// _ "k8s.io/client-go/pkg/apis/extensions/install"
)

var wg sync.WaitGroup

var mu sync.Mutex

//f, err := os.Create("output.txt")

func main() {
	//var Int32Ptr := Int32
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filename := d + "/log1.txt"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Println(err)
	}
	//check(err)
	//defer f.Close()
	// fmt.Println(file)
	//ubeconfig := flag.String("kubeconfig", "~.kube/config", "location to your kubeconfig file")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to config file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "(optional) absolute path to config file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig) //clientcmd module is used to build the clientSet by passing the local kubeconfig file
	if err != nil {
		// handle error
		fmt.Printf("erorr %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}
	//ctx := context.Background()
	//pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	wg.Add(5)
	go Getthepods(clientset, err, &wg, f)
	go Getthedeplyment(clientset, err, &wg, f)
	go Gettheservice(clientset, err, &wg, f)
	go Getthereplicaset(clientset, err, &wg, f)
	go Gettheconfigmap(clientset, err, &wg, f)

	// 	Create the config from the kubeconfig file
	// Create the clientset from the config
	// Create the resource(s) (configMap) using the appropriate API structure (core/v1 and meta/v1)
	// Identify if the resource exists using Get() and the right resource method. If it exist, use Create(), otherwise use Update()

	////////////////////////////////////////new deplyment
	// deployment := getdeployment()
	// deploymentsClient := clientset.AppsV1().Deployments(core.NamespaceDefault)
	// fmt.Println("Creating deployment...")
	// result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	//////////////////////////////service

	// servicenew := createserviceObject()
	// //serviceClient :=clientset.CoreV1().Services(core.NamespaceDefault).Create(context.TODO(),servicenew, metav1.CreateOptions{})
	// newservice, err := clientset.CoreV1().Services(servicenew.Namespace).Create(context.TODO(), servicenew, metav1.CreateOptions{}) //clientset.CoreV1().Pods(podcreate.Namespace).Create(podcreate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pod created successfully...")
	// fmt.Printf("%s", newservice.Name)

	//////////////////////////////////////create node pod
	//build the pod defination we want to deploy
	// podcreate := getPodObject()
	// fmt.Printf("%s", podcreate)
	// fmt.Printf("%s", podcreate.Name)

	// newpod, err := clientset.CoreV1().Pods(podcreate.Namespace).Create(context.TODO(), podcreate, metav1.CreateOptions{}) //clientset.CoreV1().Pods(podcreate.Namespace).Create(podcreate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pod created successfully...")
	// fmt.Printf("%s", newpod.Name)
	//deployment:=getdeployment()
	// deploymentsClient := clientset.AppsV1().Deployments(core.NamespaceDefault)

	// fmt.Println("Creating deployment...")
	// result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	// fmt.Println("rest api")

	// fmt.Println("hellllllllllllllllllllllllllllllllo")

	// url := "http://localhost:8001/api/v1/namespaces/default/configmaps/config-example-values"

	// resp, getErr := http.Get(url)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }
	// body, readErr := ioutil.ReadAll(resp.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }
	// fmt.Println(string(body))

	// decode := scheme.Codecs.UniversalDeserializer().Decode
	// //decode:=scheme.Codecs.UniversalDecoder().Decode()

	// obj, _, err := decode([]byte(string(body)), nil, nil)
	// //f//mt.Println(decode)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	// }
	// //cm:= obj.(*v1beta1.Deployment)
	// //deployment := obj.(*v1beta1.Deployment)
	// fmt.Println("helllllllllllllllllllllllllllllaalllo")
	// fmt.Printf("%#v\n", obj)

	// configmap := obj.(*v1beta1.Deployment)

	// fmt.Printf("%#v\n", deployment)
	wg.Wait()

}

func Gettheconfigmap(clientset *kubernetes.Clientset, err error, w *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	x := "configmap"
	//mu.Lock()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("configmap")
	cm, err := clientset.CoreV1().ConfigMaps("default").List(context.Background(), metav1.ListOptions{})
	//config,err := clientset.CoreV1().ConfigMaps("default").List(context.Background(),metav1.ListOptions{})
	if err != nil {
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	for _, c := range cm.Items {

		fmt.Printf("%s", c.Name)
		//fmt.Println("%s",c.Name)
		WriteToFile(c.Name, f, w, x)
		//fmt.Fprintf(w, "%v\n", i)
		//fmt.FPrintf(w, "%v\n", c.Name)
		fmt.Println("")
	}
	//mu.Unlock()

}

func Getthereplicaset(clientset *kubernetes.Clientset, err error, w *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	x := "replicaset"
	//mu.Lock()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("relicaset are")

	replicaset, err := clientset.AppsV1().ReplicaSets("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	for _, r := range replicaset.Items {

		fmt.Printf("%s", r.Name)
		WriteToFile(r.Name, f, w, x)
		fmt.Println("")
	}
	//mu.Unlock()

}

func Gettheservice(clientset *kubernetes.Clientset, err error, w *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	//mu.Lock()
	x := "service"
	fmt.Println("")
	fmt.Println("")
	fmt.Println("service are")
	services, err := clientset.CoreV1().Services("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	for _, s := range services.Items {

		fmt.Printf("%s", s.Name)
		WriteToFile(s.Name, f, w, x)
		fmt.Println("")
	}
	//mu.Unlock()
}

func Getthedeplyment(clientset *kubernetes.Clientset, err error, w *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	//mu.Lock()
	x := "deployment"
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Deployments are")
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	for _, d := range deployments.Items {

		fmt.Printf("%s", d.Name)
		//fmt.Printf("%+v", d)
		WriteToFile(d.Name, f, w, x)
		fmt.Println("")
	}
	//mu.Unlock()
}

func Getthepods(clientset *kubernetes.Clientset, err error, w *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	x := "pod"
	//mu.Lock()
	fmt.Println("Pods from default namespace")

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// handle error
		fmt.Printf("error %s, while listing all the pods from default namespace\n", err.Error())
		fmt.Println("inside")
	}
	for _, pod := range pods.Items {
		fmt.Println("")
		fmt.Printf("%s", pod.Name)
		WriteToFile(pod.Name, f, w, x)
	}
	//mu.Unlock()
}

// func getPodObject() *core.Pod {
// 	return &core.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "my-test-pod-new",
// 			Namespace: "default",
// 			Labels: map[string]string{
// 				"app": "demo",
// 			},
// 		},
// 		Spec: core.PodSpec{
// 			Containers: []core.Container{
// 				{
// 					Name:            "busybox",
// 					Image:           "busybox",
// 					ImagePullPolicy: core.PullIfNotPresent,
// 					Command: []string{
// 						"sleep",
// 						"3600",
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// /////////////////////////////////////////////////////////////////////////////////////////////

// func getdeployment() *appsv1.Deployment {

// 	return &appsv1.Deployment{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "demo-deployment",
// 		},
// 		Spec: appsv1.DeploymentSpec{
// 			Replicas: int32ptr(2),
// 			Selector: &metav1.LabelSelector{
// 				MatchLabels: map[string]string{
// 					"app": "demo",
// 				},
// 			},
// 			Template: core.PodTemplateSpec{
// 				ObjectMeta: metav1.ObjectMeta{
// 					Labels: map[string]string{
// 						"app": "demo",
// 					},
// 				},
// 				Spec: core.PodSpec{
// 					Containers: []core.Container{
// 						{
// 							Name:  "web",
// 							Image: "nginx:1.12",
// 							Ports: []core.ContainerPort{
// 								{
// 									Name:          "http",
// 									Protocol:      core.ProtocolTCP,
// 									ContainerPort: 80,
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// }

// func int32ptr(i int32) *int32 {
// 	return &i
// }

// //////////////////////////////service
// func createserviceObject() *core.Service {
// 	return &core.Service{

// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "myservice",
// 			Namespace: "default",
// 			Labels: map[string]string{
// 				"app": "app",
// 			},
// 		},
// 		Spec: core.ServiceSpec{
// 			Ports:     nil,
// 			Selector:  nil,
// 			ClusterIP: "",
// 		},
// 	}

// }

// // spec:
// //   selector:
// //     app: simple
// //   type: LoadBalancer
// //   ports:
// //   - protocol: TCP
// //     port: 5000
// //     targetPort : 8000
// //     nodePort: 31110

// // Create Service
func WriteToFile(i string, f *os.File, w *sync.WaitGroup, x string) {
	if x == "pod" {

		mu.Lock()
		//write to the file
		fmt.Fprintln(f, "pods")
		fmt.Fprintf(f, "Printing out: %s\n", i)
		//write to stdout
		fmt.Printf("Printing out: %s\n", i)
		mu.Unlock()

	} else if x == "service" {
		mu.Lock()
		//write to the file
		fmt.Fprintln(f, "service")
		fmt.Fprintf(f, "Printing out: %s\n", i)
		//write to stdout
		fmt.Printf("Printing out: %s\n", i)
		mu.Unlock()

	} else if x == "deployment" {
		mu.Lock()
		//write to the file
		fmt.Fprintln(f, "deployment")
		fmt.Fprintf(f, "Printing out: %s\n", i)
		//write to stdout
		fmt.Printf("Printing out: %s\n", i)
		mu.Unlock()

	} else if x == "replicaset" {
		mu.Lock()
		//write to the file
		fmt.Fprintln(f, "replicaset")
		fmt.Fprintf(f, "Printing out: %s\n", i)
		//write to stdout
		fmt.Printf("Printing out: %s\n", i)
		mu.Unlock()

	} else if x == "configmap" {

		mu.Lock()
		//write to the file
		fmt.Fprintln(f, "configmap")
		fmt.Fprintf(f, "Printing out: %s\n", i)
		//write to stdout
		fmt.Printf("Printing out: %s\n", i)
		mu.Unlock()
	}

}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetOutClusterConfig(kubeconfig_abs_path *string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig_abs_path)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func GetInClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}

func GetDeployments(config *rest.Config) []v1.Deployment {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return list.Items
	// for _, item := range list.Items {
	// 	fmt.Printf(" * %s (%d replicas)\n", item.GetName(), *item.Spec.Replicas)
	// }
}

//-------------------
// dynamimc client

var dummySitesResource = schema.GroupVersionResource{
	Group:    "stable.anth",
	Version:  "v1",
	Resource: "dummysites",
}

func NewDynamicClient(config *rest.Config) (dynamic.Interface, error) {
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return dynClient, nil
}

func ListDummySites(
	ctx context.Context,
	client dynamic.Interface,
	namespace string,
) ([]unstructured.Unstructured, error) {
	// GET /apis/mongodbcommunity.mongodb.com/v1/namespaces/{namespace}/mongodbcommunity/
	list, err := client.
		Resource(dummySitesResource).
		Namespace(namespace).
		List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

func WatchDummySites(
	client dynamic.Interface,
) {
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(client, 0, metav1.NamespaceAll, nil)
	informer := informerFactory.ForResource(dummySitesResource).Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			name, url, ns := getDummySiteDataFromUnstructuredObj(obj)
			log.Printf("dummysite found: \nname:%s url:%s ns:%s \n", url, name, ns)

			dummySiteDeployment, err := CreateDummySiteDeployment(
				client,
				"k3d-myregistry.localhost:12345/dummysite",
				name+"-deployment",
				"default",
				"8080",
				1,
			)

			if err != nil {
				log.Printf("Failed to create deployment: \n%v \n", err)
				return
			}

			log.Printf("Created deployment: %s \n", dummySiteDeployment.GetName())
			// dummySiteSvc, err := CreateDummySiteDeploymentService(
			// 	client,
			// 	name+"-svc",
			// 	name+"-deployment",
			// 	"default",
			// 	"NodePort",
			// 	30080,
			// 	1234,
			// 	8080,
			// )

			// if err != nil {
			// 	log.Printf("Failed to create service: \n%v \n", err)
			// 	return
			// }

			// log.Printf("Created Service: %s \n", dummySiteSvc.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			old_url, old_name, old_ns := getDummySiteDataFromUnstructuredObj(oldObj)
			fmt.Printf("old_url: %s old_name:%s old_ns:%s \n", old_url, old_name, old_ns)

			new_url, new_name, new_ns := getDummySiteDataFromUnstructuredObj(newObj)
			fmt.Printf("new_url: %s new_name:%s new_ns:%s \n", new_url, new_name, new_ns)
		},
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	informer.Run(ctx.Done())

	fmt.Print("done")
}

func getDummySiteDataFromUnstructuredObj(obj interface{}) (string, string, string) {
	// we should cast this entyre object. Too lazy to write an struct for this
	// AddFunc: func(obj unstructured.Unstructured) {
	// converting the dynamic object to your CRD struct
	// typedObj := obj.(*unstructured.Unstructured)
	// bytes, _ := typedObj.MarshalJSON()

	// var crdObj *crd.CRD
	// json.Unmarshal(bytes, &crdObj)
	uobj := obj.(*unstructured.Unstructured)

	i_spec := uobj.Object["spec"]
	map_spec := i_spec.(map[string]interface{})
	url := map_spec["url"].(string)

	i_metadata := uobj.Object["metadata"]
	map_metadata := i_metadata.(map[string]interface{})
	name := map_metadata["name"].(string)
	ns := map_metadata["namespace"].(string)

	return name, url, ns
}

func CreateDummySiteDeployment(
	client dynamic.Interface,
	image, name, namespace, port string,
	replicas int32,
) (*unstructured.Unstructured, error) {
	dummySiteDeployment := DummySiteDeployment(image, name, namespace, port, replicas)
	createdDeployment, err := client.Resource(*Deployment).Namespace(namespace).
		Create(context.TODO(), dummySiteDeployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return createdDeployment, nil
}

func CreateDummySiteDeploymentService(
	client dynamic.Interface,
	name, selector, namespace, svcType string,
	nodePort, port, targetPort int32,
) (*unstructured.Unstructured, error) {
	dummySiteDeploymentSvc := DummySiteDeploymentService(name, selector, svcType, nodePort, port, targetPort)
	createdService, err := client.Resource(Service).Namespace(namespace).
		Create(context.TODO(), dummySiteDeploymentSvc, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return createdService, nil
}

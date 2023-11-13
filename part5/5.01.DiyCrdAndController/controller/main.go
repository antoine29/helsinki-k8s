package main

import (
	"log"
)

func main() {
	// k8sConfigAbsPath := "/home/antoine/.kube/config"
	// k8sConfig := GetOutClusterConfig(&k8sConfigAbsPath)

	k8sConfig := GetInClusterConfig()

	client, err := NewDynamicClient(k8sConfig)
	if err != nil {
		log.Println("error creating dynamic cient")
		log.Println(err.Error())
	}

	WatchDummySites(client)
}

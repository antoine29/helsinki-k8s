package main

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var Deployment = &schema.GroupVersionResource{
	Group:    "apps",
	Version:  "v1",
	Resource: "deployments",
}

func DummySiteDeployment(
	image, name, namespace, port, url string,
	replicas int32,
) *unstructured.Unstructured {
	deployment := unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"replicas": replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": name,
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": name,
						},
					},
					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  name,
								"image": image,
								"env": []map[string]interface{}{
									{
										"name":  "PORT",
										"value": port,
									},
									{
										"name":  "URL",
										"value": url
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return &deployment
}

var Service = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "services",
}

func DummySiteDeploymentService(
	name, appSelector, svcType string,
	nodePort, port, targetPort int32,
) *unstructured.Unstructured {
	service := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"type": svcType,
				"selector": map[string]interface{}{
					"app": appSelector,
				},
				"ports": []map[string]interface{}{
					{
						"name":       "http",
						"protocol":   "TCP",
						"targetPort": targetPort,
						"nodePort":   nodePort,
						"port":       port,
					},
				},
			},
		},
	}

	return service
}

package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	r "com.justin.k8s.api/deploy/k8sapi/resources"
)

const (
	NAMESPACE = "blog-system"
)

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	r.CreateNamespace(clientset, NAMESPACE)
	// r.DeleteNamespace(clientset, NAMESPACE)

	userDeployment := &r.DeploymentResource{
		Namespace:   NAMESPACE,
		Name:        "srv-user-deployment",
		Replicas:    2,
		Matchlabels: map[string]string{"app": "srv-user"},
		Labels:      map[string]string{"app": "srv-user"},
		Containers: []r.Container{{
			Name:  "srv-user",
			Image: "srv-user",
			Ports: []r.Port{{
				Name: "http",
				Port: 8888,
			}},
		}},
	}
	r.CreateDeployment(clientset, userDeployment)
	// r.DeleteDeployment(clientset, userDeployment.Namespace, userDeployment.Name)

	userService := &r.ServiceResource{
		Namespace: NAMESPACE,
		Name:      "srv-user-service",
		Ports: []r.Port{{
			Name:     "http",
			Port:     8888,
			NodePort: 30088,
		}},
		Selector: map[string]string{"app": "srv-user"},
	}
	r.CreateService(clientset, userService)
	// r.DeleteService(clientset, userService.Namespace, userService.Name)
}

package main

import (
	"flag"
	"io/ioutil"
	"os"
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

	/**
	 * 创建 Namespace
	 */
	r.CreateNamespace(clientset, NAMESPACE)
	// r.DeleteNamespace(clientset, NAMESPACE)

	/**
	 * 创建 ConfigMap
	 */
	workDir, _ := os.Getwd()
	userConfig, _ := ioutil.ReadFile(workDir + "/config/srv-user.yaml")
	articleConfig, _ := ioutil.ReadFile(workDir + "/config/srv-article.yaml")

	userConfigmap := &r.ConfigmapResource{
		Namespace: NAMESPACE,
		Name:      "srv-user-configmap",
		BinaryData: map[string][]byte{
			"application.yaml": userConfig,
		},
	}
	articleConfigmap := &r.ConfigmapResource{
		Namespace: NAMESPACE,
		Name:      "srv-article-configmap",
		BinaryData: map[string][]byte{
			"application.yaml": articleConfig,
		},
	}

	r.CreateConfigmap(clientset, userConfigmap)
	r.CreateConfigmap(clientset, articleConfigmap)

	// r.UpdateConfigmap(clientset, userConfigmap)
	// r.UpdateConfigmap(clientset, articleConfigmap)

	// r.DeleteConfigmap(clientset, userConfigmap.Namespace, userConfigmap.Name)
	// r.DeleteConfigmap(clientset, articleConfigmap.Namespace, articleConfigmap.Name)

	/**
	 * 创建 Deployment
	 */
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
		VolumesName:   "config",
		ConfigmapName: userConfigmap.Name,
		MountPath:     "/etc/config",
	}
	articleDeployment := &r.DeploymentResource{
		Namespace:   NAMESPACE,
		Name:        "srv-article-deployment",
		Replicas:    2,
		Matchlabels: map[string]string{"app": "srv-article"},
		Labels:      map[string]string{"app": "srv-article"},
		Containers: []r.Container{{
			Name:  "srv-article",
			Image: "srv-article",
			Ports: []r.Port{{
				Name: "http",
				Port: 8899,
			}},
		}},
		VolumesName:   "config",
		ConfigmapName: articleConfigmap.Name,
		MountPath:     "/etc/config",
	}

	r.CreateDeployment(clientset, userDeployment)
	r.CreateDeployment(clientset, articleDeployment)

	// r.DeleteDeployment(clientset, userDeployment.Namespace, userDeployment.Name)
	// r.DeleteDeployment(clientset, articleConfigmap.Namespace, articleConfigmap.Name)

	/**
	 * 创建 Service
	 */
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
	articleService := &r.ServiceResource{
		Namespace: NAMESPACE,
		Name:      "srv-article-service",
		Ports: []r.Port{{
			Name:     "http",
			Port:     8899,
			NodePort: 30099,
		}},
		Selector: map[string]string{"app": "srv-article"},
	}

	r.CreateService(clientset, userService)
	r.CreateService(clientset, articleService)

	// r.DeleteService(clientset, userService.Namespace, userService.Name)
	// r.DeleteService(clientset, articleConfigmap.Namespace, articleConfigmap.Name)
}

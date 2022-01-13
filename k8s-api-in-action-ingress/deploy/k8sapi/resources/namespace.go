package resources

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNamespace(clientset *kubernetes.Clientset, name string) {
	namespacesClient := clientset.CoreV1().Namespaces()

	namespace, _ := namespacesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if namespace.Name == name {
		fmt.Printf("Namespace: %s is already exist.\n", name)
		return
	} else {
		namespace = &apiv1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}

		result, err := namespacesClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create namespace: %s success\n", result.GetName())
	}
}

func DeleteNamespace(clientset *kubernetes.Clientset, name string) {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

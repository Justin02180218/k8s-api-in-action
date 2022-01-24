package resources

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateConfigmap(clientset *kubernetes.Clientset, cmr *ConfigmapResource) {
	configmapsClient := clientset.CoreV1().ConfigMaps(cmr.Namespace)

	configmap, _ := configmapsClient.Get(context.TODO(), cmr.Name, metav1.GetOptions{})
	if configmap.Name == cmr.Name {
		fmt.Printf("Configmap: %s is already exist.\n", cmr.Name)
		return
	}

	configmap = &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: cmr.Name,
		},
		// Data: cmr.Data,
		BinaryData: cmr.BinaryData,
	}

	result, err := configmapsClient.Create(context.TODO(), configmap, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create configmap: %s success\n", result.GetName())
}

func UpdateConfigmap(clientset *kubernetes.Clientset, cmr *ConfigmapResource) {
	configmapsClient := clientset.CoreV1().ConfigMaps(cmr.Namespace)

	configmap, _ := configmapsClient.Get(context.TODO(), cmr.Name, metav1.GetOptions{})
	if configmap.Name == cmr.Name {
		configmap = &apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: cmr.Name,
			},
			BinaryData: cmr.BinaryData,
		}

		result, err := configmapsClient.Update(context.TODO(), configmap, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Update configmap: %s success\n", result.GetName())
	}
}

func DeleteConfigmap(clientset *kubernetes.Clientset, namespace, name string) {
	err := clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

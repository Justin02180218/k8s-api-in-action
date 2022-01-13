package resources

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateService(clientset *kubernetes.Clientset, sr *ServiceResource) {
	servicesClient := clientset.CoreV1().Services(sr.Namespace)

	service, _ := servicesClient.Get(context.TODO(), sr.Name, metav1.GetOptions{})
	if service.Name == sr.Name {
		fmt.Printf("Service: %s is already exist.\n", sr.Name)
		return
	}

	service = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: sr.Name,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{{
				Name: sr.Ports[0].Name,
				Port: sr.Ports[0].Port,
				// NodePort: sr.Ports[0].NodePort,
			}},
			Selector: sr.Selector,
			Type:     apiv1.ServiceTypeClusterIP,
		},
	}

	result, err := servicesClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create service: %s success\n", result.GetName())

}

func DeleteService(clientset *kubernetes.Clientset, namespace, name string) {
	err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

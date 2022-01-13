package resources

import (
	"context"
	"fmt"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateIngress(clientset *kubernetes.Clientset, ir *IngressResource) {
	ingresssClient := clientset.NetworkingV1().Ingresses(ir.Namespace)

	ingress, _ := ingresssClient.Get(context.TODO(), ir.Name, metav1.GetOptions{})
	if ingress.Name == ir.Name {
		fmt.Printf("Ingress: %s is already exist.\n", ir.Name)
		return
	}

	prefix := netv1.PathTypePrefix
	ingress = &netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: ir.Name,
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "nginx",
			},
		},
		Spec: netv1.IngressSpec{
			DefaultBackend: &netv1.IngressBackend{
				Service: &netv1.IngressServiceBackend{
					Name: ir.ServiceName[0],
					Port: netv1.ServiceBackendPort{
						Number: ir.ServicePort[0],
					},
				},
			},
			Rules: []netv1.IngressRule{
				{
					Host: ir.Host[0],
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{{
								Path:     "/",
								PathType: &prefix,
								Backend: netv1.IngressBackend{
									Service: &netv1.IngressServiceBackend{
										Name: ir.ServiceName[0],
										Port: netv1.ServiceBackendPort{
											Number: ir.ServicePort[0],
										},
									},
								},
							}},
						},
					},
				},
				{
					Host: ir.Host[1],
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{{
								Path:     "/",
								PathType: &prefix,
								Backend: netv1.IngressBackend{
									Service: &netv1.IngressServiceBackend{
										Name: ir.ServiceName[1],
										Port: netv1.ServiceBackendPort{
											Number: ir.ServicePort[1],
										},
									},
								},
							}},
						},
					},
				},
			},
		},
	}

	result, err := ingresssClient.Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create ingress: %s success\n", result.GetName())
}

func DeleteIngress(clientset *kubernetes.Clientset, namespace, name string) {
	err := clientset.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

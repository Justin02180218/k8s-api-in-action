package resources

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateDeployment(clientset *kubernetes.Clientset, dr *DeploymentResource) {
	deploymentsClient := clientset.AppsV1().Deployments(dr.Namespace)

	deployment, _ := deploymentsClient.Get(context.TODO(), dr.Name, metav1.GetOptions{})
	if deployment.Name == dr.Name {
		fmt.Printf("Deployment: %s is already exist.\n", dr.Name)
		return
	}

	deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dr.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &dr.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: dr.Matchlabels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: dr.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{{
						Name:            dr.Containers[0].Name,
						Image:           dr.Containers[0].Image,
						ImagePullPolicy: "IfNotPresent",
						Ports: []apiv1.ContainerPort{{
							Name:          dr.Containers[0].Ports[0].Name,
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: dr.Containers[0].Ports[0].Port,
						}},
						VolumeMounts: []apiv1.VolumeMount{{
							Name:      dr.VolumesName,
							MountPath: dr.MountPath,
						}},
					}},
					Volumes: []apiv1.Volume{{
						Name: dr.VolumesName,
						VolumeSource: apiv1.VolumeSource{
							ConfigMap: &apiv1.ConfigMapVolumeSource{
								LocalObjectReference: apiv1.LocalObjectReference{
									Name: dr.ConfigmapName,
								},
							},
						},
					}},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create deployment: %s success\n", result.GetName())

}

// 扩缩容，滚动升级
// TODO

func DeleteDeployment(clientset *kubernetes.Clientset, namespace, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		panic(err)
	}
}

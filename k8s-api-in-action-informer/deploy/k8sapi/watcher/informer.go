package watcher

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func PodWatcher(clientset *kubernetes.Clientset, namespace string) {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)
			fmt.Printf("Add pod: %s\n", pod.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*apiv1.Pod)
			newPod := newObj.(*apiv1.Pod)
			fmt.Printf("Update old pod:%s to new pod:%s\n", oldPod.Name, newPod.Name)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)
			fmt.Printf("Delete pod: %s\n", pod.Name)
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)
	informerFactory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, podInformer.Informer().HasSynced) {
		return
	}

	// lister := podInformer.Lister()
	// pods, err := lister.Pods(namespace).List(labels.Everything())
	// if err != nil {
	// 	panic(err)
	// }
	// for _, pod := range pods {
	// 	fmt.Printf("Namespace:%s, =>Pod: %s\n", namespace, pod.Name)
	// }

	<-stopCh
}

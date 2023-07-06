package libs

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ClientSetFunc(file string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", file)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

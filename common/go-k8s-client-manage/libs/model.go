package libs

import (
	"k8s.io/client-go/kubernetes"
)

type ClientSet struct {
	Clientset   *kubernetes.Clientset
	NameSpace   string
	ServiceName string
	Works       string
}

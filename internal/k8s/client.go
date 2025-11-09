package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type client struct {
	kube *kubernetes.Clientset
}

//init client

func NewKubeClient(kubeclien string) (*client, error) {
	var cfg *rest.Config
	fmt.Printf("cfg: %v\n", cfg)
	return nil, nil
}

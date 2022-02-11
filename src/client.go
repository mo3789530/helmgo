package src

import "k8s.io/client-go/kubernetes"

type K8SClient interface {
	GetNameSpaces() ([]string, error)
	// GetPods(namespace string) ([]string, error)
	// CreateNameSpace(namespace string) error
}

type k8sClient struct {
	Clientset *kubernetes.Clientset
}

func Newk8sClient(clientset *kubernetes.Clientset) K8SClient {
	return &k8sClient{
		Clientset: clientset,
	}
}

func (c *k8sClient) GetNameSpaces() ([]string, error) {
	c.Clientset.CoreV1().Namespaces()
}

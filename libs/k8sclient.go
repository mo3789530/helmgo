package libs

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8SClient interface {
	GetNameSpaces() ([]string, error)
	GetPods(ns string) ([]string, error)
	CreateNameSpace(Name string) error
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
	ns, err := c.Clientset.CoreV1().Namespaces().List(
		context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var arr []string
	for _, v := range ns.Items {
		arr = append(arr, v.ObjectMeta.Name)
	}
	return arr, err
}

func (c *k8sClient) GetPods(ns string) ([]string, error) {
	pods, err := c.Clientset.CoreV1().Pods(ns).List(
		context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var arr []string
	for _, v := range pods.Items {
		arr = append(arr, v.ObjectMeta.Name)
	}
	return arr, nil
}

func (c *k8sClient) CreateNameSpace(name string) error {
	object := &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: corev1.NamespaceSpec{},
	}

	_, err := c.Clientset.CoreV1().Namespaces().Create(
		context.TODO(), object, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

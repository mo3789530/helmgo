package main

import (
	"flag"
	"gopls-workspace/libs"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func inCluster() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return config, err
}

func outCluster() (*rest.Config, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}
	return config, err
}

func main() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Not found env local")
	}

	isIncluster, err := strconv.ParseBool(os.Getenv("IN_CLUSTER"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	var config *rest.Config
	if isIncluster {
		config, err = inCluster()
	} else {
		config, err = outCluster()
	}
	if err != nil {
		log.Fatalf(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	k8sClient := libs.Newk8sClient(clientset)

	// debug create new namespace
	// err = k8sClient.CreateNameSpace("test1")
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// debug list namespace
	ns, err := k8sClient.GetNameSpaces()
	if err != nil {
		log.Println(err.Error())
	}
	for _, v := range ns {
		println(v)
	}

}

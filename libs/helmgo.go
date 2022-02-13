package libs

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

const (
	chartPath = "/charts/keycloak"
)

type HelmGo interface {
	Install(ns string) error
}

type helmgo struct {
	Config *rest.Config
}

func NewHelmGo(config *rest.Config) HelmGo {
	return &helmgo{
		Config: config,
	}
}

func (c *helmgo) Install(ns string) error {

	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	releaseName := ns
	releaseNamespace := ns
	namespace := ns

	kubeConfig := genericclioptions.NewConfigFlags(false)
	kubeConfig.APIServer = &c.Config.Host
	kubeConfig.BearerToken = &c.Config.BearerToken
	kubeConfig.CAFile = &c.Config.CAFile
	kubeConfig.Namespace = &namespace

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kubeConfig, releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		log.Println(err.Error())
		return err
	}

	iCli := action.NewInstall(actionConfig)
	iCli.Namespace = releaseNamespace
	iCli.ReleaseName = releaseName
	rel, err := iCli.Run(chart, nil)
	if err != nil {
		return err
	}
	fmt.Println("Successfully installed release: ", rel.Name)
	return nil
}

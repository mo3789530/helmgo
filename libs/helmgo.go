package libs

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var storage = repo.File{}

const (
	defaultCachePath            = "/tmp/.helmcache"
	defaultRepositoryConfigPath = "/tmp/.helmrepo"
)

type HelmGo interface {
}

type helmgo struct {
	Client *HelmClient
}

func NewHelmGo(options *Options) HelmGo {
	settings := cli.New()

	err := setEnvSettings(options, settings)
	if err != nil {
		log.Fatalf(err.Error())
	}
	clinet, err := newClient(options, settings.RESTClientGetter(), settings)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &helmgo{
		Client: clinet,
	}
}

func newClient(options *Options, clientGetter genericclioptions.RESTClientGetter, settings *cli.EnvSettings) (*HelmClient, error) {
	err := setEnvSettings(options, settings)
	if err != nil {
		return nil, err
	}

	debugLog := options.DebugLog
	if debugLog == nil {
		debugLog = func(format string, v ...interface{}) {
			log.Printf(format, v...)
		}
	}

	actionConfig := new(action.Configuration)
	err = actionConfig.Init(
		clientGetter,
		settings.Namespace(),
		os.Getenv("HELM_DRIVER"),
		debugLog,
	)
	if err != nil {
		return nil, err
	}

	getter.All(settings)
	return &HelmClient{
		Settings:     settings,
		Providers:    getter.All(settings),
		storage:      &storage,
		ActionConfig: actionConfig,
		linting:      options.Linting,
		DebugLog:     debugLog,
	}, nil
}

func setEnvSettings(options *Options, settings *cli.EnvSettings) error {
	if options == nil {
		options = &Options{
			RepositoryConfig: defaultRepositoryConfigPath,
			RepositoryCache:  defaultCachePath,
			Linting:          true,
		}
	}

	// set the namespace with this ugly workaround because cli.EnvSettings.namespace is private
	// thank you helm!
	if options.Namespace != "" {
		pflags := pflag.NewFlagSet("", pflag.ContinueOnError)
		settings.AddFlags(pflags)
		err := pflags.Parse([]string{"-n", options.Namespace})
		if err != nil {
			return err
		}
	}

	if options.RepositoryConfig == "" {
		options.RepositoryConfig = defaultRepositoryConfigPath
	}

	if options.RepositoryCache == "" {
		options.RepositoryCache = defaultCachePath
	}

	settings.RepositoryCache = options.RepositoryCache
	settings.RepositoryConfig = options.RepositoryConfig
	settings.Debug = options.Debug

	return nil
}

func (c *HelmClient) InitHelm() {

}

func Install() {

}

func DeptUp() {

}

func (c *HelmClient) GetRelease(name string) (*release.Release, error) {
	return c.GetRelease(name)
}

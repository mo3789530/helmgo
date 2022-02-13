package libs

// var storage = repo.File{}

// const (
// 	defaultCachePath            = "/tmp/.helmcache"
// 	defaultRepositoryConfigPath = "/tmp/.helmrepo"
// )

// type HelmClient interface {
// 	GetRelease(name string) (*release.Release, error)
// }

// type helmclient struct {
// 	Client *HelmClient
// }

// func NewHelmClient(options *Options) HelmGo {
// 	settings := cli.New()

// 	err := setEnvSettings(options, settings)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	clinet, err := newClient(options, settings.RESTClientGetter(), settings)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	return &helmgo{
// 		Client: clinet,
// 	}
// }

// func newClient(options *Options, clientGetter genericclioptions.RESTClientGetter, settings *cli.EnvSettings) (*HelmClient, error) {
// 	err := setEnvSettings(options, settings)
// 	if err != nil {
// 		return nil, err
// 	}

// 	debugLog := options.DebugLog
// 	if debugLog == nil {
// 		debugLog = func(format string, v ...interface{}) {
// 			log.Printf(format, v...)
// 		}
// 	}

// 	actionConfig := new(action.Configuration)
// 	err = actionConfig.Init(
// 		clientGetter,
// 		settings.Namespace(),
// 		os.Getenv("HELM_DRIVER"),
// 		debugLog,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	providers := getter.All(settings)
// 	return &HelmClient{
// 		Settings:     settings,
// 		Providers:    providers,
// 		storage:      &storage,
// 		ActionConfig: actionConfig,
// 		linting:      options.Linting,
// 		DebugLog:     debugLog,
// 	}, nil
// }

// func setEnvSettings(options *Options, settings *cli.EnvSettings) error {
// 	if options == nil {
// 		options = &Options{
// 			RepositoryConfig: defaultRepositoryConfigPath,
// 			RepositoryCache:  defaultCachePath,
// 			Linting:          true,
// 		}
// 	}

// 	if options.Namespace != "" {
// 		pflags := pflag.NewFlagSet("", pflag.ContinueOnError)
// 		settings.AddFlags(pflags)
// 		err := pflags.Parse([]string{"-n", options.Namespace})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	if options.RepositoryConfig == "" {
// 		options.RepositoryConfig = defaultRepositoryConfigPath
// 	}

// 	if options.RepositoryCache == "" {
// 		options.RepositoryCache = defaultCachePath
// 	}

// 	settings.RepositoryCache = options.RepositoryCache
// 	settings.RepositoryConfig = options.RepositoryConfig
// 	settings.Debug = options.Debug

// 	return nil
// }

// func (c *helmgo) InitHelm() {

// }

// func (c *helmgo) Install() {

// }

// func DeptUp() {

// }

// func (c *helmgo) UninstallReleaseByName(name string) error {
// 	return c.uninstallReleaseByName(name)
// }

// func (c *helmgo) GetRelease(name string) (*release.Release, error) {
// 	return c.getRelease(name)
// }

// // func (c *helmgo) install(ctx context.Context, spec *ChartSpec) (*release.Release, error) {
// // 	client := action.NewInstall(c.Client.ActionConfig)
// // 	mergeInstallOptions(spec, client)

// // 	if client.Version == "" {
// // 		client.Version = ">0.0.0-0"
// // 	}

// // 	helmChart, chartPath, err := c.Client.getChart(spec.ChartName, &client.ChartPathOptions)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	if helmChart.Metadata.Type != "" && helmChart.Metadata.Type != "application" {
// // 		return nil, fmt.Errorf(
// // 			"chart %q has an unsupported type and is not installable: %q",
// // 			helmChart.Metadata.Name,
// // 			helmChart.Metadata.Type,
// // 		)
// // 	}

// // 	if req := helmChart.Metadata.Dependencies; req != nil {
// // 		if err := action.CheckDependencies(helmChart, req); err != nil {
// // 			if client.DependencyUpdate {
// // 				man := &downloader.Manager{
// // 					ChartPath:        chartPath,
// // 					Keyring:          client.ChartPathOptions.Keyring,
// // 					SkipUpdate:       false,
// // 					Getters:          c.Client.Providers,
// // 					RepositoryConfig: c.Client.Settings.RepositoryConfig,
// // 					RepositoryCache:  c.Client.Settings.RepositoryCache,
// // 				}
// // 				if err := man.Update(); err != nil {
// // 					return nil, err
// // 				}
// // 			} else {
// // 				return nil, err
// // 			}
// // 		}
// // 	}

// // 	values, err := spec.GetValuesMap()
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	if c.Client.linting {
// // 		err = c.Client.lint(chartPath, values)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 	}

// // 	rel, err := client.RunWithContext(ctx, helmChart, values)
// // 	if err != nil {
// // 		return rel, err
// // 	}

// // 	c.Client.DebugLog("release installed successfully: %s/%s-%s", rel.Name, rel.Chart.Metadata.Name, rel.Chart.Metadata.Version)

// // 	return rel, nil

// // }

// func (spec *ChartSpec) GetValuesMap() (map[string]interface{}, error) {
// 	var values map[string]interface{}

// 	err := yaml.Unmarshal([]byte(spec.ValuesYaml), &values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return values, nil
// }

// // func (c *HelmClient) getChart(chartName string, chartPathOptions *action.ChartPathOptions) (*chart.Chart, string, error) {
// // 	chartPath, err := chartPathOptions.LocateChart(chartName, c.Settings)
// // 	if err != nil {
// // 		return nil, "", err
// // 	}

// // 	helmChart, err := loader.Load(chartPath)
// // 	if err != nil {
// // 		return nil, "", err
// // 	}

// // 	if helmChart.Metadata.Deprecated {
// // 		c.DebugLog("WARNING: This chart (%q) is deprecated", helmChart.Metadata.Name)
// // 	}

// // 	return helmChart, chartPath, err
// // }

// func (c *helmgo) uninstallReleaseByName(name string) error {
// 	client := action.NewUninstall(c.Client.ActionConfig)

// 	resp, err := client.Run(name)
// 	if err != nil {
// 		return err
// 	}

// 	c.Client.DebugLog("release uninstalled, response: %v", resp)

// 	return nil
// }

// func (c *helmgo) getRelease(name string) (*release.Release, error) {
// 	getReleaseClient := action.NewGet(c.Client.ActionConfig)
// 	return getReleaseClient.Run(name)
// }

// // func mergeInstallOptions(chartSpec *ChartSpec, installOptions *action.Install) {
// // 	installOptions.CreateNamespace = chartSpec.CreateNamespace
// // 	installOptions.DisableHooks = chartSpec.DisableHooks
// // 	installOptions.Replace = chartSpec.Replace
// // 	installOptions.Wait = chartSpec.Wait
// // 	installOptions.DependencyUpdate = chartSpec.DependencyUpdate
// // 	installOptions.Timeout = chartSpec.Timeout
// // 	installOptions.Namespace = chartSpec.Namespace
// // 	installOptions.ReleaseName = chartSpec.ReleaseName
// // 	installOptions.Version = chartSpec.Version
// // 	installOptions.GenerateName = chartSpec.GenerateName
// // 	installOptions.NameTemplate = chartSpec.NameTemplate
// // 	installOptions.Atomic = chartSpec.Atomic
// // 	installOptions.SkipCRDs = chartSpec.SkipCRDs
// // 	installOptions.DryRun = chartSpec.DryRun
// // 	installOptions.SubNotes = chartSpec.SubNotes
// // }

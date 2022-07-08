package models

import (
	"errors"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func GetHelmReleaseInfo(releaseName string, namespace string, kubeConfigPath string) (*release.Info, error) {
	actionConfig := new(action.Configuration)
	settings := cli.New()
	settings.KubeConfig = kubeConfigPath
	// You can pass an empty string instead of settings.Namespace() to list
	// all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace,
		"secret", log.Printf); err != nil {
		log.Printf("%+v", err)
		return nil, errors.New("kube config not valid")
	}
	client := action.NewGet(actionConfig)
	//var out io.Writer
	res, err := client.Run(releaseName)
	if err != nil {
		return nil, err
	}
	return res.Info, nil
}

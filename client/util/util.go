package util

import (
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-wrapper/application"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/fileutils"
)

func GetAppsFromDisk() ([]*application.Application, error) {
	currentDir, err := fileutils.GetCurrentDir()
	if err != nil {
		return nil, err
	}
	return application.NewAppsFromDisk(currentDir+config.GetRootFolderOrDefault(), config.GetNamespaceOrDefault())
}

func GetAppFromDisk(appName string) (*application.Application, error) {
	currentDir, err := fileutils.GetCurrentDir()
	if err != nil {
		return nil, err
	}
	return application.NewAppFromDisk(currentDir+config.GetRootFolderOrDefault(), appName, config.GetNamespaceOrDefault())
}

func GetAppFromNamespace(appName string, labelSelector string) (*application.Application, error) {
	restConf, err := client.GetRestConfig(config.IsInClusterModeOrDefault())
	if err != nil {
		return nil, err
	}
	return application.NewAppFromNamespace(appName, config.GetNamespaceOrDefault(), restConf, labelSelector)
}

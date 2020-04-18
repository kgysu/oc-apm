package cmdsApp

import (
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-wrapper/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func RunCliCmdAppDelete(input []string) error {
	name := ""
	if len(input) > 1 {
		name = input[1]
	}
	prj, err := util.GetAppFromDisk(name)
	if err != nil {
		return err
	}
	restConf, err := client.GetRestConfig(config.IsInClusterModeOrDefault())
	if err != nil {
		return err
	}
	prj.DeleteItems(os.Stdout, config.GetNamespaceOrDefault(), restConf, &v1.DeleteOptions{})
	return nil
}

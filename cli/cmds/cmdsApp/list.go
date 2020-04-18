package cmdsApp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
)

func RunCliCmdAppList(input []string) error {
	name := ""
	if len(input) > 1 {
		name = input[1]
	}

	apps, err := util.GetAppsFromDisk()
	if err != nil {
		return err
	}
	for i, prj := range apps {
		if prj.Name == name || name == "" {
			fmt.Printf("%d: App [%s] (%d) \n", i, prj.Name, len(prj.Items))
		}
	}
	return nil
}

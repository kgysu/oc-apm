package cmdsApp

import (
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-wrapper/templates"
	"os"
)

func RunCliCmdAppNew(input []string) error {
	subCmd, name, labelSelector := parseArgs(input)
	writer := os.Stdout
	if subCmd == "sample" {
		sampleApp := templates.GetTemplateApp("sample")
		sampleApp.Save(writer, config.GetRootFolderOrDefault())
	}
	if subCmd == "template" {
		tempApp := templates.GetTemplateApp(name)
		tempApp.Save(writer, config.GetRootFolderOrDefault())
	}
	if subCmd == "from" {
		prj, err := util.GetAppFromNamespace(name, labelSelector)
		if err != nil {
			return err
		}
		prj.Save(writer, config.GetRootFolderOrDefault())
	}
	return nil
}

func parseArgs(args []string) (string, string, string) {
	subCmd := args[1]
	args = args[1:]
	name := "sample"
	if len(args) > 1 {
		name = args[1]
	}
	labelSelector := ""
	if len(args) > 2 {
		labelSelector = args[2]
	}
	return subCmd, name, labelSelector
}

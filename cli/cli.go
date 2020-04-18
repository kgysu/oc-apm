package cli

import (
	"github.com/kgysu/oc-apm/cli/cmds"
	"github.com/kgysu/oc-apm/cli/cmds/cmdsApp"
	"github.com/kgysu/oc-apm/cli/cmds/cmdsList"
)

func RunCli(input []string) error {
	if len(input) > 1 {
		cmd := input[1]
		args := input[1:]

		if cmd == "start" {
			return cmds.RunCliCmdStartServer(args)
		}
		if cmd == "app" {
			return cmdsApp.RunCliCmdApp(args)
		}
		if cmd == "list" {
			return cmdsList.RunCliCmdList(args)
		}
		if cmd == "install" {
			return cmds.RunCliCmdInstallSelf()
		}
		if cmd == "remove" {
			return cmds.RunCliCmdRemoveSelf()
		}
	}
	cmds.RunCliCmdHelp(input)
	return nil
}

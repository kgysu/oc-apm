package cmds

import (
	"fmt"
	"github.com/kgysu/oc-apm/config"
)

func RunCliCmdHelp(args []string) {
	fmt.Print(config.HelpText)
}

package cmds

import (
	"fmt"
	"github.com/kgysu/oc-apm/config"
)

func RunCliCmdVersion(input []string) error {
	fmt.Printf("OC-APM version: %s\n", config.GetVersion())
	return nil
}

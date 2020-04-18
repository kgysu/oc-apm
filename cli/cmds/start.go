package cmds

import (
	"github.com/kgysu/oc-apm/web/server"
)

const defaultLocalPort = ":8082"

func RunCliCmdStartServer(input []string) error {
	port := parseParamsForStart(input[1:])
	if port == "" {
		port = defaultLocalPort
	}
	err := server.RunServer(port)
	if err != nil {
		return err
	}
	return nil
}

func parseParamsForStart(input []string) string {
	if len(input) == 0 {
		return ""
	}
	return input[0]
}

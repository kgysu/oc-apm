package main

import (
	"github.com/kgysu/oc-apm/cli"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-apm/web/server"
	"os"
	"strings"
)

func main() {
	args := os.Args
	index := extractAndSetNamespace(args)
	if index != 0 {
		args = append(args[:index], args[index+1:]...)
	}

	runInClusterEnv := os.Getenv(config.GetInClusterEnvNameOrDefault())
	if runInClusterEnv != "" {
		config.SetInClusterMode(true)
		config.SetNamespace(os.Getenv(config.GetNamespaceEnvNameOrDefault()))
	}

	if config.IsInClusterModeOrDefault() {
		err := server.RunServer(config.GetPortOrDefault())
		if err != nil {
			panic(err)
		}
	} else {
		err := cli.RunCli(args)
		if err != nil {
			panic(err)
		}
	}
}

func extractAndSetNamespace(inputs []string) int {
	for i, in := range inputs {
		if strings.HasPrefix(in, "namespace=") {
			config.SetNamespace(strings.Replace(in, "namespace=", "", 1))
			return i
		}
		if strings.HasPrefix(in, "ns=") {
			config.SetNamespace(strings.Replace(in, "ns=", "", 1))
			return i
		}
	}
	return 0
}

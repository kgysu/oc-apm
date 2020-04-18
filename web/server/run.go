package server

import (
	"fmt"
	"github.com/kgysu/oc-apm/config"
	"net/http"
)

// Run an HTTP server which will serve the web pages.
func RunServer(port string) error {
	mux := http.NewServeMux()
	RegisterAllHandlers(mux)

	fmt.Printf("Start server on [%v] namespace=[%s] inCluster=[%v]\n", port, config.GetNamespaceOrDefault(), config.IsInClusterModeOrDefault())
	err := http.ListenAndServe(port, mux)
	if err != nil {
		return err
	}
	return nil
}

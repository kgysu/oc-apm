package cmdsapp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"github.com/kgysu/oc-wrapper/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strings"
)

func HandleCmdAppDelete(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/plain")
	name := strings.Replace(req.RequestURI, urls.CmdAppDeleteUrl, "", 1)

	prj, err := util.GetAppFromDisk(name)
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	restConf, err := client.GetRestConfig(config.IsInClusterModeOrDefault())
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}
	prj.DeleteItems(rw, config.GetNamespaceOrDefault(), restConf, &v1.DeleteOptions{})
}

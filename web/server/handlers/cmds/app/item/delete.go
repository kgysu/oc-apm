package cmdsappitem

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

func HandleCmdAppItemDelete(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/plain")
	urlEnding := strings.Replace(req.RequestURI, urls.CmdAppItemDeleteUrl, "", 1)
	appName, itemKind, itemName := extractFromUrl(urlEnding)

	prj, err := util.GetAppFromDisk(appName)
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	restConf, err := client.GetRestConfig(config.IsInClusterModeOrDefault())
	if err != nil {
		fmt.Fprintln(rw, err.Error())
		return
	}

	item := prj.GetItem(itemKind, itemName)
	if item != nil {
		err = item.Delete(config.GetNamespaceOrDefault(), restConf, &v1.DeleteOptions{})
		if err != nil {
			fmt.Fprintln(rw, err.Error())
		} else {
			fmt.Fprintf(rw, "%s deleted \n", item.Info())
		}
	} else {
		fmt.Fprintln(rw, "item not found")
	}
}

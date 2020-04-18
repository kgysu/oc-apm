package cmdsappitem

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"github.com/kgysu/oc-wrapper/client"
	"net/http"
	"strings"
)

func HandleCmdAppItemCreate(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/plain")
	urlEnding := strings.Replace(req.RequestURI, urls.CmdAppItemCreateUrl, "", 1)
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
		err = item.Create(config.GetNamespaceOrDefault(), restConf)
		if err != nil {
			fmt.Fprintln(rw, err.Error())
		} else {
			fmt.Fprintf(rw, "%s created \n", item.Info())
		}
	} else {
		fmt.Fprintln(rw, "item not found")
	}
}

func extractFromUrl(url string) (string, string, string) {
	splitted := strings.Split(url, "/")
	appName := "default"
	itemKind := "none"
	itemName := "none"
	if len(splitted) > 0 {
		appName = splitted[0]
	}
	if len(splitted) > 1 {
		itemKind = splitted[1]
	}
	if len(splitted) > 2 {
		itemName = splitted[2]
	}
	return appName, itemKind, itemName
}

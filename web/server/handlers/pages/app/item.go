package handlerapp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/web/html/pages/app"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
	"strings"
)

func HandleAppItemDetailsPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")
	urlEnding := strings.Replace(req.RequestURI, urls.AppItemDetailsPageUrl, "", 1)
	appName, itemKind, itemName := extractFromUrl(urlEnding)

	prj, err := util.GetAppFromDisk(appName)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}

	item := prj.GetItem(itemKind, itemName)
	if item != nil {
		content := htmlapp.CreateAppItemDetailsHtml(prj, item)
		_, err := fmt.Fprint(rw, content)
		if err != nil {
			fmt.Print(err)
			fmt.Fprint(rw, err)
			return
		}
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

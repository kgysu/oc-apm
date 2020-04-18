package handlerapp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	htmlapp "github.com/kgysu/oc-apm/web/html/pages/app"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
	"strings"
)

func HandleAppDetailsPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")
	name := strings.Replace(req.RequestURI, urls.AppDetailsPageUrl, "", 1)

	appFromDisk, err := util.GetAppFromDisk(name)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
	content := htmlapp.CreateAppDetailsHtml(appFromDisk)
	_, err = fmt.Fprint(rw, content)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
}

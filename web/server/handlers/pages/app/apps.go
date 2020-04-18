package handlerapp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	htmlapp "github.com/kgysu/oc-apm/web/html/pages/app"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
)

func HandleAppListPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")

	apps, err := util.GetAppsFromDisk()
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
	content := htmlapp.CreateAppListHtml(apps)
	_, err = fmt.Fprint(rw, content)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
}

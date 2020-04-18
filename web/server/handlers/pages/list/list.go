package handlerlist

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/web/html/pages/list"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
)

func HandleListPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")
	labelSelector, kinds := parseReq(req)

	appFromNamespace, err := util.GetAppFromNamespace("list", labelSelector)
	if err != nil {
		fmt.Fprint(rw, err)
		return
	}

	content := htmllist.CreateItemsListHtml(appFromNamespace.GetItemsByKinds(kinds), labelSelector, kinds)
	_, err = fmt.Fprint(rw, content)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
}

func parseReq(req *http.Request) (string, string) {
	label := req.FormValue("label")
	kinds := req.FormValue("kinds")
	return label, kinds
}

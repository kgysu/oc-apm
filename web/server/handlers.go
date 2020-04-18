package server

import (
	"fmt"
	htmlPages "github.com/kgysu/oc-apm/web/html/pages"
	"github.com/kgysu/oc-apm/web/server/handlers/cmds"
	cmdsapp "github.com/kgysu/oc-apm/web/server/handlers/cmds/app"
	cmdsappitem "github.com/kgysu/oc-apm/web/server/handlers/cmds/app/item"
	"github.com/kgysu/oc-apm/web/server/handlers/pages"
	"github.com/kgysu/oc-apm/web/server/handlers/pages/app"
	"github.com/kgysu/oc-apm/web/server/handlers/pages/list"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
)

func RegisterAllHandlers(mux *http.ServeMux) {
	mux.HandleFunc(urls.AppDetailsPageUrl, handlerapp.HandleAppDetailsPage)
	mux.HandleFunc(urls.AppNewPageUrl, handlerapp.HandleAppNewPage)
	mux.HandleFunc(urls.AppListPageUrl, handlerapp.HandleAppListPage)
	mux.HandleFunc(urls.AppItemDetailsPageUrl, handlerapp.HandleAppItemDetailsPage)
	mux.HandleFunc(urls.CmdAppNewUrl, cmdsapp.HandleCmdAppNew)
	mux.HandleFunc(urls.CmdAppCreateUrl, cmdsapp.HandleCmdAppCreate)
	mux.HandleFunc(urls.CmdAppUpdateUrl, cmdsapp.HandleCmdAppUpdate)
	mux.HandleFunc(urls.CmdAppDeleteUrl, cmdsapp.HandleCmdAppDelete)
	mux.HandleFunc(urls.CmdAppStartUrl, cmdsapp.HandleCmdAppStart)
	mux.HandleFunc(urls.CmdAppStopUrl, cmdsapp.HandleCmdAppStop)
	mux.HandleFunc(urls.CmdAppItemCreateUrl, cmdsappitem.HandleCmdAppItemCreate)
	mux.HandleFunc(urls.CmdAppItemUpdateUrl, cmdsappitem.HandleCmdAppItemUpdate)
	mux.HandleFunc(urls.CmdAppItemDeleteUrl, cmdsappitem.HandleCmdAppItemDelete)
	mux.HandleFunc(urls.CmdAppItemStartUrl, cmdsappitem.HandleCmdAppItemStart)
	mux.HandleFunc(urls.CmdAppItemStopUrl, cmdsappitem.HandleCmdAppItemStop)
	mux.HandleFunc(urls.OcCmdList, handlerlist.HandleListPage)
	mux.HandleFunc(urls.LoadFromGitUrl, cmds.HandleLoadFromGit)
	mux.HandleFunc(urls.GitPageUrl, pages.HandleGitPage)
	mux.HandleFunc("/", handleRootPath)
}

func handleRootPath(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")

	htmlContent := htmlPages.CreateHomePage()
	_, err := fmt.Fprintln(rw, htmlContent)
	if err != nil {
		fmt.Fprint(rw, err)
		fmt.Print(err)
		return
	}
}

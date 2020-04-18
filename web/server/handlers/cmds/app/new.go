package cmdsapp

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"github.com/kgysu/oc-wrapper/templates"
	"net/http"
)

func HandleCmdAppNew(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/plain")
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintln(rw, err.Error())
	}
	name, label, from := parseReq(req)

	if name == "" {
		fmt.Fprintln(rw, fmt.Errorf("error: name is empty, name=[%s]\n", name))
		return
	}
	if from == "on" {
		prj, err := util.GetAppFromNamespace(name, label)
		if err != nil {
			fmt.Fprintln(rw, err.Error())
			return
		}
		prj.Save(rw, config.GetRootFolderOrDefault())
	} else {
		tempApp := templates.GetTemplateApp(name)
		tempApp.Save(rw, config.GetRootFolderOrDefault())
	}
}

func parseReq(req *http.Request) (string, string, string) {
	name := req.PostForm.Get("name")
	label := req.PostForm.Get("label")
	from := req.PostForm.Get("from")
	return name, label, from
}

package cmds

import (
	"fmt"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"github.com/kgysu/oc-wrapper/fileutils"
	"github.com/kgysu/oc-wrapper/gitutils"
	"net/http"
)

func HandleLoadFromGit(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/plain")
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintln(rw, err.Error())
		return
	}
	subPath, repoUrl, tag, branch := parseReq(req)
	fmt.Fprintf(rw, "Load from [%s%s] tag=[%s] branch=[%s]\n", repoUrl, subPath, tag, branch)

	if repoUrl == "" {
		fmt.Fprintln(rw, fmt.Errorf("missing url"))
		return
	}

	currentDir, err := fileutils.GetCurrentDir()
	if err != nil {
		fmt.Fprintln(rw, err.Error())
		return
	}

	err = gitutils.LoadFromGitRepo(rw, currentDir, subPath, repoUrl, tag, branch)
	if err != nil {
		fmt.Fprintln(rw, err.Error())
		return
	}
}

func parseReq(req *http.Request) (string, string, string, string) {
	subPath := req.PostForm.Get("subPath")
	repoUrl := req.PostForm.Get("repoUrl")
	tag := req.PostForm.Get("tag")
	branch := req.PostForm.Get("branch")
	return subPath, repoUrl, tag, branch
}

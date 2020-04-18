package pages

import (
	"fmt"
	htmlPages "github.com/kgysu/oc-apm/web/html/pages"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
)

func HandleGitPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")

	content := htmlPages.CreateGitPageHtml()
	_, err := fmt.Fprint(rw, content)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
}

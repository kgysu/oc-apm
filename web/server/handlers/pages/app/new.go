package handlerapp

import (
	"fmt"
	htmlapp "github.com/kgysu/oc-apm/web/html/pages/app"
	"github.com/kgysu/oc-apm/web/server/serverutil"
	"net/http"
)

func HandleAppNewPage(rw http.ResponseWriter, req *http.Request) {
	serverutil.SetHeaders(rw, req, "text/html")

	content := htmlapp.CreateAppNewHtml()
	_, err := fmt.Fprint(rw, content)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(rw, err)
		return
	}
}

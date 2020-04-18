package serverutil

import "net/http"

func SetHeaders(rw http.ResponseWriter, req *http.Request, contentType string) {
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Content-Type", contentType)
}

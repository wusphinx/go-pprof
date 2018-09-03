package register

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HttpRouter(router *httprouter.Router) {
	router.Handler(http.MethodGet, "/debug/pprof/:item", http.DefaultServeMux)
}

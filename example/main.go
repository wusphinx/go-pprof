package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wusphinx/go-pprof/register"
)

func PprofTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pprof test! \n")
}

func main() {
	router := httprouter.New()
	register.HttpRouter(router)

	router.GET("/", PprofTest)
	fmt.Println(http.ListenAndServe(":8000", router))
}

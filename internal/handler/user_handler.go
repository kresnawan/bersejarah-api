package handler

import (
	"app/internal/storage"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexPage(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	res.Header().Set("Content-Type", "text/json")
	var data []byte = storage.DbConnect()
	res.Write(data)
}

func HelloPage(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(res, "Hello there!, %s!", ps.ByName("name"))
}

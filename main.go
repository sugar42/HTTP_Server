package main

import (
	"net/http"

	"github.com/sugar42/HTTP_Server/api"
)


func main(){
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)	

}
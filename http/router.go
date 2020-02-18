package router

import (
	"net/http"
)
type Router interface{
	GET(url string,f func(response http.ResponseWriter, request *http.Request))
	POST(url string,f func(response http.ResponseWriter, request *http.Request))
	SERVE(port string)
}


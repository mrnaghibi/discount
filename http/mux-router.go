package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type muxRouter struct{}


var(
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router{
	return &muxRouter{}
}


func (*muxRouter) GET(url string,f func(response http.ResponseWriter, request *http.Request)){
	muxDispatcher.HandleFunc(url,f).Methods("GET")
}
func (*muxRouter) POST(url string,f func(response http.ResponseWriter, request *http.Request)){
	muxDispatcher.HandleFunc(url,f).Methods("POST")
}
func (*muxRouter) SERVE(port string){
	log.Println("Mux - Server Up And Running On Port: ",port)
	log.Fatal(http.ListenAndServe(port,muxDispatcher))
}
package main

import {
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
}

func Index(2 http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcom! v1\n")
}

func main(){
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080",router))
}
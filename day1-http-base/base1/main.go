package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)

	log.Fatal(http.ListenAndServe(":9999",nil))

}

func indexHandler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
}


func helloHandler(w http.ResponseWriter,req *http.Request){
	for s, strings := range req.Header {

		fmt.Fprintf(w,"Heder[%q]=%q\n",s,strings)

	}
}
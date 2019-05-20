package main

import(
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter,request *http.Request){
	fmt.FPrintf(writer, "Hello World, %s", request.URL.Path[1:])
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

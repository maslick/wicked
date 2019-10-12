package main

import (
	rest "github.com/maslick/gowiki/logic"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", rest.ViewHandler)
	http.HandleFunc("/edit/", rest.EditHandler)
	http.HandleFunc("/save/", rest.SaveHandler)
	http.HandleFunc("/", rest.AllHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

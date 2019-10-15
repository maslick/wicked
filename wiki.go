package main

import (
	"github.com/google/logger"
	rest "github.com/maslick/gowiki/logic"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", rest.ViewHandler)
	http.HandleFunc("/edit/", rest.EditHandler)
	http.HandleFunc("/save/", rest.SaveHandler)
	http.HandleFunc("/all/", rest.AllHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	logger.Init("LoggerExample", true, false, ioutil.Discard)
}

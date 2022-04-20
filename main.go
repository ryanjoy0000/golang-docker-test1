package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var tpl *template.Template

const (
	webServerAdd = "localhost:5555"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("App running on arch: ", runtime.GOARCH, runtime.GOOS)
	fmt.Println("App listening on port: ", webServerAdd)

	chiMux := chi.NewRouter()
	chiMux.Use(middleware.Logger)
	chiMux.Get("/", homePage)
	chiMux.Get("/contact", contactPage)

	fmt.Println("Starting web server...")
	err := http.ListenAndServe(webServerAdd, chiMux)
	handleErr(err)
}

func homePage(respW http.ResponseWriter, req *http.Request) {
	respW.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(respW, "index.html", nil)
	handleHttpErr(err, respW)
}

func contactPage(respW http.ResponseWriter, req *http.Request) {
	respW.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(respW, "contact.html", nil)
	handleHttpErr(err, respW)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func handleHttpErr(err error, respW http.ResponseWriter) {
	if err != nil {
		http.Error(respW, "", http.StatusInternalServerError)
		log.Println("Error in http", err)
		return
	}
}

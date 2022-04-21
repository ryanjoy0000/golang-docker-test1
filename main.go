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
	webServerAdd = ":5555"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("App running on arch: ", runtime.GOARCH, runtime.GOOS)
	fmt.Println("App listening on port: ", webServerAdd)

	fmt.Println("Starting web server...")
	// normalServer()
	chiServer()
}

func normalServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe(webServerAdd, nil)
	handleErr(err)
}

func chiServer() {
	chiMux := chi.NewRouter()
	chiMux.Use(middleware.Logger)
	chiMux.Get("/", homePage)
	chiMux.Get("/contact", contactPage)

	err := http.ListenAndServe(webServerAdd, chiMux)
	handleErr(err)
}

func homePage(respW http.ResponseWriter, req *http.Request) {
	// msg := "Home Page"
	// respW.Write([]byte(msg))

	respW.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(respW, "index.html", nil)
	handleHttpErr(err, respW)
}

func contactPage(respW http.ResponseWriter, req *http.Request) {
	// msg := "Contact Us Page"
	// respW.Write([]byte(msg))

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

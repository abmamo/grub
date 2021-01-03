package main

import (
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/home.html"))
	tpl.Execute(w, nil)
}

func main() {
	// get port to serve UI
	port := getEnvironment("PORT", ".env")
	// if no port specified in .env file
	if port == "" {
		// use 3000 as the default
		port = "3000"
	}
	// init mux router
	mux := http.NewServeMux()
	// init file server
	fs := http.FileServer(http.Dir("assets"))
	// connect file server to mux router
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// handlers
	mux.HandleFunc("/", homeHandler)
	// serve
	http.ListenAndServe(":"+port, mux)
}

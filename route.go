package main

/*
gophish - Open-Source Phishing Framework

The MIT License (MIT)

Copyright (c) 2013 Jordan Wright

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func createRouter() http.Handler {
	router := mux.NewRouter()
	// Base Front-end routes
	router.HandleFunc("/", Base)
	router.HandleFunc("/login", Login)
	router.HandleFunc("/register", Register)
	router.HandleFunc("/campaigns", Base_Campaigns)

	// Create the API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", API)
	api.HandleFunc("/campaigns", API_Campaigns)
	api.HandleFunc("/campaigns/{id}", API_Campaigns_Id_Post).Methods("POST")
	api.HandleFunc("/campaigns/{id}", API_Campaigns_Id_Get).Methods("GET")

	//Setup static file serving
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return router
}

func Register(w http.ResponseWriter, r *http.Request) {
	// If it is a post request, attempt to register the account
	// Now that we are all registered, we can log the user in
	Login(w, r)
}

func Base(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	renderTemplate(w, "index")
}

func Base_Campaigns(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, _ := template.ParseFiles("static/templates/" + tmpl + ".html")
	t.Execute(w, "T")
}

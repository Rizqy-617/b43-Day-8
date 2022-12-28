package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))


	route.HandleFunc("/", homePage).Methods("GET")
	route.HandleFunc("/project", projectPage).Methods("GET")
	route.HandleFunc("/project", addProject).Methods("POST")
	route.HandleFunc("/contact", contactPage).Methods("GET")

	fmt.Println("Server running on port:8080")
	http.ListenAndServe("localhost:8080", route)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func projectPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("view/myProject.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

type dataType struct {
	ID int
	Projectname string
	Description string
	Technologies []string
	Startdate string
	Enddate string
	Image string
	Duration string
}

var dataSubmit = []data{

}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)

	if err != nil {
		log.Fatal(err)
	}

	projectname := r.PostForm.Get("project-name")
	startDate := r.Form["technologies"]

	// fmt.Println("Project Name : " + r.PostForm.Get("project-name"))
	// fmt.Println("Start-date : " + r.PostForm.Get("start-date"))
	// fmt.Println("End-date : " + r.PostForm.Get("end-date"))
	// fmt.Println("Description : " + r.PostForm.Get("description"))
	
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("view/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

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

type dataReceive struct {
	ID int
	Projectname string
	Description string
	Technologies []string
	Startdate string
	Enddate string
	Image string
	Duration string
}

var dataSubmit = []dataReceive{

}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)

	if err != nil {
		log.Fatal(err)
	}

	projectname := r.PostForm.Get("project-name")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	description := r.PostForm.Get("description")
	technologies := r.Form["technologies"]


	//Buat Durasi
	const timeFormat = "2006-01-02"
	timeStartDate, err := time.Parse(timeFormat, startDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	timeEndDate, err := time.Parse(timeFormat, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	distance := timeEndDate.Sub(timeStartDate)
	dayDistance := int64(distance.Hours() / 24)

	var duration string
	if dayDistance >= 0 {
		duration =strconv.FormatInt(dayDistance, 10) + " Days"
	}

	img, imgname, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	defer img.Close()
	dir, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	filename := imgname.Filename
	fileLocation := filepath.Join(dir, "public/uploaded-image", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	defer targetFile.Close()
	if _, err := io.Copy(targetFile, img); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var newData = dataReceive{
		Projectname: projectname,
		Description: description,
		Technologies: technologies,
		Startdate: startDate,
		Enddate: endDate,
		Duration: duration,
		Image: imgname.Filename,
	} 

	fmt.Println("Project Name : " + projectname)
	fmt.Println("Start-date : " + startDate)
	fmt.Println("End-date : " + endDate)
	fmt.Println("Description : " + description)
	fmt.Println("Technologies : ", r.Form["technologies"] )
	fmt.Println("Duration : " + duration)
	fmt.Println("Image : " + imgname.Filename)

	dataSubmit = append(dataSubmit, newData)
	
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
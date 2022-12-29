package main

import (
	"fmt"
	"log"
	"net/http"
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
	route.HandleFunc("/project/{id}", detailProject).Methods("GET")
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

	dataCaller := map[string]interface{} {
		"Projects": dataSubmit,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, dataCaller)
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

// Struct buat menentukan variable sama tipe data nya
type dataReceive struct {
	ID int
	Projectname string
	Description string
	Technologies []string
	Startdate string
	Enddate string
	Duration string
}

// Nanti si variable dataSubmit ini bakal di isi sama value dari function di bawah
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
	timeStartDate, _:= time.Parse(timeFormat, startDate)
	timeEndDate, _:= time.Parse(timeFormat, endDate)

	// Hitung jaraka
	distance := timeEndDate.Sub(timeStartDate)

	//Ubah milisecond menjadi bulan, minggu dan hari
	monthDistance := int(distance.Hours() / 24 / 30)
	weekDistance := int(distance.Hours() / 24 / 7)
	daysDistance := int(distance.Hours() / 24)

	var duration string
	if monthDistance >= 1 && daysDistance <= 0{
		duration = strconv.Itoa(monthDistance) + " months"
	} else if monthDistance < 1 && weekDistance >= 1 {
		duration = strconv.Itoa(weekDistance) + " weeks"
	} else if monthDistance < 1 && daysDistance >= 0 {
		duration = strconv.Itoa(daysDistance) + " days"
	} else {
		duration = "0 days"
	}

	// Input Image Start
	// img, imgname, err := r.FormFile("image")// Buat ngambil datanya doang
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Message : " + err.Error()))
	// 	return
	// }

	// defer img.Close()
	// dir, err := os.Getwd() // buat c:/download
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Message : " + err.Error()))
	// 	return
	// }

	// filename := imgname.Filename // Buat negbuat nama file nya
	// fileLocation := filepath.Join(dir, "public/uploaded-image", filename)
	// targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Message : " + err.Error()))
	// 	return
	// }

	// defer targetFile.Close()
	// if _, err := io.Copy(targetFile, img); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Message : " + err.Error()))
	// 	return
	// }
	// Input Image End

	var newData = dataReceive{
		Projectname: projectname,
		Description: description,
		Technologies: technologies,
		Startdate: startDate,
		Enddate: endDate,
		Duration: duration,
	} 

	fmt.Println("Project Name : " + projectname)
	fmt.Println("Start-date : " + startDate)
	fmt.Println("End-date : " + endDate)
	fmt.Println("Description : " + description)
	fmt.Println("Technologies : ", r.Form["technologies"] )
	fmt.Println("Duration : " + duration)
	// fmt.Println("Image : " + imgname.Filename)

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

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("view/project-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
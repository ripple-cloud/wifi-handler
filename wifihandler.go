package main

import (
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var (
	formTmpl *template.Template
	joinTmpl *template.Template
)

func init() {
	formTmpl = template.Must(template.New("form").ParseFiles("templates/form.html"))
	joinTmpl = template.Must(template.New("join").ParseFiles("templates/join.html"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := formTmpl.ExecuteTemplate(w, "form.html", struct{}{})
	if err != nil {
		log.Fatal("WiFi Handler: ", err)
	}
}

func networkHandler(w http.ResponseWriter, r *http.Request) {
	res, execErr := exec.Command("sh", "network.sh").Output()
	if execErr != nil {
		log.Println(execErr)
	}

	w.Write(res)
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	ssid := r.FormValue("ssid")
	password := r.FormValue("password")
	encryption := "psk2" // TODO: should be possible to change

	go func() {
		cmd := exec.Command("/usr/bin/connect-wifi", ssid, encryption, password)
		cmd.Run()
	}()

	err := joinTmpl.ExecuteTemplate(w, "join.html", struct {
		SSID string
	}{
		ssid,
	})
	if err != nil {
		log.Fatal("WiFi Handler: ", err)
	}
}

func main() {
	http.HandleFunc("/index", formHandler)
	http.HandleFunc("/join", joinHandler)

	http.HandleFunc("/network", networkHandler)

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("/templates"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

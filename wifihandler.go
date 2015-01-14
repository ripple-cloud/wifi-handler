package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

var (
	formTmpl *template.Template
	joinTmpl *template.Template
)

func init() {
	formBindata, err := Asset("templates/form.html")
	if err != nil {
		log.Fatal(err)
	}

	joinBindata, err := Asset("templates/join.html")
	if err != nil {
		log.Fatal(err)
	}

	formTmpl = template.Must(template.New("form").Parse(string(formBindata)))
	joinTmpl = template.Must(template.New("join").Parse(string(joinBindata)))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	bootstrapJS, bsErr := Asset("public/js/bootstrap.min.js")
	if bsErr != nil {
		log.Fatal(bsErr)
	}
	formJS, formErr := Asset("public/js/form.js")
	if formErr != nil {
		log.Fatal(formErr)
	}
	jqueryJS, jErr := Asset("public/js/jquery.min.js")
	if jErr != nil {
		log.Fatal(jErr)
	}
	typeaheadJS, tErr := Asset("public/js/typeahead.bundle.min.js")
	if tErr != nil {
		log.Fatal(tErr)
	}
	bootstrapCSS, bErr := Asset("public/stylesheets/bootstrap.min.css")
	if bErr != nil {
		log.Fatal(bErr)
	}
	formCSS, fErr := Asset("public/stylesheets/form.css")
	if fErr != nil {
		log.Fatal(fErr)
	}

	err := formTmpl.Execute(w, struct {
		SafeBootstrapJS  template.JS
		SafeFormJS       template.JS
		SafeJqueryJS     template.JS
		SafeTypeaheadJS  template.JS
		SafeBootstrapCSS template.CSS
		SafeFormCSS      template.CSS
	}{
		template.JS(bootstrapJS),
		template.JS(formJS),
		template.JS(jqueryJS),
		template.JS(typeaheadJS),
		template.CSS(bootstrapCSS),
		template.CSS(formCSS),
	})
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
	security := r.FormValue("security")

	var encryption string

	switch security {
	case "None":
		encryption = "none"
	case "WEP":
		encryption = "wep"
	case "WPA":
		encryption = "psk"
	default:
		encryption = "psk2"
	}

	go func() {
		cmd := exec.Command("/usr/bin/connect-wifi", ssid, encryption, password)
		cmd.Run()
	}()

	err := joinTmpl.Execute(w, struct {
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

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

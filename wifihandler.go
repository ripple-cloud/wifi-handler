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

const formTmplContents = `
<html>
		<body>
			<h1>Join WiFi Network</h1>
			<form method="post" action="/join" id="loginForm">
				<label for="ssid">Network Name</label>
				<input type="text" id="ssid" name="ssid"/>
				<label for="password">Password</label>
				<input type="password" id="password" name="password"/>
					
				<button type="submit">Join</button>
			</form>
		</body>
</html>
`

const joinTmplContents = `
<html>
		<body>
			<h1>Connecting to <b>{{.SSID}}</b>...</h1>
			<p>You can now switch back to <b>{{.SSID}}</b> network. To find Pi's IP address go to your router configuration page and check DHCP leases.</p>
			<p>Once you know the IP address of Pi, you can access it by running <b>telnet 192.168.X.X</b>.</p>
		</body>
</html>
`

func init() {
	formTmpl = template.Must(template.New("form").Parse(formTmplContents))
	joinTmpl = template.Must(template.New("join").Parse(joinTmplContents))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := formTmpl.Execute(w, struct{}{})
	if err != nil {
		log.Fatal("WiFi Handler: ", err)
	}
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	ssid := r.FormValue("ssid")
	password := r.FormValue("password")
	encryption := "psk2" // TODO: should be possible to change

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
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/join", joinHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

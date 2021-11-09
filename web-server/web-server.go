package webserver

import (
	"fmt"
	"net/http"
)

func WebServer() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)

}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to web")
}

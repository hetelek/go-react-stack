package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func gethostname(request *http.Request) []byte {
	hostname, err := os.Hostname()
	if err != nil {
		return []byte("error getting hostname")
	}
	return []byte(hostname)
}

// StartServer starts the server at localhost:PORT, where PORT is an env variable, or defaulted to 8000
func StartServer(staticDirectory string) {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	if staticDirectory != "" {
		fs := http.FileServer(http.Dir(staticDirectory))
		http.Handle("/", http.StripPrefix("/", fs))
	}

	RegisterHandler("/api/hostname", HTTPGET, gethostname)

	fmt.Println("Serving at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

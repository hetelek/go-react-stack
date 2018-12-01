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
func StartServer(staticFilesDirectory string, port int) {
	if staticFilesDirectory != "" {
		fs := http.FileServer(http.Dir(staticFilesDirectory))
		http.Handle("/", http.StripPrefix("/", fs))
	}

	RegisterHandler("/api/hostname", HTTPGET, gethostname)

	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("Starting server at http://localhost" + portStr)
	log.Fatal(http.ListenAndServe(portStr, nil))
}

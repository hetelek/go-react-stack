package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"./backend"
)

var frontendDirectory = "./frontend"
var backendDirectory = "./backend"
var staticFilesDirectory = "./frontend/build"

func startFrontend() {
	for {
		fmt.Println("Starting frontend")
		cmd := exec.Command("npm", "start")
		cmd.Dir = frontendDirectory
		out, _ := cmd.CombinedOutput()
		log.Println(string(out))
	}
}

func main() {
	var development bool
	var port int

	flag.BoolVar(&development, "dev", false, "If set, runs in development mode")
	flag.IntVar(&port, "port", 8000, "Port to run the backend on, defaults to 8000")
	flag.Parse()

	if development {
		fmt.Println("[ RUNNING IN DEVELOPMENT MODE ]")
		fmt.Println()

		go startFrontend()
		server.StartServer("", port)
	} else {
		server.StartServer(staticFilesDirectory, port)
	}
}

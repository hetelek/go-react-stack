package main

import (
	"flag"
	"fmt"
	"os/exec"

	"./backend"
)

var frontendDirectory = "./frontend"
var staticFiles = "./frontend/build"

func startFrontend() {
	for {
		cmd := exec.Command("npm", "start")
		cmd.Dir = frontendDirectory
		cmd.Run()
	}
}

func main() {
	development := flag.Bool("dev", false, "If set, runs the server in development mode")
	flag.Parse()

	if *development {
		fmt.Println("[ RUNNING IN DEVELOPMENT MODE ]")
		fmt.Println()
		go startFrontend()

		server.StartServer("")
	} else {
		server.StartServer(staticFiles)
	}
}

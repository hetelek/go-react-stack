package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"./backend"
	"github.com/radovskyb/watcher"
)

var dirsToWatch = []string{"./frontend/public", "./frontend/src"}
var buildWorkingDirectory = "./frontend"
var buildDirectory = "./frontend/build"

func build(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	cmd := exec.Command("npm", "run-script", "build")
	cmd.Dir = buildWorkingDirectory
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatalf("Rebuild failed with %s\n", err)
	}
}

func startWatching() {
	w := watcher.New()

	w.SetMaxEvents(1)
	go func() {
		mutex := sync.Mutex{}
		for {
			select {
			case <-w.Event:
				fmt.Println("[DEV] Change detected, re-building...")
				build(&mutex)
				fmt.Println("[DEV] Done building.")
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	fmt.Printf("[DEV] Watching: %v", dirsToWatch)
	fmt.Println()
	for _, dir := range dirsToWatch {
		if err := w.AddRecursive(dir); err != nil {
			log.Fatalln(err)
		}
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	development := flag.Bool("dev", false, "If set, runs the server in development mode")
	flag.Parse()
	if *development {
		fmt.Println("[ RUNNING IN DEVELOPMENT MODE ]")
		fmt.Println()
		go startWatching()
	}

	server.StartServer(buildDirectory)
}

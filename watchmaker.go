package main

import (
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

/* would be nice to be able to use gitignore or similar */
var ignoredRegexp = regexp.MustCompile(`(^|/)(\.[^/]|4913$)|~$|\.sw[px]+$|\.lock$|\.log`)

func ignored(path string) bool {
	return ignoredRegexp.MatchString(path)
}

func Watch(watcher *fsnotify.Watcher, ch chan struct{}) {
	for {
		select {
		case event := <-watcher.Events:
			if !ignored(event.Name) {
				ch <- struct{}{}
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func Build() {
	cmd := exec.Command("make")
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
	cmd.Run()
}

func Wait(in chan struct{}) {
	waitch := make(<-chan time.Time)
	for {
		select {
		case <-in:
			waitch = time.After(100 * time.Millisecond)
		case <-waitch:
			Build()
		}
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	eventch := make(chan struct{})

	go Watch(watcher, eventch)
	go Wait(eventch)

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	<-done
}

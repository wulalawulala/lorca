package main

import (
	"log"
	"net/url"
	"runtime"

	"github.com/zserge/lorca"
)

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	args = append(args, "--kiosk")

	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 800, 600, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

package main

import (
	"surlac.org/vest/internal/httpserver"
)

func main() {
	run()
}

func run() {
	srv := httpserver.HTTPServer{}
	srv.Start()
}

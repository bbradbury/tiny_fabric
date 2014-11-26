package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome too the home page, %s!", c.URLParams["name"])
}

func createService(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating service %s...", c.URLParams["name"])
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Get("/services", func() {})
	goji.Post("/services/:service_name", createService)
	goji.Post("/services/:service_name/instances/:instance_id", func() {})
	goji.Post("/services/:service_name/instances/:instance_id/logs", func() {})
	goji.Post("/services/:service_name/instances/:instance_id/logs/:log_id", func() {})
	goji.Serve()
}

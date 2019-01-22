package main

import (
	"fmt"
	"net/http"

	"github.com/ortense/godesafio/core"
)

func main() {
	app := core.CreateApp()

	app.Router(
		core.Route{
			Path:   "/",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{ "message": "It's works!" }`)
			}},
		core.Route{
			Path:   "/resource/:id",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{ "message": "GET /resource/:id match with `+r.URL.Path+`" }`)
			}},
		core.Route{
			Path:   "/resource/:id",
			Method: "POST",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{ "message": "POST /resource/:id match with `+r.URL.Path+`" }`)
			}})

	app.Start(":3000")
}

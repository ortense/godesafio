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
			}})

	app.Start(":3000")
}

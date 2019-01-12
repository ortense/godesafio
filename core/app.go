package core

import (
	"fmt"
	"log"
	"net/http"
)

// HandleFunc represent a http handler function
type HandleFunc func(w http.ResponseWriter, r *http.Request)

// Route a struct that represent a http route
type Route struct {
	Method  string
	Path    string
	Handler HandleFunc
}

// App a abstraction of http server
type App struct {
	routes map[string]map[string]HandleFunc
}

func (app *App) addRoute(route Route) {

	if app.routes[route.Path] == nil {
		app.routes[route.Path] = make(map[string]HandleFunc)
	}

	app.routes[route.Path][route.Method] = route.Handler
}

// Router adds variadic number of routes to app
func (app *App) Router(routes ...Route) {
	for _, route := range routes {
		app.addRoute(route)
	}
}

// Start initialize http server
func (app *App) Start(address string) {

	handleNotFound := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{ "mesage": "%s" }`, "Not found "+r.Method+" "+r.URL.Path)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.URL.Path)

		if app.routes[r.URL.Path] == nil || app.routes[r.URL.Path][r.Method] == nil {
			handleNotFound(w, r)
			return
		}

		app.routes[r.URL.Path][r.Method](w, r)
		return
	})

	fmt.Println("Start server at " + address)
	http.ListenAndServe(address, nil)
}

// CreateApp is a http app factory
func CreateApp() App {
	app := App{make(map[string]map[string]HandleFunc)}
	return app
}

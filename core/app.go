package core

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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
	routes   map[string]map[string]HandleFunc
	patterns map[string]map[string]HandleFunc
}

func (app *App) addRoute(route Route) {
	if app.routes[route.Path] == nil {
		app.routes[route.Path] = make(map[string]HandleFunc)
	}

	app.routes[route.Path][route.Method] = route.Handler
}

func isRoutePattern(path string) bool {
	rgx := regexp.MustCompile(".*\\/:\\w+.*")
	return rgx.MatchString(path)
}

func createRoutePattern(path string) string {
	wildcard := "([^/]+)"
	matchParams := regexp.MustCompile(":" + wildcard)
	rgx := strings.Replace(matchParams.ReplaceAllLiteralString(path, wildcard), "/", "\\/", -1)
	return "^" + rgx + "$"
}

func (app *App) addPattern(route Route) {
	pattern := createRoutePattern(route.Path)

	if app.patterns[pattern] == nil {
		app.patterns[pattern] = make(map[string]HandleFunc)
	}

	app.patterns[pattern][route.Method] = route.Handler
}

// Router adds variadic number of routes to app
func (app *App) Router(routes ...Route) {
	for _, route := range routes {
		if isRoutePattern(route.Path) {
			app.addPattern(route)
		} else {
			app.addRoute(route)
		}
	}
}

// HandlerFunc is a HTTP handler function to be attached to http.HandleFunc
func (app *App) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	if app.routes[r.URL.Path] != nil || app.routes[r.URL.Path][r.Method] != nil {
		app.routes[r.URL.Path][r.Method](w, r)
		return
	}

	for pattern, handlers := range app.patterns {
		rgx := regexp.MustCompile(pattern)
		if rgx.MatchString(r.URL.Path) && handlers[r.Method] != nil {
			handlers[r.Method](w, r)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{ "mesage": "%s" }`, "Not found "+r.Method+" "+r.URL.Path)
	return
}

// Start initialize http server
func (app *App) Start(address string) {
	http.HandleFunc("/", app.HandlerFunc)
	fmt.Println("Start server at " + address)
	http.ListenAndServe(address, nil)
}

// CreateApp is a http app factory
func CreateApp() App {
	return App{
		routes:   make(map[string]map[string]HandleFunc),
		patterns: make(map[string]map[string]HandleFunc)}
}

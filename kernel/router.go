package kernel

import (
	"log"
	"net/http"
	"regexp"
)

type Router struct {
	Routes []*Route
}

func NewRouter() *Router {
	return &Router{
		Routes: nil,
	}
}

func (o *Router) AddRoute(route *Route) {
	o.Routes = append(o.Routes, route)
}

func (o *Router) Handler(pattern *regexp.Regexp, handler http.Handler) {
	o.AddRoute(NewRoute(pattern, handler))
}

func (o *Router) HandleFunc(pattern *regexp.Regexp, fn func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(fn)
	route := NewRoute(pattern, handler)
	o.AddRoute(route)
}

func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	for _, route := range o.Routes {
		if route.pattern.MatchString(url) {
			log.Println("Route matched:", url, route.pattern)
			route.handler.ServeHTTP(w, r)
			return
		} else {
			log.Println("Route not matched:", url, route.pattern)
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

package nlike

import (
	"log"
	"net/http"
	"regexp"
)

type Router struct {
	Routes []*Route
}

func NewRoute(pattern *regexp.Regexp, handler http.Handler) *Route {
	return &Route{
		pattern: pattern,
		handler: handler,
	}
}

func NewRouter() *Router {
	return &Router{
		Routes: nil,
	}
}

func (o *Router) Handler(pattern *regexp.Regexp, handler http.Handler) {
	o.Routes = append(o.Routes, NewRoute(pattern, handler))
}

func (o *Router) HandleFunc(pattern *regexp.Regexp, fn func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(fn)
	route := NewRoute(pattern, handler)
	o.Routes = append(o.Routes, route)
}

func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	for _, route := range o.Routes {
		if route.pattern.MatchString(url) {
			route.handler.ServeHTTP(w, r)
			return
		} else {
			log.Println("Route not matched:", url, route.pattern)
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

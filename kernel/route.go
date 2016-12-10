package kernel

import (
	"net/http"
	"regexp"
)

type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
	child   []*Route
}

func NewRoute(pattern *regexp.Regexp, handler http.Handler) *Route {
	return &Route{
		pattern: pattern,
		handler: handler,
	}
}

func (o *Route) getPattern() *regexp.Regexp {
	return o.pattern
}

func (o *Route) getHandler() http.Handler {
	return o.handler
}

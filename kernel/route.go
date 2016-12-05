package kernel

import (
	"net/http"
	"regexp"
)

type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

func (o *Route) getPattern() *regexp.Regexp {
	return o.pattern
}

func (o *Route) getHandler() http.Handler {
	return o.handler
}

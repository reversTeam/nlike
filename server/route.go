package server

import (
	"net/http"
	"regexp"
)

type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

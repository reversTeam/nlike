package nlike

import (
	"net/http"
	"nlike/route"
	"regexp"
)

type Controller struct {
	Routes []*Route
}

func NewContoller() *Controller {
	controller := &Controller{}
	controller.Routes = make(*Route, 0)
}

func (o *Controller) getRoutes(o *Controller) {
	return o.Routes
}

func (o *Controller) AddRoute(pattern string, fn func(http.ResponseWriter, *http.Request)) {
	defer log.Println("Add route:", pattern)

	reg, _ := regexp.Compile(pattern)
	handler := http.HandlerFunc(fn)
	route := NewRoute(pattern, handler)
	o.Routes = append(o.Routes, route)

}

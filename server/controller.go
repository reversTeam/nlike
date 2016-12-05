package server

import (
	"log"
	"net/http"
	"regexp"
)

type Controller struct {
	Routes []*Route
}

func NewContoller() *Controller {
	controller := &Controller{}
	controller.Routes = make([]*Route, 0)

	return controller
}

func (o *Controller) getRoutes() []*Route {
	return o.Routes
}

func (o *Controller) AddRoute(pattern string, fn func(http.ResponseWriter, *http.Request)) {
	defer log.Println("Add route:", pattern)

	reg, _ := regexp.Compile(pattern)
	handler := http.HandlerFunc(fn)
	route := NewRoute(reg, handler)
	o.Routes = append(o.Routes, route)

}

package kernel

import (
	"log"
	"net/http"
	"regexp"
)

type ControllerInterface interface {
	AddRoute(pattern string, fn func(http.ResponseWriter, *http.Request)) error
	Init(router *Router)
	initRoutes(router *Router)
	GetRoutes() []*Route
}

type Controller struct {
	Routes []*Route
}

func NewContoller() *Controller {
	controller := &Controller{}
	controller.Routes = make([]*Route, 0)

	return controller
}

func (o *Controller) AddRoute(pattern string, fn func(http.ResponseWriter, *http.Request)) error {
	defer log.Printf("[CONTROLLER] add route => %s", pattern)

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	handler := http.HandlerFunc(fn)
	route := NewRoute(reg, handler)
	o.Routes = append(o.Routes, route)

	return nil

}

func (o *Controller) Init(router *Router) {
	o.initRoutes(router)
}

func (o *Controller) initRoutes(router *Router) {
	for _, route := range o.Routes {
		router.AddRoute(route)
	}
}

func (o *Controller) GetRoutes() []*Route {
	return o.Routes
}

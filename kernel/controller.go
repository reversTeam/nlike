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
	Stop()
}

type Controller struct {
	name   string
	Routes []*Route
}

func NewController(name string) *Controller {
	controller := &Controller{
		name: name,
	}
	controller.Routes = make([]*Route, 0)

	return controller
}

func (o *Controller) AddRoute(pattern string, fn func(http.ResponseWriter, *http.Request)) error {
	defer log.Printf("[%s]: Added route - %s", o.name, pattern)

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

func (o *Controller) Stop() {
}

func (o *Controller) initRoutes(router *Router) {
	for _, route := range o.Routes {
		router.AddRoute(route)
	}
}

func (o *Controller) GetRoutes() []*Route {
	return o.Routes
}

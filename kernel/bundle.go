package kernel

import (
	"log"
)

type BundleInterface interface {
	Init(router *Router)
	AddController(controller ControllerInterface)
	BootstrapEvent()
	Stop()
	stopControllers()
	stopGrpcs()
}

type Bundle struct {
	name        string
	controllers []ControllerInterface
	grpcs       []GrpcInterface
}

func NewBundle(name string) *Bundle {
	bundle := &Bundle{
		name:        name,
		controllers: make([]ControllerInterface, 0),
		grpcs:       make([]GrpcInterface, 0),
	}
	return bundle
}

func (o *Bundle) Init(router *Router) {
	log.Printf("[%s]: Initialization", o.name)
	o.initControllers(router)
	o.initGrpcs()
}

func (o *Bundle) Stop() {
	defer log.Printf("[%s]: Stoped", o.name)
	o.stopControllers()
	o.stopGrpcs()
}

func (o *Bundle) stopControllers() {
	for _, controller := range o.controllers {
		controller.Stop()
	}
}

func (o *Bundle) stopGrpcs() {
	for _, grpc := range o.grpcs {
		grpc.Stop()
	}
}

func (o *Bundle) initControllers(router *Router) {
	for _, controller := range o.controllers {
		controller.Init(router)
	}
}

func (o *Bundle) initGrpcs() {
	for _, grpc := range o.grpcs {
		grpc.Init()
	}
}

func (o *Bundle) AddController(controller ControllerInterface) {
	log.Printf("[%s]: Add controller", o.name)
	o.controllers = append(o.controllers, controller)
}

func (o *Bundle) AddGrpc(grpc GrpcInterface) {
	o.grpcs = append(o.grpcs, grpc)
}

func (o *Bundle) BootstrapEvent() {

}

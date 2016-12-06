package kernel

import "log"

type BundleInterface interface {
	Init(router *Router)
	AddController(controller ControllerInterface)
	GetName() string
	BootstrapEvent()
}

type Bundle struct {
	name        string
	controllers []ControllerInterface
}

func NewBundle(name string) *Bundle {
	bundle := &Bundle{
		name:        name,
		controllers: make([]ControllerInterface, 0),
	}
	return bundle
}

func (o *Bundle) Init(router *Router) {
	defer log.Printf("[BUNDLE]: %s initialized", o.name)
	log.Printf("[BUNDLE]: %s initialization", o.name)
	o.initControllers(router)
}

func (o *Bundle) initControllers(router *Router) {
	for _, controller := range o.controllers {
		controller.Init(router)
	}
}

func (o *Bundle) AddController(controller ControllerInterface) {
	o.controllers = append(o.controllers, controller)
}

func (o *Bundle) BootstrapEvent() {
	defer log.Printf("[BUNDLE] %s initialized Bootstrap Event", o.name)
}

func (o *Bundle) GetName() string {
	return o.name
}

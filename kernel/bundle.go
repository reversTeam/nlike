package kernel

type Bundle struct {
	name        string
	controllers []ControllerInterface
}

func NewBundle(name string) *Bundle {
	return &Bundle{
		name:        name,
		controllers: make([]ControllerInterface, 0),
	}
}

func (o *Bundle) Init(router *Router) {
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

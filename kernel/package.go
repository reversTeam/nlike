package kernel

import (
	"log"
)

type PackageInterface interface {
	Init(router *Router)
	initBundles(router *Router)
	AddBundle(bundle BundleInterface)
	BootstrapEvent()
}

type Package struct {
	name    string
	bundles []BundleInterface
}

func NewPackage(name string) *Package {
	return &Package{
		name:    name,
		bundles: make([]BundleInterface, 0),
	}
}

func (o *Package) BootstrapEvent() {
	defer log.Printf("[PACKAGE]: %s Bootstraped", o.name)
}

func (o *Package) Init(router *Router) {
	o.initBundles(router)
}

func (o *Package) initBundles(router *Router) {
	for _, bundle := range o.bundles {
		bundle.Init(router)
	}
}

func (o *Package) AddBundle(bundle BundleInterface) {
	defer log.Printf("[PACKAGE]: %s Package add bundle %s", o.name, bundle.GetName())
	o.bundles = append(o.bundles, bundle)
}

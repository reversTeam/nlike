package kernel

import (
	"log"
)

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

func (o *Package) Init(router *Router) {
	defer log.Println("[PACKAGE]: %s initialized", o.name)
	log.Printf("[PACKAGE]: %s initialization", o.name)

	o.initBundles(router)
}

func (o *Package) initBundles(router *Router) {
	for _, bundle := range o.bundles {
		bundle.Init(router)
	}
}

func (o *Package) AddBundle(bundle BundleInterface) {
	defer log.Printf("[%s] Package add bundle %s", o.name, bundle.GetName())
	o.bundles = append(o.bundles, bundle)
}

package kernel

type Package struct {
	name    string
	bundles []*Bundle
}

func NewPackage(name string) *Package {
	return &Package{
		name:    name,
		bundles: make([]*Bundle, 0),
	}
}

func (o *Package) Init(router *Router) {
	o.initBundles(router)
}

func (o *Package) initBundles(router *Router) {
	for _, bundle := range o.bundles {
		bundle.Init(router)
	}
}

func (o *Package) AddBundle(bundle *Bundle) {
	o.bundles = append(o.bundles, bundle)
}

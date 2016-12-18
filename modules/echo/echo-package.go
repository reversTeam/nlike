package echo

import (
	"github.com/reversTeam/nlike/kernel"
	"github.com/reversTeam/nlike/modules/echo/bundles/echo"
)

type EchoPackage struct {
	kernel.Package
}

func NewPackage() *EchoPackage {
	return &EchoPackage{
		*kernel.NewPackage("EchoPackage"),
	}
}

func (o *EchoPackage) BootstrapEvent() {
	o.Package.BootstrapEvent()
	bundle := topic.NewBundle()
	bundle.BootstrapEvent()
	o.AddBundle(bundle)
}

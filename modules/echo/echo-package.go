package echo

import (
	nlike "github.com/reversTeam/nlike/kernel"
	echoBundle "github.com/reversTeam/nlike/modules/echo/bundles/echo"
)

type EchoPackage struct {
	nlike.Package
}

func NewPackage() *EchoPackage {
	return &EchoPackage{
		*nlike.NewPackage("EchoPackage"),
	}
}

func (o *EchoPackage) BootstrapEvent() {
	o.Package.BootstrapEvent()
	bundle := echoBundle.NewBundle()
	bundle.BootstrapEvent()
	o.AddBundle(bundle)
}

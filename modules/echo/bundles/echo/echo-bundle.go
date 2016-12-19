package topic

import (
	"github.com/reversTeam/nlike/kernel"
	echoController "github.com/reversTeam/nlike/modules/echo/bundles/echo/controller"
	echoGrpc "github.com/reversTeam/nlike/modules/echo/bundles/echo/grpc"
)

type EchoBundle struct {
	kernel.Bundle
}

func NewBundle() *EchoBundle {
	bundle := &EchoBundle{
		*kernel.NewBundle("EchoBundle"),
	}

	return bundle
}

func (o *EchoBundle) BootstrapEvent() {
	o.Bundle.BootstrapEvent()
	o.AddControllers()
	o.AddGrpcs()
}

func (o *EchoBundle) AddControllers() {
	o.AddController(echoController.NewController())
}

func (o *EchoBundle) AddGrpcs() {
	grpc := echoGrpc.NewGrpc()
	grpc.InitRequests()
	o.AddGrpc(grpc)
}

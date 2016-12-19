package topic

import (
	"github.com/reversTeam/nlike/kernel"
	topicController "github.com/reversTeam/nlike/modules/topic/bundles/topic/controller"
	topicGrpc "github.com/reversTeam/nlike/modules/topic/bundles/topic/grpc"
)

type TopicBundle struct {
	kernel.Bundle
}

func NewBundle() *TopicBundle {
	bundle := &TopicBundle{
		*kernel.NewBundle("TopicBundle"),
	}

	return bundle
}

func (o *TopicBundle) BootstrapEvent() {
	o.Bundle.BootstrapEvent()
	o.AddControllers()
	o.AddGrpcs()
}

func (o *TopicBundle) AddControllers() {
	o.AddController(topicController.NewController())
}

func (o *TopicBundle) AddGrpcs() {
	grpc := topicGrpc.NewGrpc()
	grpc.InitRequests()
	o.AddGrpc(grpc)
}

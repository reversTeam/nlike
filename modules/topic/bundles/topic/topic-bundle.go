package topic

import (
	"github.com/reversTeam/nlike/kernel"
	topicController "github.com/reversTeam/nlike/modules/topic/bundles/topic/controller"
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
	o.AddController(topicController.NewController())
}

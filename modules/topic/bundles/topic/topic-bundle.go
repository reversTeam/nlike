package topic

import (
	"github.com/reversTeam/nlike/kernel"
	topicController "github.com/reversTeam/nlike/modules/topic/bundles/topic/controller"
	"log"
)

type TopicBundle struct {
	*kernel.Bundle
}

func (o *TopicBundle) NewBundle() kernel.BundleInterface {
	bundle := &TopicBundle{
		kernel.NewBundle("TopicBundle"),
	}

	return bundle
}

func (o *TopicBundle) BootstrapEvent() {
	log.Println("###################################################################")
	o.AddController(topicController.NewController())
}

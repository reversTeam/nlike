package topic

import (
	"github.com/reversTeam/nlike/kernel"
	"github.com/reversTeam/nlike/modules/topic/bundles/topic"
)

type TopicPackage struct {
	kernel.Package
}

func NewPackage() *TopicPackage {
	return &TopicPackage{
		*kernel.NewPackage("TopicPackage"),
	}
}

func (o *TopicPackage) BootstrapEvent() {
	o.Package.BootstrapEvent()
	bundle := topic.NewBundle()
	bundle.BootstrapEvent()
	o.AddBundle(bundle)
}

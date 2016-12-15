package modules

import (
	"github.com/reversTeam/nlike/kernel"
	"github.com/reversTeam/nlike/modules/topic"
)

func Init(server *kernel.Server) {
	server.AddPackage(topic.NewPackage())
}

package request

import (
	"github.com/reversTeam/nlike/kernel"
	echoProto "github.com/reversTeam/nlike/modules/echo/bundles/echo/proto/build"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoRequest struct {
	kernel.grpc.Request
}

func NewRequest() *EchoRequest {
	return &EchoRequest{
		*kernel.grpc.NewRequest(),
	}
}

func (o *EchoRequest) Echo(ctx context.Context, echoMessage *echoProto.EchoMessage) (*echoProto.EchoMessage, error) {
	return echoMessage, nil
}

func (o *EchoRequest) InitServices(s *grpc.Server) {
	o.Request.InitServices(s)
	// TODO : Can't register this server because segfault
	echoProto.RegisterEchoServiceServer(s, o)
}

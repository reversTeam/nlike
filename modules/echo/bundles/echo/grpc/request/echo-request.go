package request

import (
	nlikeGrpc "github.com/reversTeam/nlike/kernel/grpc"
	echoProto "github.com/reversTeam/nlike/modules/echo/bundles/echo/proto/build"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoRequest struct {
	nlikeGrpc.Request
}

func NewRequest() *EchoRequest {
	return &EchoRequest{
		*nlikeGrpc.NewRequest(),
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

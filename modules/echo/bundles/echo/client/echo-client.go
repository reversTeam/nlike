package client

import (
	echoProto "github.com/reversTeam/nlike/modules/echo/bundles/echo/proto/build"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	conn echoProto.EchoServiceClient
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("163.172.32.48:4244", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Client{echoProto.NewEchoServiceClient(conn)}, nil
}

func (o *Client) Echo(msg string) {
	ctx := context.Background()
	protoMsg := &echoProto.EchoMessage{msg}
	_, err := o.conn.Echo(ctx, protoMsg)
	if err != nil {
		log.Fatal(err)
	}
	//	log.Println(res.GetValue())
}

package client

import (
	"fmt"
	echoProto "github.com/reversTeam/nlike/modules/echo/bundles/echo/proto/build"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	conn echoProto.EchoServiceClient
	host string
	port int
}

func NewClient(host string, port int) (*Client, error) {
	tmp := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(tmp, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Client{
		echoProto.NewEchoServiceClient(conn),
		host,
		port,
	}, nil
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

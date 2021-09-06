package pb

import (
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	Service UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

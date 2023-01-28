package cacher

import (
	"time"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/patrickmn/go-cache"
)

type Client struct {
	client ec2iface.EC2API
	cache  *cache.Cache
}

func NewClient(client ec2iface.EC2API) *Client {
	return &Client{
		client: client,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}

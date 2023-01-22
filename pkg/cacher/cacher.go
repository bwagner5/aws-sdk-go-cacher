package cacher

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/mitchellh/hashstructure/v2"
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

func (c *Client) DescribeInstancesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	return c.client.DescribeInstancesWithContext(ctx, input, opts...)
}

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudHSMV2 struct {
	cloudhsmv2iface.CloudHSMV2API
	cache *cache.Cache
}

func NewCloudHSMV2(cloudhsmv2api cloudhsmv2iface.CloudHSMV2API) *CloudHSMV2 {
	return &CloudHSMV2{
		CloudHSMV2API: cloudhsmv2api,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudHSMV2) DescribeBackups(input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.DescribeBackupsOutput), nil
	}
	out, err := c.CloudHSMV2API.DescribeBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSMV2) DescribeBackupsPages(input *cloudhsmv2.DescribeBackupsInput, fn func(*cloudhsmv2.DescribeBackupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.DescribeBackupsOutput{}
	fnCacher := func(out *cloudhsmv2.DescribeBackupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.DescribeBackupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) DescribeBackupsPagesWithContext(ctx context.Context, input *cloudhsmv2.DescribeBackupsInput, fn func(*cloudhsmv2.DescribeBackupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.DescribeBackupsOutput{}
	fnCacher := func(out *cloudhsmv2.DescribeBackupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.DescribeBackupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) DescribeBackupsWithContext(ctx context.Context, input *cloudhsmv2.DescribeBackupsInput, opts ...request.Option) (*cloudhsmv2.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.DescribeBackupsOutput), nil
	}
	out, err := c.CloudHSMV2API.DescribeBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSMV2) DescribeClusters(input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.DescribeClustersOutput), nil
	}
	out, err := c.CloudHSMV2API.DescribeClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSMV2) DescribeClustersPages(input *cloudhsmv2.DescribeClustersInput, fn func(*cloudhsmv2.DescribeClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.DescribeClustersOutput{}
	fnCacher := func(out *cloudhsmv2.DescribeClustersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.DescribeClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) DescribeClustersPagesWithContext(ctx context.Context, input *cloudhsmv2.DescribeClustersInput, fn func(*cloudhsmv2.DescribeClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.DescribeClustersOutput{}
	fnCacher := func(out *cloudhsmv2.DescribeClustersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.DescribeClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) DescribeClustersWithContext(ctx context.Context, input *cloudhsmv2.DescribeClustersInput, opts ...request.Option) (*cloudhsmv2.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.DescribeClustersOutput), nil
	}
	out, err := c.CloudHSMV2API.DescribeClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSMV2) ListTags(input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.ListTagsOutput), nil
	}
	out, err := c.CloudHSMV2API.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSMV2) ListTagsPages(input *cloudhsmv2.ListTagsInput, fn func(*cloudhsmv2.ListTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.ListTagsOutput{}
	fnCacher := func(out *cloudhsmv2.ListTagsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.ListTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) ListTagsPagesWithContext(ctx context.Context, input *cloudhsmv2.ListTagsInput, fn func(*cloudhsmv2.ListTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudhsmv2.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudhsmv2.ListTagsOutput{}
	fnCacher := func(out *cloudhsmv2.ListTagsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.CloudHSMV2API.ListTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudHSMV2) ListTagsWithContext(ctx context.Context, input *cloudhsmv2.ListTagsInput, opts ...request.Option) (*cloudhsmv2.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsmv2.ListTagsOutput), nil
	}
	out, err := c.CloudHSMV2API.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package cloud9cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloud9/cloud9iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Cloud9 struct {
	cloud9iface.Cloud9API
	cache *cache.Cache
}

func New(cloud9api cloud9iface.Cloud9API) *Cloud9 {
	return &Cloud9{
		Cloud9API: cloud9api,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Cloud9) DescribeEnvironmentMemberships(input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentMembershipsOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironmentMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) DescribeEnvironmentMembershipsPages(input *cloud9.DescribeEnvironmentMembershipsInput, fn func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloud9.DescribeEnvironmentMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &cloud9.DescribeEnvironmentMembershipsOutput{}
	fnCacher := func(out *cloud9.DescribeEnvironmentMembershipsOutput, more bool) bool {
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
	if err := c.Cloud9API.DescribeEnvironmentMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Cloud9) DescribeEnvironmentMembershipsPagesWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, fn func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloud9.DescribeEnvironmentMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &cloud9.DescribeEnvironmentMembershipsOutput{}
	fnCacher := func(out *cloud9.DescribeEnvironmentMembershipsOutput, more bool) bool {
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
	if err := c.Cloud9API.DescribeEnvironmentMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Cloud9) DescribeEnvironmentMembershipsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentMembershipsOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironmentMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) DescribeEnvironmentStatus(input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentStatusOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironmentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) DescribeEnvironmentStatusWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentStatusInput, opts ...request.Option) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentStatusOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironmentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) DescribeEnvironments(input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentsOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) DescribeEnvironmentsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.DescribeEnvironmentsOutput), nil
	}
	out, err := c.Cloud9API.DescribeEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) ListEnvironments(input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.ListEnvironmentsOutput), nil
	}
	out, err := c.Cloud9API.ListEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) ListEnvironmentsPages(input *cloud9.ListEnvironmentsInput, fn func(*cloud9.ListEnvironmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloud9.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloud9.ListEnvironmentsOutput{}
	fnCacher := func(out *cloud9.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.Cloud9API.ListEnvironmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Cloud9) ListEnvironmentsPagesWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, fn func(*cloud9.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloud9.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloud9.ListEnvironmentsOutput{}
	fnCacher := func(out *cloud9.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.Cloud9API.ListEnvironmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Cloud9) ListEnvironmentsWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, opts ...request.Option) (*cloud9.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.ListEnvironmentsOutput), nil
	}
	out, err := c.Cloud9API.ListEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) ListTagsForResource(input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.ListTagsForResourceOutput), nil
	}
	out, err := c.Cloud9API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Cloud9) ListTagsForResourceWithContext(ctx context.Context, input *cloud9.ListTagsForResourceInput, opts ...request.Option) (*cloud9.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloud9.ListTagsForResourceOutput), nil
	}
	out, err := c.Cloud9API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package codestarconnectionscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeStarConnections struct {
	codestarconnectionsiface.CodeStarConnectionsAPI
	cache *cache.Cache
}

func New(codestarconnectionsapi codestarconnectionsiface.CodeStarConnectionsAPI) *CodeStarConnections {
	return &CodeStarConnections{
		CodeStarConnectionsAPI: codestarconnectionsapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeStarConnections) GetConnection(input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.GetConnectionOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.GetConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) GetConnectionWithContext(ctx context.Context, input *codestarconnections.GetConnectionInput, opts ...request.Option) (*codestarconnections.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.GetConnectionOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.GetConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) GetHost(input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.GetHostOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.GetHost(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) GetHostWithContext(ctx context.Context, input *codestarconnections.GetHostInput, opts ...request.Option) (*codestarconnections.GetHostOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.GetHostOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.GetHostWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListConnections(input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListConnectionsOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListConnectionsPages(input *codestarconnections.ListConnectionsInput, fn func(*codestarconnections.ListConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarconnections.ListConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarconnections.ListConnectionsOutput{}
	fnCacher := func(out *codestarconnections.ListConnectionsOutput, more bool) bool {
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
	if err := c.CodeStarConnectionsAPI.ListConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarConnections) ListConnectionsPagesWithContext(ctx context.Context, input *codestarconnections.ListConnectionsInput, fn func(*codestarconnections.ListConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarconnections.ListConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarconnections.ListConnectionsOutput{}
	fnCacher := func(out *codestarconnections.ListConnectionsOutput, more bool) bool {
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
	if err := c.CodeStarConnectionsAPI.ListConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarConnections) ListConnectionsWithContext(ctx context.Context, input *codestarconnections.ListConnectionsInput, opts ...request.Option) (*codestarconnections.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListConnectionsOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListHosts(input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListHostsOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListHostsPages(input *codestarconnections.ListHostsInput, fn func(*codestarconnections.ListHostsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarconnections.ListHostsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarconnections.ListHostsOutput{}
	fnCacher := func(out *codestarconnections.ListHostsOutput, more bool) bool {
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
	if err := c.CodeStarConnectionsAPI.ListHostsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarConnections) ListHostsPagesWithContext(ctx context.Context, input *codestarconnections.ListHostsInput, fn func(*codestarconnections.ListHostsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarconnections.ListHostsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarconnections.ListHostsOutput{}
	fnCacher := func(out *codestarconnections.ListHostsOutput, more bool) bool {
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
	if err := c.CodeStarConnectionsAPI.ListHostsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarConnections) ListHostsWithContext(ctx context.Context, input *codestarconnections.ListHostsInput, opts ...request.Option) (*codestarconnections.ListHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListHostsOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListTagsForResource(input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarConnections) ListTagsForResourceWithContext(ctx context.Context, input *codestarconnections.ListTagsForResourceInput, opts ...request.Option) (*codestarconnections.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarconnections.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeStarConnectionsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

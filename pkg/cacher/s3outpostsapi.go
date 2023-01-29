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
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"github.com/aws/aws-sdk-go/service/s3outposts/s3outpostsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type S3Outposts struct {
	s3outpostsiface.S3OutpostsAPI
	cache *cache.Cache
}

func NewS3Outposts(s3outpostsapi s3outpostsiface.S3OutpostsAPI) *S3Outposts {
	return &S3Outposts{
		S3OutpostsAPI: s3outpostsapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *S3Outposts) ListEndpoints(input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3outposts.ListEndpointsOutput), nil
	}
	out, err := c.S3OutpostsAPI.ListEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Outposts) ListEndpointsPages(input *s3outposts.ListEndpointsInput, fn func(*s3outposts.ListEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3outposts.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3outposts.ListEndpointsOutput{}
	fnCacher := func(out *s3outposts.ListEndpointsOutput, more bool) bool {
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
	if err := c.S3OutpostsAPI.ListEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Outposts) ListEndpointsPagesWithContext(ctx context.Context, input *s3outposts.ListEndpointsInput, fn func(*s3outposts.ListEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3outposts.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3outposts.ListEndpointsOutput{}
	fnCacher := func(out *s3outposts.ListEndpointsOutput, more bool) bool {
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
	if err := c.S3OutpostsAPI.ListEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Outposts) ListEndpointsWithContext(ctx context.Context, input *s3outposts.ListEndpointsInput, opts ...request.Option) (*s3outposts.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3outposts.ListEndpointsOutput), nil
	}
	out, err := c.S3OutpostsAPI.ListEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Outposts) ListSharedEndpoints(input *s3outposts.ListSharedEndpointsInput) (*s3outposts.ListSharedEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3outposts.ListSharedEndpointsOutput), nil
	}
	out, err := c.S3OutpostsAPI.ListSharedEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Outposts) ListSharedEndpointsPages(input *s3outposts.ListSharedEndpointsInput, fn func(*s3outposts.ListSharedEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3outposts.ListSharedEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3outposts.ListSharedEndpointsOutput{}
	fnCacher := func(out *s3outposts.ListSharedEndpointsOutput, more bool) bool {
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
	if err := c.S3OutpostsAPI.ListSharedEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Outposts) ListSharedEndpointsPagesWithContext(ctx context.Context, input *s3outposts.ListSharedEndpointsInput, fn func(*s3outposts.ListSharedEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3outposts.ListSharedEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3outposts.ListSharedEndpointsOutput{}
	fnCacher := func(out *s3outposts.ListSharedEndpointsOutput, more bool) bool {
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
	if err := c.S3OutpostsAPI.ListSharedEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Outposts) ListSharedEndpointsWithContext(ctx context.Context, input *s3outposts.ListSharedEndpointsInput, opts ...request.Option) (*s3outposts.ListSharedEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3outposts.ListSharedEndpointsOutput), nil
	}
	out, err := c.S3OutpostsAPI.ListSharedEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

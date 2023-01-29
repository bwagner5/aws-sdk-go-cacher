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
	"github.com/aws/aws-sdk-go/service/mwaa"
	"github.com/aws/aws-sdk-go/service/mwaa/mwaaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MWAA struct {
	mwaaiface.MWAAAPI
	cache *cache.Cache
}

func NewMWAA(mwaaapi mwaaiface.MWAAAPI) *MWAA {
	return &MWAA{
		MWAAAPI: mwaaapi,
		cache:   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MWAA) GetEnvironment(input *mwaa.GetEnvironmentInput) (*mwaa.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.GetEnvironmentOutput), nil
	}
	out, err := c.MWAAAPI.GetEnvironment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MWAA) GetEnvironmentWithContext(ctx context.Context, input *mwaa.GetEnvironmentInput, opts ...request.Option) (*mwaa.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.GetEnvironmentOutput), nil
	}
	out, err := c.MWAAAPI.GetEnvironmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MWAA) ListEnvironments(input *mwaa.ListEnvironmentsInput) (*mwaa.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.ListEnvironmentsOutput), nil
	}
	out, err := c.MWAAAPI.ListEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MWAA) ListEnvironmentsPages(input *mwaa.ListEnvironmentsInput, fn func(*mwaa.ListEnvironmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mwaa.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &mwaa.ListEnvironmentsOutput{}
	fnCacher := func(out *mwaa.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.MWAAAPI.ListEnvironmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MWAA) ListEnvironmentsPagesWithContext(ctx context.Context, input *mwaa.ListEnvironmentsInput, fn func(*mwaa.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mwaa.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &mwaa.ListEnvironmentsOutput{}
	fnCacher := func(out *mwaa.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.MWAAAPI.ListEnvironmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MWAA) ListEnvironmentsWithContext(ctx context.Context, input *mwaa.ListEnvironmentsInput, opts ...request.Option) (*mwaa.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.ListEnvironmentsOutput), nil
	}
	out, err := c.MWAAAPI.ListEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MWAA) ListTagsForResource(input *mwaa.ListTagsForResourceInput) (*mwaa.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.ListTagsForResourceOutput), nil
	}
	out, err := c.MWAAAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MWAA) ListTagsForResourceWithContext(ctx context.Context, input *mwaa.ListTagsForResourceInput, opts ...request.Option) (*mwaa.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mwaa.ListTagsForResourceOutput), nil
	}
	out, err := c.MWAAAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/finspace"
	"github.com/aws/aws-sdk-go/service/finspace/finspaceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Finspace struct {
	finspaceiface.FinspaceAPI
	cache *cache.Cache
}

func NewFinspace(finspaceapi finspaceiface.FinspaceAPI) *Finspace {
	return &Finspace{
		FinspaceAPI: finspaceapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Finspace) GetEnvironment(input *finspace.GetEnvironmentInput) (*finspace.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.GetEnvironmentOutput), nil
	}
	out, err := c.FinspaceAPI.GetEnvironment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Finspace) GetEnvironmentWithContext(ctx context.Context, input *finspace.GetEnvironmentInput, opts ...request.Option) (*finspace.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.GetEnvironmentOutput), nil
	}
	out, err := c.FinspaceAPI.GetEnvironmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Finspace) ListEnvironments(input *finspace.ListEnvironmentsInput) (*finspace.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.ListEnvironmentsOutput), nil
	}
	out, err := c.FinspaceAPI.ListEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Finspace) ListEnvironmentsWithContext(ctx context.Context, input *finspace.ListEnvironmentsInput, opts ...request.Option) (*finspace.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.ListEnvironmentsOutput), nil
	}
	out, err := c.FinspaceAPI.ListEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Finspace) ListTagsForResource(input *finspace.ListTagsForResourceInput) (*finspace.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.ListTagsForResourceOutput), nil
	}
	out, err := c.FinspaceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Finspace) ListTagsForResourceWithContext(ctx context.Context, input *finspace.ListTagsForResourceInput, opts ...request.Option) (*finspace.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspace.ListTagsForResourceOutput), nil
	}
	out, err := c.FinspaceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package simspaceweavercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/simspaceweaver"
	"github.com/aws/aws-sdk-go/service/simspaceweaver/simspaceweaveriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SimSpaceWeaver struct {
	simspaceweaveriface.SimSpaceWeaverAPI
	cache *cache.Cache
}

func New(simspaceweaverapi simspaceweaveriface.SimSpaceWeaverAPI) *SimSpaceWeaver {
	return &SimSpaceWeaver{
		SimSpaceWeaverAPI: simspaceweaverapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SimSpaceWeaver) DescribeApp(input *simspaceweaver.DescribeAppInput) (*simspaceweaver.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.DescribeAppOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.DescribeApp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) DescribeAppWithContext(ctx context.Context, input *simspaceweaver.DescribeAppInput, opts ...request.Option) (*simspaceweaver.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.DescribeAppOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.DescribeAppWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) DescribeSimulation(input *simspaceweaver.DescribeSimulationInput) (*simspaceweaver.DescribeSimulationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.DescribeSimulationOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.DescribeSimulation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) DescribeSimulationWithContext(ctx context.Context, input *simspaceweaver.DescribeSimulationInput, opts ...request.Option) (*simspaceweaver.DescribeSimulationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.DescribeSimulationOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.DescribeSimulationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListApps(input *simspaceweaver.ListAppsInput) (*simspaceweaver.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListAppsOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListAppsPages(input *simspaceweaver.ListAppsInput, fn func(*simspaceweaver.ListAppsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*simspaceweaver.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &simspaceweaver.ListAppsOutput{}
	fnCacher := func(out *simspaceweaver.ListAppsOutput, more bool) bool {
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
	if err := c.SimSpaceWeaverAPI.ListAppsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SimSpaceWeaver) ListAppsPagesWithContext(ctx context.Context, input *simspaceweaver.ListAppsInput, fn func(*simspaceweaver.ListAppsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*simspaceweaver.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &simspaceweaver.ListAppsOutput{}
	fnCacher := func(out *simspaceweaver.ListAppsOutput, more bool) bool {
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
	if err := c.SimSpaceWeaverAPI.ListAppsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SimSpaceWeaver) ListAppsWithContext(ctx context.Context, input *simspaceweaver.ListAppsInput, opts ...request.Option) (*simspaceweaver.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListAppsOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListSimulations(input *simspaceweaver.ListSimulationsInput) (*simspaceweaver.ListSimulationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListSimulationsOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListSimulations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListSimulationsPages(input *simspaceweaver.ListSimulationsInput, fn func(*simspaceweaver.ListSimulationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*simspaceweaver.ListSimulationsOutput), false)
		return nil
	}
	cachable := true
	output := &simspaceweaver.ListSimulationsOutput{}
	fnCacher := func(out *simspaceweaver.ListSimulationsOutput, more bool) bool {
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
	if err := c.SimSpaceWeaverAPI.ListSimulationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SimSpaceWeaver) ListSimulationsPagesWithContext(ctx context.Context, input *simspaceweaver.ListSimulationsInput, fn func(*simspaceweaver.ListSimulationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*simspaceweaver.ListSimulationsOutput), false)
		return nil
	}
	cachable := true
	output := &simspaceweaver.ListSimulationsOutput{}
	fnCacher := func(out *simspaceweaver.ListSimulationsOutput, more bool) bool {
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
	if err := c.SimSpaceWeaverAPI.ListSimulationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SimSpaceWeaver) ListSimulationsWithContext(ctx context.Context, input *simspaceweaver.ListSimulationsInput, opts ...request.Option) (*simspaceweaver.ListSimulationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListSimulationsOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListSimulationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListTagsForResource(input *simspaceweaver.ListTagsForResourceInput) (*simspaceweaver.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListTagsForResourceOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SimSpaceWeaver) ListTagsForResourceWithContext(ctx context.Context, input *simspaceweaver.ListTagsForResourceInput, opts ...request.Option) (*simspaceweaver.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*simspaceweaver.ListTagsForResourceOutput), nil
	}
	out, err := c.SimSpaceWeaverAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

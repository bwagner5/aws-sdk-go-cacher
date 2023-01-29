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
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoT1ClickProjects struct {
	iot1clickprojectsiface.IoT1ClickProjectsAPI
	cache *cache.Cache
}

func NewIoT1ClickProjects(iot1clickprojectsapi iot1clickprojectsiface.IoT1ClickProjectsAPI) *IoT1ClickProjects {
	return &IoT1ClickProjects{
		IoT1ClickProjectsAPI: iot1clickprojectsapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoT1ClickProjects) DescribePlacement(input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.DescribePlacementOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.DescribePlacement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) DescribePlacementWithContext(ctx context.Context, input *iot1clickprojects.DescribePlacementInput, opts ...request.Option) (*iot1clickprojects.DescribePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.DescribePlacementOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.DescribePlacementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) DescribeProject(input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.DescribeProjectOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) DescribeProjectWithContext(ctx context.Context, input *iot1clickprojects.DescribeProjectInput, opts ...request.Option) (*iot1clickprojects.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.DescribeProjectOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) GetDevicesInPlacement(input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.GetDevicesInPlacementOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.GetDevicesInPlacement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) GetDevicesInPlacementWithContext(ctx context.Context, input *iot1clickprojects.GetDevicesInPlacementInput, opts ...request.Option) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.GetDevicesInPlacementOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.GetDevicesInPlacementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListPlacements(input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListPlacementsOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListPlacements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListPlacementsPages(input *iot1clickprojects.ListPlacementsInput, fn func(*iot1clickprojects.ListPlacementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot1clickprojects.ListPlacementsOutput), false)
		return nil
	}
	cachable := true
	output := &iot1clickprojects.ListPlacementsOutput{}
	fnCacher := func(out *iot1clickprojects.ListPlacementsOutput, more bool) bool {
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
	if err := c.IoT1ClickProjectsAPI.ListPlacementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT1ClickProjects) ListPlacementsPagesWithContext(ctx context.Context, input *iot1clickprojects.ListPlacementsInput, fn func(*iot1clickprojects.ListPlacementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot1clickprojects.ListPlacementsOutput), false)
		return nil
	}
	cachable := true
	output := &iot1clickprojects.ListPlacementsOutput{}
	fnCacher := func(out *iot1clickprojects.ListPlacementsOutput, more bool) bool {
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
	if err := c.IoT1ClickProjectsAPI.ListPlacementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT1ClickProjects) ListPlacementsWithContext(ctx context.Context, input *iot1clickprojects.ListPlacementsInput, opts ...request.Option) (*iot1clickprojects.ListPlacementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListPlacementsOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListPlacementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListProjects(input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListProjectsOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListProjectsPages(input *iot1clickprojects.ListProjectsInput, fn func(*iot1clickprojects.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot1clickprojects.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &iot1clickprojects.ListProjectsOutput{}
	fnCacher := func(out *iot1clickprojects.ListProjectsOutput, more bool) bool {
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
	if err := c.IoT1ClickProjectsAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT1ClickProjects) ListProjectsPagesWithContext(ctx context.Context, input *iot1clickprojects.ListProjectsInput, fn func(*iot1clickprojects.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot1clickprojects.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &iot1clickprojects.ListProjectsOutput{}
	fnCacher := func(out *iot1clickprojects.ListProjectsOutput, more bool) bool {
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
	if err := c.IoT1ClickProjectsAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT1ClickProjects) ListProjectsWithContext(ctx context.Context, input *iot1clickprojects.ListProjectsInput, opts ...request.Option) (*iot1clickprojects.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListProjectsOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListTagsForResource(input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListTagsForResourceOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickProjects) ListTagsForResourceWithContext(ctx context.Context, input *iot1clickprojects.ListTagsForResourceInput, opts ...request.Option) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickprojects.ListTagsForResourceOutput), nil
	}
	out, err := c.IoT1ClickProjectsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

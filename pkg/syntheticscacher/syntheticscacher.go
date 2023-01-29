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
package syntheticscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Synthetics struct {
	syntheticsiface.SyntheticsAPI
	cache *cache.Cache
}

func New(syntheticsapi syntheticsiface.SyntheticsAPI) *Synthetics {
	return &Synthetics{
		SyntheticsAPI: syntheticsapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Synthetics) DescribeCanaries(input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeCanariesOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeCanaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) DescribeCanariesLastRun(input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeCanariesLastRunOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeCanariesLastRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) DescribeCanariesLastRunPages(input *synthetics.DescribeCanariesLastRunInput, fn func(*synthetics.DescribeCanariesLastRunOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeCanariesLastRunOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeCanariesLastRunOutput{}
	fnCacher := func(out *synthetics.DescribeCanariesLastRunOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeCanariesLastRunPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeCanariesLastRunPagesWithContext(ctx context.Context, input *synthetics.DescribeCanariesLastRunInput, fn func(*synthetics.DescribeCanariesLastRunOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeCanariesLastRunOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeCanariesLastRunOutput{}
	fnCacher := func(out *synthetics.DescribeCanariesLastRunOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeCanariesLastRunPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeCanariesLastRunWithContext(ctx context.Context, input *synthetics.DescribeCanariesLastRunInput, opts ...request.Option) (*synthetics.DescribeCanariesLastRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeCanariesLastRunOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeCanariesLastRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) DescribeCanariesPages(input *synthetics.DescribeCanariesInput, fn func(*synthetics.DescribeCanariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeCanariesOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeCanariesOutput{}
	fnCacher := func(out *synthetics.DescribeCanariesOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeCanariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeCanariesPagesWithContext(ctx context.Context, input *synthetics.DescribeCanariesInput, fn func(*synthetics.DescribeCanariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeCanariesOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeCanariesOutput{}
	fnCacher := func(out *synthetics.DescribeCanariesOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeCanariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeCanariesWithContext(ctx context.Context, input *synthetics.DescribeCanariesInput, opts ...request.Option) (*synthetics.DescribeCanariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeCanariesOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeCanariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) DescribeRuntimeVersions(input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeRuntimeVersionsOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeRuntimeVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) DescribeRuntimeVersionsPages(input *synthetics.DescribeRuntimeVersionsInput, fn func(*synthetics.DescribeRuntimeVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeRuntimeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeRuntimeVersionsOutput{}
	fnCacher := func(out *synthetics.DescribeRuntimeVersionsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeRuntimeVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeRuntimeVersionsPagesWithContext(ctx context.Context, input *synthetics.DescribeRuntimeVersionsInput, fn func(*synthetics.DescribeRuntimeVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.DescribeRuntimeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.DescribeRuntimeVersionsOutput{}
	fnCacher := func(out *synthetics.DescribeRuntimeVersionsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.DescribeRuntimeVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) DescribeRuntimeVersionsWithContext(ctx context.Context, input *synthetics.DescribeRuntimeVersionsInput, opts ...request.Option) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.DescribeRuntimeVersionsOutput), nil
	}
	out, err := c.SyntheticsAPI.DescribeRuntimeVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetCanary(input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetCanaryOutput), nil
	}
	out, err := c.SyntheticsAPI.GetCanary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetCanaryRuns(input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetCanaryRunsOutput), nil
	}
	out, err := c.SyntheticsAPI.GetCanaryRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetCanaryRunsPages(input *synthetics.GetCanaryRunsInput, fn func(*synthetics.GetCanaryRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.GetCanaryRunsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.GetCanaryRunsOutput{}
	fnCacher := func(out *synthetics.GetCanaryRunsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.GetCanaryRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) GetCanaryRunsPagesWithContext(ctx context.Context, input *synthetics.GetCanaryRunsInput, fn func(*synthetics.GetCanaryRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.GetCanaryRunsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.GetCanaryRunsOutput{}
	fnCacher := func(out *synthetics.GetCanaryRunsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.GetCanaryRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) GetCanaryRunsWithContext(ctx context.Context, input *synthetics.GetCanaryRunsInput, opts ...request.Option) (*synthetics.GetCanaryRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetCanaryRunsOutput), nil
	}
	out, err := c.SyntheticsAPI.GetCanaryRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetCanaryWithContext(ctx context.Context, input *synthetics.GetCanaryInput, opts ...request.Option) (*synthetics.GetCanaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetCanaryOutput), nil
	}
	out, err := c.SyntheticsAPI.GetCanaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetGroup(input *synthetics.GetGroupInput) (*synthetics.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetGroupOutput), nil
	}
	out, err := c.SyntheticsAPI.GetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) GetGroupWithContext(ctx context.Context, input *synthetics.GetGroupInput, opts ...request.Option) (*synthetics.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.GetGroupOutput), nil
	}
	out, err := c.SyntheticsAPI.GetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListAssociatedGroups(input *synthetics.ListAssociatedGroupsInput) (*synthetics.ListAssociatedGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListAssociatedGroupsOutput), nil
	}
	out, err := c.SyntheticsAPI.ListAssociatedGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListAssociatedGroupsPages(input *synthetics.ListAssociatedGroupsInput, fn func(*synthetics.ListAssociatedGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListAssociatedGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListAssociatedGroupsOutput{}
	fnCacher := func(out *synthetics.ListAssociatedGroupsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListAssociatedGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListAssociatedGroupsPagesWithContext(ctx context.Context, input *synthetics.ListAssociatedGroupsInput, fn func(*synthetics.ListAssociatedGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListAssociatedGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListAssociatedGroupsOutput{}
	fnCacher := func(out *synthetics.ListAssociatedGroupsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListAssociatedGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListAssociatedGroupsWithContext(ctx context.Context, input *synthetics.ListAssociatedGroupsInput, opts ...request.Option) (*synthetics.ListAssociatedGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListAssociatedGroupsOutput), nil
	}
	out, err := c.SyntheticsAPI.ListAssociatedGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListGroupResources(input *synthetics.ListGroupResourcesInput) (*synthetics.ListGroupResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListGroupResourcesOutput), nil
	}
	out, err := c.SyntheticsAPI.ListGroupResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListGroupResourcesPages(input *synthetics.ListGroupResourcesInput, fn func(*synthetics.ListGroupResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListGroupResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListGroupResourcesOutput{}
	fnCacher := func(out *synthetics.ListGroupResourcesOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListGroupResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListGroupResourcesPagesWithContext(ctx context.Context, input *synthetics.ListGroupResourcesInput, fn func(*synthetics.ListGroupResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListGroupResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListGroupResourcesOutput{}
	fnCacher := func(out *synthetics.ListGroupResourcesOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListGroupResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListGroupResourcesWithContext(ctx context.Context, input *synthetics.ListGroupResourcesInput, opts ...request.Option) (*synthetics.ListGroupResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListGroupResourcesOutput), nil
	}
	out, err := c.SyntheticsAPI.ListGroupResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListGroups(input *synthetics.ListGroupsInput) (*synthetics.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListGroupsOutput), nil
	}
	out, err := c.SyntheticsAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListGroupsPages(input *synthetics.ListGroupsInput, fn func(*synthetics.ListGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListGroupsOutput{}
	fnCacher := func(out *synthetics.ListGroupsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListGroupsPagesWithContext(ctx context.Context, input *synthetics.ListGroupsInput, fn func(*synthetics.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*synthetics.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &synthetics.ListGroupsOutput{}
	fnCacher := func(out *synthetics.ListGroupsOutput, more bool) bool {
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
	if err := c.SyntheticsAPI.ListGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Synthetics) ListGroupsWithContext(ctx context.Context, input *synthetics.ListGroupsInput, opts ...request.Option) (*synthetics.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListGroupsOutput), nil
	}
	out, err := c.SyntheticsAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListTagsForResource(input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListTagsForResourceOutput), nil
	}
	out, err := c.SyntheticsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Synthetics) ListTagsForResourceWithContext(ctx context.Context, input *synthetics.ListTagsForResourceInput, opts ...request.Option) (*synthetics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*synthetics.ListTagsForResourceOutput), nil
	}
	out, err := c.SyntheticsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

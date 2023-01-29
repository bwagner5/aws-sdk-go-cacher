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
package augmentedairuntimecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime/augmentedairuntimeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AugmentedAIRuntime struct {
	augmentedairuntimeiface.AugmentedAIRuntimeAPI
	cache *cache.Cache
}

func New(augmentedairuntimeapi augmentedairuntimeiface.AugmentedAIRuntimeAPI) *AugmentedAIRuntime {
	return &AugmentedAIRuntime{
		AugmentedAIRuntimeAPI: augmentedairuntimeapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AugmentedAIRuntime) DescribeHumanLoop(input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*augmentedairuntime.DescribeHumanLoopOutput), nil
	}
	out, err := c.AugmentedAIRuntimeAPI.DescribeHumanLoop(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AugmentedAIRuntime) DescribeHumanLoopWithContext(ctx context.Context, input *augmentedairuntime.DescribeHumanLoopInput, opts ...request.Option) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*augmentedairuntime.DescribeHumanLoopOutput), nil
	}
	out, err := c.AugmentedAIRuntimeAPI.DescribeHumanLoopWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AugmentedAIRuntime) ListHumanLoops(input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*augmentedairuntime.ListHumanLoopsOutput), nil
	}
	out, err := c.AugmentedAIRuntimeAPI.ListHumanLoops(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AugmentedAIRuntime) ListHumanLoopsPages(input *augmentedairuntime.ListHumanLoopsInput, fn func(*augmentedairuntime.ListHumanLoopsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*augmentedairuntime.ListHumanLoopsOutput), false)
		return nil
	}
	cachable := true
	output := &augmentedairuntime.ListHumanLoopsOutput{}
	fnCacher := func(out *augmentedairuntime.ListHumanLoopsOutput, more bool) bool {
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
	if err := c.AugmentedAIRuntimeAPI.ListHumanLoopsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AugmentedAIRuntime) ListHumanLoopsPagesWithContext(ctx context.Context, input *augmentedairuntime.ListHumanLoopsInput, fn func(*augmentedairuntime.ListHumanLoopsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*augmentedairuntime.ListHumanLoopsOutput), false)
		return nil
	}
	cachable := true
	output := &augmentedairuntime.ListHumanLoopsOutput{}
	fnCacher := func(out *augmentedairuntime.ListHumanLoopsOutput, more bool) bool {
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
	if err := c.AugmentedAIRuntimeAPI.ListHumanLoopsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AugmentedAIRuntime) ListHumanLoopsWithContext(ctx context.Context, input *augmentedairuntime.ListHumanLoopsInput, opts ...request.Option) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*augmentedairuntime.ListHumanLoopsOutput), nil
	}
	out, err := c.AugmentedAIRuntimeAPI.ListHumanLoopsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

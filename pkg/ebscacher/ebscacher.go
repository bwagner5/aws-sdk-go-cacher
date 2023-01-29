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
package ebscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/service/ebs/ebsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EBS struct {
	ebsiface.EBSAPI
	cache *cache.Cache
}

func New(ebsapi ebsiface.EBSAPI) *EBS {
	return &EBS{
		EBSAPI: ebsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EBS) GetSnapshotBlock(input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.GetSnapshotBlockOutput), nil
	}
	out, err := c.EBSAPI.GetSnapshotBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EBS) GetSnapshotBlockWithContext(ctx context.Context, input *ebs.GetSnapshotBlockInput, opts ...request.Option) (*ebs.GetSnapshotBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.GetSnapshotBlockOutput), nil
	}
	out, err := c.EBSAPI.GetSnapshotBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EBS) ListChangedBlocks(input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.ListChangedBlocksOutput), nil
	}
	out, err := c.EBSAPI.ListChangedBlocks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EBS) ListChangedBlocksPages(input *ebs.ListChangedBlocksInput, fn func(*ebs.ListChangedBlocksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ebs.ListChangedBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &ebs.ListChangedBlocksOutput{}
	fnCacher := func(out *ebs.ListChangedBlocksOutput, more bool) bool {
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
	if err := c.EBSAPI.ListChangedBlocksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EBS) ListChangedBlocksPagesWithContext(ctx context.Context, input *ebs.ListChangedBlocksInput, fn func(*ebs.ListChangedBlocksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ebs.ListChangedBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &ebs.ListChangedBlocksOutput{}
	fnCacher := func(out *ebs.ListChangedBlocksOutput, more bool) bool {
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
	if err := c.EBSAPI.ListChangedBlocksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EBS) ListChangedBlocksWithContext(ctx context.Context, input *ebs.ListChangedBlocksInput, opts ...request.Option) (*ebs.ListChangedBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.ListChangedBlocksOutput), nil
	}
	out, err := c.EBSAPI.ListChangedBlocksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EBS) ListSnapshotBlocks(input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.ListSnapshotBlocksOutput), nil
	}
	out, err := c.EBSAPI.ListSnapshotBlocks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EBS) ListSnapshotBlocksPages(input *ebs.ListSnapshotBlocksInput, fn func(*ebs.ListSnapshotBlocksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ebs.ListSnapshotBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &ebs.ListSnapshotBlocksOutput{}
	fnCacher := func(out *ebs.ListSnapshotBlocksOutput, more bool) bool {
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
	if err := c.EBSAPI.ListSnapshotBlocksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EBS) ListSnapshotBlocksPagesWithContext(ctx context.Context, input *ebs.ListSnapshotBlocksInput, fn func(*ebs.ListSnapshotBlocksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ebs.ListSnapshotBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &ebs.ListSnapshotBlocksOutput{}
	fnCacher := func(out *ebs.ListSnapshotBlocksOutput, more bool) bool {
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
	if err := c.EBSAPI.ListSnapshotBlocksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EBS) ListSnapshotBlocksWithContext(ctx context.Context, input *ebs.ListSnapshotBlocksInput, opts ...request.Option) (*ebs.ListSnapshotBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ebs.ListSnapshotBlocksOutput), nil
	}
	out, err := c.EBSAPI.ListSnapshotBlocksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

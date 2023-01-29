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
	"github.com/aws/aws-sdk-go/service/controltower"
	"github.com/aws/aws-sdk-go/service/controltower/controltoweriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ControlTower struct {
	controltoweriface.ControlTowerAPI
	cache *cache.Cache
}

func NewControlTower(controltowerapi controltoweriface.ControlTowerAPI) *ControlTower {
	return &ControlTower{
		ControlTowerAPI: controltowerapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ControlTower) GetControlOperation(input *controltower.GetControlOperationInput) (*controltower.GetControlOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*controltower.GetControlOperationOutput), nil
	}
	out, err := c.ControlTowerAPI.GetControlOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ControlTower) GetControlOperationWithContext(ctx context.Context, input *controltower.GetControlOperationInput, opts ...request.Option) (*controltower.GetControlOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*controltower.GetControlOperationOutput), nil
	}
	out, err := c.ControlTowerAPI.GetControlOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ControlTower) ListEnabledControls(input *controltower.ListEnabledControlsInput) (*controltower.ListEnabledControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*controltower.ListEnabledControlsOutput), nil
	}
	out, err := c.ControlTowerAPI.ListEnabledControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ControlTower) ListEnabledControlsPages(input *controltower.ListEnabledControlsInput, fn func(*controltower.ListEnabledControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*controltower.ListEnabledControlsOutput), false)
		return nil
	}
	cachable := true
	output := &controltower.ListEnabledControlsOutput{}
	fnCacher := func(out *controltower.ListEnabledControlsOutput, more bool) bool {
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
	if err := c.ControlTowerAPI.ListEnabledControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ControlTower) ListEnabledControlsPagesWithContext(ctx context.Context, input *controltower.ListEnabledControlsInput, fn func(*controltower.ListEnabledControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*controltower.ListEnabledControlsOutput), false)
		return nil
	}
	cachable := true
	output := &controltower.ListEnabledControlsOutput{}
	fnCacher := func(out *controltower.ListEnabledControlsOutput, more bool) bool {
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
	if err := c.ControlTowerAPI.ListEnabledControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ControlTower) ListEnabledControlsWithContext(ctx context.Context, input *controltower.ListEnabledControlsInput, opts ...request.Option) (*controltower.ListEnabledControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*controltower.ListEnabledControlsOutput), nil
	}
	out, err := c.ControlTowerAPI.ListEnabledControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package migrationhubconfigcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MigrationHubConfig struct {
	migrationhubconfigiface.MigrationHubConfigAPI
	cache *cache.Cache
}

func New(migrationhubconfigapi migrationhubconfigiface.MigrationHubConfigAPI) *MigrationHubConfig {
	return &MigrationHubConfig{
		MigrationHubConfigAPI: migrationhubconfigapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MigrationHubConfig) DescribeHomeRegionControls(input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubconfig.DescribeHomeRegionControlsOutput), nil
	}
	out, err := c.MigrationHubConfigAPI.DescribeHomeRegionControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubConfig) DescribeHomeRegionControlsPages(input *migrationhubconfig.DescribeHomeRegionControlsInput, fn func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubconfig.DescribeHomeRegionControlsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubconfig.DescribeHomeRegionControlsOutput{}
	fnCacher := func(out *migrationhubconfig.DescribeHomeRegionControlsOutput, more bool) bool {
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
	if err := c.MigrationHubConfigAPI.DescribeHomeRegionControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubConfig) DescribeHomeRegionControlsPagesWithContext(ctx context.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput, fn func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubconfig.DescribeHomeRegionControlsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubconfig.DescribeHomeRegionControlsOutput{}
	fnCacher := func(out *migrationhubconfig.DescribeHomeRegionControlsOutput, more bool) bool {
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
	if err := c.MigrationHubConfigAPI.DescribeHomeRegionControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubConfig) DescribeHomeRegionControlsWithContext(ctx context.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput, opts ...request.Option) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubconfig.DescribeHomeRegionControlsOutput), nil
	}
	out, err := c.MigrationHubConfigAPI.DescribeHomeRegionControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubConfig) GetHomeRegion(input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubconfig.GetHomeRegionOutput), nil
	}
	out, err := c.MigrationHubConfigAPI.GetHomeRegion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubConfig) GetHomeRegionWithContext(ctx context.Context, input *migrationhubconfig.GetHomeRegionInput, opts ...request.Option) (*migrationhubconfig.GetHomeRegionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubconfig.GetHomeRegionOutput), nil
	}
	out, err := c.MigrationHubConfigAPI.GetHomeRegionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

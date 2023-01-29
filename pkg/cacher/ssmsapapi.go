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
	"github.com/aws/aws-sdk-go/service/ssmsap"
	"github.com/aws/aws-sdk-go/service/ssmsap/ssmsapiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SsmSap struct {
	ssmsapiface.SsmSapAPI
	cache *cache.Cache
}

func NewSsmSap(ssmsapapi ssmsapiface.SsmSapAPI) *SsmSap {
	return &SsmSap{
		SsmSapAPI: ssmsapapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SsmSap) GetApplication(input *ssmsap.GetApplicationInput) (*ssmsap.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetApplicationOutput), nil
	}
	out, err := c.SsmSapAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetApplicationWithContext(ctx context.Context, input *ssmsap.GetApplicationInput, opts ...request.Option) (*ssmsap.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetApplicationOutput), nil
	}
	out, err := c.SsmSapAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetComponent(input *ssmsap.GetComponentInput) (*ssmsap.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetComponentOutput), nil
	}
	out, err := c.SsmSapAPI.GetComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetComponentWithContext(ctx context.Context, input *ssmsap.GetComponentInput, opts ...request.Option) (*ssmsap.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetComponentOutput), nil
	}
	out, err := c.SsmSapAPI.GetComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetDatabase(input *ssmsap.GetDatabaseInput) (*ssmsap.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetDatabaseOutput), nil
	}
	out, err := c.SsmSapAPI.GetDatabase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetDatabaseWithContext(ctx context.Context, input *ssmsap.GetDatabaseInput, opts ...request.Option) (*ssmsap.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetDatabaseOutput), nil
	}
	out, err := c.SsmSapAPI.GetDatabaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetOperation(input *ssmsap.GetOperationInput) (*ssmsap.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetOperationOutput), nil
	}
	out, err := c.SsmSapAPI.GetOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetOperationWithContext(ctx context.Context, input *ssmsap.GetOperationInput, opts ...request.Option) (*ssmsap.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetOperationOutput), nil
	}
	out, err := c.SsmSapAPI.GetOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetResourcePermission(input *ssmsap.GetResourcePermissionInput) (*ssmsap.GetResourcePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetResourcePermissionOutput), nil
	}
	out, err := c.SsmSapAPI.GetResourcePermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) GetResourcePermissionWithContext(ctx context.Context, input *ssmsap.GetResourcePermissionInput, opts ...request.Option) (*ssmsap.GetResourcePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.GetResourcePermissionOutput), nil
	}
	out, err := c.SsmSapAPI.GetResourcePermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListApplications(input *ssmsap.ListApplicationsInput) (*ssmsap.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListApplicationsOutput), nil
	}
	out, err := c.SsmSapAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListApplicationsPages(input *ssmsap.ListApplicationsInput, fn func(*ssmsap.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListApplicationsOutput{}
	fnCacher := func(out *ssmsap.ListApplicationsOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListApplicationsPagesWithContext(ctx context.Context, input *ssmsap.ListApplicationsInput, fn func(*ssmsap.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListApplicationsOutput{}
	fnCacher := func(out *ssmsap.ListApplicationsOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListApplicationsWithContext(ctx context.Context, input *ssmsap.ListApplicationsInput, opts ...request.Option) (*ssmsap.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListApplicationsOutput), nil
	}
	out, err := c.SsmSapAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListComponents(input *ssmsap.ListComponentsInput) (*ssmsap.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListComponentsOutput), nil
	}
	out, err := c.SsmSapAPI.ListComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListComponentsPages(input *ssmsap.ListComponentsInput, fn func(*ssmsap.ListComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListComponentsOutput{}
	fnCacher := func(out *ssmsap.ListComponentsOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListComponentsPagesWithContext(ctx context.Context, input *ssmsap.ListComponentsInput, fn func(*ssmsap.ListComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListComponentsOutput{}
	fnCacher := func(out *ssmsap.ListComponentsOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListComponentsWithContext(ctx context.Context, input *ssmsap.ListComponentsInput, opts ...request.Option) (*ssmsap.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListComponentsOutput), nil
	}
	out, err := c.SsmSapAPI.ListComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListDatabases(input *ssmsap.ListDatabasesInput) (*ssmsap.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListDatabasesOutput), nil
	}
	out, err := c.SsmSapAPI.ListDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListDatabasesPages(input *ssmsap.ListDatabasesInput, fn func(*ssmsap.ListDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListDatabasesOutput{}
	fnCacher := func(out *ssmsap.ListDatabasesOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListDatabasesPagesWithContext(ctx context.Context, input *ssmsap.ListDatabasesInput, fn func(*ssmsap.ListDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmsap.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &ssmsap.ListDatabasesOutput{}
	fnCacher := func(out *ssmsap.ListDatabasesOutput, more bool) bool {
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
	if err := c.SsmSapAPI.ListDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SsmSap) ListDatabasesWithContext(ctx context.Context, input *ssmsap.ListDatabasesInput, opts ...request.Option) (*ssmsap.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListDatabasesOutput), nil
	}
	out, err := c.SsmSapAPI.ListDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListTagsForResource(input *ssmsap.ListTagsForResourceInput) (*ssmsap.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListTagsForResourceOutput), nil
	}
	out, err := c.SsmSapAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SsmSap) ListTagsForResourceWithContext(ctx context.Context, input *ssmsap.ListTagsForResourceInput, opts ...request.Option) (*ssmsap.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmsap.ListTagsForResourceOutput), nil
	}
	out, err := c.SsmSapAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

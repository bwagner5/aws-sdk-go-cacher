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
package appregistrycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appregistry"
	"github.com/aws/aws-sdk-go/service/appregistry/appregistryiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppRegistry struct {
	appregistryiface.AppRegistryAPI
	cache *cache.Cache
}

func New(appregistryapi appregistryiface.AppRegistryAPI) *AppRegistry {
	return &AppRegistry{
		AppRegistryAPI: appregistryapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppRegistry) GetApplication(input *appregistry.GetApplicationInput) (*appregistry.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetApplicationOutput), nil
	}
	out, err := c.AppRegistryAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetApplicationWithContext(ctx context.Context, input *appregistry.GetApplicationInput, opts ...request.Option) (*appregistry.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetApplicationOutput), nil
	}
	out, err := c.AppRegistryAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetAssociatedResource(input *appregistry.GetAssociatedResourceInput) (*appregistry.GetAssociatedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetAssociatedResourceOutput), nil
	}
	out, err := c.AppRegistryAPI.GetAssociatedResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetAssociatedResourceWithContext(ctx context.Context, input *appregistry.GetAssociatedResourceInput, opts ...request.Option) (*appregistry.GetAssociatedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetAssociatedResourceOutput), nil
	}
	out, err := c.AppRegistryAPI.GetAssociatedResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetAttributeGroup(input *appregistry.GetAttributeGroupInput) (*appregistry.GetAttributeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetAttributeGroupOutput), nil
	}
	out, err := c.AppRegistryAPI.GetAttributeGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetAttributeGroupWithContext(ctx context.Context, input *appregistry.GetAttributeGroupInput, opts ...request.Option) (*appregistry.GetAttributeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetAttributeGroupOutput), nil
	}
	out, err := c.AppRegistryAPI.GetAttributeGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetConfiguration(input *appregistry.GetConfigurationInput) (*appregistry.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetConfigurationOutput), nil
	}
	out, err := c.AppRegistryAPI.GetConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) GetConfigurationWithContext(ctx context.Context, input *appregistry.GetConfigurationInput, opts ...request.Option) (*appregistry.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.GetConfigurationOutput), nil
	}
	out, err := c.AppRegistryAPI.GetConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListApplications(input *appregistry.ListApplicationsInput) (*appregistry.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListApplicationsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListApplicationsPages(input *appregistry.ListApplicationsInput, fn func(*appregistry.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListApplicationsOutput{}
	fnCacher := func(out *appregistry.ListApplicationsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListApplicationsPagesWithContext(ctx context.Context, input *appregistry.ListApplicationsInput, fn func(*appregistry.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListApplicationsOutput{}
	fnCacher := func(out *appregistry.ListApplicationsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListApplicationsWithContext(ctx context.Context, input *appregistry.ListApplicationsInput, opts ...request.Option) (*appregistry.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListApplicationsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAssociatedAttributeGroups(input *appregistry.ListAssociatedAttributeGroupsInput) (*appregistry.ListAssociatedAttributeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAssociatedAttributeGroupsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAssociatedAttributeGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAssociatedAttributeGroupsPages(input *appregistry.ListAssociatedAttributeGroupsInput, fn func(*appregistry.ListAssociatedAttributeGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAssociatedAttributeGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAssociatedAttributeGroupsOutput{}
	fnCacher := func(out *appregistry.ListAssociatedAttributeGroupsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAssociatedAttributeGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAssociatedAttributeGroupsPagesWithContext(ctx context.Context, input *appregistry.ListAssociatedAttributeGroupsInput, fn func(*appregistry.ListAssociatedAttributeGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAssociatedAttributeGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAssociatedAttributeGroupsOutput{}
	fnCacher := func(out *appregistry.ListAssociatedAttributeGroupsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAssociatedAttributeGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAssociatedAttributeGroupsWithContext(ctx context.Context, input *appregistry.ListAssociatedAttributeGroupsInput, opts ...request.Option) (*appregistry.ListAssociatedAttributeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAssociatedAttributeGroupsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAssociatedAttributeGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAssociatedResources(input *appregistry.ListAssociatedResourcesInput) (*appregistry.ListAssociatedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAssociatedResourcesOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAssociatedResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAssociatedResourcesPages(input *appregistry.ListAssociatedResourcesInput, fn func(*appregistry.ListAssociatedResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAssociatedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAssociatedResourcesOutput{}
	fnCacher := func(out *appregistry.ListAssociatedResourcesOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAssociatedResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAssociatedResourcesPagesWithContext(ctx context.Context, input *appregistry.ListAssociatedResourcesInput, fn func(*appregistry.ListAssociatedResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAssociatedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAssociatedResourcesOutput{}
	fnCacher := func(out *appregistry.ListAssociatedResourcesOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAssociatedResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAssociatedResourcesWithContext(ctx context.Context, input *appregistry.ListAssociatedResourcesInput, opts ...request.Option) (*appregistry.ListAssociatedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAssociatedResourcesOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAssociatedResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAttributeGroups(input *appregistry.ListAttributeGroupsInput) (*appregistry.ListAttributeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAttributeGroupsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAttributeGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAttributeGroupsForApplication(input *appregistry.ListAttributeGroupsForApplicationInput) (*appregistry.ListAttributeGroupsForApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAttributeGroupsForApplicationOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAttributeGroupsForApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAttributeGroupsForApplicationPages(input *appregistry.ListAttributeGroupsForApplicationInput, fn func(*appregistry.ListAttributeGroupsForApplicationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAttributeGroupsForApplicationOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAttributeGroupsForApplicationOutput{}
	fnCacher := func(out *appregistry.ListAttributeGroupsForApplicationOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAttributeGroupsForApplicationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAttributeGroupsForApplicationPagesWithContext(ctx context.Context, input *appregistry.ListAttributeGroupsForApplicationInput, fn func(*appregistry.ListAttributeGroupsForApplicationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAttributeGroupsForApplicationOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAttributeGroupsForApplicationOutput{}
	fnCacher := func(out *appregistry.ListAttributeGroupsForApplicationOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAttributeGroupsForApplicationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAttributeGroupsForApplicationWithContext(ctx context.Context, input *appregistry.ListAttributeGroupsForApplicationInput, opts ...request.Option) (*appregistry.ListAttributeGroupsForApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAttributeGroupsForApplicationOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAttributeGroupsForApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListAttributeGroupsPages(input *appregistry.ListAttributeGroupsInput, fn func(*appregistry.ListAttributeGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAttributeGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAttributeGroupsOutput{}
	fnCacher := func(out *appregistry.ListAttributeGroupsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAttributeGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAttributeGroupsPagesWithContext(ctx context.Context, input *appregistry.ListAttributeGroupsInput, fn func(*appregistry.ListAttributeGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appregistry.ListAttributeGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &appregistry.ListAttributeGroupsOutput{}
	fnCacher := func(out *appregistry.ListAttributeGroupsOutput, more bool) bool {
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
	if err := c.AppRegistryAPI.ListAttributeGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRegistry) ListAttributeGroupsWithContext(ctx context.Context, input *appregistry.ListAttributeGroupsInput, opts ...request.Option) (*appregistry.ListAttributeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListAttributeGroupsOutput), nil
	}
	out, err := c.AppRegistryAPI.ListAttributeGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListTagsForResource(input *appregistry.ListTagsForResourceInput) (*appregistry.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListTagsForResourceOutput), nil
	}
	out, err := c.AppRegistryAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRegistry) ListTagsForResourceWithContext(ctx context.Context, input *appregistry.ListTagsForResourceInput, opts ...request.Option) (*appregistry.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appregistry.ListTagsForResourceOutput), nil
	}
	out, err := c.AppRegistryAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

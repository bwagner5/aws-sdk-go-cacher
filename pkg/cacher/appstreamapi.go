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
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appstream/appstreamiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppStream struct {
	appstreamiface.AppStreamAPI
	cache *cache.Cache
}

func NewAppStream(appstreamapi appstreamiface.AppStreamAPI) *AppStream {
	return &AppStream{
		AppStreamAPI: appstreamapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppStream) BatchAssociateUserStack(input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.BatchAssociateUserStackOutput), nil
	}
	out, err := c.AppStreamAPI.BatchAssociateUserStack(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) BatchAssociateUserStackWithContext(ctx context.Context, input *appstream.BatchAssociateUserStackInput, opts ...request.Option) (*appstream.BatchAssociateUserStackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.BatchAssociateUserStackOutput), nil
	}
	out, err := c.AppStreamAPI.BatchAssociateUserStackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) BatchDisassociateUserStack(input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.BatchDisassociateUserStackOutput), nil
	}
	out, err := c.AppStreamAPI.BatchDisassociateUserStack(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) BatchDisassociateUserStackWithContext(ctx context.Context, input *appstream.BatchDisassociateUserStackInput, opts ...request.Option) (*appstream.BatchDisassociateUserStackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.BatchDisassociateUserStackOutput), nil
	}
	out, err := c.AppStreamAPI.BatchDisassociateUserStackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeAppBlocks(input *appstream.DescribeAppBlocksInput) (*appstream.DescribeAppBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeAppBlocksOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeAppBlocks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeAppBlocksWithContext(ctx context.Context, input *appstream.DescribeAppBlocksInput, opts ...request.Option) (*appstream.DescribeAppBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeAppBlocksOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeAppBlocksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeApplicationFleetAssociations(input *appstream.DescribeApplicationFleetAssociationsInput) (*appstream.DescribeApplicationFleetAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeApplicationFleetAssociationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeApplicationFleetAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeApplicationFleetAssociationsWithContext(ctx context.Context, input *appstream.DescribeApplicationFleetAssociationsInput, opts ...request.Option) (*appstream.DescribeApplicationFleetAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeApplicationFleetAssociationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeApplicationFleetAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeApplications(input *appstream.DescribeApplicationsInput) (*appstream.DescribeApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeApplicationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeApplicationsWithContext(ctx context.Context, input *appstream.DescribeApplicationsInput, opts ...request.Option) (*appstream.DescribeApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeApplicationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeDirectoryConfigs(input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeDirectoryConfigsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeDirectoryConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeDirectoryConfigsWithContext(ctx context.Context, input *appstream.DescribeDirectoryConfigsInput, opts ...request.Option) (*appstream.DescribeDirectoryConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeDirectoryConfigsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeDirectoryConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeEntitlements(input *appstream.DescribeEntitlementsInput) (*appstream.DescribeEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeEntitlementsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeEntitlements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeEntitlementsWithContext(ctx context.Context, input *appstream.DescribeEntitlementsInput, opts ...request.Option) (*appstream.DescribeEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeEntitlementsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeEntitlementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeFleets(input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeFleetsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeFleetsWithContext(ctx context.Context, input *appstream.DescribeFleetsInput, opts ...request.Option) (*appstream.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeFleetsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImageBuilders(input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImageBuildersOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImageBuilders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImageBuildersWithContext(ctx context.Context, input *appstream.DescribeImageBuildersInput, opts ...request.Option) (*appstream.DescribeImageBuildersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImageBuildersOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImageBuildersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImagePermissions(input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImagePermissionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImagePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImagePermissionsPages(input *appstream.DescribeImagePermissionsInput, fn func(*appstream.DescribeImagePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appstream.DescribeImagePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &appstream.DescribeImagePermissionsOutput{}
	fnCacher := func(out *appstream.DescribeImagePermissionsOutput, more bool) bool {
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
	if err := c.AppStreamAPI.DescribeImagePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppStream) DescribeImagePermissionsPagesWithContext(ctx context.Context, input *appstream.DescribeImagePermissionsInput, fn func(*appstream.DescribeImagePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appstream.DescribeImagePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &appstream.DescribeImagePermissionsOutput{}
	fnCacher := func(out *appstream.DescribeImagePermissionsOutput, more bool) bool {
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
	if err := c.AppStreamAPI.DescribeImagePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppStream) DescribeImagePermissionsWithContext(ctx context.Context, input *appstream.DescribeImagePermissionsInput, opts ...request.Option) (*appstream.DescribeImagePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImagePermissionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImagePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImages(input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImagesOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeImagesPages(input *appstream.DescribeImagesInput, fn func(*appstream.DescribeImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appstream.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &appstream.DescribeImagesOutput{}
	fnCacher := func(out *appstream.DescribeImagesOutput, more bool) bool {
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
	if err := c.AppStreamAPI.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppStream) DescribeImagesPagesWithContext(ctx context.Context, input *appstream.DescribeImagesInput, fn func(*appstream.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appstream.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &appstream.DescribeImagesOutput{}
	fnCacher := func(out *appstream.DescribeImagesOutput, more bool) bool {
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
	if err := c.AppStreamAPI.DescribeImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppStream) DescribeImagesWithContext(ctx context.Context, input *appstream.DescribeImagesInput, opts ...request.Option) (*appstream.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeImagesOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeSessions(input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeSessionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeSessionsWithContext(ctx context.Context, input *appstream.DescribeSessionsInput, opts ...request.Option) (*appstream.DescribeSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeSessionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeStacks(input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeStacksOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeStacksWithContext(ctx context.Context, input *appstream.DescribeStacksInput, opts ...request.Option) (*appstream.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeStacksOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUsageReportSubscriptions(input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUsageReportSubscriptionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUsageReportSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUsageReportSubscriptionsWithContext(ctx context.Context, input *appstream.DescribeUsageReportSubscriptionsInput, opts ...request.Option) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUsageReportSubscriptionsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUsageReportSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUserStackAssociations(input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUserStackAssociationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUserStackAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUserStackAssociationsWithContext(ctx context.Context, input *appstream.DescribeUserStackAssociationsInput, opts ...request.Option) (*appstream.DescribeUserStackAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUserStackAssociationsOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUserStackAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUsers(input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUsersOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) DescribeUsersWithContext(ctx context.Context, input *appstream.DescribeUsersInput, opts ...request.Option) (*appstream.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.DescribeUsersOutput), nil
	}
	out, err := c.AppStreamAPI.DescribeUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListAssociatedFleets(input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListAssociatedFleetsOutput), nil
	}
	out, err := c.AppStreamAPI.ListAssociatedFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListAssociatedFleetsWithContext(ctx context.Context, input *appstream.ListAssociatedFleetsInput, opts ...request.Option) (*appstream.ListAssociatedFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListAssociatedFleetsOutput), nil
	}
	out, err := c.AppStreamAPI.ListAssociatedFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListAssociatedStacks(input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListAssociatedStacksOutput), nil
	}
	out, err := c.AppStreamAPI.ListAssociatedStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListAssociatedStacksWithContext(ctx context.Context, input *appstream.ListAssociatedStacksInput, opts ...request.Option) (*appstream.ListAssociatedStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListAssociatedStacksOutput), nil
	}
	out, err := c.AppStreamAPI.ListAssociatedStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListEntitledApplications(input *appstream.ListEntitledApplicationsInput) (*appstream.ListEntitledApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListEntitledApplicationsOutput), nil
	}
	out, err := c.AppStreamAPI.ListEntitledApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListEntitledApplicationsWithContext(ctx context.Context, input *appstream.ListEntitledApplicationsInput, opts ...request.Option) (*appstream.ListEntitledApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListEntitledApplicationsOutput), nil
	}
	out, err := c.AppStreamAPI.ListEntitledApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListTagsForResource(input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListTagsForResourceOutput), nil
	}
	out, err := c.AppStreamAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppStream) ListTagsForResourceWithContext(ctx context.Context, input *appstream.ListTagsForResourceInput, opts ...request.Option) (*appstream.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appstream.ListTagsForResourceOutput), nil
	}
	out, err := c.AppStreamAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

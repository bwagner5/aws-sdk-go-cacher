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
package finspacedatacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/finspacedata"
	"github.com/aws/aws-sdk-go/service/finspacedata/finspacedataiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type FinSpaceData struct {
	finspacedataiface.FinSpaceDataAPI
	cache *cache.Cache
}

func New(finspacedataapi finspacedataiface.FinSpaceDataAPI) *FinSpaceData {
	return &FinSpaceData{
		FinSpaceDataAPI: finspacedataapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *FinSpaceData) GetChangeset(input *finspacedata.GetChangesetInput) (*finspacedata.GetChangesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetChangesetOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetChangeset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetChangesetWithContext(ctx context.Context, input *finspacedata.GetChangesetInput, opts ...request.Option) (*finspacedata.GetChangesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetChangesetOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetChangesetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetDataView(input *finspacedata.GetDataViewInput) (*finspacedata.GetDataViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetDataViewOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetDataView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetDataViewWithContext(ctx context.Context, input *finspacedata.GetDataViewInput, opts ...request.Option) (*finspacedata.GetDataViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetDataViewOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetDataViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetDataset(input *finspacedata.GetDatasetInput) (*finspacedata.GetDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetDatasetOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetDatasetWithContext(ctx context.Context, input *finspacedata.GetDatasetInput, opts ...request.Option) (*finspacedata.GetDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetDatasetOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetExternalDataViewAccessDetails(input *finspacedata.GetExternalDataViewAccessDetailsInput) (*finspacedata.GetExternalDataViewAccessDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetExternalDataViewAccessDetailsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetExternalDataViewAccessDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetExternalDataViewAccessDetailsWithContext(ctx context.Context, input *finspacedata.GetExternalDataViewAccessDetailsInput, opts ...request.Option) (*finspacedata.GetExternalDataViewAccessDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetExternalDataViewAccessDetailsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetExternalDataViewAccessDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetPermissionGroup(input *finspacedata.GetPermissionGroupInput) (*finspacedata.GetPermissionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetPermissionGroupOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetPermissionGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetPermissionGroupWithContext(ctx context.Context, input *finspacedata.GetPermissionGroupInput, opts ...request.Option) (*finspacedata.GetPermissionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetPermissionGroupOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetPermissionGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetProgrammaticAccessCredentials(input *finspacedata.GetProgrammaticAccessCredentialsInput) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetProgrammaticAccessCredentialsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetProgrammaticAccessCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetProgrammaticAccessCredentialsWithContext(ctx context.Context, input *finspacedata.GetProgrammaticAccessCredentialsInput, opts ...request.Option) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetProgrammaticAccessCredentialsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetProgrammaticAccessCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetUser(input *finspacedata.GetUserInput) (*finspacedata.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetUserOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetUserWithContext(ctx context.Context, input *finspacedata.GetUserInput, opts ...request.Option) (*finspacedata.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetUserOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetWorkingLocation(input *finspacedata.GetWorkingLocationInput) (*finspacedata.GetWorkingLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetWorkingLocationOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetWorkingLocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) GetWorkingLocationWithContext(ctx context.Context, input *finspacedata.GetWorkingLocationInput, opts ...request.Option) (*finspacedata.GetWorkingLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.GetWorkingLocationOutput), nil
	}
	out, err := c.FinSpaceDataAPI.GetWorkingLocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListChangesets(input *finspacedata.ListChangesetsInput) (*finspacedata.ListChangesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListChangesetsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListChangesets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListChangesetsPages(input *finspacedata.ListChangesetsInput, fn func(*finspacedata.ListChangesetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListChangesetsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListChangesetsOutput{}
	fnCacher := func(out *finspacedata.ListChangesetsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListChangesetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListChangesetsPagesWithContext(ctx context.Context, input *finspacedata.ListChangesetsInput, fn func(*finspacedata.ListChangesetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListChangesetsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListChangesetsOutput{}
	fnCacher := func(out *finspacedata.ListChangesetsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListChangesetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListChangesetsWithContext(ctx context.Context, input *finspacedata.ListChangesetsInput, opts ...request.Option) (*finspacedata.ListChangesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListChangesetsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListChangesetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListDataViews(input *finspacedata.ListDataViewsInput) (*finspacedata.ListDataViewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListDataViewsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListDataViews(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListDataViewsPages(input *finspacedata.ListDataViewsInput, fn func(*finspacedata.ListDataViewsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListDataViewsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListDataViewsOutput{}
	fnCacher := func(out *finspacedata.ListDataViewsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListDataViewsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListDataViewsPagesWithContext(ctx context.Context, input *finspacedata.ListDataViewsInput, fn func(*finspacedata.ListDataViewsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListDataViewsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListDataViewsOutput{}
	fnCacher := func(out *finspacedata.ListDataViewsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListDataViewsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListDataViewsWithContext(ctx context.Context, input *finspacedata.ListDataViewsInput, opts ...request.Option) (*finspacedata.ListDataViewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListDataViewsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListDataViewsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListDatasets(input *finspacedata.ListDatasetsInput) (*finspacedata.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListDatasetsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListDatasetsPages(input *finspacedata.ListDatasetsInput, fn func(*finspacedata.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListDatasetsOutput{}
	fnCacher := func(out *finspacedata.ListDatasetsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListDatasetsPagesWithContext(ctx context.Context, input *finspacedata.ListDatasetsInput, fn func(*finspacedata.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListDatasetsOutput{}
	fnCacher := func(out *finspacedata.ListDatasetsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListDatasetsWithContext(ctx context.Context, input *finspacedata.ListDatasetsInput, opts ...request.Option) (*finspacedata.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListDatasetsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListPermissionGroups(input *finspacedata.ListPermissionGroupsInput) (*finspacedata.ListPermissionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListPermissionGroupsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListPermissionGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListPermissionGroupsByUser(input *finspacedata.ListPermissionGroupsByUserInput) (*finspacedata.ListPermissionGroupsByUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListPermissionGroupsByUserOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListPermissionGroupsByUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListPermissionGroupsByUserWithContext(ctx context.Context, input *finspacedata.ListPermissionGroupsByUserInput, opts ...request.Option) (*finspacedata.ListPermissionGroupsByUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListPermissionGroupsByUserOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListPermissionGroupsByUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListPermissionGroupsPages(input *finspacedata.ListPermissionGroupsInput, fn func(*finspacedata.ListPermissionGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListPermissionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListPermissionGroupsOutput{}
	fnCacher := func(out *finspacedata.ListPermissionGroupsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListPermissionGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListPermissionGroupsPagesWithContext(ctx context.Context, input *finspacedata.ListPermissionGroupsInput, fn func(*finspacedata.ListPermissionGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListPermissionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListPermissionGroupsOutput{}
	fnCacher := func(out *finspacedata.ListPermissionGroupsOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListPermissionGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListPermissionGroupsWithContext(ctx context.Context, input *finspacedata.ListPermissionGroupsInput, opts ...request.Option) (*finspacedata.ListPermissionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListPermissionGroupsOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListPermissionGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListUsers(input *finspacedata.ListUsersInput) (*finspacedata.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListUsersOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListUsersByPermissionGroup(input *finspacedata.ListUsersByPermissionGroupInput) (*finspacedata.ListUsersByPermissionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListUsersByPermissionGroupOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListUsersByPermissionGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListUsersByPermissionGroupWithContext(ctx context.Context, input *finspacedata.ListUsersByPermissionGroupInput, opts ...request.Option) (*finspacedata.ListUsersByPermissionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListUsersByPermissionGroupOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListUsersByPermissionGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FinSpaceData) ListUsersPages(input *finspacedata.ListUsersInput, fn func(*finspacedata.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListUsersOutput{}
	fnCacher := func(out *finspacedata.ListUsersOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListUsersPagesWithContext(ctx context.Context, input *finspacedata.ListUsersInput, fn func(*finspacedata.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*finspacedata.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &finspacedata.ListUsersOutput{}
	fnCacher := func(out *finspacedata.ListUsersOutput, more bool) bool {
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
	if err := c.FinSpaceDataAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FinSpaceData) ListUsersWithContext(ctx context.Context, input *finspacedata.ListUsersInput, opts ...request.Option) (*finspacedata.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*finspacedata.ListUsersOutput), nil
	}
	out, err := c.FinSpaceDataAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LakeFormation struct {
	lakeformationiface.LakeFormationAPI
	cache *cache.Cache
}

func NewLakeFormation(lakeformationapi lakeformationiface.LakeFormationAPI) *LakeFormation {
	return &LakeFormation{
		LakeFormationAPI: lakeformationapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LakeFormation) BatchGrantPermissions(input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.BatchGrantPermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.BatchGrantPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) BatchGrantPermissionsWithContext(ctx context.Context, input *lakeformation.BatchGrantPermissionsInput, opts ...request.Option) (*lakeformation.BatchGrantPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.BatchGrantPermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.BatchGrantPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) BatchRevokePermissions(input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.BatchRevokePermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.BatchRevokePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) BatchRevokePermissionsWithContext(ctx context.Context, input *lakeformation.BatchRevokePermissionsInput, opts ...request.Option) (*lakeformation.BatchRevokePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.BatchRevokePermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.BatchRevokePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) DescribeResource(input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.DescribeResourceOutput), nil
	}
	out, err := c.LakeFormationAPI.DescribeResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) DescribeResourceWithContext(ctx context.Context, input *lakeformation.DescribeResourceInput, opts ...request.Option) (*lakeformation.DescribeResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.DescribeResourceOutput), nil
	}
	out, err := c.LakeFormationAPI.DescribeResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) DescribeTransaction(input *lakeformation.DescribeTransactionInput) (*lakeformation.DescribeTransactionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.DescribeTransactionOutput), nil
	}
	out, err := c.LakeFormationAPI.DescribeTransaction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) DescribeTransactionWithContext(ctx context.Context, input *lakeformation.DescribeTransactionInput, opts ...request.Option) (*lakeformation.DescribeTransactionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.DescribeTransactionOutput), nil
	}
	out, err := c.LakeFormationAPI.DescribeTransactionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetDataLakeSettings(input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetDataLakeSettingsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetDataLakeSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetDataLakeSettingsWithContext(ctx context.Context, input *lakeformation.GetDataLakeSettingsInput, opts ...request.Option) (*lakeformation.GetDataLakeSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetDataLakeSettingsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetDataLakeSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetEffectivePermissionsForPath(input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetEffectivePermissionsForPathOutput), nil
	}
	out, err := c.LakeFormationAPI.GetEffectivePermissionsForPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetEffectivePermissionsForPathPages(input *lakeformation.GetEffectivePermissionsForPathInput, fn func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetEffectivePermissionsForPathOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetEffectivePermissionsForPathOutput{}
	fnCacher := func(out *lakeformation.GetEffectivePermissionsForPathOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetEffectivePermissionsForPathPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetEffectivePermissionsForPathPagesWithContext(ctx context.Context, input *lakeformation.GetEffectivePermissionsForPathInput, fn func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetEffectivePermissionsForPathOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetEffectivePermissionsForPathOutput{}
	fnCacher := func(out *lakeformation.GetEffectivePermissionsForPathOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetEffectivePermissionsForPathPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetEffectivePermissionsForPathWithContext(ctx context.Context, input *lakeformation.GetEffectivePermissionsForPathInput, opts ...request.Option) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetEffectivePermissionsForPathOutput), nil
	}
	out, err := c.LakeFormationAPI.GetEffectivePermissionsForPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetLFTag(input *lakeformation.GetLFTagInput) (*lakeformation.GetLFTagOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetLFTagOutput), nil
	}
	out, err := c.LakeFormationAPI.GetLFTag(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetLFTagWithContext(ctx context.Context, input *lakeformation.GetLFTagInput, opts ...request.Option) (*lakeformation.GetLFTagOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetLFTagOutput), nil
	}
	out, err := c.LakeFormationAPI.GetLFTagWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetQueryState(input *lakeformation.GetQueryStateInput) (*lakeformation.GetQueryStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetQueryStateOutput), nil
	}
	out, err := c.LakeFormationAPI.GetQueryState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetQueryStateWithContext(ctx context.Context, input *lakeformation.GetQueryStateInput, opts ...request.Option) (*lakeformation.GetQueryStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetQueryStateOutput), nil
	}
	out, err := c.LakeFormationAPI.GetQueryStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetQueryStatistics(input *lakeformation.GetQueryStatisticsInput) (*lakeformation.GetQueryStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetQueryStatisticsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetQueryStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetQueryStatisticsWithContext(ctx context.Context, input *lakeformation.GetQueryStatisticsInput, opts ...request.Option) (*lakeformation.GetQueryStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetQueryStatisticsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetQueryStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetResourceLFTags(input *lakeformation.GetResourceLFTagsInput) (*lakeformation.GetResourceLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetResourceLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetResourceLFTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetResourceLFTagsWithContext(ctx context.Context, input *lakeformation.GetResourceLFTagsInput, opts ...request.Option) (*lakeformation.GetResourceLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetResourceLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetResourceLFTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTableObjects(input *lakeformation.GetTableObjectsInput) (*lakeformation.GetTableObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTableObjectsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTableObjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTableObjectsPages(input *lakeformation.GetTableObjectsInput, fn func(*lakeformation.GetTableObjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetTableObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetTableObjectsOutput{}
	fnCacher := func(out *lakeformation.GetTableObjectsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetTableObjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetTableObjectsPagesWithContext(ctx context.Context, input *lakeformation.GetTableObjectsInput, fn func(*lakeformation.GetTableObjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetTableObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetTableObjectsOutput{}
	fnCacher := func(out *lakeformation.GetTableObjectsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetTableObjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetTableObjectsWithContext(ctx context.Context, input *lakeformation.GetTableObjectsInput, opts ...request.Option) (*lakeformation.GetTableObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTableObjectsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTableObjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTemporaryGluePartitionCredentials(input *lakeformation.GetTemporaryGluePartitionCredentialsInput) (*lakeformation.GetTemporaryGluePartitionCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTemporaryGluePartitionCredentialsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTemporaryGluePartitionCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTemporaryGluePartitionCredentialsWithContext(ctx context.Context, input *lakeformation.GetTemporaryGluePartitionCredentialsInput, opts ...request.Option) (*lakeformation.GetTemporaryGluePartitionCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTemporaryGluePartitionCredentialsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTemporaryGluePartitionCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTemporaryGlueTableCredentials(input *lakeformation.GetTemporaryGlueTableCredentialsInput) (*lakeformation.GetTemporaryGlueTableCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTemporaryGlueTableCredentialsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTemporaryGlueTableCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetTemporaryGlueTableCredentialsWithContext(ctx context.Context, input *lakeformation.GetTemporaryGlueTableCredentialsInput, opts ...request.Option) (*lakeformation.GetTemporaryGlueTableCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetTemporaryGlueTableCredentialsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetTemporaryGlueTableCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetWorkUnitResults(input *lakeformation.GetWorkUnitResultsInput) (*lakeformation.GetWorkUnitResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetWorkUnitResultsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetWorkUnitResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetWorkUnitResultsWithContext(ctx context.Context, input *lakeformation.GetWorkUnitResultsInput, opts ...request.Option) (*lakeformation.GetWorkUnitResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetWorkUnitResultsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetWorkUnitResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetWorkUnits(input *lakeformation.GetWorkUnitsInput) (*lakeformation.GetWorkUnitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetWorkUnitsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetWorkUnits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) GetWorkUnitsPages(input *lakeformation.GetWorkUnitsInput, fn func(*lakeformation.GetWorkUnitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetWorkUnitsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetWorkUnitsOutput{}
	fnCacher := func(out *lakeformation.GetWorkUnitsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetWorkUnitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetWorkUnitsPagesWithContext(ctx context.Context, input *lakeformation.GetWorkUnitsInput, fn func(*lakeformation.GetWorkUnitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.GetWorkUnitsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.GetWorkUnitsOutput{}
	fnCacher := func(out *lakeformation.GetWorkUnitsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.GetWorkUnitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) GetWorkUnitsWithContext(ctx context.Context, input *lakeformation.GetWorkUnitsInput, opts ...request.Option) (*lakeformation.GetWorkUnitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.GetWorkUnitsOutput), nil
	}
	out, err := c.LakeFormationAPI.GetWorkUnitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListDataCellsFilter(input *lakeformation.ListDataCellsFilterInput) (*lakeformation.ListDataCellsFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListDataCellsFilterOutput), nil
	}
	out, err := c.LakeFormationAPI.ListDataCellsFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListDataCellsFilterPages(input *lakeformation.ListDataCellsFilterInput, fn func(*lakeformation.ListDataCellsFilterOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListDataCellsFilterOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListDataCellsFilterOutput{}
	fnCacher := func(out *lakeformation.ListDataCellsFilterOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListDataCellsFilterPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListDataCellsFilterPagesWithContext(ctx context.Context, input *lakeformation.ListDataCellsFilterInput, fn func(*lakeformation.ListDataCellsFilterOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListDataCellsFilterOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListDataCellsFilterOutput{}
	fnCacher := func(out *lakeformation.ListDataCellsFilterOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListDataCellsFilterPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListDataCellsFilterWithContext(ctx context.Context, input *lakeformation.ListDataCellsFilterInput, opts ...request.Option) (*lakeformation.ListDataCellsFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListDataCellsFilterOutput), nil
	}
	out, err := c.LakeFormationAPI.ListDataCellsFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListLFTags(input *lakeformation.ListLFTagsInput) (*lakeformation.ListLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListLFTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListLFTagsPages(input *lakeformation.ListLFTagsInput, fn func(*lakeformation.ListLFTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListLFTagsOutput{}
	fnCacher := func(out *lakeformation.ListLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListLFTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListLFTagsPagesWithContext(ctx context.Context, input *lakeformation.ListLFTagsInput, fn func(*lakeformation.ListLFTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListLFTagsOutput{}
	fnCacher := func(out *lakeformation.ListLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListLFTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListLFTagsWithContext(ctx context.Context, input *lakeformation.ListLFTagsInput, opts ...request.Option) (*lakeformation.ListLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListLFTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListPermissions(input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListPermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListPermissionsPages(input *lakeformation.ListPermissionsInput, fn func(*lakeformation.ListPermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListPermissionsOutput{}
	fnCacher := func(out *lakeformation.ListPermissionsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListPermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListPermissionsPagesWithContext(ctx context.Context, input *lakeformation.ListPermissionsInput, fn func(*lakeformation.ListPermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListPermissionsOutput{}
	fnCacher := func(out *lakeformation.ListPermissionsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListPermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListPermissionsWithContext(ctx context.Context, input *lakeformation.ListPermissionsInput, opts ...request.Option) (*lakeformation.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListPermissionsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListResources(input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListResourcesOutput), nil
	}
	out, err := c.LakeFormationAPI.ListResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListResourcesPages(input *lakeformation.ListResourcesInput, fn func(*lakeformation.ListResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListResourcesOutput{}
	fnCacher := func(out *lakeformation.ListResourcesOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListResourcesPagesWithContext(ctx context.Context, input *lakeformation.ListResourcesInput, fn func(*lakeformation.ListResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListResourcesOutput{}
	fnCacher := func(out *lakeformation.ListResourcesOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListResourcesWithContext(ctx context.Context, input *lakeformation.ListResourcesInput, opts ...request.Option) (*lakeformation.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListResourcesOutput), nil
	}
	out, err := c.LakeFormationAPI.ListResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListTableStorageOptimizers(input *lakeformation.ListTableStorageOptimizersInput) (*lakeformation.ListTableStorageOptimizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListTableStorageOptimizersOutput), nil
	}
	out, err := c.LakeFormationAPI.ListTableStorageOptimizers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListTableStorageOptimizersPages(input *lakeformation.ListTableStorageOptimizersInput, fn func(*lakeformation.ListTableStorageOptimizersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListTableStorageOptimizersOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListTableStorageOptimizersOutput{}
	fnCacher := func(out *lakeformation.ListTableStorageOptimizersOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListTableStorageOptimizersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListTableStorageOptimizersPagesWithContext(ctx context.Context, input *lakeformation.ListTableStorageOptimizersInput, fn func(*lakeformation.ListTableStorageOptimizersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListTableStorageOptimizersOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListTableStorageOptimizersOutput{}
	fnCacher := func(out *lakeformation.ListTableStorageOptimizersOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListTableStorageOptimizersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListTableStorageOptimizersWithContext(ctx context.Context, input *lakeformation.ListTableStorageOptimizersInput, opts ...request.Option) (*lakeformation.ListTableStorageOptimizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListTableStorageOptimizersOutput), nil
	}
	out, err := c.LakeFormationAPI.ListTableStorageOptimizersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListTransactions(input *lakeformation.ListTransactionsInput) (*lakeformation.ListTransactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListTransactionsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListTransactions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) ListTransactionsPages(input *lakeformation.ListTransactionsInput, fn func(*lakeformation.ListTransactionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListTransactionsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListTransactionsOutput{}
	fnCacher := func(out *lakeformation.ListTransactionsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListTransactionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListTransactionsPagesWithContext(ctx context.Context, input *lakeformation.ListTransactionsInput, fn func(*lakeformation.ListTransactionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.ListTransactionsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.ListTransactionsOutput{}
	fnCacher := func(out *lakeformation.ListTransactionsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.ListTransactionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) ListTransactionsWithContext(ctx context.Context, input *lakeformation.ListTransactionsInput, opts ...request.Option) (*lakeformation.ListTransactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.ListTransactionsOutput), nil
	}
	out, err := c.LakeFormationAPI.ListTransactionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) SearchDatabasesByLFTags(input *lakeformation.SearchDatabasesByLFTagsInput) (*lakeformation.SearchDatabasesByLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.SearchDatabasesByLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.SearchDatabasesByLFTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) SearchDatabasesByLFTagsPages(input *lakeformation.SearchDatabasesByLFTagsInput, fn func(*lakeformation.SearchDatabasesByLFTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.SearchDatabasesByLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.SearchDatabasesByLFTagsOutput{}
	fnCacher := func(out *lakeformation.SearchDatabasesByLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.SearchDatabasesByLFTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) SearchDatabasesByLFTagsPagesWithContext(ctx context.Context, input *lakeformation.SearchDatabasesByLFTagsInput, fn func(*lakeformation.SearchDatabasesByLFTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.SearchDatabasesByLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.SearchDatabasesByLFTagsOutput{}
	fnCacher := func(out *lakeformation.SearchDatabasesByLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.SearchDatabasesByLFTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) SearchDatabasesByLFTagsWithContext(ctx context.Context, input *lakeformation.SearchDatabasesByLFTagsInput, opts ...request.Option) (*lakeformation.SearchDatabasesByLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.SearchDatabasesByLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.SearchDatabasesByLFTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) SearchTablesByLFTags(input *lakeformation.SearchTablesByLFTagsInput) (*lakeformation.SearchTablesByLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.SearchTablesByLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.SearchTablesByLFTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LakeFormation) SearchTablesByLFTagsPages(input *lakeformation.SearchTablesByLFTagsInput, fn func(*lakeformation.SearchTablesByLFTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.SearchTablesByLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.SearchTablesByLFTagsOutput{}
	fnCacher := func(out *lakeformation.SearchTablesByLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.SearchTablesByLFTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) SearchTablesByLFTagsPagesWithContext(ctx context.Context, input *lakeformation.SearchTablesByLFTagsInput, fn func(*lakeformation.SearchTablesByLFTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lakeformation.SearchTablesByLFTagsOutput), false)
		return nil
	}
	cachable := true
	output := &lakeformation.SearchTablesByLFTagsOutput{}
	fnCacher := func(out *lakeformation.SearchTablesByLFTagsOutput, more bool) bool {
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
	if err := c.LakeFormationAPI.SearchTablesByLFTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LakeFormation) SearchTablesByLFTagsWithContext(ctx context.Context, input *lakeformation.SearchTablesByLFTagsInput, opts ...request.Option) (*lakeformation.SearchTablesByLFTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lakeformation.SearchTablesByLFTagsOutput), nil
	}
	out, err := c.LakeFormationAPI.SearchTablesByLFTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

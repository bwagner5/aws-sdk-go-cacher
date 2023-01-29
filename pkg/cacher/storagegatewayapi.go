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
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type StorageGateway struct {
	storagegatewayiface.StorageGatewayAPI
	cache *cache.Cache
}

func NewStorageGateway(storagegatewayapi storagegatewayiface.StorageGatewayAPI) *StorageGateway {
	return &StorageGateway{
		StorageGatewayAPI: storagegatewayapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *StorageGateway) DescribeAvailabilityMonitorTest(input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeAvailabilityMonitorTestOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeAvailabilityMonitorTest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeAvailabilityMonitorTestWithContext(ctx context.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput, opts ...request.Option) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeAvailabilityMonitorTestOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeAvailabilityMonitorTestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeBandwidthRateLimit(input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeBandwidthRateLimitOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeBandwidthRateLimit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeBandwidthRateLimitSchedule(input *storagegateway.DescribeBandwidthRateLimitScheduleInput) (*storagegateway.DescribeBandwidthRateLimitScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeBandwidthRateLimitScheduleOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeBandwidthRateLimitSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeBandwidthRateLimitScheduleWithContext(ctx context.Context, input *storagegateway.DescribeBandwidthRateLimitScheduleInput, opts ...request.Option) (*storagegateway.DescribeBandwidthRateLimitScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeBandwidthRateLimitScheduleOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeBandwidthRateLimitScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeBandwidthRateLimitWithContext(ctx context.Context, input *storagegateway.DescribeBandwidthRateLimitInput, opts ...request.Option) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeBandwidthRateLimitOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeBandwidthRateLimitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeCache(input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeCacheOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeCache(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeCacheWithContext(ctx context.Context, input *storagegateway.DescribeCacheInput, opts ...request.Option) (*storagegateway.DescribeCacheOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeCacheOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeCacheWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeCachediSCSIVolumes(input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeCachediSCSIVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeCachediSCSIVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeCachediSCSIVolumesWithContext(ctx context.Context, input *storagegateway.DescribeCachediSCSIVolumesInput, opts ...request.Option) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeCachediSCSIVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeCachediSCSIVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeChapCredentials(input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeChapCredentialsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeChapCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeChapCredentialsWithContext(ctx context.Context, input *storagegateway.DescribeChapCredentialsInput, opts ...request.Option) (*storagegateway.DescribeChapCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeChapCredentialsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeChapCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeFileSystemAssociations(input *storagegateway.DescribeFileSystemAssociationsInput) (*storagegateway.DescribeFileSystemAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeFileSystemAssociationsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeFileSystemAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeFileSystemAssociationsWithContext(ctx context.Context, input *storagegateway.DescribeFileSystemAssociationsInput, opts ...request.Option) (*storagegateway.DescribeFileSystemAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeFileSystemAssociationsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeFileSystemAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeGatewayInformation(input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeGatewayInformationOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeGatewayInformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeGatewayInformationWithContext(ctx context.Context, input *storagegateway.DescribeGatewayInformationInput, opts ...request.Option) (*storagegateway.DescribeGatewayInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeGatewayInformationOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeGatewayInformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeMaintenanceStartTime(input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeMaintenanceStartTimeOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeMaintenanceStartTime(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeMaintenanceStartTimeWithContext(ctx context.Context, input *storagegateway.DescribeMaintenanceStartTimeInput, opts ...request.Option) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeMaintenanceStartTimeOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeMaintenanceStartTimeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeNFSFileShares(input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeNFSFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeNFSFileShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeNFSFileSharesWithContext(ctx context.Context, input *storagegateway.DescribeNFSFileSharesInput, opts ...request.Option) (*storagegateway.DescribeNFSFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeNFSFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeNFSFileSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSMBFileShares(input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSMBFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSMBFileShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSMBFileSharesWithContext(ctx context.Context, input *storagegateway.DescribeSMBFileSharesInput, opts ...request.Option) (*storagegateway.DescribeSMBFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSMBFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSMBFileSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSMBSettings(input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSMBSettingsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSMBSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSMBSettingsWithContext(ctx context.Context, input *storagegateway.DescribeSMBSettingsInput, opts ...request.Option) (*storagegateway.DescribeSMBSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSMBSettingsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSMBSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSnapshotSchedule(input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSnapshotScheduleOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSnapshotSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeSnapshotScheduleWithContext(ctx context.Context, input *storagegateway.DescribeSnapshotScheduleInput, opts ...request.Option) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeSnapshotScheduleOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeSnapshotScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeStorediSCSIVolumes(input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeStorediSCSIVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeStorediSCSIVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeStorediSCSIVolumesWithContext(ctx context.Context, input *storagegateway.DescribeStorediSCSIVolumesInput, opts ...request.Option) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeStorediSCSIVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeStorediSCSIVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapeArchives(input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapeArchivesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapeArchives(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapeArchivesPages(input *storagegateway.DescribeTapeArchivesInput, fn func(*storagegateway.DescribeTapeArchivesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapeArchivesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapeArchivesOutput{}
	fnCacher := func(out *storagegateway.DescribeTapeArchivesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapeArchivesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapeArchivesPagesWithContext(ctx context.Context, input *storagegateway.DescribeTapeArchivesInput, fn func(*storagegateway.DescribeTapeArchivesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapeArchivesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapeArchivesOutput{}
	fnCacher := func(out *storagegateway.DescribeTapeArchivesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapeArchivesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapeArchivesWithContext(ctx context.Context, input *storagegateway.DescribeTapeArchivesInput, opts ...request.Option) (*storagegateway.DescribeTapeArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapeArchivesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapeArchivesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapeRecoveryPoints(input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapeRecoveryPointsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapeRecoveryPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapeRecoveryPointsPages(input *storagegateway.DescribeTapeRecoveryPointsInput, fn func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapeRecoveryPointsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapeRecoveryPointsOutput{}
	fnCacher := func(out *storagegateway.DescribeTapeRecoveryPointsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapeRecoveryPointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapeRecoveryPointsPagesWithContext(ctx context.Context, input *storagegateway.DescribeTapeRecoveryPointsInput, fn func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapeRecoveryPointsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapeRecoveryPointsOutput{}
	fnCacher := func(out *storagegateway.DescribeTapeRecoveryPointsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapeRecoveryPointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapeRecoveryPointsWithContext(ctx context.Context, input *storagegateway.DescribeTapeRecoveryPointsInput, opts ...request.Option) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapeRecoveryPointsOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapeRecoveryPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapes(input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeTapesPages(input *storagegateway.DescribeTapesInput, fn func(*storagegateway.DescribeTapesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapesOutput{}
	fnCacher := func(out *storagegateway.DescribeTapesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapesPagesWithContext(ctx context.Context, input *storagegateway.DescribeTapesInput, fn func(*storagegateway.DescribeTapesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeTapesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeTapesOutput{}
	fnCacher := func(out *storagegateway.DescribeTapesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeTapesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeTapesWithContext(ctx context.Context, input *storagegateway.DescribeTapesInput, opts ...request.Option) (*storagegateway.DescribeTapesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeTapesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeTapesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeUploadBuffer(input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeUploadBufferOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeUploadBuffer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeUploadBufferWithContext(ctx context.Context, input *storagegateway.DescribeUploadBufferInput, opts ...request.Option) (*storagegateway.DescribeUploadBufferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeUploadBufferOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeUploadBufferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeVTLDevices(input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeVTLDevicesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeVTLDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeVTLDevicesPages(input *storagegateway.DescribeVTLDevicesInput, fn func(*storagegateway.DescribeVTLDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeVTLDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeVTLDevicesOutput{}
	fnCacher := func(out *storagegateway.DescribeVTLDevicesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeVTLDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeVTLDevicesPagesWithContext(ctx context.Context, input *storagegateway.DescribeVTLDevicesInput, fn func(*storagegateway.DescribeVTLDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.DescribeVTLDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.DescribeVTLDevicesOutput{}
	fnCacher := func(out *storagegateway.DescribeVTLDevicesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.DescribeVTLDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) DescribeVTLDevicesWithContext(ctx context.Context, input *storagegateway.DescribeVTLDevicesInput, opts ...request.Option) (*storagegateway.DescribeVTLDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeVTLDevicesOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeVTLDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeWorkingStorage(input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeWorkingStorageOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeWorkingStorage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) DescribeWorkingStorageWithContext(ctx context.Context, input *storagegateway.DescribeWorkingStorageInput, opts ...request.Option) (*storagegateway.DescribeWorkingStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.DescribeWorkingStorageOutput), nil
	}
	out, err := c.StorageGatewayAPI.DescribeWorkingStorageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListAutomaticTapeCreationPolicies(input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListAutomaticTapeCreationPoliciesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListAutomaticTapeCreationPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListAutomaticTapeCreationPoliciesWithContext(ctx context.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput, opts ...request.Option) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListAutomaticTapeCreationPoliciesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListAutomaticTapeCreationPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListFileShares(input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListFileShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListFileSharesPages(input *storagegateway.ListFileSharesInput, fn func(*storagegateway.ListFileSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListFileSharesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListFileSharesOutput{}
	fnCacher := func(out *storagegateway.ListFileSharesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListFileSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListFileSharesPagesWithContext(ctx context.Context, input *storagegateway.ListFileSharesInput, fn func(*storagegateway.ListFileSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListFileSharesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListFileSharesOutput{}
	fnCacher := func(out *storagegateway.ListFileSharesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListFileSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListFileSharesWithContext(ctx context.Context, input *storagegateway.ListFileSharesInput, opts ...request.Option) (*storagegateway.ListFileSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListFileSharesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListFileSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListFileSystemAssociations(input *storagegateway.ListFileSystemAssociationsInput) (*storagegateway.ListFileSystemAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListFileSystemAssociationsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListFileSystemAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListFileSystemAssociationsPages(input *storagegateway.ListFileSystemAssociationsInput, fn func(*storagegateway.ListFileSystemAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListFileSystemAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListFileSystemAssociationsOutput{}
	fnCacher := func(out *storagegateway.ListFileSystemAssociationsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListFileSystemAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListFileSystemAssociationsPagesWithContext(ctx context.Context, input *storagegateway.ListFileSystemAssociationsInput, fn func(*storagegateway.ListFileSystemAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListFileSystemAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListFileSystemAssociationsOutput{}
	fnCacher := func(out *storagegateway.ListFileSystemAssociationsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListFileSystemAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListFileSystemAssociationsWithContext(ctx context.Context, input *storagegateway.ListFileSystemAssociationsInput, opts ...request.Option) (*storagegateway.ListFileSystemAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListFileSystemAssociationsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListFileSystemAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListGateways(input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListGatewaysOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListGatewaysPages(input *storagegateway.ListGatewaysInput, fn func(*storagegateway.ListGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListGatewaysOutput{}
	fnCacher := func(out *storagegateway.ListGatewaysOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListGatewaysPagesWithContext(ctx context.Context, input *storagegateway.ListGatewaysInput, fn func(*storagegateway.ListGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListGatewaysOutput{}
	fnCacher := func(out *storagegateway.ListGatewaysOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListGatewaysWithContext(ctx context.Context, input *storagegateway.ListGatewaysInput, opts ...request.Option) (*storagegateway.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListGatewaysOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListLocalDisks(input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListLocalDisksOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListLocalDisks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListLocalDisksWithContext(ctx context.Context, input *storagegateway.ListLocalDisksInput, opts ...request.Option) (*storagegateway.ListLocalDisksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListLocalDisksOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListLocalDisksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTagsForResource(input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTagsForResourceOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTagsForResourcePages(input *storagegateway.ListTagsForResourceInput, fn func(*storagegateway.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTagsForResourceOutput{}
	fnCacher := func(out *storagegateway.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTagsForResourcePagesWithContext(ctx context.Context, input *storagegateway.ListTagsForResourceInput, fn func(*storagegateway.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTagsForResourceOutput{}
	fnCacher := func(out *storagegateway.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTagsForResourceWithContext(ctx context.Context, input *storagegateway.ListTagsForResourceInput, opts ...request.Option) (*storagegateway.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTagsForResourceOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTapePools(input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTapePoolsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTapePools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTapePoolsPages(input *storagegateway.ListTapePoolsInput, fn func(*storagegateway.ListTapePoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTapePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTapePoolsOutput{}
	fnCacher := func(out *storagegateway.ListTapePoolsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTapePoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTapePoolsPagesWithContext(ctx context.Context, input *storagegateway.ListTapePoolsInput, fn func(*storagegateway.ListTapePoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTapePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTapePoolsOutput{}
	fnCacher := func(out *storagegateway.ListTapePoolsOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTapePoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTapePoolsWithContext(ctx context.Context, input *storagegateway.ListTapePoolsInput, opts ...request.Option) (*storagegateway.ListTapePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTapePoolsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTapePoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTapes(input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTapesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTapes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListTapesPages(input *storagegateway.ListTapesInput, fn func(*storagegateway.ListTapesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTapesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTapesOutput{}
	fnCacher := func(out *storagegateway.ListTapesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTapesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTapesPagesWithContext(ctx context.Context, input *storagegateway.ListTapesInput, fn func(*storagegateway.ListTapesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListTapesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListTapesOutput{}
	fnCacher := func(out *storagegateway.ListTapesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListTapesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListTapesWithContext(ctx context.Context, input *storagegateway.ListTapesInput, opts ...request.Option) (*storagegateway.ListTapesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListTapesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListTapesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumeInitiators(input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumeInitiatorsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumeInitiators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumeInitiatorsWithContext(ctx context.Context, input *storagegateway.ListVolumeInitiatorsInput, opts ...request.Option) (*storagegateway.ListVolumeInitiatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumeInitiatorsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumeInitiatorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumeRecoveryPoints(input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumeRecoveryPointsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumeRecoveryPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumeRecoveryPointsWithContext(ctx context.Context, input *storagegateway.ListVolumeRecoveryPointsInput, opts ...request.Option) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumeRecoveryPointsOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumeRecoveryPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumes(input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *StorageGateway) ListVolumesPages(input *storagegateway.ListVolumesInput, fn func(*storagegateway.ListVolumesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListVolumesOutput{}
	fnCacher := func(out *storagegateway.ListVolumesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListVolumesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListVolumesPagesWithContext(ctx context.Context, input *storagegateway.ListVolumesInput, fn func(*storagegateway.ListVolumesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*storagegateway.ListVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &storagegateway.ListVolumesOutput{}
	fnCacher := func(out *storagegateway.ListVolumesOutput, more bool) bool {
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
	if err := c.StorageGatewayAPI.ListVolumesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *StorageGateway) ListVolumesWithContext(ctx context.Context, input *storagegateway.ListVolumesInput, opts ...request.Option) (*storagegateway.ListVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*storagegateway.ListVolumesOutput), nil
	}
	out, err := c.StorageGatewayAPI.ListVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

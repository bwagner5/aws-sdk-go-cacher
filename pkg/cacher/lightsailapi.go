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
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Lightsail struct {
	lightsailiface.LightsailAPI
	cache *cache.Cache
}

func NewLightsail(lightsailapi lightsailiface.LightsailAPI) *Lightsail {
	return &Lightsail{
		LightsailAPI: lightsailapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Lightsail) GetActiveNames(input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetActiveNamesOutput), nil
	}
	out, err := c.LightsailAPI.GetActiveNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetActiveNamesWithContext(ctx context.Context, input *lightsail.GetActiveNamesInput, opts ...request.Option) (*lightsail.GetActiveNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetActiveNamesOutput), nil
	}
	out, err := c.LightsailAPI.GetActiveNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetAlarms(input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetAlarmsOutput), nil
	}
	out, err := c.LightsailAPI.GetAlarms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetAlarmsWithContext(ctx context.Context, input *lightsail.GetAlarmsInput, opts ...request.Option) (*lightsail.GetAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetAlarmsOutput), nil
	}
	out, err := c.LightsailAPI.GetAlarmsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetAutoSnapshots(input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetAutoSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetAutoSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetAutoSnapshotsWithContext(ctx context.Context, input *lightsail.GetAutoSnapshotsInput, opts ...request.Option) (*lightsail.GetAutoSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetAutoSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetAutoSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBlueprints(input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBlueprintsOutput), nil
	}
	out, err := c.LightsailAPI.GetBlueprints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBlueprintsWithContext(ctx context.Context, input *lightsail.GetBlueprintsInput, opts ...request.Option) (*lightsail.GetBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBlueprintsOutput), nil
	}
	out, err := c.LightsailAPI.GetBlueprintsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketAccessKeys(input *lightsail.GetBucketAccessKeysInput) (*lightsail.GetBucketAccessKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketAccessKeysOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketAccessKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketAccessKeysWithContext(ctx context.Context, input *lightsail.GetBucketAccessKeysInput, opts ...request.Option) (*lightsail.GetBucketAccessKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketAccessKeysOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketAccessKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketBundles(input *lightsail.GetBucketBundlesInput) (*lightsail.GetBucketBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketBundles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketBundlesWithContext(ctx context.Context, input *lightsail.GetBucketBundlesInput, opts ...request.Option) (*lightsail.GetBucketBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketBundlesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketMetricData(input *lightsail.GetBucketMetricDataInput) (*lightsail.GetBucketMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketMetricDataWithContext(ctx context.Context, input *lightsail.GetBucketMetricDataInput, opts ...request.Option) (*lightsail.GetBucketMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBuckets(input *lightsail.GetBucketsInput) (*lightsail.GetBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketsOutput), nil
	}
	out, err := c.LightsailAPI.GetBuckets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBucketsWithContext(ctx context.Context, input *lightsail.GetBucketsInput, opts ...request.Option) (*lightsail.GetBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBucketsOutput), nil
	}
	out, err := c.LightsailAPI.GetBucketsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBundles(input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetBundles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetBundlesWithContext(ctx context.Context, input *lightsail.GetBundlesInput, opts ...request.Option) (*lightsail.GetBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetBundlesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetCertificates(input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetCertificatesOutput), nil
	}
	out, err := c.LightsailAPI.GetCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetCertificatesWithContext(ctx context.Context, input *lightsail.GetCertificatesInput, opts ...request.Option) (*lightsail.GetCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetCertificatesOutput), nil
	}
	out, err := c.LightsailAPI.GetCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetCloudFormationStackRecords(input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetCloudFormationStackRecordsOutput), nil
	}
	out, err := c.LightsailAPI.GetCloudFormationStackRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetCloudFormationStackRecordsWithContext(ctx context.Context, input *lightsail.GetCloudFormationStackRecordsInput, opts ...request.Option) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetCloudFormationStackRecordsOutput), nil
	}
	out, err := c.LightsailAPI.GetCloudFormationStackRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContactMethods(input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContactMethodsOutput), nil
	}
	out, err := c.LightsailAPI.GetContactMethods(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContactMethodsWithContext(ctx context.Context, input *lightsail.GetContactMethodsInput, opts ...request.Option) (*lightsail.GetContactMethodsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContactMethodsOutput), nil
	}
	out, err := c.LightsailAPI.GetContactMethodsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerAPIMetadata(input *lightsail.GetContainerAPIMetadataInput) (*lightsail.GetContainerAPIMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerAPIMetadataOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerAPIMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerAPIMetadataWithContext(ctx context.Context, input *lightsail.GetContainerAPIMetadataInput, opts ...request.Option) (*lightsail.GetContainerAPIMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerAPIMetadataOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerAPIMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerImages(input *lightsail.GetContainerImagesInput) (*lightsail.GetContainerImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerImagesOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerImagesWithContext(ctx context.Context, input *lightsail.GetContainerImagesInput, opts ...request.Option) (*lightsail.GetContainerImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerImagesOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerLog(input *lightsail.GetContainerLogInput) (*lightsail.GetContainerLogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerLogOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerLog(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerLogWithContext(ctx context.Context, input *lightsail.GetContainerLogInput, opts ...request.Option) (*lightsail.GetContainerLogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerLogOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerLogWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServiceDeployments(input *lightsail.GetContainerServiceDeploymentsInput) (*lightsail.GetContainerServiceDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServiceDeploymentsOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServiceDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServiceDeploymentsWithContext(ctx context.Context, input *lightsail.GetContainerServiceDeploymentsInput, opts ...request.Option) (*lightsail.GetContainerServiceDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServiceDeploymentsOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServiceDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServiceMetricData(input *lightsail.GetContainerServiceMetricDataInput) (*lightsail.GetContainerServiceMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServiceMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServiceMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServiceMetricDataWithContext(ctx context.Context, input *lightsail.GetContainerServiceMetricDataInput, opts ...request.Option) (*lightsail.GetContainerServiceMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServiceMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServiceMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServicePowers(input *lightsail.GetContainerServicePowersInput) (*lightsail.GetContainerServicePowersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServicePowersOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServicePowers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServicePowersWithContext(ctx context.Context, input *lightsail.GetContainerServicePowersInput, opts ...request.Option) (*lightsail.GetContainerServicePowersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServicePowersOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServicePowersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServices(input *lightsail.GetContainerServicesInput) (*lightsail.GetContainerServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServicesOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetContainerServicesWithContext(ctx context.Context, input *lightsail.GetContainerServicesInput, opts ...request.Option) (*lightsail.GetContainerServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetContainerServicesOutput), nil
	}
	out, err := c.LightsailAPI.GetContainerServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDisk(input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskOutput), nil
	}
	out, err := c.LightsailAPI.GetDisk(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDiskSnapshot(input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetDiskSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDiskSnapshotWithContext(ctx context.Context, input *lightsail.GetDiskSnapshotInput, opts ...request.Option) (*lightsail.GetDiskSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetDiskSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDiskSnapshots(input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetDiskSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDiskSnapshotsWithContext(ctx context.Context, input *lightsail.GetDiskSnapshotsInput, opts ...request.Option) (*lightsail.GetDiskSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetDiskSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDiskWithContext(ctx context.Context, input *lightsail.GetDiskInput, opts ...request.Option) (*lightsail.GetDiskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDiskOutput), nil
	}
	out, err := c.LightsailAPI.GetDiskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDisks(input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDisksOutput), nil
	}
	out, err := c.LightsailAPI.GetDisks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDisksWithContext(ctx context.Context, input *lightsail.GetDisksInput, opts ...request.Option) (*lightsail.GetDisksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDisksOutput), nil
	}
	out, err := c.LightsailAPI.GetDisksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionBundles(input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionBundles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionBundlesWithContext(ctx context.Context, input *lightsail.GetDistributionBundlesInput, opts ...request.Option) (*lightsail.GetDistributionBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionBundlesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionLatestCacheReset(input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionLatestCacheResetOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionLatestCacheReset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionLatestCacheResetWithContext(ctx context.Context, input *lightsail.GetDistributionLatestCacheResetInput, opts ...request.Option) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionLatestCacheResetOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionLatestCacheResetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionMetricData(input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionMetricDataWithContext(ctx context.Context, input *lightsail.GetDistributionMetricDataInput, opts ...request.Option) (*lightsail.GetDistributionMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributions(input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionsOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDistributionsWithContext(ctx context.Context, input *lightsail.GetDistributionsInput, opts ...request.Option) (*lightsail.GetDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDistributionsOutput), nil
	}
	out, err := c.LightsailAPI.GetDistributionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDomain(input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDomainOutput), nil
	}
	out, err := c.LightsailAPI.GetDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDomainWithContext(ctx context.Context, input *lightsail.GetDomainInput, opts ...request.Option) (*lightsail.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDomainOutput), nil
	}
	out, err := c.LightsailAPI.GetDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDomains(input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDomainsOutput), nil
	}
	out, err := c.LightsailAPI.GetDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetDomainsWithContext(ctx context.Context, input *lightsail.GetDomainsInput, opts ...request.Option) (*lightsail.GetDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetDomainsOutput), nil
	}
	out, err := c.LightsailAPI.GetDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetExportSnapshotRecords(input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetExportSnapshotRecordsOutput), nil
	}
	out, err := c.LightsailAPI.GetExportSnapshotRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetExportSnapshotRecordsWithContext(ctx context.Context, input *lightsail.GetExportSnapshotRecordsInput, opts ...request.Option) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetExportSnapshotRecordsOutput), nil
	}
	out, err := c.LightsailAPI.GetExportSnapshotRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstance(input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceOutput), nil
	}
	out, err := c.LightsailAPI.GetInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceAccessDetails(input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceAccessDetailsOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceAccessDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceAccessDetailsWithContext(ctx context.Context, input *lightsail.GetInstanceAccessDetailsInput, opts ...request.Option) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceAccessDetailsOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceAccessDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceMetricData(input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceMetricDataWithContext(ctx context.Context, input *lightsail.GetInstanceMetricDataInput, opts ...request.Option) (*lightsail.GetInstanceMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstancePortStates(input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstancePortStatesOutput), nil
	}
	out, err := c.LightsailAPI.GetInstancePortStates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstancePortStatesWithContext(ctx context.Context, input *lightsail.GetInstancePortStatesInput, opts ...request.Option) (*lightsail.GetInstancePortStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstancePortStatesOutput), nil
	}
	out, err := c.LightsailAPI.GetInstancePortStatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceSnapshot(input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceSnapshotWithContext(ctx context.Context, input *lightsail.GetInstanceSnapshotInput, opts ...request.Option) (*lightsail.GetInstanceSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceSnapshots(input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceSnapshotsWithContext(ctx context.Context, input *lightsail.GetInstanceSnapshotsInput, opts ...request.Option) (*lightsail.GetInstanceSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceState(input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceStateOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceStateWithContext(ctx context.Context, input *lightsail.GetInstanceStateInput, opts ...request.Option) (*lightsail.GetInstanceStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceStateOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstanceWithContext(ctx context.Context, input *lightsail.GetInstanceInput, opts ...request.Option) (*lightsail.GetInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstanceOutput), nil
	}
	out, err := c.LightsailAPI.GetInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstances(input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstancesOutput), nil
	}
	out, err := c.LightsailAPI.GetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetInstancesWithContext(ctx context.Context, input *lightsail.GetInstancesInput, opts ...request.Option) (*lightsail.GetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetInstancesOutput), nil
	}
	out, err := c.LightsailAPI.GetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetKeyPair(input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetKeyPairOutput), nil
	}
	out, err := c.LightsailAPI.GetKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetKeyPairWithContext(ctx context.Context, input *lightsail.GetKeyPairInput, opts ...request.Option) (*lightsail.GetKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetKeyPairOutput), nil
	}
	out, err := c.LightsailAPI.GetKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetKeyPairs(input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetKeyPairsOutput), nil
	}
	out, err := c.LightsailAPI.GetKeyPairs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetKeyPairsWithContext(ctx context.Context, input *lightsail.GetKeyPairsInput, opts ...request.Option) (*lightsail.GetKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetKeyPairsOutput), nil
	}
	out, err := c.LightsailAPI.GetKeyPairsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancer(input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerMetricData(input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerMetricDataWithContext(ctx context.Context, input *lightsail.GetLoadBalancerMetricDataInput, opts ...request.Option) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerTlsCertificates(input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerTlsCertificatesOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerTlsCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerTlsCertificatesWithContext(ctx context.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput, opts ...request.Option) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerTlsCertificatesOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerTlsCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerTlsPolicies(input *lightsail.GetLoadBalancerTlsPoliciesInput) (*lightsail.GetLoadBalancerTlsPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerTlsPoliciesOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerTlsPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerTlsPoliciesWithContext(ctx context.Context, input *lightsail.GetLoadBalancerTlsPoliciesInput, opts ...request.Option) (*lightsail.GetLoadBalancerTlsPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerTlsPoliciesOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerTlsPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancerWithContext(ctx context.Context, input *lightsail.GetLoadBalancerInput, opts ...request.Option) (*lightsail.GetLoadBalancerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancerOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancers(input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancersOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetLoadBalancersWithContext(ctx context.Context, input *lightsail.GetLoadBalancersInput, opts ...request.Option) (*lightsail.GetLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetLoadBalancersOutput), nil
	}
	out, err := c.LightsailAPI.GetLoadBalancersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperation(input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationOutput), nil
	}
	out, err := c.LightsailAPI.GetOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperationWithContext(ctx context.Context, input *lightsail.GetOperationInput, opts ...request.Option) (*lightsail.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationOutput), nil
	}
	out, err := c.LightsailAPI.GetOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperations(input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationsOutput), nil
	}
	out, err := c.LightsailAPI.GetOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperationsForResource(input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationsForResourceOutput), nil
	}
	out, err := c.LightsailAPI.GetOperationsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperationsForResourceWithContext(ctx context.Context, input *lightsail.GetOperationsForResourceInput, opts ...request.Option) (*lightsail.GetOperationsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationsForResourceOutput), nil
	}
	out, err := c.LightsailAPI.GetOperationsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetOperationsWithContext(ctx context.Context, input *lightsail.GetOperationsInput, opts ...request.Option) (*lightsail.GetOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetOperationsOutput), nil
	}
	out, err := c.LightsailAPI.GetOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRegions(input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRegionsOutput), nil
	}
	out, err := c.LightsailAPI.GetRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRegionsWithContext(ctx context.Context, input *lightsail.GetRegionsInput, opts ...request.Option) (*lightsail.GetRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRegionsOutput), nil
	}
	out, err := c.LightsailAPI.GetRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabase(input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseBlueprints(input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseBlueprintsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseBlueprints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseBlueprintsWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseBlueprintsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseBlueprintsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseBundles(input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseBundles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseBundlesWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseBundlesInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseBundlesOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseBundlesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseEvents(input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseEventsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseEventsWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseEventsInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseEventsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseLogEvents(input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseLogEventsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseLogEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseLogEventsWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseLogEventsInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseLogEventsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseLogEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseLogStreams(input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseLogStreamsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseLogStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseLogStreamsWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseLogStreamsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseLogStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseMasterUserPassword(input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseMasterUserPasswordOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseMasterUserPassword(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseMasterUserPasswordWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseMasterUserPasswordOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseMasterUserPasswordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseMetricData(input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseMetricDataWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseMetricDataInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseMetricDataOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseParameters(input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseParametersOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseParametersWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseParametersInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseParametersOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseSnapshot(input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseSnapshotWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseSnapshotOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseSnapshots(input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseSnapshotsWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseSnapshotsOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabaseWithContext(ctx context.Context, input *lightsail.GetRelationalDatabaseInput, opts ...request.Option) (*lightsail.GetRelationalDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabaseOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabases(input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabasesOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetRelationalDatabasesWithContext(ctx context.Context, input *lightsail.GetRelationalDatabasesInput, opts ...request.Option) (*lightsail.GetRelationalDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetRelationalDatabasesOutput), nil
	}
	out, err := c.LightsailAPI.GetRelationalDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetStaticIp(input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetStaticIpOutput), nil
	}
	out, err := c.LightsailAPI.GetStaticIp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetStaticIpWithContext(ctx context.Context, input *lightsail.GetStaticIpInput, opts ...request.Option) (*lightsail.GetStaticIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetStaticIpOutput), nil
	}
	out, err := c.LightsailAPI.GetStaticIpWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetStaticIps(input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetStaticIpsOutput), nil
	}
	out, err := c.LightsailAPI.GetStaticIps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lightsail) GetStaticIpsWithContext(ctx context.Context, input *lightsail.GetStaticIpsInput, opts ...request.Option) (*lightsail.GetStaticIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lightsail.GetStaticIpsOutput), nil
	}
	out, err := c.LightsailAPI.GetStaticIpsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

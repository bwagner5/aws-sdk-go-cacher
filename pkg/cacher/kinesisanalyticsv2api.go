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
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2/kinesisanalyticsv2iface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisAnalyticsV2 struct {
	kinesisanalyticsv2iface.KinesisAnalyticsV2API
	cache *cache.Cache
}

func NewKinesisAnalyticsV2(kinesisanalyticsv2api kinesisanalyticsv2iface.KinesisAnalyticsV2API) *KinesisAnalyticsV2 {
	return &KinesisAnalyticsV2{
		KinesisAnalyticsV2API: kinesisanalyticsv2api,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisAnalyticsV2) DescribeApplication(input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) DescribeApplicationSnapshot(input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationSnapshotOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplicationSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) DescribeApplicationSnapshotWithContext(ctx context.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput, opts ...request.Option) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationSnapshotOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplicationSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) DescribeApplicationVersion(input *kinesisanalyticsv2.DescribeApplicationVersionInput) (*kinesisanalyticsv2.DescribeApplicationVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationVersionOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplicationVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) DescribeApplicationVersionWithContext(ctx context.Context, input *kinesisanalyticsv2.DescribeApplicationVersionInput, opts ...request.Option) (*kinesisanalyticsv2.DescribeApplicationVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationVersionOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplicationVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) DescribeApplicationWithContext(ctx context.Context, input *kinesisanalyticsv2.DescribeApplicationInput, opts ...request.Option) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.DescribeApplicationOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.DescribeApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplicationSnapshots(input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationSnapshotsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplicationSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplicationSnapshotsWithContext(ctx context.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput, opts ...request.Option) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationSnapshotsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplicationSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplicationVersions(input *kinesisanalyticsv2.ListApplicationVersionsInput) (*kinesisanalyticsv2.ListApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationVersionsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplicationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplicationVersionsWithContext(ctx context.Context, input *kinesisanalyticsv2.ListApplicationVersionsInput, opts ...request.Option) (*kinesisanalyticsv2.ListApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationVersionsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplicationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplications(input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListApplicationsWithContext(ctx context.Context, input *kinesisanalyticsv2.ListApplicationsInput, opts ...request.Option) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListApplicationsOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListTagsForResource(input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalyticsV2) ListTagsForResourceWithContext(ctx context.Context, input *kinesisanalyticsv2.ListTagsForResourceInput, opts ...request.Option) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalyticsv2.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisAnalyticsV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package snowballcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/snowball/snowballiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Snowball struct {
	snowballiface.SnowballAPI
	cache *cache.Cache
}

func New(snowballapi snowballiface.SnowballAPI) *Snowball {
	return &Snowball{
		SnowballAPI: snowballapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Snowball) DescribeAddress(input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeAddressOutput), nil
	}
	out, err := c.SnowballAPI.DescribeAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeAddressWithContext(ctx context.Context, input *snowball.DescribeAddressInput, opts ...request.Option) (*snowball.DescribeAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeAddressOutput), nil
	}
	out, err := c.SnowballAPI.DescribeAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeAddresses(input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeAddressesOutput), nil
	}
	out, err := c.SnowballAPI.DescribeAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeAddressesPages(input *snowball.DescribeAddressesInput, fn func(*snowball.DescribeAddressesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.DescribeAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.DescribeAddressesOutput{}
	fnCacher := func(out *snowball.DescribeAddressesOutput, more bool) bool {
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
	if err := c.SnowballAPI.DescribeAddressesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) DescribeAddressesPagesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, fn func(*snowball.DescribeAddressesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.DescribeAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.DescribeAddressesOutput{}
	fnCacher := func(out *snowball.DescribeAddressesOutput, more bool) bool {
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
	if err := c.SnowballAPI.DescribeAddressesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) DescribeAddressesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, opts ...request.Option) (*snowball.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeAddressesOutput), nil
	}
	out, err := c.SnowballAPI.DescribeAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeCluster(input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeClusterOutput), nil
	}
	out, err := c.SnowballAPI.DescribeCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeClusterWithContext(ctx context.Context, input *snowball.DescribeClusterInput, opts ...request.Option) (*snowball.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeClusterOutput), nil
	}
	out, err := c.SnowballAPI.DescribeClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeJob(input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeJobOutput), nil
	}
	out, err := c.SnowballAPI.DescribeJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeJobWithContext(ctx context.Context, input *snowball.DescribeJobInput, opts ...request.Option) (*snowball.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeJobOutput), nil
	}
	out, err := c.SnowballAPI.DescribeJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeReturnShippingLabel(input *snowball.DescribeReturnShippingLabelInput) (*snowball.DescribeReturnShippingLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeReturnShippingLabelOutput), nil
	}
	out, err := c.SnowballAPI.DescribeReturnShippingLabel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) DescribeReturnShippingLabelWithContext(ctx context.Context, input *snowball.DescribeReturnShippingLabelInput, opts ...request.Option) (*snowball.DescribeReturnShippingLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.DescribeReturnShippingLabelOutput), nil
	}
	out, err := c.SnowballAPI.DescribeReturnShippingLabelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetJobManifest(input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetJobManifestOutput), nil
	}
	out, err := c.SnowballAPI.GetJobManifest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetJobManifestWithContext(ctx context.Context, input *snowball.GetJobManifestInput, opts ...request.Option) (*snowball.GetJobManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetJobManifestOutput), nil
	}
	out, err := c.SnowballAPI.GetJobManifestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetJobUnlockCode(input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetJobUnlockCodeOutput), nil
	}
	out, err := c.SnowballAPI.GetJobUnlockCode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetJobUnlockCodeWithContext(ctx context.Context, input *snowball.GetJobUnlockCodeInput, opts ...request.Option) (*snowball.GetJobUnlockCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetJobUnlockCodeOutput), nil
	}
	out, err := c.SnowballAPI.GetJobUnlockCodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetSnowballUsage(input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetSnowballUsageOutput), nil
	}
	out, err := c.SnowballAPI.GetSnowballUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetSnowballUsageWithContext(ctx context.Context, input *snowball.GetSnowballUsageInput, opts ...request.Option) (*snowball.GetSnowballUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetSnowballUsageOutput), nil
	}
	out, err := c.SnowballAPI.GetSnowballUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetSoftwareUpdates(input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetSoftwareUpdatesOutput), nil
	}
	out, err := c.SnowballAPI.GetSoftwareUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) GetSoftwareUpdatesWithContext(ctx context.Context, input *snowball.GetSoftwareUpdatesInput, opts ...request.Option) (*snowball.GetSoftwareUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.GetSoftwareUpdatesOutput), nil
	}
	out, err := c.SnowballAPI.GetSoftwareUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListClusterJobs(input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListClusterJobsOutput), nil
	}
	out, err := c.SnowballAPI.ListClusterJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListClusterJobsPages(input *snowball.ListClusterJobsInput, fn func(*snowball.ListClusterJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListClusterJobsOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListClusterJobsOutput{}
	fnCacher := func(out *snowball.ListClusterJobsOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListClusterJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListClusterJobsPagesWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, fn func(*snowball.ListClusterJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListClusterJobsOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListClusterJobsOutput{}
	fnCacher := func(out *snowball.ListClusterJobsOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListClusterJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListClusterJobsWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, opts ...request.Option) (*snowball.ListClusterJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListClusterJobsOutput), nil
	}
	out, err := c.SnowballAPI.ListClusterJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListClusters(input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListClustersOutput), nil
	}
	out, err := c.SnowballAPI.ListClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListClustersPages(input *snowball.ListClustersInput, fn func(*snowball.ListClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListClustersOutput{}
	fnCacher := func(out *snowball.ListClustersOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListClustersPagesWithContext(ctx context.Context, input *snowball.ListClustersInput, fn func(*snowball.ListClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListClustersOutput{}
	fnCacher := func(out *snowball.ListClustersOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListClustersWithContext(ctx context.Context, input *snowball.ListClustersInput, opts ...request.Option) (*snowball.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListClustersOutput), nil
	}
	out, err := c.SnowballAPI.ListClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListCompatibleImages(input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListCompatibleImagesOutput), nil
	}
	out, err := c.SnowballAPI.ListCompatibleImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListCompatibleImagesPages(input *snowball.ListCompatibleImagesInput, fn func(*snowball.ListCompatibleImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListCompatibleImagesOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListCompatibleImagesOutput{}
	fnCacher := func(out *snowball.ListCompatibleImagesOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListCompatibleImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListCompatibleImagesPagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, fn func(*snowball.ListCompatibleImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListCompatibleImagesOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListCompatibleImagesOutput{}
	fnCacher := func(out *snowball.ListCompatibleImagesOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListCompatibleImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListCompatibleImagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, opts ...request.Option) (*snowball.ListCompatibleImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListCompatibleImagesOutput), nil
	}
	out, err := c.SnowballAPI.ListCompatibleImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListJobs(input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListJobsOutput), nil
	}
	out, err := c.SnowballAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListJobsPages(input *snowball.ListJobsInput, fn func(*snowball.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListJobsOutput{}
	fnCacher := func(out *snowball.ListJobsOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListJobsPagesWithContext(ctx context.Context, input *snowball.ListJobsInput, fn func(*snowball.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListJobsOutput{}
	fnCacher := func(out *snowball.ListJobsOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListJobsWithContext(ctx context.Context, input *snowball.ListJobsInput, opts ...request.Option) (*snowball.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListJobsOutput), nil
	}
	out, err := c.SnowballAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListLongTermPricing(input *snowball.ListLongTermPricingInput) (*snowball.ListLongTermPricingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListLongTermPricingOutput), nil
	}
	out, err := c.SnowballAPI.ListLongTermPricing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Snowball) ListLongTermPricingPages(input *snowball.ListLongTermPricingInput, fn func(*snowball.ListLongTermPricingOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListLongTermPricingOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListLongTermPricingOutput{}
	fnCacher := func(out *snowball.ListLongTermPricingOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListLongTermPricingPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListLongTermPricingPagesWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, fn func(*snowball.ListLongTermPricingOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowball.ListLongTermPricingOutput), false)
		return nil
	}
	cachable := true
	output := &snowball.ListLongTermPricingOutput{}
	fnCacher := func(out *snowball.ListLongTermPricingOutput, more bool) bool {
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
	if err := c.SnowballAPI.ListLongTermPricingPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Snowball) ListLongTermPricingWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, opts ...request.Option) (*snowball.ListLongTermPricingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowball.ListLongTermPricingOutput), nil
	}
	out, err := c.SnowballAPI.ListLongTermPricingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

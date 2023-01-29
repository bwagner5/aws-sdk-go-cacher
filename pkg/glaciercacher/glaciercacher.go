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
package glaciercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/glacier/glacieriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Glacier struct {
	glacieriface.GlacierAPI
	cache *cache.Cache
}

func New(glacierapi glacieriface.GlacierAPI) *Glacier {
	return &Glacier{
		GlacierAPI: glacierapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Glacier) DescribeVault(input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.DescribeVaultOutput), nil
	}
	out, err := c.GlacierAPI.DescribeVault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) DescribeVaultWithContext(ctx context.Context, input *glacier.DescribeVaultInput, opts ...request.Option) (*glacier.DescribeVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.DescribeVaultOutput), nil
	}
	out, err := c.GlacierAPI.DescribeVaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetDataRetrievalPolicy(input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetDataRetrievalPolicyOutput), nil
	}
	out, err := c.GlacierAPI.GetDataRetrievalPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetDataRetrievalPolicyWithContext(ctx context.Context, input *glacier.GetDataRetrievalPolicyInput, opts ...request.Option) (*glacier.GetDataRetrievalPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetDataRetrievalPolicyOutput), nil
	}
	out, err := c.GlacierAPI.GetDataRetrievalPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetJobOutput(input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetJobOutputOutput), nil
	}
	out, err := c.GlacierAPI.GetJobOutput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetJobOutputWithContext(ctx context.Context, input *glacier.GetJobOutputInput, opts ...request.Option) (*glacier.GetJobOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetJobOutputOutput), nil
	}
	out, err := c.GlacierAPI.GetJobOutputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultAccessPolicy(input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultAccessPolicyOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultAccessPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultAccessPolicyWithContext(ctx context.Context, input *glacier.GetVaultAccessPolicyInput, opts ...request.Option) (*glacier.GetVaultAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultAccessPolicyOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultAccessPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultLock(input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultLockOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultLock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultLockWithContext(ctx context.Context, input *glacier.GetVaultLockInput, opts ...request.Option) (*glacier.GetVaultLockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultLockOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultLockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultNotifications(input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultNotificationsOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) GetVaultNotificationsWithContext(ctx context.Context, input *glacier.GetVaultNotificationsInput, opts ...request.Option) (*glacier.GetVaultNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.GetVaultNotificationsOutput), nil
	}
	out, err := c.GlacierAPI.GetVaultNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListJobs(input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListJobsOutput), nil
	}
	out, err := c.GlacierAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListJobsPages(input *glacier.ListJobsInput, fn func(*glacier.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListJobsOutput{}
	fnCacher := func(out *glacier.ListJobsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListJobsPagesWithContext(ctx context.Context, input *glacier.ListJobsInput, fn func(*glacier.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListJobsOutput{}
	fnCacher := func(out *glacier.ListJobsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListJobsWithContext(ctx context.Context, input *glacier.ListJobsInput, opts ...request.Option) (*glacier.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListJobsOutput), nil
	}
	out, err := c.GlacierAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListMultipartUploads(input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListMultipartUploadsOutput), nil
	}
	out, err := c.GlacierAPI.ListMultipartUploads(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListMultipartUploadsPages(input *glacier.ListMultipartUploadsInput, fn func(*glacier.ListMultipartUploadsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListMultipartUploadsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListMultipartUploadsOutput{}
	fnCacher := func(out *glacier.ListMultipartUploadsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListMultipartUploadsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListMultipartUploadsPagesWithContext(ctx context.Context, input *glacier.ListMultipartUploadsInput, fn func(*glacier.ListMultipartUploadsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListMultipartUploadsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListMultipartUploadsOutput{}
	fnCacher := func(out *glacier.ListMultipartUploadsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListMultipartUploadsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListMultipartUploadsWithContext(ctx context.Context, input *glacier.ListMultipartUploadsInput, opts ...request.Option) (*glacier.ListMultipartUploadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListMultipartUploadsOutput), nil
	}
	out, err := c.GlacierAPI.ListMultipartUploadsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListParts(input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListPartsOutput), nil
	}
	out, err := c.GlacierAPI.ListParts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListPartsPages(input *glacier.ListPartsInput, fn func(*glacier.ListPartsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListPartsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListPartsOutput{}
	fnCacher := func(out *glacier.ListPartsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListPartsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListPartsPagesWithContext(ctx context.Context, input *glacier.ListPartsInput, fn func(*glacier.ListPartsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListPartsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListPartsOutput{}
	fnCacher := func(out *glacier.ListPartsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListPartsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListPartsWithContext(ctx context.Context, input *glacier.ListPartsInput, opts ...request.Option) (*glacier.ListPartsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListPartsOutput), nil
	}
	out, err := c.GlacierAPI.ListPartsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListProvisionedCapacity(input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListProvisionedCapacityOutput), nil
	}
	out, err := c.GlacierAPI.ListProvisionedCapacity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListProvisionedCapacityWithContext(ctx context.Context, input *glacier.ListProvisionedCapacityInput, opts ...request.Option) (*glacier.ListProvisionedCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListProvisionedCapacityOutput), nil
	}
	out, err := c.GlacierAPI.ListProvisionedCapacityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListTagsForVault(input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListTagsForVaultOutput), nil
	}
	out, err := c.GlacierAPI.ListTagsForVault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListTagsForVaultWithContext(ctx context.Context, input *glacier.ListTagsForVaultInput, opts ...request.Option) (*glacier.ListTagsForVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListTagsForVaultOutput), nil
	}
	out, err := c.GlacierAPI.ListTagsForVaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListVaults(input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListVaultsOutput), nil
	}
	out, err := c.GlacierAPI.ListVaults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glacier) ListVaultsPages(input *glacier.ListVaultsInput, fn func(*glacier.ListVaultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListVaultsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListVaultsOutput{}
	fnCacher := func(out *glacier.ListVaultsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListVaultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListVaultsPagesWithContext(ctx context.Context, input *glacier.ListVaultsInput, fn func(*glacier.ListVaultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glacier.ListVaultsOutput), false)
		return nil
	}
	cachable := true
	output := &glacier.ListVaultsOutput{}
	fnCacher := func(out *glacier.ListVaultsOutput, more bool) bool {
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
	if err := c.GlacierAPI.ListVaultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glacier) ListVaultsWithContext(ctx context.Context, input *glacier.ListVaultsInput, opts ...request.Option) (*glacier.ListVaultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glacier.ListVaultsOutput), nil
	}
	out, err := c.GlacierAPI.ListVaultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

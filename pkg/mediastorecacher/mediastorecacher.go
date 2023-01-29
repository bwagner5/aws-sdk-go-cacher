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
package mediastorecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaStore struct {
	mediastoreiface.MediaStoreAPI
	cache *cache.Cache
}

func New(mediastoreapi mediastoreiface.MediaStoreAPI) *MediaStore {
	return &MediaStore{
		MediaStoreAPI: mediastoreapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaStore) DescribeContainer(input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.DescribeContainerOutput), nil
	}
	out, err := c.MediaStoreAPI.DescribeContainer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) DescribeContainerWithContext(ctx context.Context, input *mediastore.DescribeContainerInput, opts ...request.Option) (*mediastore.DescribeContainerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.DescribeContainerOutput), nil
	}
	out, err := c.MediaStoreAPI.DescribeContainerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetContainerPolicy(input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetContainerPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetContainerPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetContainerPolicyWithContext(ctx context.Context, input *mediastore.GetContainerPolicyInput, opts ...request.Option) (*mediastore.GetContainerPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetContainerPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetContainerPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetCorsPolicy(input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetCorsPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetCorsPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetCorsPolicyWithContext(ctx context.Context, input *mediastore.GetCorsPolicyInput, opts ...request.Option) (*mediastore.GetCorsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetCorsPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetCorsPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetLifecyclePolicy(input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetLifecyclePolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetLifecyclePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetLifecyclePolicyWithContext(ctx context.Context, input *mediastore.GetLifecyclePolicyInput, opts ...request.Option) (*mediastore.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetLifecyclePolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetLifecyclePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetMetricPolicy(input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetMetricPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetMetricPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) GetMetricPolicyWithContext(ctx context.Context, input *mediastore.GetMetricPolicyInput, opts ...request.Option) (*mediastore.GetMetricPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.GetMetricPolicyOutput), nil
	}
	out, err := c.MediaStoreAPI.GetMetricPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) ListContainers(input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.ListContainersOutput), nil
	}
	out, err := c.MediaStoreAPI.ListContainers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) ListContainersPages(input *mediastore.ListContainersInput, fn func(*mediastore.ListContainersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediastore.ListContainersOutput), false)
		return nil
	}
	cachable := true
	output := &mediastore.ListContainersOutput{}
	fnCacher := func(out *mediastore.ListContainersOutput, more bool) bool {
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
	if err := c.MediaStoreAPI.ListContainersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaStore) ListContainersPagesWithContext(ctx context.Context, input *mediastore.ListContainersInput, fn func(*mediastore.ListContainersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediastore.ListContainersOutput), false)
		return nil
	}
	cachable := true
	output := &mediastore.ListContainersOutput{}
	fnCacher := func(out *mediastore.ListContainersOutput, more bool) bool {
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
	if err := c.MediaStoreAPI.ListContainersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaStore) ListContainersWithContext(ctx context.Context, input *mediastore.ListContainersInput, opts ...request.Option) (*mediastore.ListContainersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.ListContainersOutput), nil
	}
	out, err := c.MediaStoreAPI.ListContainersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) ListTagsForResource(input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaStoreAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStore) ListTagsForResourceWithContext(ctx context.Context, input *mediastore.ListTagsForResourceInput, opts ...request.Option) (*mediastore.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastore.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaStoreAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

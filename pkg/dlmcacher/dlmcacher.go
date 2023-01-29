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
package dlmcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DLM struct {
	dlmiface.DLMAPI
	cache *cache.Cache
}

func New(dlmapi dlmiface.DLMAPI) *DLM {
	return &DLM{
		DLMAPI: dlmapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DLM) GetLifecyclePolicies(input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.GetLifecyclePoliciesOutput), nil
	}
	out, err := c.DLMAPI.GetLifecyclePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DLM) GetLifecyclePoliciesWithContext(ctx context.Context, input *dlm.GetLifecyclePoliciesInput, opts ...request.Option) (*dlm.GetLifecyclePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.GetLifecyclePoliciesOutput), nil
	}
	out, err := c.DLMAPI.GetLifecyclePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DLM) GetLifecyclePolicy(input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.GetLifecyclePolicyOutput), nil
	}
	out, err := c.DLMAPI.GetLifecyclePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DLM) GetLifecyclePolicyWithContext(ctx context.Context, input *dlm.GetLifecyclePolicyInput, opts ...request.Option) (*dlm.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.GetLifecyclePolicyOutput), nil
	}
	out, err := c.DLMAPI.GetLifecyclePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DLM) ListTagsForResource(input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.ListTagsForResourceOutput), nil
	}
	out, err := c.DLMAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DLM) ListTagsForResourceWithContext(ctx context.Context, input *dlm.ListTagsForResourceInput, opts ...request.Option) (*dlm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dlm.ListTagsForResourceOutput), nil
	}
	out, err := c.DLMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

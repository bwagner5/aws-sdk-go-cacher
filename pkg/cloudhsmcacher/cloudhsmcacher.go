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
package cloudhsmcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudHSM struct {
	cloudhsmiface.CloudHSMAPI
	cache *cache.Cache
}

func New(cloudhsmapi cloudhsmiface.CloudHSMAPI) *CloudHSM {
	return &CloudHSM{
		CloudHSMAPI: cloudhsmapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudHSM) DescribeHapg(input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeHapgOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeHapg(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) DescribeHapgWithContext(ctx context.Context, input *cloudhsm.DescribeHapgInput, opts ...request.Option) (*cloudhsm.DescribeHapgOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeHapgOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeHapgWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) DescribeHsm(input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeHsmOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeHsm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) DescribeHsmWithContext(ctx context.Context, input *cloudhsm.DescribeHsmInput, opts ...request.Option) (*cloudhsm.DescribeHsmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeHsmOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeHsmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) DescribeLunaClient(input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeLunaClientOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeLunaClient(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) DescribeLunaClientWithContext(ctx context.Context, input *cloudhsm.DescribeLunaClientInput, opts ...request.Option) (*cloudhsm.DescribeLunaClientOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.DescribeLunaClientOutput), nil
	}
	out, err := c.CloudHSMAPI.DescribeLunaClientWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) GetConfig(input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.GetConfigOutput), nil
	}
	out, err := c.CloudHSMAPI.GetConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) GetConfigWithContext(ctx context.Context, input *cloudhsm.GetConfigInput, opts ...request.Option) (*cloudhsm.GetConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.GetConfigOutput), nil
	}
	out, err := c.CloudHSMAPI.GetConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListAvailableZones(input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListAvailableZonesOutput), nil
	}
	out, err := c.CloudHSMAPI.ListAvailableZones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListAvailableZonesWithContext(ctx context.Context, input *cloudhsm.ListAvailableZonesInput, opts ...request.Option) (*cloudhsm.ListAvailableZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListAvailableZonesOutput), nil
	}
	out, err := c.CloudHSMAPI.ListAvailableZonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListHapgs(input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListHapgsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListHapgs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListHapgsWithContext(ctx context.Context, input *cloudhsm.ListHapgsInput, opts ...request.Option) (*cloudhsm.ListHapgsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListHapgsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListHapgsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListHsms(input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListHsmsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListHsms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListHsmsWithContext(ctx context.Context, input *cloudhsm.ListHsmsInput, opts ...request.Option) (*cloudhsm.ListHsmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListHsmsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListHsmsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListLunaClients(input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListLunaClientsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListLunaClients(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListLunaClientsWithContext(ctx context.Context, input *cloudhsm.ListLunaClientsInput, opts ...request.Option) (*cloudhsm.ListLunaClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListLunaClientsOutput), nil
	}
	out, err := c.CloudHSMAPI.ListLunaClientsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListTagsForResource(input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudHSMAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudHSM) ListTagsForResourceWithContext(ctx context.Context, input *cloudhsm.ListTagsForResourceInput, opts ...request.Option) (*cloudhsm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudhsm.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudHSMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

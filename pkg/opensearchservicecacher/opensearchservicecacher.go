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
package opensearchservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/opensearchservice"
	"github.com/aws/aws-sdk-go/service/opensearchservice/opensearchserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type OpenSearchService struct {
	opensearchserviceiface.OpenSearchServiceAPI
	cache *cache.Cache
}

func New(opensearchserviceapi opensearchserviceiface.OpenSearchServiceAPI) *OpenSearchService {
	return &OpenSearchService{
		OpenSearchServiceAPI: opensearchserviceapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *OpenSearchService) DescribeDomain(input *opensearchservice.DescribeDomainInput) (*opensearchservice.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainAutoTunes(input *opensearchservice.DescribeDomainAutoTunesInput) (*opensearchservice.DescribeDomainAutoTunesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainAutoTunesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainAutoTunes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainAutoTunesPages(input *opensearchservice.DescribeDomainAutoTunesInput, fn func(*opensearchservice.DescribeDomainAutoTunesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeDomainAutoTunesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeDomainAutoTunesOutput{}
	fnCacher := func(out *opensearchservice.DescribeDomainAutoTunesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeDomainAutoTunesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeDomainAutoTunesPagesWithContext(ctx context.Context, input *opensearchservice.DescribeDomainAutoTunesInput, fn func(*opensearchservice.DescribeDomainAutoTunesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeDomainAutoTunesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeDomainAutoTunesOutput{}
	fnCacher := func(out *opensearchservice.DescribeDomainAutoTunesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeDomainAutoTunesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeDomainAutoTunesWithContext(ctx context.Context, input *opensearchservice.DescribeDomainAutoTunesInput, opts ...request.Option) (*opensearchservice.DescribeDomainAutoTunesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainAutoTunesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainAutoTunesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainChangeProgress(input *opensearchservice.DescribeDomainChangeProgressInput) (*opensearchservice.DescribeDomainChangeProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainChangeProgressOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainChangeProgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainChangeProgressWithContext(ctx context.Context, input *opensearchservice.DescribeDomainChangeProgressInput, opts ...request.Option) (*opensearchservice.DescribeDomainChangeProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainChangeProgressOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainChangeProgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainConfig(input *opensearchservice.DescribeDomainConfigInput) (*opensearchservice.DescribeDomainConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainConfigOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainConfigWithContext(ctx context.Context, input *opensearchservice.DescribeDomainConfigInput, opts ...request.Option) (*opensearchservice.DescribeDomainConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainConfigOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainWithContext(ctx context.Context, input *opensearchservice.DescribeDomainInput, opts ...request.Option) (*opensearchservice.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomains(input *opensearchservice.DescribeDomainsInput) (*opensearchservice.DescribeDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDomainsWithContext(ctx context.Context, input *opensearchservice.DescribeDomainsInput, opts ...request.Option) (*opensearchservice.DescribeDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDomainsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDryRunProgress(input *opensearchservice.DescribeDryRunProgressInput) (*opensearchservice.DescribeDryRunProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDryRunProgressOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDryRunProgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeDryRunProgressWithContext(ctx context.Context, input *opensearchservice.DescribeDryRunProgressInput, opts ...request.Option) (*opensearchservice.DescribeDryRunProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeDryRunProgressOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeDryRunProgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeInboundConnections(input *opensearchservice.DescribeInboundConnectionsInput) (*opensearchservice.DescribeInboundConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeInboundConnectionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeInboundConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeInboundConnectionsPages(input *opensearchservice.DescribeInboundConnectionsInput, fn func(*opensearchservice.DescribeInboundConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeInboundConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeInboundConnectionsOutput{}
	fnCacher := func(out *opensearchservice.DescribeInboundConnectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeInboundConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeInboundConnectionsPagesWithContext(ctx context.Context, input *opensearchservice.DescribeInboundConnectionsInput, fn func(*opensearchservice.DescribeInboundConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeInboundConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeInboundConnectionsOutput{}
	fnCacher := func(out *opensearchservice.DescribeInboundConnectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeInboundConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeInboundConnectionsWithContext(ctx context.Context, input *opensearchservice.DescribeInboundConnectionsInput, opts ...request.Option) (*opensearchservice.DescribeInboundConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeInboundConnectionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeInboundConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeInstanceTypeLimits(input *opensearchservice.DescribeInstanceTypeLimitsInput) (*opensearchservice.DescribeInstanceTypeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeInstanceTypeLimitsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeInstanceTypeLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeInstanceTypeLimitsWithContext(ctx context.Context, input *opensearchservice.DescribeInstanceTypeLimitsInput, opts ...request.Option) (*opensearchservice.DescribeInstanceTypeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeInstanceTypeLimitsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeInstanceTypeLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeOutboundConnections(input *opensearchservice.DescribeOutboundConnectionsInput) (*opensearchservice.DescribeOutboundConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeOutboundConnectionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeOutboundConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeOutboundConnectionsPages(input *opensearchservice.DescribeOutboundConnectionsInput, fn func(*opensearchservice.DescribeOutboundConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeOutboundConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeOutboundConnectionsOutput{}
	fnCacher := func(out *opensearchservice.DescribeOutboundConnectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeOutboundConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeOutboundConnectionsPagesWithContext(ctx context.Context, input *opensearchservice.DescribeOutboundConnectionsInput, fn func(*opensearchservice.DescribeOutboundConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeOutboundConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeOutboundConnectionsOutput{}
	fnCacher := func(out *opensearchservice.DescribeOutboundConnectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeOutboundConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeOutboundConnectionsWithContext(ctx context.Context, input *opensearchservice.DescribeOutboundConnectionsInput, opts ...request.Option) (*opensearchservice.DescribeOutboundConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeOutboundConnectionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeOutboundConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribePackages(input *opensearchservice.DescribePackagesInput) (*opensearchservice.DescribePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribePackagesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribePackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribePackagesPages(input *opensearchservice.DescribePackagesInput, fn func(*opensearchservice.DescribePackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribePackagesOutput{}
	fnCacher := func(out *opensearchservice.DescribePackagesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribePackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribePackagesPagesWithContext(ctx context.Context, input *opensearchservice.DescribePackagesInput, fn func(*opensearchservice.DescribePackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribePackagesOutput{}
	fnCacher := func(out *opensearchservice.DescribePackagesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribePackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribePackagesWithContext(ctx context.Context, input *opensearchservice.DescribePackagesInput, opts ...request.Option) (*opensearchservice.DescribePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribePackagesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribePackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeReservedInstanceOfferings(input *opensearchservice.DescribeReservedInstanceOfferingsInput) (*opensearchservice.DescribeReservedInstanceOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeReservedInstanceOfferingsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeReservedInstanceOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeReservedInstanceOfferingsPages(input *opensearchservice.DescribeReservedInstanceOfferingsInput, fn func(*opensearchservice.DescribeReservedInstanceOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeReservedInstanceOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeReservedInstanceOfferingsOutput{}
	fnCacher := func(out *opensearchservice.DescribeReservedInstanceOfferingsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeReservedInstanceOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeReservedInstanceOfferingsPagesWithContext(ctx context.Context, input *opensearchservice.DescribeReservedInstanceOfferingsInput, fn func(*opensearchservice.DescribeReservedInstanceOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeReservedInstanceOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeReservedInstanceOfferingsOutput{}
	fnCacher := func(out *opensearchservice.DescribeReservedInstanceOfferingsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeReservedInstanceOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeReservedInstanceOfferingsWithContext(ctx context.Context, input *opensearchservice.DescribeReservedInstanceOfferingsInput, opts ...request.Option) (*opensearchservice.DescribeReservedInstanceOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeReservedInstanceOfferingsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeReservedInstanceOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeReservedInstances(input *opensearchservice.DescribeReservedInstancesInput) (*opensearchservice.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeReservedInstancesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeReservedInstancesPages(input *opensearchservice.DescribeReservedInstancesInput, fn func(*opensearchservice.DescribeReservedInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeReservedInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeReservedInstancesOutput{}
	fnCacher := func(out *opensearchservice.DescribeReservedInstancesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeReservedInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeReservedInstancesPagesWithContext(ctx context.Context, input *opensearchservice.DescribeReservedInstancesInput, fn func(*opensearchservice.DescribeReservedInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.DescribeReservedInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.DescribeReservedInstancesOutput{}
	fnCacher := func(out *opensearchservice.DescribeReservedInstancesOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.DescribeReservedInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) DescribeReservedInstancesWithContext(ctx context.Context, input *opensearchservice.DescribeReservedInstancesInput, opts ...request.Option) (*opensearchservice.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeReservedInstancesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeVpcEndpoints(input *opensearchservice.DescribeVpcEndpointsInput) (*opensearchservice.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) DescribeVpcEndpointsWithContext(ctx context.Context, input *opensearchservice.DescribeVpcEndpointsInput, opts ...request.Option) (*opensearchservice.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.DescribeVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetCompatibleVersions(input *opensearchservice.GetCompatibleVersionsInput) (*opensearchservice.GetCompatibleVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetCompatibleVersionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetCompatibleVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetCompatibleVersionsWithContext(ctx context.Context, input *opensearchservice.GetCompatibleVersionsInput, opts ...request.Option) (*opensearchservice.GetCompatibleVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetCompatibleVersionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetCompatibleVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetPackageVersionHistory(input *opensearchservice.GetPackageVersionHistoryInput) (*opensearchservice.GetPackageVersionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetPackageVersionHistoryOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetPackageVersionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetPackageVersionHistoryPages(input *opensearchservice.GetPackageVersionHistoryInput, fn func(*opensearchservice.GetPackageVersionHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.GetPackageVersionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.GetPackageVersionHistoryOutput{}
	fnCacher := func(out *opensearchservice.GetPackageVersionHistoryOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.GetPackageVersionHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) GetPackageVersionHistoryPagesWithContext(ctx context.Context, input *opensearchservice.GetPackageVersionHistoryInput, fn func(*opensearchservice.GetPackageVersionHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.GetPackageVersionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.GetPackageVersionHistoryOutput{}
	fnCacher := func(out *opensearchservice.GetPackageVersionHistoryOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.GetPackageVersionHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) GetPackageVersionHistoryWithContext(ctx context.Context, input *opensearchservice.GetPackageVersionHistoryInput, opts ...request.Option) (*opensearchservice.GetPackageVersionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetPackageVersionHistoryOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetPackageVersionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetUpgradeHistory(input *opensearchservice.GetUpgradeHistoryInput) (*opensearchservice.GetUpgradeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetUpgradeHistoryOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetUpgradeHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetUpgradeHistoryPages(input *opensearchservice.GetUpgradeHistoryInput, fn func(*opensearchservice.GetUpgradeHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.GetUpgradeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.GetUpgradeHistoryOutput{}
	fnCacher := func(out *opensearchservice.GetUpgradeHistoryOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.GetUpgradeHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) GetUpgradeHistoryPagesWithContext(ctx context.Context, input *opensearchservice.GetUpgradeHistoryInput, fn func(*opensearchservice.GetUpgradeHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.GetUpgradeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.GetUpgradeHistoryOutput{}
	fnCacher := func(out *opensearchservice.GetUpgradeHistoryOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.GetUpgradeHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) GetUpgradeHistoryWithContext(ctx context.Context, input *opensearchservice.GetUpgradeHistoryInput, opts ...request.Option) (*opensearchservice.GetUpgradeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetUpgradeHistoryOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetUpgradeHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetUpgradeStatus(input *opensearchservice.GetUpgradeStatusInput) (*opensearchservice.GetUpgradeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetUpgradeStatusOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetUpgradeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) GetUpgradeStatusWithContext(ctx context.Context, input *opensearchservice.GetUpgradeStatusInput, opts ...request.Option) (*opensearchservice.GetUpgradeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.GetUpgradeStatusOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.GetUpgradeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListDomainNames(input *opensearchservice.ListDomainNamesInput) (*opensearchservice.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListDomainNamesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListDomainNamesWithContext(ctx context.Context, input *opensearchservice.ListDomainNamesInput, opts ...request.Option) (*opensearchservice.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListDomainNamesOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListDomainsForPackage(input *opensearchservice.ListDomainsForPackageInput) (*opensearchservice.ListDomainsForPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListDomainsForPackageOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListDomainsForPackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListDomainsForPackagePages(input *opensearchservice.ListDomainsForPackageInput, fn func(*opensearchservice.ListDomainsForPackageOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListDomainsForPackageOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListDomainsForPackageOutput{}
	fnCacher := func(out *opensearchservice.ListDomainsForPackageOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListDomainsForPackagePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListDomainsForPackagePagesWithContext(ctx context.Context, input *opensearchservice.ListDomainsForPackageInput, fn func(*opensearchservice.ListDomainsForPackageOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListDomainsForPackageOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListDomainsForPackageOutput{}
	fnCacher := func(out *opensearchservice.ListDomainsForPackageOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListDomainsForPackagePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListDomainsForPackageWithContext(ctx context.Context, input *opensearchservice.ListDomainsForPackageInput, opts ...request.Option) (*opensearchservice.ListDomainsForPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListDomainsForPackageOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListDomainsForPackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListInstanceTypeDetails(input *opensearchservice.ListInstanceTypeDetailsInput) (*opensearchservice.ListInstanceTypeDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListInstanceTypeDetailsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListInstanceTypeDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListInstanceTypeDetailsPages(input *opensearchservice.ListInstanceTypeDetailsInput, fn func(*opensearchservice.ListInstanceTypeDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListInstanceTypeDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListInstanceTypeDetailsOutput{}
	fnCacher := func(out *opensearchservice.ListInstanceTypeDetailsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListInstanceTypeDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListInstanceTypeDetailsPagesWithContext(ctx context.Context, input *opensearchservice.ListInstanceTypeDetailsInput, fn func(*opensearchservice.ListInstanceTypeDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListInstanceTypeDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListInstanceTypeDetailsOutput{}
	fnCacher := func(out *opensearchservice.ListInstanceTypeDetailsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListInstanceTypeDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListInstanceTypeDetailsWithContext(ctx context.Context, input *opensearchservice.ListInstanceTypeDetailsInput, opts ...request.Option) (*opensearchservice.ListInstanceTypeDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListInstanceTypeDetailsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListInstanceTypeDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListPackagesForDomain(input *opensearchservice.ListPackagesForDomainInput) (*opensearchservice.ListPackagesForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListPackagesForDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListPackagesForDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListPackagesForDomainPages(input *opensearchservice.ListPackagesForDomainInput, fn func(*opensearchservice.ListPackagesForDomainOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListPackagesForDomainOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListPackagesForDomainOutput{}
	fnCacher := func(out *opensearchservice.ListPackagesForDomainOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListPackagesForDomainPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListPackagesForDomainPagesWithContext(ctx context.Context, input *opensearchservice.ListPackagesForDomainInput, fn func(*opensearchservice.ListPackagesForDomainOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListPackagesForDomainOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListPackagesForDomainOutput{}
	fnCacher := func(out *opensearchservice.ListPackagesForDomainOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListPackagesForDomainPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListPackagesForDomainWithContext(ctx context.Context, input *opensearchservice.ListPackagesForDomainInput, opts ...request.Option) (*opensearchservice.ListPackagesForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListPackagesForDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListPackagesForDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListTags(input *opensearchservice.ListTagsInput) (*opensearchservice.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListTagsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListTagsWithContext(ctx context.Context, input *opensearchservice.ListTagsInput, opts ...request.Option) (*opensearchservice.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListTagsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVersions(input *opensearchservice.ListVersionsInput) (*opensearchservice.ListVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVersionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVersionsPages(input *opensearchservice.ListVersionsInput, fn func(*opensearchservice.ListVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListVersionsOutput{}
	fnCacher := func(out *opensearchservice.ListVersionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListVersionsPagesWithContext(ctx context.Context, input *opensearchservice.ListVersionsInput, fn func(*opensearchservice.ListVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchservice.ListVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchservice.ListVersionsOutput{}
	fnCacher := func(out *opensearchservice.ListVersionsOutput, more bool) bool {
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
	if err := c.OpenSearchServiceAPI.ListVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchService) ListVersionsWithContext(ctx context.Context, input *opensearchservice.ListVersionsInput, opts ...request.Option) (*opensearchservice.ListVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVersionsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpointAccess(input *opensearchservice.ListVpcEndpointAccessInput) (*opensearchservice.ListVpcEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointAccessOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpointAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpointAccessWithContext(ctx context.Context, input *opensearchservice.ListVpcEndpointAccessInput, opts ...request.Option) (*opensearchservice.ListVpcEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointAccessOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpointAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpoints(input *opensearchservice.ListVpcEndpointsInput) (*opensearchservice.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpointsForDomain(input *opensearchservice.ListVpcEndpointsForDomainInput) (*opensearchservice.ListVpcEndpointsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointsForDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpointsForDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpointsForDomainWithContext(ctx context.Context, input *opensearchservice.ListVpcEndpointsForDomainInput, opts ...request.Option) (*opensearchservice.ListVpcEndpointsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointsForDomainOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpointsForDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchService) ListVpcEndpointsWithContext(ctx context.Context, input *opensearchservice.ListVpcEndpointsInput, opts ...request.Option) (*opensearchservice.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchservice.ListVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServiceAPI.ListVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

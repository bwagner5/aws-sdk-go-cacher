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
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ElasticsearchService struct {
	elasticsearchserviceiface.ElasticsearchServiceAPI
	cache *cache.Cache
}

func NewElasticsearchService(elasticsearchserviceapi elasticsearchserviceiface.ElasticsearchServiceAPI) *ElasticsearchService {
	return &ElasticsearchService{
		ElasticsearchServiceAPI: elasticsearchserviceapi,
		cache:                   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ElasticsearchService) DescribeDomainAutoTunes(input *elasticsearchservice.DescribeDomainAutoTunesInput) (*elasticsearchservice.DescribeDomainAutoTunesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeDomainAutoTunesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeDomainAutoTunes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeDomainAutoTunesPages(input *elasticsearchservice.DescribeDomainAutoTunesInput, fn func(*elasticsearchservice.DescribeDomainAutoTunesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeDomainAutoTunesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeDomainAutoTunesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeDomainAutoTunesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeDomainAutoTunesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeDomainAutoTunesPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribeDomainAutoTunesInput, fn func(*elasticsearchservice.DescribeDomainAutoTunesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeDomainAutoTunesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeDomainAutoTunesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeDomainAutoTunesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeDomainAutoTunesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeDomainAutoTunesWithContext(ctx context.Context, input *elasticsearchservice.DescribeDomainAutoTunesInput, opts ...request.Option) (*elasticsearchservice.DescribeDomainAutoTunesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeDomainAutoTunesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeDomainAutoTunesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeDomainChangeProgress(input *elasticsearchservice.DescribeDomainChangeProgressInput) (*elasticsearchservice.DescribeDomainChangeProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeDomainChangeProgressOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeDomainChangeProgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeDomainChangeProgressWithContext(ctx context.Context, input *elasticsearchservice.DescribeDomainChangeProgressInput, opts ...request.Option) (*elasticsearchservice.DescribeDomainChangeProgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeDomainChangeProgressOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeDomainChangeProgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomain(input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomainConfig(input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainConfigOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomainConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomainConfigWithContext(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput, opts ...request.Option) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainConfigOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomainConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomainWithContext(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput, opts ...request.Option) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomains(input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchDomainsWithContext(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput, opts ...request.Option) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchDomainsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchInstanceTypeLimits(input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchInstanceTypeLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeElasticsearchInstanceTypeLimitsWithContext(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput, opts ...request.Option) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeElasticsearchInstanceTypeLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeInboundCrossClusterSearchConnections(input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeInboundCrossClusterSearchConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeInboundCrossClusterSearchConnectionsPages(input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput, fn func(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeInboundCrossClusterSearchConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeInboundCrossClusterSearchConnectionsPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput, fn func(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeInboundCrossClusterSearchConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeInboundCrossClusterSearchConnectionsWithContext(ctx context.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput, opts ...request.Option) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeInboundCrossClusterSearchConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeOutboundCrossClusterSearchConnections(input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeOutboundCrossClusterSearchConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeOutboundCrossClusterSearchConnectionsPages(input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput, fn func(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeOutboundCrossClusterSearchConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeOutboundCrossClusterSearchConnectionsPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput, fn func(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeOutboundCrossClusterSearchConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeOutboundCrossClusterSearchConnectionsWithContext(ctx context.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput, opts ...request.Option) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeOutboundCrossClusterSearchConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribePackages(input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribePackagesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribePackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribePackagesPages(input *elasticsearchservice.DescribePackagesInput, fn func(*elasticsearchservice.DescribePackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribePackagesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribePackagesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribePackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribePackagesPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribePackagesInput, fn func(*elasticsearchservice.DescribePackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribePackagesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribePackagesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribePackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribePackagesWithContext(ctx context.Context, input *elasticsearchservice.DescribePackagesInput, opts ...request.Option) (*elasticsearchservice.DescribePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribePackagesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribePackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstanceOfferings(input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstanceOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstanceOfferingsPages(input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput, fn func(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstanceOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstanceOfferingsPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput, fn func(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstanceOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstanceOfferingsWithContext(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput, opts ...request.Option) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstanceOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstances(input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstancesPages(input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput, fn func(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeReservedElasticsearchInstancesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstancesPagesWithContext(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput, fn func(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.DescribeReservedElasticsearchInstancesOutput{}
	fnCacher := func(out *elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) DescribeReservedElasticsearchInstancesWithContext(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput, opts ...request.Option) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeReservedElasticsearchInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeVpcEndpoints(input *elasticsearchservice.DescribeVpcEndpointsInput) (*elasticsearchservice.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) DescribeVpcEndpointsWithContext(ctx context.Context, input *elasticsearchservice.DescribeVpcEndpointsInput, opts ...request.Option) (*elasticsearchservice.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.DescribeVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetCompatibleElasticsearchVersions(input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetCompatibleElasticsearchVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetCompatibleElasticsearchVersionsWithContext(ctx context.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput, opts ...request.Option) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetCompatibleElasticsearchVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetPackageVersionHistory(input *elasticsearchservice.GetPackageVersionHistoryInput) (*elasticsearchservice.GetPackageVersionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetPackageVersionHistoryOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetPackageVersionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetPackageVersionHistoryPages(input *elasticsearchservice.GetPackageVersionHistoryInput, fn func(*elasticsearchservice.GetPackageVersionHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.GetPackageVersionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.GetPackageVersionHistoryOutput{}
	fnCacher := func(out *elasticsearchservice.GetPackageVersionHistoryOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.GetPackageVersionHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) GetPackageVersionHistoryPagesWithContext(ctx context.Context, input *elasticsearchservice.GetPackageVersionHistoryInput, fn func(*elasticsearchservice.GetPackageVersionHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.GetPackageVersionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.GetPackageVersionHistoryOutput{}
	fnCacher := func(out *elasticsearchservice.GetPackageVersionHistoryOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.GetPackageVersionHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) GetPackageVersionHistoryWithContext(ctx context.Context, input *elasticsearchservice.GetPackageVersionHistoryInput, opts ...request.Option) (*elasticsearchservice.GetPackageVersionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetPackageVersionHistoryOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetPackageVersionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetUpgradeHistory(input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetUpgradeHistoryOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetUpgradeHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetUpgradeHistoryPages(input *elasticsearchservice.GetUpgradeHistoryInput, fn func(*elasticsearchservice.GetUpgradeHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.GetUpgradeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.GetUpgradeHistoryOutput{}
	fnCacher := func(out *elasticsearchservice.GetUpgradeHistoryOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.GetUpgradeHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) GetUpgradeHistoryPagesWithContext(ctx context.Context, input *elasticsearchservice.GetUpgradeHistoryInput, fn func(*elasticsearchservice.GetUpgradeHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.GetUpgradeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.GetUpgradeHistoryOutput{}
	fnCacher := func(out *elasticsearchservice.GetUpgradeHistoryOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.GetUpgradeHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) GetUpgradeHistoryWithContext(ctx context.Context, input *elasticsearchservice.GetUpgradeHistoryInput, opts ...request.Option) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetUpgradeHistoryOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetUpgradeHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetUpgradeStatus(input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetUpgradeStatusOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetUpgradeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) GetUpgradeStatusWithContext(ctx context.Context, input *elasticsearchservice.GetUpgradeStatusInput, opts ...request.Option) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.GetUpgradeStatusOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.GetUpgradeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListDomainNames(input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListDomainNamesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListDomainNamesWithContext(ctx context.Context, input *elasticsearchservice.ListDomainNamesInput, opts ...request.Option) (*elasticsearchservice.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListDomainNamesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListDomainsForPackage(input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListDomainsForPackageOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListDomainsForPackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListDomainsForPackagePages(input *elasticsearchservice.ListDomainsForPackageInput, fn func(*elasticsearchservice.ListDomainsForPackageOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListDomainsForPackageOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListDomainsForPackageOutput{}
	fnCacher := func(out *elasticsearchservice.ListDomainsForPackageOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListDomainsForPackagePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListDomainsForPackagePagesWithContext(ctx context.Context, input *elasticsearchservice.ListDomainsForPackageInput, fn func(*elasticsearchservice.ListDomainsForPackageOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListDomainsForPackageOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListDomainsForPackageOutput{}
	fnCacher := func(out *elasticsearchservice.ListDomainsForPackageOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListDomainsForPackagePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListDomainsForPackageWithContext(ctx context.Context, input *elasticsearchservice.ListDomainsForPackageInput, opts ...request.Option) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListDomainsForPackageOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListDomainsForPackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListElasticsearchInstanceTypes(input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListElasticsearchInstanceTypesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListElasticsearchInstanceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListElasticsearchInstanceTypesPages(input *elasticsearchservice.ListElasticsearchInstanceTypesInput, fn func(*elasticsearchservice.ListElasticsearchInstanceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListElasticsearchInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListElasticsearchInstanceTypesOutput{}
	fnCacher := func(out *elasticsearchservice.ListElasticsearchInstanceTypesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListElasticsearchInstanceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListElasticsearchInstanceTypesPagesWithContext(ctx context.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput, fn func(*elasticsearchservice.ListElasticsearchInstanceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListElasticsearchInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListElasticsearchInstanceTypesOutput{}
	fnCacher := func(out *elasticsearchservice.ListElasticsearchInstanceTypesOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListElasticsearchInstanceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListElasticsearchInstanceTypesWithContext(ctx context.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput, opts ...request.Option) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListElasticsearchInstanceTypesOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListElasticsearchInstanceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListElasticsearchVersions(input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListElasticsearchVersionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListElasticsearchVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListElasticsearchVersionsPages(input *elasticsearchservice.ListElasticsearchVersionsInput, fn func(*elasticsearchservice.ListElasticsearchVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListElasticsearchVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListElasticsearchVersionsOutput{}
	fnCacher := func(out *elasticsearchservice.ListElasticsearchVersionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListElasticsearchVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListElasticsearchVersionsPagesWithContext(ctx context.Context, input *elasticsearchservice.ListElasticsearchVersionsInput, fn func(*elasticsearchservice.ListElasticsearchVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListElasticsearchVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListElasticsearchVersionsOutput{}
	fnCacher := func(out *elasticsearchservice.ListElasticsearchVersionsOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListElasticsearchVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListElasticsearchVersionsWithContext(ctx context.Context, input *elasticsearchservice.ListElasticsearchVersionsInput, opts ...request.Option) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListElasticsearchVersionsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListElasticsearchVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListPackagesForDomain(input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListPackagesForDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListPackagesForDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListPackagesForDomainPages(input *elasticsearchservice.ListPackagesForDomainInput, fn func(*elasticsearchservice.ListPackagesForDomainOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListPackagesForDomainOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListPackagesForDomainOutput{}
	fnCacher := func(out *elasticsearchservice.ListPackagesForDomainOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListPackagesForDomainPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListPackagesForDomainPagesWithContext(ctx context.Context, input *elasticsearchservice.ListPackagesForDomainInput, fn func(*elasticsearchservice.ListPackagesForDomainOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticsearchservice.ListPackagesForDomainOutput), false)
		return nil
	}
	cachable := true
	output := &elasticsearchservice.ListPackagesForDomainOutput{}
	fnCacher := func(out *elasticsearchservice.ListPackagesForDomainOutput, more bool) bool {
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
	if err := c.ElasticsearchServiceAPI.ListPackagesForDomainPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticsearchService) ListPackagesForDomainWithContext(ctx context.Context, input *elasticsearchservice.ListPackagesForDomainInput, opts ...request.Option) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListPackagesForDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListPackagesForDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListTags(input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListTagsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListTagsWithContext(ctx context.Context, input *elasticsearchservice.ListTagsInput, opts ...request.Option) (*elasticsearchservice.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListTagsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpointAccess(input *elasticsearchservice.ListVpcEndpointAccessInput) (*elasticsearchservice.ListVpcEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointAccessOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpointAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpointAccessWithContext(ctx context.Context, input *elasticsearchservice.ListVpcEndpointAccessInput, opts ...request.Option) (*elasticsearchservice.ListVpcEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointAccessOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpointAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpoints(input *elasticsearchservice.ListVpcEndpointsInput) (*elasticsearchservice.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpointsForDomain(input *elasticsearchservice.ListVpcEndpointsForDomainInput) (*elasticsearchservice.ListVpcEndpointsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointsForDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpointsForDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpointsForDomainWithContext(ctx context.Context, input *elasticsearchservice.ListVpcEndpointsForDomainInput, opts ...request.Option) (*elasticsearchservice.ListVpcEndpointsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointsForDomainOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpointsForDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticsearchService) ListVpcEndpointsWithContext(ctx context.Context, input *elasticsearchservice.ListVpcEndpointsInput, opts ...request.Option) (*elasticsearchservice.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticsearchservice.ListVpcEndpointsOutput), nil
	}
	out, err := c.ElasticsearchServiceAPI.ListVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

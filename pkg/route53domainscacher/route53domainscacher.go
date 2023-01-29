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
package route53domainscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53Domains struct {
	route53domainsiface.Route53DomainsAPI
	cache *cache.Cache
}

func New(route53domainsapi route53domainsiface.Route53DomainsAPI) *Route53Domains {
	return &Route53Domains{
		Route53DomainsAPI: route53domainsapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53Domains) GetContactReachabilityStatus(input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetContactReachabilityStatusOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetContactReachabilityStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetContactReachabilityStatusWithContext(ctx context.Context, input *route53domains.GetContactReachabilityStatusInput, opts ...request.Option) (*route53domains.GetContactReachabilityStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetContactReachabilityStatusOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetContactReachabilityStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetDomainDetail(input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetDomainDetailOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetDomainDetail(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetDomainDetailWithContext(ctx context.Context, input *route53domains.GetDomainDetailInput, opts ...request.Option) (*route53domains.GetDomainDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetDomainDetailOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetDomainDetailWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetDomainSuggestions(input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetDomainSuggestionsOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetDomainSuggestions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetDomainSuggestionsWithContext(ctx context.Context, input *route53domains.GetDomainSuggestionsInput, opts ...request.Option) (*route53domains.GetDomainSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetDomainSuggestionsOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetDomainSuggestionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetOperationDetail(input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetOperationDetailOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetOperationDetail(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) GetOperationDetailWithContext(ctx context.Context, input *route53domains.GetOperationDetailInput, opts ...request.Option) (*route53domains.GetOperationDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.GetOperationDetailOutput), nil
	}
	out, err := c.Route53DomainsAPI.GetOperationDetailWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListDomains(input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListDomainsOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListDomainsPages(input *route53domains.ListDomainsInput, fn func(*route53domains.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListDomainsOutput{}
	fnCacher := func(out *route53domains.ListDomainsOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListDomainsPagesWithContext(ctx context.Context, input *route53domains.ListDomainsInput, fn func(*route53domains.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListDomainsOutput{}
	fnCacher := func(out *route53domains.ListDomainsOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListDomainsWithContext(ctx context.Context, input *route53domains.ListDomainsInput, opts ...request.Option) (*route53domains.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListDomainsOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListOperations(input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListOperationsOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListOperationsPages(input *route53domains.ListOperationsInput, fn func(*route53domains.ListOperationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListOperationsOutput{}
	fnCacher := func(out *route53domains.ListOperationsOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListOperationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListOperationsPagesWithContext(ctx context.Context, input *route53domains.ListOperationsInput, fn func(*route53domains.ListOperationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListOperationsOutput{}
	fnCacher := func(out *route53domains.ListOperationsOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListOperationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListOperationsWithContext(ctx context.Context, input *route53domains.ListOperationsInput, opts ...request.Option) (*route53domains.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListOperationsOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListPrices(input *route53domains.ListPricesInput) (*route53domains.ListPricesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListPricesOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListPrices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListPricesPages(input *route53domains.ListPricesInput, fn func(*route53domains.ListPricesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListPricesOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListPricesOutput{}
	fnCacher := func(out *route53domains.ListPricesOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListPricesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListPricesPagesWithContext(ctx context.Context, input *route53domains.ListPricesInput, fn func(*route53domains.ListPricesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ListPricesOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ListPricesOutput{}
	fnCacher := func(out *route53domains.ListPricesOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ListPricesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ListPricesWithContext(ctx context.Context, input *route53domains.ListPricesInput, opts ...request.Option) (*route53domains.ListPricesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListPricesOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListPricesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListTagsForDomain(input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListTagsForDomainOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListTagsForDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ListTagsForDomainWithContext(ctx context.Context, input *route53domains.ListTagsForDomainInput, opts ...request.Option) (*route53domains.ListTagsForDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ListTagsForDomainOutput), nil
	}
	out, err := c.Route53DomainsAPI.ListTagsForDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ViewBilling(input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ViewBillingOutput), nil
	}
	out, err := c.Route53DomainsAPI.ViewBilling(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Domains) ViewBillingPages(input *route53domains.ViewBillingInput, fn func(*route53domains.ViewBillingOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ViewBillingOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ViewBillingOutput{}
	fnCacher := func(out *route53domains.ViewBillingOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ViewBillingPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ViewBillingPagesWithContext(ctx context.Context, input *route53domains.ViewBillingInput, fn func(*route53domains.ViewBillingOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53domains.ViewBillingOutput), false)
		return nil
	}
	cachable := true
	output := &route53domains.ViewBillingOutput{}
	fnCacher := func(out *route53domains.ViewBillingOutput, more bool) bool {
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
	if err := c.Route53DomainsAPI.ViewBillingPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Domains) ViewBillingWithContext(ctx context.Context, input *route53domains.ViewBillingInput, opts ...request.Option) (*route53domains.ViewBillingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53domains.ViewBillingOutput), nil
	}
	out, err := c.Route53DomainsAPI.ViewBillingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

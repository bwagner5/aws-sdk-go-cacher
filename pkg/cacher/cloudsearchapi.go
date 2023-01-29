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
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudSearch struct {
	cloudsearchiface.CloudSearchAPI
	cache *cache.Cache
}

func NewCloudSearch(cloudsearchapi cloudsearchiface.CloudSearchAPI) *CloudSearch {
	return &CloudSearch{
		CloudSearchAPI: cloudsearchapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudSearch) DescribeAnalysisSchemes(input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeAnalysisSchemesOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeAnalysisSchemes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeAnalysisSchemesWithContext(ctx context.Context, input *cloudsearch.DescribeAnalysisSchemesInput, opts ...request.Option) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeAnalysisSchemesOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeAnalysisSchemesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeAvailabilityOptions(input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeAvailabilityOptionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeAvailabilityOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeAvailabilityOptionsWithContext(ctx context.Context, input *cloudsearch.DescribeAvailabilityOptionsInput, opts ...request.Option) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeAvailabilityOptionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeAvailabilityOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeDomainEndpointOptions(input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeDomainEndpointOptionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeDomainEndpointOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeDomainEndpointOptionsWithContext(ctx context.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput, opts ...request.Option) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeDomainEndpointOptionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeDomainEndpointOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeDomains(input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeDomainsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeDomainsWithContext(ctx context.Context, input *cloudsearch.DescribeDomainsInput, opts ...request.Option) (*cloudsearch.DescribeDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeDomainsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeExpressions(input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeExpressionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeExpressions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeExpressionsWithContext(ctx context.Context, input *cloudsearch.DescribeExpressionsInput, opts ...request.Option) (*cloudsearch.DescribeExpressionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeExpressionsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeExpressionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeIndexFields(input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeIndexFieldsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeIndexFields(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeIndexFieldsWithContext(ctx context.Context, input *cloudsearch.DescribeIndexFieldsInput, opts ...request.Option) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeIndexFieldsOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeIndexFieldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeScalingParameters(input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeScalingParametersOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeScalingParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeScalingParametersWithContext(ctx context.Context, input *cloudsearch.DescribeScalingParametersInput, opts ...request.Option) (*cloudsearch.DescribeScalingParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeScalingParametersOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeScalingParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeServiceAccessPolicies(input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeServiceAccessPoliciesOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeServiceAccessPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeServiceAccessPoliciesWithContext(ctx context.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput, opts ...request.Option) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeServiceAccessPoliciesOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeServiceAccessPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeSuggesters(input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeSuggestersOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeSuggesters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) DescribeSuggestersWithContext(ctx context.Context, input *cloudsearch.DescribeSuggestersInput, opts ...request.Option) (*cloudsearch.DescribeSuggestersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.DescribeSuggestersOutput), nil
	}
	out, err := c.CloudSearchAPI.DescribeSuggestersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) ListDomainNames(input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.ListDomainNamesOutput), nil
	}
	out, err := c.CloudSearchAPI.ListDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearch) ListDomainNamesWithContext(ctx context.Context, input *cloudsearch.ListDomainNamesInput, opts ...request.Option) (*cloudsearch.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearch.ListDomainNamesOutput), nil
	}
	out, err := c.CloudSearchAPI.ListDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

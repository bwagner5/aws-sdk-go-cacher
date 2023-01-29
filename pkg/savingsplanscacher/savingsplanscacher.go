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
package savingsplanscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/savingsplans/savingsplansiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SavingsPlans struct {
	savingsplansiface.SavingsPlansAPI
	cache *cache.Cache
}

func New(savingsplansapi savingsplansiface.SavingsPlansAPI) *SavingsPlans {
	return &SavingsPlans{
		SavingsPlansAPI: savingsplansapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SavingsPlans) DescribeSavingsPlanRates(input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlanRatesOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlanRates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlanRatesWithContext(ctx context.Context, input *savingsplans.DescribeSavingsPlanRatesInput, opts ...request.Option) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlanRatesOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlanRatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlans(input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlansOfferingRates(input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOfferingRatesOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlansOfferingRates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlansOfferingRatesWithContext(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput, opts ...request.Option) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOfferingRatesOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlansOfferingRatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlansOfferings(input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOfferingsOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlansOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlansOfferingsWithContext(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput, opts ...request.Option) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOfferingsOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlansOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) DescribeSavingsPlansWithContext(ctx context.Context, input *savingsplans.DescribeSavingsPlansInput, opts ...request.Option) (*savingsplans.DescribeSavingsPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.DescribeSavingsPlansOutput), nil
	}
	out, err := c.SavingsPlansAPI.DescribeSavingsPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) ListTagsForResource(input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.ListTagsForResourceOutput), nil
	}
	out, err := c.SavingsPlansAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SavingsPlans) ListTagsForResourceWithContext(ctx context.Context, input *savingsplans.ListTagsForResourceInput, opts ...request.Option) (*savingsplans.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*savingsplans.ListTagsForResourceOutput), nil
	}
	out, err := c.SavingsPlansAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

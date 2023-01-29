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
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MarketplaceMetering struct {
	marketplacemeteringiface.MarketplaceMeteringAPI
	cache *cache.Cache
}

func NewMarketplaceMetering(marketplacemeteringapi marketplacemeteringiface.MarketplaceMeteringAPI) *MarketplaceMetering {
	return &MarketplaceMetering{
		MarketplaceMeteringAPI: marketplacemeteringapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MarketplaceMetering) BatchMeterUsage(input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacemetering.BatchMeterUsageOutput), nil
	}
	out, err := c.MarketplaceMeteringAPI.BatchMeterUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceMetering) BatchMeterUsageWithContext(ctx context.Context, input *marketplacemetering.BatchMeterUsageInput, opts ...request.Option) (*marketplacemetering.BatchMeterUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacemetering.BatchMeterUsageOutput), nil
	}
	out, err := c.MarketplaceMeteringAPI.BatchMeterUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

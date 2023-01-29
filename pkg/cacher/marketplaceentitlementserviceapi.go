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
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MarketplaceEntitlementService struct {
	marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
	cache *cache.Cache
}

func NewMarketplaceEntitlementService(marketplaceentitlementserviceapi marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI) *MarketplaceEntitlementService {
	return &MarketplaceEntitlementService{
		MarketplaceEntitlementServiceAPI: marketplaceentitlementserviceapi,
		cache:                            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MarketplaceEntitlementService) GetEntitlements(input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplaceentitlementservice.GetEntitlementsOutput), nil
	}
	out, err := c.MarketplaceEntitlementServiceAPI.GetEntitlements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceEntitlementService) GetEntitlementsWithContext(ctx context.Context, input *marketplaceentitlementservice.GetEntitlementsInput, opts ...request.Option) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplaceentitlementservice.GetEntitlementsOutput), nil
	}
	out, err := c.MarketplaceEntitlementServiceAPI.GetEntitlementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

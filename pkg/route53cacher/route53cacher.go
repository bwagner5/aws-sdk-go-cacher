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
package route53cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53 struct {
	route53iface.Route53API
	cache *cache.Cache
}

func New(route53api route53iface.Route53API) *Route53 {
	return &Route53{
		Route53API: route53api,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53) GetAccountLimit(input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetAccountLimitOutput), nil
	}
	out, err := c.Route53API.GetAccountLimit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetAccountLimitWithContext(ctx context.Context, input *route53.GetAccountLimitInput, opts ...request.Option) (*route53.GetAccountLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetAccountLimitOutput), nil
	}
	out, err := c.Route53API.GetAccountLimitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetChange(input *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetChangeOutput), nil
	}
	out, err := c.Route53API.GetChange(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetChangeWithContext(ctx context.Context, input *route53.GetChangeInput, opts ...request.Option) (*route53.GetChangeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetChangeOutput), nil
	}
	out, err := c.Route53API.GetChangeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetCheckerIpRanges(input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetCheckerIpRangesOutput), nil
	}
	out, err := c.Route53API.GetCheckerIpRanges(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetCheckerIpRangesWithContext(ctx context.Context, input *route53.GetCheckerIpRangesInput, opts ...request.Option) (*route53.GetCheckerIpRangesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetCheckerIpRangesOutput), nil
	}
	out, err := c.Route53API.GetCheckerIpRangesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetDNSSEC(input *route53.GetDNSSECInput) (*route53.GetDNSSECOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetDNSSECOutput), nil
	}
	out, err := c.Route53API.GetDNSSEC(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetDNSSECWithContext(ctx context.Context, input *route53.GetDNSSECInput, opts ...request.Option) (*route53.GetDNSSECOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetDNSSECOutput), nil
	}
	out, err := c.Route53API.GetDNSSECWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetGeoLocation(input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetGeoLocationOutput), nil
	}
	out, err := c.Route53API.GetGeoLocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetGeoLocationWithContext(ctx context.Context, input *route53.GetGeoLocationInput, opts ...request.Option) (*route53.GetGeoLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetGeoLocationOutput), nil
	}
	out, err := c.Route53API.GetGeoLocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheck(input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckOutput), nil
	}
	out, err := c.Route53API.GetHealthCheck(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckCount(input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckCountOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckCountWithContext(ctx context.Context, input *route53.GetHealthCheckCountInput, opts ...request.Option) (*route53.GetHealthCheckCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckCountOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckLastFailureReason(input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckLastFailureReasonOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckLastFailureReason(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckLastFailureReasonWithContext(ctx context.Context, input *route53.GetHealthCheckLastFailureReasonInput, opts ...request.Option) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckLastFailureReasonOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckLastFailureReasonWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckStatus(input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckStatusOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckStatusWithContext(ctx context.Context, input *route53.GetHealthCheckStatusInput, opts ...request.Option) (*route53.GetHealthCheckStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckStatusOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHealthCheckWithContext(ctx context.Context, input *route53.GetHealthCheckInput, opts ...request.Option) (*route53.GetHealthCheckOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHealthCheckOutput), nil
	}
	out, err := c.Route53API.GetHealthCheckWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZone(input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneOutput), nil
	}
	out, err := c.Route53API.GetHostedZone(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZoneCount(input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneCountOutput), nil
	}
	out, err := c.Route53API.GetHostedZoneCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZoneCountWithContext(ctx context.Context, input *route53.GetHostedZoneCountInput, opts ...request.Option) (*route53.GetHostedZoneCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneCountOutput), nil
	}
	out, err := c.Route53API.GetHostedZoneCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZoneLimit(input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneLimitOutput), nil
	}
	out, err := c.Route53API.GetHostedZoneLimit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZoneLimitWithContext(ctx context.Context, input *route53.GetHostedZoneLimitInput, opts ...request.Option) (*route53.GetHostedZoneLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneLimitOutput), nil
	}
	out, err := c.Route53API.GetHostedZoneLimitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetHostedZoneWithContext(ctx context.Context, input *route53.GetHostedZoneInput, opts ...request.Option) (*route53.GetHostedZoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetHostedZoneOutput), nil
	}
	out, err := c.Route53API.GetHostedZoneWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetQueryLoggingConfig(input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetQueryLoggingConfigOutput), nil
	}
	out, err := c.Route53API.GetQueryLoggingConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetQueryLoggingConfigWithContext(ctx context.Context, input *route53.GetQueryLoggingConfigInput, opts ...request.Option) (*route53.GetQueryLoggingConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetQueryLoggingConfigOutput), nil
	}
	out, err := c.Route53API.GetQueryLoggingConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetReusableDelegationSet(input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetReusableDelegationSetOutput), nil
	}
	out, err := c.Route53API.GetReusableDelegationSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetReusableDelegationSetLimit(input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetReusableDelegationSetLimitOutput), nil
	}
	out, err := c.Route53API.GetReusableDelegationSetLimit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetReusableDelegationSetLimitWithContext(ctx context.Context, input *route53.GetReusableDelegationSetLimitInput, opts ...request.Option) (*route53.GetReusableDelegationSetLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetReusableDelegationSetLimitOutput), nil
	}
	out, err := c.Route53API.GetReusableDelegationSetLimitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetReusableDelegationSetWithContext(ctx context.Context, input *route53.GetReusableDelegationSetInput, opts ...request.Option) (*route53.GetReusableDelegationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetReusableDelegationSetOutput), nil
	}
	out, err := c.Route53API.GetReusableDelegationSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicy(input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicyInstance(input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyInstanceOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicyInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicyInstanceCount(input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyInstanceCountOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicyInstanceCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicyInstanceCountWithContext(ctx context.Context, input *route53.GetTrafficPolicyInstanceCountInput, opts ...request.Option) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyInstanceCountOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicyInstanceCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicyInstanceWithContext(ctx context.Context, input *route53.GetTrafficPolicyInstanceInput, opts ...request.Option) (*route53.GetTrafficPolicyInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyInstanceOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicyInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) GetTrafficPolicyWithContext(ctx context.Context, input *route53.GetTrafficPolicyInput, opts ...request.Option) (*route53.GetTrafficPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.GetTrafficPolicyOutput), nil
	}
	out, err := c.Route53API.GetTrafficPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrBlocks(input *route53.ListCidrBlocksInput) (*route53.ListCidrBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrBlocksOutput), nil
	}
	out, err := c.Route53API.ListCidrBlocks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrBlocksPages(input *route53.ListCidrBlocksInput, fn func(*route53.ListCidrBlocksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrBlocksOutput{}
	fnCacher := func(out *route53.ListCidrBlocksOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrBlocksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrBlocksPagesWithContext(ctx context.Context, input *route53.ListCidrBlocksInput, fn func(*route53.ListCidrBlocksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrBlocksOutput{}
	fnCacher := func(out *route53.ListCidrBlocksOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrBlocksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrBlocksWithContext(ctx context.Context, input *route53.ListCidrBlocksInput, opts ...request.Option) (*route53.ListCidrBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrBlocksOutput), nil
	}
	out, err := c.Route53API.ListCidrBlocksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrCollections(input *route53.ListCidrCollectionsInput) (*route53.ListCidrCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrCollectionsOutput), nil
	}
	out, err := c.Route53API.ListCidrCollections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrCollectionsPages(input *route53.ListCidrCollectionsInput, fn func(*route53.ListCidrCollectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrCollectionsOutput{}
	fnCacher := func(out *route53.ListCidrCollectionsOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrCollectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrCollectionsPagesWithContext(ctx context.Context, input *route53.ListCidrCollectionsInput, fn func(*route53.ListCidrCollectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrCollectionsOutput{}
	fnCacher := func(out *route53.ListCidrCollectionsOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrCollectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrCollectionsWithContext(ctx context.Context, input *route53.ListCidrCollectionsInput, opts ...request.Option) (*route53.ListCidrCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrCollectionsOutput), nil
	}
	out, err := c.Route53API.ListCidrCollectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrLocations(input *route53.ListCidrLocationsInput) (*route53.ListCidrLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrLocationsOutput), nil
	}
	out, err := c.Route53API.ListCidrLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListCidrLocationsPages(input *route53.ListCidrLocationsInput, fn func(*route53.ListCidrLocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrLocationsOutput{}
	fnCacher := func(out *route53.ListCidrLocationsOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrLocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrLocationsPagesWithContext(ctx context.Context, input *route53.ListCidrLocationsInput, fn func(*route53.ListCidrLocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListCidrLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListCidrLocationsOutput{}
	fnCacher := func(out *route53.ListCidrLocationsOutput, more bool) bool {
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
	if err := c.Route53API.ListCidrLocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListCidrLocationsWithContext(ctx context.Context, input *route53.ListCidrLocationsInput, opts ...request.Option) (*route53.ListCidrLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListCidrLocationsOutput), nil
	}
	out, err := c.Route53API.ListCidrLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListGeoLocations(input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListGeoLocationsOutput), nil
	}
	out, err := c.Route53API.ListGeoLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListGeoLocationsWithContext(ctx context.Context, input *route53.ListGeoLocationsInput, opts ...request.Option) (*route53.ListGeoLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListGeoLocationsOutput), nil
	}
	out, err := c.Route53API.ListGeoLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHealthChecks(input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHealthChecksOutput), nil
	}
	out, err := c.Route53API.ListHealthChecks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHealthChecksPages(input *route53.ListHealthChecksInput, fn func(*route53.ListHealthChecksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListHealthChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListHealthChecksOutput{}
	fnCacher := func(out *route53.ListHealthChecksOutput, more bool) bool {
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
	if err := c.Route53API.ListHealthChecksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListHealthChecksPagesWithContext(ctx context.Context, input *route53.ListHealthChecksInput, fn func(*route53.ListHealthChecksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListHealthChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListHealthChecksOutput{}
	fnCacher := func(out *route53.ListHealthChecksOutput, more bool) bool {
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
	if err := c.Route53API.ListHealthChecksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListHealthChecksWithContext(ctx context.Context, input *route53.ListHealthChecksInput, opts ...request.Option) (*route53.ListHealthChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHealthChecksOutput), nil
	}
	out, err := c.Route53API.ListHealthChecksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZones(input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesOutput), nil
	}
	out, err := c.Route53API.ListHostedZones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesByNameOutput), nil
	}
	out, err := c.Route53API.ListHostedZonesByName(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZonesByNameWithContext(ctx context.Context, input *route53.ListHostedZonesByNameInput, opts ...request.Option) (*route53.ListHostedZonesByNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesByNameOutput), nil
	}
	out, err := c.Route53API.ListHostedZonesByNameWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZonesByVPC(input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesByVPCOutput), nil
	}
	out, err := c.Route53API.ListHostedZonesByVPC(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZonesByVPCWithContext(ctx context.Context, input *route53.ListHostedZonesByVPCInput, opts ...request.Option) (*route53.ListHostedZonesByVPCOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesByVPCOutput), nil
	}
	out, err := c.Route53API.ListHostedZonesByVPCWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListHostedZonesPages(input *route53.ListHostedZonesInput, fn func(*route53.ListHostedZonesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListHostedZonesOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListHostedZonesOutput{}
	fnCacher := func(out *route53.ListHostedZonesOutput, more bool) bool {
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
	if err := c.Route53API.ListHostedZonesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListHostedZonesPagesWithContext(ctx context.Context, input *route53.ListHostedZonesInput, fn func(*route53.ListHostedZonesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListHostedZonesOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListHostedZonesOutput{}
	fnCacher := func(out *route53.ListHostedZonesOutput, more bool) bool {
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
	if err := c.Route53API.ListHostedZonesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListHostedZonesWithContext(ctx context.Context, input *route53.ListHostedZonesInput, opts ...request.Option) (*route53.ListHostedZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListHostedZonesOutput), nil
	}
	out, err := c.Route53API.ListHostedZonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListQueryLoggingConfigs(input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListQueryLoggingConfigsOutput), nil
	}
	out, err := c.Route53API.ListQueryLoggingConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListQueryLoggingConfigsPages(input *route53.ListQueryLoggingConfigsInput, fn func(*route53.ListQueryLoggingConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListQueryLoggingConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListQueryLoggingConfigsOutput{}
	fnCacher := func(out *route53.ListQueryLoggingConfigsOutput, more bool) bool {
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
	if err := c.Route53API.ListQueryLoggingConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListQueryLoggingConfigsPagesWithContext(ctx context.Context, input *route53.ListQueryLoggingConfigsInput, fn func(*route53.ListQueryLoggingConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListQueryLoggingConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListQueryLoggingConfigsOutput{}
	fnCacher := func(out *route53.ListQueryLoggingConfigsOutput, more bool) bool {
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
	if err := c.Route53API.ListQueryLoggingConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListQueryLoggingConfigsWithContext(ctx context.Context, input *route53.ListQueryLoggingConfigsInput, opts ...request.Option) (*route53.ListQueryLoggingConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListQueryLoggingConfigsOutput), nil
	}
	out, err := c.Route53API.ListQueryLoggingConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListResourceRecordSets(input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListResourceRecordSetsOutput), nil
	}
	out, err := c.Route53API.ListResourceRecordSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(*route53.ListResourceRecordSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListResourceRecordSetsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListResourceRecordSetsOutput{}
	fnCacher := func(out *route53.ListResourceRecordSetsOutput, more bool) bool {
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
	if err := c.Route53API.ListResourceRecordSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListResourceRecordSetsPagesWithContext(ctx context.Context, input *route53.ListResourceRecordSetsInput, fn func(*route53.ListResourceRecordSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53.ListResourceRecordSetsOutput), false)
		return nil
	}
	cachable := true
	output := &route53.ListResourceRecordSetsOutput{}
	fnCacher := func(out *route53.ListResourceRecordSetsOutput, more bool) bool {
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
	if err := c.Route53API.ListResourceRecordSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53) ListResourceRecordSetsWithContext(ctx context.Context, input *route53.ListResourceRecordSetsInput, opts ...request.Option) (*route53.ListResourceRecordSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListResourceRecordSetsOutput), nil
	}
	out, err := c.Route53API.ListResourceRecordSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListReusableDelegationSets(input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListReusableDelegationSetsOutput), nil
	}
	out, err := c.Route53API.ListReusableDelegationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListReusableDelegationSetsWithContext(ctx context.Context, input *route53.ListReusableDelegationSetsInput, opts ...request.Option) (*route53.ListReusableDelegationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListReusableDelegationSetsOutput), nil
	}
	out, err := c.Route53API.ListReusableDelegationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTagsForResource(input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTagsForResourceWithContext(ctx context.Context, input *route53.ListTagsForResourceInput, opts ...request.Option) (*route53.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTagsForResources(input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTagsForResourcesOutput), nil
	}
	out, err := c.Route53API.ListTagsForResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTagsForResourcesWithContext(ctx context.Context, input *route53.ListTagsForResourcesInput, opts ...request.Option) (*route53.ListTagsForResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTagsForResourcesOutput), nil
	}
	out, err := c.Route53API.ListTagsForResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicies(input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPoliciesOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPoliciesWithContext(ctx context.Context, input *route53.ListTrafficPoliciesInput, opts ...request.Option) (*route53.ListTrafficPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPoliciesOutput), nil
	}
	out, err := c.Route53API.ListTrafficPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstances(input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstancesByHostedZone(input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesByHostedZoneOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstancesByHostedZone(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstancesByHostedZoneWithContext(ctx context.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput, opts ...request.Option) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesByHostedZoneOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstancesByHostedZoneWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstancesByPolicy(input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesByPolicyOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstancesByPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstancesByPolicyWithContext(ctx context.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput, opts ...request.Option) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesByPolicyOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstancesByPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyInstancesWithContext(ctx context.Context, input *route53.ListTrafficPolicyInstancesInput, opts ...request.Option) (*route53.ListTrafficPolicyInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyInstancesOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyVersions(input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyVersionsOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListTrafficPolicyVersionsWithContext(ctx context.Context, input *route53.ListTrafficPolicyVersionsInput, opts ...request.Option) (*route53.ListTrafficPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListTrafficPolicyVersionsOutput), nil
	}
	out, err := c.Route53API.ListTrafficPolicyVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListVPCAssociationAuthorizations(input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListVPCAssociationAuthorizationsOutput), nil
	}
	out, err := c.Route53API.ListVPCAssociationAuthorizations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53) ListVPCAssociationAuthorizationsWithContext(ctx context.Context, input *route53.ListVPCAssociationAuthorizationsInput, opts ...request.Option) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53.ListVPCAssociationAuthorizationsOutput), nil
	}
	out, err := c.Route53API.ListVPCAssociationAuthorizationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go/service/route53recoverycontrolconfig/route53recoverycontrolconfigiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53RecoveryControlConfig struct {
	route53recoverycontrolconfigiface.Route53RecoveryControlConfigAPI
	cache *cache.Cache
}

func NewRoute53RecoveryControlConfig(route53recoverycontrolconfigapi route53recoverycontrolconfigiface.Route53RecoveryControlConfigAPI) *Route53RecoveryControlConfig {
	return &Route53RecoveryControlConfig{
		Route53RecoveryControlConfigAPI: route53recoverycontrolconfigapi,
		cache:                           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53RecoveryControlConfig) DescribeCluster(input *route53recoverycontrolconfig.DescribeClusterInput) (*route53recoverycontrolconfig.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeClusterOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeClusterWithContext(ctx context.Context, input *route53recoverycontrolconfig.DescribeClusterInput, opts ...request.Option) (*route53recoverycontrolconfig.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeClusterOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeControlPanel(input *route53recoverycontrolconfig.DescribeControlPanelInput) (*route53recoverycontrolconfig.DescribeControlPanelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeControlPanelOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeControlPanel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeControlPanelWithContext(ctx context.Context, input *route53recoverycontrolconfig.DescribeControlPanelInput, opts ...request.Option) (*route53recoverycontrolconfig.DescribeControlPanelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeControlPanelOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeControlPanelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeRoutingControl(input *route53recoverycontrolconfig.DescribeRoutingControlInput) (*route53recoverycontrolconfig.DescribeRoutingControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeRoutingControlOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeRoutingControl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeRoutingControlWithContext(ctx context.Context, input *route53recoverycontrolconfig.DescribeRoutingControlInput, opts ...request.Option) (*route53recoverycontrolconfig.DescribeRoutingControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeRoutingControlOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeRoutingControlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeSafetyRule(input *route53recoverycontrolconfig.DescribeSafetyRuleInput) (*route53recoverycontrolconfig.DescribeSafetyRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeSafetyRuleOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeSafetyRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) DescribeSafetyRuleWithContext(ctx context.Context, input *route53recoverycontrolconfig.DescribeSafetyRuleInput, opts ...request.Option) (*route53recoverycontrolconfig.DescribeSafetyRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.DescribeSafetyRuleOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.DescribeSafetyRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListAssociatedRoute53HealthChecks(input *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput) (*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListAssociatedRoute53HealthChecks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListAssociatedRoute53HealthChecksPages(input *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput, fn func(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListAssociatedRoute53HealthChecksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListAssociatedRoute53HealthChecksPagesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput, fn func(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListAssociatedRoute53HealthChecksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListAssociatedRoute53HealthChecksWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput, opts ...request.Option) (*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListAssociatedRoute53HealthChecksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListClusters(input *route53recoverycontrolconfig.ListClustersInput) (*route53recoverycontrolconfig.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListClustersOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListClustersPages(input *route53recoverycontrolconfig.ListClustersInput, fn func(*route53recoverycontrolconfig.ListClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListClustersOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListClustersOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListClustersPagesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListClustersInput, fn func(*route53recoverycontrolconfig.ListClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListClustersOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListClustersOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListClustersWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListClustersInput, opts ...request.Option) (*route53recoverycontrolconfig.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListClustersOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListControlPanels(input *route53recoverycontrolconfig.ListControlPanelsInput) (*route53recoverycontrolconfig.ListControlPanelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListControlPanelsOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListControlPanels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListControlPanelsPages(input *route53recoverycontrolconfig.ListControlPanelsInput, fn func(*route53recoverycontrolconfig.ListControlPanelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListControlPanelsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListControlPanelsOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListControlPanelsOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListControlPanelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListControlPanelsPagesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListControlPanelsInput, fn func(*route53recoverycontrolconfig.ListControlPanelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListControlPanelsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListControlPanelsOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListControlPanelsOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListControlPanelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListControlPanelsWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListControlPanelsInput, opts ...request.Option) (*route53recoverycontrolconfig.ListControlPanelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListControlPanelsOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListControlPanelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListRoutingControls(input *route53recoverycontrolconfig.ListRoutingControlsInput) (*route53recoverycontrolconfig.ListRoutingControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListRoutingControlsOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListRoutingControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListRoutingControlsPages(input *route53recoverycontrolconfig.ListRoutingControlsInput, fn func(*route53recoverycontrolconfig.ListRoutingControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListRoutingControlsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListRoutingControlsOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListRoutingControlsOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListRoutingControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListRoutingControlsPagesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListRoutingControlsInput, fn func(*route53recoverycontrolconfig.ListRoutingControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListRoutingControlsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListRoutingControlsOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListRoutingControlsOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListRoutingControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListRoutingControlsWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListRoutingControlsInput, opts ...request.Option) (*route53recoverycontrolconfig.ListRoutingControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListRoutingControlsOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListRoutingControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListSafetyRules(input *route53recoverycontrolconfig.ListSafetyRulesInput) (*route53recoverycontrolconfig.ListSafetyRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListSafetyRulesOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListSafetyRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListSafetyRulesPages(input *route53recoverycontrolconfig.ListSafetyRulesInput, fn func(*route53recoverycontrolconfig.ListSafetyRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListSafetyRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListSafetyRulesOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListSafetyRulesOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListSafetyRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListSafetyRulesPagesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListSafetyRulesInput, fn func(*route53recoverycontrolconfig.ListSafetyRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycontrolconfig.ListSafetyRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycontrolconfig.ListSafetyRulesOutput{}
	fnCacher := func(out *route53recoverycontrolconfig.ListSafetyRulesOutput, more bool) bool {
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
	if err := c.Route53RecoveryControlConfigAPI.ListSafetyRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryControlConfig) ListSafetyRulesWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListSafetyRulesInput, opts ...request.Option) (*route53recoverycontrolconfig.ListSafetyRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListSafetyRulesOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListSafetyRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListTagsForResource(input *route53recoverycontrolconfig.ListTagsForResourceInput) (*route53recoverycontrolconfig.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryControlConfig) ListTagsForResourceWithContext(ctx context.Context, input *route53recoverycontrolconfig.ListTagsForResourceInput, opts ...request.Option) (*route53recoverycontrolconfig.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycontrolconfig.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53RecoveryControlConfigAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package route53recoveryreadinesscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go/service/route53recoveryreadiness/route53recoveryreadinessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53RecoveryReadiness struct {
	route53recoveryreadinessiface.Route53RecoveryReadinessAPI
	cache *cache.Cache
}

func New(route53recoveryreadinessapi route53recoveryreadinessiface.Route53RecoveryReadinessAPI) *Route53RecoveryReadiness {
	return &Route53RecoveryReadiness{
		Route53RecoveryReadinessAPI: route53recoveryreadinessapi,
		cache:                       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53RecoveryReadiness) GetArchitectureRecommendations(input *route53recoveryreadiness.GetArchitectureRecommendationsInput) (*route53recoveryreadiness.GetArchitectureRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetArchitectureRecommendationsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetArchitectureRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetArchitectureRecommendationsWithContext(ctx context.Context, input *route53recoveryreadiness.GetArchitectureRecommendationsInput, opts ...request.Option) (*route53recoveryreadiness.GetArchitectureRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetArchitectureRecommendationsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetArchitectureRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetCell(input *route53recoveryreadiness.GetCellInput) (*route53recoveryreadiness.GetCellOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetCellOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetCell(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetCellReadinessSummary(input *route53recoveryreadiness.GetCellReadinessSummaryInput) (*route53recoveryreadiness.GetCellReadinessSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetCellReadinessSummaryOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetCellReadinessSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetCellReadinessSummaryPages(input *route53recoveryreadiness.GetCellReadinessSummaryInput, fn func(*route53recoveryreadiness.GetCellReadinessSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetCellReadinessSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetCellReadinessSummaryOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetCellReadinessSummaryOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetCellReadinessSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetCellReadinessSummaryPagesWithContext(ctx context.Context, input *route53recoveryreadiness.GetCellReadinessSummaryInput, fn func(*route53recoveryreadiness.GetCellReadinessSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetCellReadinessSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetCellReadinessSummaryOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetCellReadinessSummaryOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetCellReadinessSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetCellReadinessSummaryWithContext(ctx context.Context, input *route53recoveryreadiness.GetCellReadinessSummaryInput, opts ...request.Option) (*route53recoveryreadiness.GetCellReadinessSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetCellReadinessSummaryOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetCellReadinessSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetCellWithContext(ctx context.Context, input *route53recoveryreadiness.GetCellInput, opts ...request.Option) (*route53recoveryreadiness.GetCellOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetCellOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetCellWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheck(input *route53recoveryreadiness.GetReadinessCheckInput) (*route53recoveryreadiness.GetReadinessCheckOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheck(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheckResourceStatus(input *route53recoveryreadiness.GetReadinessCheckResourceStatusInput) (*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheckResourceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheckResourceStatusPages(input *route53recoveryreadiness.GetReadinessCheckResourceStatusInput, fn func(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetReadinessCheckResourceStatusOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetReadinessCheckResourceStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetReadinessCheckResourceStatusPagesWithContext(ctx context.Context, input *route53recoveryreadiness.GetReadinessCheckResourceStatusInput, fn func(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetReadinessCheckResourceStatusOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetReadinessCheckResourceStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetReadinessCheckResourceStatusWithContext(ctx context.Context, input *route53recoveryreadiness.GetReadinessCheckResourceStatusInput, opts ...request.Option) (*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheckResourceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheckStatus(input *route53recoveryreadiness.GetReadinessCheckStatusInput) (*route53recoveryreadiness.GetReadinessCheckStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckStatusOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheckStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheckStatusPages(input *route53recoveryreadiness.GetReadinessCheckStatusInput, fn func(*route53recoveryreadiness.GetReadinessCheckStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetReadinessCheckStatusOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetReadinessCheckStatusOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetReadinessCheckStatusOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetReadinessCheckStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetReadinessCheckStatusPagesWithContext(ctx context.Context, input *route53recoveryreadiness.GetReadinessCheckStatusInput, fn func(*route53recoveryreadiness.GetReadinessCheckStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetReadinessCheckStatusOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetReadinessCheckStatusOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetReadinessCheckStatusOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetReadinessCheckStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetReadinessCheckStatusWithContext(ctx context.Context, input *route53recoveryreadiness.GetReadinessCheckStatusInput, opts ...request.Option) (*route53recoveryreadiness.GetReadinessCheckStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckStatusOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheckStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetReadinessCheckWithContext(ctx context.Context, input *route53recoveryreadiness.GetReadinessCheckInput, opts ...request.Option) (*route53recoveryreadiness.GetReadinessCheckOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetReadinessCheckOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetReadinessCheckWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetRecoveryGroup(input *route53recoveryreadiness.GetRecoveryGroupInput) (*route53recoveryreadiness.GetRecoveryGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetRecoveryGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetRecoveryGroupReadinessSummary(input *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput) (*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetRecoveryGroupReadinessSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetRecoveryGroupReadinessSummaryPages(input *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput, fn func(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetRecoveryGroupReadinessSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetRecoveryGroupReadinessSummaryPagesWithContext(ctx context.Context, input *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput, fn func(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput{}
	fnCacher := func(out *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.GetRecoveryGroupReadinessSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) GetRecoveryGroupReadinessSummaryWithContext(ctx context.Context, input *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput, opts ...request.Option) (*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetRecoveryGroupReadinessSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetRecoveryGroupWithContext(ctx context.Context, input *route53recoveryreadiness.GetRecoveryGroupInput, opts ...request.Option) (*route53recoveryreadiness.GetRecoveryGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetRecoveryGroupOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetRecoveryGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetResourceSet(input *route53recoveryreadiness.GetResourceSetInput) (*route53recoveryreadiness.GetResourceSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetResourceSetOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetResourceSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) GetResourceSetWithContext(ctx context.Context, input *route53recoveryreadiness.GetResourceSetInput, opts ...request.Option) (*route53recoveryreadiness.GetResourceSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.GetResourceSetOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.GetResourceSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListCells(input *route53recoveryreadiness.ListCellsInput) (*route53recoveryreadiness.ListCellsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListCellsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListCells(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListCellsPages(input *route53recoveryreadiness.ListCellsInput, fn func(*route53recoveryreadiness.ListCellsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListCellsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListCellsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListCellsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListCellsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListCellsPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListCellsInput, fn func(*route53recoveryreadiness.ListCellsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListCellsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListCellsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListCellsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListCellsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListCellsWithContext(ctx context.Context, input *route53recoveryreadiness.ListCellsInput, opts ...request.Option) (*route53recoveryreadiness.ListCellsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListCellsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListCellsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListCrossAccountAuthorizations(input *route53recoveryreadiness.ListCrossAccountAuthorizationsInput) (*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListCrossAccountAuthorizations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListCrossAccountAuthorizationsPages(input *route53recoveryreadiness.ListCrossAccountAuthorizationsInput, fn func(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListCrossAccountAuthorizationsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListCrossAccountAuthorizationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListCrossAccountAuthorizationsPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListCrossAccountAuthorizationsInput, fn func(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListCrossAccountAuthorizationsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListCrossAccountAuthorizationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListCrossAccountAuthorizationsWithContext(ctx context.Context, input *route53recoveryreadiness.ListCrossAccountAuthorizationsInput, opts ...request.Option) (*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListCrossAccountAuthorizationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListReadinessChecks(input *route53recoveryreadiness.ListReadinessChecksInput) (*route53recoveryreadiness.ListReadinessChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListReadinessChecksOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListReadinessChecks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListReadinessChecksPages(input *route53recoveryreadiness.ListReadinessChecksInput, fn func(*route53recoveryreadiness.ListReadinessChecksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListReadinessChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListReadinessChecksOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListReadinessChecksOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListReadinessChecksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListReadinessChecksPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListReadinessChecksInput, fn func(*route53recoveryreadiness.ListReadinessChecksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListReadinessChecksOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListReadinessChecksOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListReadinessChecksOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListReadinessChecksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListReadinessChecksWithContext(ctx context.Context, input *route53recoveryreadiness.ListReadinessChecksInput, opts ...request.Option) (*route53recoveryreadiness.ListReadinessChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListReadinessChecksOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListReadinessChecksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListRecoveryGroups(input *route53recoveryreadiness.ListRecoveryGroupsInput) (*route53recoveryreadiness.ListRecoveryGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListRecoveryGroupsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListRecoveryGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListRecoveryGroupsPages(input *route53recoveryreadiness.ListRecoveryGroupsInput, fn func(*route53recoveryreadiness.ListRecoveryGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListRecoveryGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListRecoveryGroupsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListRecoveryGroupsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListRecoveryGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListRecoveryGroupsPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListRecoveryGroupsInput, fn func(*route53recoveryreadiness.ListRecoveryGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListRecoveryGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListRecoveryGroupsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListRecoveryGroupsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListRecoveryGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListRecoveryGroupsWithContext(ctx context.Context, input *route53recoveryreadiness.ListRecoveryGroupsInput, opts ...request.Option) (*route53recoveryreadiness.ListRecoveryGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListRecoveryGroupsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListRecoveryGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListResourceSets(input *route53recoveryreadiness.ListResourceSetsInput) (*route53recoveryreadiness.ListResourceSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListResourceSetsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListResourceSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListResourceSetsPages(input *route53recoveryreadiness.ListResourceSetsInput, fn func(*route53recoveryreadiness.ListResourceSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListResourceSetsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListResourceSetsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListResourceSetsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListResourceSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListResourceSetsPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListResourceSetsInput, fn func(*route53recoveryreadiness.ListResourceSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListResourceSetsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListResourceSetsOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListResourceSetsOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListResourceSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListResourceSetsWithContext(ctx context.Context, input *route53recoveryreadiness.ListResourceSetsInput, opts ...request.Option) (*route53recoveryreadiness.ListResourceSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListResourceSetsOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListResourceSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListRules(input *route53recoveryreadiness.ListRulesInput) (*route53recoveryreadiness.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListRulesOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListRulesPages(input *route53recoveryreadiness.ListRulesInput, fn func(*route53recoveryreadiness.ListRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListRulesOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListRulesOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListRulesPagesWithContext(ctx context.Context, input *route53recoveryreadiness.ListRulesInput, fn func(*route53recoveryreadiness.ListRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoveryreadiness.ListRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoveryreadiness.ListRulesOutput{}
	fnCacher := func(out *route53recoveryreadiness.ListRulesOutput, more bool) bool {
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
	if err := c.Route53RecoveryReadinessAPI.ListRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryReadiness) ListRulesWithContext(ctx context.Context, input *route53recoveryreadiness.ListRulesInput, opts ...request.Option) (*route53recoveryreadiness.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListRulesOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListTagsForResources(input *route53recoveryreadiness.ListTagsForResourcesInput) (*route53recoveryreadiness.ListTagsForResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListTagsForResourcesOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListTagsForResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryReadiness) ListTagsForResourcesWithContext(ctx context.Context, input *route53recoveryreadiness.ListTagsForResourcesInput, opts ...request.Option) (*route53recoveryreadiness.ListTagsForResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoveryreadiness.ListTagsForResourcesOutput), nil
	}
	out, err := c.Route53RecoveryReadinessAPI.ListTagsForResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

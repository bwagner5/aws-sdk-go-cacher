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
package supportcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Support struct {
	supportiface.SupportAPI
	cache *cache.Cache
}

func New(supportapi supportiface.SupportAPI) *Support {
	return &Support{
		SupportAPI: supportapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Support) DescribeAttachment(input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeAttachmentOutput), nil
	}
	out, err := c.SupportAPI.DescribeAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeAttachmentWithContext(ctx context.Context, input *support.DescribeAttachmentInput, opts ...request.Option) (*support.DescribeAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeAttachmentOutput), nil
	}
	out, err := c.SupportAPI.DescribeAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeCases(input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeCasesOutput), nil
	}
	out, err := c.SupportAPI.DescribeCases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeCasesPages(input *support.DescribeCasesInput, fn func(*support.DescribeCasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*support.DescribeCasesOutput), false)
		return nil
	}
	cachable := true
	output := &support.DescribeCasesOutput{}
	fnCacher := func(out *support.DescribeCasesOutput, more bool) bool {
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
	if err := c.SupportAPI.DescribeCasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Support) DescribeCasesPagesWithContext(ctx context.Context, input *support.DescribeCasesInput, fn func(*support.DescribeCasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*support.DescribeCasesOutput), false)
		return nil
	}
	cachable := true
	output := &support.DescribeCasesOutput{}
	fnCacher := func(out *support.DescribeCasesOutput, more bool) bool {
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
	if err := c.SupportAPI.DescribeCasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Support) DescribeCasesWithContext(ctx context.Context, input *support.DescribeCasesInput, opts ...request.Option) (*support.DescribeCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeCasesOutput), nil
	}
	out, err := c.SupportAPI.DescribeCasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeCommunications(input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeCommunicationsOutput), nil
	}
	out, err := c.SupportAPI.DescribeCommunications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeCommunicationsPages(input *support.DescribeCommunicationsInput, fn func(*support.DescribeCommunicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*support.DescribeCommunicationsOutput), false)
		return nil
	}
	cachable := true
	output := &support.DescribeCommunicationsOutput{}
	fnCacher := func(out *support.DescribeCommunicationsOutput, more bool) bool {
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
	if err := c.SupportAPI.DescribeCommunicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Support) DescribeCommunicationsPagesWithContext(ctx context.Context, input *support.DescribeCommunicationsInput, fn func(*support.DescribeCommunicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*support.DescribeCommunicationsOutput), false)
		return nil
	}
	cachable := true
	output := &support.DescribeCommunicationsOutput{}
	fnCacher := func(out *support.DescribeCommunicationsOutput, more bool) bool {
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
	if err := c.SupportAPI.DescribeCommunicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Support) DescribeCommunicationsWithContext(ctx context.Context, input *support.DescribeCommunicationsInput, opts ...request.Option) (*support.DescribeCommunicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeCommunicationsOutput), nil
	}
	out, err := c.SupportAPI.DescribeCommunicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeServices(input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeServicesOutput), nil
	}
	out, err := c.SupportAPI.DescribeServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeServicesWithContext(ctx context.Context, input *support.DescribeServicesInput, opts ...request.Option) (*support.DescribeServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeServicesOutput), nil
	}
	out, err := c.SupportAPI.DescribeServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeSeverityLevels(input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeSeverityLevelsOutput), nil
	}
	out, err := c.SupportAPI.DescribeSeverityLevels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeSeverityLevelsWithContext(ctx context.Context, input *support.DescribeSeverityLevelsInput, opts ...request.Option) (*support.DescribeSeverityLevelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeSeverityLevelsOutput), nil
	}
	out, err := c.SupportAPI.DescribeSeverityLevelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatuses(input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckRefreshStatuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckResult(input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckResultOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckResult(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckResultWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckResultInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckResultOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckResultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckSummaries(input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckSummariesOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorCheckSummariesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorCheckSummariesOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorCheckSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorChecks(input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorChecksOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorChecks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Support) DescribeTrustedAdvisorChecksWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorChecksInput, opts ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*support.DescribeTrustedAdvisorChecksOutput), nil
	}
	out, err := c.SupportAPI.DescribeTrustedAdvisorChecksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package prometheusservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/aws/aws-sdk-go/service/prometheusservice/prometheusserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PrometheusService struct {
	prometheusserviceiface.PrometheusServiceAPI
	cache *cache.Cache
}

func New(prometheusserviceapi prometheusserviceiface.PrometheusServiceAPI) *PrometheusService {
	return &PrometheusService{
		PrometheusServiceAPI: prometheusserviceapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PrometheusService) DescribeAlertManagerDefinition(input *prometheusservice.DescribeAlertManagerDefinitionInput) (*prometheusservice.DescribeAlertManagerDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeAlertManagerDefinitionOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeAlertManagerDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeAlertManagerDefinitionWithContext(ctx context.Context, input *prometheusservice.DescribeAlertManagerDefinitionInput, opts ...request.Option) (*prometheusservice.DescribeAlertManagerDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeAlertManagerDefinitionOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeAlertManagerDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeLoggingConfiguration(input *prometheusservice.DescribeLoggingConfigurationInput) (*prometheusservice.DescribeLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeLoggingConfigurationOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeLoggingConfigurationWithContext(ctx context.Context, input *prometheusservice.DescribeLoggingConfigurationInput, opts ...request.Option) (*prometheusservice.DescribeLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeLoggingConfigurationOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeRuleGroupsNamespace(input *prometheusservice.DescribeRuleGroupsNamespaceInput) (*prometheusservice.DescribeRuleGroupsNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeRuleGroupsNamespaceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeRuleGroupsNamespace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeRuleGroupsNamespaceWithContext(ctx context.Context, input *prometheusservice.DescribeRuleGroupsNamespaceInput, opts ...request.Option) (*prometheusservice.DescribeRuleGroupsNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeRuleGroupsNamespaceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeRuleGroupsNamespaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeWorkspace(input *prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeWorkspaceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeWorkspace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) DescribeWorkspaceWithContext(ctx context.Context, input *prometheusservice.DescribeWorkspaceInput, opts ...request.Option) (*prometheusservice.DescribeWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.DescribeWorkspaceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.DescribeWorkspaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListRuleGroupsNamespaces(input *prometheusservice.ListRuleGroupsNamespacesInput) (*prometheusservice.ListRuleGroupsNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListRuleGroupsNamespacesOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListRuleGroupsNamespaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListRuleGroupsNamespacesPages(input *prometheusservice.ListRuleGroupsNamespacesInput, fn func(*prometheusservice.ListRuleGroupsNamespacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*prometheusservice.ListRuleGroupsNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &prometheusservice.ListRuleGroupsNamespacesOutput{}
	fnCacher := func(out *prometheusservice.ListRuleGroupsNamespacesOutput, more bool) bool {
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
	if err := c.PrometheusServiceAPI.ListRuleGroupsNamespacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrometheusService) ListRuleGroupsNamespacesPagesWithContext(ctx context.Context, input *prometheusservice.ListRuleGroupsNamespacesInput, fn func(*prometheusservice.ListRuleGroupsNamespacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*prometheusservice.ListRuleGroupsNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &prometheusservice.ListRuleGroupsNamespacesOutput{}
	fnCacher := func(out *prometheusservice.ListRuleGroupsNamespacesOutput, more bool) bool {
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
	if err := c.PrometheusServiceAPI.ListRuleGroupsNamespacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrometheusService) ListRuleGroupsNamespacesWithContext(ctx context.Context, input *prometheusservice.ListRuleGroupsNamespacesInput, opts ...request.Option) (*prometheusservice.ListRuleGroupsNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListRuleGroupsNamespacesOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListRuleGroupsNamespacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListTagsForResource(input *prometheusservice.ListTagsForResourceInput) (*prometheusservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListTagsForResourceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListTagsForResourceWithContext(ctx context.Context, input *prometheusservice.ListTagsForResourceInput, opts ...request.Option) (*prometheusservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListTagsForResourceOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListWorkspaces(input *prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListWorkspacesOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListWorkspaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrometheusService) ListWorkspacesPages(input *prometheusservice.ListWorkspacesInput, fn func(*prometheusservice.ListWorkspacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*prometheusservice.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &prometheusservice.ListWorkspacesOutput{}
	fnCacher := func(out *prometheusservice.ListWorkspacesOutput, more bool) bool {
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
	if err := c.PrometheusServiceAPI.ListWorkspacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrometheusService) ListWorkspacesPagesWithContext(ctx context.Context, input *prometheusservice.ListWorkspacesInput, fn func(*prometheusservice.ListWorkspacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*prometheusservice.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &prometheusservice.ListWorkspacesOutput{}
	fnCacher := func(out *prometheusservice.ListWorkspacesOutput, more bool) bool {
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
	if err := c.PrometheusServiceAPI.ListWorkspacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrometheusService) ListWorkspacesWithContext(ctx context.Context, input *prometheusservice.ListWorkspacesInput, opts ...request.Option) (*prometheusservice.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*prometheusservice.ListWorkspacesOutput), nil
	}
	out, err := c.PrometheusServiceAPI.ListWorkspacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

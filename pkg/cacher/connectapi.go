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
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Connect struct {
	connectiface.ConnectAPI
	cache *cache.Cache
}

func NewConnect(connectapi connectiface.ConnectAPI) *Connect {
	return &Connect{
		ConnectAPI: connectapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Connect) DescribeAgentStatus(input *connect.DescribeAgentStatusInput) (*connect.DescribeAgentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeAgentStatusOutput), nil
	}
	out, err := c.ConnectAPI.DescribeAgentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeAgentStatusWithContext(ctx context.Context, input *connect.DescribeAgentStatusInput, opts ...request.Option) (*connect.DescribeAgentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeAgentStatusOutput), nil
	}
	out, err := c.ConnectAPI.DescribeAgentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContact(input *connect.DescribeContactInput) (*connect.DescribeContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContactFlow(input *connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactFlowOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContactFlow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContactFlowModule(input *connect.DescribeContactFlowModuleInput) (*connect.DescribeContactFlowModuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactFlowModuleOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContactFlowModule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContactFlowModuleWithContext(ctx context.Context, input *connect.DescribeContactFlowModuleInput, opts ...request.Option) (*connect.DescribeContactFlowModuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactFlowModuleOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContactFlowModuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContactFlowWithContext(ctx context.Context, input *connect.DescribeContactFlowInput, opts ...request.Option) (*connect.DescribeContactFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactFlowOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContactFlowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeContactWithContext(ctx context.Context, input *connect.DescribeContactInput, opts ...request.Option) (*connect.DescribeContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeContactOutput), nil
	}
	out, err := c.ConnectAPI.DescribeContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeHoursOfOperation(input *connect.DescribeHoursOfOperationInput) (*connect.DescribeHoursOfOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeHoursOfOperationOutput), nil
	}
	out, err := c.ConnectAPI.DescribeHoursOfOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeHoursOfOperationWithContext(ctx context.Context, input *connect.DescribeHoursOfOperationInput, opts ...request.Option) (*connect.DescribeHoursOfOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeHoursOfOperationOutput), nil
	}
	out, err := c.ConnectAPI.DescribeHoursOfOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstance(input *connect.DescribeInstanceInput) (*connect.DescribeInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstanceAttribute(input *connect.DescribeInstanceAttributeInput) (*connect.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstanceAttributeWithContext(ctx context.Context, input *connect.DescribeInstanceAttributeInput, opts ...request.Option) (*connect.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstanceStorageConfig(input *connect.DescribeInstanceStorageConfigInput) (*connect.DescribeInstanceStorageConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceStorageConfigOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstanceStorageConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstanceStorageConfigWithContext(ctx context.Context, input *connect.DescribeInstanceStorageConfigInput, opts ...request.Option) (*connect.DescribeInstanceStorageConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceStorageConfigOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstanceStorageConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeInstanceWithContext(ctx context.Context, input *connect.DescribeInstanceInput, opts ...request.Option) (*connect.DescribeInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeInstanceOutput), nil
	}
	out, err := c.ConnectAPI.DescribeInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribePhoneNumber(input *connect.DescribePhoneNumberInput) (*connect.DescribePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribePhoneNumberOutput), nil
	}
	out, err := c.ConnectAPI.DescribePhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribePhoneNumberWithContext(ctx context.Context, input *connect.DescribePhoneNumberInput, opts ...request.Option) (*connect.DescribePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribePhoneNumberOutput), nil
	}
	out, err := c.ConnectAPI.DescribePhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeQueue(input *connect.DescribeQueueInput) (*connect.DescribeQueueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeQueueOutput), nil
	}
	out, err := c.ConnectAPI.DescribeQueue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeQueueWithContext(ctx context.Context, input *connect.DescribeQueueInput, opts ...request.Option) (*connect.DescribeQueueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeQueueOutput), nil
	}
	out, err := c.ConnectAPI.DescribeQueueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeQuickConnect(input *connect.DescribeQuickConnectInput) (*connect.DescribeQuickConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeQuickConnectOutput), nil
	}
	out, err := c.ConnectAPI.DescribeQuickConnect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeQuickConnectWithContext(ctx context.Context, input *connect.DescribeQuickConnectInput, opts ...request.Option) (*connect.DescribeQuickConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeQuickConnectOutput), nil
	}
	out, err := c.ConnectAPI.DescribeQuickConnectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeRoutingProfile(input *connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeRoutingProfileOutput), nil
	}
	out, err := c.ConnectAPI.DescribeRoutingProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeRoutingProfileWithContext(ctx context.Context, input *connect.DescribeRoutingProfileInput, opts ...request.Option) (*connect.DescribeRoutingProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeRoutingProfileOutput), nil
	}
	out, err := c.ConnectAPI.DescribeRoutingProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeRule(input *connect.DescribeRuleInput) (*connect.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeRuleOutput), nil
	}
	out, err := c.ConnectAPI.DescribeRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeRuleWithContext(ctx context.Context, input *connect.DescribeRuleInput, opts ...request.Option) (*connect.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeRuleOutput), nil
	}
	out, err := c.ConnectAPI.DescribeRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeSecurityProfile(input *connect.DescribeSecurityProfileInput) (*connect.DescribeSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeSecurityProfileOutput), nil
	}
	out, err := c.ConnectAPI.DescribeSecurityProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeSecurityProfileWithContext(ctx context.Context, input *connect.DescribeSecurityProfileInput, opts ...request.Option) (*connect.DescribeSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeSecurityProfileOutput), nil
	}
	out, err := c.ConnectAPI.DescribeSecurityProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeTrafficDistributionGroup(input *connect.DescribeTrafficDistributionGroupInput) (*connect.DescribeTrafficDistributionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeTrafficDistributionGroupOutput), nil
	}
	out, err := c.ConnectAPI.DescribeTrafficDistributionGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeTrafficDistributionGroupWithContext(ctx context.Context, input *connect.DescribeTrafficDistributionGroupInput, opts ...request.Option) (*connect.DescribeTrafficDistributionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeTrafficDistributionGroupOutput), nil
	}
	out, err := c.ConnectAPI.DescribeTrafficDistributionGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUser(input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUserHierarchyGroup(input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserHierarchyGroupOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUserHierarchyGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUserHierarchyGroupWithContext(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput, opts ...request.Option) (*connect.DescribeUserHierarchyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserHierarchyGroupOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUserHierarchyGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUserHierarchyStructure(input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserHierarchyStructureOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUserHierarchyStructure(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUserHierarchyStructureWithContext(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput, opts ...request.Option) (*connect.DescribeUserHierarchyStructureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserHierarchyStructureOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUserHierarchyStructureWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeUserWithContext(ctx context.Context, input *connect.DescribeUserInput, opts ...request.Option) (*connect.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeUserOutput), nil
	}
	out, err := c.ConnectAPI.DescribeUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeVocabulary(input *connect.DescribeVocabularyInput) (*connect.DescribeVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeVocabularyOutput), nil
	}
	out, err := c.ConnectAPI.DescribeVocabulary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) DescribeVocabularyWithContext(ctx context.Context, input *connect.DescribeVocabularyInput, opts ...request.Option) (*connect.DescribeVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.DescribeVocabularyOutput), nil
	}
	out, err := c.ConnectAPI.DescribeVocabularyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetContactAttributes(input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetContactAttributesOutput), nil
	}
	out, err := c.ConnectAPI.GetContactAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetContactAttributesWithContext(ctx context.Context, input *connect.GetContactAttributesInput, opts ...request.Option) (*connect.GetContactAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetContactAttributesOutput), nil
	}
	out, err := c.ConnectAPI.GetContactAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetCurrentMetricData(input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetCurrentMetricDataOutput), nil
	}
	out, err := c.ConnectAPI.GetCurrentMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetCurrentMetricDataPages(input *connect.GetCurrentMetricDataInput, fn func(*connect.GetCurrentMetricDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetCurrentMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetCurrentMetricDataOutput{}
	fnCacher := func(out *connect.GetCurrentMetricDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetCurrentMetricDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetCurrentMetricDataPagesWithContext(ctx context.Context, input *connect.GetCurrentMetricDataInput, fn func(*connect.GetCurrentMetricDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetCurrentMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetCurrentMetricDataOutput{}
	fnCacher := func(out *connect.GetCurrentMetricDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetCurrentMetricDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetCurrentMetricDataWithContext(ctx context.Context, input *connect.GetCurrentMetricDataInput, opts ...request.Option) (*connect.GetCurrentMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetCurrentMetricDataOutput), nil
	}
	out, err := c.ConnectAPI.GetCurrentMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetCurrentUserData(input *connect.GetCurrentUserDataInput) (*connect.GetCurrentUserDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetCurrentUserDataOutput), nil
	}
	out, err := c.ConnectAPI.GetCurrentUserData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetCurrentUserDataPages(input *connect.GetCurrentUserDataInput, fn func(*connect.GetCurrentUserDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetCurrentUserDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetCurrentUserDataOutput{}
	fnCacher := func(out *connect.GetCurrentUserDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetCurrentUserDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetCurrentUserDataPagesWithContext(ctx context.Context, input *connect.GetCurrentUserDataInput, fn func(*connect.GetCurrentUserDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetCurrentUserDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetCurrentUserDataOutput{}
	fnCacher := func(out *connect.GetCurrentUserDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetCurrentUserDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetCurrentUserDataWithContext(ctx context.Context, input *connect.GetCurrentUserDataInput, opts ...request.Option) (*connect.GetCurrentUserDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetCurrentUserDataOutput), nil
	}
	out, err := c.ConnectAPI.GetCurrentUserDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetFederationToken(input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetFederationTokenOutput), nil
	}
	out, err := c.ConnectAPI.GetFederationToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetFederationTokenWithContext(ctx context.Context, input *connect.GetFederationTokenInput, opts ...request.Option) (*connect.GetFederationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetFederationTokenOutput), nil
	}
	out, err := c.ConnectAPI.GetFederationTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetMetricData(input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetMetricDataOutput), nil
	}
	out, err := c.ConnectAPI.GetMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetMetricDataPages(input *connect.GetMetricDataInput, fn func(*connect.GetMetricDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetMetricDataOutput{}
	fnCacher := func(out *connect.GetMetricDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetMetricDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetMetricDataPagesWithContext(ctx context.Context, input *connect.GetMetricDataInput, fn func(*connect.GetMetricDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.GetMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &connect.GetMetricDataOutput{}
	fnCacher := func(out *connect.GetMetricDataOutput, more bool) bool {
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
	if err := c.ConnectAPI.GetMetricDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) GetMetricDataWithContext(ctx context.Context, input *connect.GetMetricDataInput, opts ...request.Option) (*connect.GetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetMetricDataOutput), nil
	}
	out, err := c.ConnectAPI.GetMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetTaskTemplate(input *connect.GetTaskTemplateInput) (*connect.GetTaskTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetTaskTemplateOutput), nil
	}
	out, err := c.ConnectAPI.GetTaskTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetTaskTemplateWithContext(ctx context.Context, input *connect.GetTaskTemplateInput, opts ...request.Option) (*connect.GetTaskTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetTaskTemplateOutput), nil
	}
	out, err := c.ConnectAPI.GetTaskTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetTrafficDistribution(input *connect.GetTrafficDistributionInput) (*connect.GetTrafficDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetTrafficDistributionOutput), nil
	}
	out, err := c.ConnectAPI.GetTrafficDistribution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) GetTrafficDistributionWithContext(ctx context.Context, input *connect.GetTrafficDistributionInput, opts ...request.Option) (*connect.GetTrafficDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.GetTrafficDistributionOutput), nil
	}
	out, err := c.ConnectAPI.GetTrafficDistributionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListAgentStatuses(input *connect.ListAgentStatusesInput) (*connect.ListAgentStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListAgentStatusesOutput), nil
	}
	out, err := c.ConnectAPI.ListAgentStatuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListAgentStatusesPages(input *connect.ListAgentStatusesInput, fn func(*connect.ListAgentStatusesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListAgentStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListAgentStatusesOutput{}
	fnCacher := func(out *connect.ListAgentStatusesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListAgentStatusesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListAgentStatusesPagesWithContext(ctx context.Context, input *connect.ListAgentStatusesInput, fn func(*connect.ListAgentStatusesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListAgentStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListAgentStatusesOutput{}
	fnCacher := func(out *connect.ListAgentStatusesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListAgentStatusesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListAgentStatusesWithContext(ctx context.Context, input *connect.ListAgentStatusesInput, opts ...request.Option) (*connect.ListAgentStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListAgentStatusesOutput), nil
	}
	out, err := c.ConnectAPI.ListAgentStatusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListApprovedOrigins(input *connect.ListApprovedOriginsInput) (*connect.ListApprovedOriginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListApprovedOriginsOutput), nil
	}
	out, err := c.ConnectAPI.ListApprovedOrigins(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListApprovedOriginsPages(input *connect.ListApprovedOriginsInput, fn func(*connect.ListApprovedOriginsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListApprovedOriginsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListApprovedOriginsOutput{}
	fnCacher := func(out *connect.ListApprovedOriginsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListApprovedOriginsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListApprovedOriginsPagesWithContext(ctx context.Context, input *connect.ListApprovedOriginsInput, fn func(*connect.ListApprovedOriginsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListApprovedOriginsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListApprovedOriginsOutput{}
	fnCacher := func(out *connect.ListApprovedOriginsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListApprovedOriginsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListApprovedOriginsWithContext(ctx context.Context, input *connect.ListApprovedOriginsInput, opts ...request.Option) (*connect.ListApprovedOriginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListApprovedOriginsOutput), nil
	}
	out, err := c.ConnectAPI.ListApprovedOriginsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListBots(input *connect.ListBotsInput) (*connect.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListBotsOutput), nil
	}
	out, err := c.ConnectAPI.ListBots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListBotsPages(input *connect.ListBotsInput, fn func(*connect.ListBotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListBotsOutput{}
	fnCacher := func(out *connect.ListBotsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListBotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListBotsPagesWithContext(ctx context.Context, input *connect.ListBotsInput, fn func(*connect.ListBotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListBotsOutput{}
	fnCacher := func(out *connect.ListBotsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListBotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListBotsWithContext(ctx context.Context, input *connect.ListBotsInput, opts ...request.Option) (*connect.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListBotsOutput), nil
	}
	out, err := c.ConnectAPI.ListBotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactFlowModules(input *connect.ListContactFlowModulesInput) (*connect.ListContactFlowModulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactFlowModulesOutput), nil
	}
	out, err := c.ConnectAPI.ListContactFlowModules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactFlowModulesPages(input *connect.ListContactFlowModulesInput, fn func(*connect.ListContactFlowModulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactFlowModulesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactFlowModulesOutput{}
	fnCacher := func(out *connect.ListContactFlowModulesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactFlowModulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactFlowModulesPagesWithContext(ctx context.Context, input *connect.ListContactFlowModulesInput, fn func(*connect.ListContactFlowModulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactFlowModulesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactFlowModulesOutput{}
	fnCacher := func(out *connect.ListContactFlowModulesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactFlowModulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactFlowModulesWithContext(ctx context.Context, input *connect.ListContactFlowModulesInput, opts ...request.Option) (*connect.ListContactFlowModulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactFlowModulesOutput), nil
	}
	out, err := c.ConnectAPI.ListContactFlowModulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactFlows(input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactFlowsOutput), nil
	}
	out, err := c.ConnectAPI.ListContactFlows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactFlowsPages(input *connect.ListContactFlowsInput, fn func(*connect.ListContactFlowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactFlowsOutput{}
	fnCacher := func(out *connect.ListContactFlowsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactFlowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactFlowsPagesWithContext(ctx context.Context, input *connect.ListContactFlowsInput, fn func(*connect.ListContactFlowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactFlowsOutput{}
	fnCacher := func(out *connect.ListContactFlowsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactFlowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactFlowsWithContext(ctx context.Context, input *connect.ListContactFlowsInput, opts ...request.Option) (*connect.ListContactFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactFlowsOutput), nil
	}
	out, err := c.ConnectAPI.ListContactFlowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactReferences(input *connect.ListContactReferencesInput) (*connect.ListContactReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactReferencesOutput), nil
	}
	out, err := c.ConnectAPI.ListContactReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListContactReferencesPages(input *connect.ListContactReferencesInput, fn func(*connect.ListContactReferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactReferencesOutput{}
	fnCacher := func(out *connect.ListContactReferencesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactReferencesPagesWithContext(ctx context.Context, input *connect.ListContactReferencesInput, fn func(*connect.ListContactReferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListContactReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListContactReferencesOutput{}
	fnCacher := func(out *connect.ListContactReferencesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListContactReferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListContactReferencesWithContext(ctx context.Context, input *connect.ListContactReferencesInput, opts ...request.Option) (*connect.ListContactReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListContactReferencesOutput), nil
	}
	out, err := c.ConnectAPI.ListContactReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListDefaultVocabularies(input *connect.ListDefaultVocabulariesInput) (*connect.ListDefaultVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListDefaultVocabulariesOutput), nil
	}
	out, err := c.ConnectAPI.ListDefaultVocabularies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListDefaultVocabulariesPages(input *connect.ListDefaultVocabulariesInput, fn func(*connect.ListDefaultVocabulariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListDefaultVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListDefaultVocabulariesOutput{}
	fnCacher := func(out *connect.ListDefaultVocabulariesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListDefaultVocabulariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListDefaultVocabulariesPagesWithContext(ctx context.Context, input *connect.ListDefaultVocabulariesInput, fn func(*connect.ListDefaultVocabulariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListDefaultVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListDefaultVocabulariesOutput{}
	fnCacher := func(out *connect.ListDefaultVocabulariesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListDefaultVocabulariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListDefaultVocabulariesWithContext(ctx context.Context, input *connect.ListDefaultVocabulariesInput, opts ...request.Option) (*connect.ListDefaultVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListDefaultVocabulariesOutput), nil
	}
	out, err := c.ConnectAPI.ListDefaultVocabulariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListHoursOfOperations(input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListHoursOfOperationsOutput), nil
	}
	out, err := c.ConnectAPI.ListHoursOfOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListHoursOfOperationsPages(input *connect.ListHoursOfOperationsInput, fn func(*connect.ListHoursOfOperationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListHoursOfOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListHoursOfOperationsOutput{}
	fnCacher := func(out *connect.ListHoursOfOperationsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListHoursOfOperationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListHoursOfOperationsPagesWithContext(ctx context.Context, input *connect.ListHoursOfOperationsInput, fn func(*connect.ListHoursOfOperationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListHoursOfOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListHoursOfOperationsOutput{}
	fnCacher := func(out *connect.ListHoursOfOperationsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListHoursOfOperationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListHoursOfOperationsWithContext(ctx context.Context, input *connect.ListHoursOfOperationsInput, opts ...request.Option) (*connect.ListHoursOfOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListHoursOfOperationsOutput), nil
	}
	out, err := c.ConnectAPI.ListHoursOfOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstanceAttributes(input *connect.ListInstanceAttributesInput) (*connect.ListInstanceAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstanceAttributesOutput), nil
	}
	out, err := c.ConnectAPI.ListInstanceAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstanceAttributesPages(input *connect.ListInstanceAttributesInput, fn func(*connect.ListInstanceAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstanceAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstanceAttributesOutput{}
	fnCacher := func(out *connect.ListInstanceAttributesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstanceAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstanceAttributesPagesWithContext(ctx context.Context, input *connect.ListInstanceAttributesInput, fn func(*connect.ListInstanceAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstanceAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstanceAttributesOutput{}
	fnCacher := func(out *connect.ListInstanceAttributesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstanceAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstanceAttributesWithContext(ctx context.Context, input *connect.ListInstanceAttributesInput, opts ...request.Option) (*connect.ListInstanceAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstanceAttributesOutput), nil
	}
	out, err := c.ConnectAPI.ListInstanceAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstanceStorageConfigs(input *connect.ListInstanceStorageConfigsInput) (*connect.ListInstanceStorageConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstanceStorageConfigsOutput), nil
	}
	out, err := c.ConnectAPI.ListInstanceStorageConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstanceStorageConfigsPages(input *connect.ListInstanceStorageConfigsInput, fn func(*connect.ListInstanceStorageConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstanceStorageConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstanceStorageConfigsOutput{}
	fnCacher := func(out *connect.ListInstanceStorageConfigsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstanceStorageConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstanceStorageConfigsPagesWithContext(ctx context.Context, input *connect.ListInstanceStorageConfigsInput, fn func(*connect.ListInstanceStorageConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstanceStorageConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstanceStorageConfigsOutput{}
	fnCacher := func(out *connect.ListInstanceStorageConfigsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstanceStorageConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstanceStorageConfigsWithContext(ctx context.Context, input *connect.ListInstanceStorageConfigsInput, opts ...request.Option) (*connect.ListInstanceStorageConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstanceStorageConfigsOutput), nil
	}
	out, err := c.ConnectAPI.ListInstanceStorageConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstances(input *connect.ListInstancesInput) (*connect.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstancesOutput), nil
	}
	out, err := c.ConnectAPI.ListInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListInstancesPages(input *connect.ListInstancesInput, fn func(*connect.ListInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstancesOutput{}
	fnCacher := func(out *connect.ListInstancesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstancesPagesWithContext(ctx context.Context, input *connect.ListInstancesInput, fn func(*connect.ListInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListInstancesOutput{}
	fnCacher := func(out *connect.ListInstancesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListInstancesWithContext(ctx context.Context, input *connect.ListInstancesInput, opts ...request.Option) (*connect.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListInstancesOutput), nil
	}
	out, err := c.ConnectAPI.ListInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListIntegrationAssociations(input *connect.ListIntegrationAssociationsInput) (*connect.ListIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListIntegrationAssociationsOutput), nil
	}
	out, err := c.ConnectAPI.ListIntegrationAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListIntegrationAssociationsPages(input *connect.ListIntegrationAssociationsInput, fn func(*connect.ListIntegrationAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListIntegrationAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListIntegrationAssociationsOutput{}
	fnCacher := func(out *connect.ListIntegrationAssociationsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListIntegrationAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListIntegrationAssociationsPagesWithContext(ctx context.Context, input *connect.ListIntegrationAssociationsInput, fn func(*connect.ListIntegrationAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListIntegrationAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListIntegrationAssociationsOutput{}
	fnCacher := func(out *connect.ListIntegrationAssociationsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListIntegrationAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListIntegrationAssociationsWithContext(ctx context.Context, input *connect.ListIntegrationAssociationsInput, opts ...request.Option) (*connect.ListIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListIntegrationAssociationsOutput), nil
	}
	out, err := c.ConnectAPI.ListIntegrationAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListLambdaFunctions(input *connect.ListLambdaFunctionsInput) (*connect.ListLambdaFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListLambdaFunctionsOutput), nil
	}
	out, err := c.ConnectAPI.ListLambdaFunctions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListLambdaFunctionsPages(input *connect.ListLambdaFunctionsInput, fn func(*connect.ListLambdaFunctionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListLambdaFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListLambdaFunctionsOutput{}
	fnCacher := func(out *connect.ListLambdaFunctionsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListLambdaFunctionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListLambdaFunctionsPagesWithContext(ctx context.Context, input *connect.ListLambdaFunctionsInput, fn func(*connect.ListLambdaFunctionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListLambdaFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListLambdaFunctionsOutput{}
	fnCacher := func(out *connect.ListLambdaFunctionsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListLambdaFunctionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListLambdaFunctionsWithContext(ctx context.Context, input *connect.ListLambdaFunctionsInput, opts ...request.Option) (*connect.ListLambdaFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListLambdaFunctionsOutput), nil
	}
	out, err := c.ConnectAPI.ListLambdaFunctionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListLexBots(input *connect.ListLexBotsInput) (*connect.ListLexBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListLexBotsOutput), nil
	}
	out, err := c.ConnectAPI.ListLexBots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListLexBotsPages(input *connect.ListLexBotsInput, fn func(*connect.ListLexBotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListLexBotsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListLexBotsOutput{}
	fnCacher := func(out *connect.ListLexBotsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListLexBotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListLexBotsPagesWithContext(ctx context.Context, input *connect.ListLexBotsInput, fn func(*connect.ListLexBotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListLexBotsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListLexBotsOutput{}
	fnCacher := func(out *connect.ListLexBotsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListLexBotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListLexBotsWithContext(ctx context.Context, input *connect.ListLexBotsInput, opts ...request.Option) (*connect.ListLexBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListLexBotsOutput), nil
	}
	out, err := c.ConnectAPI.ListLexBotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPhoneNumbers(input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPhoneNumbersOutput), nil
	}
	out, err := c.ConnectAPI.ListPhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPhoneNumbersPages(input *connect.ListPhoneNumbersInput, fn func(*connect.ListPhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListPhoneNumbersOutput{}
	fnCacher := func(out *connect.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListPhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPhoneNumbersPagesWithContext(ctx context.Context, input *connect.ListPhoneNumbersInput, fn func(*connect.ListPhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListPhoneNumbersOutput{}
	fnCacher := func(out *connect.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListPhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPhoneNumbersV2(input *connect.ListPhoneNumbersV2Input) (*connect.ListPhoneNumbersV2Output, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPhoneNumbersV2Output), nil
	}
	out, err := c.ConnectAPI.ListPhoneNumbersV2(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPhoneNumbersV2Pages(input *connect.ListPhoneNumbersV2Input, fn func(*connect.ListPhoneNumbersV2Output, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPhoneNumbersV2Output), false)
		return nil
	}
	cachable := true
	output := &connect.ListPhoneNumbersV2Output{}
	fnCacher := func(out *connect.ListPhoneNumbersV2Output, more bool) bool {
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
	if err := c.ConnectAPI.ListPhoneNumbersV2Pages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPhoneNumbersV2PagesWithContext(ctx context.Context, input *connect.ListPhoneNumbersV2Input, fn func(*connect.ListPhoneNumbersV2Output, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPhoneNumbersV2Output), false)
		return nil
	}
	cachable := true
	output := &connect.ListPhoneNumbersV2Output{}
	fnCacher := func(out *connect.ListPhoneNumbersV2Output, more bool) bool {
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
	if err := c.ConnectAPI.ListPhoneNumbersV2PagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPhoneNumbersV2WithContext(ctx context.Context, input *connect.ListPhoneNumbersV2Input, opts ...request.Option) (*connect.ListPhoneNumbersV2Output, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPhoneNumbersV2Output), nil
	}
	out, err := c.ConnectAPI.ListPhoneNumbersV2WithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPhoneNumbersWithContext(ctx context.Context, input *connect.ListPhoneNumbersInput, opts ...request.Option) (*connect.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPhoneNumbersOutput), nil
	}
	out, err := c.ConnectAPI.ListPhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPrompts(input *connect.ListPromptsInput) (*connect.ListPromptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPromptsOutput), nil
	}
	out, err := c.ConnectAPI.ListPrompts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListPromptsPages(input *connect.ListPromptsInput, fn func(*connect.ListPromptsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPromptsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListPromptsOutput{}
	fnCacher := func(out *connect.ListPromptsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListPromptsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPromptsPagesWithContext(ctx context.Context, input *connect.ListPromptsInput, fn func(*connect.ListPromptsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListPromptsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListPromptsOutput{}
	fnCacher := func(out *connect.ListPromptsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListPromptsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListPromptsWithContext(ctx context.Context, input *connect.ListPromptsInput, opts ...request.Option) (*connect.ListPromptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListPromptsOutput), nil
	}
	out, err := c.ConnectAPI.ListPromptsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQueueQuickConnects(input *connect.ListQueueQuickConnectsInput) (*connect.ListQueueQuickConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQueueQuickConnectsOutput), nil
	}
	out, err := c.ConnectAPI.ListQueueQuickConnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQueueQuickConnectsPages(input *connect.ListQueueQuickConnectsInput, fn func(*connect.ListQueueQuickConnectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQueueQuickConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQueueQuickConnectsOutput{}
	fnCacher := func(out *connect.ListQueueQuickConnectsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQueueQuickConnectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQueueQuickConnectsPagesWithContext(ctx context.Context, input *connect.ListQueueQuickConnectsInput, fn func(*connect.ListQueueQuickConnectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQueueQuickConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQueueQuickConnectsOutput{}
	fnCacher := func(out *connect.ListQueueQuickConnectsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQueueQuickConnectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQueueQuickConnectsWithContext(ctx context.Context, input *connect.ListQueueQuickConnectsInput, opts ...request.Option) (*connect.ListQueueQuickConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQueueQuickConnectsOutput), nil
	}
	out, err := c.ConnectAPI.ListQueueQuickConnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQueues(input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQueuesOutput), nil
	}
	out, err := c.ConnectAPI.ListQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQueuesPages(input *connect.ListQueuesInput, fn func(*connect.ListQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQueuesOutput{}
	fnCacher := func(out *connect.ListQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQueuesPagesWithContext(ctx context.Context, input *connect.ListQueuesInput, fn func(*connect.ListQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQueuesOutput{}
	fnCacher := func(out *connect.ListQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQueuesWithContext(ctx context.Context, input *connect.ListQueuesInput, opts ...request.Option) (*connect.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQueuesOutput), nil
	}
	out, err := c.ConnectAPI.ListQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQuickConnects(input *connect.ListQuickConnectsInput) (*connect.ListQuickConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQuickConnectsOutput), nil
	}
	out, err := c.ConnectAPI.ListQuickConnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListQuickConnectsPages(input *connect.ListQuickConnectsInput, fn func(*connect.ListQuickConnectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQuickConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQuickConnectsOutput{}
	fnCacher := func(out *connect.ListQuickConnectsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQuickConnectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQuickConnectsPagesWithContext(ctx context.Context, input *connect.ListQuickConnectsInput, fn func(*connect.ListQuickConnectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListQuickConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListQuickConnectsOutput{}
	fnCacher := func(out *connect.ListQuickConnectsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListQuickConnectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListQuickConnectsWithContext(ctx context.Context, input *connect.ListQuickConnectsInput, opts ...request.Option) (*connect.ListQuickConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListQuickConnectsOutput), nil
	}
	out, err := c.ConnectAPI.ListQuickConnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRoutingProfileQueues(input *connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRoutingProfileQueuesOutput), nil
	}
	out, err := c.ConnectAPI.ListRoutingProfileQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRoutingProfileQueuesPages(input *connect.ListRoutingProfileQueuesInput, fn func(*connect.ListRoutingProfileQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRoutingProfileQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRoutingProfileQueuesOutput{}
	fnCacher := func(out *connect.ListRoutingProfileQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRoutingProfileQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRoutingProfileQueuesPagesWithContext(ctx context.Context, input *connect.ListRoutingProfileQueuesInput, fn func(*connect.ListRoutingProfileQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRoutingProfileQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRoutingProfileQueuesOutput{}
	fnCacher := func(out *connect.ListRoutingProfileQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRoutingProfileQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRoutingProfileQueuesWithContext(ctx context.Context, input *connect.ListRoutingProfileQueuesInput, opts ...request.Option) (*connect.ListRoutingProfileQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRoutingProfileQueuesOutput), nil
	}
	out, err := c.ConnectAPI.ListRoutingProfileQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRoutingProfiles(input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRoutingProfilesOutput), nil
	}
	out, err := c.ConnectAPI.ListRoutingProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRoutingProfilesPages(input *connect.ListRoutingProfilesInput, fn func(*connect.ListRoutingProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRoutingProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRoutingProfilesOutput{}
	fnCacher := func(out *connect.ListRoutingProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRoutingProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRoutingProfilesPagesWithContext(ctx context.Context, input *connect.ListRoutingProfilesInput, fn func(*connect.ListRoutingProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRoutingProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRoutingProfilesOutput{}
	fnCacher := func(out *connect.ListRoutingProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRoutingProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRoutingProfilesWithContext(ctx context.Context, input *connect.ListRoutingProfilesInput, opts ...request.Option) (*connect.ListRoutingProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRoutingProfilesOutput), nil
	}
	out, err := c.ConnectAPI.ListRoutingProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRules(input *connect.ListRulesInput) (*connect.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRulesOutput), nil
	}
	out, err := c.ConnectAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListRulesPages(input *connect.ListRulesInput, fn func(*connect.ListRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRulesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRulesOutput{}
	fnCacher := func(out *connect.ListRulesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRulesPagesWithContext(ctx context.Context, input *connect.ListRulesInput, fn func(*connect.ListRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListRulesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListRulesOutput{}
	fnCacher := func(out *connect.ListRulesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListRulesWithContext(ctx context.Context, input *connect.ListRulesInput, opts ...request.Option) (*connect.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListRulesOutput), nil
	}
	out, err := c.ConnectAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityKeys(input *connect.ListSecurityKeysInput) (*connect.ListSecurityKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityKeysOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityKeysPages(input *connect.ListSecurityKeysInput, fn func(*connect.ListSecurityKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityKeysOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityKeysOutput{}
	fnCacher := func(out *connect.ListSecurityKeysOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityKeysPagesWithContext(ctx context.Context, input *connect.ListSecurityKeysInput, fn func(*connect.ListSecurityKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityKeysOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityKeysOutput{}
	fnCacher := func(out *connect.ListSecurityKeysOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityKeysWithContext(ctx context.Context, input *connect.ListSecurityKeysInput, opts ...request.Option) (*connect.ListSecurityKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityKeysOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityProfilePermissions(input *connect.ListSecurityProfilePermissionsInput) (*connect.ListSecurityProfilePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityProfilePermissionsOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityProfilePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityProfilePermissionsPages(input *connect.ListSecurityProfilePermissionsInput, fn func(*connect.ListSecurityProfilePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityProfilePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityProfilePermissionsOutput{}
	fnCacher := func(out *connect.ListSecurityProfilePermissionsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityProfilePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityProfilePermissionsPagesWithContext(ctx context.Context, input *connect.ListSecurityProfilePermissionsInput, fn func(*connect.ListSecurityProfilePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityProfilePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityProfilePermissionsOutput{}
	fnCacher := func(out *connect.ListSecurityProfilePermissionsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityProfilePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityProfilePermissionsWithContext(ctx context.Context, input *connect.ListSecurityProfilePermissionsInput, opts ...request.Option) (*connect.ListSecurityProfilePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityProfilePermissionsOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityProfilePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityProfiles(input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityProfilesOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListSecurityProfilesPages(input *connect.ListSecurityProfilesInput, fn func(*connect.ListSecurityProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityProfilesOutput{}
	fnCacher := func(out *connect.ListSecurityProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityProfilesPagesWithContext(ctx context.Context, input *connect.ListSecurityProfilesInput, fn func(*connect.ListSecurityProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListSecurityProfilesOutput{}
	fnCacher := func(out *connect.ListSecurityProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListSecurityProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListSecurityProfilesWithContext(ctx context.Context, input *connect.ListSecurityProfilesInput, opts ...request.Option) (*connect.ListSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListSecurityProfilesOutput), nil
	}
	out, err := c.ConnectAPI.ListSecurityProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTagsForResource(input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTagsForResourceWithContext(ctx context.Context, input *connect.ListTagsForResourceInput, opts ...request.Option) (*connect.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTaskTemplates(input *connect.ListTaskTemplatesInput) (*connect.ListTaskTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTaskTemplatesOutput), nil
	}
	out, err := c.ConnectAPI.ListTaskTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTaskTemplatesPages(input *connect.ListTaskTemplatesInput, fn func(*connect.ListTaskTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListTaskTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListTaskTemplatesOutput{}
	fnCacher := func(out *connect.ListTaskTemplatesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListTaskTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListTaskTemplatesPagesWithContext(ctx context.Context, input *connect.ListTaskTemplatesInput, fn func(*connect.ListTaskTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListTaskTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListTaskTemplatesOutput{}
	fnCacher := func(out *connect.ListTaskTemplatesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListTaskTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListTaskTemplatesWithContext(ctx context.Context, input *connect.ListTaskTemplatesInput, opts ...request.Option) (*connect.ListTaskTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTaskTemplatesOutput), nil
	}
	out, err := c.ConnectAPI.ListTaskTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTrafficDistributionGroups(input *connect.ListTrafficDistributionGroupsInput) (*connect.ListTrafficDistributionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTrafficDistributionGroupsOutput), nil
	}
	out, err := c.ConnectAPI.ListTrafficDistributionGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListTrafficDistributionGroupsPages(input *connect.ListTrafficDistributionGroupsInput, fn func(*connect.ListTrafficDistributionGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListTrafficDistributionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListTrafficDistributionGroupsOutput{}
	fnCacher := func(out *connect.ListTrafficDistributionGroupsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListTrafficDistributionGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListTrafficDistributionGroupsPagesWithContext(ctx context.Context, input *connect.ListTrafficDistributionGroupsInput, fn func(*connect.ListTrafficDistributionGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListTrafficDistributionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListTrafficDistributionGroupsOutput{}
	fnCacher := func(out *connect.ListTrafficDistributionGroupsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListTrafficDistributionGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListTrafficDistributionGroupsWithContext(ctx context.Context, input *connect.ListTrafficDistributionGroupsInput, opts ...request.Option) (*connect.ListTrafficDistributionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListTrafficDistributionGroupsOutput), nil
	}
	out, err := c.ConnectAPI.ListTrafficDistributionGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUseCases(input *connect.ListUseCasesInput) (*connect.ListUseCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUseCasesOutput), nil
	}
	out, err := c.ConnectAPI.ListUseCases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUseCasesPages(input *connect.ListUseCasesInput, fn func(*connect.ListUseCasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUseCasesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUseCasesOutput{}
	fnCacher := func(out *connect.ListUseCasesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUseCasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUseCasesPagesWithContext(ctx context.Context, input *connect.ListUseCasesInput, fn func(*connect.ListUseCasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUseCasesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUseCasesOutput{}
	fnCacher := func(out *connect.ListUseCasesOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUseCasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUseCasesWithContext(ctx context.Context, input *connect.ListUseCasesInput, opts ...request.Option) (*connect.ListUseCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUseCasesOutput), nil
	}
	out, err := c.ConnectAPI.ListUseCasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUserHierarchyGroups(input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUserHierarchyGroupsOutput), nil
	}
	out, err := c.ConnectAPI.ListUserHierarchyGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUserHierarchyGroupsPages(input *connect.ListUserHierarchyGroupsInput, fn func(*connect.ListUserHierarchyGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUserHierarchyGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUserHierarchyGroupsOutput{}
	fnCacher := func(out *connect.ListUserHierarchyGroupsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUserHierarchyGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUserHierarchyGroupsPagesWithContext(ctx context.Context, input *connect.ListUserHierarchyGroupsInput, fn func(*connect.ListUserHierarchyGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUserHierarchyGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUserHierarchyGroupsOutput{}
	fnCacher := func(out *connect.ListUserHierarchyGroupsOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUserHierarchyGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUserHierarchyGroupsWithContext(ctx context.Context, input *connect.ListUserHierarchyGroupsInput, opts ...request.Option) (*connect.ListUserHierarchyGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUserHierarchyGroupsOutput), nil
	}
	out, err := c.ConnectAPI.ListUserHierarchyGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUsers(input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUsersOutput), nil
	}
	out, err := c.ConnectAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) ListUsersPages(input *connect.ListUsersInput, fn func(*connect.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUsersOutput{}
	fnCacher := func(out *connect.ListUsersOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUsersPagesWithContext(ctx context.Context, input *connect.ListUsersInput, fn func(*connect.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.ListUsersOutput{}
	fnCacher := func(out *connect.ListUsersOutput, more bool) bool {
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
	if err := c.ConnectAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) ListUsersWithContext(ctx context.Context, input *connect.ListUsersInput, opts ...request.Option) (*connect.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.ListUsersOutput), nil
	}
	out, err := c.ConnectAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchAvailablePhoneNumbers(input *connect.SearchAvailablePhoneNumbersInput) (*connect.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ConnectAPI.SearchAvailablePhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchAvailablePhoneNumbersPages(input *connect.SearchAvailablePhoneNumbersInput, fn func(*connect.SearchAvailablePhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *connect.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchAvailablePhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchAvailablePhoneNumbersPagesWithContext(ctx context.Context, input *connect.SearchAvailablePhoneNumbersInput, fn func(*connect.SearchAvailablePhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *connect.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchAvailablePhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchAvailablePhoneNumbersWithContext(ctx context.Context, input *connect.SearchAvailablePhoneNumbersInput, opts ...request.Option) (*connect.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ConnectAPI.SearchAvailablePhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchQueues(input *connect.SearchQueuesInput) (*connect.SearchQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchQueuesOutput), nil
	}
	out, err := c.ConnectAPI.SearchQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchQueuesPages(input *connect.SearchQueuesInput, fn func(*connect.SearchQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchQueuesOutput{}
	fnCacher := func(out *connect.SearchQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchQueuesPagesWithContext(ctx context.Context, input *connect.SearchQueuesInput, fn func(*connect.SearchQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchQueuesOutput{}
	fnCacher := func(out *connect.SearchQueuesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchQueuesWithContext(ctx context.Context, input *connect.SearchQueuesInput, opts ...request.Option) (*connect.SearchQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchQueuesOutput), nil
	}
	out, err := c.ConnectAPI.SearchQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchRoutingProfiles(input *connect.SearchRoutingProfilesInput) (*connect.SearchRoutingProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchRoutingProfilesOutput), nil
	}
	out, err := c.ConnectAPI.SearchRoutingProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchRoutingProfilesPages(input *connect.SearchRoutingProfilesInput, fn func(*connect.SearchRoutingProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchRoutingProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchRoutingProfilesOutput{}
	fnCacher := func(out *connect.SearchRoutingProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchRoutingProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchRoutingProfilesPagesWithContext(ctx context.Context, input *connect.SearchRoutingProfilesInput, fn func(*connect.SearchRoutingProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchRoutingProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchRoutingProfilesOutput{}
	fnCacher := func(out *connect.SearchRoutingProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchRoutingProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchRoutingProfilesWithContext(ctx context.Context, input *connect.SearchRoutingProfilesInput, opts ...request.Option) (*connect.SearchRoutingProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchRoutingProfilesOutput), nil
	}
	out, err := c.ConnectAPI.SearchRoutingProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchSecurityProfiles(input *connect.SearchSecurityProfilesInput) (*connect.SearchSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchSecurityProfilesOutput), nil
	}
	out, err := c.ConnectAPI.SearchSecurityProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchSecurityProfilesPages(input *connect.SearchSecurityProfilesInput, fn func(*connect.SearchSecurityProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchSecurityProfilesOutput{}
	fnCacher := func(out *connect.SearchSecurityProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchSecurityProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchSecurityProfilesPagesWithContext(ctx context.Context, input *connect.SearchSecurityProfilesInput, fn func(*connect.SearchSecurityProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchSecurityProfilesOutput{}
	fnCacher := func(out *connect.SearchSecurityProfilesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchSecurityProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchSecurityProfilesWithContext(ctx context.Context, input *connect.SearchSecurityProfilesInput, opts ...request.Option) (*connect.SearchSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchSecurityProfilesOutput), nil
	}
	out, err := c.ConnectAPI.SearchSecurityProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchUsers(input *connect.SearchUsersInput) (*connect.SearchUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchUsersOutput), nil
	}
	out, err := c.ConnectAPI.SearchUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchUsersPages(input *connect.SearchUsersInput, fn func(*connect.SearchUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchUsersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchUsersOutput{}
	fnCacher := func(out *connect.SearchUsersOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchUsersPagesWithContext(ctx context.Context, input *connect.SearchUsersInput, fn func(*connect.SearchUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchUsersOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchUsersOutput{}
	fnCacher := func(out *connect.SearchUsersOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchUsersWithContext(ctx context.Context, input *connect.SearchUsersInput, opts ...request.Option) (*connect.SearchUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchUsersOutput), nil
	}
	out, err := c.ConnectAPI.SearchUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchVocabularies(input *connect.SearchVocabulariesInput) (*connect.SearchVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchVocabulariesOutput), nil
	}
	out, err := c.ConnectAPI.SearchVocabularies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Connect) SearchVocabulariesPages(input *connect.SearchVocabulariesInput, fn func(*connect.SearchVocabulariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchVocabulariesOutput{}
	fnCacher := func(out *connect.SearchVocabulariesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchVocabulariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchVocabulariesPagesWithContext(ctx context.Context, input *connect.SearchVocabulariesInput, fn func(*connect.SearchVocabulariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connect.SearchVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &connect.SearchVocabulariesOutput{}
	fnCacher := func(out *connect.SearchVocabulariesOutput, more bool) bool {
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
	if err := c.ConnectAPI.SearchVocabulariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Connect) SearchVocabulariesWithContext(ctx context.Context, input *connect.SearchVocabulariesInput, opts ...request.Option) (*connect.SearchVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connect.SearchVocabulariesOutput), nil
	}
	out, err := c.ConnectAPI.SearchVocabulariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

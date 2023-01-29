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
package cloudwatcheventscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudWatchEvents struct {
	cloudwatcheventsiface.CloudWatchEventsAPI
	cache *cache.Cache
}

func New(cloudwatcheventsapi cloudwatcheventsiface.CloudWatchEventsAPI) *CloudWatchEvents {
	return &CloudWatchEvents{
		CloudWatchEventsAPI: cloudwatcheventsapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudWatchEvents) DescribeApiDestination(input *cloudwatchevents.DescribeApiDestinationInput) (*cloudwatchevents.DescribeApiDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeApiDestinationOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeApiDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeApiDestinationWithContext(ctx context.Context, input *cloudwatchevents.DescribeApiDestinationInput, opts ...request.Option) (*cloudwatchevents.DescribeApiDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeApiDestinationOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeApiDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeArchive(input *cloudwatchevents.DescribeArchiveInput) (*cloudwatchevents.DescribeArchiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeArchiveOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeArchive(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeArchiveWithContext(ctx context.Context, input *cloudwatchevents.DescribeArchiveInput, opts ...request.Option) (*cloudwatchevents.DescribeArchiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeArchiveOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeArchiveWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeConnection(input *cloudwatchevents.DescribeConnectionInput) (*cloudwatchevents.DescribeConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeConnectionOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeConnectionWithContext(ctx context.Context, input *cloudwatchevents.DescribeConnectionInput, opts ...request.Option) (*cloudwatchevents.DescribeConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeConnectionOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeEventBus(input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeEventBusOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeEventBus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeEventBusWithContext(ctx context.Context, input *cloudwatchevents.DescribeEventBusInput, opts ...request.Option) (*cloudwatchevents.DescribeEventBusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeEventBusOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeEventBusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeEventSource(input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeEventSourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeEventSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeEventSourceWithContext(ctx context.Context, input *cloudwatchevents.DescribeEventSourceInput, opts ...request.Option) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeEventSourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeEventSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribePartnerEventSource(input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribePartnerEventSourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribePartnerEventSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribePartnerEventSourceWithContext(ctx context.Context, input *cloudwatchevents.DescribePartnerEventSourceInput, opts ...request.Option) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribePartnerEventSourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribePartnerEventSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeReplay(input *cloudwatchevents.DescribeReplayInput) (*cloudwatchevents.DescribeReplayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeReplayOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeReplay(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeReplayWithContext(ctx context.Context, input *cloudwatchevents.DescribeReplayInput, opts ...request.Option) (*cloudwatchevents.DescribeReplayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeReplayOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeReplayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeRule(input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeRuleOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) DescribeRuleWithContext(ctx context.Context, input *cloudwatchevents.DescribeRuleInput, opts ...request.Option) (*cloudwatchevents.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.DescribeRuleOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.DescribeRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListApiDestinations(input *cloudwatchevents.ListApiDestinationsInput) (*cloudwatchevents.ListApiDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListApiDestinationsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListApiDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListApiDestinationsWithContext(ctx context.Context, input *cloudwatchevents.ListApiDestinationsInput, opts ...request.Option) (*cloudwatchevents.ListApiDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListApiDestinationsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListApiDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListArchives(input *cloudwatchevents.ListArchivesInput) (*cloudwatchevents.ListArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListArchivesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListArchives(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListArchivesWithContext(ctx context.Context, input *cloudwatchevents.ListArchivesInput, opts ...request.Option) (*cloudwatchevents.ListArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListArchivesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListArchivesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListConnections(input *cloudwatchevents.ListConnectionsInput) (*cloudwatchevents.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListConnectionsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListConnectionsWithContext(ctx context.Context, input *cloudwatchevents.ListConnectionsInput, opts ...request.Option) (*cloudwatchevents.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListConnectionsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListEventBuses(input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListEventBusesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListEventBuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListEventBusesWithContext(ctx context.Context, input *cloudwatchevents.ListEventBusesInput, opts ...request.Option) (*cloudwatchevents.ListEventBusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListEventBusesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListEventBusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListEventSources(input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListEventSourcesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListEventSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListEventSourcesWithContext(ctx context.Context, input *cloudwatchevents.ListEventSourcesInput, opts ...request.Option) (*cloudwatchevents.ListEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListEventSourcesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListEventSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListPartnerEventSourceAccounts(input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListPartnerEventSourceAccountsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListPartnerEventSourceAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListPartnerEventSourceAccountsWithContext(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput, opts ...request.Option) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListPartnerEventSourceAccountsOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListPartnerEventSourceAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListPartnerEventSources(input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListPartnerEventSourcesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListPartnerEventSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListPartnerEventSourcesWithContext(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourcesInput, opts ...request.Option) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListPartnerEventSourcesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListPartnerEventSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListReplays(input *cloudwatchevents.ListReplaysInput) (*cloudwatchevents.ListReplaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListReplaysOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListReplays(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListReplaysWithContext(ctx context.Context, input *cloudwatchevents.ListReplaysInput, opts ...request.Option) (*cloudwatchevents.ListReplaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListReplaysOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListReplaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListRuleNamesByTarget(input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListRuleNamesByTargetOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListRuleNamesByTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListRuleNamesByTargetWithContext(ctx context.Context, input *cloudwatchevents.ListRuleNamesByTargetInput, opts ...request.Option) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListRuleNamesByTargetOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListRuleNamesByTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListRules(input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListRulesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListRulesWithContext(ctx context.Context, input *cloudwatchevents.ListRulesInput, opts ...request.Option) (*cloudwatchevents.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListRulesOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListTagsForResource(input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatchevents.ListTagsForResourceInput, opts ...request.Option) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListTargetsByRule(input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListTargetsByRuleOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListTargetsByRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvents) ListTargetsByRuleWithContext(ctx context.Context, input *cloudwatchevents.ListTargetsByRuleInput, opts ...request.Option) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevents.ListTargetsByRuleOutput), nil
	}
	out, err := c.CloudWatchEventsAPI.ListTargetsByRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

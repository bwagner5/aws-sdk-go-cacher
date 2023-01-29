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
package eventbridgecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EventBridge struct {
	eventbridgeiface.EventBridgeAPI
	cache *cache.Cache
}

func New(eventbridgeapi eventbridgeiface.EventBridgeAPI) *EventBridge {
	return &EventBridge{
		EventBridgeAPI: eventbridgeapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EventBridge) DescribeApiDestination(input *eventbridge.DescribeApiDestinationInput) (*eventbridge.DescribeApiDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeApiDestinationOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeApiDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeApiDestinationWithContext(ctx context.Context, input *eventbridge.DescribeApiDestinationInput, opts ...request.Option) (*eventbridge.DescribeApiDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeApiDestinationOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeApiDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeArchive(input *eventbridge.DescribeArchiveInput) (*eventbridge.DescribeArchiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeArchiveOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeArchive(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeArchiveWithContext(ctx context.Context, input *eventbridge.DescribeArchiveInput, opts ...request.Option) (*eventbridge.DescribeArchiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeArchiveOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeArchiveWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeConnection(input *eventbridge.DescribeConnectionInput) (*eventbridge.DescribeConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeConnectionOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeConnectionWithContext(ctx context.Context, input *eventbridge.DescribeConnectionInput, opts ...request.Option) (*eventbridge.DescribeConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeConnectionOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEndpoint(input *eventbridge.DescribeEndpointInput) (*eventbridge.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEndpointOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEndpointWithContext(ctx context.Context, input *eventbridge.DescribeEndpointInput, opts ...request.Option) (*eventbridge.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEndpointOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEventBus(input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEventBusOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEventBus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEventBusWithContext(ctx context.Context, input *eventbridge.DescribeEventBusInput, opts ...request.Option) (*eventbridge.DescribeEventBusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEventBusOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEventBusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEventSource(input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEventSourceOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEventSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeEventSourceWithContext(ctx context.Context, input *eventbridge.DescribeEventSourceInput, opts ...request.Option) (*eventbridge.DescribeEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeEventSourceOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeEventSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribePartnerEventSource(input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribePartnerEventSourceOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribePartnerEventSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.DescribePartnerEventSourceInput, opts ...request.Option) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribePartnerEventSourceOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribePartnerEventSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeReplay(input *eventbridge.DescribeReplayInput) (*eventbridge.DescribeReplayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeReplayOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeReplay(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeReplayWithContext(ctx context.Context, input *eventbridge.DescribeReplayInput, opts ...request.Option) (*eventbridge.DescribeReplayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeReplayOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeReplayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeRule(input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeRuleOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) DescribeRuleWithContext(ctx context.Context, input *eventbridge.DescribeRuleInput, opts ...request.Option) (*eventbridge.DescribeRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.DescribeRuleOutput), nil
	}
	out, err := c.EventBridgeAPI.DescribeRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListApiDestinations(input *eventbridge.ListApiDestinationsInput) (*eventbridge.ListApiDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListApiDestinationsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListApiDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListApiDestinationsWithContext(ctx context.Context, input *eventbridge.ListApiDestinationsInput, opts ...request.Option) (*eventbridge.ListApiDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListApiDestinationsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListApiDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListArchives(input *eventbridge.ListArchivesInput) (*eventbridge.ListArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListArchivesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListArchives(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListArchivesWithContext(ctx context.Context, input *eventbridge.ListArchivesInput, opts ...request.Option) (*eventbridge.ListArchivesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListArchivesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListArchivesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListConnections(input *eventbridge.ListConnectionsInput) (*eventbridge.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListConnectionsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListConnectionsWithContext(ctx context.Context, input *eventbridge.ListConnectionsInput, opts ...request.Option) (*eventbridge.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListConnectionsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEndpoints(input *eventbridge.ListEndpointsInput) (*eventbridge.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEndpointsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEndpointsWithContext(ctx context.Context, input *eventbridge.ListEndpointsInput, opts ...request.Option) (*eventbridge.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEndpointsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEventBuses(input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEventBusesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEventBuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEventBusesWithContext(ctx context.Context, input *eventbridge.ListEventBusesInput, opts ...request.Option) (*eventbridge.ListEventBusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEventBusesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEventBusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEventSources(input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEventSourcesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEventSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListEventSourcesWithContext(ctx context.Context, input *eventbridge.ListEventSourcesInput, opts ...request.Option) (*eventbridge.ListEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListEventSourcesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListEventSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListPartnerEventSourceAccounts(input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListPartnerEventSourceAccountsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListPartnerEventSourceAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListPartnerEventSourceAccountsWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourceAccountsInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListPartnerEventSourceAccountsOutput), nil
	}
	out, err := c.EventBridgeAPI.ListPartnerEventSourceAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListPartnerEventSources(input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListPartnerEventSourcesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListPartnerEventSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListPartnerEventSourcesWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourcesInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListPartnerEventSourcesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListPartnerEventSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListReplays(input *eventbridge.ListReplaysInput) (*eventbridge.ListReplaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListReplaysOutput), nil
	}
	out, err := c.EventBridgeAPI.ListReplays(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListReplaysWithContext(ctx context.Context, input *eventbridge.ListReplaysInput, opts ...request.Option) (*eventbridge.ListReplaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListReplaysOutput), nil
	}
	out, err := c.EventBridgeAPI.ListReplaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListRuleNamesByTarget(input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListRuleNamesByTargetOutput), nil
	}
	out, err := c.EventBridgeAPI.ListRuleNamesByTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListRuleNamesByTargetWithContext(ctx context.Context, input *eventbridge.ListRuleNamesByTargetInput, opts ...request.Option) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListRuleNamesByTargetOutput), nil
	}
	out, err := c.EventBridgeAPI.ListRuleNamesByTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListRules(input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListRulesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListRulesWithContext(ctx context.Context, input *eventbridge.ListRulesInput, opts ...request.Option) (*eventbridge.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListRulesOutput), nil
	}
	out, err := c.EventBridgeAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListTagsForResource(input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListTagsForResourceOutput), nil
	}
	out, err := c.EventBridgeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListTagsForResourceWithContext(ctx context.Context, input *eventbridge.ListTagsForResourceInput, opts ...request.Option) (*eventbridge.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListTagsForResourceOutput), nil
	}
	out, err := c.EventBridgeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListTargetsByRule(input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListTargetsByRuleOutput), nil
	}
	out, err := c.EventBridgeAPI.ListTargetsByRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EventBridge) ListTargetsByRuleWithContext(ctx context.Context, input *eventbridge.ListTargetsByRuleInput, opts ...request.Option) (*eventbridge.ListTargetsByRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eventbridge.ListTargetsByRuleOutput), nil
	}
	out, err := c.EventBridgeAPI.ListTargetsByRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

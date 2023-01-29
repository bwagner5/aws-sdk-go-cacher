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
package networkmanagercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/networkmanager/networkmanageriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type NetworkManager struct {
	networkmanageriface.NetworkManagerAPI
	cache *cache.Cache
}

func New(networkmanagerapi networkmanageriface.NetworkManagerAPI) *NetworkManager {
	return &NetworkManager{
		NetworkManagerAPI: networkmanagerapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *NetworkManager) DescribeGlobalNetworks(input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.DescribeGlobalNetworksOutput), nil
	}
	out, err := c.NetworkManagerAPI.DescribeGlobalNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) DescribeGlobalNetworksPages(input *networkmanager.DescribeGlobalNetworksInput, fn func(*networkmanager.DescribeGlobalNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.DescribeGlobalNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.DescribeGlobalNetworksOutput{}
	fnCacher := func(out *networkmanager.DescribeGlobalNetworksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.DescribeGlobalNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) DescribeGlobalNetworksPagesWithContext(ctx context.Context, input *networkmanager.DescribeGlobalNetworksInput, fn func(*networkmanager.DescribeGlobalNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.DescribeGlobalNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.DescribeGlobalNetworksOutput{}
	fnCacher := func(out *networkmanager.DescribeGlobalNetworksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.DescribeGlobalNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) DescribeGlobalNetworksWithContext(ctx context.Context, input *networkmanager.DescribeGlobalNetworksInput, opts ...request.Option) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.DescribeGlobalNetworksOutput), nil
	}
	out, err := c.NetworkManagerAPI.DescribeGlobalNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectAttachment(input *networkmanager.GetConnectAttachmentInput) (*networkmanager.GetConnectAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectAttachmentWithContext(ctx context.Context, input *networkmanager.GetConnectAttachmentInput, opts ...request.Option) (*networkmanager.GetConnectAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectPeer(input *networkmanager.GetConnectPeerInput) (*networkmanager.GetConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectPeerOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectPeer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectPeerAssociations(input *networkmanager.GetConnectPeerAssociationsInput) (*networkmanager.GetConnectPeerAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectPeerAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectPeerAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectPeerAssociationsPages(input *networkmanager.GetConnectPeerAssociationsInput, fn func(*networkmanager.GetConnectPeerAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetConnectPeerAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetConnectPeerAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetConnectPeerAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetConnectPeerAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetConnectPeerAssociationsPagesWithContext(ctx context.Context, input *networkmanager.GetConnectPeerAssociationsInput, fn func(*networkmanager.GetConnectPeerAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetConnectPeerAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetConnectPeerAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetConnectPeerAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetConnectPeerAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetConnectPeerAssociationsWithContext(ctx context.Context, input *networkmanager.GetConnectPeerAssociationsInput, opts ...request.Option) (*networkmanager.GetConnectPeerAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectPeerAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectPeerAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectPeerWithContext(ctx context.Context, input *networkmanager.GetConnectPeerInput, opts ...request.Option) (*networkmanager.GetConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectPeerOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectPeerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnections(input *networkmanager.GetConnectionsInput) (*networkmanager.GetConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectionsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetConnectionsPages(input *networkmanager.GetConnectionsInput, fn func(*networkmanager.GetConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetConnectionsOutput{}
	fnCacher := func(out *networkmanager.GetConnectionsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetConnectionsPagesWithContext(ctx context.Context, input *networkmanager.GetConnectionsInput, fn func(*networkmanager.GetConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetConnectionsOutput{}
	fnCacher := func(out *networkmanager.GetConnectionsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetConnectionsWithContext(ctx context.Context, input *networkmanager.GetConnectionsInput, opts ...request.Option) (*networkmanager.GetConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetConnectionsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetwork(input *networkmanager.GetCoreNetworkInput) (*networkmanager.GetCoreNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkChangeEvents(input *networkmanager.GetCoreNetworkChangeEventsInput) (*networkmanager.GetCoreNetworkChangeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkChangeEventsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkChangeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkChangeEventsPages(input *networkmanager.GetCoreNetworkChangeEventsInput, fn func(*networkmanager.GetCoreNetworkChangeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCoreNetworkChangeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCoreNetworkChangeEventsOutput{}
	fnCacher := func(out *networkmanager.GetCoreNetworkChangeEventsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCoreNetworkChangeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCoreNetworkChangeEventsPagesWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkChangeEventsInput, fn func(*networkmanager.GetCoreNetworkChangeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCoreNetworkChangeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCoreNetworkChangeEventsOutput{}
	fnCacher := func(out *networkmanager.GetCoreNetworkChangeEventsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCoreNetworkChangeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCoreNetworkChangeEventsWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkChangeEventsInput, opts ...request.Option) (*networkmanager.GetCoreNetworkChangeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkChangeEventsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkChangeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkChangeSet(input *networkmanager.GetCoreNetworkChangeSetInput) (*networkmanager.GetCoreNetworkChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkChangeSetOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkChangeSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkChangeSetPages(input *networkmanager.GetCoreNetworkChangeSetInput, fn func(*networkmanager.GetCoreNetworkChangeSetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCoreNetworkChangeSetOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCoreNetworkChangeSetOutput{}
	fnCacher := func(out *networkmanager.GetCoreNetworkChangeSetOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCoreNetworkChangeSetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCoreNetworkChangeSetPagesWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkChangeSetInput, fn func(*networkmanager.GetCoreNetworkChangeSetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCoreNetworkChangeSetOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCoreNetworkChangeSetOutput{}
	fnCacher := func(out *networkmanager.GetCoreNetworkChangeSetOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCoreNetworkChangeSetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCoreNetworkChangeSetWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkChangeSetInput, opts ...request.Option) (*networkmanager.GetCoreNetworkChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkChangeSetOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkChangeSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkPolicy(input *networkmanager.GetCoreNetworkPolicyInput) (*networkmanager.GetCoreNetworkPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkPolicyOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkPolicyWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkPolicyInput, opts ...request.Option) (*networkmanager.GetCoreNetworkPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkPolicyOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCoreNetworkWithContext(ctx context.Context, input *networkmanager.GetCoreNetworkInput, opts ...request.Option) (*networkmanager.GetCoreNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCoreNetworkOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCoreNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCustomerGatewayAssociations(input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCustomerGatewayAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCustomerGatewayAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetCustomerGatewayAssociationsPages(input *networkmanager.GetCustomerGatewayAssociationsInput, fn func(*networkmanager.GetCustomerGatewayAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCustomerGatewayAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCustomerGatewayAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetCustomerGatewayAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCustomerGatewayAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCustomerGatewayAssociationsPagesWithContext(ctx context.Context, input *networkmanager.GetCustomerGatewayAssociationsInput, fn func(*networkmanager.GetCustomerGatewayAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetCustomerGatewayAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetCustomerGatewayAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetCustomerGatewayAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetCustomerGatewayAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetCustomerGatewayAssociationsWithContext(ctx context.Context, input *networkmanager.GetCustomerGatewayAssociationsInput, opts ...request.Option) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetCustomerGatewayAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetCustomerGatewayAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetDevices(input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetDevicesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetDevicesPages(input *networkmanager.GetDevicesInput, fn func(*networkmanager.GetDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetDevicesOutput{}
	fnCacher := func(out *networkmanager.GetDevicesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetDevicesPagesWithContext(ctx context.Context, input *networkmanager.GetDevicesInput, fn func(*networkmanager.GetDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetDevicesOutput{}
	fnCacher := func(out *networkmanager.GetDevicesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetDevicesWithContext(ctx context.Context, input *networkmanager.GetDevicesInput, opts ...request.Option) (*networkmanager.GetDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetDevicesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetLinkAssociations(input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetLinkAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetLinkAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetLinkAssociationsPages(input *networkmanager.GetLinkAssociationsInput, fn func(*networkmanager.GetLinkAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetLinkAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetLinkAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetLinkAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetLinkAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetLinkAssociationsPagesWithContext(ctx context.Context, input *networkmanager.GetLinkAssociationsInput, fn func(*networkmanager.GetLinkAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetLinkAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetLinkAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetLinkAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetLinkAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetLinkAssociationsWithContext(ctx context.Context, input *networkmanager.GetLinkAssociationsInput, opts ...request.Option) (*networkmanager.GetLinkAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetLinkAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetLinkAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetLinks(input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetLinksOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetLinks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetLinksPages(input *networkmanager.GetLinksInput, fn func(*networkmanager.GetLinksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetLinksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetLinksOutput{}
	fnCacher := func(out *networkmanager.GetLinksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetLinksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetLinksPagesWithContext(ctx context.Context, input *networkmanager.GetLinksInput, fn func(*networkmanager.GetLinksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetLinksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetLinksOutput{}
	fnCacher := func(out *networkmanager.GetLinksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetLinksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetLinksWithContext(ctx context.Context, input *networkmanager.GetLinksInput, opts ...request.Option) (*networkmanager.GetLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetLinksOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetLinksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResourceCounts(input *networkmanager.GetNetworkResourceCountsInput) (*networkmanager.GetNetworkResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourceCountsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResourceCounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResourceCountsPages(input *networkmanager.GetNetworkResourceCountsInput, fn func(*networkmanager.GetNetworkResourceCountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourceCountsOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourceCountsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourceCountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourceCountsPagesWithContext(ctx context.Context, input *networkmanager.GetNetworkResourceCountsInput, fn func(*networkmanager.GetNetworkResourceCountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourceCountsOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourceCountsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourceCountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourceCountsWithContext(ctx context.Context, input *networkmanager.GetNetworkResourceCountsInput, opts ...request.Option) (*networkmanager.GetNetworkResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourceCountsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResourceCountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResourceRelationships(input *networkmanager.GetNetworkResourceRelationshipsInput) (*networkmanager.GetNetworkResourceRelationshipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourceRelationshipsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResourceRelationships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResourceRelationshipsPages(input *networkmanager.GetNetworkResourceRelationshipsInput, fn func(*networkmanager.GetNetworkResourceRelationshipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourceRelationshipsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourceRelationshipsOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourceRelationshipsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourceRelationshipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourceRelationshipsPagesWithContext(ctx context.Context, input *networkmanager.GetNetworkResourceRelationshipsInput, fn func(*networkmanager.GetNetworkResourceRelationshipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourceRelationshipsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourceRelationshipsOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourceRelationshipsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourceRelationshipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourceRelationshipsWithContext(ctx context.Context, input *networkmanager.GetNetworkResourceRelationshipsInput, opts ...request.Option) (*networkmanager.GetNetworkResourceRelationshipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourceRelationshipsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResourceRelationshipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResources(input *networkmanager.GetNetworkResourcesInput) (*networkmanager.GetNetworkResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourcesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkResourcesPages(input *networkmanager.GetNetworkResourcesInput, fn func(*networkmanager.GetNetworkResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourcesOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourcesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourcesPagesWithContext(ctx context.Context, input *networkmanager.GetNetworkResourcesInput, fn func(*networkmanager.GetNetworkResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkResourcesOutput{}
	fnCacher := func(out *networkmanager.GetNetworkResourcesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkResourcesWithContext(ctx context.Context, input *networkmanager.GetNetworkResourcesInput, opts ...request.Option) (*networkmanager.GetNetworkResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkResourcesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkRoutes(input *networkmanager.GetNetworkRoutesInput) (*networkmanager.GetNetworkRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkRoutesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkRoutesWithContext(ctx context.Context, input *networkmanager.GetNetworkRoutesInput, opts ...request.Option) (*networkmanager.GetNetworkRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkRoutesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkTelemetry(input *networkmanager.GetNetworkTelemetryInput) (*networkmanager.GetNetworkTelemetryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkTelemetryOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkTelemetry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetNetworkTelemetryPages(input *networkmanager.GetNetworkTelemetryInput, fn func(*networkmanager.GetNetworkTelemetryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkTelemetryOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkTelemetryOutput{}
	fnCacher := func(out *networkmanager.GetNetworkTelemetryOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkTelemetryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkTelemetryPagesWithContext(ctx context.Context, input *networkmanager.GetNetworkTelemetryInput, fn func(*networkmanager.GetNetworkTelemetryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetNetworkTelemetryOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetNetworkTelemetryOutput{}
	fnCacher := func(out *networkmanager.GetNetworkTelemetryOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetNetworkTelemetryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetNetworkTelemetryWithContext(ctx context.Context, input *networkmanager.GetNetworkTelemetryInput, opts ...request.Option) (*networkmanager.GetNetworkTelemetryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetNetworkTelemetryOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetNetworkTelemetryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetResourcePolicy(input *networkmanager.GetResourcePolicyInput) (*networkmanager.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetResourcePolicyOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetResourcePolicyWithContext(ctx context.Context, input *networkmanager.GetResourcePolicyInput, opts ...request.Option) (*networkmanager.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetResourcePolicyOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetRouteAnalysis(input *networkmanager.GetRouteAnalysisInput) (*networkmanager.GetRouteAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetRouteAnalysisOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetRouteAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetRouteAnalysisWithContext(ctx context.Context, input *networkmanager.GetRouteAnalysisInput, opts ...request.Option) (*networkmanager.GetRouteAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetRouteAnalysisOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetRouteAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetSiteToSiteVpnAttachment(input *networkmanager.GetSiteToSiteVpnAttachmentInput) (*networkmanager.GetSiteToSiteVpnAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetSiteToSiteVpnAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetSiteToSiteVpnAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetSiteToSiteVpnAttachmentWithContext(ctx context.Context, input *networkmanager.GetSiteToSiteVpnAttachmentInput, opts ...request.Option) (*networkmanager.GetSiteToSiteVpnAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetSiteToSiteVpnAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetSiteToSiteVpnAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetSites(input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetSitesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetSites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetSitesPages(input *networkmanager.GetSitesInput, fn func(*networkmanager.GetSitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetSitesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetSitesOutput{}
	fnCacher := func(out *networkmanager.GetSitesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetSitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetSitesPagesWithContext(ctx context.Context, input *networkmanager.GetSitesInput, fn func(*networkmanager.GetSitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetSitesOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetSitesOutput{}
	fnCacher := func(out *networkmanager.GetSitesOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetSitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetSitesWithContext(ctx context.Context, input *networkmanager.GetSitesInput, opts ...request.Option) (*networkmanager.GetSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetSitesOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetSitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayConnectPeerAssociations(input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput) (*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayConnectPeerAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayConnectPeerAssociationsPages(input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput, fn func(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetTransitGatewayConnectPeerAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetTransitGatewayConnectPeerAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetTransitGatewayConnectPeerAssociationsPagesWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput, fn func(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetTransitGatewayConnectPeerAssociationsOutput{}
	fnCacher := func(out *networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetTransitGatewayConnectPeerAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetTransitGatewayConnectPeerAssociationsWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput, opts ...request.Option) (*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayConnectPeerAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayPeering(input *networkmanager.GetTransitGatewayPeeringInput) (*networkmanager.GetTransitGatewayPeeringOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayPeeringOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayPeering(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayPeeringWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayPeeringInput, opts ...request.Option) (*networkmanager.GetTransitGatewayPeeringOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayPeeringOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayPeeringWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayRegistrations(input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayRegistrationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayRegistrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayRegistrationsPages(input *networkmanager.GetTransitGatewayRegistrationsInput, fn func(*networkmanager.GetTransitGatewayRegistrationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetTransitGatewayRegistrationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetTransitGatewayRegistrationsOutput{}
	fnCacher := func(out *networkmanager.GetTransitGatewayRegistrationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetTransitGatewayRegistrationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetTransitGatewayRegistrationsPagesWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayRegistrationsInput, fn func(*networkmanager.GetTransitGatewayRegistrationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.GetTransitGatewayRegistrationsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.GetTransitGatewayRegistrationsOutput{}
	fnCacher := func(out *networkmanager.GetTransitGatewayRegistrationsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.GetTransitGatewayRegistrationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) GetTransitGatewayRegistrationsWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayRegistrationsInput, opts ...request.Option) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayRegistrationsOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayRegistrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayRouteTableAttachment(input *networkmanager.GetTransitGatewayRouteTableAttachmentInput) (*networkmanager.GetTransitGatewayRouteTableAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayRouteTableAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayRouteTableAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetTransitGatewayRouteTableAttachmentWithContext(ctx context.Context, input *networkmanager.GetTransitGatewayRouteTableAttachmentInput, opts ...request.Option) (*networkmanager.GetTransitGatewayRouteTableAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetTransitGatewayRouteTableAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetTransitGatewayRouteTableAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetVpcAttachment(input *networkmanager.GetVpcAttachmentInput) (*networkmanager.GetVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetVpcAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) GetVpcAttachmentWithContext(ctx context.Context, input *networkmanager.GetVpcAttachmentInput, opts ...request.Option) (*networkmanager.GetVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.GetVpcAttachmentOutput), nil
	}
	out, err := c.NetworkManagerAPI.GetVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListAttachments(input *networkmanager.ListAttachmentsInput) (*networkmanager.ListAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListAttachmentsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListAttachmentsPages(input *networkmanager.ListAttachmentsInput, fn func(*networkmanager.ListAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListAttachmentsOutput{}
	fnCacher := func(out *networkmanager.ListAttachmentsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListAttachmentsPagesWithContext(ctx context.Context, input *networkmanager.ListAttachmentsInput, fn func(*networkmanager.ListAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListAttachmentsOutput{}
	fnCacher := func(out *networkmanager.ListAttachmentsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListAttachmentsWithContext(ctx context.Context, input *networkmanager.ListAttachmentsInput, opts ...request.Option) (*networkmanager.ListAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListAttachmentsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListConnectPeers(input *networkmanager.ListConnectPeersInput) (*networkmanager.ListConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListConnectPeersOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListConnectPeers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListConnectPeersPages(input *networkmanager.ListConnectPeersInput, fn func(*networkmanager.ListConnectPeersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListConnectPeersOutput{}
	fnCacher := func(out *networkmanager.ListConnectPeersOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListConnectPeersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListConnectPeersPagesWithContext(ctx context.Context, input *networkmanager.ListConnectPeersInput, fn func(*networkmanager.ListConnectPeersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListConnectPeersOutput{}
	fnCacher := func(out *networkmanager.ListConnectPeersOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListConnectPeersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListConnectPeersWithContext(ctx context.Context, input *networkmanager.ListConnectPeersInput, opts ...request.Option) (*networkmanager.ListConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListConnectPeersOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListConnectPeersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListCoreNetworkPolicyVersions(input *networkmanager.ListCoreNetworkPolicyVersionsInput) (*networkmanager.ListCoreNetworkPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListCoreNetworkPolicyVersionsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListCoreNetworkPolicyVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListCoreNetworkPolicyVersionsPages(input *networkmanager.ListCoreNetworkPolicyVersionsInput, fn func(*networkmanager.ListCoreNetworkPolicyVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListCoreNetworkPolicyVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListCoreNetworkPolicyVersionsOutput{}
	fnCacher := func(out *networkmanager.ListCoreNetworkPolicyVersionsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListCoreNetworkPolicyVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListCoreNetworkPolicyVersionsPagesWithContext(ctx context.Context, input *networkmanager.ListCoreNetworkPolicyVersionsInput, fn func(*networkmanager.ListCoreNetworkPolicyVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListCoreNetworkPolicyVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListCoreNetworkPolicyVersionsOutput{}
	fnCacher := func(out *networkmanager.ListCoreNetworkPolicyVersionsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListCoreNetworkPolicyVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListCoreNetworkPolicyVersionsWithContext(ctx context.Context, input *networkmanager.ListCoreNetworkPolicyVersionsInput, opts ...request.Option) (*networkmanager.ListCoreNetworkPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListCoreNetworkPolicyVersionsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListCoreNetworkPolicyVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListCoreNetworks(input *networkmanager.ListCoreNetworksInput) (*networkmanager.ListCoreNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListCoreNetworksOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListCoreNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListCoreNetworksPages(input *networkmanager.ListCoreNetworksInput, fn func(*networkmanager.ListCoreNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListCoreNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListCoreNetworksOutput{}
	fnCacher := func(out *networkmanager.ListCoreNetworksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListCoreNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListCoreNetworksPagesWithContext(ctx context.Context, input *networkmanager.ListCoreNetworksInput, fn func(*networkmanager.ListCoreNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListCoreNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListCoreNetworksOutput{}
	fnCacher := func(out *networkmanager.ListCoreNetworksOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListCoreNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListCoreNetworksWithContext(ctx context.Context, input *networkmanager.ListCoreNetworksInput, opts ...request.Option) (*networkmanager.ListCoreNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListCoreNetworksOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListCoreNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListOrganizationServiceAccessStatus(input *networkmanager.ListOrganizationServiceAccessStatusInput) (*networkmanager.ListOrganizationServiceAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListOrganizationServiceAccessStatusOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListOrganizationServiceAccessStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListOrganizationServiceAccessStatusWithContext(ctx context.Context, input *networkmanager.ListOrganizationServiceAccessStatusInput, opts ...request.Option) (*networkmanager.ListOrganizationServiceAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListOrganizationServiceAccessStatusOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListOrganizationServiceAccessStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListPeerings(input *networkmanager.ListPeeringsInput) (*networkmanager.ListPeeringsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListPeeringsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListPeerings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListPeeringsPages(input *networkmanager.ListPeeringsInput, fn func(*networkmanager.ListPeeringsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListPeeringsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListPeeringsOutput{}
	fnCacher := func(out *networkmanager.ListPeeringsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListPeeringsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListPeeringsPagesWithContext(ctx context.Context, input *networkmanager.ListPeeringsInput, fn func(*networkmanager.ListPeeringsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkmanager.ListPeeringsOutput), false)
		return nil
	}
	cachable := true
	output := &networkmanager.ListPeeringsOutput{}
	fnCacher := func(out *networkmanager.ListPeeringsOutput, more bool) bool {
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
	if err := c.NetworkManagerAPI.ListPeeringsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkManager) ListPeeringsWithContext(ctx context.Context, input *networkmanager.ListPeeringsInput, opts ...request.Option) (*networkmanager.ListPeeringsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListPeeringsOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListPeeringsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListTagsForResource(input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListTagsForResourceOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkManager) ListTagsForResourceWithContext(ctx context.Context, input *networkmanager.ListTagsForResourceInput, opts ...request.Option) (*networkmanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkmanager.ListTagsForResourceOutput), nil
	}
	out, err := c.NetworkManagerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

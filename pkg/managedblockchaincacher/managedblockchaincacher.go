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
package managedblockchaincacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ManagedBlockchain struct {
	managedblockchainiface.ManagedBlockchainAPI
	cache *cache.Cache
}

func New(managedblockchainapi managedblockchainiface.ManagedBlockchainAPI) *ManagedBlockchain {
	return &ManagedBlockchain{
		ManagedBlockchainAPI: managedblockchainapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ManagedBlockchain) GetAccessor(input *managedblockchain.GetAccessorInput) (*managedblockchain.GetAccessorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetAccessorOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetAccessor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetAccessorWithContext(ctx context.Context, input *managedblockchain.GetAccessorInput, opts ...request.Option) (*managedblockchain.GetAccessorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetAccessorOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetAccessorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetMember(input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetMemberOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetMemberWithContext(ctx context.Context, input *managedblockchain.GetMemberInput, opts ...request.Option) (*managedblockchain.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetMemberOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetNetwork(input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetNetworkOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetNetworkWithContext(ctx context.Context, input *managedblockchain.GetNetworkInput, opts ...request.Option) (*managedblockchain.GetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetNetworkOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetNode(input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetNodeOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetNode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetNodeWithContext(ctx context.Context, input *managedblockchain.GetNodeInput, opts ...request.Option) (*managedblockchain.GetNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetNodeOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetNodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetProposal(input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetProposalOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetProposal(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) GetProposalWithContext(ctx context.Context, input *managedblockchain.GetProposalInput, opts ...request.Option) (*managedblockchain.GetProposalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.GetProposalOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.GetProposalWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListAccessors(input *managedblockchain.ListAccessorsInput) (*managedblockchain.ListAccessorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListAccessorsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListAccessors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListAccessorsPages(input *managedblockchain.ListAccessorsInput, fn func(*managedblockchain.ListAccessorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListAccessorsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListAccessorsOutput{}
	fnCacher := func(out *managedblockchain.ListAccessorsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListAccessorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListAccessorsPagesWithContext(ctx context.Context, input *managedblockchain.ListAccessorsInput, fn func(*managedblockchain.ListAccessorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListAccessorsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListAccessorsOutput{}
	fnCacher := func(out *managedblockchain.ListAccessorsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListAccessorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListAccessorsWithContext(ctx context.Context, input *managedblockchain.ListAccessorsInput, opts ...request.Option) (*managedblockchain.ListAccessorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListAccessorsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListAccessorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListInvitations(input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListInvitationsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListInvitationsPages(input *managedblockchain.ListInvitationsInput, fn func(*managedblockchain.ListInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListInvitationsOutput{}
	fnCacher := func(out *managedblockchain.ListInvitationsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListInvitationsPagesWithContext(ctx context.Context, input *managedblockchain.ListInvitationsInput, fn func(*managedblockchain.ListInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListInvitationsOutput{}
	fnCacher := func(out *managedblockchain.ListInvitationsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListInvitationsWithContext(ctx context.Context, input *managedblockchain.ListInvitationsInput, opts ...request.Option) (*managedblockchain.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListInvitationsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListMembers(input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListMembersOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListMembersPages(input *managedblockchain.ListMembersInput, fn func(*managedblockchain.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListMembersOutput{}
	fnCacher := func(out *managedblockchain.ListMembersOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListMembersPagesWithContext(ctx context.Context, input *managedblockchain.ListMembersInput, fn func(*managedblockchain.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListMembersOutput{}
	fnCacher := func(out *managedblockchain.ListMembersOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListMembersWithContext(ctx context.Context, input *managedblockchain.ListMembersInput, opts ...request.Option) (*managedblockchain.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListMembersOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListNetworks(input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListNetworksOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListNetworksPages(input *managedblockchain.ListNetworksInput, fn func(*managedblockchain.ListNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListNetworksOutput{}
	fnCacher := func(out *managedblockchain.ListNetworksOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListNetworksPagesWithContext(ctx context.Context, input *managedblockchain.ListNetworksInput, fn func(*managedblockchain.ListNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListNetworksOutput{}
	fnCacher := func(out *managedblockchain.ListNetworksOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListNetworksWithContext(ctx context.Context, input *managedblockchain.ListNetworksInput, opts ...request.Option) (*managedblockchain.ListNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListNetworksOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListNodes(input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListNodesOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListNodesPages(input *managedblockchain.ListNodesInput, fn func(*managedblockchain.ListNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListNodesOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListNodesOutput{}
	fnCacher := func(out *managedblockchain.ListNodesOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListNodesPagesWithContext(ctx context.Context, input *managedblockchain.ListNodesInput, fn func(*managedblockchain.ListNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListNodesOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListNodesOutput{}
	fnCacher := func(out *managedblockchain.ListNodesOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListNodesWithContext(ctx context.Context, input *managedblockchain.ListNodesInput, opts ...request.Option) (*managedblockchain.ListNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListNodesOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListProposalVotes(input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListProposalVotesOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListProposalVotes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListProposalVotesPages(input *managedblockchain.ListProposalVotesInput, fn func(*managedblockchain.ListProposalVotesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListProposalVotesOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListProposalVotesOutput{}
	fnCacher := func(out *managedblockchain.ListProposalVotesOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListProposalVotesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListProposalVotesPagesWithContext(ctx context.Context, input *managedblockchain.ListProposalVotesInput, fn func(*managedblockchain.ListProposalVotesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListProposalVotesOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListProposalVotesOutput{}
	fnCacher := func(out *managedblockchain.ListProposalVotesOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListProposalVotesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListProposalVotesWithContext(ctx context.Context, input *managedblockchain.ListProposalVotesInput, opts ...request.Option) (*managedblockchain.ListProposalVotesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListProposalVotesOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListProposalVotesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListProposals(input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListProposalsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListProposals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListProposalsPages(input *managedblockchain.ListProposalsInput, fn func(*managedblockchain.ListProposalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListProposalsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListProposalsOutput{}
	fnCacher := func(out *managedblockchain.ListProposalsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListProposalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListProposalsPagesWithContext(ctx context.Context, input *managedblockchain.ListProposalsInput, fn func(*managedblockchain.ListProposalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedblockchain.ListProposalsOutput), false)
		return nil
	}
	cachable := true
	output := &managedblockchain.ListProposalsOutput{}
	fnCacher := func(out *managedblockchain.ListProposalsOutput, more bool) bool {
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
	if err := c.ManagedBlockchainAPI.ListProposalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedBlockchain) ListProposalsWithContext(ctx context.Context, input *managedblockchain.ListProposalsInput, opts ...request.Option) (*managedblockchain.ListProposalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListProposalsOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListProposalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListTagsForResource(input *managedblockchain.ListTagsForResourceInput) (*managedblockchain.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListTagsForResourceOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedBlockchain) ListTagsForResourceWithContext(ctx context.Context, input *managedblockchain.ListTagsForResourceInput, opts ...request.Option) (*managedblockchain.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedblockchain.ListTagsForResourceOutput), nil
	}
	out, err := c.ManagedBlockchainAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

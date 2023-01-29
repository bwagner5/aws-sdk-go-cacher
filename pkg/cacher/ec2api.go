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
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EC2 struct {
	ec2iface.EC2API
	cache *cache.Cache
}

func New(ec2api ec2iface.EC2API) *EC2 {
	return &EC2{
		EC2API: ec2api,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}

func (c *EC2) AcceptAddressTransfer(input *ec2.AcceptAddressTransferInput) (*ec2.AcceptAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptAddressTransferOutput), nil
	}
	out, err := c.EC2API.AcceptAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptAddressTransferWithContext(ctx context.Context, input *ec2.AcceptAddressTransferInput, opts ...request.Option) (*ec2.AcceptAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptAddressTransferOutput), nil
	}
	out, err := c.EC2API.AcceptAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptReservedInstancesExchangeQuote(input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.AcceptReservedInstancesExchangeQuote(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptReservedInstancesExchangeQuoteWithContext(ctx context.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput, opts ...request.Option) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.AcceptReservedInstancesExchangeQuoteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayMulticastDomainAssociations(input *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayPeeringAttachment(input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayVpcAttachment(input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.AcceptTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptVpcEndpointConnections(input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.AcceptVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.AcceptVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.AcceptVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptVpcPeeringConnection(input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.AcceptVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AcceptVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.AcceptVpcPeeringConnectionInput, opts ...request.Option) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.AcceptVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AdvertiseByoipCidr(input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AdvertiseByoipCidrOutput), nil
	}
	out, err := c.EC2API.AdvertiseByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AdvertiseByoipCidrWithContext(ctx context.Context, input *ec2.AdvertiseByoipCidrInput, opts ...request.Option) (*ec2.AdvertiseByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AdvertiseByoipCidrOutput), nil
	}
	out, err := c.EC2API.AdvertiseByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateAddress(input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateAddressOutput), nil
	}
	out, err := c.EC2API.AllocateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateAddressWithContext(ctx context.Context, input *ec2.AllocateAddressInput, opts ...request.Option) (*ec2.AllocateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateAddressOutput), nil
	}
	out, err := c.EC2API.AllocateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateHosts(input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateHostsOutput), nil
	}
	out, err := c.EC2API.AllocateHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateHostsWithContext(ctx context.Context, input *ec2.AllocateHostsInput, opts ...request.Option) (*ec2.AllocateHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateHostsOutput), nil
	}
	out, err := c.EC2API.AllocateHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateIpamPoolCidr(input *ec2.AllocateIpamPoolCidrInput) (*ec2.AllocateIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.AllocateIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AllocateIpamPoolCidrWithContext(ctx context.Context, input *ec2.AllocateIpamPoolCidrInput, opts ...request.Option) (*ec2.AllocateIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.AllocateIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ApplySecurityGroupsToClientVpnTargetNetwork(input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.ApplySecurityGroupsToClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ApplySecurityGroupsToClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.ApplySecurityGroupsToClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssignIpv6Addresses(input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignIpv6AddressesOutput), nil
	}
	out, err := c.EC2API.AssignIpv6Addresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssignIpv6AddressesWithContext(ctx context.Context, input *ec2.AssignIpv6AddressesInput, opts ...request.Option) (*ec2.AssignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignIpv6AddressesOutput), nil
	}
	out, err := c.EC2API.AssignIpv6AddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssignPrivateIpAddresses(input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignPrivateIpAddressesOutput), nil
	}
	out, err := c.EC2API.AssignPrivateIpAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssignPrivateIpAddressesWithContext(ctx context.Context, input *ec2.AssignPrivateIpAddressesInput, opts ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignPrivateIpAddressesOutput), nil
	}
	out, err := c.EC2API.AssignPrivateIpAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateAddress(input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateAddressOutput), nil
	}
	out, err := c.EC2API.AssociateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateAddressWithContext(ctx context.Context, input *ec2.AssociateAddressInput, opts ...request.Option) (*ec2.AssociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateAddressOutput), nil
	}
	out, err := c.EC2API.AssociateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateClientVpnTargetNetwork(input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.AssociateClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.AssociateClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.AssociateClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateDhcpOptions(input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.AssociateDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateDhcpOptionsWithContext(ctx context.Context, input *ec2.AssociateDhcpOptionsInput, opts ...request.Option) (*ec2.AssociateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.AssociateDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateEnclaveCertificateIamRole(input *ec2.AssociateEnclaveCertificateIamRoleInput) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.EC2API.AssociateEnclaveCertificateIamRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateEnclaveCertificateIamRoleWithContext(ctx context.Context, input *ec2.AssociateEnclaveCertificateIamRoleInput, opts ...request.Option) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.EC2API.AssociateEnclaveCertificateIamRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateIamInstanceProfile(input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateIamInstanceProfileOutput), nil
	}
	out, err := c.EC2API.AssociateIamInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateIamInstanceProfileWithContext(ctx context.Context, input *ec2.AssociateIamInstanceProfileInput, opts ...request.Option) (*ec2.AssociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateIamInstanceProfileOutput), nil
	}
	out, err := c.EC2API.AssociateIamInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateInstanceEventWindow(input *ec2.AssociateInstanceEventWindowInput) (*ec2.AssociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.AssociateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateInstanceEventWindowWithContext(ctx context.Context, input *ec2.AssociateInstanceEventWindowInput, opts ...request.Option) (*ec2.AssociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.AssociateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateRouteTable(input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateRouteTableOutput), nil
	}
	out, err := c.EC2API.AssociateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateRouteTableWithContext(ctx context.Context, input *ec2.AssociateRouteTableInput, opts ...request.Option) (*ec2.AssociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateRouteTableOutput), nil
	}
	out, err := c.EC2API.AssociateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateSubnetCidrBlock(input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateSubnetCidrBlockOutput), nil
	}
	out, err := c.EC2API.AssociateSubnetCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateSubnetCidrBlockWithContext(ctx context.Context, input *ec2.AssociateSubnetCidrBlockInput, opts ...request.Option) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateSubnetCidrBlockOutput), nil
	}
	out, err := c.EC2API.AssociateSubnetCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayMulticastDomain(input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayPolicyTable(input *ec2.AssociateTransitGatewayPolicyTableInput) (*ec2.AssociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.AssociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayRouteTable(input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.AssociateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTrunkInterface(input *ec2.AssociateTrunkInterfaceInput) (*ec2.AssociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTrunkInterfaceOutput), nil
	}
	out, err := c.EC2API.AssociateTrunkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateTrunkInterfaceWithContext(ctx context.Context, input *ec2.AssociateTrunkInterfaceInput, opts ...request.Option) (*ec2.AssociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTrunkInterfaceOutput), nil
	}
	out, err := c.EC2API.AssociateTrunkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateVpcCidrBlock(input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateVpcCidrBlockOutput), nil
	}
	out, err := c.EC2API.AssociateVpcCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AssociateVpcCidrBlockWithContext(ctx context.Context, input *ec2.AssociateVpcCidrBlockInput, opts ...request.Option) (*ec2.AssociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateVpcCidrBlockOutput), nil
	}
	out, err := c.EC2API.AssociateVpcCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachClassicLinkVpc(input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachClassicLinkVpcOutput), nil
	}
	out, err := c.EC2API.AttachClassicLinkVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachClassicLinkVpcWithContext(ctx context.Context, input *ec2.AttachClassicLinkVpcInput, opts ...request.Option) (*ec2.AttachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachClassicLinkVpcOutput), nil
	}
	out, err := c.EC2API.AttachClassicLinkVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachInternetGateway(input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachInternetGatewayOutput), nil
	}
	out, err := c.EC2API.AttachInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachInternetGatewayWithContext(ctx context.Context, input *ec2.AttachInternetGatewayInput, opts ...request.Option) (*ec2.AttachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachInternetGatewayOutput), nil
	}
	out, err := c.EC2API.AttachInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachNetworkInterface(input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.AttachNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachNetworkInterfaceWithContext(ctx context.Context, input *ec2.AttachNetworkInterfaceInput, opts ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.AttachNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachVerifiedAccessTrustProvider(input *ec2.AttachVerifiedAccessTrustProviderInput) (*ec2.AttachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.AttachVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.AttachVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.AttachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.AttachVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachVpnGateway(input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVpnGatewayOutput), nil
	}
	out, err := c.EC2API.AttachVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AttachVpnGatewayWithContext(ctx context.Context, input *ec2.AttachVpnGatewayInput, opts ...request.Option) (*ec2.AttachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVpnGatewayOutput), nil
	}
	out, err := c.EC2API.AttachVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeClientVpnIngress(input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeClientVpnIngressOutput), nil
	}
	out, err := c.EC2API.AuthorizeClientVpnIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeClientVpnIngressWithContext(ctx context.Context, input *ec2.AuthorizeClientVpnIngressInput, opts ...request.Option) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeClientVpnIngressOutput), nil
	}
	out, err := c.EC2API.AuthorizeClientVpnIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeSecurityGroupEgress(input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupEgressOutput), nil
	}
	out, err := c.EC2API.AuthorizeSecurityGroupEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeSecurityGroupEgressWithContext(ctx context.Context, input *ec2.AuthorizeSecurityGroupEgressInput, opts ...request.Option) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupEgressOutput), nil
	}
	out, err := c.EC2API.AuthorizeSecurityGroupEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupIngressOutput), nil
	}
	out, err := c.EC2API.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) AuthorizeSecurityGroupIngressWithContext(ctx context.Context, input *ec2.AuthorizeSecurityGroupIngressInput, opts ...request.Option) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupIngressOutput), nil
	}
	out, err := c.EC2API.AuthorizeSecurityGroupIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) BundleInstance(input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.BundleInstanceOutput), nil
	}
	out, err := c.EC2API.BundleInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) BundleInstanceWithContext(ctx context.Context, input *ec2.BundleInstanceInput, opts ...request.Option) (*ec2.BundleInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.BundleInstanceOutput), nil
	}
	out, err := c.EC2API.BundleInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelBundleTask(input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelBundleTaskOutput), nil
	}
	out, err := c.EC2API.CancelBundleTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelBundleTaskWithContext(ctx context.Context, input *ec2.CancelBundleTaskInput, opts ...request.Option) (*ec2.CancelBundleTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelBundleTaskOutput), nil
	}
	out, err := c.EC2API.CancelBundleTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelCapacityReservation(input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationOutput), nil
	}
	out, err := c.EC2API.CancelCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelCapacityReservationFleets(input *ec2.CancelCapacityReservationFleetsInput) (*ec2.CancelCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.CancelCapacityReservationFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelCapacityReservationFleetsWithContext(ctx context.Context, input *ec2.CancelCapacityReservationFleetsInput, opts ...request.Option) (*ec2.CancelCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.CancelCapacityReservationFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelCapacityReservationWithContext(ctx context.Context, input *ec2.CancelCapacityReservationInput, opts ...request.Option) (*ec2.CancelCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationOutput), nil
	}
	out, err := c.EC2API.CancelCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelConversionTask(input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelConversionTaskOutput), nil
	}
	out, err := c.EC2API.CancelConversionTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelConversionTaskWithContext(ctx context.Context, input *ec2.CancelConversionTaskInput, opts ...request.Option) (*ec2.CancelConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelConversionTaskOutput), nil
	}
	out, err := c.EC2API.CancelConversionTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelExportTask(input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelExportTaskOutput), nil
	}
	out, err := c.EC2API.CancelExportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelExportTaskWithContext(ctx context.Context, input *ec2.CancelExportTaskInput, opts ...request.Option) (*ec2.CancelExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelExportTaskOutput), nil
	}
	out, err := c.EC2API.CancelExportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelImageLaunchPermission(input *ec2.CancelImageLaunchPermissionInput) (*ec2.CancelImageLaunchPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImageLaunchPermissionOutput), nil
	}
	out, err := c.EC2API.CancelImageLaunchPermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelImageLaunchPermissionWithContext(ctx context.Context, input *ec2.CancelImageLaunchPermissionInput, opts ...request.Option) (*ec2.CancelImageLaunchPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImageLaunchPermissionOutput), nil
	}
	out, err := c.EC2API.CancelImageLaunchPermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelImportTask(input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImportTaskOutput), nil
	}
	out, err := c.EC2API.CancelImportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelImportTaskWithContext(ctx context.Context, input *ec2.CancelImportTaskInput, opts ...request.Option) (*ec2.CancelImportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImportTaskOutput), nil
	}
	out, err := c.EC2API.CancelImportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelReservedInstancesListing(input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelReservedInstancesListingOutput), nil
	}
	out, err := c.EC2API.CancelReservedInstancesListing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelReservedInstancesListingWithContext(ctx context.Context, input *ec2.CancelReservedInstancesListingInput, opts ...request.Option) (*ec2.CancelReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelReservedInstancesListingOutput), nil
	}
	out, err := c.EC2API.CancelReservedInstancesListingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelSpotFleetRequests(input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.CancelSpotFleetRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelSpotFleetRequestsWithContext(ctx context.Context, input *ec2.CancelSpotFleetRequestsInput, opts ...request.Option) (*ec2.CancelSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.CancelSpotFleetRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelSpotInstanceRequests(input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.CancelSpotInstanceRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CancelSpotInstanceRequestsWithContext(ctx context.Context, input *ec2.CancelSpotInstanceRequestsInput, opts ...request.Option) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.CancelSpotInstanceRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ConfirmProductInstance(input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ConfirmProductInstanceOutput), nil
	}
	out, err := c.EC2API.ConfirmProductInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ConfirmProductInstanceWithContext(ctx context.Context, input *ec2.ConfirmProductInstanceInput, opts ...request.Option) (*ec2.ConfirmProductInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ConfirmProductInstanceOutput), nil
	}
	out, err := c.EC2API.ConfirmProductInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopyFpgaImage(input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyFpgaImageOutput), nil
	}
	out, err := c.EC2API.CopyFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopyFpgaImageWithContext(ctx context.Context, input *ec2.CopyFpgaImageInput, opts ...request.Option) (*ec2.CopyFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyFpgaImageOutput), nil
	}
	out, err := c.EC2API.CopyFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopyImage(input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyImageOutput), nil
	}
	out, err := c.EC2API.CopyImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopyImageWithContext(ctx context.Context, input *ec2.CopyImageInput, opts ...request.Option) (*ec2.CopyImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyImageOutput), nil
	}
	out, err := c.EC2API.CopyImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopySnapshot(input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopySnapshotOutput), nil
	}
	out, err := c.EC2API.CopySnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CopySnapshotWithContext(ctx context.Context, input *ec2.CopySnapshotInput, opts ...request.Option) (*ec2.CopySnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopySnapshotOutput), nil
	}
	out, err := c.EC2API.CopySnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCapacityReservation(input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationOutput), nil
	}
	out, err := c.EC2API.CreateCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCapacityReservationFleet(input *ec2.CreateCapacityReservationFleetInput) (*ec2.CreateCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationFleetOutput), nil
	}
	out, err := c.EC2API.CreateCapacityReservationFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCapacityReservationFleetWithContext(ctx context.Context, input *ec2.CreateCapacityReservationFleetInput, opts ...request.Option) (*ec2.CreateCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationFleetOutput), nil
	}
	out, err := c.EC2API.CreateCapacityReservationFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCapacityReservationWithContext(ctx context.Context, input *ec2.CreateCapacityReservationInput, opts ...request.Option) (*ec2.CreateCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationOutput), nil
	}
	out, err := c.EC2API.CreateCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCarrierGateway(input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCarrierGatewayOutput), nil
	}
	out, err := c.EC2API.CreateCarrierGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCarrierGatewayWithContext(ctx context.Context, input *ec2.CreateCarrierGatewayInput, opts ...request.Option) (*ec2.CreateCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCarrierGatewayOutput), nil
	}
	out, err := c.EC2API.CreateCarrierGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateClientVpnEndpoint(input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.CreateClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateClientVpnEndpointWithContext(ctx context.Context, input *ec2.CreateClientVpnEndpointInput, opts ...request.Option) (*ec2.CreateClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.CreateClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateClientVpnRoute(input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnRouteOutput), nil
	}
	out, err := c.EC2API.CreateClientVpnRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateClientVpnRouteWithContext(ctx context.Context, input *ec2.CreateClientVpnRouteInput, opts ...request.Option) (*ec2.CreateClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnRouteOutput), nil
	}
	out, err := c.EC2API.CreateClientVpnRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCoipCidr(input *ec2.CreateCoipCidrInput) (*ec2.CreateCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipCidrOutput), nil
	}
	out, err := c.EC2API.CreateCoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCoipCidrWithContext(ctx context.Context, input *ec2.CreateCoipCidrInput, opts ...request.Option) (*ec2.CreateCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipCidrOutput), nil
	}
	out, err := c.EC2API.CreateCoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCoipPool(input *ec2.CreateCoipPoolInput) (*ec2.CreateCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipPoolOutput), nil
	}
	out, err := c.EC2API.CreateCoipPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCoipPoolWithContext(ctx context.Context, input *ec2.CreateCoipPoolInput, opts ...request.Option) (*ec2.CreateCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipPoolOutput), nil
	}
	out, err := c.EC2API.CreateCoipPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCustomerGateway(input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCustomerGatewayOutput), nil
	}
	out, err := c.EC2API.CreateCustomerGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateCustomerGatewayWithContext(ctx context.Context, input *ec2.CreateCustomerGatewayInput, opts ...request.Option) (*ec2.CreateCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCustomerGatewayOutput), nil
	}
	out, err := c.EC2API.CreateCustomerGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDefaultSubnet(input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultSubnetOutput), nil
	}
	out, err := c.EC2API.CreateDefaultSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDefaultSubnetWithContext(ctx context.Context, input *ec2.CreateDefaultSubnetInput, opts ...request.Option) (*ec2.CreateDefaultSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultSubnetOutput), nil
	}
	out, err := c.EC2API.CreateDefaultSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDefaultVpc(input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultVpcOutput), nil
	}
	out, err := c.EC2API.CreateDefaultVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDefaultVpcWithContext(ctx context.Context, input *ec2.CreateDefaultVpcInput, opts ...request.Option) (*ec2.CreateDefaultVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultVpcOutput), nil
	}
	out, err := c.EC2API.CreateDefaultVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDhcpOptions(input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.CreateDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateDhcpOptionsWithContext(ctx context.Context, input *ec2.CreateDhcpOptionsInput, opts ...request.Option) (*ec2.CreateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.CreateDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateEgressOnlyInternetGateway(input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.EC2API.CreateEgressOnlyInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateEgressOnlyInternetGatewayWithContext(ctx context.Context, input *ec2.CreateEgressOnlyInternetGatewayInput, opts ...request.Option) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.EC2API.CreateEgressOnlyInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFleet(input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFleetOutput), nil
	}
	out, err := c.EC2API.CreateFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFleetWithContext(ctx context.Context, input *ec2.CreateFleetInput, opts ...request.Option) (*ec2.CreateFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFleetOutput), nil
	}
	out, err := c.EC2API.CreateFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFlowLogs(input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFlowLogsOutput), nil
	}
	out, err := c.EC2API.CreateFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFlowLogsWithContext(ctx context.Context, input *ec2.CreateFlowLogsInput, opts ...request.Option) (*ec2.CreateFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFlowLogsOutput), nil
	}
	out, err := c.EC2API.CreateFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFpgaImage(input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFpgaImageOutput), nil
	}
	out, err := c.EC2API.CreateFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateFpgaImageWithContext(ctx context.Context, input *ec2.CreateFpgaImageInput, opts ...request.Option) (*ec2.CreateFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFpgaImageOutput), nil
	}
	out, err := c.EC2API.CreateFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateImage(input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateImageOutput), nil
	}
	out, err := c.EC2API.CreateImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateImageWithContext(ctx context.Context, input *ec2.CreateImageInput, opts ...request.Option) (*ec2.CreateImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateImageOutput), nil
	}
	out, err := c.EC2API.CreateImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInstanceEventWindow(input *ec2.CreateInstanceEventWindowInput) (*ec2.CreateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.CreateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInstanceEventWindowWithContext(ctx context.Context, input *ec2.CreateInstanceEventWindowInput, opts ...request.Option) (*ec2.CreateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.CreateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInstanceExportTask(input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceExportTaskOutput), nil
	}
	out, err := c.EC2API.CreateInstanceExportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInstanceExportTaskWithContext(ctx context.Context, input *ec2.CreateInstanceExportTaskInput, opts ...request.Option) (*ec2.CreateInstanceExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceExportTaskOutput), nil
	}
	out, err := c.EC2API.CreateInstanceExportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInternetGateway(input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInternetGatewayOutput), nil
	}
	out, err := c.EC2API.CreateInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateInternetGatewayWithContext(ctx context.Context, input *ec2.CreateInternetGatewayInput, opts ...request.Option) (*ec2.CreateInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInternetGatewayOutput), nil
	}
	out, err := c.EC2API.CreateInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpam(input *ec2.CreateIpamInput) (*ec2.CreateIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamOutput), nil
	}
	out, err := c.EC2API.CreateIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpamPool(input *ec2.CreateIpamPoolInput) (*ec2.CreateIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamPoolOutput), nil
	}
	out, err := c.EC2API.CreateIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpamPoolWithContext(ctx context.Context, input *ec2.CreateIpamPoolInput, opts ...request.Option) (*ec2.CreateIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamPoolOutput), nil
	}
	out, err := c.EC2API.CreateIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpamScope(input *ec2.CreateIpamScopeInput) (*ec2.CreateIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamScopeOutput), nil
	}
	out, err := c.EC2API.CreateIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpamScopeWithContext(ctx context.Context, input *ec2.CreateIpamScopeInput, opts ...request.Option) (*ec2.CreateIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamScopeOutput), nil
	}
	out, err := c.EC2API.CreateIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateIpamWithContext(ctx context.Context, input *ec2.CreateIpamInput, opts ...request.Option) (*ec2.CreateIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamOutput), nil
	}
	out, err := c.EC2API.CreateIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateKeyPair(input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateKeyPairOutput), nil
	}
	out, err := c.EC2API.CreateKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateKeyPairWithContext(ctx context.Context, input *ec2.CreateKeyPairInput, opts ...request.Option) (*ec2.CreateKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateKeyPairOutput), nil
	}
	out, err := c.EC2API.CreateKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLaunchTemplate(input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.CreateLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLaunchTemplateVersion(input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateVersionOutput), nil
	}
	out, err := c.EC2API.CreateLaunchTemplateVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLaunchTemplateVersionWithContext(ctx context.Context, input *ec2.CreateLaunchTemplateVersionInput, opts ...request.Option) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateVersionOutput), nil
	}
	out, err := c.EC2API.CreateLaunchTemplateVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLaunchTemplateWithContext(ctx context.Context, input *ec2.CreateLaunchTemplateInput, opts ...request.Option) (*ec2.CreateLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.CreateLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRoute(input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTable(input *ec2.CreateLocalGatewayRouteTableInput) (*ec2.CreateLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input *ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput) (*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTableVpcAssociation(input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTableVpcAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTableVpcAssociationWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTableVpcAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteTableWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateLocalGatewayRouteWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.CreateLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateManagedPrefixList(input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.CreateManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateManagedPrefixListWithContext(ctx context.Context, input *ec2.CreateManagedPrefixListInput, opts ...request.Option) (*ec2.CreateManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.CreateManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNatGateway(input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNatGatewayOutput), nil
	}
	out, err := c.EC2API.CreateNatGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNatGatewayWithContext(ctx context.Context, input *ec2.CreateNatGatewayInput, opts ...request.Option) (*ec2.CreateNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNatGatewayOutput), nil
	}
	out, err := c.EC2API.CreateNatGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkAcl(input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclOutput), nil
	}
	out, err := c.EC2API.CreateNetworkAcl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkAclEntry(input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.CreateNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkAclEntryWithContext(ctx context.Context, input *ec2.CreateNetworkAclEntryInput, opts ...request.Option) (*ec2.CreateNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.CreateNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkAclWithContext(ctx context.Context, input *ec2.CreateNetworkAclInput, opts ...request.Option) (*ec2.CreateNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclOutput), nil
	}
	out, err := c.EC2API.CreateNetworkAclWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInsightsAccessScope(input *ec2.CreateNetworkInsightsAccessScopeInput) (*ec2.CreateNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInsightsAccessScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInsightsAccessScopeWithContext(ctx context.Context, input *ec2.CreateNetworkInsightsAccessScopeInput, opts ...request.Option) (*ec2.CreateNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInsightsAccessScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInsightsPath(input *ec2.CreateNetworkInsightsPathInput) (*ec2.CreateNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsPathOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInsightsPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInsightsPathWithContext(ctx context.Context, input *ec2.CreateNetworkInsightsPathInput, opts ...request.Option) (*ec2.CreateNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsPathOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInsightsPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInterface(input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInterfacePermission(input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfacePermissionOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInterfacePermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInterfacePermissionWithContext(ctx context.Context, input *ec2.CreateNetworkInterfacePermissionInput, opts ...request.Option) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfacePermissionOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInterfacePermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateNetworkInterfaceWithContext(ctx context.Context, input *ec2.CreateNetworkInterfaceInput, opts ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.CreateNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreatePlacementGroup(input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePlacementGroupOutput), nil
	}
	out, err := c.EC2API.CreatePlacementGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreatePlacementGroupWithContext(ctx context.Context, input *ec2.CreatePlacementGroupInput, opts ...request.Option) (*ec2.CreatePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePlacementGroupOutput), nil
	}
	out, err := c.EC2API.CreatePlacementGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreatePublicIpv4Pool(input *ec2.CreatePublicIpv4PoolInput) (*ec2.CreatePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePublicIpv4PoolOutput), nil
	}
	out, err := c.EC2API.CreatePublicIpv4Pool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreatePublicIpv4PoolWithContext(ctx context.Context, input *ec2.CreatePublicIpv4PoolInput, opts ...request.Option) (*ec2.CreatePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePublicIpv4PoolOutput), nil
	}
	out, err := c.EC2API.CreatePublicIpv4PoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateReplaceRootVolumeTask(input *ec2.CreateReplaceRootVolumeTaskInput) (*ec2.CreateReplaceRootVolumeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReplaceRootVolumeTaskOutput), nil
	}
	out, err := c.EC2API.CreateReplaceRootVolumeTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateReplaceRootVolumeTaskWithContext(ctx context.Context, input *ec2.CreateReplaceRootVolumeTaskInput, opts ...request.Option) (*ec2.CreateReplaceRootVolumeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReplaceRootVolumeTaskOutput), nil
	}
	out, err := c.EC2API.CreateReplaceRootVolumeTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateReservedInstancesListing(input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReservedInstancesListingOutput), nil
	}
	out, err := c.EC2API.CreateReservedInstancesListing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateReservedInstancesListingWithContext(ctx context.Context, input *ec2.CreateReservedInstancesListingInput, opts ...request.Option) (*ec2.CreateReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReservedInstancesListingOutput), nil
	}
	out, err := c.EC2API.CreateReservedInstancesListingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRestoreImageTask(input *ec2.CreateRestoreImageTaskInput) (*ec2.CreateRestoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRestoreImageTaskOutput), nil
	}
	out, err := c.EC2API.CreateRestoreImageTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRestoreImageTaskWithContext(ctx context.Context, input *ec2.CreateRestoreImageTaskInput, opts ...request.Option) (*ec2.CreateRestoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRestoreImageTaskOutput), nil
	}
	out, err := c.EC2API.CreateRestoreImageTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRoute(input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteOutput), nil
	}
	out, err := c.EC2API.CreateRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRouteTable(input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRouteTableWithContext(ctx context.Context, input *ec2.CreateRouteTableInput, opts ...request.Option) (*ec2.CreateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateRouteWithContext(ctx context.Context, input *ec2.CreateRouteInput, opts ...request.Option) (*ec2.CreateRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteOutput), nil
	}
	out, err := c.EC2API.CreateRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSecurityGroup(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSecurityGroupOutput), nil
	}
	out, err := c.EC2API.CreateSecurityGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSecurityGroupWithContext(ctx context.Context, input *ec2.CreateSecurityGroupInput, opts ...request.Option) (*ec2.CreateSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSecurityGroupOutput), nil
	}
	out, err := c.EC2API.CreateSecurityGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSnapshots(input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSnapshotsOutput), nil
	}
	out, err := c.EC2API.CreateSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSnapshotsWithContext(ctx context.Context, input *ec2.CreateSnapshotsInput, opts ...request.Option) (*ec2.CreateSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSnapshotsOutput), nil
	}
	out, err := c.EC2API.CreateSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSpotDatafeedSubscription(input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.CreateSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.CreateSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.CreateSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateStoreImageTask(input *ec2.CreateStoreImageTaskInput) (*ec2.CreateStoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateStoreImageTaskOutput), nil
	}
	out, err := c.EC2API.CreateStoreImageTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateStoreImageTaskWithContext(ctx context.Context, input *ec2.CreateStoreImageTaskInput, opts ...request.Option) (*ec2.CreateStoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateStoreImageTaskOutput), nil
	}
	out, err := c.EC2API.CreateStoreImageTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSubnet(input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetOutput), nil
	}
	out, err := c.EC2API.CreateSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSubnetCidrReservation(input *ec2.CreateSubnetCidrReservationInput) (*ec2.CreateSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetCidrReservationOutput), nil
	}
	out, err := c.EC2API.CreateSubnetCidrReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSubnetCidrReservationWithContext(ctx context.Context, input *ec2.CreateSubnetCidrReservationInput, opts ...request.Option) (*ec2.CreateSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetCidrReservationOutput), nil
	}
	out, err := c.EC2API.CreateSubnetCidrReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateSubnetWithContext(ctx context.Context, input *ec2.CreateSubnetInput, opts ...request.Option) (*ec2.CreateSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetOutput), nil
	}
	out, err := c.EC2API.CreateSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTags(input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTagsOutput), nil
	}
	out, err := c.EC2API.CreateTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTagsWithContext(ctx context.Context, input *ec2.CreateTagsInput, opts ...request.Option) (*ec2.CreateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTagsOutput), nil
	}
	out, err := c.EC2API.CreateTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorFilter(input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorFilterRule(input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorFilterWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorFilterInput, opts ...request.Option) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorSession(input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorSessionInput, opts ...request.Option) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorTarget(input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorTargetOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTrafficMirrorTargetWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorTargetInput, opts ...request.Option) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorTargetOutput), nil
	}
	out, err := c.EC2API.CreateTrafficMirrorTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGateway(input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayOutput), nil
	}
	out, err := c.EC2API.CreateTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayConnect(input *ec2.CreateTransitGatewayConnectInput) (*ec2.CreateTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayConnect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayConnectPeer(input *ec2.CreateTransitGatewayConnectPeerInput) (*ec2.CreateTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayConnectPeer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayConnectPeerWithContext(ctx context.Context, input *ec2.CreateTransitGatewayConnectPeerInput, opts ...request.Option) (*ec2.CreateTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayConnectPeerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayConnectWithContext(ctx context.Context, input *ec2.CreateTransitGatewayConnectInput, opts ...request.Option) (*ec2.CreateTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayConnectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayMulticastDomain(input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.CreateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPeeringAttachment(input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPolicyTable(input *ec2.CreateTransitGatewayPolicyTableInput) (*ec2.CreateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.CreateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPrefixListReference(input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRoute(input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRouteTable(input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRouteTableAnnouncement(input *ec2.CreateTransitGatewayRouteTableAnnouncementInput) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRouteTableAnnouncement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRouteTableAnnouncementWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteTableAnnouncementInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRouteTableAnnouncementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayRouteWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayVpcAttachment(input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateTransitGatewayWithContext(ctx context.Context, input *ec2.CreateTransitGatewayInput, opts ...request.Option) (*ec2.CreateTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayOutput), nil
	}
	out, err := c.EC2API.CreateTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessEndpoint(input *ec2.CreateVerifiedAccessEndpointInput) (*ec2.CreateVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.CreateVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessGroup(input *ec2.CreateVerifiedAccessGroupInput) (*ec2.CreateVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessGroupInput, opts ...request.Option) (*ec2.CreateVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessInstance(input *ec2.CreateVerifiedAccessInstanceInput) (*ec2.CreateVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.CreateVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessTrustProvider(input *ec2.CreateVerifiedAccessTrustProviderInput) (*ec2.CreateVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.CreateVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.CreateVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpc(input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcOutput), nil
	}
	out, err := c.EC2API.CreateVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpoint(input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpointConnectionNotification(input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpointConnectionNotification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpointConnectionNotificationWithContext(ctx context.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput, opts ...request.Option) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpointConnectionNotificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpointServiceConfiguration(input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpointServiceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpointServiceConfigurationWithContext(ctx context.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput, opts ...request.Option) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpointServiceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcEndpointWithContext(ctx context.Context, input *ec2.CreateVpcEndpointInput, opts ...request.Option) (*ec2.CreateVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointOutput), nil
	}
	out, err := c.EC2API.CreateVpcEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcPeeringConnection(input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.CreateVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.CreateVpcPeeringConnectionInput, opts ...request.Option) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.CreateVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpcWithContext(ctx context.Context, input *ec2.CreateVpcInput, opts ...request.Option) (*ec2.CreateVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcOutput), nil
	}
	out, err := c.EC2API.CreateVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnConnection(input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionOutput), nil
	}
	out, err := c.EC2API.CreateVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnConnectionRoute(input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionRouteOutput), nil
	}
	out, err := c.EC2API.CreateVpnConnectionRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnConnectionRouteWithContext(ctx context.Context, input *ec2.CreateVpnConnectionRouteInput, opts ...request.Option) (*ec2.CreateVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionRouteOutput), nil
	}
	out, err := c.EC2API.CreateVpnConnectionRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnConnectionWithContext(ctx context.Context, input *ec2.CreateVpnConnectionInput, opts ...request.Option) (*ec2.CreateVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionOutput), nil
	}
	out, err := c.EC2API.CreateVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnGateway(input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnGatewayOutput), nil
	}
	out, err := c.EC2API.CreateVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) CreateVpnGatewayWithContext(ctx context.Context, input *ec2.CreateVpnGatewayInput, opts ...request.Option) (*ec2.CreateVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnGatewayOutput), nil
	}
	out, err := c.EC2API.CreateVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCarrierGateway(input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCarrierGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteCarrierGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCarrierGatewayWithContext(ctx context.Context, input *ec2.DeleteCarrierGatewayInput, opts ...request.Option) (*ec2.DeleteCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCarrierGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteCarrierGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteClientVpnEndpoint(input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.DeleteClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteClientVpnEndpointWithContext(ctx context.Context, input *ec2.DeleteClientVpnEndpointInput, opts ...request.Option) (*ec2.DeleteClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.DeleteClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteClientVpnRoute(input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnRouteOutput), nil
	}
	out, err := c.EC2API.DeleteClientVpnRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteClientVpnRouteWithContext(ctx context.Context, input *ec2.DeleteClientVpnRouteInput, opts ...request.Option) (*ec2.DeleteClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnRouteOutput), nil
	}
	out, err := c.EC2API.DeleteClientVpnRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCoipCidr(input *ec2.DeleteCoipCidrInput) (*ec2.DeleteCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipCidrOutput), nil
	}
	out, err := c.EC2API.DeleteCoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCoipCidrWithContext(ctx context.Context, input *ec2.DeleteCoipCidrInput, opts ...request.Option) (*ec2.DeleteCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipCidrOutput), nil
	}
	out, err := c.EC2API.DeleteCoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCoipPool(input *ec2.DeleteCoipPoolInput) (*ec2.DeleteCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipPoolOutput), nil
	}
	out, err := c.EC2API.DeleteCoipPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCoipPoolWithContext(ctx context.Context, input *ec2.DeleteCoipPoolInput, opts ...request.Option) (*ec2.DeleteCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipPoolOutput), nil
	}
	out, err := c.EC2API.DeleteCoipPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCustomerGateway(input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCustomerGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteCustomerGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteCustomerGatewayWithContext(ctx context.Context, input *ec2.DeleteCustomerGatewayInput, opts ...request.Option) (*ec2.DeleteCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCustomerGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteCustomerGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteDhcpOptions(input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DeleteDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteDhcpOptionsWithContext(ctx context.Context, input *ec2.DeleteDhcpOptionsInput, opts ...request.Option) (*ec2.DeleteDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DeleteDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteEgressOnlyInternetGateway(input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteEgressOnlyInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteEgressOnlyInternetGatewayWithContext(ctx context.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput, opts ...request.Option) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteEgressOnlyInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFleets(input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFleetsOutput), nil
	}
	out, err := c.EC2API.DeleteFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFleetsWithContext(ctx context.Context, input *ec2.DeleteFleetsInput, opts ...request.Option) (*ec2.DeleteFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFleetsOutput), nil
	}
	out, err := c.EC2API.DeleteFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFlowLogs(input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFlowLogsOutput), nil
	}
	out, err := c.EC2API.DeleteFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFlowLogsWithContext(ctx context.Context, input *ec2.DeleteFlowLogsInput, opts ...request.Option) (*ec2.DeleteFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFlowLogsOutput), nil
	}
	out, err := c.EC2API.DeleteFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFpgaImage(input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFpgaImageOutput), nil
	}
	out, err := c.EC2API.DeleteFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteFpgaImageWithContext(ctx context.Context, input *ec2.DeleteFpgaImageInput, opts ...request.Option) (*ec2.DeleteFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFpgaImageOutput), nil
	}
	out, err := c.EC2API.DeleteFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteInstanceEventWindow(input *ec2.DeleteInstanceEventWindowInput) (*ec2.DeleteInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.DeleteInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteInstanceEventWindowWithContext(ctx context.Context, input *ec2.DeleteInstanceEventWindowInput, opts ...request.Option) (*ec2.DeleteInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.DeleteInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteInternetGateway(input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteInternetGatewayWithContext(ctx context.Context, input *ec2.DeleteInternetGatewayInput, opts ...request.Option) (*ec2.DeleteInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpam(input *ec2.DeleteIpamInput) (*ec2.DeleteIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamOutput), nil
	}
	out, err := c.EC2API.DeleteIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpamPool(input *ec2.DeleteIpamPoolInput) (*ec2.DeleteIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamPoolOutput), nil
	}
	out, err := c.EC2API.DeleteIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpamPoolWithContext(ctx context.Context, input *ec2.DeleteIpamPoolInput, opts ...request.Option) (*ec2.DeleteIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamPoolOutput), nil
	}
	out, err := c.EC2API.DeleteIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpamScope(input *ec2.DeleteIpamScopeInput) (*ec2.DeleteIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamScopeOutput), nil
	}
	out, err := c.EC2API.DeleteIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpamScopeWithContext(ctx context.Context, input *ec2.DeleteIpamScopeInput, opts ...request.Option) (*ec2.DeleteIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamScopeOutput), nil
	}
	out, err := c.EC2API.DeleteIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteIpamWithContext(ctx context.Context, input *ec2.DeleteIpamInput, opts ...request.Option) (*ec2.DeleteIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamOutput), nil
	}
	out, err := c.EC2API.DeleteIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteKeyPairOutput), nil
	}
	out, err := c.EC2API.DeleteKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteKeyPairWithContext(ctx context.Context, input *ec2.DeleteKeyPairInput, opts ...request.Option) (*ec2.DeleteKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteKeyPairOutput), nil
	}
	out, err := c.EC2API.DeleteKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLaunchTemplate(input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.DeleteLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLaunchTemplateVersions(input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DeleteLaunchTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLaunchTemplateVersionsWithContext(ctx context.Context, input *ec2.DeleteLaunchTemplateVersionsInput, opts ...request.Option) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DeleteLaunchTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLaunchTemplateWithContext(ctx context.Context, input *ec2.DeleteLaunchTemplateInput, opts ...request.Option) (*ec2.DeleteLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.DeleteLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRoute(input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTable(input *ec2.DeleteLocalGatewayRouteTableInput) (*ec2.DeleteLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input *ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTableVpcAssociation(input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTableVpcAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTableVpcAssociationWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTableVpcAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteTableWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteLocalGatewayRouteWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.DeleteLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteManagedPrefixList(input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.DeleteManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteManagedPrefixListWithContext(ctx context.Context, input *ec2.DeleteManagedPrefixListInput, opts ...request.Option) (*ec2.DeleteManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.DeleteManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNatGateway(input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNatGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteNatGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNatGatewayWithContext(ctx context.Context, input *ec2.DeleteNatGatewayInput, opts ...request.Option) (*ec2.DeleteNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNatGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteNatGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkAcl(input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkAcl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkAclEntry(input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkAclEntryWithContext(ctx context.Context, input *ec2.DeleteNetworkAclEntryInput, opts ...request.Option) (*ec2.DeleteNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkAclWithContext(ctx context.Context, input *ec2.DeleteNetworkAclInput, opts ...request.Option) (*ec2.DeleteNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkAclWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAccessScope(input *ec2.DeleteNetworkInsightsAccessScopeInput) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAccessScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAccessScopeAnalysis(input *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAccessScopeAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAccessScopeAnalysisWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAccessScopeAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAccessScopeWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAccessScopeInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAccessScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAnalysis(input *ec2.DeleteNetworkInsightsAnalysisInput) (*ec2.DeleteNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsAnalysisWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAnalysisInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsPath(input *ec2.DeleteNetworkInsightsPathInput) (*ec2.DeleteNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsPathOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInsightsPathWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsPathInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsPathOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInsightsPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInterface(input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInterfacePermission(input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfacePermissionOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInterfacePermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInterfacePermissionWithContext(ctx context.Context, input *ec2.DeleteNetworkInterfacePermissionInput, opts ...request.Option) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfacePermissionOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInterfacePermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteNetworkInterfaceWithContext(ctx context.Context, input *ec2.DeleteNetworkInterfaceInput, opts ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.DeleteNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeletePlacementGroup(input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePlacementGroupOutput), nil
	}
	out, err := c.EC2API.DeletePlacementGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeletePlacementGroupWithContext(ctx context.Context, input *ec2.DeletePlacementGroupInput, opts ...request.Option) (*ec2.DeletePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePlacementGroupOutput), nil
	}
	out, err := c.EC2API.DeletePlacementGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeletePublicIpv4Pool(input *ec2.DeletePublicIpv4PoolInput) (*ec2.DeletePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePublicIpv4PoolOutput), nil
	}
	out, err := c.EC2API.DeletePublicIpv4Pool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeletePublicIpv4PoolWithContext(ctx context.Context, input *ec2.DeletePublicIpv4PoolInput, opts ...request.Option) (*ec2.DeletePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePublicIpv4PoolOutput), nil
	}
	out, err := c.EC2API.DeletePublicIpv4PoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteQueuedReservedInstances(input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteQueuedReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DeleteQueuedReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteQueuedReservedInstancesWithContext(ctx context.Context, input *ec2.DeleteQueuedReservedInstancesInput, opts ...request.Option) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteQueuedReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DeleteQueuedReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteRoute(input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteOutput), nil
	}
	out, err := c.EC2API.DeleteRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteRouteTable(input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteRouteTableWithContext(ctx context.Context, input *ec2.DeleteRouteTableInput, opts ...request.Option) (*ec2.DeleteRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteRouteWithContext(ctx context.Context, input *ec2.DeleteRouteInput, opts ...request.Option) (*ec2.DeleteRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteOutput), nil
	}
	out, err := c.EC2API.DeleteRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSecurityGroupOutput), nil
	}
	out, err := c.EC2API.DeleteSecurityGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSecurityGroupWithContext(ctx context.Context, input *ec2.DeleteSecurityGroupInput, opts ...request.Option) (*ec2.DeleteSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSecurityGroupOutput), nil
	}
	out, err := c.EC2API.DeleteSecurityGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSnapshot(input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSnapshotOutput), nil
	}
	out, err := c.EC2API.DeleteSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSnapshotWithContext(ctx context.Context, input *ec2.DeleteSnapshotInput, opts ...request.Option) (*ec2.DeleteSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSnapshotOutput), nil
	}
	out, err := c.EC2API.DeleteSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSpotDatafeedSubscription(input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DeleteSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DeleteSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSubnet(input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetOutput), nil
	}
	out, err := c.EC2API.DeleteSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSubnetCidrReservation(input *ec2.DeleteSubnetCidrReservationInput) (*ec2.DeleteSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetCidrReservationOutput), nil
	}
	out, err := c.EC2API.DeleteSubnetCidrReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSubnetCidrReservationWithContext(ctx context.Context, input *ec2.DeleteSubnetCidrReservationInput, opts ...request.Option) (*ec2.DeleteSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetCidrReservationOutput), nil
	}
	out, err := c.EC2API.DeleteSubnetCidrReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteSubnetWithContext(ctx context.Context, input *ec2.DeleteSubnetInput, opts ...request.Option) (*ec2.DeleteSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetOutput), nil
	}
	out, err := c.EC2API.DeleteSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTagsOutput), nil
	}
	out, err := c.EC2API.DeleteTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTagsWithContext(ctx context.Context, input *ec2.DeleteTagsInput, opts ...request.Option) (*ec2.DeleteTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTagsOutput), nil
	}
	out, err := c.EC2API.DeleteTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorFilter(input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorFilterRule(input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorFilterWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorSession(input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorSessionInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorTarget(input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorTargetOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTrafficMirrorTargetWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorTargetInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorTargetOutput), nil
	}
	out, err := c.EC2API.DeleteTrafficMirrorTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGateway(input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayConnect(input *ec2.DeleteTransitGatewayConnectInput) (*ec2.DeleteTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayConnect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayConnectPeer(input *ec2.DeleteTransitGatewayConnectPeerInput) (*ec2.DeleteTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayConnectPeer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayConnectPeerWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayConnectPeerInput, opts ...request.Option) (*ec2.DeleteTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayConnectPeerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayConnectWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayConnectInput, opts ...request.Option) (*ec2.DeleteTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayConnectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayMulticastDomain(input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPeeringAttachment(input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPolicyTable(input *ec2.DeleteTransitGatewayPolicyTableInput) (*ec2.DeleteTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPrefixListReference(input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRoute(input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRouteTable(input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRouteTableAnnouncement(input *ec2.DeleteTransitGatewayRouteTableAnnouncementInput) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRouteTableAnnouncement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRouteTableAnnouncementWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteTableAnnouncementInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRouteTableAnnouncementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayRouteWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayVpcAttachment(input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteTransitGatewayWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayInput, opts ...request.Option) (*ec2.DeleteTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessEndpoint(input *ec2.DeleteVerifiedAccessEndpointInput) (*ec2.DeleteVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessGroup(input *ec2.DeleteVerifiedAccessGroupInput) (*ec2.DeleteVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessGroupInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessInstance(input *ec2.DeleteVerifiedAccessInstanceInput) (*ec2.DeleteVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessTrustProvider(input *ec2.DeleteVerifiedAccessTrustProviderInput) (*ec2.DeleteVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.DeleteVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVolumeOutput), nil
	}
	out, err := c.EC2API.DeleteVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVolumeWithContext(ctx context.Context, input *ec2.DeleteVolumeInput, opts ...request.Option) (*ec2.DeleteVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVolumeOutput), nil
	}
	out, err := c.EC2API.DeleteVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpc(input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcOutput), nil
	}
	out, err := c.EC2API.DeleteVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpointConnectionNotifications(input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpointConnectionNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpointConnectionNotificationsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpointConnectionNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpointServiceConfigurations(input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpointServiceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpointServiceConfigurationsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpointServiceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpoints(input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcEndpointsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DeleteVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcPeeringConnection(input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.DeleteVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.DeleteVpcPeeringConnectionInput, opts ...request.Option) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.DeleteVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpcWithContext(ctx context.Context, input *ec2.DeleteVpcInput, opts ...request.Option) (*ec2.DeleteVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcOutput), nil
	}
	out, err := c.EC2API.DeleteVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnConnection(input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionOutput), nil
	}
	out, err := c.EC2API.DeleteVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnConnectionRoute(input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionRouteOutput), nil
	}
	out, err := c.EC2API.DeleteVpnConnectionRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnConnectionRouteWithContext(ctx context.Context, input *ec2.DeleteVpnConnectionRouteInput, opts ...request.Option) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionRouteOutput), nil
	}
	out, err := c.EC2API.DeleteVpnConnectionRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnConnectionWithContext(ctx context.Context, input *ec2.DeleteVpnConnectionInput, opts ...request.Option) (*ec2.DeleteVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionOutput), nil
	}
	out, err := c.EC2API.DeleteVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnGateway(input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeleteVpnGatewayWithContext(ctx context.Context, input *ec2.DeleteVpnGatewayInput, opts ...request.Option) (*ec2.DeleteVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnGatewayOutput), nil
	}
	out, err := c.EC2API.DeleteVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionByoipCidr(input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionByoipCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionByoipCidrWithContext(ctx context.Context, input *ec2.DeprovisionByoipCidrInput, opts ...request.Option) (*ec2.DeprovisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionByoipCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionIpamPoolCidr(input *ec2.DeprovisionIpamPoolCidrInput) (*ec2.DeprovisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionIpamPoolCidrWithContext(ctx context.Context, input *ec2.DeprovisionIpamPoolCidrInput, opts ...request.Option) (*ec2.DeprovisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionPublicIpv4PoolCidr(input *ec2.DeprovisionPublicIpv4PoolCidrInput) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionPublicIpv4PoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeprovisionPublicIpv4PoolCidrWithContext(ctx context.Context, input *ec2.DeprovisionPublicIpv4PoolCidrInput, opts ...request.Option) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.EC2API.DeprovisionPublicIpv4PoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterImage(input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterImageOutput), nil
	}
	out, err := c.EC2API.DeregisterImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterImageWithContext(ctx context.Context, input *ec2.DeregisterImageInput, opts ...request.Option) (*ec2.DeregisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterImageOutput), nil
	}
	out, err := c.EC2API.DeregisterImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterInstanceEventNotificationAttributes(input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DeregisterInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DeregisterInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterTransitGatewayMulticastGroupMembers(input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.EC2API.DeregisterTransitGatewayMulticastGroupMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterTransitGatewayMulticastGroupMembersWithContext(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput, opts ...request.Option) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.EC2API.DeregisterTransitGatewayMulticastGroupMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterTransitGatewayMulticastGroupSources(input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.EC2API.DeregisterTransitGatewayMulticastGroupSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DeregisterTransitGatewayMulticastGroupSourcesWithContext(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput, opts ...request.Option) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.EC2API.DeregisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAccountAttributes(input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAccountAttributesWithContext(ctx context.Context, input *ec2.DescribeAccountAttributesInput, opts ...request.Option) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressTransfers(input *ec2.DescribeAddressTransfersInput) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.EC2API.DescribeAddressTransfers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressTransfersPages(input *ec2.DescribeAddressTransfersInput, fn func(*ec2.DescribeAddressTransfersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressTransfersOutput{}
	fnCacher := func(out *ec2.DescribeAddressTransfersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressTransfersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressTransfersPagesWithContext(ctx context.Context, input *ec2.DescribeAddressTransfersInput, fn func(*ec2.DescribeAddressTransfersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressTransfersOutput{}
	fnCacher := func(out *ec2.DescribeAddressTransfersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressTransfersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressTransfersWithContext(ctx context.Context, input *ec2.DescribeAddressTransfersInput, opts ...request.Option) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.EC2API.DescribeAddressTransfersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesAttribute(input *ec2.DescribeAddressesAttributeInput) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesAttributePages(input *ec2.DescribeAddressesAttributeInput, fn func(*ec2.DescribeAddressesAttributeOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressesAttributeOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressesAttributeOutput{}
	fnCacher := func(out *ec2.DescribeAddressesAttributeOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressesAttributePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressesAttributePagesWithContext(ctx context.Context, input *ec2.DescribeAddressesAttributeInput, fn func(*ec2.DescribeAddressesAttributeOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressesAttributeOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressesAttributeOutput{}
	fnCacher := func(out *ec2.DescribeAddressesAttributeOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressesAttributePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressesAttributeWithContext(ctx context.Context, input *ec2.DescribeAddressesAttributeInput, opts ...request.Option) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesWithContext(ctx context.Context, input *ec2.DescribeAddressesInput, opts ...request.Option) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAggregateIdFormat(input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeAggregateIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAggregateIdFormatWithContext(ctx context.Context, input *ec2.DescribeAggregateIdFormatInput, opts ...request.Option) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeAggregateIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAvailabilityZones(input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.EC2API.DescribeAvailabilityZones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAvailabilityZonesWithContext(ctx context.Context, input *ec2.DescribeAvailabilityZonesInput, opts ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.EC2API.DescribeAvailabilityZonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptions(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, fn func(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput{}
	fnCacher := func(out *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsPagesWithContext(ctx context.Context, input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, fn func(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput{}
	fnCacher := func(out *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx context.Context, input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, opts ...request.Option) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeBundleTasks(input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.EC2API.DescribeBundleTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeBundleTasksWithContext(ctx context.Context, input *ec2.DescribeBundleTasksInput, opts ...request.Option) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.EC2API.DescribeBundleTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeByoipCidrs(input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.EC2API.DescribeByoipCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeByoipCidrsPages(input *ec2.DescribeByoipCidrsInput, fn func(*ec2.DescribeByoipCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeByoipCidrsOutput{}
	fnCacher := func(out *ec2.DescribeByoipCidrsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeByoipCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeByoipCidrsPagesWithContext(ctx context.Context, input *ec2.DescribeByoipCidrsInput, fn func(*ec2.DescribeByoipCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeByoipCidrsOutput{}
	fnCacher := func(out *ec2.DescribeByoipCidrsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeByoipCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeByoipCidrsWithContext(ctx context.Context, input *ec2.DescribeByoipCidrsInput, opts ...request.Option) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.EC2API.DescribeByoipCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationFleets(input *ec2.DescribeCapacityReservationFleetsInput) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationFleetsPages(input *ec2.DescribeCapacityReservationFleetsInput, fn func(*ec2.DescribeCapacityReservationFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationFleetsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationFleetsPagesWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationFleetsInput, fn func(*ec2.DescribeCapacityReservationFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationFleetsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationFleetsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationFleetsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservations(input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationsPages(input *ec2.DescribeCapacityReservationsInput, fn func(*ec2.DescribeCapacityReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationsPagesWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationsInput, fn func(*ec2.DescribeCapacityReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCarrierGateways(input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCarrierGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCarrierGatewaysPages(input *ec2.DescribeCarrierGatewaysInput, fn func(*ec2.DescribeCarrierGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCarrierGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeCarrierGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCarrierGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCarrierGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput, fn func(*ec2.DescribeCarrierGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCarrierGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeCarrierGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCarrierGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCarrierGatewaysWithContext(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput, opts ...request.Option) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCarrierGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClassicLinkInstances(input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeClassicLinkInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClassicLinkInstancesPages(input *ec2.DescribeClassicLinkInstancesInput, fn func(*ec2.DescribeClassicLinkInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClassicLinkInstancesOutput{}
	fnCacher := func(out *ec2.DescribeClassicLinkInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClassicLinkInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClassicLinkInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput, fn func(*ec2.DescribeClassicLinkInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClassicLinkInstancesOutput{}
	fnCacher := func(out *ec2.DescribeClassicLinkInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClassicLinkInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClassicLinkInstancesWithContext(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput, opts ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeClassicLinkInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnAuthorizationRules(input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnAuthorizationRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnAuthorizationRulesPages(input *ec2.DescribeClientVpnAuthorizationRulesInput, fn func(*ec2.DescribeClientVpnAuthorizationRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnAuthorizationRulesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnAuthorizationRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnAuthorizationRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnAuthorizationRulesPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput, fn func(*ec2.DescribeClientVpnAuthorizationRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnAuthorizationRulesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnAuthorizationRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnAuthorizationRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnAuthorizationRulesWithContext(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput, opts ...request.Option) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnAuthorizationRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnConnections(input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnConnectionsPages(input *ec2.DescribeClientVpnConnectionsInput, fn func(*ec2.DescribeClientVpnConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput, fn func(*ec2.DescribeClientVpnConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnEndpoints(input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnEndpointsPages(input *ec2.DescribeClientVpnEndpointsInput, fn func(*ec2.DescribeClientVpnEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput, fn func(*ec2.DescribeClientVpnEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnEndpointsWithContext(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput, opts ...request.Option) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnRoutes(input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnRoutesPages(input *ec2.DescribeClientVpnRoutesInput, fn func(*ec2.DescribeClientVpnRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnRoutesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnRoutesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnRoutesPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput, fn func(*ec2.DescribeClientVpnRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnRoutesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnRoutesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnRoutesWithContext(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput, opts ...request.Option) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnTargetNetworks(input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnTargetNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnTargetNetworksPages(input *ec2.DescribeClientVpnTargetNetworksInput, fn func(*ec2.DescribeClientVpnTargetNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnTargetNetworksOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnTargetNetworksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnTargetNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnTargetNetworksPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput, fn func(*ec2.DescribeClientVpnTargetNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnTargetNetworksOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnTargetNetworksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnTargetNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnTargetNetworksWithContext(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput, opts ...request.Option) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnTargetNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCoipPools(input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeCoipPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCoipPoolsPages(input *ec2.DescribeCoipPoolsInput, fn func(*ec2.DescribeCoipPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCoipPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCoipPoolsOutput{}
	fnCacher := func(out *ec2.DescribeCoipPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCoipPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCoipPoolsPagesWithContext(ctx context.Context, input *ec2.DescribeCoipPoolsInput, fn func(*ec2.DescribeCoipPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCoipPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCoipPoolsOutput{}
	fnCacher := func(out *ec2.DescribeCoipPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCoipPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCoipPoolsWithContext(ctx context.Context, input *ec2.DescribeCoipPoolsInput, opts ...request.Option) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeCoipPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeConversionTasks(input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.EC2API.DescribeConversionTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeConversionTasksWithContext(ctx context.Context, input *ec2.DescribeConversionTasksInput, opts ...request.Option) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.EC2API.DescribeConversionTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCustomerGateways(input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCustomerGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCustomerGatewaysWithContext(ctx context.Context, input *ec2.DescribeCustomerGatewaysInput, opts ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCustomerGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeDhcpOptions(input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DescribeDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeDhcpOptionsPages(input *ec2.DescribeDhcpOptionsInput, fn func(*ec2.DescribeDhcpOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeDhcpOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeDhcpOptionsOutput{}
	fnCacher := func(out *ec2.DescribeDhcpOptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeDhcpOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeDhcpOptionsPagesWithContext(ctx context.Context, input *ec2.DescribeDhcpOptionsInput, fn func(*ec2.DescribeDhcpOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeDhcpOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeDhcpOptionsOutput{}
	fnCacher := func(out *ec2.DescribeDhcpOptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeDhcpOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeDhcpOptionsWithContext(ctx context.Context, input *ec2.DescribeDhcpOptionsInput, opts ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DescribeDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeEgressOnlyInternetGateways(input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeEgressOnlyInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysPages(input *ec2.DescribeEgressOnlyInternetGatewaysInput, fn func(*ec2.DescribeEgressOnlyInternetGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeEgressOnlyInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeEgressOnlyInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput, fn func(*ec2.DescribeEgressOnlyInternetGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeEgressOnlyInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeEgressOnlyInternetGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeEgressOnlyInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeElasticGpus(input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.EC2API.DescribeElasticGpus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeElasticGpusWithContext(ctx context.Context, input *ec2.DescribeElasticGpusInput, opts ...request.Option) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.EC2API.DescribeElasticGpusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportImageTasks(input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportImageTasksPages(input *ec2.DescribeExportImageTasksInput, fn func(*ec2.DescribeExportImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeExportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeExportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeExportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeExportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeExportImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeExportImageTasksInput, fn func(*ec2.DescribeExportImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeExportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeExportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeExportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeExportImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeExportImageTasksWithContext(ctx context.Context, input *ec2.DescribeExportImageTasksInput, opts ...request.Option) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportTasks(input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportTasksWithContext(ctx context.Context, input *ec2.DescribeExportTasksInput, opts ...request.Option) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastLaunchImages(input *ec2.DescribeFastLaunchImagesInput) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFastLaunchImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastLaunchImagesPages(input *ec2.DescribeFastLaunchImagesInput, fn func(*ec2.DescribeFastLaunchImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastLaunchImagesOutput{}
	fnCacher := func(out *ec2.DescribeFastLaunchImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastLaunchImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastLaunchImagesPagesWithContext(ctx context.Context, input *ec2.DescribeFastLaunchImagesInput, fn func(*ec2.DescribeFastLaunchImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastLaunchImagesOutput{}
	fnCacher := func(out *ec2.DescribeFastLaunchImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastLaunchImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastLaunchImagesWithContext(ctx context.Context, input *ec2.DescribeFastLaunchImagesInput, opts ...request.Option) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFastLaunchImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastSnapshotRestores(input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DescribeFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastSnapshotRestoresPages(input *ec2.DescribeFastSnapshotRestoresInput, fn func(*ec2.DescribeFastSnapshotRestoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastSnapshotRestoresOutput{}
	fnCacher := func(out *ec2.DescribeFastSnapshotRestoresOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastSnapshotRestoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastSnapshotRestoresPagesWithContext(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput, fn func(*ec2.DescribeFastSnapshotRestoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastSnapshotRestoresOutput{}
	fnCacher := func(out *ec2.DescribeFastSnapshotRestoresOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastSnapshotRestoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput, opts ...request.Option) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DescribeFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetHistory(input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeFleetHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetHistoryWithContext(ctx context.Context, input *ec2.DescribeFleetHistoryInput, opts ...request.Option) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeFleetHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetInstances(input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeFleetInstancesInput, opts ...request.Option) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleets(input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetsPages(input *ec2.DescribeFleetsInput, fn func(*ec2.DescribeFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFleetsOutput{}
	fnCacher := func(out *ec2.DescribeFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFleetsPagesWithContext(ctx context.Context, input *ec2.DescribeFleetsInput, fn func(*ec2.DescribeFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFleetsOutput{}
	fnCacher := func(out *ec2.DescribeFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFleetsWithContext(ctx context.Context, input *ec2.DescribeFleetsInput, opts ...request.Option) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFlowLogs(input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.EC2API.DescribeFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFlowLogsPages(input *ec2.DescribeFlowLogsInput, fn func(*ec2.DescribeFlowLogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFlowLogsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFlowLogsOutput{}
	fnCacher := func(out *ec2.DescribeFlowLogsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFlowLogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFlowLogsPagesWithContext(ctx context.Context, input *ec2.DescribeFlowLogsInput, fn func(*ec2.DescribeFlowLogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFlowLogsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFlowLogsOutput{}
	fnCacher := func(out *ec2.DescribeFlowLogsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFlowLogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFlowLogsWithContext(ctx context.Context, input *ec2.DescribeFlowLogsInput, opts ...request.Option) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.EC2API.DescribeFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImageAttribute(input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImageAttributeWithContext(ctx context.Context, input *ec2.DescribeFpgaImageAttributeInput, opts ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImages(input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImagesPages(input *ec2.DescribeFpgaImagesInput, fn func(*ec2.DescribeFpgaImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFpgaImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFpgaImagesOutput{}
	fnCacher := func(out *ec2.DescribeFpgaImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFpgaImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFpgaImagesPagesWithContext(ctx context.Context, input *ec2.DescribeFpgaImagesInput, fn func(*ec2.DescribeFpgaImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFpgaImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFpgaImagesOutput{}
	fnCacher := func(out *ec2.DescribeFpgaImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFpgaImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFpgaImagesWithContext(ctx context.Context, input *ec2.DescribeFpgaImagesInput, opts ...request.Option) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationOfferings(input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationOfferingsPages(input *ec2.DescribeHostReservationOfferingsInput, fn func(*ec2.DescribeHostReservationOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput, fn func(*ec2.DescribeHostReservationOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationOfferingsWithContext(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput, opts ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservations(input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationsPages(input *ec2.DescribeHostReservationsInput, fn func(*ec2.DescribeHostReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationsPagesWithContext(ctx context.Context, input *ec2.DescribeHostReservationsInput, fn func(*ec2.DescribeHostReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationsWithContext(ctx context.Context, input *ec2.DescribeHostReservationsInput, opts ...request.Option) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHosts(input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.EC2API.DescribeHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostsPages(input *ec2.DescribeHostsInput, fn func(*ec2.DescribeHostsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostsOutput{}
	fnCacher := func(out *ec2.DescribeHostsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostsPagesWithContext(ctx context.Context, input *ec2.DescribeHostsInput, fn func(*ec2.DescribeHostsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostsOutput{}
	fnCacher := func(out *ec2.DescribeHostsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostsWithContext(ctx context.Context, input *ec2.DescribeHostsInput, opts ...request.Option) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.EC2API.DescribeHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIamInstanceProfileAssociations(input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeIamInstanceProfileAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIamInstanceProfileAssociationsPages(input *ec2.DescribeIamInstanceProfileAssociationsInput, fn func(*ec2.DescribeIamInstanceProfileAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIamInstanceProfileAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeIamInstanceProfileAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIamInstanceProfileAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIamInstanceProfileAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput, fn func(*ec2.DescribeIamInstanceProfileAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIamInstanceProfileAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeIamInstanceProfileAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIamInstanceProfileAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIamInstanceProfileAssociationsWithContext(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput, opts ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeIamInstanceProfileAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdFormat(input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdFormatInput, opts ...request.Option) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdentityIdFormat(input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdentityIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdentityIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdentityIdFormatInput, opts ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdentityIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImageAttribute(input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImageAttributeWithContext(ctx context.Context, input *ec2.DescribeImageAttributeInput, opts ...request.Option) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.EC2API.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImagesPages(input *ec2.DescribeImagesInput, fn func(*ec2.DescribeImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImagesOutput{}
	fnCacher := func(out *ec2.DescribeImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImagesPagesWithContext(ctx context.Context, input *ec2.DescribeImagesInput, fn func(*ec2.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImagesOutput{}
	fnCacher := func(out *ec2.DescribeImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImagesWithContext(ctx context.Context, input *ec2.DescribeImagesInput, opts ...request.Option) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.EC2API.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportImageTasks(input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportImageTasksPages(input *ec2.DescribeImportImageTasksInput, fn func(*ec2.DescribeImportImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeImportImageTasksInput, fn func(*ec2.DescribeImportImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportImageTasksWithContext(ctx context.Context, input *ec2.DescribeImportImageTasksInput, opts ...request.Option) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportSnapshotTasks(input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportSnapshotTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportSnapshotTasksPages(input *ec2.DescribeImportSnapshotTasksInput, fn func(*ec2.DescribeImportSnapshotTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportSnapshotTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportSnapshotTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportSnapshotTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportSnapshotTasksPagesWithContext(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput, fn func(*ec2.DescribeImportSnapshotTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportSnapshotTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportSnapshotTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportSnapshotTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportSnapshotTasksWithContext(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput, opts ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportSnapshotTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceAttribute(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceAttributeWithContext(ctx context.Context, input *ec2.DescribeInstanceAttributeInput, opts ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceCreditSpecifications(input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceCreditSpecifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceCreditSpecificationsPages(input *ec2.DescribeInstanceCreditSpecificationsInput, fn func(*ec2.DescribeInstanceCreditSpecificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceCreditSpecificationsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceCreditSpecificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceCreditSpecificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceCreditSpecificationsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput, fn func(*ec2.DescribeInstanceCreditSpecificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceCreditSpecificationsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceCreditSpecificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceCreditSpecificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceCreditSpecificationsWithContext(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput, opts ...request.Option) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceCreditSpecificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventNotificationAttributes(input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventWindows(input *ec2.DescribeInstanceEventWindowsInput) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventWindows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventWindowsPages(input *ec2.DescribeInstanceEventWindowsInput, fn func(*ec2.DescribeInstanceEventWindowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceEventWindowsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceEventWindowsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceEventWindowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceEventWindowsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceEventWindowsInput, fn func(*ec2.DescribeInstanceEventWindowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceEventWindowsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceEventWindowsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceEventWindowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceEventWindowsWithContext(ctx context.Context, input *ec2.DescribeInstanceEventWindowsInput, opts ...request.Option) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventWindowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceStatusPages(input *ec2.DescribeInstanceStatusInput, fn func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceStatusOutput{}
	fnCacher := func(out *ec2.DescribeInstanceStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceStatusPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceStatusInput, fn func(*ec2.DescribeInstanceStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceStatusOutput{}
	fnCacher := func(out *ec2.DescribeInstanceStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceStatusWithContext(ctx context.Context, input *ec2.DescribeInstanceStatusInput, opts ...request.Option) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypeOfferings(input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypeOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypeOfferingsPages(input *ec2.DescribeInstanceTypeOfferingsInput, fn func(*ec2.DescribeInstanceTypeOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypeOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypeOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypeOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypeOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput, fn func(*ec2.DescribeInstanceTypeOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypeOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypeOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypeOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypeOfferingsWithContext(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput, opts ...request.Option) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypeOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypes(input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypesPages(input *ec2.DescribeInstanceTypesInput, fn func(*ec2.DescribeInstanceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypesOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypesPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypesInput, fn func(*ec2.DescribeInstanceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypesOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypesInput, opts ...request.Option) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstancesPages(input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstancesOutput{}
	fnCacher := func(out *ec2.DescribeInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstancesOutput{}
	fnCacher := func(out *ec2.DescribeInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstancesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInternetGateways(input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInternetGatewaysPages(input *ec2.DescribeInternetGatewaysInput, fn func(*ec2.DescribeInternetGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInternetGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeInternetGatewaysInput, fn func(*ec2.DescribeInternetGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInternetGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamPools(input *ec2.DescribeIpamPoolsInput) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamPoolsPages(input *ec2.DescribeIpamPoolsInput, fn func(*ec2.DescribeIpamPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamPoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpamPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamPoolsPagesWithContext(ctx context.Context, input *ec2.DescribeIpamPoolsInput, fn func(*ec2.DescribeIpamPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamPoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpamPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamPoolsWithContext(ctx context.Context, input *ec2.DescribeIpamPoolsInput, opts ...request.Option) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamScopes(input *ec2.DescribeIpamScopesInput) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.EC2API.DescribeIpamScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamScopesPages(input *ec2.DescribeIpamScopesInput, fn func(*ec2.DescribeIpamScopesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamScopesOutput{}
	fnCacher := func(out *ec2.DescribeIpamScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamScopesPagesWithContext(ctx context.Context, input *ec2.DescribeIpamScopesInput, fn func(*ec2.DescribeIpamScopesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamScopesOutput{}
	fnCacher := func(out *ec2.DescribeIpamScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamScopesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamScopesWithContext(ctx context.Context, input *ec2.DescribeIpamScopesInput, opts ...request.Option) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.EC2API.DescribeIpamScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpams(input *ec2.DescribeIpamsInput) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.EC2API.DescribeIpams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamsPages(input *ec2.DescribeIpamsInput, fn func(*ec2.DescribeIpamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamsOutput{}
	fnCacher := func(out *ec2.DescribeIpamsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamsPagesWithContext(ctx context.Context, input *ec2.DescribeIpamsInput, fn func(*ec2.DescribeIpamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamsOutput{}
	fnCacher := func(out *ec2.DescribeIpamsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamsWithContext(ctx context.Context, input *ec2.DescribeIpamsInput, opts ...request.Option) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpv6Pools(input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpv6Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpv6PoolsPages(input *ec2.DescribeIpv6PoolsInput, fn func(*ec2.DescribeIpv6PoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpv6PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpv6PoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpv6PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpv6PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpv6PoolsPagesWithContext(ctx context.Context, input *ec2.DescribeIpv6PoolsInput, fn func(*ec2.DescribeIpv6PoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpv6PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpv6PoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpv6PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpv6PoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpv6PoolsWithContext(ctx context.Context, input *ec2.DescribeIpv6PoolsInput, opts ...request.Option) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpv6PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.EC2API.DescribeKeyPairs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeKeyPairsWithContext(ctx context.Context, input *ec2.DescribeKeyPairsInput, opts ...request.Option) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.EC2API.DescribeKeyPairsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplateVersions(input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplateVersionsPages(input *ec2.DescribeLaunchTemplateVersionsInput, fn func(*ec2.DescribeLaunchTemplateVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplateVersionsOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplateVersionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplateVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplateVersionsPagesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput, fn func(*ec2.DescribeLaunchTemplateVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplateVersionsOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplateVersionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplateVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplateVersionsWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput, opts ...request.Option) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplates(input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplatesPages(input *ec2.DescribeLaunchTemplatesInput, fn func(*ec2.DescribeLaunchTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplatesOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplatesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplatesPagesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput, fn func(*ec2.DescribeLaunchTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplatesOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplatesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplatesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput, opts ...request.Option) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociations(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTables(input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTablesPages(input *ec2.DescribeLocalGatewayRouteTablesInput, fn func(*ec2.DescribeLocalGatewayRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput, fn func(*ec2.DescribeLocalGatewayRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroups(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsPages(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaces(input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesPages(input *ec2.DescribeLocalGatewayVirtualInterfacesInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGateways(input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewaysPages(input *ec2.DescribeLocalGatewaysInput, fn func(*ec2.DescribeLocalGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewaysInput, fn func(*ec2.DescribeLocalGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewaysWithContext(ctx context.Context, input *ec2.DescribeLocalGatewaysInput, opts ...request.Option) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeManagedPrefixLists(input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribeManagedPrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeManagedPrefixListsPages(input *ec2.DescribeManagedPrefixListsInput, fn func(*ec2.DescribeManagedPrefixListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeManagedPrefixListsOutput{}
	fnCacher := func(out *ec2.DescribeManagedPrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeManagedPrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeManagedPrefixListsPagesWithContext(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput, fn func(*ec2.DescribeManagedPrefixListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeManagedPrefixListsOutput{}
	fnCacher := func(out *ec2.DescribeManagedPrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeManagedPrefixListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeManagedPrefixListsWithContext(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput, opts ...request.Option) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribeManagedPrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeMovingAddresses(input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeMovingAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeMovingAddressesPages(input *ec2.DescribeMovingAddressesInput, fn func(*ec2.DescribeMovingAddressesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeMovingAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeMovingAddressesOutput{}
	fnCacher := func(out *ec2.DescribeMovingAddressesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeMovingAddressesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeMovingAddressesPagesWithContext(ctx context.Context, input *ec2.DescribeMovingAddressesInput, fn func(*ec2.DescribeMovingAddressesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeMovingAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeMovingAddressesOutput{}
	fnCacher := func(out *ec2.DescribeMovingAddressesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeMovingAddressesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeMovingAddressesWithContext(ctx context.Context, input *ec2.DescribeMovingAddressesInput, opts ...request.Option) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeMovingAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNatGateways(input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeNatGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNatGatewaysPages(input *ec2.DescribeNatGatewaysInput, fn func(*ec2.DescribeNatGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNatGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNatGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeNatGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNatGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNatGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeNatGatewaysInput, fn func(*ec2.DescribeNatGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNatGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNatGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeNatGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNatGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNatGatewaysWithContext(ctx context.Context, input *ec2.DescribeNatGatewaysInput, opts ...request.Option) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeNatGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkAcls(input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkAcls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkAclsPages(input *ec2.DescribeNetworkAclsInput, fn func(*ec2.DescribeNetworkAclsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkAclsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkAclsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkAclsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkAclsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkAclsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkAclsInput, fn func(*ec2.DescribeNetworkAclsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkAclsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkAclsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkAclsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkAclsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkAclsWithContext(ctx context.Context, input *ec2.DescribeNetworkAclsInput, opts ...request.Option) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkAclsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalyses(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesPages(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopes(input *ec2.DescribeNetworkInsightsAccessScopesInput) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopesPages(input *ec2.DescribeNetworkInsightsAccessScopesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAnalyses(input *ec2.DescribeNetworkInsightsAnalysesInput) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAnalysesPages(input *ec2.DescribeNetworkInsightsAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAnalysesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsPaths(input *ec2.DescribeNetworkInsightsPathsInput) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsPaths(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsPathsPages(input *ec2.DescribeNetworkInsightsPathsInput, fn func(*ec2.DescribeNetworkInsightsPathsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsPathsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsPathsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsPathsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsPathsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsPathsInput, fn func(*ec2.DescribeNetworkInsightsPathsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsPathsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsPathsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsPathsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsPathsWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsPathsInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsPathsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaceAttribute(input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacePermissions(input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacePermissionsPages(input *ec2.DescribeNetworkInterfacePermissionsInput, fn func(*ec2.DescribeNetworkInterfacePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacePermissionsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput, fn func(*ec2.DescribeNetworkInterfacePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacePermissionsWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacesPages(input *ec2.DescribeNetworkInterfacesInput, fn func(*ec2.DescribeNetworkInterfacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput, fn func(*ec2.DescribeNetworkInterfacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePlacementGroups(input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.EC2API.DescribePlacementGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePlacementGroupsWithContext(ctx context.Context, input *ec2.DescribePlacementGroupsInput, opts ...request.Option) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.EC2API.DescribePlacementGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrefixLists(input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribePrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrefixListsPages(input *ec2.DescribePrefixListsInput, fn func(*ec2.DescribePrefixListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrefixListsOutput{}
	fnCacher := func(out *ec2.DescribePrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrefixListsPagesWithContext(ctx context.Context, input *ec2.DescribePrefixListsInput, fn func(*ec2.DescribePrefixListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrefixListsOutput{}
	fnCacher := func(out *ec2.DescribePrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrefixListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrefixListsWithContext(ctx context.Context, input *ec2.DescribePrefixListsInput, opts ...request.Option) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribePrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrincipalIdFormat(input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribePrincipalIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrincipalIdFormatPages(input *ec2.DescribePrincipalIdFormatInput, fn func(*ec2.DescribePrincipalIdFormatOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrincipalIdFormatOutput{}
	fnCacher := func(out *ec2.DescribePrincipalIdFormatOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrincipalIdFormatPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrincipalIdFormatPagesWithContext(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput, fn func(*ec2.DescribePrincipalIdFormatOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrincipalIdFormatOutput{}
	fnCacher := func(out *ec2.DescribePrincipalIdFormatOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrincipalIdFormatPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrincipalIdFormatWithContext(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput, opts ...request.Option) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribePrincipalIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePublicIpv4Pools(input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.EC2API.DescribePublicIpv4Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePublicIpv4PoolsPages(input *ec2.DescribePublicIpv4PoolsInput, fn func(*ec2.DescribePublicIpv4PoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePublicIpv4PoolsOutput{}
	fnCacher := func(out *ec2.DescribePublicIpv4PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePublicIpv4PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePublicIpv4PoolsPagesWithContext(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput, fn func(*ec2.DescribePublicIpv4PoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePublicIpv4PoolsOutput{}
	fnCacher := func(out *ec2.DescribePublicIpv4PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePublicIpv4PoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePublicIpv4PoolsWithContext(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput, opts ...request.Option) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.EC2API.DescribePublicIpv4PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRegions(input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.EC2API.DescribeRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRegionsWithContext(ctx context.Context, input *ec2.DescribeRegionsInput, opts ...request.Option) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.EC2API.DescribeRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReplaceRootVolumeTasks(input *ec2.DescribeReplaceRootVolumeTasksInput) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.EC2API.DescribeReplaceRootVolumeTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReplaceRootVolumeTasksPages(input *ec2.DescribeReplaceRootVolumeTasksInput, fn func(*ec2.DescribeReplaceRootVolumeTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReplaceRootVolumeTasksOutput{}
	fnCacher := func(out *ec2.DescribeReplaceRootVolumeTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReplaceRootVolumeTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReplaceRootVolumeTasksPagesWithContext(ctx context.Context, input *ec2.DescribeReplaceRootVolumeTasksInput, fn func(*ec2.DescribeReplaceRootVolumeTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReplaceRootVolumeTasksOutput{}
	fnCacher := func(out *ec2.DescribeReplaceRootVolumeTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReplaceRootVolumeTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReplaceRootVolumeTasksWithContext(ctx context.Context, input *ec2.DescribeReplaceRootVolumeTasksInput, opts ...request.Option) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.EC2API.DescribeReplaceRootVolumeTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstances(input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesListings(input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesListings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesListingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesListingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesListingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesModifications(input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesModificationsPages(input *ec2.DescribeReservedInstancesModificationsInput, fn func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesModificationsPagesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput, fn func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesModificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesModificationsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesOfferings(input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesOfferingsPages(input *ec2.DescribeReservedInstancesOfferingsInput, fn func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput, fn func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesOfferingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRouteTables(input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRouteTablesPages(input *ec2.DescribeRouteTablesInput, fn func(*ec2.DescribeRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeRouteTablesInput, fn func(*ec2.DescribeRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeRouteTablesWithContext(ctx context.Context, input *ec2.DescribeRouteTablesInput, opts ...request.Option) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstanceAvailability(input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstanceAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstanceAvailabilityPages(input *ec2.DescribeScheduledInstanceAvailabilityInput, fn func(*ec2.DescribeScheduledInstanceAvailabilityOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstanceAvailabilityOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstanceAvailabilityOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstanceAvailabilityPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstanceAvailabilityPagesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput, fn func(*ec2.DescribeScheduledInstanceAvailabilityOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstanceAvailabilityOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstanceAvailabilityOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstanceAvailabilityPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstanceAvailabilityWithContext(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput, opts ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstanceAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstances(input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstancesPages(input *ec2.DescribeScheduledInstancesInput, fn func(*ec2.DescribeScheduledInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstancesOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstancesInput, fn func(*ec2.DescribeScheduledInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstancesOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstancesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstancesInput, opts ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupReferences(input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupReferencesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupReferencesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupRules(input *ec2.DescribeSecurityGroupRulesInput) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupRulesPages(input *ec2.DescribeSecurityGroupRulesInput, fn func(*ec2.DescribeSecurityGroupRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupRulesOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupRulesPagesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupRulesInput, fn func(*ec2.DescribeSecurityGroupRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupRulesOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupRulesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupRulesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupsPages(input *ec2.DescribeSecurityGroupsInput, fn func(*ec2.DescribeSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupsInput, fn func(*ec2.DescribeSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotAttribute(input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotAttributeWithContext(ctx context.Context, input *ec2.DescribeSnapshotAttributeInput, opts ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotTierStatus(input *ec2.DescribeSnapshotTierStatusInput) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotTierStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotTierStatusPages(input *ec2.DescribeSnapshotTierStatusInput, fn func(*ec2.DescribeSnapshotTierStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotTierStatusOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotTierStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotTierStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotTierStatusPagesWithContext(ctx context.Context, input *ec2.DescribeSnapshotTierStatusInput, fn func(*ec2.DescribeSnapshotTierStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotTierStatusOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotTierStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotTierStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotTierStatusWithContext(ctx context.Context, input *ec2.DescribeSnapshotTierStatusInput, opts ...request.Option) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotTierStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshots(input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotsPages(input *ec2.DescribeSnapshotsInput, fn func(*ec2.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotsOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotsPagesWithContext(ctx context.Context, input *ec2.DescribeSnapshotsInput, fn func(*ec2.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotsOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotsWithContext(ctx context.Context, input *ec2.DescribeSnapshotsInput, opts ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotDatafeedSubscription(input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DescribeSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DescribeSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetInstances(input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeSpotFleetInstancesInput, opts ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestHistory(input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestHistoryInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequests(input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestsPages(input *ec2.DescribeSpotFleetRequestsInput, fn func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotFleetRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotFleetRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotFleetRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotFleetRequestsPagesWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput, fn func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotFleetRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotFleetRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotFleetRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotFleetRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotInstanceRequests(input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotInstanceRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotInstanceRequestsPages(input *ec2.DescribeSpotInstanceRequestsInput, fn func(*ec2.DescribeSpotInstanceRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotInstanceRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotInstanceRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotInstanceRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotInstanceRequestsPagesWithContext(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput, fn func(*ec2.DescribeSpotInstanceRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotInstanceRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotInstanceRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotInstanceRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotInstanceRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput, opts ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotInstanceRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotPriceHistory(input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotPriceHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotPriceHistoryPages(input *ec2.DescribeSpotPriceHistoryInput, fn func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotPriceHistoryOutput{}
	fnCacher := func(out *ec2.DescribeSpotPriceHistoryOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotPriceHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotPriceHistoryPagesWithContext(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput, fn func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotPriceHistoryOutput{}
	fnCacher := func(out *ec2.DescribeSpotPriceHistoryOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotPriceHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotPriceHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput, opts ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotPriceHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStaleSecurityGroups(input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeStaleSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStaleSecurityGroupsPages(input *ec2.DescribeStaleSecurityGroupsInput, fn func(*ec2.DescribeStaleSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStaleSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeStaleSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStaleSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStaleSecurityGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput, fn func(*ec2.DescribeStaleSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStaleSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeStaleSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStaleSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStaleSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeStaleSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStoreImageTasks(input *ec2.DescribeStoreImageTasksInput) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeStoreImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStoreImageTasksPages(input *ec2.DescribeStoreImageTasksInput, fn func(*ec2.DescribeStoreImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStoreImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStoreImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeStoreImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStoreImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStoreImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeStoreImageTasksInput, fn func(*ec2.DescribeStoreImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStoreImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStoreImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeStoreImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStoreImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStoreImageTasksWithContext(ctx context.Context, input *ec2.DescribeStoreImageTasksInput, opts ...request.Option) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeStoreImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.EC2API.DescribeSubnets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSubnetsPages(input *ec2.DescribeSubnetsInput, fn func(*ec2.DescribeSubnetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSubnetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSubnetsOutput{}
	fnCacher := func(out *ec2.DescribeSubnetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSubnetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSubnetsPagesWithContext(ctx context.Context, input *ec2.DescribeSubnetsInput, fn func(*ec2.DescribeSubnetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSubnetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSubnetsOutput{}
	fnCacher := func(out *ec2.DescribeSubnetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSubnetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSubnetsWithContext(ctx context.Context, input *ec2.DescribeSubnetsInput, opts ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.EC2API.DescribeSubnetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.EC2API.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTagsPages(input *ec2.DescribeTagsInput, fn func(*ec2.DescribeTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTagsOutput{}
	fnCacher := func(out *ec2.DescribeTagsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTagsPagesWithContext(ctx context.Context, input *ec2.DescribeTagsInput, fn func(*ec2.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTagsOutput{}
	fnCacher := func(out *ec2.DescribeTagsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTagsWithContext(ctx context.Context, input *ec2.DescribeTagsInput, opts ...request.Option) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.EC2API.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorFilters(input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorFiltersPages(input *ec2.DescribeTrafficMirrorFiltersInput, fn func(*ec2.DescribeTrafficMirrorFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorFiltersOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorFiltersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorFiltersPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput, fn func(*ec2.DescribeTrafficMirrorFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorFiltersOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorFiltersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorFiltersWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorSessions(input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorSessionsPages(input *ec2.DescribeTrafficMirrorSessionsInput, fn func(*ec2.DescribeTrafficMirrorSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorSessionsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorSessionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorSessionsPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput, fn func(*ec2.DescribeTrafficMirrorSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorSessionsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorSessionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorSessionsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorTargets(input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorTargetsPages(input *ec2.DescribeTrafficMirrorTargetsInput, fn func(*ec2.DescribeTrafficMirrorTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorTargetsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorTargetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorTargetsPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput, fn func(*ec2.DescribeTrafficMirrorTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorTargetsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorTargetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorTargetsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayAttachments(input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayAttachmentsPages(input *ec2.DescribeTransitGatewayAttachmentsInput, fn func(*ec2.DescribeTransitGatewayAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput, fn func(*ec2.DescribeTransitGatewayAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectPeers(input *ec2.DescribeTransitGatewayConnectPeersInput) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectPeers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectPeersPages(input *ec2.DescribeTransitGatewayConnectPeersInput, fn func(*ec2.DescribeTransitGatewayConnectPeersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectPeersOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectPeersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectPeersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectPeersPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectPeersInput, fn func(*ec2.DescribeTransitGatewayConnectPeersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectPeersOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectPeersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectPeersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectPeersWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectPeersInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectPeersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnects(input *ec2.DescribeTransitGatewayConnectsInput) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectsPages(input *ec2.DescribeTransitGatewayConnectsInput, fn func(*ec2.DescribeTransitGatewayConnectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectsInput, fn func(*ec2.DescribeTransitGatewayConnectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayMulticastDomains(input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayMulticastDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsPages(input *ec2.DescribeTransitGatewayMulticastDomainsInput, fn func(*ec2.DescribeTransitGatewayMulticastDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayMulticastDomainsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayMulticastDomainsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayMulticastDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput, fn func(*ec2.DescribeTransitGatewayMulticastDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayMulticastDomainsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayMulticastDomainsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayMulticastDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayMulticastDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPeeringAttachments(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPeeringAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsPages(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, fn func(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPeeringAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPeeringAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, fn func(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPeeringAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPeeringAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPolicyTables(input *ec2.DescribeTransitGatewayPolicyTablesInput) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPolicyTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPolicyTablesPages(input *ec2.DescribeTransitGatewayPolicyTablesInput, fn func(*ec2.DescribeTransitGatewayPolicyTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPolicyTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPolicyTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPolicyTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPolicyTablesPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPolicyTablesInput, fn func(*ec2.DescribeTransitGatewayPolicyTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPolicyTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPolicyTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPolicyTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPolicyTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPolicyTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPolicyTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncements(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsPages(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, fn func(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, fn func(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTables(input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTablesPages(input *ec2.DescribeTransitGatewayRouteTablesInput, fn func(*ec2.DescribeTransitGatewayRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput, fn func(*ec2.DescribeTransitGatewayRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayVpcAttachments(input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayVpcAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsPages(input *ec2.DescribeTransitGatewayVpcAttachmentsInput, fn func(*ec2.DescribeTransitGatewayVpcAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayVpcAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayVpcAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayVpcAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput, fn func(*ec2.DescribeTransitGatewayVpcAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayVpcAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayVpcAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayVpcAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayVpcAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGateways(input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewaysPages(input *ec2.DescribeTransitGatewaysInput, fn func(*ec2.DescribeTransitGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewaysInput, fn func(*ec2.DescribeTransitGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewaysWithContext(ctx context.Context, input *ec2.DescribeTransitGatewaysInput, opts ...request.Option) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrunkInterfaceAssociations(input *ec2.DescribeTrunkInterfaceAssociationsInput) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeTrunkInterfaceAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrunkInterfaceAssociationsPages(input *ec2.DescribeTrunkInterfaceAssociationsInput, fn func(*ec2.DescribeTrunkInterfaceAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrunkInterfaceAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeTrunkInterfaceAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrunkInterfaceAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrunkInterfaceAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeTrunkInterfaceAssociationsInput, fn func(*ec2.DescribeTrunkInterfaceAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrunkInterfaceAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeTrunkInterfaceAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrunkInterfaceAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrunkInterfaceAssociationsWithContext(ctx context.Context, input *ec2.DescribeTrunkInterfaceAssociationsInput, opts ...request.Option) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeTrunkInterfaceAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessEndpoints(input *ec2.DescribeVerifiedAccessEndpointsInput) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessEndpointsPages(input *ec2.DescribeVerifiedAccessEndpointsInput, fn func(*ec2.DescribeVerifiedAccessEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessEndpointsInput, fn func(*ec2.DescribeVerifiedAccessEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessEndpointsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessEndpointsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessGroups(input *ec2.DescribeVerifiedAccessGroupsInput) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessGroupsPages(input *ec2.DescribeVerifiedAccessGroupsInput, fn func(*ec2.DescribeVerifiedAccessGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessGroupsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessGroupsInput, fn func(*ec2.DescribeVerifiedAccessGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessGroupsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessGroupsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessGroupsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurations(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, fn func(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, fn func(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstances(input *ec2.DescribeVerifiedAccessInstancesInput) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstancesPages(input *ec2.DescribeVerifiedAccessInstancesInput, fn func(*ec2.DescribeVerifiedAccessInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstancesOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstancesInput, fn func(*ec2.DescribeVerifiedAccessInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstancesOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstancesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstancesInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessTrustProviders(input *ec2.DescribeVerifiedAccessTrustProvidersInput) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessTrustProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersPages(input *ec2.DescribeVerifiedAccessTrustProvidersInput, fn func(*ec2.DescribeVerifiedAccessTrustProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessTrustProvidersOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessTrustProvidersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessTrustProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessTrustProvidersInput, fn func(*ec2.DescribeVerifiedAccessTrustProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessTrustProvidersOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessTrustProvidersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessTrustProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessTrustProvidersInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessTrustProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeAttribute(input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeAttributeWithContext(ctx context.Context, input *ec2.DescribeVolumeAttributeInput, opts ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeStatus(input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeStatusPages(input *ec2.DescribeVolumeStatusInput, fn func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumeStatusOutput{}
	fnCacher := func(out *ec2.DescribeVolumeStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumeStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumeStatusPagesWithContext(ctx context.Context, input *ec2.DescribeVolumeStatusInput, fn func(*ec2.DescribeVolumeStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumeStatusOutput{}
	fnCacher := func(out *ec2.DescribeVolumeStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumeStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumeStatusWithContext(ctx context.Context, input *ec2.DescribeVolumeStatusInput, opts ...request.Option) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.EC2API.DescribeVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesModifications(input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesModificationsPages(input *ec2.DescribeVolumesModificationsInput, fn func(*ec2.DescribeVolumesModificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeVolumesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesModificationsPagesWithContext(ctx context.Context, input *ec2.DescribeVolumesModificationsInput, fn func(*ec2.DescribeVolumesModificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeVolumesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesModificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesModificationsWithContext(ctx context.Context, input *ec2.DescribeVolumesModificationsInput, opts ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesPages(input *ec2.DescribeVolumesInput, fn func(*ec2.DescribeVolumesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesOutput{}
	fnCacher := func(out *ec2.DescribeVolumesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesPagesWithContext(ctx context.Context, input *ec2.DescribeVolumesInput, fn func(*ec2.DescribeVolumesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesOutput{}
	fnCacher := func(out *ec2.DescribeVolumesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesWithContext(ctx context.Context, input *ec2.DescribeVolumesInput, opts ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcAttribute(input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVpcAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcAttributeWithContext(ctx context.Context, input *ec2.DescribeVpcAttributeInput, opts ...request.Option) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVpcAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLink(input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkDnsSupport(input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportPages(input *ec2.DescribeVpcClassicLinkDnsSupportInput, fn func(*ec2.DescribeVpcClassicLinkDnsSupportOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcClassicLinkDnsSupportOutput{}
	fnCacher := func(out *ec2.DescribeVpcClassicLinkDnsSupportOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcClassicLinkDnsSupportPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportPagesWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput, fn func(*ec2.DescribeVpcClassicLinkDnsSupportOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcClassicLinkDnsSupportOutput{}
	fnCacher := func(out *ec2.DescribeVpcClassicLinkDnsSupportOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcClassicLinkDnsSupportPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionNotifications(input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsPages(input *ec2.DescribeVpcEndpointConnectionNotificationsInput, fn func(*ec2.DescribeVpcEndpointConnectionNotificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionNotificationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionNotificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionNotificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput, fn func(*ec2.DescribeVpcEndpointConnectionNotificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionNotificationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionNotificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionNotificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnections(input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionsPages(input *ec2.DescribeVpcEndpointConnectionsInput, fn func(*ec2.DescribeVpcEndpointConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput, fn func(*ec2.DescribeVpcEndpointConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServiceConfigurations(input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServiceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsPages(input *ec2.DescribeVpcEndpointServiceConfigurationsInput, fn func(*ec2.DescribeVpcEndpointServiceConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServiceConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServiceConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServiceConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, fn func(*ec2.DescribeVpcEndpointServiceConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServiceConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServiceConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServiceConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServiceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicePermissions(input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicePermissionsPages(input *ec2.DescribeVpcEndpointServicePermissionsInput, fn func(*ec2.DescribeVpcEndpointServicePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServicePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServicePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServicePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServicePermissionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, fn func(*ec2.DescribeVpcEndpointServicePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServicePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServicePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServicePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServicePermissionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServices(input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicesInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpoints(input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointsPages(input *ec2.DescribeVpcEndpointsInput, fn func(*ec2.DescribeVpcEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointsInput, fn func(*ec2.DescribeVpcEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcPeeringConnections(input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcPeeringConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcPeeringConnectionsPages(input *ec2.DescribeVpcPeeringConnectionsInput, fn func(*ec2.DescribeVpcPeeringConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcPeeringConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcPeeringConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcPeeringConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcPeeringConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput, fn func(*ec2.DescribeVpcPeeringConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcPeeringConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcPeeringConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcPeeringConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcPeeringConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcPeeringConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcsPages(input *ec2.DescribeVpcsInput, fn func(*ec2.DescribeVpcsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcsOutput{}
	fnCacher := func(out *ec2.DescribeVpcsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcsInput, fn func(*ec2.DescribeVpcsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcsOutput{}
	fnCacher := func(out *ec2.DescribeVpcsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcsWithContext(ctx context.Context, input *ec2.DescribeVpcsInput, opts ...request.Option) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnConnections(input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnGateways(input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeVpnGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnGatewaysWithContext(ctx context.Context, input *ec2.DescribeVpnGatewaysInput, opts ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeVpnGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachClassicLinkVpc(input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachClassicLinkVpcOutput), nil
	}
	out, err := c.EC2API.DetachClassicLinkVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachClassicLinkVpcWithContext(ctx context.Context, input *ec2.DetachClassicLinkVpcInput, opts ...request.Option) (*ec2.DetachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachClassicLinkVpcOutput), nil
	}
	out, err := c.EC2API.DetachClassicLinkVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachInternetGateway(input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DetachInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachInternetGatewayWithContext(ctx context.Context, input *ec2.DetachInternetGatewayInput, opts ...request.Option) (*ec2.DetachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachInternetGatewayOutput), nil
	}
	out, err := c.EC2API.DetachInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachNetworkInterface(input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.DetachNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachNetworkInterfaceWithContext(ctx context.Context, input *ec2.DetachNetworkInterfaceInput, opts ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachNetworkInterfaceOutput), nil
	}
	out, err := c.EC2API.DetachNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachVerifiedAccessTrustProvider(input *ec2.DetachVerifiedAccessTrustProviderInput) (*ec2.DetachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.DetachVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.DetachVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.DetachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.DetachVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachVpnGateway(input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVpnGatewayOutput), nil
	}
	out, err := c.EC2API.DetachVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DetachVpnGatewayWithContext(ctx context.Context, input *ec2.DetachVpnGatewayInput, opts ...request.Option) (*ec2.DetachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVpnGatewayOutput), nil
	}
	out, err := c.EC2API.DetachVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableAddressTransfer(input *ec2.DisableAddressTransferInput) (*ec2.DisableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAddressTransferOutput), nil
	}
	out, err := c.EC2API.DisableAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableAddressTransferWithContext(ctx context.Context, input *ec2.DisableAddressTransferInput, opts ...request.Option) (*ec2.DisableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAddressTransferOutput), nil
	}
	out, err := c.EC2API.DisableAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableAwsNetworkPerformanceMetricSubscription(input *ec2.DisableAwsNetworkPerformanceMetricSubscriptionInput) (*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.EC2API.DisableAwsNetworkPerformanceMetricSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx context.Context, input *ec2.DisableAwsNetworkPerformanceMetricSubscriptionInput, opts ...request.Option) (*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.EC2API.DisableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableEbsEncryptionByDefault(input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.DisableEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.DisableEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.DisableEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableFastLaunch(input *ec2.DisableFastLaunchInput) (*ec2.DisableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastLaunchOutput), nil
	}
	out, err := c.EC2API.DisableFastLaunch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableFastLaunchWithContext(ctx context.Context, input *ec2.DisableFastLaunchInput, opts ...request.Option) (*ec2.DisableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastLaunchOutput), nil
	}
	out, err := c.EC2API.DisableFastLaunchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableFastSnapshotRestores(input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DisableFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.DisableFastSnapshotRestoresInput, opts ...request.Option) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DisableFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableImageDeprecation(input *ec2.DisableImageDeprecationInput) (*ec2.DisableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableImageDeprecationOutput), nil
	}
	out, err := c.EC2API.DisableImageDeprecation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableImageDeprecationWithContext(ctx context.Context, input *ec2.DisableImageDeprecationInput, opts ...request.Option) (*ec2.DisableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableImageDeprecationOutput), nil
	}
	out, err := c.EC2API.DisableImageDeprecationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableIpamOrganizationAdminAccount(input *ec2.DisableIpamOrganizationAdminAccountInput) (*ec2.DisableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.EC2API.DisableIpamOrganizationAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableIpamOrganizationAdminAccountWithContext(ctx context.Context, input *ec2.DisableIpamOrganizationAdminAccountInput, opts ...request.Option) (*ec2.DisableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.EC2API.DisableIpamOrganizationAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableSerialConsoleAccess(input *ec2.DisableSerialConsoleAccessInput) (*ec2.DisableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableSerialConsoleAccessOutput), nil
	}
	out, err := c.EC2API.DisableSerialConsoleAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableSerialConsoleAccessWithContext(ctx context.Context, input *ec2.DisableSerialConsoleAccessInput, opts ...request.Option) (*ec2.DisableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableSerialConsoleAccessOutput), nil
	}
	out, err := c.EC2API.DisableSerialConsoleAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableTransitGatewayRouteTablePropagation(input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.EC2API.DisableTransitGatewayRouteTablePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableTransitGatewayRouteTablePropagationWithContext(ctx context.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput, opts ...request.Option) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.EC2API.DisableTransitGatewayRouteTablePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVgwRoutePropagation(input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVgwRoutePropagationOutput), nil
	}
	out, err := c.EC2API.DisableVgwRoutePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVgwRoutePropagationWithContext(ctx context.Context, input *ec2.DisableVgwRoutePropagationInput, opts ...request.Option) (*ec2.DisableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVgwRoutePropagationOutput), nil
	}
	out, err := c.EC2API.DisableVgwRoutePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVpcClassicLink(input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DisableVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVpcClassicLinkDnsSupport(input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DisableVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DisableVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisableVpcClassicLinkWithContext(ctx context.Context, input *ec2.DisableVpcClassicLinkInput, opts ...request.Option) (*ec2.DisableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DisableVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateAddress(input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateAddressOutput), nil
	}
	out, err := c.EC2API.DisassociateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateAddressWithContext(ctx context.Context, input *ec2.DisassociateAddressInput, opts ...request.Option) (*ec2.DisassociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateAddressOutput), nil
	}
	out, err := c.EC2API.DisassociateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateClientVpnTargetNetwork(input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.DisassociateClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.DisassociateClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.EC2API.DisassociateClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateEnclaveCertificateIamRole(input *ec2.DisassociateEnclaveCertificateIamRoleInput) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.EC2API.DisassociateEnclaveCertificateIamRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateEnclaveCertificateIamRoleWithContext(ctx context.Context, input *ec2.DisassociateEnclaveCertificateIamRoleInput, opts ...request.Option) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.EC2API.DisassociateEnclaveCertificateIamRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateIamInstanceProfile(input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateIamInstanceProfileOutput), nil
	}
	out, err := c.EC2API.DisassociateIamInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateIamInstanceProfileWithContext(ctx context.Context, input *ec2.DisassociateIamInstanceProfileInput, opts ...request.Option) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateIamInstanceProfileOutput), nil
	}
	out, err := c.EC2API.DisassociateIamInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateInstanceEventWindow(input *ec2.DisassociateInstanceEventWindowInput) (*ec2.DisassociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.DisassociateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateInstanceEventWindowWithContext(ctx context.Context, input *ec2.DisassociateInstanceEventWindowInput, opts ...request.Option) (*ec2.DisassociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.DisassociateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateRouteTable(input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateRouteTableOutput), nil
	}
	out, err := c.EC2API.DisassociateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateRouteTableWithContext(ctx context.Context, input *ec2.DisassociateRouteTableInput, opts ...request.Option) (*ec2.DisassociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateRouteTableOutput), nil
	}
	out, err := c.EC2API.DisassociateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateSubnetCidrBlock(input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateSubnetCidrBlockOutput), nil
	}
	out, err := c.EC2API.DisassociateSubnetCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateSubnetCidrBlockWithContext(ctx context.Context, input *ec2.DisassociateSubnetCidrBlockInput, opts ...request.Option) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateSubnetCidrBlockOutput), nil
	}
	out, err := c.EC2API.DisassociateSubnetCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayMulticastDomain(input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayPolicyTable(input *ec2.DisassociateTransitGatewayPolicyTableInput) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayRouteTable(input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.EC2API.DisassociateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTrunkInterface(input *ec2.DisassociateTrunkInterfaceInput) (*ec2.DisassociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTrunkInterfaceOutput), nil
	}
	out, err := c.EC2API.DisassociateTrunkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateTrunkInterfaceWithContext(ctx context.Context, input *ec2.DisassociateTrunkInterfaceInput, opts ...request.Option) (*ec2.DisassociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTrunkInterfaceOutput), nil
	}
	out, err := c.EC2API.DisassociateTrunkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateVpcCidrBlock(input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateVpcCidrBlockOutput), nil
	}
	out, err := c.EC2API.DisassociateVpcCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DisassociateVpcCidrBlockWithContext(ctx context.Context, input *ec2.DisassociateVpcCidrBlockInput, opts ...request.Option) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateVpcCidrBlockOutput), nil
	}
	out, err := c.EC2API.DisassociateVpcCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableAddressTransfer(input *ec2.EnableAddressTransferInput) (*ec2.EnableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAddressTransferOutput), nil
	}
	out, err := c.EC2API.EnableAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableAddressTransferWithContext(ctx context.Context, input *ec2.EnableAddressTransferInput, opts ...request.Option) (*ec2.EnableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAddressTransferOutput), nil
	}
	out, err := c.EC2API.EnableAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableAwsNetworkPerformanceMetricSubscription(input *ec2.EnableAwsNetworkPerformanceMetricSubscriptionInput) (*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.EC2API.EnableAwsNetworkPerformanceMetricSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx context.Context, input *ec2.EnableAwsNetworkPerformanceMetricSubscriptionInput, opts ...request.Option) (*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.EC2API.EnableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableEbsEncryptionByDefault(input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.EnableEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.EnableEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.EnableEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableFastLaunch(input *ec2.EnableFastLaunchInput) (*ec2.EnableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastLaunchOutput), nil
	}
	out, err := c.EC2API.EnableFastLaunch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableFastLaunchWithContext(ctx context.Context, input *ec2.EnableFastLaunchInput, opts ...request.Option) (*ec2.EnableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastLaunchOutput), nil
	}
	out, err := c.EC2API.EnableFastLaunchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableFastSnapshotRestores(input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.EnableFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.EnableFastSnapshotRestoresInput, opts ...request.Option) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.EnableFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableImageDeprecation(input *ec2.EnableImageDeprecationInput) (*ec2.EnableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableImageDeprecationOutput), nil
	}
	out, err := c.EC2API.EnableImageDeprecation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableImageDeprecationWithContext(ctx context.Context, input *ec2.EnableImageDeprecationInput, opts ...request.Option) (*ec2.EnableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableImageDeprecationOutput), nil
	}
	out, err := c.EC2API.EnableImageDeprecationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableIpamOrganizationAdminAccount(input *ec2.EnableIpamOrganizationAdminAccountInput) (*ec2.EnableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.EC2API.EnableIpamOrganizationAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableIpamOrganizationAdminAccountWithContext(ctx context.Context, input *ec2.EnableIpamOrganizationAdminAccountInput, opts ...request.Option) (*ec2.EnableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.EC2API.EnableIpamOrganizationAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableReachabilityAnalyzerOrganizationSharing(input *ec2.EnableReachabilityAnalyzerOrganizationSharingInput) (*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput), nil
	}
	out, err := c.EC2API.EnableReachabilityAnalyzerOrganizationSharing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableReachabilityAnalyzerOrganizationSharingWithContext(ctx context.Context, input *ec2.EnableReachabilityAnalyzerOrganizationSharingInput, opts ...request.Option) (*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput), nil
	}
	out, err := c.EC2API.EnableReachabilityAnalyzerOrganizationSharingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableSerialConsoleAccess(input *ec2.EnableSerialConsoleAccessInput) (*ec2.EnableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableSerialConsoleAccessOutput), nil
	}
	out, err := c.EC2API.EnableSerialConsoleAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableSerialConsoleAccessWithContext(ctx context.Context, input *ec2.EnableSerialConsoleAccessInput, opts ...request.Option) (*ec2.EnableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableSerialConsoleAccessOutput), nil
	}
	out, err := c.EC2API.EnableSerialConsoleAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableTransitGatewayRouteTablePropagation(input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.EC2API.EnableTransitGatewayRouteTablePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableTransitGatewayRouteTablePropagationWithContext(ctx context.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput, opts ...request.Option) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.EC2API.EnableTransitGatewayRouteTablePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVgwRoutePropagation(input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVgwRoutePropagationOutput), nil
	}
	out, err := c.EC2API.EnableVgwRoutePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVgwRoutePropagationWithContext(ctx context.Context, input *ec2.EnableVgwRoutePropagationInput, opts ...request.Option) (*ec2.EnableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVgwRoutePropagationOutput), nil
	}
	out, err := c.EC2API.EnableVgwRoutePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVolumeIO(input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVolumeIOOutput), nil
	}
	out, err := c.EC2API.EnableVolumeIO(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVolumeIOWithContext(ctx context.Context, input *ec2.EnableVolumeIOInput, opts ...request.Option) (*ec2.EnableVolumeIOOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVolumeIOOutput), nil
	}
	out, err := c.EC2API.EnableVolumeIOWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVpcClassicLink(input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.EnableVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVpcClassicLinkDnsSupport(input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.EnableVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.EnableVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) EnableVpcClassicLinkWithContext(ctx context.Context, input *ec2.EnableVpcClassicLinkInput, opts ...request.Option) (*ec2.EnableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.EnableVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportClientVpnClientCertificateRevocationList(input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.EC2API.ExportClientVpnClientCertificateRevocationList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportClientVpnClientCertificateRevocationListWithContext(ctx context.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput, opts ...request.Option) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.EC2API.ExportClientVpnClientCertificateRevocationListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportClientVpnClientConfiguration(input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientConfigurationOutput), nil
	}
	out, err := c.EC2API.ExportClientVpnClientConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportClientVpnClientConfigurationWithContext(ctx context.Context, input *ec2.ExportClientVpnClientConfigurationInput, opts ...request.Option) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientConfigurationOutput), nil
	}
	out, err := c.EC2API.ExportClientVpnClientConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportImage(input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportImageOutput), nil
	}
	out, err := c.EC2API.ExportImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportImageWithContext(ctx context.Context, input *ec2.ExportImageInput, opts ...request.Option) (*ec2.ExportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportImageOutput), nil
	}
	out, err := c.EC2API.ExportImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportTransitGatewayRoutes(input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.ExportTransitGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ExportTransitGatewayRoutesWithContext(ctx context.Context, input *ec2.ExportTransitGatewayRoutesInput, opts ...request.Option) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.ExportTransitGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedEnclaveCertificateIamRoles(input *ec2.GetAssociatedEnclaveCertificateIamRolesInput) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.EC2API.GetAssociatedEnclaveCertificateIamRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedEnclaveCertificateIamRolesWithContext(ctx context.Context, input *ec2.GetAssociatedEnclaveCertificateIamRolesInput, opts ...request.Option) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.EC2API.GetAssociatedEnclaveCertificateIamRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedIpv6PoolCidrs(input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetAssociatedIpv6PoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedIpv6PoolCidrsPages(input *ec2.GetAssociatedIpv6PoolCidrsInput, fn func(*ec2.GetAssociatedIpv6PoolCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAssociatedIpv6PoolCidrsOutput{}
	fnCacher := func(out *ec2.GetAssociatedIpv6PoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetAssociatedIpv6PoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAssociatedIpv6PoolCidrsPagesWithContext(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput, fn func(*ec2.GetAssociatedIpv6PoolCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAssociatedIpv6PoolCidrsOutput{}
	fnCacher := func(out *ec2.GetAssociatedIpv6PoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetAssociatedIpv6PoolCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAssociatedIpv6PoolCidrsWithContext(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput, opts ...request.Option) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetAssociatedIpv6PoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAwsNetworkPerformanceData(input *ec2.GetAwsNetworkPerformanceDataInput) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.EC2API.GetAwsNetworkPerformanceData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAwsNetworkPerformanceDataPages(input *ec2.GetAwsNetworkPerformanceDataInput, fn func(*ec2.GetAwsNetworkPerformanceDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAwsNetworkPerformanceDataOutput{}
	fnCacher := func(out *ec2.GetAwsNetworkPerformanceDataOutput, more bool) bool {
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
	if err := c.EC2API.GetAwsNetworkPerformanceDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAwsNetworkPerformanceDataPagesWithContext(ctx context.Context, input *ec2.GetAwsNetworkPerformanceDataInput, fn func(*ec2.GetAwsNetworkPerformanceDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAwsNetworkPerformanceDataOutput{}
	fnCacher := func(out *ec2.GetAwsNetworkPerformanceDataOutput, more bool) bool {
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
	if err := c.EC2API.GetAwsNetworkPerformanceDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAwsNetworkPerformanceDataWithContext(ctx context.Context, input *ec2.GetAwsNetworkPerformanceDataInput, opts ...request.Option) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.EC2API.GetAwsNetworkPerformanceDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCapacityReservationUsage(input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.EC2API.GetCapacityReservationUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCapacityReservationUsageWithContext(ctx context.Context, input *ec2.GetCapacityReservationUsageInput, opts ...request.Option) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.EC2API.GetCapacityReservationUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCoipPoolUsage(input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.EC2API.GetCoipPoolUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCoipPoolUsageWithContext(ctx context.Context, input *ec2.GetCoipPoolUsageInput, opts ...request.Option) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.EC2API.GetCoipPoolUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleOutput(input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.EC2API.GetConsoleOutput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleOutputWithContext(ctx context.Context, input *ec2.GetConsoleOutputInput, opts ...request.Option) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.EC2API.GetConsoleOutputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleScreenshot(input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.EC2API.GetConsoleScreenshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleScreenshotWithContext(ctx context.Context, input *ec2.GetConsoleScreenshotInput, opts ...request.Option) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.EC2API.GetConsoleScreenshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetDefaultCreditSpecification(input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.GetDefaultCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetDefaultCreditSpecificationWithContext(ctx context.Context, input *ec2.GetDefaultCreditSpecificationInput, opts ...request.Option) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.GetDefaultCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsDefaultKmsKeyId(input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.GetEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.GetEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.GetEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsEncryptionByDefault(input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.GetEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.GetEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.GetEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetFlowLogsIntegrationTemplate(input *ec2.GetFlowLogsIntegrationTemplateInput) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.EC2API.GetFlowLogsIntegrationTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetFlowLogsIntegrationTemplateWithContext(ctx context.Context, input *ec2.GetFlowLogsIntegrationTemplateInput, opts ...request.Option) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.EC2API.GetFlowLogsIntegrationTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetGroupsForCapacityReservation(input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.EC2API.GetGroupsForCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetGroupsForCapacityReservationPages(input *ec2.GetGroupsForCapacityReservationInput, fn func(*ec2.GetGroupsForCapacityReservationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetGroupsForCapacityReservationOutput{}
	fnCacher := func(out *ec2.GetGroupsForCapacityReservationOutput, more bool) bool {
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
	if err := c.EC2API.GetGroupsForCapacityReservationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetGroupsForCapacityReservationPagesWithContext(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput, fn func(*ec2.GetGroupsForCapacityReservationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetGroupsForCapacityReservationOutput{}
	fnCacher := func(out *ec2.GetGroupsForCapacityReservationOutput, more bool) bool {
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
	if err := c.EC2API.GetGroupsForCapacityReservationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetGroupsForCapacityReservationWithContext(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput, opts ...request.Option) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.EC2API.GetGroupsForCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetHostReservationPurchasePreview(input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.EC2API.GetHostReservationPurchasePreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetHostReservationPurchasePreviewWithContext(ctx context.Context, input *ec2.GetHostReservationPurchasePreviewInput, opts ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.EC2API.GetHostReservationPurchasePreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceTypesFromInstanceRequirements(input *ec2.GetInstanceTypesFromInstanceRequirementsInput) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.EC2API.GetInstanceTypesFromInstanceRequirements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsPages(input *ec2.GetInstanceTypesFromInstanceRequirementsInput, fn func(*ec2.GetInstanceTypesFromInstanceRequirementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetInstanceTypesFromInstanceRequirementsOutput{}
	fnCacher := func(out *ec2.GetInstanceTypesFromInstanceRequirementsOutput, more bool) bool {
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
	if err := c.EC2API.GetInstanceTypesFromInstanceRequirementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsPagesWithContext(ctx context.Context, input *ec2.GetInstanceTypesFromInstanceRequirementsInput, fn func(*ec2.GetInstanceTypesFromInstanceRequirementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetInstanceTypesFromInstanceRequirementsOutput{}
	fnCacher := func(out *ec2.GetInstanceTypesFromInstanceRequirementsOutput, more bool) bool {
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
	if err := c.EC2API.GetInstanceTypesFromInstanceRequirementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsWithContext(ctx context.Context, input *ec2.GetInstanceTypesFromInstanceRequirementsInput, opts ...request.Option) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.EC2API.GetInstanceTypesFromInstanceRequirementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceUefiData(input *ec2.GetInstanceUefiDataInput) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.EC2API.GetInstanceUefiData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceUefiDataWithContext(ctx context.Context, input *ec2.GetInstanceUefiDataInput, opts ...request.Option) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.EC2API.GetInstanceUefiDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamAddressHistory(input *ec2.GetIpamAddressHistoryInput) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.EC2API.GetIpamAddressHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamAddressHistoryPages(input *ec2.GetIpamAddressHistoryInput, fn func(*ec2.GetIpamAddressHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamAddressHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamAddressHistoryOutput{}
	fnCacher := func(out *ec2.GetIpamAddressHistoryOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamAddressHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamAddressHistoryPagesWithContext(ctx context.Context, input *ec2.GetIpamAddressHistoryInput, fn func(*ec2.GetIpamAddressHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamAddressHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamAddressHistoryOutput{}
	fnCacher := func(out *ec2.GetIpamAddressHistoryOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamAddressHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamAddressHistoryWithContext(ctx context.Context, input *ec2.GetIpamAddressHistoryInput, opts ...request.Option) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.EC2API.GetIpamAddressHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolAllocations(input *ec2.GetIpamPoolAllocationsInput) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolAllocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolAllocationsPages(input *ec2.GetIpamPoolAllocationsInput, fn func(*ec2.GetIpamPoolAllocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolAllocationsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolAllocationsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolAllocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolAllocationsPagesWithContext(ctx context.Context, input *ec2.GetIpamPoolAllocationsInput, fn func(*ec2.GetIpamPoolAllocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolAllocationsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolAllocationsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolAllocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolAllocationsWithContext(ctx context.Context, input *ec2.GetIpamPoolAllocationsInput, opts ...request.Option) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolAllocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolCidrs(input *ec2.GetIpamPoolCidrsInput) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolCidrsPages(input *ec2.GetIpamPoolCidrsInput, fn func(*ec2.GetIpamPoolCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolCidrsPagesWithContext(ctx context.Context, input *ec2.GetIpamPoolCidrsInput, fn func(*ec2.GetIpamPoolCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolCidrsWithContext(ctx context.Context, input *ec2.GetIpamPoolCidrsInput, opts ...request.Option) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamResourceCidrs(input *ec2.GetIpamResourceCidrsInput) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamResourceCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamResourceCidrsPages(input *ec2.GetIpamResourceCidrsInput, fn func(*ec2.GetIpamResourceCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamResourceCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamResourceCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamResourceCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamResourceCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamResourceCidrsPagesWithContext(ctx context.Context, input *ec2.GetIpamResourceCidrsInput, fn func(*ec2.GetIpamResourceCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamResourceCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamResourceCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamResourceCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamResourceCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamResourceCidrsWithContext(ctx context.Context, input *ec2.GetIpamResourceCidrsInput, opts ...request.Option) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamResourceCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetLaunchTemplateData(input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.EC2API.GetLaunchTemplateData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetLaunchTemplateDataWithContext(ctx context.Context, input *ec2.GetLaunchTemplateDataInput, opts ...request.Option) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.EC2API.GetLaunchTemplateDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListAssociations(input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListAssociationsPages(input *ec2.GetManagedPrefixListAssociationsInput, fn func(*ec2.GetManagedPrefixListAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListAssociationsOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListAssociationsPagesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput, fn func(*ec2.GetManagedPrefixListAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListAssociationsOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListAssociationsWithContext(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput, opts ...request.Option) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListEntries(input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListEntriesPages(input *ec2.GetManagedPrefixListEntriesInput, fn func(*ec2.GetManagedPrefixListEntriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListEntriesOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListEntriesOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListEntriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListEntriesPagesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput, fn func(*ec2.GetManagedPrefixListEntriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListEntriesOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListEntriesOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListEntriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListEntriesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput, opts ...request.Option) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeAnalysisFindings(input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeAnalysisFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeContent(input *ec2.GetNetworkInsightsAccessScopeContentInput) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeContentWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeContentInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetPasswordData(input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.EC2API.GetPasswordData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetPasswordDataWithContext(ctx context.Context, input *ec2.GetPasswordDataInput, opts ...request.Option) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.EC2API.GetPasswordDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetReservedInstancesExchangeQuote(input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.GetReservedInstancesExchangeQuote(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetReservedInstancesExchangeQuoteWithContext(ctx context.Context, input *ec2.GetReservedInstancesExchangeQuoteInput, opts ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.GetReservedInstancesExchangeQuoteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSerialConsoleAccessStatus(input *ec2.GetSerialConsoleAccessStatusInput) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.EC2API.GetSerialConsoleAccessStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSerialConsoleAccessStatusWithContext(ctx context.Context, input *ec2.GetSerialConsoleAccessStatusInput, opts ...request.Option) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.EC2API.GetSerialConsoleAccessStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSpotPlacementScores(input *ec2.GetSpotPlacementScoresInput) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.EC2API.GetSpotPlacementScores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSpotPlacementScoresPages(input *ec2.GetSpotPlacementScoresInput, fn func(*ec2.GetSpotPlacementScoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetSpotPlacementScoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetSpotPlacementScoresOutput{}
	fnCacher := func(out *ec2.GetSpotPlacementScoresOutput, more bool) bool {
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
	if err := c.EC2API.GetSpotPlacementScoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetSpotPlacementScoresPagesWithContext(ctx context.Context, input *ec2.GetSpotPlacementScoresInput, fn func(*ec2.GetSpotPlacementScoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetSpotPlacementScoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetSpotPlacementScoresOutput{}
	fnCacher := func(out *ec2.GetSpotPlacementScoresOutput, more bool) bool {
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
	if err := c.EC2API.GetSpotPlacementScoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetSpotPlacementScoresWithContext(ctx context.Context, input *ec2.GetSpotPlacementScoresInput, opts ...request.Option) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.EC2API.GetSpotPlacementScoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSubnetCidrReservations(input *ec2.GetSubnetCidrReservationsInput) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.EC2API.GetSubnetCidrReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSubnetCidrReservationsWithContext(ctx context.Context, input *ec2.GetSubnetCidrReservationsInput, opts ...request.Option) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.EC2API.GetSubnetCidrReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayAttachmentPropagations(input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayAttachmentPropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsPages(input *ec2.GetTransitGatewayAttachmentPropagationsInput, fn func(*ec2.GetTransitGatewayAttachmentPropagationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayAttachmentPropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayAttachmentPropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayAttachmentPropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput, fn func(*ec2.GetTransitGatewayAttachmentPropagationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayAttachmentPropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayAttachmentPropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayAttachmentPropagationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayAttachmentPropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociations(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsPages(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, fn func(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayMulticastDomainAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayMulticastDomainAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, fn func(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayMulticastDomainAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayMulticastDomainAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableAssociations(input *ec2.GetTransitGatewayPolicyTableAssociationsInput) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsPages(input *ec2.GetTransitGatewayPolicyTableAssociationsInput, fn func(*ec2.GetTransitGatewayPolicyTableAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPolicyTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPolicyTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPolicyTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableAssociationsInput, fn func(*ec2.GetTransitGatewayPolicyTableAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPolicyTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPolicyTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPolicyTableAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableEntries(input *ec2.GetTransitGatewayPolicyTableEntriesInput) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableEntriesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableEntriesInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPrefixListReferences(input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPrefixListReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPrefixListReferencesPages(input *ec2.GetTransitGatewayPrefixListReferencesInput, fn func(*ec2.GetTransitGatewayPrefixListReferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPrefixListReferencesOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPrefixListReferencesOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPrefixListReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPrefixListReferencesPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput, fn func(*ec2.GetTransitGatewayPrefixListReferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPrefixListReferencesOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPrefixListReferencesOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPrefixListReferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPrefixListReferencesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput, opts ...request.Option) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPrefixListReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTableAssociations(input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsPages(input *ec2.GetTransitGatewayRouteTableAssociationsInput, fn func(*ec2.GetTransitGatewayRouteTableAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput, fn func(*ec2.GetTransitGatewayRouteTableAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTableAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTablePropagations(input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTablePropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsPages(input *ec2.GetTransitGatewayRouteTablePropagationsInput, fn func(*ec2.GetTransitGatewayRouteTablePropagationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTablePropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTablePropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTablePropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput, fn func(*ec2.GetTransitGatewayRouteTablePropagationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTablePropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTablePropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTablePropagationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTablePropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessEndpointPolicy(input *ec2.GetVerifiedAccessEndpointPolicyInput) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessEndpointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessEndpointPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessEndpointPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessEndpointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessGroupPolicy(input *ec2.GetVerifiedAccessGroupPolicyInput) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessGroupPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessGroupPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceSampleConfiguration(input *ec2.GetVpnConnectionDeviceSampleConfigurationInput) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceSampleConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceSampleConfigurationWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceSampleConfigurationInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceSampleConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceTypes(input *ec2.GetVpnConnectionDeviceTypesInput) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceTypesPages(input *ec2.GetVpnConnectionDeviceTypesInput, fn func(*ec2.GetVpnConnectionDeviceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetVpnConnectionDeviceTypesOutput{}
	fnCacher := func(out *ec2.GetVpnConnectionDeviceTypesOutput, more bool) bool {
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
	if err := c.EC2API.GetVpnConnectionDeviceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetVpnConnectionDeviceTypesPagesWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceTypesInput, fn func(*ec2.GetVpnConnectionDeviceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetVpnConnectionDeviceTypesOutput{}
	fnCacher := func(out *ec2.GetVpnConnectionDeviceTypesOutput, more bool) bool {
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
	if err := c.EC2API.GetVpnConnectionDeviceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetVpnConnectionDeviceTypesWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceTypesInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportClientVpnClientCertificateRevocationList(input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.EC2API.ImportClientVpnClientCertificateRevocationList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportClientVpnClientCertificateRevocationListWithContext(ctx context.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput, opts ...request.Option) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.EC2API.ImportClientVpnClientCertificateRevocationListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportImage(input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportImageOutput), nil
	}
	out, err := c.EC2API.ImportImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportImageWithContext(ctx context.Context, input *ec2.ImportImageInput, opts ...request.Option) (*ec2.ImportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportImageOutput), nil
	}
	out, err := c.EC2API.ImportImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportInstance(input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportInstanceOutput), nil
	}
	out, err := c.EC2API.ImportInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportInstanceWithContext(ctx context.Context, input *ec2.ImportInstanceInput, opts ...request.Option) (*ec2.ImportInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportInstanceOutput), nil
	}
	out, err := c.EC2API.ImportInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportKeyPair(input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportKeyPairOutput), nil
	}
	out, err := c.EC2API.ImportKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportKeyPairWithContext(ctx context.Context, input *ec2.ImportKeyPairInput, opts ...request.Option) (*ec2.ImportKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportKeyPairOutput), nil
	}
	out, err := c.EC2API.ImportKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportSnapshot(input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportSnapshotOutput), nil
	}
	out, err := c.EC2API.ImportSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportSnapshotWithContext(ctx context.Context, input *ec2.ImportSnapshotInput, opts ...request.Option) (*ec2.ImportSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportSnapshotOutput), nil
	}
	out, err := c.EC2API.ImportSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportVolume(input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportVolumeOutput), nil
	}
	out, err := c.EC2API.ImportVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ImportVolumeWithContext(ctx context.Context, input *ec2.ImportVolumeInput, opts ...request.Option) (*ec2.ImportVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportVolumeOutput), nil
	}
	out, err := c.EC2API.ImportVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListImagesInRecycleBin(input *ec2.ListImagesInRecycleBinInput) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListImagesInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListImagesInRecycleBinPages(input *ec2.ListImagesInRecycleBinInput, fn func(*ec2.ListImagesInRecycleBinOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListImagesInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListImagesInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListImagesInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListImagesInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListImagesInRecycleBinPagesWithContext(ctx context.Context, input *ec2.ListImagesInRecycleBinInput, fn func(*ec2.ListImagesInRecycleBinOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListImagesInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListImagesInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListImagesInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListImagesInRecycleBinPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListImagesInRecycleBinWithContext(ctx context.Context, input *ec2.ListImagesInRecycleBinInput, opts ...request.Option) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListImagesInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListSnapshotsInRecycleBin(input *ec2.ListSnapshotsInRecycleBinInput) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListSnapshotsInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListSnapshotsInRecycleBinPages(input *ec2.ListSnapshotsInRecycleBinInput, fn func(*ec2.ListSnapshotsInRecycleBinOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListSnapshotsInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListSnapshotsInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListSnapshotsInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListSnapshotsInRecycleBinPagesWithContext(ctx context.Context, input *ec2.ListSnapshotsInRecycleBinInput, fn func(*ec2.ListSnapshotsInRecycleBinOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListSnapshotsInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListSnapshotsInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListSnapshotsInRecycleBinPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListSnapshotsInRecycleBinWithContext(ctx context.Context, input *ec2.ListSnapshotsInRecycleBinInput, opts ...request.Option) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListSnapshotsInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyAddressAttribute(input *ec2.ModifyAddressAttributeInput) (*ec2.ModifyAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAddressAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyAddressAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyAddressAttributeWithContext(ctx context.Context, input *ec2.ModifyAddressAttributeInput, opts ...request.Option) (*ec2.ModifyAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAddressAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyAddressAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyAvailabilityZoneGroup(input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAvailabilityZoneGroupOutput), nil
	}
	out, err := c.EC2API.ModifyAvailabilityZoneGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyAvailabilityZoneGroupWithContext(ctx context.Context, input *ec2.ModifyAvailabilityZoneGroupInput, opts ...request.Option) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAvailabilityZoneGroupOutput), nil
	}
	out, err := c.EC2API.ModifyAvailabilityZoneGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyCapacityReservation(input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationOutput), nil
	}
	out, err := c.EC2API.ModifyCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyCapacityReservationFleet(input *ec2.ModifyCapacityReservationFleetInput) (*ec2.ModifyCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationFleetOutput), nil
	}
	out, err := c.EC2API.ModifyCapacityReservationFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyCapacityReservationFleetWithContext(ctx context.Context, input *ec2.ModifyCapacityReservationFleetInput, opts ...request.Option) (*ec2.ModifyCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationFleetOutput), nil
	}
	out, err := c.EC2API.ModifyCapacityReservationFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyCapacityReservationWithContext(ctx context.Context, input *ec2.ModifyCapacityReservationInput, opts ...request.Option) (*ec2.ModifyCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationOutput), nil
	}
	out, err := c.EC2API.ModifyCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyClientVpnEndpoint(input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyClientVpnEndpointWithContext(ctx context.Context, input *ec2.ModifyClientVpnEndpointInput, opts ...request.Option) (*ec2.ModifyClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyClientVpnEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyDefaultCreditSpecification(input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.ModifyDefaultCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyDefaultCreditSpecificationWithContext(ctx context.Context, input *ec2.ModifyDefaultCreditSpecificationInput, opts ...request.Option) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.ModifyDefaultCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyEbsDefaultKmsKeyId(input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.ModifyEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.ModifyEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyFleet(input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFleetOutput), nil
	}
	out, err := c.EC2API.ModifyFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyFleetWithContext(ctx context.Context, input *ec2.ModifyFleetInput, opts ...request.Option) (*ec2.ModifyFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFleetOutput), nil
	}
	out, err := c.EC2API.ModifyFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyFpgaImageAttribute(input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyFpgaImageAttributeWithContext(ctx context.Context, input *ec2.ModifyFpgaImageAttributeInput, opts ...request.Option) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyHosts(input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyHostsOutput), nil
	}
	out, err := c.EC2API.ModifyHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyHostsWithContext(ctx context.Context, input *ec2.ModifyHostsInput, opts ...request.Option) (*ec2.ModifyHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyHostsOutput), nil
	}
	out, err := c.EC2API.ModifyHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIdFormat(input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdFormatOutput), nil
	}
	out, err := c.EC2API.ModifyIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIdFormatWithContext(ctx context.Context, input *ec2.ModifyIdFormatInput, opts ...request.Option) (*ec2.ModifyIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdFormatOutput), nil
	}
	out, err := c.EC2API.ModifyIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIdentityIdFormat(input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.ModifyIdentityIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIdentityIdFormatWithContext(ctx context.Context, input *ec2.ModifyIdentityIdFormatInput, opts ...request.Option) (*ec2.ModifyIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.ModifyIdentityIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyImageAttribute(input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyImageAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyImageAttributeWithContext(ctx context.Context, input *ec2.ModifyImageAttributeInput, opts ...request.Option) (*ec2.ModifyImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyImageAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceAttribute(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceAttributeWithContext(ctx context.Context, input *ec2.ModifyInstanceAttributeInput, opts ...request.Option) (*ec2.ModifyInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceCapacityReservationAttributes(input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCapacityReservationAttributesOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceCapacityReservationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceCapacityReservationAttributesWithContext(ctx context.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput, opts ...request.Option) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCapacityReservationAttributesOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceCapacityReservationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceCreditSpecification(input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceCreditSpecificationWithContext(ctx context.Context, input *ec2.ModifyInstanceCreditSpecificationInput, opts ...request.Option) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceEventStartTime(input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventStartTimeOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceEventStartTime(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceEventStartTimeWithContext(ctx context.Context, input *ec2.ModifyInstanceEventStartTimeInput, opts ...request.Option) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventStartTimeOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceEventStartTimeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceEventWindow(input *ec2.ModifyInstanceEventWindowInput) (*ec2.ModifyInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceEventWindowWithContext(ctx context.Context, input *ec2.ModifyInstanceEventWindowInput, opts ...request.Option) (*ec2.ModifyInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventWindowOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceMaintenanceOptions(input *ec2.ModifyInstanceMaintenanceOptionsInput) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMaintenanceOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceMaintenanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceMaintenanceOptionsWithContext(ctx context.Context, input *ec2.ModifyInstanceMaintenanceOptionsInput, opts ...request.Option) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMaintenanceOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceMaintenanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceMetadataOptions(input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMetadataOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceMetadataOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstanceMetadataOptionsWithContext(ctx context.Context, input *ec2.ModifyInstanceMetadataOptionsInput, opts ...request.Option) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMetadataOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyInstanceMetadataOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstancePlacement(input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstancePlacementOutput), nil
	}
	out, err := c.EC2API.ModifyInstancePlacement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyInstancePlacementWithContext(ctx context.Context, input *ec2.ModifyInstancePlacementInput, opts ...request.Option) (*ec2.ModifyInstancePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstancePlacementOutput), nil
	}
	out, err := c.EC2API.ModifyInstancePlacementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpam(input *ec2.ModifyIpamInput) (*ec2.ModifyIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamOutput), nil
	}
	out, err := c.EC2API.ModifyIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamPool(input *ec2.ModifyIpamPoolInput) (*ec2.ModifyIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamPoolOutput), nil
	}
	out, err := c.EC2API.ModifyIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamPoolWithContext(ctx context.Context, input *ec2.ModifyIpamPoolInput, opts ...request.Option) (*ec2.ModifyIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamPoolOutput), nil
	}
	out, err := c.EC2API.ModifyIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamResourceCidr(input *ec2.ModifyIpamResourceCidrInput) (*ec2.ModifyIpamResourceCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamResourceCidrOutput), nil
	}
	out, err := c.EC2API.ModifyIpamResourceCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamResourceCidrWithContext(ctx context.Context, input *ec2.ModifyIpamResourceCidrInput, opts ...request.Option) (*ec2.ModifyIpamResourceCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamResourceCidrOutput), nil
	}
	out, err := c.EC2API.ModifyIpamResourceCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamScope(input *ec2.ModifyIpamScopeInput) (*ec2.ModifyIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamScopeOutput), nil
	}
	out, err := c.EC2API.ModifyIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamScopeWithContext(ctx context.Context, input *ec2.ModifyIpamScopeInput, opts ...request.Option) (*ec2.ModifyIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamScopeOutput), nil
	}
	out, err := c.EC2API.ModifyIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyIpamWithContext(ctx context.Context, input *ec2.ModifyIpamInput, opts ...request.Option) (*ec2.ModifyIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamOutput), nil
	}
	out, err := c.EC2API.ModifyIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyLaunchTemplate(input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.ModifyLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyLaunchTemplateWithContext(ctx context.Context, input *ec2.ModifyLaunchTemplateInput, opts ...request.Option) (*ec2.ModifyLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLaunchTemplateOutput), nil
	}
	out, err := c.EC2API.ModifyLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyLocalGatewayRoute(input *ec2.ModifyLocalGatewayRouteInput) (*ec2.ModifyLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.ModifyLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyLocalGatewayRouteWithContext(ctx context.Context, input *ec2.ModifyLocalGatewayRouteInput, opts ...request.Option) (*ec2.ModifyLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLocalGatewayRouteOutput), nil
	}
	out, err := c.EC2API.ModifyLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyManagedPrefixList(input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.ModifyManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyManagedPrefixListWithContext(ctx context.Context, input *ec2.ModifyManagedPrefixListInput, opts ...request.Option) (*ec2.ModifyManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyManagedPrefixListOutput), nil
	}
	out, err := c.EC2API.ModifyManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyNetworkInterfaceAttribute(input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.ModifyNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyPrivateDnsNameOptions(input *ec2.ModifyPrivateDnsNameOptionsInput) (*ec2.ModifyPrivateDnsNameOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyPrivateDnsNameOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyPrivateDnsNameOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyPrivateDnsNameOptionsWithContext(ctx context.Context, input *ec2.ModifyPrivateDnsNameOptionsInput, opts ...request.Option) (*ec2.ModifyPrivateDnsNameOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyPrivateDnsNameOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyPrivateDnsNameOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyReservedInstances(input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyReservedInstancesOutput), nil
	}
	out, err := c.EC2API.ModifyReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyReservedInstancesWithContext(ctx context.Context, input *ec2.ModifyReservedInstancesInput, opts ...request.Option) (*ec2.ModifyReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyReservedInstancesOutput), nil
	}
	out, err := c.EC2API.ModifyReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySecurityGroupRules(input *ec2.ModifySecurityGroupRulesInput) (*ec2.ModifySecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.ModifySecurityGroupRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySecurityGroupRulesWithContext(ctx context.Context, input *ec2.ModifySecurityGroupRulesInput, opts ...request.Option) (*ec2.ModifySecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.ModifySecurityGroupRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySnapshotAttribute(input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.ModifySnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySnapshotAttributeWithContext(ctx context.Context, input *ec2.ModifySnapshotAttributeInput, opts ...request.Option) (*ec2.ModifySnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.ModifySnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySnapshotTier(input *ec2.ModifySnapshotTierInput) (*ec2.ModifySnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotTierOutput), nil
	}
	out, err := c.EC2API.ModifySnapshotTier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySnapshotTierWithContext(ctx context.Context, input *ec2.ModifySnapshotTierInput, opts ...request.Option) (*ec2.ModifySnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotTierOutput), nil
	}
	out, err := c.EC2API.ModifySnapshotTierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySpotFleetRequestWithContext(ctx context.Context, input *ec2.ModifySpotFleetRequestInput, opts ...request.Option) (*ec2.ModifySpotFleetRequestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySpotFleetRequestOutput), nil
	}
	out, err := c.EC2API.ModifySpotFleetRequestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySubnetAttribute(input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySubnetAttributeOutput), nil
	}
	out, err := c.EC2API.ModifySubnetAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifySubnetAttributeWithContext(ctx context.Context, input *ec2.ModifySubnetAttributeInput, opts ...request.Option) (*ec2.ModifySubnetAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySubnetAttributeOutput), nil
	}
	out, err := c.EC2API.ModifySubnetAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorFilterNetworkServices(input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorFilterNetworkServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorFilterRule(input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorSession(input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorSessionInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorSessionOutput), nil
	}
	out, err := c.EC2API.ModifyTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGateway(input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGatewayPrefixListReference(input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGatewayVpcAttachment(input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyTransitGatewayWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayInput, opts ...request.Option) (*ec2.ModifyTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayOutput), nil
	}
	out, err := c.EC2API.ModifyTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessEndpoint(input *ec2.ModifyVerifiedAccessEndpointInput) (*ec2.ModifyVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessEndpointPolicy(input *ec2.ModifyVerifiedAccessEndpointPolicyInput) (*ec2.ModifyVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessEndpointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessEndpointPolicyWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessEndpointPolicyInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessEndpointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessGroup(input *ec2.ModifyVerifiedAccessGroupInput) (*ec2.ModifyVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessGroupPolicy(input *ec2.ModifyVerifiedAccessGroupPolicyInput) (*ec2.ModifyVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessGroupPolicyWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessGroupPolicyInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessGroupInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessInstance(input *ec2.ModifyVerifiedAccessInstanceInput) (*ec2.ModifyVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessInstanceLoggingConfiguration(input *ec2.ModifyVerifiedAccessInstanceLoggingConfigurationInput) (*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessInstanceLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessInstanceLoggingConfigurationWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessInstanceLoggingConfigurationInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessInstanceLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessTrustProvider(input *ec2.ModifyVerifiedAccessTrustProviderInput) (*ec2.ModifyVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.EC2API.ModifyVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVolume(input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeOutput), nil
	}
	out, err := c.EC2API.ModifyVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVolumeAttribute(input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyVolumeAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVolumeAttributeWithContext(ctx context.Context, input *ec2.ModifyVolumeAttributeInput, opts ...request.Option) (*ec2.ModifyVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyVolumeAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVolumeWithContext(ctx context.Context, input *ec2.ModifyVolumeInput, opts ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeOutput), nil
	}
	out, err := c.EC2API.ModifyVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcAttribute(input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyVpcAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcAttributeWithContext(ctx context.Context, input *ec2.ModifyVpcAttributeInput, opts ...request.Option) (*ec2.ModifyVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcAttributeOutput), nil
	}
	out, err := c.EC2API.ModifyVpcAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpoint(input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointConnectionNotification(input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointConnectionNotification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointConnectionNotificationWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput, opts ...request.Option) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointConnectionNotificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServiceConfiguration(input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServiceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServiceConfigurationWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServiceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServicePayerResponsibility(input *ec2.ModifyVpcEndpointServicePayerResponsibilityInput) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServicePayerResponsibility(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServicePayerResponsibilityWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServicePayerResponsibilityInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServicePayerResponsibilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServicePermissions(input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServicePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointServicePermissionsWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointServicePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcEndpointWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointInput, opts ...request.Option) (*ec2.ModifyVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointOutput), nil
	}
	out, err := c.EC2API.ModifyVpcEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcPeeringConnectionOptions(input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcPeeringConnectionOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpcPeeringConnectionOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcPeeringConnectionOptionsWithContext(ctx context.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput, opts ...request.Option) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcPeeringConnectionOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpcPeeringConnectionOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcTenancy(input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcTenancyOutput), nil
	}
	out, err := c.EC2API.ModifyVpcTenancy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpcTenancyWithContext(ctx context.Context, input *ec2.ModifyVpcTenancyInput, opts ...request.Option) (*ec2.ModifyVpcTenancyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcTenancyOutput), nil
	}
	out, err := c.EC2API.ModifyVpcTenancyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnConnection(input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOutput), nil
	}
	out, err := c.EC2API.ModifyVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnConnectionOptions(input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpnConnectionOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnConnectionOptionsWithContext(ctx context.Context, input *ec2.ModifyVpnConnectionOptionsInput, opts ...request.Option) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpnConnectionOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnConnectionWithContext(ctx context.Context, input *ec2.ModifyVpnConnectionInput, opts ...request.Option) (*ec2.ModifyVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOutput), nil
	}
	out, err := c.EC2API.ModifyVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnTunnelCertificate(input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelCertificateOutput), nil
	}
	out, err := c.EC2API.ModifyVpnTunnelCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnTunnelCertificateWithContext(ctx context.Context, input *ec2.ModifyVpnTunnelCertificateInput, opts ...request.Option) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelCertificateOutput), nil
	}
	out, err := c.EC2API.ModifyVpnTunnelCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnTunnelOptions(input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpnTunnelOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ModifyVpnTunnelOptionsWithContext(ctx context.Context, input *ec2.ModifyVpnTunnelOptionsInput, opts ...request.Option) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelOptionsOutput), nil
	}
	out, err := c.EC2API.ModifyVpnTunnelOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MonitorInstances(input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MonitorInstancesOutput), nil
	}
	out, err := c.EC2API.MonitorInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MonitorInstancesWithContext(ctx context.Context, input *ec2.MonitorInstancesInput, opts ...request.Option) (*ec2.MonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MonitorInstancesOutput), nil
	}
	out, err := c.EC2API.MonitorInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MoveAddressToVpc(input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveAddressToVpcOutput), nil
	}
	out, err := c.EC2API.MoveAddressToVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MoveAddressToVpcWithContext(ctx context.Context, input *ec2.MoveAddressToVpcInput, opts ...request.Option) (*ec2.MoveAddressToVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveAddressToVpcOutput), nil
	}
	out, err := c.EC2API.MoveAddressToVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MoveByoipCidrToIpam(input *ec2.MoveByoipCidrToIpamInput) (*ec2.MoveByoipCidrToIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveByoipCidrToIpamOutput), nil
	}
	out, err := c.EC2API.MoveByoipCidrToIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) MoveByoipCidrToIpamWithContext(ctx context.Context, input *ec2.MoveByoipCidrToIpamInput, opts ...request.Option) (*ec2.MoveByoipCidrToIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveByoipCidrToIpamOutput), nil
	}
	out, err := c.EC2API.MoveByoipCidrToIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionByoipCidr(input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionByoipCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionByoipCidrWithContext(ctx context.Context, input *ec2.ProvisionByoipCidrInput, opts ...request.Option) (*ec2.ProvisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionByoipCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionIpamPoolCidr(input *ec2.ProvisionIpamPoolCidrInput) (*ec2.ProvisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionIpamPoolCidrWithContext(ctx context.Context, input *ec2.ProvisionIpamPoolCidrInput, opts ...request.Option) (*ec2.ProvisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionIpamPoolCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionPublicIpv4PoolCidr(input *ec2.ProvisionPublicIpv4PoolCidrInput) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionPublicIpv4PoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ProvisionPublicIpv4PoolCidrWithContext(ctx context.Context, input *ec2.ProvisionPublicIpv4PoolCidrInput, opts ...request.Option) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.EC2API.ProvisionPublicIpv4PoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseHostReservation(input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseHostReservationOutput), nil
	}
	out, err := c.EC2API.PurchaseHostReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseHostReservationWithContext(ctx context.Context, input *ec2.PurchaseHostReservationInput, opts ...request.Option) (*ec2.PurchaseHostReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseHostReservationOutput), nil
	}
	out, err := c.EC2API.PurchaseHostReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseReservedInstancesOffering(input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseReservedInstancesOfferingOutput), nil
	}
	out, err := c.EC2API.PurchaseReservedInstancesOffering(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseReservedInstancesOfferingWithContext(ctx context.Context, input *ec2.PurchaseReservedInstancesOfferingInput, opts ...request.Option) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseReservedInstancesOfferingOutput), nil
	}
	out, err := c.EC2API.PurchaseReservedInstancesOfferingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseScheduledInstances(input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.PurchaseScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) PurchaseScheduledInstancesWithContext(ctx context.Context, input *ec2.PurchaseScheduledInstancesInput, opts ...request.Option) (*ec2.PurchaseScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.PurchaseScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RebootInstances(input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RebootInstancesOutput), nil
	}
	out, err := c.EC2API.RebootInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RebootInstancesWithContext(ctx context.Context, input *ec2.RebootInstancesInput, opts ...request.Option) (*ec2.RebootInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RebootInstancesOutput), nil
	}
	out, err := c.EC2API.RebootInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterImage(input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterImageOutput), nil
	}
	out, err := c.EC2API.RegisterImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterImageWithContext(ctx context.Context, input *ec2.RegisterImageInput, opts ...request.Option) (*ec2.RegisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterImageOutput), nil
	}
	out, err := c.EC2API.RegisterImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterInstanceEventNotificationAttributes(input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.RegisterInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.RegisterInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterTransitGatewayMulticastGroupMembers(input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.EC2API.RegisterTransitGatewayMulticastGroupMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterTransitGatewayMulticastGroupMembersWithContext(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput, opts ...request.Option) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.EC2API.RegisterTransitGatewayMulticastGroupMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterTransitGatewayMulticastGroupSources(input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.EC2API.RegisterTransitGatewayMulticastGroupSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RegisterTransitGatewayMulticastGroupSourcesWithContext(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput, opts ...request.Option) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.EC2API.RegisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayMulticastDomainAssociations(input *ec2.RejectTransitGatewayMulticastDomainAssociationsInput) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.RejectTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayPeeringAttachment(input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayVpcAttachment(input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.EC2API.RejectTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectVpcEndpointConnections(input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.RejectVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.RejectVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.RejectVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectVpcPeeringConnection(input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.RejectVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RejectVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.RejectVpcPeeringConnectionInput, opts ...request.Option) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcPeeringConnectionOutput), nil
	}
	out, err := c.EC2API.RejectVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseAddress(input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseAddressOutput), nil
	}
	out, err := c.EC2API.ReleaseAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseAddressWithContext(ctx context.Context, input *ec2.ReleaseAddressInput, opts ...request.Option) (*ec2.ReleaseAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseAddressOutput), nil
	}
	out, err := c.EC2API.ReleaseAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseHosts(input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseHostsOutput), nil
	}
	out, err := c.EC2API.ReleaseHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseHostsWithContext(ctx context.Context, input *ec2.ReleaseHostsInput, opts ...request.Option) (*ec2.ReleaseHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseHostsOutput), nil
	}
	out, err := c.EC2API.ReleaseHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseIpamPoolAllocation(input *ec2.ReleaseIpamPoolAllocationInput) (*ec2.ReleaseIpamPoolAllocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseIpamPoolAllocationOutput), nil
	}
	out, err := c.EC2API.ReleaseIpamPoolAllocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReleaseIpamPoolAllocationWithContext(ctx context.Context, input *ec2.ReleaseIpamPoolAllocationInput, opts ...request.Option) (*ec2.ReleaseIpamPoolAllocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseIpamPoolAllocationOutput), nil
	}
	out, err := c.EC2API.ReleaseIpamPoolAllocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceIamInstanceProfileAssociation(input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceIamInstanceProfileAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceIamInstanceProfileAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceIamInstanceProfileAssociationWithContext(ctx context.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput, opts ...request.Option) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceIamInstanceProfileAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceIamInstanceProfileAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceNetworkAclAssociation(input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceNetworkAclAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceNetworkAclAssociationWithContext(ctx context.Context, input *ec2.ReplaceNetworkAclAssociationInput, opts ...request.Option) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceNetworkAclAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceNetworkAclEntry(input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.ReplaceNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceNetworkAclEntryWithContext(ctx context.Context, input *ec2.ReplaceNetworkAclEntryInput, opts ...request.Option) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclEntryOutput), nil
	}
	out, err := c.EC2API.ReplaceNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceRoute(input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteOutput), nil
	}
	out, err := c.EC2API.ReplaceRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceRouteTableAssociation(input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteTableAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceRouteTableAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceRouteTableAssociationWithContext(ctx context.Context, input *ec2.ReplaceRouteTableAssociationInput, opts ...request.Option) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteTableAssociationOutput), nil
	}
	out, err := c.EC2API.ReplaceRouteTableAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceRouteWithContext(ctx context.Context, input *ec2.ReplaceRouteInput, opts ...request.Option) (*ec2.ReplaceRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteOutput), nil
	}
	out, err := c.EC2API.ReplaceRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceTransitGatewayRoute(input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.ReplaceTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReplaceTransitGatewayRouteWithContext(ctx context.Context, input *ec2.ReplaceTransitGatewayRouteInput, opts ...request.Option) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceTransitGatewayRouteOutput), nil
	}
	out, err := c.EC2API.ReplaceTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReportInstanceStatus(input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReportInstanceStatusOutput), nil
	}
	out, err := c.EC2API.ReportInstanceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ReportInstanceStatusWithContext(ctx context.Context, input *ec2.ReportInstanceStatusInput, opts ...request.Option) (*ec2.ReportInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReportInstanceStatusOutput), nil
	}
	out, err := c.EC2API.ReportInstanceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RequestSpotFleet(input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotFleetOutput), nil
	}
	out, err := c.EC2API.RequestSpotFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RequestSpotFleetWithContext(ctx context.Context, input *ec2.RequestSpotFleetInput, opts ...request.Option) (*ec2.RequestSpotFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotFleetOutput), nil
	}
	out, err := c.EC2API.RequestSpotFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RequestSpotInstances(input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotInstancesOutput), nil
	}
	out, err := c.EC2API.RequestSpotInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RequestSpotInstancesWithContext(ctx context.Context, input *ec2.RequestSpotInstancesInput, opts ...request.Option) (*ec2.RequestSpotInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotInstancesOutput), nil
	}
	out, err := c.EC2API.RequestSpotInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetAddressAttribute(input *ec2.ResetAddressAttributeInput) (*ec2.ResetAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetAddressAttributeOutput), nil
	}
	out, err := c.EC2API.ResetAddressAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetAddressAttributeWithContext(ctx context.Context, input *ec2.ResetAddressAttributeInput, opts ...request.Option) (*ec2.ResetAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetAddressAttributeOutput), nil
	}
	out, err := c.EC2API.ResetAddressAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetEbsDefaultKmsKeyId(input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.ResetEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.ResetEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetFpgaImageAttribute(input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.ResetFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetFpgaImageAttributeWithContext(ctx context.Context, input *ec2.ResetFpgaImageAttributeInput, opts ...request.Option) (*ec2.ResetFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.ResetFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetImageAttribute(input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetImageAttributeOutput), nil
	}
	out, err := c.EC2API.ResetImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetImageAttributeWithContext(ctx context.Context, input *ec2.ResetImageAttributeInput, opts ...request.Option) (*ec2.ResetImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetImageAttributeOutput), nil
	}
	out, err := c.EC2API.ResetImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetInstanceAttribute(input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.ResetInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetInstanceAttributeWithContext(ctx context.Context, input *ec2.ResetInstanceAttributeInput, opts ...request.Option) (*ec2.ResetInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.ResetInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetNetworkInterfaceAttribute(input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.ResetNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.ResetNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.ResetNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetSnapshotAttribute(input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.ResetSnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ResetSnapshotAttributeWithContext(ctx context.Context, input *ec2.ResetSnapshotAttributeInput, opts ...request.Option) (*ec2.ResetSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.ResetSnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreAddressToClassic(input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreAddressToClassicOutput), nil
	}
	out, err := c.EC2API.RestoreAddressToClassic(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreAddressToClassicWithContext(ctx context.Context, input *ec2.RestoreAddressToClassicInput, opts ...request.Option) (*ec2.RestoreAddressToClassicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreAddressToClassicOutput), nil
	}
	out, err := c.EC2API.RestoreAddressToClassicWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreImageFromRecycleBin(input *ec2.RestoreImageFromRecycleBinInput) (*ec2.RestoreImageFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreImageFromRecycleBinOutput), nil
	}
	out, err := c.EC2API.RestoreImageFromRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreImageFromRecycleBinWithContext(ctx context.Context, input *ec2.RestoreImageFromRecycleBinInput, opts ...request.Option) (*ec2.RestoreImageFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreImageFromRecycleBinOutput), nil
	}
	out, err := c.EC2API.RestoreImageFromRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreManagedPrefixListVersion(input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreManagedPrefixListVersionOutput), nil
	}
	out, err := c.EC2API.RestoreManagedPrefixListVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreManagedPrefixListVersionWithContext(ctx context.Context, input *ec2.RestoreManagedPrefixListVersionInput, opts ...request.Option) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreManagedPrefixListVersionOutput), nil
	}
	out, err := c.EC2API.RestoreManagedPrefixListVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreSnapshotFromRecycleBin(input *ec2.RestoreSnapshotFromRecycleBinInput) (*ec2.RestoreSnapshotFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotFromRecycleBinOutput), nil
	}
	out, err := c.EC2API.RestoreSnapshotFromRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreSnapshotFromRecycleBinWithContext(ctx context.Context, input *ec2.RestoreSnapshotFromRecycleBinInput, opts ...request.Option) (*ec2.RestoreSnapshotFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotFromRecycleBinOutput), nil
	}
	out, err := c.EC2API.RestoreSnapshotFromRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreSnapshotTier(input *ec2.RestoreSnapshotTierInput) (*ec2.RestoreSnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotTierOutput), nil
	}
	out, err := c.EC2API.RestoreSnapshotTier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RestoreSnapshotTierWithContext(ctx context.Context, input *ec2.RestoreSnapshotTierInput, opts ...request.Option) (*ec2.RestoreSnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotTierOutput), nil
	}
	out, err := c.EC2API.RestoreSnapshotTierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeClientVpnIngress(input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeClientVpnIngressOutput), nil
	}
	out, err := c.EC2API.RevokeClientVpnIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeClientVpnIngressWithContext(ctx context.Context, input *ec2.RevokeClientVpnIngressInput, opts ...request.Option) (*ec2.RevokeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeClientVpnIngressOutput), nil
	}
	out, err := c.EC2API.RevokeClientVpnIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeSecurityGroupEgress(input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupEgressOutput), nil
	}
	out, err := c.EC2API.RevokeSecurityGroupEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeSecurityGroupEgressWithContext(ctx context.Context, input *ec2.RevokeSecurityGroupEgressInput, opts ...request.Option) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupEgressOutput), nil
	}
	out, err := c.EC2API.RevokeSecurityGroupEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupIngressOutput), nil
	}
	out, err := c.EC2API.RevokeSecurityGroupIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RevokeSecurityGroupIngressWithContext(ctx context.Context, input *ec2.RevokeSecurityGroupIngressInput, opts ...request.Option) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupIngressOutput), nil
	}
	out, err := c.EC2API.RevokeSecurityGroupIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RunScheduledInstances(input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RunScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.RunScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) RunScheduledInstancesWithContext(ctx context.Context, input *ec2.RunScheduledInstancesInput, opts ...request.Option) (*ec2.RunScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RunScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.RunScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchLocalGatewayRoutes(input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchLocalGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchLocalGatewayRoutesPages(input *ec2.SearchLocalGatewayRoutesInput, fn func(*ec2.SearchLocalGatewayRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchLocalGatewayRoutesOutput{}
	fnCacher := func(out *ec2.SearchLocalGatewayRoutesOutput, more bool) bool {
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
	if err := c.EC2API.SearchLocalGatewayRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchLocalGatewayRoutesPagesWithContext(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput, fn func(*ec2.SearchLocalGatewayRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchLocalGatewayRoutesOutput{}
	fnCacher := func(out *ec2.SearchLocalGatewayRoutesOutput, more bool) bool {
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
	if err := c.EC2API.SearchLocalGatewayRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchLocalGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput, opts ...request.Option) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchLocalGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayMulticastGroups(input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayMulticastGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayMulticastGroupsPages(input *ec2.SearchTransitGatewayMulticastGroupsInput, fn func(*ec2.SearchTransitGatewayMulticastGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchTransitGatewayMulticastGroupsOutput{}
	fnCacher := func(out *ec2.SearchTransitGatewayMulticastGroupsOutput, more bool) bool {
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
	if err := c.EC2API.SearchTransitGatewayMulticastGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchTransitGatewayMulticastGroupsPagesWithContext(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput, fn func(*ec2.SearchTransitGatewayMulticastGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchTransitGatewayMulticastGroupsOutput{}
	fnCacher := func(out *ec2.SearchTransitGatewayMulticastGroupsOutput, more bool) bool {
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
	if err := c.EC2API.SearchTransitGatewayMulticastGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchTransitGatewayMulticastGroupsWithContext(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput, opts ...request.Option) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayMulticastGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayRoutes(input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchTransitGatewayRoutesInput, opts ...request.Option) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SendDiagnosticInterrupt(input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SendDiagnosticInterruptOutput), nil
	}
	out, err := c.EC2API.SendDiagnosticInterrupt(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SendDiagnosticInterruptWithContext(ctx context.Context, input *ec2.SendDiagnosticInterruptInput, opts ...request.Option) (*ec2.SendDiagnosticInterruptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SendDiagnosticInterruptOutput), nil
	}
	out, err := c.EC2API.SendDiagnosticInterruptWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartInstances(input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartInstancesOutput), nil
	}
	out, err := c.EC2API.StartInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartInstancesWithContext(ctx context.Context, input *ec2.StartInstancesInput, opts ...request.Option) (*ec2.StartInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartInstancesOutput), nil
	}
	out, err := c.EC2API.StartInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartNetworkInsightsAccessScopeAnalysis(input *ec2.StartNetworkInsightsAccessScopeAnalysisInput) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.EC2API.StartNetworkInsightsAccessScopeAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartNetworkInsightsAccessScopeAnalysisWithContext(ctx context.Context, input *ec2.StartNetworkInsightsAccessScopeAnalysisInput, opts ...request.Option) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.EC2API.StartNetworkInsightsAccessScopeAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartNetworkInsightsAnalysis(input *ec2.StartNetworkInsightsAnalysisInput) (*ec2.StartNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.EC2API.StartNetworkInsightsAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartNetworkInsightsAnalysisWithContext(ctx context.Context, input *ec2.StartNetworkInsightsAnalysisInput, opts ...request.Option) (*ec2.StartNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.EC2API.StartNetworkInsightsAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartVpcEndpointServicePrivateDnsVerification(input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput), nil
	}
	out, err := c.EC2API.StartVpcEndpointServicePrivateDnsVerification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StartVpcEndpointServicePrivateDnsVerificationWithContext(ctx context.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput, opts ...request.Option) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput), nil
	}
	out, err := c.EC2API.StartVpcEndpointServicePrivateDnsVerificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StopInstances(input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StopInstancesOutput), nil
	}
	out, err := c.EC2API.StopInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) StopInstancesWithContext(ctx context.Context, input *ec2.StopInstancesInput, opts ...request.Option) (*ec2.StopInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StopInstancesOutput), nil
	}
	out, err := c.EC2API.StopInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) TerminateClientVpnConnections(input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.TerminateClientVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) TerminateClientVpnConnectionsWithContext(ctx context.Context, input *ec2.TerminateClientVpnConnectionsInput, opts ...request.Option) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.TerminateClientVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateInstancesOutput), nil
	}
	out, err := c.EC2API.TerminateInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) TerminateInstancesWithContext(ctx context.Context, input *ec2.TerminateInstancesInput, opts ...request.Option) (*ec2.TerminateInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateInstancesOutput), nil
	}
	out, err := c.EC2API.TerminateInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnassignIpv6Addresses(input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignIpv6AddressesOutput), nil
	}
	out, err := c.EC2API.UnassignIpv6Addresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnassignIpv6AddressesWithContext(ctx context.Context, input *ec2.UnassignIpv6AddressesInput, opts ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignIpv6AddressesOutput), nil
	}
	out, err := c.EC2API.UnassignIpv6AddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnassignPrivateIpAddresses(input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignPrivateIpAddressesOutput), nil
	}
	out, err := c.EC2API.UnassignPrivateIpAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnassignPrivateIpAddressesWithContext(ctx context.Context, input *ec2.UnassignPrivateIpAddressesInput, opts ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignPrivateIpAddressesOutput), nil
	}
	out, err := c.EC2API.UnassignPrivateIpAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnmonitorInstances(input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnmonitorInstancesOutput), nil
	}
	out, err := c.EC2API.UnmonitorInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UnmonitorInstancesWithContext(ctx context.Context, input *ec2.UnmonitorInstancesInput, opts ...request.Option) (*ec2.UnmonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnmonitorInstancesOutput), nil
	}
	out, err := c.EC2API.UnmonitorInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UpdateSecurityGroupRuleDescriptionsEgress(input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput), nil
	}
	out, err := c.EC2API.UpdateSecurityGroupRuleDescriptionsEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UpdateSecurityGroupRuleDescriptionsEgressWithContext(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, opts ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput), nil
	}
	out, err := c.EC2API.UpdateSecurityGroupRuleDescriptionsEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UpdateSecurityGroupRuleDescriptionsIngress(input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput), nil
	}
	out, err := c.EC2API.UpdateSecurityGroupRuleDescriptionsIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) UpdateSecurityGroupRuleDescriptionsIngressWithContext(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, opts ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput), nil
	}
	out, err := c.EC2API.UpdateSecurityGroupRuleDescriptionsIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) WithdrawByoipCidr(input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.WithdrawByoipCidrOutput), nil
	}
	out, err := c.EC2API.WithdrawByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) WithdrawByoipCidrWithContext(ctx context.Context, input *ec2.WithdrawByoipCidrInput, opts ...request.Option) (*ec2.WithdrawByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.WithdrawByoipCidrOutput), nil
	}
	out, err := c.EC2API.WithdrawByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}


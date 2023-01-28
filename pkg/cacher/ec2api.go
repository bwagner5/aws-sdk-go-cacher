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

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
)

func (c *Client) AcceptAddressTransfer(input *ec2.AcceptAddressTransferInput) (*ec2.AcceptAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptAddressTransferOutput), nil
	}
	out, err := c.client.AcceptAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptAddressTransferWithContext(ctx context.Context, input *ec2.AcceptAddressTransferInput, opts ...request.Option) (*ec2.AcceptAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptAddressTransferOutput), nil
	}
	out, err := c.client.AcceptAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptReservedInstancesExchangeQuote(input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.client.AcceptReservedInstancesExchangeQuote(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptReservedInstancesExchangeQuoteWithContext(ctx context.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput, opts ...request.Option) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.client.AcceptReservedInstancesExchangeQuoteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayMulticastDomainAssociations(input *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayPeeringAttachment(input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayVpcAttachment(input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.AcceptTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptVpcEndpointConnections(input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.AcceptVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.AcceptVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.AcceptVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptVpcPeeringConnection(input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.AcceptVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AcceptVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.AcceptVpcPeeringConnectionInput, opts ...request.Option) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AcceptVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.AcceptVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AdvertiseByoipCidr(input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AdvertiseByoipCidrOutput), nil
	}
	out, err := c.client.AdvertiseByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AdvertiseByoipCidrWithContext(ctx context.Context, input *ec2.AdvertiseByoipCidrInput, opts ...request.Option) (*ec2.AdvertiseByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AdvertiseByoipCidrOutput), nil
	}
	out, err := c.client.AdvertiseByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateAddress(input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateAddressOutput), nil
	}
	out, err := c.client.AllocateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateAddressWithContext(ctx context.Context, input *ec2.AllocateAddressInput, opts ...request.Option) (*ec2.AllocateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateAddressOutput), nil
	}
	out, err := c.client.AllocateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateHosts(input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateHostsOutput), nil
	}
	out, err := c.client.AllocateHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateHostsWithContext(ctx context.Context, input *ec2.AllocateHostsInput, opts ...request.Option) (*ec2.AllocateHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateHostsOutput), nil
	}
	out, err := c.client.AllocateHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateIpamPoolCidr(input *ec2.AllocateIpamPoolCidrInput) (*ec2.AllocateIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateIpamPoolCidrOutput), nil
	}
	out, err := c.client.AllocateIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AllocateIpamPoolCidrWithContext(ctx context.Context, input *ec2.AllocateIpamPoolCidrInput, opts ...request.Option) (*ec2.AllocateIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AllocateIpamPoolCidrOutput), nil
	}
	out, err := c.client.AllocateIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetwork(input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.ApplySecurityGroupsToClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.ApplySecurityGroupsToClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssignIpv6Addresses(input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignIpv6AddressesOutput), nil
	}
	out, err := c.client.AssignIpv6Addresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssignIpv6AddressesWithContext(ctx context.Context, input *ec2.AssignIpv6AddressesInput, opts ...request.Option) (*ec2.AssignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignIpv6AddressesOutput), nil
	}
	out, err := c.client.AssignIpv6AddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssignPrivateIpAddresses(input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignPrivateIpAddressesOutput), nil
	}
	out, err := c.client.AssignPrivateIpAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssignPrivateIpAddressesWithContext(ctx context.Context, input *ec2.AssignPrivateIpAddressesInput, opts ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssignPrivateIpAddressesOutput), nil
	}
	out, err := c.client.AssignPrivateIpAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateAddress(input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateAddressOutput), nil
	}
	out, err := c.client.AssociateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateAddressWithContext(ctx context.Context, input *ec2.AssociateAddressInput, opts ...request.Option) (*ec2.AssociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateAddressOutput), nil
	}
	out, err := c.client.AssociateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateClientVpnTargetNetwork(input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.AssociateClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.AssociateClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.AssociateClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateDhcpOptions(input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateDhcpOptionsOutput), nil
	}
	out, err := c.client.AssociateDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateDhcpOptionsWithContext(ctx context.Context, input *ec2.AssociateDhcpOptionsInput, opts ...request.Option) (*ec2.AssociateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateDhcpOptionsOutput), nil
	}
	out, err := c.client.AssociateDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateEnclaveCertificateIamRole(input *ec2.AssociateEnclaveCertificateIamRoleInput) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.client.AssociateEnclaveCertificateIamRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateEnclaveCertificateIamRoleWithContext(ctx context.Context, input *ec2.AssociateEnclaveCertificateIamRoleInput, opts ...request.Option) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.client.AssociateEnclaveCertificateIamRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateIamInstanceProfile(input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateIamInstanceProfileOutput), nil
	}
	out, err := c.client.AssociateIamInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateIamInstanceProfileWithContext(ctx context.Context, input *ec2.AssociateIamInstanceProfileInput, opts ...request.Option) (*ec2.AssociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateIamInstanceProfileOutput), nil
	}
	out, err := c.client.AssociateIamInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateInstanceEventWindow(input *ec2.AssociateInstanceEventWindowInput) (*ec2.AssociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateInstanceEventWindowOutput), nil
	}
	out, err := c.client.AssociateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateInstanceEventWindowWithContext(ctx context.Context, input *ec2.AssociateInstanceEventWindowInput, opts ...request.Option) (*ec2.AssociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateInstanceEventWindowOutput), nil
	}
	out, err := c.client.AssociateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateRouteTable(input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateRouteTableOutput), nil
	}
	out, err := c.client.AssociateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateRouteTableWithContext(ctx context.Context, input *ec2.AssociateRouteTableInput, opts ...request.Option) (*ec2.AssociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateRouteTableOutput), nil
	}
	out, err := c.client.AssociateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateSubnetCidrBlock(input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateSubnetCidrBlockOutput), nil
	}
	out, err := c.client.AssociateSubnetCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateSubnetCidrBlockWithContext(ctx context.Context, input *ec2.AssociateSubnetCidrBlockInput, opts ...request.Option) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateSubnetCidrBlockOutput), nil
	}
	out, err := c.client.AssociateSubnetCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayMulticastDomain(input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayPolicyTable(input *ec2.AssociateTransitGatewayPolicyTableInput) (*ec2.AssociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.AssociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayRouteTable(input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.AssociateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.AssociateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTrunkInterface(input *ec2.AssociateTrunkInterfaceInput) (*ec2.AssociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTrunkInterfaceOutput), nil
	}
	out, err := c.client.AssociateTrunkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateTrunkInterfaceWithContext(ctx context.Context, input *ec2.AssociateTrunkInterfaceInput, opts ...request.Option) (*ec2.AssociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateTrunkInterfaceOutput), nil
	}
	out, err := c.client.AssociateTrunkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateVpcCidrBlock(input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateVpcCidrBlockOutput), nil
	}
	out, err := c.client.AssociateVpcCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AssociateVpcCidrBlockWithContext(ctx context.Context, input *ec2.AssociateVpcCidrBlockInput, opts ...request.Option) (*ec2.AssociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AssociateVpcCidrBlockOutput), nil
	}
	out, err := c.client.AssociateVpcCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachClassicLinkVpc(input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachClassicLinkVpcOutput), nil
	}
	out, err := c.client.AttachClassicLinkVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachClassicLinkVpcWithContext(ctx context.Context, input *ec2.AttachClassicLinkVpcInput, opts ...request.Option) (*ec2.AttachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachClassicLinkVpcOutput), nil
	}
	out, err := c.client.AttachClassicLinkVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachInternetGateway(input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachInternetGatewayOutput), nil
	}
	out, err := c.client.AttachInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachInternetGatewayWithContext(ctx context.Context, input *ec2.AttachInternetGatewayInput, opts ...request.Option) (*ec2.AttachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachInternetGatewayOutput), nil
	}
	out, err := c.client.AttachInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachNetworkInterface(input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachNetworkInterfaceOutput), nil
	}
	out, err := c.client.AttachNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, input *ec2.AttachNetworkInterfaceInput, opts ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachNetworkInterfaceOutput), nil
	}
	out, err := c.client.AttachNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachVerifiedAccessTrustProvider(input *ec2.AttachVerifiedAccessTrustProviderInput) (*ec2.AttachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.AttachVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.AttachVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.AttachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.AttachVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachVpnGateway(input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVpnGatewayOutput), nil
	}
	out, err := c.client.AttachVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AttachVpnGatewayWithContext(ctx context.Context, input *ec2.AttachVpnGatewayInput, opts ...request.Option) (*ec2.AttachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AttachVpnGatewayOutput), nil
	}
	out, err := c.client.AttachVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeClientVpnIngress(input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeClientVpnIngressOutput), nil
	}
	out, err := c.client.AuthorizeClientVpnIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeClientVpnIngressWithContext(ctx context.Context, input *ec2.AuthorizeClientVpnIngressInput, opts ...request.Option) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeClientVpnIngressOutput), nil
	}
	out, err := c.client.AuthorizeClientVpnIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeSecurityGroupEgress(input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupEgressOutput), nil
	}
	out, err := c.client.AuthorizeSecurityGroupEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeSecurityGroupEgressWithContext(ctx context.Context, input *ec2.AuthorizeSecurityGroupEgressInput, opts ...request.Option) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupEgressOutput), nil
	}
	out, err := c.client.AuthorizeSecurityGroupEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupIngressOutput), nil
	}
	out, err := c.client.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) AuthorizeSecurityGroupIngressWithContext(ctx context.Context, input *ec2.AuthorizeSecurityGroupIngressInput, opts ...request.Option) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.AuthorizeSecurityGroupIngressOutput), nil
	}
	out, err := c.client.AuthorizeSecurityGroupIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) BundleInstance(input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.BundleInstanceOutput), nil
	}
	out, err := c.client.BundleInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) BundleInstanceWithContext(ctx context.Context, input *ec2.BundleInstanceInput, opts ...request.Option) (*ec2.BundleInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.BundleInstanceOutput), nil
	}
	out, err := c.client.BundleInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelBundleTask(input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelBundleTaskOutput), nil
	}
	out, err := c.client.CancelBundleTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelBundleTaskWithContext(ctx context.Context, input *ec2.CancelBundleTaskInput, opts ...request.Option) (*ec2.CancelBundleTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelBundleTaskOutput), nil
	}
	out, err := c.client.CancelBundleTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelCapacityReservation(input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationOutput), nil
	}
	out, err := c.client.CancelCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelCapacityReservationFleets(input *ec2.CancelCapacityReservationFleetsInput) (*ec2.CancelCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationFleetsOutput), nil
	}
	out, err := c.client.CancelCapacityReservationFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelCapacityReservationFleetsWithContext(ctx context.Context, input *ec2.CancelCapacityReservationFleetsInput, opts ...request.Option) (*ec2.CancelCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationFleetsOutput), nil
	}
	out, err := c.client.CancelCapacityReservationFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelCapacityReservationWithContext(ctx context.Context, input *ec2.CancelCapacityReservationInput, opts ...request.Option) (*ec2.CancelCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelCapacityReservationOutput), nil
	}
	out, err := c.client.CancelCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelConversionTask(input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelConversionTaskOutput), nil
	}
	out, err := c.client.CancelConversionTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelConversionTaskWithContext(ctx context.Context, input *ec2.CancelConversionTaskInput, opts ...request.Option) (*ec2.CancelConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelConversionTaskOutput), nil
	}
	out, err := c.client.CancelConversionTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelExportTask(input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelExportTaskOutput), nil
	}
	out, err := c.client.CancelExportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelExportTaskWithContext(ctx context.Context, input *ec2.CancelExportTaskInput, opts ...request.Option) (*ec2.CancelExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelExportTaskOutput), nil
	}
	out, err := c.client.CancelExportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelImageLaunchPermission(input *ec2.CancelImageLaunchPermissionInput) (*ec2.CancelImageLaunchPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImageLaunchPermissionOutput), nil
	}
	out, err := c.client.CancelImageLaunchPermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelImageLaunchPermissionWithContext(ctx context.Context, input *ec2.CancelImageLaunchPermissionInput, opts ...request.Option) (*ec2.CancelImageLaunchPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImageLaunchPermissionOutput), nil
	}
	out, err := c.client.CancelImageLaunchPermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelImportTask(input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImportTaskOutput), nil
	}
	out, err := c.client.CancelImportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelImportTaskWithContext(ctx context.Context, input *ec2.CancelImportTaskInput, opts ...request.Option) (*ec2.CancelImportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelImportTaskOutput), nil
	}
	out, err := c.client.CancelImportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelReservedInstancesListing(input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelReservedInstancesListingOutput), nil
	}
	out, err := c.client.CancelReservedInstancesListing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelReservedInstancesListingWithContext(ctx context.Context, input *ec2.CancelReservedInstancesListingInput, opts ...request.Option) (*ec2.CancelReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelReservedInstancesListingOutput), nil
	}
	out, err := c.client.CancelReservedInstancesListingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelSpotFleetRequests(input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotFleetRequestsOutput), nil
	}
	out, err := c.client.CancelSpotFleetRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelSpotFleetRequestsWithContext(ctx context.Context, input *ec2.CancelSpotFleetRequestsInput, opts ...request.Option) (*ec2.CancelSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotFleetRequestsOutput), nil
	}
	out, err := c.client.CancelSpotFleetRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelSpotInstanceRequests(input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotInstanceRequestsOutput), nil
	}
	out, err := c.client.CancelSpotInstanceRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CancelSpotInstanceRequestsWithContext(ctx context.Context, input *ec2.CancelSpotInstanceRequestsInput, opts ...request.Option) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CancelSpotInstanceRequestsOutput), nil
	}
	out, err := c.client.CancelSpotInstanceRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ConfirmProductInstance(input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ConfirmProductInstanceOutput), nil
	}
	out, err := c.client.ConfirmProductInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ConfirmProductInstanceWithContext(ctx context.Context, input *ec2.ConfirmProductInstanceInput, opts ...request.Option) (*ec2.ConfirmProductInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ConfirmProductInstanceOutput), nil
	}
	out, err := c.client.ConfirmProductInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopyFpgaImage(input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyFpgaImageOutput), nil
	}
	out, err := c.client.CopyFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopyFpgaImageWithContext(ctx context.Context, input *ec2.CopyFpgaImageInput, opts ...request.Option) (*ec2.CopyFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyFpgaImageOutput), nil
	}
	out, err := c.client.CopyFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopyImage(input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyImageOutput), nil
	}
	out, err := c.client.CopyImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopyImageWithContext(ctx context.Context, input *ec2.CopyImageInput, opts ...request.Option) (*ec2.CopyImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopyImageOutput), nil
	}
	out, err := c.client.CopyImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopySnapshot(input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopySnapshotOutput), nil
	}
	out, err := c.client.CopySnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CopySnapshotWithContext(ctx context.Context, input *ec2.CopySnapshotInput, opts ...request.Option) (*ec2.CopySnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CopySnapshotOutput), nil
	}
	out, err := c.client.CopySnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCapacityReservation(input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationOutput), nil
	}
	out, err := c.client.CreateCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCapacityReservationFleet(input *ec2.CreateCapacityReservationFleetInput) (*ec2.CreateCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationFleetOutput), nil
	}
	out, err := c.client.CreateCapacityReservationFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCapacityReservationFleetWithContext(ctx context.Context, input *ec2.CreateCapacityReservationFleetInput, opts ...request.Option) (*ec2.CreateCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationFleetOutput), nil
	}
	out, err := c.client.CreateCapacityReservationFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCapacityReservationWithContext(ctx context.Context, input *ec2.CreateCapacityReservationInput, opts ...request.Option) (*ec2.CreateCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCapacityReservationOutput), nil
	}
	out, err := c.client.CreateCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCarrierGateway(input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCarrierGatewayOutput), nil
	}
	out, err := c.client.CreateCarrierGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCarrierGatewayWithContext(ctx context.Context, input *ec2.CreateCarrierGatewayInput, opts ...request.Option) (*ec2.CreateCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCarrierGatewayOutput), nil
	}
	out, err := c.client.CreateCarrierGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateClientVpnEndpoint(input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnEndpointOutput), nil
	}
	out, err := c.client.CreateClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateClientVpnEndpointWithContext(ctx context.Context, input *ec2.CreateClientVpnEndpointInput, opts ...request.Option) (*ec2.CreateClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnEndpointOutput), nil
	}
	out, err := c.client.CreateClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateClientVpnRoute(input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnRouteOutput), nil
	}
	out, err := c.client.CreateClientVpnRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateClientVpnRouteWithContext(ctx context.Context, input *ec2.CreateClientVpnRouteInput, opts ...request.Option) (*ec2.CreateClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateClientVpnRouteOutput), nil
	}
	out, err := c.client.CreateClientVpnRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCoipCidr(input *ec2.CreateCoipCidrInput) (*ec2.CreateCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipCidrOutput), nil
	}
	out, err := c.client.CreateCoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCoipCidrWithContext(ctx context.Context, input *ec2.CreateCoipCidrInput, opts ...request.Option) (*ec2.CreateCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipCidrOutput), nil
	}
	out, err := c.client.CreateCoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCoipPool(input *ec2.CreateCoipPoolInput) (*ec2.CreateCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipPoolOutput), nil
	}
	out, err := c.client.CreateCoipPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCoipPoolWithContext(ctx context.Context, input *ec2.CreateCoipPoolInput, opts ...request.Option) (*ec2.CreateCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCoipPoolOutput), nil
	}
	out, err := c.client.CreateCoipPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCustomerGateway(input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCustomerGatewayOutput), nil
	}
	out, err := c.client.CreateCustomerGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateCustomerGatewayWithContext(ctx context.Context, input *ec2.CreateCustomerGatewayInput, opts ...request.Option) (*ec2.CreateCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateCustomerGatewayOutput), nil
	}
	out, err := c.client.CreateCustomerGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDefaultSubnet(input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultSubnetOutput), nil
	}
	out, err := c.client.CreateDefaultSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDefaultSubnetWithContext(ctx context.Context, input *ec2.CreateDefaultSubnetInput, opts ...request.Option) (*ec2.CreateDefaultSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultSubnetOutput), nil
	}
	out, err := c.client.CreateDefaultSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDefaultVpc(input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultVpcOutput), nil
	}
	out, err := c.client.CreateDefaultVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDefaultVpcWithContext(ctx context.Context, input *ec2.CreateDefaultVpcInput, opts ...request.Option) (*ec2.CreateDefaultVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDefaultVpcOutput), nil
	}
	out, err := c.client.CreateDefaultVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDhcpOptions(input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDhcpOptionsOutput), nil
	}
	out, err := c.client.CreateDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateDhcpOptionsWithContext(ctx context.Context, input *ec2.CreateDhcpOptionsInput, opts ...request.Option) (*ec2.CreateDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateDhcpOptionsOutput), nil
	}
	out, err := c.client.CreateDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateEgressOnlyInternetGateway(input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.client.CreateEgressOnlyInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateEgressOnlyInternetGatewayWithContext(ctx context.Context, input *ec2.CreateEgressOnlyInternetGatewayInput, opts ...request.Option) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.client.CreateEgressOnlyInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFleet(input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFleetOutput), nil
	}
	out, err := c.client.CreateFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFleetWithContext(ctx context.Context, input *ec2.CreateFleetInput, opts ...request.Option) (*ec2.CreateFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFleetOutput), nil
	}
	out, err := c.client.CreateFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFlowLogs(input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFlowLogsOutput), nil
	}
	out, err := c.client.CreateFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFlowLogsWithContext(ctx context.Context, input *ec2.CreateFlowLogsInput, opts ...request.Option) (*ec2.CreateFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFlowLogsOutput), nil
	}
	out, err := c.client.CreateFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFpgaImage(input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFpgaImageOutput), nil
	}
	out, err := c.client.CreateFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateFpgaImageWithContext(ctx context.Context, input *ec2.CreateFpgaImageInput, opts ...request.Option) (*ec2.CreateFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateFpgaImageOutput), nil
	}
	out, err := c.client.CreateFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateImage(input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateImageOutput), nil
	}
	out, err := c.client.CreateImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateImageWithContext(ctx context.Context, input *ec2.CreateImageInput, opts ...request.Option) (*ec2.CreateImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateImageOutput), nil
	}
	out, err := c.client.CreateImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInstanceEventWindow(input *ec2.CreateInstanceEventWindowInput) (*ec2.CreateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceEventWindowOutput), nil
	}
	out, err := c.client.CreateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInstanceEventWindowWithContext(ctx context.Context, input *ec2.CreateInstanceEventWindowInput, opts ...request.Option) (*ec2.CreateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceEventWindowOutput), nil
	}
	out, err := c.client.CreateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInstanceExportTask(input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceExportTaskOutput), nil
	}
	out, err := c.client.CreateInstanceExportTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInstanceExportTaskWithContext(ctx context.Context, input *ec2.CreateInstanceExportTaskInput, opts ...request.Option) (*ec2.CreateInstanceExportTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInstanceExportTaskOutput), nil
	}
	out, err := c.client.CreateInstanceExportTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInternetGateway(input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInternetGatewayOutput), nil
	}
	out, err := c.client.CreateInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateInternetGatewayWithContext(ctx context.Context, input *ec2.CreateInternetGatewayInput, opts ...request.Option) (*ec2.CreateInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateInternetGatewayOutput), nil
	}
	out, err := c.client.CreateInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpam(input *ec2.CreateIpamInput) (*ec2.CreateIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamOutput), nil
	}
	out, err := c.client.CreateIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpamPool(input *ec2.CreateIpamPoolInput) (*ec2.CreateIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamPoolOutput), nil
	}
	out, err := c.client.CreateIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpamPoolWithContext(ctx context.Context, input *ec2.CreateIpamPoolInput, opts ...request.Option) (*ec2.CreateIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamPoolOutput), nil
	}
	out, err := c.client.CreateIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpamScope(input *ec2.CreateIpamScopeInput) (*ec2.CreateIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamScopeOutput), nil
	}
	out, err := c.client.CreateIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpamScopeWithContext(ctx context.Context, input *ec2.CreateIpamScopeInput, opts ...request.Option) (*ec2.CreateIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamScopeOutput), nil
	}
	out, err := c.client.CreateIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateIpamWithContext(ctx context.Context, input *ec2.CreateIpamInput, opts ...request.Option) (*ec2.CreateIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateIpamOutput), nil
	}
	out, err := c.client.CreateIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateKeyPair(input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateKeyPairOutput), nil
	}
	out, err := c.client.CreateKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateKeyPairWithContext(ctx context.Context, input *ec2.CreateKeyPairInput, opts ...request.Option) (*ec2.CreateKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateKeyPairOutput), nil
	}
	out, err := c.client.CreateKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLaunchTemplate(input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateOutput), nil
	}
	out, err := c.client.CreateLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLaunchTemplateVersion(input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateVersionOutput), nil
	}
	out, err := c.client.CreateLaunchTemplateVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLaunchTemplateVersionWithContext(ctx context.Context, input *ec2.CreateLaunchTemplateVersionInput, opts ...request.Option) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateVersionOutput), nil
	}
	out, err := c.client.CreateLaunchTemplateVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLaunchTemplateWithContext(ctx context.Context, input *ec2.CreateLaunchTemplateInput, opts ...request.Option) (*ec2.CreateLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLaunchTemplateOutput), nil
	}
	out, err := c.client.CreateLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRoute(input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTable(input *ec2.CreateLocalGatewayRouteTableInput) (*ec2.CreateLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input *ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput) (*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTableVpcAssociation(input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTableVpcAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTableVpcAssociationWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTableVpcAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteTableWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteTableOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateLocalGatewayRouteWithContext(ctx context.Context, input *ec2.CreateLocalGatewayRouteInput, opts ...request.Option) (*ec2.CreateLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateLocalGatewayRouteOutput), nil
	}
	out, err := c.client.CreateLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateManagedPrefixList(input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateManagedPrefixListOutput), nil
	}
	out, err := c.client.CreateManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateManagedPrefixListWithContext(ctx context.Context, input *ec2.CreateManagedPrefixListInput, opts ...request.Option) (*ec2.CreateManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateManagedPrefixListOutput), nil
	}
	out, err := c.client.CreateManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNatGateway(input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNatGatewayOutput), nil
	}
	out, err := c.client.CreateNatGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNatGatewayWithContext(ctx context.Context, input *ec2.CreateNatGatewayInput, opts ...request.Option) (*ec2.CreateNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNatGatewayOutput), nil
	}
	out, err := c.client.CreateNatGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkAcl(input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclOutput), nil
	}
	out, err := c.client.CreateNetworkAcl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkAclEntry(input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclEntryOutput), nil
	}
	out, err := c.client.CreateNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkAclEntryWithContext(ctx context.Context, input *ec2.CreateNetworkAclEntryInput, opts ...request.Option) (*ec2.CreateNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclEntryOutput), nil
	}
	out, err := c.client.CreateNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkAclWithContext(ctx context.Context, input *ec2.CreateNetworkAclInput, opts ...request.Option) (*ec2.CreateNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkAclOutput), nil
	}
	out, err := c.client.CreateNetworkAclWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInsightsAccessScope(input *ec2.CreateNetworkInsightsAccessScopeInput) (*ec2.CreateNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.client.CreateNetworkInsightsAccessScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInsightsAccessScopeWithContext(ctx context.Context, input *ec2.CreateNetworkInsightsAccessScopeInput, opts ...request.Option) (*ec2.CreateNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.client.CreateNetworkInsightsAccessScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInsightsPath(input *ec2.CreateNetworkInsightsPathInput) (*ec2.CreateNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsPathOutput), nil
	}
	out, err := c.client.CreateNetworkInsightsPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInsightsPathWithContext(ctx context.Context, input *ec2.CreateNetworkInsightsPathInput, opts ...request.Option) (*ec2.CreateNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInsightsPathOutput), nil
	}
	out, err := c.client.CreateNetworkInsightsPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInterface(input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfaceOutput), nil
	}
	out, err := c.client.CreateNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInterfacePermission(input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfacePermissionOutput), nil
	}
	out, err := c.client.CreateNetworkInterfacePermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInterfacePermissionWithContext(ctx context.Context, input *ec2.CreateNetworkInterfacePermissionInput, opts ...request.Option) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfacePermissionOutput), nil
	}
	out, err := c.client.CreateNetworkInterfacePermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateNetworkInterfaceWithContext(ctx context.Context, input *ec2.CreateNetworkInterfaceInput, opts ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateNetworkInterfaceOutput), nil
	}
	out, err := c.client.CreateNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreatePlacementGroup(input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePlacementGroupOutput), nil
	}
	out, err := c.client.CreatePlacementGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreatePlacementGroupWithContext(ctx context.Context, input *ec2.CreatePlacementGroupInput, opts ...request.Option) (*ec2.CreatePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePlacementGroupOutput), nil
	}
	out, err := c.client.CreatePlacementGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreatePublicIpv4Pool(input *ec2.CreatePublicIpv4PoolInput) (*ec2.CreatePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePublicIpv4PoolOutput), nil
	}
	out, err := c.client.CreatePublicIpv4Pool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreatePublicIpv4PoolWithContext(ctx context.Context, input *ec2.CreatePublicIpv4PoolInput, opts ...request.Option) (*ec2.CreatePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreatePublicIpv4PoolOutput), nil
	}
	out, err := c.client.CreatePublicIpv4PoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateReplaceRootVolumeTask(input *ec2.CreateReplaceRootVolumeTaskInput) (*ec2.CreateReplaceRootVolumeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReplaceRootVolumeTaskOutput), nil
	}
	out, err := c.client.CreateReplaceRootVolumeTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateReplaceRootVolumeTaskWithContext(ctx context.Context, input *ec2.CreateReplaceRootVolumeTaskInput, opts ...request.Option) (*ec2.CreateReplaceRootVolumeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReplaceRootVolumeTaskOutput), nil
	}
	out, err := c.client.CreateReplaceRootVolumeTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateReservedInstancesListing(input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReservedInstancesListingOutput), nil
	}
	out, err := c.client.CreateReservedInstancesListing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateReservedInstancesListingWithContext(ctx context.Context, input *ec2.CreateReservedInstancesListingInput, opts ...request.Option) (*ec2.CreateReservedInstancesListingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateReservedInstancesListingOutput), nil
	}
	out, err := c.client.CreateReservedInstancesListingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRestoreImageTask(input *ec2.CreateRestoreImageTaskInput) (*ec2.CreateRestoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRestoreImageTaskOutput), nil
	}
	out, err := c.client.CreateRestoreImageTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRestoreImageTaskWithContext(ctx context.Context, input *ec2.CreateRestoreImageTaskInput, opts ...request.Option) (*ec2.CreateRestoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRestoreImageTaskOutput), nil
	}
	out, err := c.client.CreateRestoreImageTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRoute(input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteOutput), nil
	}
	out, err := c.client.CreateRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRouteTable(input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteTableOutput), nil
	}
	out, err := c.client.CreateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRouteTableWithContext(ctx context.Context, input *ec2.CreateRouteTableInput, opts ...request.Option) (*ec2.CreateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteTableOutput), nil
	}
	out, err := c.client.CreateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateRouteWithContext(ctx context.Context, input *ec2.CreateRouteInput, opts ...request.Option) (*ec2.CreateRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateRouteOutput), nil
	}
	out, err := c.client.CreateRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSecurityGroup(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSecurityGroupOutput), nil
	}
	out, err := c.client.CreateSecurityGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, input *ec2.CreateSecurityGroupInput, opts ...request.Option) (*ec2.CreateSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSecurityGroupOutput), nil
	}
	out, err := c.client.CreateSecurityGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSnapshots(input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSnapshotsOutput), nil
	}
	out, err := c.client.CreateSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSnapshotsWithContext(ctx context.Context, input *ec2.CreateSnapshotsInput, opts ...request.Option) (*ec2.CreateSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSnapshotsOutput), nil
	}
	out, err := c.client.CreateSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSpotDatafeedSubscription(input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.CreateSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.CreateSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.CreateSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateStoreImageTask(input *ec2.CreateStoreImageTaskInput) (*ec2.CreateStoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateStoreImageTaskOutput), nil
	}
	out, err := c.client.CreateStoreImageTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateStoreImageTaskWithContext(ctx context.Context, input *ec2.CreateStoreImageTaskInput, opts ...request.Option) (*ec2.CreateStoreImageTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateStoreImageTaskOutput), nil
	}
	out, err := c.client.CreateStoreImageTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSubnet(input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetOutput), nil
	}
	out, err := c.client.CreateSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSubnetCidrReservation(input *ec2.CreateSubnetCidrReservationInput) (*ec2.CreateSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetCidrReservationOutput), nil
	}
	out, err := c.client.CreateSubnetCidrReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSubnetCidrReservationWithContext(ctx context.Context, input *ec2.CreateSubnetCidrReservationInput, opts ...request.Option) (*ec2.CreateSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetCidrReservationOutput), nil
	}
	out, err := c.client.CreateSubnetCidrReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateSubnetWithContext(ctx context.Context, input *ec2.CreateSubnetInput, opts ...request.Option) (*ec2.CreateSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateSubnetOutput), nil
	}
	out, err := c.client.CreateSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTags(input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTagsOutput), nil
	}
	out, err := c.client.CreateTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTagsWithContext(ctx context.Context, input *ec2.CreateTagsInput, opts ...request.Option) (*ec2.CreateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTagsOutput), nil
	}
	out, err := c.client.CreateTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorFilter(input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorFilterRule(input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorFilterWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorFilterInput, opts ...request.Option) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorFilterOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorSession(input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorSessionInput, opts ...request.Option) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorTarget(input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorTargetOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTrafficMirrorTargetWithContext(ctx context.Context, input *ec2.CreateTrafficMirrorTargetInput, opts ...request.Option) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTrafficMirrorTargetOutput), nil
	}
	out, err := c.client.CreateTrafficMirrorTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGateway(input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayOutput), nil
	}
	out, err := c.client.CreateTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayConnect(input *ec2.CreateTransitGatewayConnectInput) (*ec2.CreateTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectOutput), nil
	}
	out, err := c.client.CreateTransitGatewayConnect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayConnectPeer(input *ec2.CreateTransitGatewayConnectPeerInput) (*ec2.CreateTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.client.CreateTransitGatewayConnectPeer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayConnectPeerWithContext(ctx context.Context, input *ec2.CreateTransitGatewayConnectPeerInput, opts ...request.Option) (*ec2.CreateTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.client.CreateTransitGatewayConnectPeerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayConnectWithContext(ctx context.Context, input *ec2.CreateTransitGatewayConnectInput, opts ...request.Option) (*ec2.CreateTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayConnectOutput), nil
	}
	out, err := c.client.CreateTransitGatewayConnectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayMulticastDomain(input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.CreateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.CreateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.CreateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPeeringAttachment(input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPolicyTable(input *ec2.CreateTransitGatewayPolicyTableInput) (*ec2.CreateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.CreateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPrefixListReference(input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.CreateTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRoute(input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRouteTable(input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRouteTableAnnouncement(input *ec2.CreateTransitGatewayRouteTableAnnouncementInput) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRouteTableAnnouncement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRouteTableAnnouncementWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteTableAnnouncementInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRouteTableAnnouncementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayRouteWithContext(ctx context.Context, input *ec2.CreateTransitGatewayRouteInput, opts ...request.Option) (*ec2.CreateTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayRouteOutput), nil
	}
	out, err := c.client.CreateTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayVpcAttachment(input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.CreateTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.CreateTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateTransitGatewayWithContext(ctx context.Context, input *ec2.CreateTransitGatewayInput, opts ...request.Option) (*ec2.CreateTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateTransitGatewayOutput), nil
	}
	out, err := c.client.CreateTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessEndpoint(input *ec2.CreateVerifiedAccessEndpointInput) (*ec2.CreateVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.CreateVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessGroup(input *ec2.CreateVerifiedAccessGroupInput) (*ec2.CreateVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessGroupInput, opts ...request.Option) (*ec2.CreateVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessInstance(input *ec2.CreateVerifiedAccessInstanceInput) (*ec2.CreateVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.CreateVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessTrustProvider(input *ec2.CreateVerifiedAccessTrustProviderInput) (*ec2.CreateVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.CreateVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.CreateVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.CreateVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpc(input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcOutput), nil
	}
	out, err := c.client.CreateVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpoint(input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointOutput), nil
	}
	out, err := c.client.CreateVpcEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpointConnectionNotification(input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.client.CreateVpcEndpointConnectionNotification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpointConnectionNotificationWithContext(ctx context.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput, opts ...request.Option) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.client.CreateVpcEndpointConnectionNotificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpointServiceConfiguration(input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.client.CreateVpcEndpointServiceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpointServiceConfigurationWithContext(ctx context.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput, opts ...request.Option) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.client.CreateVpcEndpointServiceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcEndpointWithContext(ctx context.Context, input *ec2.CreateVpcEndpointInput, opts ...request.Option) (*ec2.CreateVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcEndpointOutput), nil
	}
	out, err := c.client.CreateVpcEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcPeeringConnection(input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.CreateVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.CreateVpcPeeringConnectionInput, opts ...request.Option) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.CreateVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpcWithContext(ctx context.Context, input *ec2.CreateVpcInput, opts ...request.Option) (*ec2.CreateVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpcOutput), nil
	}
	out, err := c.client.CreateVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnConnection(input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionOutput), nil
	}
	out, err := c.client.CreateVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnConnectionRoute(input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionRouteOutput), nil
	}
	out, err := c.client.CreateVpnConnectionRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnConnectionRouteWithContext(ctx context.Context, input *ec2.CreateVpnConnectionRouteInput, opts ...request.Option) (*ec2.CreateVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionRouteOutput), nil
	}
	out, err := c.client.CreateVpnConnectionRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnConnectionWithContext(ctx context.Context, input *ec2.CreateVpnConnectionInput, opts ...request.Option) (*ec2.CreateVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnConnectionOutput), nil
	}
	out, err := c.client.CreateVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnGateway(input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnGatewayOutput), nil
	}
	out, err := c.client.CreateVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) CreateVpnGatewayWithContext(ctx context.Context, input *ec2.CreateVpnGatewayInput, opts ...request.Option) (*ec2.CreateVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.CreateVpnGatewayOutput), nil
	}
	out, err := c.client.CreateVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCarrierGateway(input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCarrierGatewayOutput), nil
	}
	out, err := c.client.DeleteCarrierGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCarrierGatewayWithContext(ctx context.Context, input *ec2.DeleteCarrierGatewayInput, opts ...request.Option) (*ec2.DeleteCarrierGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCarrierGatewayOutput), nil
	}
	out, err := c.client.DeleteCarrierGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteClientVpnEndpoint(input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnEndpointOutput), nil
	}
	out, err := c.client.DeleteClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteClientVpnEndpointWithContext(ctx context.Context, input *ec2.DeleteClientVpnEndpointInput, opts ...request.Option) (*ec2.DeleteClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnEndpointOutput), nil
	}
	out, err := c.client.DeleteClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteClientVpnRoute(input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnRouteOutput), nil
	}
	out, err := c.client.DeleteClientVpnRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteClientVpnRouteWithContext(ctx context.Context, input *ec2.DeleteClientVpnRouteInput, opts ...request.Option) (*ec2.DeleteClientVpnRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteClientVpnRouteOutput), nil
	}
	out, err := c.client.DeleteClientVpnRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCoipCidr(input *ec2.DeleteCoipCidrInput) (*ec2.DeleteCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipCidrOutput), nil
	}
	out, err := c.client.DeleteCoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCoipCidrWithContext(ctx context.Context, input *ec2.DeleteCoipCidrInput, opts ...request.Option) (*ec2.DeleteCoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipCidrOutput), nil
	}
	out, err := c.client.DeleteCoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCoipPool(input *ec2.DeleteCoipPoolInput) (*ec2.DeleteCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipPoolOutput), nil
	}
	out, err := c.client.DeleteCoipPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCoipPoolWithContext(ctx context.Context, input *ec2.DeleteCoipPoolInput, opts ...request.Option) (*ec2.DeleteCoipPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCoipPoolOutput), nil
	}
	out, err := c.client.DeleteCoipPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCustomerGateway(input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCustomerGatewayOutput), nil
	}
	out, err := c.client.DeleteCustomerGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteCustomerGatewayWithContext(ctx context.Context, input *ec2.DeleteCustomerGatewayInput, opts ...request.Option) (*ec2.DeleteCustomerGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteCustomerGatewayOutput), nil
	}
	out, err := c.client.DeleteCustomerGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteDhcpOptions(input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteDhcpOptionsOutput), nil
	}
	out, err := c.client.DeleteDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteDhcpOptionsWithContext(ctx context.Context, input *ec2.DeleteDhcpOptionsInput, opts ...request.Option) (*ec2.DeleteDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteDhcpOptionsOutput), nil
	}
	out, err := c.client.DeleteDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteEgressOnlyInternetGateway(input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.client.DeleteEgressOnlyInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteEgressOnlyInternetGatewayWithContext(ctx context.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput, opts ...request.Option) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteEgressOnlyInternetGatewayOutput), nil
	}
	out, err := c.client.DeleteEgressOnlyInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFleets(input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFleetsOutput), nil
	}
	out, err := c.client.DeleteFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFleetsWithContext(ctx context.Context, input *ec2.DeleteFleetsInput, opts ...request.Option) (*ec2.DeleteFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFleetsOutput), nil
	}
	out, err := c.client.DeleteFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFlowLogs(input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFlowLogsOutput), nil
	}
	out, err := c.client.DeleteFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFlowLogsWithContext(ctx context.Context, input *ec2.DeleteFlowLogsInput, opts ...request.Option) (*ec2.DeleteFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFlowLogsOutput), nil
	}
	out, err := c.client.DeleteFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFpgaImage(input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFpgaImageOutput), nil
	}
	out, err := c.client.DeleteFpgaImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteFpgaImageWithContext(ctx context.Context, input *ec2.DeleteFpgaImageInput, opts ...request.Option) (*ec2.DeleteFpgaImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteFpgaImageOutput), nil
	}
	out, err := c.client.DeleteFpgaImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteInstanceEventWindow(input *ec2.DeleteInstanceEventWindowInput) (*ec2.DeleteInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInstanceEventWindowOutput), nil
	}
	out, err := c.client.DeleteInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteInstanceEventWindowWithContext(ctx context.Context, input *ec2.DeleteInstanceEventWindowInput, opts ...request.Option) (*ec2.DeleteInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInstanceEventWindowOutput), nil
	}
	out, err := c.client.DeleteInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteInternetGateway(input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInternetGatewayOutput), nil
	}
	out, err := c.client.DeleteInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteInternetGatewayWithContext(ctx context.Context, input *ec2.DeleteInternetGatewayInput, opts ...request.Option) (*ec2.DeleteInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteInternetGatewayOutput), nil
	}
	out, err := c.client.DeleteInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpam(input *ec2.DeleteIpamInput) (*ec2.DeleteIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamOutput), nil
	}
	out, err := c.client.DeleteIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpamPool(input *ec2.DeleteIpamPoolInput) (*ec2.DeleteIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamPoolOutput), nil
	}
	out, err := c.client.DeleteIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpamPoolWithContext(ctx context.Context, input *ec2.DeleteIpamPoolInput, opts ...request.Option) (*ec2.DeleteIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamPoolOutput), nil
	}
	out, err := c.client.DeleteIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpamScope(input *ec2.DeleteIpamScopeInput) (*ec2.DeleteIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamScopeOutput), nil
	}
	out, err := c.client.DeleteIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpamScopeWithContext(ctx context.Context, input *ec2.DeleteIpamScopeInput, opts ...request.Option) (*ec2.DeleteIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamScopeOutput), nil
	}
	out, err := c.client.DeleteIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteIpamWithContext(ctx context.Context, input *ec2.DeleteIpamInput, opts ...request.Option) (*ec2.DeleteIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteIpamOutput), nil
	}
	out, err := c.client.DeleteIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteKeyPairOutput), nil
	}
	out, err := c.client.DeleteKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteKeyPairWithContext(ctx context.Context, input *ec2.DeleteKeyPairInput, opts ...request.Option) (*ec2.DeleteKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteKeyPairOutput), nil
	}
	out, err := c.client.DeleteKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLaunchTemplate(input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateOutput), nil
	}
	out, err := c.client.DeleteLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLaunchTemplateVersions(input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateVersionsOutput), nil
	}
	out, err := c.client.DeleteLaunchTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLaunchTemplateVersionsWithContext(ctx context.Context, input *ec2.DeleteLaunchTemplateVersionsInput, opts ...request.Option) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateVersionsOutput), nil
	}
	out, err := c.client.DeleteLaunchTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLaunchTemplateWithContext(ctx context.Context, input *ec2.DeleteLaunchTemplateInput, opts ...request.Option) (*ec2.DeleteLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLaunchTemplateOutput), nil
	}
	out, err := c.client.DeleteLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRoute(input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTable(input *ec2.DeleteLocalGatewayRouteTableInput) (*ec2.DeleteLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input *ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTableVpcAssociation(input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTableVpcAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTableVpcAssociationWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTableVpcAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteTableWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteTableOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteLocalGatewayRouteWithContext(ctx context.Context, input *ec2.DeleteLocalGatewayRouteInput, opts ...request.Option) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteLocalGatewayRouteOutput), nil
	}
	out, err := c.client.DeleteLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteManagedPrefixList(input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteManagedPrefixListOutput), nil
	}
	out, err := c.client.DeleteManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteManagedPrefixListWithContext(ctx context.Context, input *ec2.DeleteManagedPrefixListInput, opts ...request.Option) (*ec2.DeleteManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteManagedPrefixListOutput), nil
	}
	out, err := c.client.DeleteManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNatGateway(input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNatGatewayOutput), nil
	}
	out, err := c.client.DeleteNatGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNatGatewayWithContext(ctx context.Context, input *ec2.DeleteNatGatewayInput, opts ...request.Option) (*ec2.DeleteNatGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNatGatewayOutput), nil
	}
	out, err := c.client.DeleteNatGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkAcl(input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclOutput), nil
	}
	out, err := c.client.DeleteNetworkAcl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkAclEntry(input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclEntryOutput), nil
	}
	out, err := c.client.DeleteNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkAclEntryWithContext(ctx context.Context, input *ec2.DeleteNetworkAclEntryInput, opts ...request.Option) (*ec2.DeleteNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclEntryOutput), nil
	}
	out, err := c.client.DeleteNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkAclWithContext(ctx context.Context, input *ec2.DeleteNetworkAclInput, opts ...request.Option) (*ec2.DeleteNetworkAclOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkAclOutput), nil
	}
	out, err := c.client.DeleteNetworkAclWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAccessScope(input *ec2.DeleteNetworkInsightsAccessScopeInput) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAccessScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAccessScopeAnalysis(input *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAccessScopeAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAccessScopeAnalysisWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAccessScopeAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAccessScopeWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAccessScopeInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAccessScopeOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAccessScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAnalysis(input *ec2.DeleteNetworkInsightsAnalysisInput) (*ec2.DeleteNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsAnalysisWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsAnalysisInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsPath(input *ec2.DeleteNetworkInsightsPathInput) (*ec2.DeleteNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsPathOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInsightsPathWithContext(ctx context.Context, input *ec2.DeleteNetworkInsightsPathInput, opts ...request.Option) (*ec2.DeleteNetworkInsightsPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInsightsPathOutput), nil
	}
	out, err := c.client.DeleteNetworkInsightsPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInterface(input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfaceOutput), nil
	}
	out, err := c.client.DeleteNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInterfacePermission(input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfacePermissionOutput), nil
	}
	out, err := c.client.DeleteNetworkInterfacePermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInterfacePermissionWithContext(ctx context.Context, input *ec2.DeleteNetworkInterfacePermissionInput, opts ...request.Option) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfacePermissionOutput), nil
	}
	out, err := c.client.DeleteNetworkInterfacePermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteNetworkInterfaceWithContext(ctx context.Context, input *ec2.DeleteNetworkInterfaceInput, opts ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteNetworkInterfaceOutput), nil
	}
	out, err := c.client.DeleteNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeletePlacementGroup(input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePlacementGroupOutput), nil
	}
	out, err := c.client.DeletePlacementGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeletePlacementGroupWithContext(ctx context.Context, input *ec2.DeletePlacementGroupInput, opts ...request.Option) (*ec2.DeletePlacementGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePlacementGroupOutput), nil
	}
	out, err := c.client.DeletePlacementGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeletePublicIpv4Pool(input *ec2.DeletePublicIpv4PoolInput) (*ec2.DeletePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePublicIpv4PoolOutput), nil
	}
	out, err := c.client.DeletePublicIpv4Pool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeletePublicIpv4PoolWithContext(ctx context.Context, input *ec2.DeletePublicIpv4PoolInput, opts ...request.Option) (*ec2.DeletePublicIpv4PoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeletePublicIpv4PoolOutput), nil
	}
	out, err := c.client.DeletePublicIpv4PoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteQueuedReservedInstances(input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteQueuedReservedInstancesOutput), nil
	}
	out, err := c.client.DeleteQueuedReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteQueuedReservedInstancesWithContext(ctx context.Context, input *ec2.DeleteQueuedReservedInstancesInput, opts ...request.Option) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteQueuedReservedInstancesOutput), nil
	}
	out, err := c.client.DeleteQueuedReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteRoute(input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteOutput), nil
	}
	out, err := c.client.DeleteRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteRouteTable(input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteTableOutput), nil
	}
	out, err := c.client.DeleteRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteRouteTableWithContext(ctx context.Context, input *ec2.DeleteRouteTableInput, opts ...request.Option) (*ec2.DeleteRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteTableOutput), nil
	}
	out, err := c.client.DeleteRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteRouteWithContext(ctx context.Context, input *ec2.DeleteRouteInput, opts ...request.Option) (*ec2.DeleteRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteRouteOutput), nil
	}
	out, err := c.client.DeleteRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSecurityGroupOutput), nil
	}
	out, err := c.client.DeleteSecurityGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, input *ec2.DeleteSecurityGroupInput, opts ...request.Option) (*ec2.DeleteSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSecurityGroupOutput), nil
	}
	out, err := c.client.DeleteSecurityGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSnapshot(input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSnapshotOutput), nil
	}
	out, err := c.client.DeleteSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSnapshotWithContext(ctx context.Context, input *ec2.DeleteSnapshotInput, opts ...request.Option) (*ec2.DeleteSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSnapshotOutput), nil
	}
	out, err := c.client.DeleteSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSpotDatafeedSubscription(input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.DeleteSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.DeleteSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSubnet(input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetOutput), nil
	}
	out, err := c.client.DeleteSubnet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSubnetCidrReservation(input *ec2.DeleteSubnetCidrReservationInput) (*ec2.DeleteSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetCidrReservationOutput), nil
	}
	out, err := c.client.DeleteSubnetCidrReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSubnetCidrReservationWithContext(ctx context.Context, input *ec2.DeleteSubnetCidrReservationInput, opts ...request.Option) (*ec2.DeleteSubnetCidrReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetCidrReservationOutput), nil
	}
	out, err := c.client.DeleteSubnetCidrReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteSubnetWithContext(ctx context.Context, input *ec2.DeleteSubnetInput, opts ...request.Option) (*ec2.DeleteSubnetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteSubnetOutput), nil
	}
	out, err := c.client.DeleteSubnetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTagsOutput), nil
	}
	out, err := c.client.DeleteTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTagsWithContext(ctx context.Context, input *ec2.DeleteTagsInput, opts ...request.Option) (*ec2.DeleteTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTagsOutput), nil
	}
	out, err := c.client.DeleteTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorFilter(input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorFilterRule(input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorFilterWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorFilterOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorSession(input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorSessionInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorTarget(input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorTargetOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTrafficMirrorTargetWithContext(ctx context.Context, input *ec2.DeleteTrafficMirrorTargetInput, opts ...request.Option) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTrafficMirrorTargetOutput), nil
	}
	out, err := c.client.DeleteTrafficMirrorTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGateway(input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayOutput), nil
	}
	out, err := c.client.DeleteTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayConnect(input *ec2.DeleteTransitGatewayConnectInput) (*ec2.DeleteTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayConnect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayConnectPeer(input *ec2.DeleteTransitGatewayConnectPeerInput) (*ec2.DeleteTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayConnectPeer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayConnectPeerWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayConnectPeerInput, opts ...request.Option) (*ec2.DeleteTransitGatewayConnectPeerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectPeerOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayConnectPeerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayConnectWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayConnectInput, opts ...request.Option) (*ec2.DeleteTransitGatewayConnectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayConnectOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayConnectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayMulticastDomain(input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPeeringAttachment(input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPolicyTable(input *ec2.DeleteTransitGatewayPolicyTableInput) (*ec2.DeleteTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPrefixListReference(input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRoute(input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRouteTable(input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRouteTableAnnouncement(input *ec2.DeleteTransitGatewayRouteTableAnnouncementInput) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRouteTableAnnouncement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRouteTableAnnouncementWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteTableAnnouncementInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRouteTableAnnouncementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayRouteWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayRouteInput, opts ...request.Option) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayRouteOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayVpcAttachment(input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteTransitGatewayWithContext(ctx context.Context, input *ec2.DeleteTransitGatewayInput, opts ...request.Option) (*ec2.DeleteTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteTransitGatewayOutput), nil
	}
	out, err := c.client.DeleteTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessEndpoint(input *ec2.DeleteVerifiedAccessEndpointInput) (*ec2.DeleteVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessGroup(input *ec2.DeleteVerifiedAccessGroupInput) (*ec2.DeleteVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessGroupInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessInstance(input *ec2.DeleteVerifiedAccessInstanceInput) (*ec2.DeleteVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessTrustProvider(input *ec2.DeleteVerifiedAccessTrustProviderInput) (*ec2.DeleteVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.DeleteVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.DeleteVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.DeleteVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVolumeOutput), nil
	}
	out, err := c.client.DeleteVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVolumeWithContext(ctx context.Context, input *ec2.DeleteVolumeInput, opts ...request.Option) (*ec2.DeleteVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVolumeOutput), nil
	}
	out, err := c.client.DeleteVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpc(input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcOutput), nil
	}
	out, err := c.client.DeleteVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpointConnectionNotifications(input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpointConnectionNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpointConnectionNotificationsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpointConnectionNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpointServiceConfigurations(input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpointServiceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpointServiceConfigurationsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpointServiceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpoints(input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcEndpointsWithContext(ctx context.Context, input *ec2.DeleteVpcEndpointsInput, opts ...request.Option) (*ec2.DeleteVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcEndpointsOutput), nil
	}
	out, err := c.client.DeleteVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcPeeringConnection(input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.DeleteVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.DeleteVpcPeeringConnectionInput, opts ...request.Option) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.DeleteVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpcWithContext(ctx context.Context, input *ec2.DeleteVpcInput, opts ...request.Option) (*ec2.DeleteVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpcOutput), nil
	}
	out, err := c.client.DeleteVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnConnection(input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionOutput), nil
	}
	out, err := c.client.DeleteVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnConnectionRoute(input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionRouteOutput), nil
	}
	out, err := c.client.DeleteVpnConnectionRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnConnectionRouteWithContext(ctx context.Context, input *ec2.DeleteVpnConnectionRouteInput, opts ...request.Option) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionRouteOutput), nil
	}
	out, err := c.client.DeleteVpnConnectionRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnConnectionWithContext(ctx context.Context, input *ec2.DeleteVpnConnectionInput, opts ...request.Option) (*ec2.DeleteVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnConnectionOutput), nil
	}
	out, err := c.client.DeleteVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnGateway(input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnGatewayOutput), nil
	}
	out, err := c.client.DeleteVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeleteVpnGatewayWithContext(ctx context.Context, input *ec2.DeleteVpnGatewayInput, opts ...request.Option) (*ec2.DeleteVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeleteVpnGatewayOutput), nil
	}
	out, err := c.client.DeleteVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionByoipCidr(input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionByoipCidrOutput), nil
	}
	out, err := c.client.DeprovisionByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionByoipCidrWithContext(ctx context.Context, input *ec2.DeprovisionByoipCidrInput, opts ...request.Option) (*ec2.DeprovisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionByoipCidrOutput), nil
	}
	out, err := c.client.DeprovisionByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionIpamPoolCidr(input *ec2.DeprovisionIpamPoolCidrInput) (*ec2.DeprovisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionIpamPoolCidrOutput), nil
	}
	out, err := c.client.DeprovisionIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionIpamPoolCidrWithContext(ctx context.Context, input *ec2.DeprovisionIpamPoolCidrInput, opts ...request.Option) (*ec2.DeprovisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionIpamPoolCidrOutput), nil
	}
	out, err := c.client.DeprovisionIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionPublicIpv4PoolCidr(input *ec2.DeprovisionPublicIpv4PoolCidrInput) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.client.DeprovisionPublicIpv4PoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeprovisionPublicIpv4PoolCidrWithContext(ctx context.Context, input *ec2.DeprovisionPublicIpv4PoolCidrInput, opts ...request.Option) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeprovisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.client.DeprovisionPublicIpv4PoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterImage(input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterImageOutput), nil
	}
	out, err := c.client.DeregisterImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterImageWithContext(ctx context.Context, input *ec2.DeregisterImageInput, opts ...request.Option) (*ec2.DeregisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterImageOutput), nil
	}
	out, err := c.client.DeregisterImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterInstanceEventNotificationAttributes(input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.DeregisterInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.DeregisterInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterTransitGatewayMulticastGroupMembers(input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.client.DeregisterTransitGatewayMulticastGroupMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterTransitGatewayMulticastGroupMembersWithContext(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput, opts ...request.Option) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.client.DeregisterTransitGatewayMulticastGroupMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterTransitGatewayMulticastGroupSources(input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.client.DeregisterTransitGatewayMulticastGroupSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DeregisterTransitGatewayMulticastGroupSourcesWithContext(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput, opts ...request.Option) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.client.DeregisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAccountAttributes(input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.client.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAccountAttributesWithContext(ctx context.Context, input *ec2.DescribeAccountAttributesInput, opts ...request.Option) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.client.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddressTransfers(input *ec2.DescribeAddressTransfersInput) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.client.DescribeAddressTransfers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddressTransfersPages(input *ec2.DescribeAddressTransfersInput, fn func(*ec2.DescribeAddressTransfersOutput, bool) bool) error {
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
	if err := c.client.DescribeAddressTransfersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeAddressTransfersWithContext(ctx context.Context, input *ec2.DescribeAddressTransfersInput, opts ...request.Option) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.client.DescribeAddressTransfersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.client.DescribeAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddressesAttribute(input *ec2.DescribeAddressesAttributeInput) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.client.DescribeAddressesAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddressesAttributePages(input *ec2.DescribeAddressesAttributeInput, fn func(*ec2.DescribeAddressesAttributeOutput, bool) bool) error {
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
	if err := c.client.DescribeAddressesAttributePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeAddressesAttributeWithContext(ctx context.Context, input *ec2.DescribeAddressesAttributeInput, opts ...request.Option) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.client.DescribeAddressesAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAddressesWithContext(ctx context.Context, input *ec2.DescribeAddressesInput, opts ...request.Option) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.client.DescribeAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAggregateIdFormat(input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.client.DescribeAggregateIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAggregateIdFormatWithContext(ctx context.Context, input *ec2.DescribeAggregateIdFormatInput, opts ...request.Option) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.client.DescribeAggregateIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAvailabilityZones(input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.client.DescribeAvailabilityZones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, input *ec2.DescribeAvailabilityZonesInput, opts ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.client.DescribeAvailabilityZonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAwsNetworkPerformanceMetricSubscriptions(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.client.DescribeAwsNetworkPerformanceMetricSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, fn func(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, bool) bool) error {
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
	if err := c.client.DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx context.Context, input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, opts ...request.Option) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.client.DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeBundleTasks(input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.client.DescribeBundleTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeBundleTasksWithContext(ctx context.Context, input *ec2.DescribeBundleTasksInput, opts ...request.Option) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.client.DescribeBundleTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeByoipCidrs(input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.client.DescribeByoipCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeByoipCidrsPages(input *ec2.DescribeByoipCidrsInput, fn func(*ec2.DescribeByoipCidrsOutput, bool) bool) error {
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
	if err := c.client.DescribeByoipCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeByoipCidrsWithContext(ctx context.Context, input *ec2.DescribeByoipCidrsInput, opts ...request.Option) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.client.DescribeByoipCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCapacityReservationFleets(input *ec2.DescribeCapacityReservationFleetsInput) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.client.DescribeCapacityReservationFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCapacityReservationFleetsPages(input *ec2.DescribeCapacityReservationFleetsInput, fn func(*ec2.DescribeCapacityReservationFleetsOutput, bool) bool) error {
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
	if err := c.client.DescribeCapacityReservationFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeCapacityReservationFleetsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationFleetsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.client.DescribeCapacityReservationFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCapacityReservations(input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.client.DescribeCapacityReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCapacityReservationsPages(input *ec2.DescribeCapacityReservationsInput, fn func(*ec2.DescribeCapacityReservationsOutput, bool) bool) error {
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
	if err := c.client.DescribeCapacityReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeCapacityReservationsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.client.DescribeCapacityReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCarrierGateways(input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.client.DescribeCarrierGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCarrierGatewaysPages(input *ec2.DescribeCarrierGatewaysInput, fn func(*ec2.DescribeCarrierGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeCarrierGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeCarrierGatewaysWithContext(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput, opts ...request.Option) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.client.DescribeCarrierGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClassicLinkInstances(input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.client.DescribeClassicLinkInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClassicLinkInstancesPages(input *ec2.DescribeClassicLinkInstancesInput, fn func(*ec2.DescribeClassicLinkInstancesOutput, bool) bool) error {
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
	if err := c.client.DescribeClassicLinkInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClassicLinkInstancesWithContext(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput, opts ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.client.DescribeClassicLinkInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnAuthorizationRules(input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.client.DescribeClientVpnAuthorizationRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnAuthorizationRulesPages(input *ec2.DescribeClientVpnAuthorizationRulesInput, fn func(*ec2.DescribeClientVpnAuthorizationRulesOutput, bool) bool) error {
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
	if err := c.client.DescribeClientVpnAuthorizationRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClientVpnAuthorizationRulesWithContext(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput, opts ...request.Option) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.client.DescribeClientVpnAuthorizationRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnConnections(input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.client.DescribeClientVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnConnectionsPages(input *ec2.DescribeClientVpnConnectionsInput, fn func(*ec2.DescribeClientVpnConnectionsOutput, bool) bool) error {
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
	if err := c.client.DescribeClientVpnConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClientVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.client.DescribeClientVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnEndpoints(input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.client.DescribeClientVpnEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnEndpointsPages(input *ec2.DescribeClientVpnEndpointsInput, fn func(*ec2.DescribeClientVpnEndpointsOutput, bool) bool) error {
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
	if err := c.client.DescribeClientVpnEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClientVpnEndpointsWithContext(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput, opts ...request.Option) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.client.DescribeClientVpnEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnRoutes(input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.client.DescribeClientVpnRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnRoutesPages(input *ec2.DescribeClientVpnRoutesInput, fn func(*ec2.DescribeClientVpnRoutesOutput, bool) bool) error {
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
	if err := c.client.DescribeClientVpnRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClientVpnRoutesWithContext(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput, opts ...request.Option) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.client.DescribeClientVpnRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnTargetNetworks(input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.client.DescribeClientVpnTargetNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeClientVpnTargetNetworksPages(input *ec2.DescribeClientVpnTargetNetworksInput, fn func(*ec2.DescribeClientVpnTargetNetworksOutput, bool) bool) error {
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
	if err := c.client.DescribeClientVpnTargetNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeClientVpnTargetNetworksWithContext(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput, opts ...request.Option) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.client.DescribeClientVpnTargetNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCoipPools(input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.client.DescribeCoipPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCoipPoolsPages(input *ec2.DescribeCoipPoolsInput, fn func(*ec2.DescribeCoipPoolsOutput, bool) bool) error {
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
	if err := c.client.DescribeCoipPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeCoipPoolsWithContext(ctx context.Context, input *ec2.DescribeCoipPoolsInput, opts ...request.Option) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.client.DescribeCoipPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeConversionTasks(input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.client.DescribeConversionTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeConversionTasksWithContext(ctx context.Context, input *ec2.DescribeConversionTasksInput, opts ...request.Option) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.client.DescribeConversionTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCustomerGateways(input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.client.DescribeCustomerGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeCustomerGatewaysWithContext(ctx context.Context, input *ec2.DescribeCustomerGatewaysInput, opts ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.client.DescribeCustomerGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeDhcpOptions(input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.client.DescribeDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeDhcpOptionsPages(input *ec2.DescribeDhcpOptionsInput, fn func(*ec2.DescribeDhcpOptionsOutput, bool) bool) error {
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
	if err := c.client.DescribeDhcpOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeDhcpOptionsWithContext(ctx context.Context, input *ec2.DescribeDhcpOptionsInput, opts ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.client.DescribeDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeEgressOnlyInternetGateways(input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.client.DescribeEgressOnlyInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeEgressOnlyInternetGatewaysPages(input *ec2.DescribeEgressOnlyInternetGatewaysInput, fn func(*ec2.DescribeEgressOnlyInternetGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeEgressOnlyInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeEgressOnlyInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.client.DescribeEgressOnlyInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeElasticGpus(input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.client.DescribeElasticGpus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeElasticGpusWithContext(ctx context.Context, input *ec2.DescribeElasticGpusInput, opts ...request.Option) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.client.DescribeElasticGpusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeExportImageTasks(input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.client.DescribeExportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeExportImageTasksPages(input *ec2.DescribeExportImageTasksInput, fn func(*ec2.DescribeExportImageTasksOutput, bool) bool) error {
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
	if err := c.client.DescribeExportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeExportImageTasksWithContext(ctx context.Context, input *ec2.DescribeExportImageTasksInput, opts ...request.Option) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.client.DescribeExportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeExportTasks(input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.client.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeExportTasksWithContext(ctx context.Context, input *ec2.DescribeExportTasksInput, opts ...request.Option) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.client.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFastLaunchImages(input *ec2.DescribeFastLaunchImagesInput) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.client.DescribeFastLaunchImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFastLaunchImagesPages(input *ec2.DescribeFastLaunchImagesInput, fn func(*ec2.DescribeFastLaunchImagesOutput, bool) bool) error {
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
	if err := c.client.DescribeFastLaunchImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeFastLaunchImagesWithContext(ctx context.Context, input *ec2.DescribeFastLaunchImagesInput, opts ...request.Option) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.client.DescribeFastLaunchImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFastSnapshotRestores(input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.DescribeFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFastSnapshotRestoresPages(input *ec2.DescribeFastSnapshotRestoresInput, fn func(*ec2.DescribeFastSnapshotRestoresOutput, bool) bool) error {
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
	if err := c.client.DescribeFastSnapshotRestoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput, opts ...request.Option) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.DescribeFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleetHistory(input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.client.DescribeFleetHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleetHistoryWithContext(ctx context.Context, input *ec2.DescribeFleetHistoryInput, opts ...request.Option) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.client.DescribeFleetHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleetInstances(input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.client.DescribeFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeFleetInstancesInput, opts ...request.Option) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.client.DescribeFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleets(input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.client.DescribeFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFleetsPages(input *ec2.DescribeFleetsInput, fn func(*ec2.DescribeFleetsOutput, bool) bool) error {
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
	if err := c.client.DescribeFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeFleetsWithContext(ctx context.Context, input *ec2.DescribeFleetsInput, opts ...request.Option) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.client.DescribeFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFlowLogs(input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.client.DescribeFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFlowLogsPages(input *ec2.DescribeFlowLogsInput, fn func(*ec2.DescribeFlowLogsOutput, bool) bool) error {
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
	if err := c.client.DescribeFlowLogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeFlowLogsWithContext(ctx context.Context, input *ec2.DescribeFlowLogsInput, opts ...request.Option) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.client.DescribeFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFpgaImageAttribute(input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.client.DescribeFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFpgaImageAttributeWithContext(ctx context.Context, input *ec2.DescribeFpgaImageAttributeInput, opts ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.client.DescribeFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFpgaImages(input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.client.DescribeFpgaImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeFpgaImagesPages(input *ec2.DescribeFpgaImagesInput, fn func(*ec2.DescribeFpgaImagesOutput, bool) bool) error {
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
	if err := c.client.DescribeFpgaImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeFpgaImagesWithContext(ctx context.Context, input *ec2.DescribeFpgaImagesInput, opts ...request.Option) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.client.DescribeFpgaImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHostReservationOfferings(input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.client.DescribeHostReservationOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHostReservationOfferingsPages(input *ec2.DescribeHostReservationOfferingsInput, fn func(*ec2.DescribeHostReservationOfferingsOutput, bool) bool) error {
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
	if err := c.client.DescribeHostReservationOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeHostReservationOfferingsWithContext(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput, opts ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.client.DescribeHostReservationOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHostReservations(input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.client.DescribeHostReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHostReservationsPages(input *ec2.DescribeHostReservationsInput, fn func(*ec2.DescribeHostReservationsOutput, bool) bool) error {
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
	if err := c.client.DescribeHostReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeHostReservationsWithContext(ctx context.Context, input *ec2.DescribeHostReservationsInput, opts ...request.Option) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.client.DescribeHostReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHosts(input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.client.DescribeHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeHostsPages(input *ec2.DescribeHostsInput, fn func(*ec2.DescribeHostsOutput, bool) bool) error {
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
	if err := c.client.DescribeHostsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeHostsWithContext(ctx context.Context, input *ec2.DescribeHostsInput, opts ...request.Option) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.client.DescribeHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIamInstanceProfileAssociations(input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.client.DescribeIamInstanceProfileAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIamInstanceProfileAssociationsPages(input *ec2.DescribeIamInstanceProfileAssociationsInput, fn func(*ec2.DescribeIamInstanceProfileAssociationsOutput, bool) bool) error {
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
	if err := c.client.DescribeIamInstanceProfileAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeIamInstanceProfileAssociationsWithContext(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput, opts ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.client.DescribeIamInstanceProfileAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIdFormat(input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.client.DescribeIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdFormatInput, opts ...request.Option) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.client.DescribeIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIdentityIdFormat(input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.client.DescribeIdentityIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIdentityIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdentityIdFormatInput, opts ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.client.DescribeIdentityIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImageAttribute(input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.client.DescribeImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImageAttributeWithContext(ctx context.Context, input *ec2.DescribeImageAttributeInput, opts ...request.Option) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.client.DescribeImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.client.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImagesPages(input *ec2.DescribeImagesInput, fn func(*ec2.DescribeImagesOutput, bool) bool) error {
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
	if err := c.client.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeImagesWithContext(ctx context.Context, input *ec2.DescribeImagesInput, opts ...request.Option) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.client.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImportImageTasks(input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.client.DescribeImportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImportImageTasksPages(input *ec2.DescribeImportImageTasksInput, fn func(*ec2.DescribeImportImageTasksOutput, bool) bool) error {
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
	if err := c.client.DescribeImportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeImportImageTasksWithContext(ctx context.Context, input *ec2.DescribeImportImageTasksInput, opts ...request.Option) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.client.DescribeImportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImportSnapshotTasks(input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.client.DescribeImportSnapshotTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeImportSnapshotTasksPages(input *ec2.DescribeImportSnapshotTasksInput, fn func(*ec2.DescribeImportSnapshotTasksOutput, bool) bool) error {
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
	if err := c.client.DescribeImportSnapshotTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeImportSnapshotTasksWithContext(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput, opts ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.client.DescribeImportSnapshotTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceAttribute(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.client.DescribeInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceAttributeWithContext(ctx context.Context, input *ec2.DescribeInstanceAttributeInput, opts ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.client.DescribeInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceCreditSpecifications(input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.client.DescribeInstanceCreditSpecifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceCreditSpecificationsPages(input *ec2.DescribeInstanceCreditSpecificationsInput, fn func(*ec2.DescribeInstanceCreditSpecificationsOutput, bool) bool) error {
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
	if err := c.client.DescribeInstanceCreditSpecificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstanceCreditSpecificationsWithContext(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput, opts ...request.Option) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.client.DescribeInstanceCreditSpecificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceEventNotificationAttributes(input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.DescribeInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.DescribeInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceEventWindows(input *ec2.DescribeInstanceEventWindowsInput) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.client.DescribeInstanceEventWindows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceEventWindowsPages(input *ec2.DescribeInstanceEventWindowsInput, fn func(*ec2.DescribeInstanceEventWindowsOutput, bool) bool) error {
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
	if err := c.client.DescribeInstanceEventWindowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstanceEventWindowsWithContext(ctx context.Context, input *ec2.DescribeInstanceEventWindowsInput, opts ...request.Option) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.client.DescribeInstanceEventWindowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.client.DescribeInstanceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceStatusPages(input *ec2.DescribeInstanceStatusInput, fn func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {
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
	if err := c.client.DescribeInstanceStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstanceStatusWithContext(ctx context.Context, input *ec2.DescribeInstanceStatusInput, opts ...request.Option) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.client.DescribeInstanceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceTypeOfferings(input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.client.DescribeInstanceTypeOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceTypeOfferingsPages(input *ec2.DescribeInstanceTypeOfferingsInput, fn func(*ec2.DescribeInstanceTypeOfferingsOutput, bool) bool) error {
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
	if err := c.client.DescribeInstanceTypeOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstanceTypeOfferingsWithContext(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput, opts ...request.Option) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.client.DescribeInstanceTypeOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceTypes(input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.client.DescribeInstanceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstanceTypesPages(input *ec2.DescribeInstanceTypesInput, fn func(*ec2.DescribeInstanceTypesOutput, bool) bool) error {
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
	if err := c.client.DescribeInstanceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstanceTypesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypesInput, opts ...request.Option) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.client.DescribeInstanceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.client.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInstancesPages(input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool) error {
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
	if err := c.client.DescribeInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInstancesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.client.DescribeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInternetGateways(input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.client.DescribeInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeInternetGatewaysPages(input *ec2.DescribeInternetGatewaysInput, fn func(*ec2.DescribeInternetGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.client.DescribeInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpamPools(input *ec2.DescribeIpamPoolsInput) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.client.DescribeIpamPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpamPoolsPages(input *ec2.DescribeIpamPoolsInput, fn func(*ec2.DescribeIpamPoolsOutput, bool) bool) error {
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
	if err := c.client.DescribeIpamPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeIpamPoolsWithContext(ctx context.Context, input *ec2.DescribeIpamPoolsInput, opts ...request.Option) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.client.DescribeIpamPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpamScopes(input *ec2.DescribeIpamScopesInput) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.client.DescribeIpamScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpamScopesPages(input *ec2.DescribeIpamScopesInput, fn func(*ec2.DescribeIpamScopesOutput, bool) bool) error {
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
	if err := c.client.DescribeIpamScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeIpamScopesWithContext(ctx context.Context, input *ec2.DescribeIpamScopesInput, opts ...request.Option) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.client.DescribeIpamScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpams(input *ec2.DescribeIpamsInput) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.client.DescribeIpams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpamsPages(input *ec2.DescribeIpamsInput, fn func(*ec2.DescribeIpamsOutput, bool) bool) error {
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
	if err := c.client.DescribeIpamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeIpamsWithContext(ctx context.Context, input *ec2.DescribeIpamsInput, opts ...request.Option) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.client.DescribeIpamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpv6Pools(input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.client.DescribeIpv6Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeIpv6PoolsPages(input *ec2.DescribeIpv6PoolsInput, fn func(*ec2.DescribeIpv6PoolsOutput, bool) bool) error {
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
	if err := c.client.DescribeIpv6PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeIpv6PoolsWithContext(ctx context.Context, input *ec2.DescribeIpv6PoolsInput, opts ...request.Option) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.client.DescribeIpv6PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.client.DescribeKeyPairs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeKeyPairsWithContext(ctx context.Context, input *ec2.DescribeKeyPairsInput, opts ...request.Option) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.client.DescribeKeyPairsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLaunchTemplateVersions(input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.client.DescribeLaunchTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLaunchTemplateVersionsPages(input *ec2.DescribeLaunchTemplateVersionsInput, fn func(*ec2.DescribeLaunchTemplateVersionsOutput, bool) bool) error {
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
	if err := c.client.DescribeLaunchTemplateVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLaunchTemplateVersionsWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput, opts ...request.Option) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.client.DescribeLaunchTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLaunchTemplates(input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.client.DescribeLaunchTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLaunchTemplatesPages(input *ec2.DescribeLaunchTemplatesInput, fn func(*ec2.DescribeLaunchTemplatesOutput, bool) bool) error {
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
	if err := c.client.DescribeLaunchTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLaunchTemplatesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput, opts ...request.Option) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.client.DescribeLaunchTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTableVpcAssociations(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTableVpcAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTableVpcAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewayRouteTableVpcAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTables(input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayRouteTablesPages(input *ec2.DescribeLocalGatewayRouteTablesInput, fn func(*ec2.DescribeLocalGatewayRouteTablesOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayVirtualInterfaceGroups(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayVirtualInterfaceGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayVirtualInterfaceGroupsPages(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewayVirtualInterfaceGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayVirtualInterfaces(input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayVirtualInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewayVirtualInterfacesPages(input *ec2.DescribeLocalGatewayVirtualInterfacesInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfacesOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewayVirtualInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewayVirtualInterfacesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.client.DescribeLocalGatewayVirtualInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGateways(input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.client.DescribeLocalGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeLocalGatewaysPages(input *ec2.DescribeLocalGatewaysInput, fn func(*ec2.DescribeLocalGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeLocalGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeLocalGatewaysWithContext(ctx context.Context, input *ec2.DescribeLocalGatewaysInput, opts ...request.Option) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.client.DescribeLocalGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeManagedPrefixLists(input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.client.DescribeManagedPrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeManagedPrefixListsPages(input *ec2.DescribeManagedPrefixListsInput, fn func(*ec2.DescribeManagedPrefixListsOutput, bool) bool) error {
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
	if err := c.client.DescribeManagedPrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeManagedPrefixListsWithContext(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput, opts ...request.Option) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.client.DescribeManagedPrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeMovingAddresses(input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.client.DescribeMovingAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeMovingAddressesPages(input *ec2.DescribeMovingAddressesInput, fn func(*ec2.DescribeMovingAddressesOutput, bool) bool) error {
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
	if err := c.client.DescribeMovingAddressesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeMovingAddressesWithContext(ctx context.Context, input *ec2.DescribeMovingAddressesInput, opts ...request.Option) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.client.DescribeMovingAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNatGateways(input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.client.DescribeNatGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNatGatewaysPages(input *ec2.DescribeNatGatewaysInput, fn func(*ec2.DescribeNatGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeNatGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNatGatewaysWithContext(ctx context.Context, input *ec2.DescribeNatGatewaysInput, opts ...request.Option) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.client.DescribeNatGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkAcls(input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.client.DescribeNetworkAcls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkAclsPages(input *ec2.DescribeNetworkAclsInput, fn func(*ec2.DescribeNetworkAclsOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkAclsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkAclsWithContext(ctx context.Context, input *ec2.DescribeNetworkAclsInput, opts ...request.Option) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.client.DescribeNetworkAclsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAccessScopeAnalyses(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAccessScopeAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAccessScopeAnalysesPages(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInsightsAccessScopeAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAccessScopes(input *ec2.DescribeNetworkInsightsAccessScopesInput) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAccessScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAccessScopesPages(input *ec2.DescribeNetworkInsightsAccessScopesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopesOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInsightsAccessScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInsightsAccessScopesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAccessScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAnalyses(input *ec2.DescribeNetworkInsightsAnalysesInput) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsAnalysesPages(input *ec2.DescribeNetworkInsightsAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAnalysesOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInsightsAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInsightsAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsPaths(input *ec2.DescribeNetworkInsightsPathsInput) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsPaths(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInsightsPathsPages(input *ec2.DescribeNetworkInsightsPathsInput, fn func(*ec2.DescribeNetworkInsightsPathsOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInsightsPathsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInsightsPathsWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsPathsInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.client.DescribeNetworkInsightsPathsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfaceAttribute(input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfacePermissions(input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfacePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfacePermissionsPages(input *ec2.DescribeNetworkInterfacePermissionsInput, fn func(*ec2.DescribeNetworkInterfacePermissionsOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInterfacePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInterfacePermissionsWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfacePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeNetworkInterfacesPages(input *ec2.DescribeNetworkInterfacesInput, fn func(*ec2.DescribeNetworkInterfacesOutput, bool) bool) error {
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
	if err := c.client.DescribeNetworkInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.client.DescribeNetworkInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePlacementGroups(input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.client.DescribePlacementGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePlacementGroupsWithContext(ctx context.Context, input *ec2.DescribePlacementGroupsInput, opts ...request.Option) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.client.DescribePlacementGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePrefixLists(input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.client.DescribePrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePrefixListsPages(input *ec2.DescribePrefixListsInput, fn func(*ec2.DescribePrefixListsOutput, bool) bool) error {
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
	if err := c.client.DescribePrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribePrefixListsWithContext(ctx context.Context, input *ec2.DescribePrefixListsInput, opts ...request.Option) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.client.DescribePrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePrincipalIdFormat(input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.client.DescribePrincipalIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePrincipalIdFormatPages(input *ec2.DescribePrincipalIdFormatInput, fn func(*ec2.DescribePrincipalIdFormatOutput, bool) bool) error {
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
	if err := c.client.DescribePrincipalIdFormatPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribePrincipalIdFormatWithContext(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput, opts ...request.Option) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.client.DescribePrincipalIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePublicIpv4Pools(input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.client.DescribePublicIpv4Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribePublicIpv4PoolsPages(input *ec2.DescribePublicIpv4PoolsInput, fn func(*ec2.DescribePublicIpv4PoolsOutput, bool) bool) error {
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
	if err := c.client.DescribePublicIpv4PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribePublicIpv4PoolsWithContext(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput, opts ...request.Option) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.client.DescribePublicIpv4PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeRegions(input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.client.DescribeRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeRegionsWithContext(ctx context.Context, input *ec2.DescribeRegionsInput, opts ...request.Option) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.client.DescribeRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReplaceRootVolumeTasks(input *ec2.DescribeReplaceRootVolumeTasksInput) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.client.DescribeReplaceRootVolumeTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReplaceRootVolumeTasksPages(input *ec2.DescribeReplaceRootVolumeTasksInput, fn func(*ec2.DescribeReplaceRootVolumeTasksOutput, bool) bool) error {
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
	if err := c.client.DescribeReplaceRootVolumeTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeReplaceRootVolumeTasksWithContext(ctx context.Context, input *ec2.DescribeReplaceRootVolumeTasksInput, opts ...request.Option) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.client.DescribeReplaceRootVolumeTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstances(input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.client.DescribeReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesListings(input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesListings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesListingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesListingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesListingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesModifications(input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesModificationsPages(input *ec2.DescribeReservedInstancesModificationsInput, fn func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {
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
	if err := c.client.DescribeReservedInstancesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeReservedInstancesModificationsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesOfferings(input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesOfferingsPages(input *ec2.DescribeReservedInstancesOfferingsInput, fn func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {
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
	if err := c.client.DescribeReservedInstancesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeReservedInstancesOfferingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeReservedInstancesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.client.DescribeReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeRouteTables(input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.client.DescribeRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeRouteTablesPages(input *ec2.DescribeRouteTablesInput, fn func(*ec2.DescribeRouteTablesOutput, bool) bool) error {
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
	if err := c.client.DescribeRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, input *ec2.DescribeRouteTablesInput, opts ...request.Option) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.client.DescribeRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeScheduledInstanceAvailability(input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.client.DescribeScheduledInstanceAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeScheduledInstanceAvailabilityPages(input *ec2.DescribeScheduledInstanceAvailabilityInput, fn func(*ec2.DescribeScheduledInstanceAvailabilityOutput, bool) bool) error {
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
	if err := c.client.DescribeScheduledInstanceAvailabilityPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeScheduledInstanceAvailabilityWithContext(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput, opts ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.client.DescribeScheduledInstanceAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeScheduledInstances(input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.client.DescribeScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeScheduledInstancesPages(input *ec2.DescribeScheduledInstancesInput, fn func(*ec2.DescribeScheduledInstancesOutput, bool) bool) error {
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
	if err := c.client.DescribeScheduledInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeScheduledInstancesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstancesInput, opts ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.client.DescribeScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroupReferences(input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.client.DescribeSecurityGroupReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroupReferencesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupReferencesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.client.DescribeSecurityGroupReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroupRules(input *ec2.DescribeSecurityGroupRulesInput) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.client.DescribeSecurityGroupRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroupRulesPages(input *ec2.DescribeSecurityGroupRulesInput, fn func(*ec2.DescribeSecurityGroupRulesOutput, bool) bool) error {
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
	if err := c.client.DescribeSecurityGroupRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSecurityGroupRulesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupRulesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.client.DescribeSecurityGroupRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.client.DescribeSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSecurityGroupsPages(input *ec2.DescribeSecurityGroupsInput, fn func(*ec2.DescribeSecurityGroupsOutput, bool) bool) error {
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
	if err := c.client.DescribeSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.client.DescribeSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshotAttribute(input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.client.DescribeSnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshotAttributeWithContext(ctx context.Context, input *ec2.DescribeSnapshotAttributeInput, opts ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.client.DescribeSnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshotTierStatus(input *ec2.DescribeSnapshotTierStatusInput) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.client.DescribeSnapshotTierStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshotTierStatusPages(input *ec2.DescribeSnapshotTierStatusInput, fn func(*ec2.DescribeSnapshotTierStatusOutput, bool) bool) error {
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
	if err := c.client.DescribeSnapshotTierStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSnapshotTierStatusWithContext(ctx context.Context, input *ec2.DescribeSnapshotTierStatusInput, opts ...request.Option) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.client.DescribeSnapshotTierStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshots(input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.client.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSnapshotsPages(input *ec2.DescribeSnapshotsInput, fn func(*ec2.DescribeSnapshotsOutput, bool) bool) error {
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
	if err := c.client.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, input *ec2.DescribeSnapshotsInput, opts ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.client.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotDatafeedSubscription(input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.DescribeSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.client.DescribeSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetInstances(input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.client.DescribeSpotFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeSpotFleetInstancesInput, opts ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.client.DescribeSpotFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetRequestHistory(input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.client.DescribeSpotFleetRequestHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetRequestHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestHistoryInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.client.DescribeSpotFleetRequestHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetRequests(input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.client.DescribeSpotFleetRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotFleetRequestsPages(input *ec2.DescribeSpotFleetRequestsInput, fn func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error {
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
	if err := c.client.DescribeSpotFleetRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSpotFleetRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.client.DescribeSpotFleetRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotInstanceRequests(input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.client.DescribeSpotInstanceRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotInstanceRequestsPages(input *ec2.DescribeSpotInstanceRequestsInput, fn func(*ec2.DescribeSpotInstanceRequestsOutput, bool) bool) error {
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
	if err := c.client.DescribeSpotInstanceRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSpotInstanceRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput, opts ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.client.DescribeSpotInstanceRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotPriceHistory(input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.client.DescribeSpotPriceHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSpotPriceHistoryPages(input *ec2.DescribeSpotPriceHistoryInput, fn func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {
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
	if err := c.client.DescribeSpotPriceHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSpotPriceHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput, opts ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.client.DescribeSpotPriceHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeStaleSecurityGroups(input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.client.DescribeStaleSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeStaleSecurityGroupsPages(input *ec2.DescribeStaleSecurityGroupsInput, fn func(*ec2.DescribeStaleSecurityGroupsOutput, bool) bool) error {
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
	if err := c.client.DescribeStaleSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeStaleSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.client.DescribeStaleSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeStoreImageTasks(input *ec2.DescribeStoreImageTasksInput) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.client.DescribeStoreImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeStoreImageTasksPages(input *ec2.DescribeStoreImageTasksInput, fn func(*ec2.DescribeStoreImageTasksOutput, bool) bool) error {
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
	if err := c.client.DescribeStoreImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeStoreImageTasksWithContext(ctx context.Context, input *ec2.DescribeStoreImageTasksInput, opts ...request.Option) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.client.DescribeStoreImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.client.DescribeSubnets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeSubnetsPages(input *ec2.DescribeSubnetsInput, fn func(*ec2.DescribeSubnetsOutput, bool) bool) error {
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
	if err := c.client.DescribeSubnetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeSubnetsWithContext(ctx context.Context, input *ec2.DescribeSubnetsInput, opts ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.client.DescribeSubnetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.client.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTagsPages(input *ec2.DescribeTagsInput, fn func(*ec2.DescribeTagsOutput, bool) bool) error {
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
	if err := c.client.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTagsWithContext(ctx context.Context, input *ec2.DescribeTagsInput, opts ...request.Option) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.client.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorFilters(input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorFiltersPages(input *ec2.DescribeTrafficMirrorFiltersInput, fn func(*ec2.DescribeTrafficMirrorFiltersOutput, bool) bool) error {
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
	if err := c.client.DescribeTrafficMirrorFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTrafficMirrorFiltersWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorSessions(input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorSessionsPages(input *ec2.DescribeTrafficMirrorSessionsInput, fn func(*ec2.DescribeTrafficMirrorSessionsOutput, bool) bool) error {
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
	if err := c.client.DescribeTrafficMirrorSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTrafficMirrorSessionsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorTargets(input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrafficMirrorTargetsPages(input *ec2.DescribeTrafficMirrorTargetsInput, fn func(*ec2.DescribeTrafficMirrorTargetsOutput, bool) bool) error {
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
	if err := c.client.DescribeTrafficMirrorTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTrafficMirrorTargetsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.client.DescribeTrafficMirrorTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayAttachments(input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayAttachmentsPages(input *ec2.DescribeTransitGatewayAttachmentsInput, fn func(*ec2.DescribeTransitGatewayAttachmentsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayConnectPeers(input *ec2.DescribeTransitGatewayConnectPeersInput) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayConnectPeers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayConnectPeersPages(input *ec2.DescribeTransitGatewayConnectPeersInput, fn func(*ec2.DescribeTransitGatewayConnectPeersOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayConnectPeersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayConnectPeersWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectPeersInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayConnectPeersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayConnects(input *ec2.DescribeTransitGatewayConnectsInput) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayConnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayConnectsPages(input *ec2.DescribeTransitGatewayConnectsInput, fn func(*ec2.DescribeTransitGatewayConnectsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayConnectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayConnectsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayConnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayMulticastDomains(input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayMulticastDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayMulticastDomainsPages(input *ec2.DescribeTransitGatewayMulticastDomainsInput, fn func(*ec2.DescribeTransitGatewayMulticastDomainsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayMulticastDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayMulticastDomainsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayMulticastDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayPeeringAttachments(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayPeeringAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayPeeringAttachmentsPages(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, fn func(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayPeeringAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayPeeringAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayPeeringAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayPolicyTables(input *ec2.DescribeTransitGatewayPolicyTablesInput) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayPolicyTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayPolicyTablesPages(input *ec2.DescribeTransitGatewayPolicyTablesInput, fn func(*ec2.DescribeTransitGatewayPolicyTablesOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayPolicyTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayPolicyTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPolicyTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayPolicyTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayRouteTableAnnouncements(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayRouteTableAnnouncements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayRouteTableAnnouncementsPages(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, fn func(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayRouteTableAnnouncementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayRouteTables(input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayRouteTablesPages(input *ec2.DescribeTransitGatewayRouteTablesInput, fn func(*ec2.DescribeTransitGatewayRouteTablesOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayVpcAttachments(input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayVpcAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewayVpcAttachmentsPages(input *ec2.DescribeTransitGatewayVpcAttachmentsInput, fn func(*ec2.DescribeTransitGatewayVpcAttachmentsOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewayVpcAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewayVpcAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.client.DescribeTransitGatewayVpcAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGateways(input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.client.DescribeTransitGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTransitGatewaysPages(input *ec2.DescribeTransitGatewaysInput, fn func(*ec2.DescribeTransitGatewaysOutput, bool) bool) error {
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
	if err := c.client.DescribeTransitGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTransitGatewaysWithContext(ctx context.Context, input *ec2.DescribeTransitGatewaysInput, opts ...request.Option) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.client.DescribeTransitGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrunkInterfaceAssociations(input *ec2.DescribeTrunkInterfaceAssociationsInput) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.client.DescribeTrunkInterfaceAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeTrunkInterfaceAssociationsPages(input *ec2.DescribeTrunkInterfaceAssociationsInput, fn func(*ec2.DescribeTrunkInterfaceAssociationsOutput, bool) bool) error {
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
	if err := c.client.DescribeTrunkInterfaceAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeTrunkInterfaceAssociationsWithContext(ctx context.Context, input *ec2.DescribeTrunkInterfaceAssociationsInput, opts ...request.Option) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.client.DescribeTrunkInterfaceAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessEndpoints(input *ec2.DescribeVerifiedAccessEndpointsInput) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessEndpointsPages(input *ec2.DescribeVerifiedAccessEndpointsInput, fn func(*ec2.DescribeVerifiedAccessEndpointsOutput, bool) bool) error {
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
	if err := c.client.DescribeVerifiedAccessEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVerifiedAccessEndpointsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessEndpointsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessGroups(input *ec2.DescribeVerifiedAccessGroupsInput) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessGroupsPages(input *ec2.DescribeVerifiedAccessGroupsInput, fn func(*ec2.DescribeVerifiedAccessGroupsOutput, bool) bool) error {
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
	if err := c.client.DescribeVerifiedAccessGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVerifiedAccessGroupsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessGroupsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessInstanceLoggingConfigurations(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessInstanceLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, fn func(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, bool) bool) error {
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
	if err := c.client.DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessInstances(input *ec2.DescribeVerifiedAccessInstancesInput) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessInstancesPages(input *ec2.DescribeVerifiedAccessInstancesInput, fn func(*ec2.DescribeVerifiedAccessInstancesOutput, bool) bool) error {
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
	if err := c.client.DescribeVerifiedAccessInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVerifiedAccessInstancesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstancesInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessTrustProviders(input *ec2.DescribeVerifiedAccessTrustProvidersInput) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessTrustProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVerifiedAccessTrustProvidersPages(input *ec2.DescribeVerifiedAccessTrustProvidersInput, fn func(*ec2.DescribeVerifiedAccessTrustProvidersOutput, bool) bool) error {
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
	if err := c.client.DescribeVerifiedAccessTrustProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVerifiedAccessTrustProvidersWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessTrustProvidersInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.client.DescribeVerifiedAccessTrustProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumeAttribute(input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.client.DescribeVolumeAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumeAttributeWithContext(ctx context.Context, input *ec2.DescribeVolumeAttributeInput, opts ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.client.DescribeVolumeAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumeStatus(input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.client.DescribeVolumeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumeStatusPages(input *ec2.DescribeVolumeStatusInput, fn func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {
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
	if err := c.client.DescribeVolumeStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVolumeStatusWithContext(ctx context.Context, input *ec2.DescribeVolumeStatusInput, opts ...request.Option) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.client.DescribeVolumeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.client.DescribeVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumesModifications(input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.client.DescribeVolumesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumesModificationsPages(input *ec2.DescribeVolumesModificationsInput, fn func(*ec2.DescribeVolumesModificationsOutput, bool) bool) error {
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
	if err := c.client.DescribeVolumesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVolumesModificationsWithContext(ctx context.Context, input *ec2.DescribeVolumesModificationsInput, opts ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.client.DescribeVolumesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVolumesPages(input *ec2.DescribeVolumesInput, fn func(*ec2.DescribeVolumesOutput, bool) bool) error {
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
	if err := c.client.DescribeVolumesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVolumesWithContext(ctx context.Context, input *ec2.DescribeVolumesInput, opts ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.client.DescribeVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcAttribute(input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.client.DescribeVpcAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcAttributeWithContext(ctx context.Context, input *ec2.DescribeVpcAttributeInput, opts ...request.Option) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.client.DescribeVpcAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcClassicLink(input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.client.DescribeVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcClassicLinkDnsSupport(input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.DescribeVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcClassicLinkDnsSupportPages(input *ec2.DescribeVpcClassicLinkDnsSupportInput, fn func(*ec2.DescribeVpcClassicLinkDnsSupportOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcClassicLinkDnsSupportPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.DescribeVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcClassicLinkWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.client.DescribeVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointConnectionNotifications(input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointConnectionNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointConnectionNotificationsPages(input *ec2.DescribeVpcEndpointConnectionNotificationsInput, fn func(*ec2.DescribeVpcEndpointConnectionNotificationsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcEndpointConnectionNotificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcEndpointConnectionNotificationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointConnectionNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointConnections(input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointConnectionsPages(input *ec2.DescribeVpcEndpointConnectionsInput, fn func(*ec2.DescribeVpcEndpointConnectionsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcEndpointConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServiceConfigurations(input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServiceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServiceConfigurationsPages(input *ec2.DescribeVpcEndpointServiceConfigurationsInput, fn func(*ec2.DescribeVpcEndpointServiceConfigurationsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcEndpointServiceConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcEndpointServiceConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServiceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServicePermissions(input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServicePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServicePermissionsPages(input *ec2.DescribeVpcEndpointServicePermissionsInput, fn func(*ec2.DescribeVpcEndpointServicePermissionsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcEndpointServicePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcEndpointServicePermissionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServicePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServices(input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointServicesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicesInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpoints(input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcEndpointsPages(input *ec2.DescribeVpcEndpointsInput, fn func(*ec2.DescribeVpcEndpointsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcEndpointsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.client.DescribeVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcPeeringConnections(input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpcPeeringConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcPeeringConnectionsPages(input *ec2.DescribeVpcPeeringConnectionsInput, fn func(*ec2.DescribeVpcPeeringConnectionsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcPeeringConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcPeeringConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpcPeeringConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.client.DescribeVpcs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpcsPages(input *ec2.DescribeVpcsInput, fn func(*ec2.DescribeVpcsOutput, bool) bool) error {
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
	if err := c.client.DescribeVpcsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) DescribeVpcsWithContext(ctx context.Context, input *ec2.DescribeVpcsInput, opts ...request.Option) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.client.DescribeVpcsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpnConnections(input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.client.DescribeVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpnGateways(input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.client.DescribeVpnGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DescribeVpnGatewaysWithContext(ctx context.Context, input *ec2.DescribeVpnGatewaysInput, opts ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.client.DescribeVpnGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachClassicLinkVpc(input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachClassicLinkVpcOutput), nil
	}
	out, err := c.client.DetachClassicLinkVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachClassicLinkVpcWithContext(ctx context.Context, input *ec2.DetachClassicLinkVpcInput, opts ...request.Option) (*ec2.DetachClassicLinkVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachClassicLinkVpcOutput), nil
	}
	out, err := c.client.DetachClassicLinkVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachInternetGateway(input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachInternetGatewayOutput), nil
	}
	out, err := c.client.DetachInternetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachInternetGatewayWithContext(ctx context.Context, input *ec2.DetachInternetGatewayInput, opts ...request.Option) (*ec2.DetachInternetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachInternetGatewayOutput), nil
	}
	out, err := c.client.DetachInternetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachNetworkInterface(input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachNetworkInterfaceOutput), nil
	}
	out, err := c.client.DetachNetworkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, input *ec2.DetachNetworkInterfaceInput, opts ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachNetworkInterfaceOutput), nil
	}
	out, err := c.client.DetachNetworkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachVerifiedAccessTrustProvider(input *ec2.DetachVerifiedAccessTrustProviderInput) (*ec2.DetachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.DetachVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.DetachVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.DetachVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.DetachVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachVpnGateway(input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVpnGatewayOutput), nil
	}
	out, err := c.client.DetachVpnGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DetachVpnGatewayWithContext(ctx context.Context, input *ec2.DetachVpnGatewayInput, opts ...request.Option) (*ec2.DetachVpnGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DetachVpnGatewayOutput), nil
	}
	out, err := c.client.DetachVpnGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableAddressTransfer(input *ec2.DisableAddressTransferInput) (*ec2.DisableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAddressTransferOutput), nil
	}
	out, err := c.client.DisableAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableAddressTransferWithContext(ctx context.Context, input *ec2.DisableAddressTransferInput, opts ...request.Option) (*ec2.DisableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAddressTransferOutput), nil
	}
	out, err := c.client.DisableAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableAwsNetworkPerformanceMetricSubscription(input *ec2.DisableAwsNetworkPerformanceMetricSubscriptionInput) (*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.client.DisableAwsNetworkPerformanceMetricSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx context.Context, input *ec2.DisableAwsNetworkPerformanceMetricSubscriptionInput, opts ...request.Option) (*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.client.DisableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableEbsEncryptionByDefault(input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.DisableEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.DisableEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.DisableEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableFastLaunch(input *ec2.DisableFastLaunchInput) (*ec2.DisableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastLaunchOutput), nil
	}
	out, err := c.client.DisableFastLaunch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableFastLaunchWithContext(ctx context.Context, input *ec2.DisableFastLaunchInput, opts ...request.Option) (*ec2.DisableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastLaunchOutput), nil
	}
	out, err := c.client.DisableFastLaunchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableFastSnapshotRestores(input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.DisableFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.DisableFastSnapshotRestoresInput, opts ...request.Option) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.DisableFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableImageDeprecation(input *ec2.DisableImageDeprecationInput) (*ec2.DisableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableImageDeprecationOutput), nil
	}
	out, err := c.client.DisableImageDeprecation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableImageDeprecationWithContext(ctx context.Context, input *ec2.DisableImageDeprecationInput, opts ...request.Option) (*ec2.DisableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableImageDeprecationOutput), nil
	}
	out, err := c.client.DisableImageDeprecationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableIpamOrganizationAdminAccount(input *ec2.DisableIpamOrganizationAdminAccountInput) (*ec2.DisableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.client.DisableIpamOrganizationAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableIpamOrganizationAdminAccountWithContext(ctx context.Context, input *ec2.DisableIpamOrganizationAdminAccountInput, opts ...request.Option) (*ec2.DisableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.client.DisableIpamOrganizationAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableSerialConsoleAccess(input *ec2.DisableSerialConsoleAccessInput) (*ec2.DisableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableSerialConsoleAccessOutput), nil
	}
	out, err := c.client.DisableSerialConsoleAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableSerialConsoleAccessWithContext(ctx context.Context, input *ec2.DisableSerialConsoleAccessInput, opts ...request.Option) (*ec2.DisableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableSerialConsoleAccessOutput), nil
	}
	out, err := c.client.DisableSerialConsoleAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableTransitGatewayRouteTablePropagation(input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.client.DisableTransitGatewayRouteTablePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableTransitGatewayRouteTablePropagationWithContext(ctx context.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput, opts ...request.Option) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.client.DisableTransitGatewayRouteTablePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVgwRoutePropagation(input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVgwRoutePropagationOutput), nil
	}
	out, err := c.client.DisableVgwRoutePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVgwRoutePropagationWithContext(ctx context.Context, input *ec2.DisableVgwRoutePropagationInput, opts ...request.Option) (*ec2.DisableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVgwRoutePropagationOutput), nil
	}
	out, err := c.client.DisableVgwRoutePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVpcClassicLink(input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkOutput), nil
	}
	out, err := c.client.DisableVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVpcClassicLinkDnsSupport(input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.DisableVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.DisableVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisableVpcClassicLinkWithContext(ctx context.Context, input *ec2.DisableVpcClassicLinkInput, opts ...request.Option) (*ec2.DisableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisableVpcClassicLinkOutput), nil
	}
	out, err := c.client.DisableVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateAddress(input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateAddressOutput), nil
	}
	out, err := c.client.DisassociateAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateAddressWithContext(ctx context.Context, input *ec2.DisassociateAddressInput, opts ...request.Option) (*ec2.DisassociateAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateAddressOutput), nil
	}
	out, err := c.client.DisassociateAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateClientVpnTargetNetwork(input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.DisassociateClientVpnTargetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateClientVpnTargetNetworkWithContext(ctx context.Context, input *ec2.DisassociateClientVpnTargetNetworkInput, opts ...request.Option) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateClientVpnTargetNetworkOutput), nil
	}
	out, err := c.client.DisassociateClientVpnTargetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateEnclaveCertificateIamRole(input *ec2.DisassociateEnclaveCertificateIamRoleInput) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.client.DisassociateEnclaveCertificateIamRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateEnclaveCertificateIamRoleWithContext(ctx context.Context, input *ec2.DisassociateEnclaveCertificateIamRoleInput, opts ...request.Option) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateEnclaveCertificateIamRoleOutput), nil
	}
	out, err := c.client.DisassociateEnclaveCertificateIamRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateIamInstanceProfile(input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateIamInstanceProfileOutput), nil
	}
	out, err := c.client.DisassociateIamInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateIamInstanceProfileWithContext(ctx context.Context, input *ec2.DisassociateIamInstanceProfileInput, opts ...request.Option) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateIamInstanceProfileOutput), nil
	}
	out, err := c.client.DisassociateIamInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateInstanceEventWindow(input *ec2.DisassociateInstanceEventWindowInput) (*ec2.DisassociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateInstanceEventWindowOutput), nil
	}
	out, err := c.client.DisassociateInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateInstanceEventWindowWithContext(ctx context.Context, input *ec2.DisassociateInstanceEventWindowInput, opts ...request.Option) (*ec2.DisassociateInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateInstanceEventWindowOutput), nil
	}
	out, err := c.client.DisassociateInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateRouteTable(input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateRouteTableOutput), nil
	}
	out, err := c.client.DisassociateRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateRouteTableWithContext(ctx context.Context, input *ec2.DisassociateRouteTableInput, opts ...request.Option) (*ec2.DisassociateRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateRouteTableOutput), nil
	}
	out, err := c.client.DisassociateRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateSubnetCidrBlock(input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateSubnetCidrBlockOutput), nil
	}
	out, err := c.client.DisassociateSubnetCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateSubnetCidrBlockWithContext(ctx context.Context, input *ec2.DisassociateSubnetCidrBlockInput, opts ...request.Option) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateSubnetCidrBlockOutput), nil
	}
	out, err := c.client.DisassociateSubnetCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayMulticastDomain(input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayMulticastDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayMulticastDomainWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayMulticastDomainOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayMulticastDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayPolicyTable(input *ec2.DisassociateTransitGatewayPolicyTableInput) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayPolicyTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayPolicyTableWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayPolicyTableInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayPolicyTableOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayPolicyTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayRouteTable(input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayRouteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTransitGatewayRouteTableWithContext(ctx context.Context, input *ec2.DisassociateTransitGatewayRouteTableInput, opts ...request.Option) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTransitGatewayRouteTableOutput), nil
	}
	out, err := c.client.DisassociateTransitGatewayRouteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTrunkInterface(input *ec2.DisassociateTrunkInterfaceInput) (*ec2.DisassociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTrunkInterfaceOutput), nil
	}
	out, err := c.client.DisassociateTrunkInterface(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateTrunkInterfaceWithContext(ctx context.Context, input *ec2.DisassociateTrunkInterfaceInput, opts ...request.Option) (*ec2.DisassociateTrunkInterfaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateTrunkInterfaceOutput), nil
	}
	out, err := c.client.DisassociateTrunkInterfaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateVpcCidrBlock(input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateVpcCidrBlockOutput), nil
	}
	out, err := c.client.DisassociateVpcCidrBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) DisassociateVpcCidrBlockWithContext(ctx context.Context, input *ec2.DisassociateVpcCidrBlockInput, opts ...request.Option) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DisassociateVpcCidrBlockOutput), nil
	}
	out, err := c.client.DisassociateVpcCidrBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableAddressTransfer(input *ec2.EnableAddressTransferInput) (*ec2.EnableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAddressTransferOutput), nil
	}
	out, err := c.client.EnableAddressTransfer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableAddressTransferWithContext(ctx context.Context, input *ec2.EnableAddressTransferInput, opts ...request.Option) (*ec2.EnableAddressTransferOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAddressTransferOutput), nil
	}
	out, err := c.client.EnableAddressTransferWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableAwsNetworkPerformanceMetricSubscription(input *ec2.EnableAwsNetworkPerformanceMetricSubscriptionInput) (*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.client.EnableAwsNetworkPerformanceMetricSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx context.Context, input *ec2.EnableAwsNetworkPerformanceMetricSubscriptionInput, opts ...request.Option) (*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableAwsNetworkPerformanceMetricSubscriptionOutput), nil
	}
	out, err := c.client.EnableAwsNetworkPerformanceMetricSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableEbsEncryptionByDefault(input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.EnableEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.EnableEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.EnableEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableFastLaunch(input *ec2.EnableFastLaunchInput) (*ec2.EnableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastLaunchOutput), nil
	}
	out, err := c.client.EnableFastLaunch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableFastLaunchWithContext(ctx context.Context, input *ec2.EnableFastLaunchInput, opts ...request.Option) (*ec2.EnableFastLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastLaunchOutput), nil
	}
	out, err := c.client.EnableFastLaunchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableFastSnapshotRestores(input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.EnableFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.EnableFastSnapshotRestoresInput, opts ...request.Option) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableFastSnapshotRestoresOutput), nil
	}
	out, err := c.client.EnableFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableImageDeprecation(input *ec2.EnableImageDeprecationInput) (*ec2.EnableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableImageDeprecationOutput), nil
	}
	out, err := c.client.EnableImageDeprecation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableImageDeprecationWithContext(ctx context.Context, input *ec2.EnableImageDeprecationInput, opts ...request.Option) (*ec2.EnableImageDeprecationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableImageDeprecationOutput), nil
	}
	out, err := c.client.EnableImageDeprecationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableIpamOrganizationAdminAccount(input *ec2.EnableIpamOrganizationAdminAccountInput) (*ec2.EnableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.client.EnableIpamOrganizationAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableIpamOrganizationAdminAccountWithContext(ctx context.Context, input *ec2.EnableIpamOrganizationAdminAccountInput, opts ...request.Option) (*ec2.EnableIpamOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableIpamOrganizationAdminAccountOutput), nil
	}
	out, err := c.client.EnableIpamOrganizationAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableReachabilityAnalyzerOrganizationSharing(input *ec2.EnableReachabilityAnalyzerOrganizationSharingInput) (*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput), nil
	}
	out, err := c.client.EnableReachabilityAnalyzerOrganizationSharing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableReachabilityAnalyzerOrganizationSharingWithContext(ctx context.Context, input *ec2.EnableReachabilityAnalyzerOrganizationSharingInput, opts ...request.Option) (*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableReachabilityAnalyzerOrganizationSharingOutput), nil
	}
	out, err := c.client.EnableReachabilityAnalyzerOrganizationSharingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableSerialConsoleAccess(input *ec2.EnableSerialConsoleAccessInput) (*ec2.EnableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableSerialConsoleAccessOutput), nil
	}
	out, err := c.client.EnableSerialConsoleAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableSerialConsoleAccessWithContext(ctx context.Context, input *ec2.EnableSerialConsoleAccessInput, opts ...request.Option) (*ec2.EnableSerialConsoleAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableSerialConsoleAccessOutput), nil
	}
	out, err := c.client.EnableSerialConsoleAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableTransitGatewayRouteTablePropagation(input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.client.EnableTransitGatewayRouteTablePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableTransitGatewayRouteTablePropagationWithContext(ctx context.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput, opts ...request.Option) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableTransitGatewayRouteTablePropagationOutput), nil
	}
	out, err := c.client.EnableTransitGatewayRouteTablePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVgwRoutePropagation(input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVgwRoutePropagationOutput), nil
	}
	out, err := c.client.EnableVgwRoutePropagation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVgwRoutePropagationWithContext(ctx context.Context, input *ec2.EnableVgwRoutePropagationInput, opts ...request.Option) (*ec2.EnableVgwRoutePropagationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVgwRoutePropagationOutput), nil
	}
	out, err := c.client.EnableVgwRoutePropagationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVolumeIO(input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVolumeIOOutput), nil
	}
	out, err := c.client.EnableVolumeIO(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVolumeIOWithContext(ctx context.Context, input *ec2.EnableVolumeIOInput, opts ...request.Option) (*ec2.EnableVolumeIOOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVolumeIOOutput), nil
	}
	out, err := c.client.EnableVolumeIOWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVpcClassicLink(input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkOutput), nil
	}
	out, err := c.client.EnableVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVpcClassicLinkDnsSupport(input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.EnableVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.client.EnableVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) EnableVpcClassicLinkWithContext(ctx context.Context, input *ec2.EnableVpcClassicLinkInput, opts ...request.Option) (*ec2.EnableVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.EnableVpcClassicLinkOutput), nil
	}
	out, err := c.client.EnableVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportClientVpnClientCertificateRevocationList(input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.client.ExportClientVpnClientCertificateRevocationList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportClientVpnClientCertificateRevocationListWithContext(ctx context.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput, opts ...request.Option) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.client.ExportClientVpnClientCertificateRevocationListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportClientVpnClientConfiguration(input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientConfigurationOutput), nil
	}
	out, err := c.client.ExportClientVpnClientConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportClientVpnClientConfigurationWithContext(ctx context.Context, input *ec2.ExportClientVpnClientConfigurationInput, opts ...request.Option) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportClientVpnClientConfigurationOutput), nil
	}
	out, err := c.client.ExportClientVpnClientConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportImage(input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportImageOutput), nil
	}
	out, err := c.client.ExportImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportImageWithContext(ctx context.Context, input *ec2.ExportImageInput, opts ...request.Option) (*ec2.ExportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportImageOutput), nil
	}
	out, err := c.client.ExportImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportTransitGatewayRoutes(input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportTransitGatewayRoutesOutput), nil
	}
	out, err := c.client.ExportTransitGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ExportTransitGatewayRoutesWithContext(ctx context.Context, input *ec2.ExportTransitGatewayRoutesInput, opts ...request.Option) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ExportTransitGatewayRoutesOutput), nil
	}
	out, err := c.client.ExportTransitGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAssociatedEnclaveCertificateIamRoles(input *ec2.GetAssociatedEnclaveCertificateIamRolesInput) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.client.GetAssociatedEnclaveCertificateIamRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAssociatedEnclaveCertificateIamRolesWithContext(ctx context.Context, input *ec2.GetAssociatedEnclaveCertificateIamRolesInput, opts ...request.Option) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.client.GetAssociatedEnclaveCertificateIamRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAssociatedIpv6PoolCidrs(input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.client.GetAssociatedIpv6PoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAssociatedIpv6PoolCidrsPages(input *ec2.GetAssociatedIpv6PoolCidrsInput, fn func(*ec2.GetAssociatedIpv6PoolCidrsOutput, bool) bool) error {
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
	if err := c.client.GetAssociatedIpv6PoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetAssociatedIpv6PoolCidrsWithContext(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput, opts ...request.Option) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.client.GetAssociatedIpv6PoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAwsNetworkPerformanceData(input *ec2.GetAwsNetworkPerformanceDataInput) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.client.GetAwsNetworkPerformanceData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetAwsNetworkPerformanceDataPages(input *ec2.GetAwsNetworkPerformanceDataInput, fn func(*ec2.GetAwsNetworkPerformanceDataOutput, bool) bool) error {
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
	if err := c.client.GetAwsNetworkPerformanceDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetAwsNetworkPerformanceDataWithContext(ctx context.Context, input *ec2.GetAwsNetworkPerformanceDataInput, opts ...request.Option) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.client.GetAwsNetworkPerformanceDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetCapacityReservationUsage(input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.client.GetCapacityReservationUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetCapacityReservationUsageWithContext(ctx context.Context, input *ec2.GetCapacityReservationUsageInput, opts ...request.Option) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.client.GetCapacityReservationUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetCoipPoolUsage(input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.client.GetCoipPoolUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetCoipPoolUsageWithContext(ctx context.Context, input *ec2.GetCoipPoolUsageInput, opts ...request.Option) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.client.GetCoipPoolUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetConsoleOutput(input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.client.GetConsoleOutput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetConsoleOutputWithContext(ctx context.Context, input *ec2.GetConsoleOutputInput, opts ...request.Option) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.client.GetConsoleOutputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetConsoleScreenshot(input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.client.GetConsoleScreenshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetConsoleScreenshotWithContext(ctx context.Context, input *ec2.GetConsoleScreenshotInput, opts ...request.Option) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.client.GetConsoleScreenshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetDefaultCreditSpecification(input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.client.GetDefaultCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetDefaultCreditSpecificationWithContext(ctx context.Context, input *ec2.GetDefaultCreditSpecificationInput, opts ...request.Option) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.client.GetDefaultCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetEbsDefaultKmsKeyId(input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.GetEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.GetEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.GetEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetEbsEncryptionByDefault(input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.GetEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.GetEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.client.GetEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetFlowLogsIntegrationTemplate(input *ec2.GetFlowLogsIntegrationTemplateInput) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.client.GetFlowLogsIntegrationTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetFlowLogsIntegrationTemplateWithContext(ctx context.Context, input *ec2.GetFlowLogsIntegrationTemplateInput, opts ...request.Option) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.client.GetFlowLogsIntegrationTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetGroupsForCapacityReservation(input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.client.GetGroupsForCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetGroupsForCapacityReservationPages(input *ec2.GetGroupsForCapacityReservationInput, fn func(*ec2.GetGroupsForCapacityReservationOutput, bool) bool) error {
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
	if err := c.client.GetGroupsForCapacityReservationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetGroupsForCapacityReservationWithContext(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput, opts ...request.Option) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.client.GetGroupsForCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetHostReservationPurchasePreview(input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.client.GetHostReservationPurchasePreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetHostReservationPurchasePreviewWithContext(ctx context.Context, input *ec2.GetHostReservationPurchasePreviewInput, opts ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.client.GetHostReservationPurchasePreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetInstanceTypesFromInstanceRequirements(input *ec2.GetInstanceTypesFromInstanceRequirementsInput) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.client.GetInstanceTypesFromInstanceRequirements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetInstanceTypesFromInstanceRequirementsPages(input *ec2.GetInstanceTypesFromInstanceRequirementsInput, fn func(*ec2.GetInstanceTypesFromInstanceRequirementsOutput, bool) bool) error {
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
	if err := c.client.GetInstanceTypesFromInstanceRequirementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetInstanceTypesFromInstanceRequirementsWithContext(ctx context.Context, input *ec2.GetInstanceTypesFromInstanceRequirementsInput, opts ...request.Option) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.client.GetInstanceTypesFromInstanceRequirementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetInstanceUefiData(input *ec2.GetInstanceUefiDataInput) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.client.GetInstanceUefiData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetInstanceUefiDataWithContext(ctx context.Context, input *ec2.GetInstanceUefiDataInput, opts ...request.Option) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.client.GetInstanceUefiDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamAddressHistory(input *ec2.GetIpamAddressHistoryInput) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.client.GetIpamAddressHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamAddressHistoryPages(input *ec2.GetIpamAddressHistoryInput, fn func(*ec2.GetIpamAddressHistoryOutput, bool) bool) error {
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
	if err := c.client.GetIpamAddressHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetIpamAddressHistoryWithContext(ctx context.Context, input *ec2.GetIpamAddressHistoryInput, opts ...request.Option) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.client.GetIpamAddressHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamPoolAllocations(input *ec2.GetIpamPoolAllocationsInput) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.client.GetIpamPoolAllocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamPoolAllocationsPages(input *ec2.GetIpamPoolAllocationsInput, fn func(*ec2.GetIpamPoolAllocationsOutput, bool) bool) error {
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
	if err := c.client.GetIpamPoolAllocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetIpamPoolAllocationsWithContext(ctx context.Context, input *ec2.GetIpamPoolAllocationsInput, opts ...request.Option) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.client.GetIpamPoolAllocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamPoolCidrs(input *ec2.GetIpamPoolCidrsInput) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.client.GetIpamPoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamPoolCidrsPages(input *ec2.GetIpamPoolCidrsInput, fn func(*ec2.GetIpamPoolCidrsOutput, bool) bool) error {
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
	if err := c.client.GetIpamPoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetIpamPoolCidrsWithContext(ctx context.Context, input *ec2.GetIpamPoolCidrsInput, opts ...request.Option) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.client.GetIpamPoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamResourceCidrs(input *ec2.GetIpamResourceCidrsInput) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.client.GetIpamResourceCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetIpamResourceCidrsPages(input *ec2.GetIpamResourceCidrsInput, fn func(*ec2.GetIpamResourceCidrsOutput, bool) bool) error {
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
	if err := c.client.GetIpamResourceCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetIpamResourceCidrsWithContext(ctx context.Context, input *ec2.GetIpamResourceCidrsInput, opts ...request.Option) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.client.GetIpamResourceCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetLaunchTemplateData(input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.client.GetLaunchTemplateData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetLaunchTemplateDataWithContext(ctx context.Context, input *ec2.GetLaunchTemplateDataInput, opts ...request.Option) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.client.GetLaunchTemplateDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetManagedPrefixListAssociations(input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.client.GetManagedPrefixListAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetManagedPrefixListAssociationsPages(input *ec2.GetManagedPrefixListAssociationsInput, fn func(*ec2.GetManagedPrefixListAssociationsOutput, bool) bool) error {
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
	if err := c.client.GetManagedPrefixListAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetManagedPrefixListAssociationsWithContext(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput, opts ...request.Option) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.client.GetManagedPrefixListAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetManagedPrefixListEntries(input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.client.GetManagedPrefixListEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetManagedPrefixListEntriesPages(input *ec2.GetManagedPrefixListEntriesInput, fn func(*ec2.GetManagedPrefixListEntriesOutput, bool) bool) error {
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
	if err := c.client.GetManagedPrefixListEntriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetManagedPrefixListEntriesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput, opts ...request.Option) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.client.GetManagedPrefixListEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetNetworkInsightsAccessScopeAnalysisFindings(input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.client.GetNetworkInsightsAccessScopeAnalysisFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.client.GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetNetworkInsightsAccessScopeContent(input *ec2.GetNetworkInsightsAccessScopeContentInput) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.client.GetNetworkInsightsAccessScopeContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetNetworkInsightsAccessScopeContentWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeContentInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.client.GetNetworkInsightsAccessScopeContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetPasswordData(input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.client.GetPasswordData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetPasswordDataWithContext(ctx context.Context, input *ec2.GetPasswordDataInput, opts ...request.Option) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.client.GetPasswordDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetReservedInstancesExchangeQuote(input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.client.GetReservedInstancesExchangeQuote(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetReservedInstancesExchangeQuoteWithContext(ctx context.Context, input *ec2.GetReservedInstancesExchangeQuoteInput, opts ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.client.GetReservedInstancesExchangeQuoteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSerialConsoleAccessStatus(input *ec2.GetSerialConsoleAccessStatusInput) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.client.GetSerialConsoleAccessStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSerialConsoleAccessStatusWithContext(ctx context.Context, input *ec2.GetSerialConsoleAccessStatusInput, opts ...request.Option) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.client.GetSerialConsoleAccessStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSpotPlacementScores(input *ec2.GetSpotPlacementScoresInput) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.client.GetSpotPlacementScores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSpotPlacementScoresPages(input *ec2.GetSpotPlacementScoresInput, fn func(*ec2.GetSpotPlacementScoresOutput, bool) bool) error {
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
	if err := c.client.GetSpotPlacementScoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetSpotPlacementScoresWithContext(ctx context.Context, input *ec2.GetSpotPlacementScoresInput, opts ...request.Option) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.client.GetSpotPlacementScoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSubnetCidrReservations(input *ec2.GetSubnetCidrReservationsInput) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.client.GetSubnetCidrReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetSubnetCidrReservationsWithContext(ctx context.Context, input *ec2.GetSubnetCidrReservationsInput, opts ...request.Option) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.client.GetSubnetCidrReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayAttachmentPropagations(input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayAttachmentPropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayAttachmentPropagationsPages(input *ec2.GetTransitGatewayAttachmentPropagationsInput, fn func(*ec2.GetTransitGatewayAttachmentPropagationsOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayAttachmentPropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayAttachmentPropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayAttachmentPropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayMulticastDomainAssociations(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayMulticastDomainAssociationsPages(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, fn func(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayMulticastDomainAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPolicyTableAssociations(input *ec2.GetTransitGatewayPolicyTableAssociationsInput) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayPolicyTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPolicyTableAssociationsPages(input *ec2.GetTransitGatewayPolicyTableAssociationsInput, fn func(*ec2.GetTransitGatewayPolicyTableAssociationsOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayPolicyTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayPolicyTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayPolicyTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPolicyTableEntries(input *ec2.GetTransitGatewayPolicyTableEntriesInput) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.client.GetTransitGatewayPolicyTableEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPolicyTableEntriesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableEntriesInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.client.GetTransitGatewayPolicyTableEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPrefixListReferences(input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.client.GetTransitGatewayPrefixListReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayPrefixListReferencesPages(input *ec2.GetTransitGatewayPrefixListReferencesInput, fn func(*ec2.GetTransitGatewayPrefixListReferencesOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayPrefixListReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayPrefixListReferencesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput, opts ...request.Option) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.client.GetTransitGatewayPrefixListReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayRouteTableAssociations(input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayRouteTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayRouteTableAssociationsPages(input *ec2.GetTransitGatewayRouteTableAssociationsInput, fn func(*ec2.GetTransitGatewayRouteTableAssociationsOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayRouteTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayRouteTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayRouteTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayRouteTablePropagations(input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayRouteTablePropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetTransitGatewayRouteTablePropagationsPages(input *ec2.GetTransitGatewayRouteTablePropagationsInput, fn func(*ec2.GetTransitGatewayRouteTablePropagationsOutput, bool) bool) error {
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
	if err := c.client.GetTransitGatewayRouteTablePropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetTransitGatewayRouteTablePropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.client.GetTransitGatewayRouteTablePropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVerifiedAccessEndpointPolicy(input *ec2.GetVerifiedAccessEndpointPolicyInput) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.client.GetVerifiedAccessEndpointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVerifiedAccessEndpointPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessEndpointPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.client.GetVerifiedAccessEndpointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVerifiedAccessGroupPolicy(input *ec2.GetVerifiedAccessGroupPolicyInput) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.client.GetVerifiedAccessGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVerifiedAccessGroupPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessGroupPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.client.GetVerifiedAccessGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVpnConnectionDeviceSampleConfiguration(input *ec2.GetVpnConnectionDeviceSampleConfigurationInput) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.client.GetVpnConnectionDeviceSampleConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVpnConnectionDeviceSampleConfigurationWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceSampleConfigurationInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.client.GetVpnConnectionDeviceSampleConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVpnConnectionDeviceTypes(input *ec2.GetVpnConnectionDeviceTypesInput) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.client.GetVpnConnectionDeviceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) GetVpnConnectionDeviceTypesPages(input *ec2.GetVpnConnectionDeviceTypesInput, fn func(*ec2.GetVpnConnectionDeviceTypesOutput, bool) bool) error {
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
	if err := c.client.GetVpnConnectionDeviceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) GetVpnConnectionDeviceTypesWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceTypesInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.client.GetVpnConnectionDeviceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportClientVpnClientCertificateRevocationList(input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.client.ImportClientVpnClientCertificateRevocationList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportClientVpnClientCertificateRevocationListWithContext(ctx context.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput, opts ...request.Option) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportClientVpnClientCertificateRevocationListOutput), nil
	}
	out, err := c.client.ImportClientVpnClientCertificateRevocationListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportImage(input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportImageOutput), nil
	}
	out, err := c.client.ImportImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportImageWithContext(ctx context.Context, input *ec2.ImportImageInput, opts ...request.Option) (*ec2.ImportImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportImageOutput), nil
	}
	out, err := c.client.ImportImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportInstance(input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportInstanceOutput), nil
	}
	out, err := c.client.ImportInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportInstanceWithContext(ctx context.Context, input *ec2.ImportInstanceInput, opts ...request.Option) (*ec2.ImportInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportInstanceOutput), nil
	}
	out, err := c.client.ImportInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportKeyPair(input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportKeyPairOutput), nil
	}
	out, err := c.client.ImportKeyPair(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportKeyPairWithContext(ctx context.Context, input *ec2.ImportKeyPairInput, opts ...request.Option) (*ec2.ImportKeyPairOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportKeyPairOutput), nil
	}
	out, err := c.client.ImportKeyPairWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportSnapshot(input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportSnapshotOutput), nil
	}
	out, err := c.client.ImportSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportSnapshotWithContext(ctx context.Context, input *ec2.ImportSnapshotInput, opts ...request.Option) (*ec2.ImportSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportSnapshotOutput), nil
	}
	out, err := c.client.ImportSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportVolume(input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportVolumeOutput), nil
	}
	out, err := c.client.ImportVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ImportVolumeWithContext(ctx context.Context, input *ec2.ImportVolumeInput, opts ...request.Option) (*ec2.ImportVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ImportVolumeOutput), nil
	}
	out, err := c.client.ImportVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ListImagesInRecycleBin(input *ec2.ListImagesInRecycleBinInput) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.client.ListImagesInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ListImagesInRecycleBinPages(input *ec2.ListImagesInRecycleBinInput, fn func(*ec2.ListImagesInRecycleBinOutput, bool) bool) error {
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
	if err := c.client.ListImagesInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) ListImagesInRecycleBinWithContext(ctx context.Context, input *ec2.ListImagesInRecycleBinInput, opts ...request.Option) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.client.ListImagesInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ListSnapshotsInRecycleBin(input *ec2.ListSnapshotsInRecycleBinInput) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.client.ListSnapshotsInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ListSnapshotsInRecycleBinPages(input *ec2.ListSnapshotsInRecycleBinInput, fn func(*ec2.ListSnapshotsInRecycleBinOutput, bool) bool) error {
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
	if err := c.client.ListSnapshotsInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) ListSnapshotsInRecycleBinWithContext(ctx context.Context, input *ec2.ListSnapshotsInRecycleBinInput, opts ...request.Option) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.client.ListSnapshotsInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyAddressAttribute(input *ec2.ModifyAddressAttributeInput) (*ec2.ModifyAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAddressAttributeOutput), nil
	}
	out, err := c.client.ModifyAddressAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyAddressAttributeWithContext(ctx context.Context, input *ec2.ModifyAddressAttributeInput, opts ...request.Option) (*ec2.ModifyAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAddressAttributeOutput), nil
	}
	out, err := c.client.ModifyAddressAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyAvailabilityZoneGroup(input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAvailabilityZoneGroupOutput), nil
	}
	out, err := c.client.ModifyAvailabilityZoneGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyAvailabilityZoneGroupWithContext(ctx context.Context, input *ec2.ModifyAvailabilityZoneGroupInput, opts ...request.Option) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyAvailabilityZoneGroupOutput), nil
	}
	out, err := c.client.ModifyAvailabilityZoneGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyCapacityReservation(input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationOutput), nil
	}
	out, err := c.client.ModifyCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyCapacityReservationFleet(input *ec2.ModifyCapacityReservationFleetInput) (*ec2.ModifyCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationFleetOutput), nil
	}
	out, err := c.client.ModifyCapacityReservationFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyCapacityReservationFleetWithContext(ctx context.Context, input *ec2.ModifyCapacityReservationFleetInput, opts ...request.Option) (*ec2.ModifyCapacityReservationFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationFleetOutput), nil
	}
	out, err := c.client.ModifyCapacityReservationFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyCapacityReservationWithContext(ctx context.Context, input *ec2.ModifyCapacityReservationInput, opts ...request.Option) (*ec2.ModifyCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyCapacityReservationOutput), nil
	}
	out, err := c.client.ModifyCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyClientVpnEndpoint(input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyClientVpnEndpointOutput), nil
	}
	out, err := c.client.ModifyClientVpnEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyClientVpnEndpointWithContext(ctx context.Context, input *ec2.ModifyClientVpnEndpointInput, opts ...request.Option) (*ec2.ModifyClientVpnEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyClientVpnEndpointOutput), nil
	}
	out, err := c.client.ModifyClientVpnEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyDefaultCreditSpecification(input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyDefaultCreditSpecificationOutput), nil
	}
	out, err := c.client.ModifyDefaultCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyDefaultCreditSpecificationWithContext(ctx context.Context, input *ec2.ModifyDefaultCreditSpecificationInput, opts ...request.Option) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyDefaultCreditSpecificationOutput), nil
	}
	out, err := c.client.ModifyDefaultCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyEbsDefaultKmsKeyId(input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.ModifyEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.ModifyEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyFleet(input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFleetOutput), nil
	}
	out, err := c.client.ModifyFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyFleetWithContext(ctx context.Context, input *ec2.ModifyFleetInput, opts ...request.Option) (*ec2.ModifyFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFleetOutput), nil
	}
	out, err := c.client.ModifyFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyFpgaImageAttribute(input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFpgaImageAttributeOutput), nil
	}
	out, err := c.client.ModifyFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyFpgaImageAttributeWithContext(ctx context.Context, input *ec2.ModifyFpgaImageAttributeInput, opts ...request.Option) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyFpgaImageAttributeOutput), nil
	}
	out, err := c.client.ModifyFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyHosts(input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyHostsOutput), nil
	}
	out, err := c.client.ModifyHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyHostsWithContext(ctx context.Context, input *ec2.ModifyHostsInput, opts ...request.Option) (*ec2.ModifyHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyHostsOutput), nil
	}
	out, err := c.client.ModifyHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIdFormat(input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdFormatOutput), nil
	}
	out, err := c.client.ModifyIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIdFormatWithContext(ctx context.Context, input *ec2.ModifyIdFormatInput, opts ...request.Option) (*ec2.ModifyIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdFormatOutput), nil
	}
	out, err := c.client.ModifyIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIdentityIdFormat(input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdentityIdFormatOutput), nil
	}
	out, err := c.client.ModifyIdentityIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIdentityIdFormatWithContext(ctx context.Context, input *ec2.ModifyIdentityIdFormatInput, opts ...request.Option) (*ec2.ModifyIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIdentityIdFormatOutput), nil
	}
	out, err := c.client.ModifyIdentityIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyImageAttribute(input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyImageAttributeOutput), nil
	}
	out, err := c.client.ModifyImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, input *ec2.ModifyImageAttributeInput, opts ...request.Option) (*ec2.ModifyImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyImageAttributeOutput), nil
	}
	out, err := c.client.ModifyImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceAttribute(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceAttributeOutput), nil
	}
	out, err := c.client.ModifyInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceAttributeWithContext(ctx context.Context, input *ec2.ModifyInstanceAttributeInput, opts ...request.Option) (*ec2.ModifyInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceAttributeOutput), nil
	}
	out, err := c.client.ModifyInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceCapacityReservationAttributes(input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCapacityReservationAttributesOutput), nil
	}
	out, err := c.client.ModifyInstanceCapacityReservationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceCapacityReservationAttributesWithContext(ctx context.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput, opts ...request.Option) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCapacityReservationAttributesOutput), nil
	}
	out, err := c.client.ModifyInstanceCapacityReservationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceCreditSpecification(input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCreditSpecificationOutput), nil
	}
	out, err := c.client.ModifyInstanceCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceCreditSpecificationWithContext(ctx context.Context, input *ec2.ModifyInstanceCreditSpecificationInput, opts ...request.Option) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceCreditSpecificationOutput), nil
	}
	out, err := c.client.ModifyInstanceCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceEventStartTime(input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventStartTimeOutput), nil
	}
	out, err := c.client.ModifyInstanceEventStartTime(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceEventStartTimeWithContext(ctx context.Context, input *ec2.ModifyInstanceEventStartTimeInput, opts ...request.Option) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventStartTimeOutput), nil
	}
	out, err := c.client.ModifyInstanceEventStartTimeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceEventWindow(input *ec2.ModifyInstanceEventWindowInput) (*ec2.ModifyInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventWindowOutput), nil
	}
	out, err := c.client.ModifyInstanceEventWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceEventWindowWithContext(ctx context.Context, input *ec2.ModifyInstanceEventWindowInput, opts ...request.Option) (*ec2.ModifyInstanceEventWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceEventWindowOutput), nil
	}
	out, err := c.client.ModifyInstanceEventWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceMaintenanceOptions(input *ec2.ModifyInstanceMaintenanceOptionsInput) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMaintenanceOptionsOutput), nil
	}
	out, err := c.client.ModifyInstanceMaintenanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceMaintenanceOptionsWithContext(ctx context.Context, input *ec2.ModifyInstanceMaintenanceOptionsInput, opts ...request.Option) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMaintenanceOptionsOutput), nil
	}
	out, err := c.client.ModifyInstanceMaintenanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceMetadataOptions(input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMetadataOptionsOutput), nil
	}
	out, err := c.client.ModifyInstanceMetadataOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstanceMetadataOptionsWithContext(ctx context.Context, input *ec2.ModifyInstanceMetadataOptionsInput, opts ...request.Option) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstanceMetadataOptionsOutput), nil
	}
	out, err := c.client.ModifyInstanceMetadataOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstancePlacement(input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstancePlacementOutput), nil
	}
	out, err := c.client.ModifyInstancePlacement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyInstancePlacementWithContext(ctx context.Context, input *ec2.ModifyInstancePlacementInput, opts ...request.Option) (*ec2.ModifyInstancePlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyInstancePlacementOutput), nil
	}
	out, err := c.client.ModifyInstancePlacementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpam(input *ec2.ModifyIpamInput) (*ec2.ModifyIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamOutput), nil
	}
	out, err := c.client.ModifyIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamPool(input *ec2.ModifyIpamPoolInput) (*ec2.ModifyIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamPoolOutput), nil
	}
	out, err := c.client.ModifyIpamPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamPoolWithContext(ctx context.Context, input *ec2.ModifyIpamPoolInput, opts ...request.Option) (*ec2.ModifyIpamPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamPoolOutput), nil
	}
	out, err := c.client.ModifyIpamPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamResourceCidr(input *ec2.ModifyIpamResourceCidrInput) (*ec2.ModifyIpamResourceCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamResourceCidrOutput), nil
	}
	out, err := c.client.ModifyIpamResourceCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamResourceCidrWithContext(ctx context.Context, input *ec2.ModifyIpamResourceCidrInput, opts ...request.Option) (*ec2.ModifyIpamResourceCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamResourceCidrOutput), nil
	}
	out, err := c.client.ModifyIpamResourceCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamScope(input *ec2.ModifyIpamScopeInput) (*ec2.ModifyIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamScopeOutput), nil
	}
	out, err := c.client.ModifyIpamScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamScopeWithContext(ctx context.Context, input *ec2.ModifyIpamScopeInput, opts ...request.Option) (*ec2.ModifyIpamScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamScopeOutput), nil
	}
	out, err := c.client.ModifyIpamScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyIpamWithContext(ctx context.Context, input *ec2.ModifyIpamInput, opts ...request.Option) (*ec2.ModifyIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyIpamOutput), nil
	}
	out, err := c.client.ModifyIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyLaunchTemplate(input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLaunchTemplateOutput), nil
	}
	out, err := c.client.ModifyLaunchTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyLaunchTemplateWithContext(ctx context.Context, input *ec2.ModifyLaunchTemplateInput, opts ...request.Option) (*ec2.ModifyLaunchTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLaunchTemplateOutput), nil
	}
	out, err := c.client.ModifyLaunchTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyLocalGatewayRoute(input *ec2.ModifyLocalGatewayRouteInput) (*ec2.ModifyLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLocalGatewayRouteOutput), nil
	}
	out, err := c.client.ModifyLocalGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyLocalGatewayRouteWithContext(ctx context.Context, input *ec2.ModifyLocalGatewayRouteInput, opts ...request.Option) (*ec2.ModifyLocalGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyLocalGatewayRouteOutput), nil
	}
	out, err := c.client.ModifyLocalGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyManagedPrefixList(input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyManagedPrefixListOutput), nil
	}
	out, err := c.client.ModifyManagedPrefixList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyManagedPrefixListWithContext(ctx context.Context, input *ec2.ModifyManagedPrefixListInput, opts ...request.Option) (*ec2.ModifyManagedPrefixListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyManagedPrefixListOutput), nil
	}
	out, err := c.client.ModifyManagedPrefixListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyNetworkInterfaceAttribute(input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.ModifyNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.ModifyNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.ModifyNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyPrivateDnsNameOptions(input *ec2.ModifyPrivateDnsNameOptionsInput) (*ec2.ModifyPrivateDnsNameOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyPrivateDnsNameOptionsOutput), nil
	}
	out, err := c.client.ModifyPrivateDnsNameOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyPrivateDnsNameOptionsWithContext(ctx context.Context, input *ec2.ModifyPrivateDnsNameOptionsInput, opts ...request.Option) (*ec2.ModifyPrivateDnsNameOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyPrivateDnsNameOptionsOutput), nil
	}
	out, err := c.client.ModifyPrivateDnsNameOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyReservedInstances(input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyReservedInstancesOutput), nil
	}
	out, err := c.client.ModifyReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyReservedInstancesWithContext(ctx context.Context, input *ec2.ModifyReservedInstancesInput, opts ...request.Option) (*ec2.ModifyReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyReservedInstancesOutput), nil
	}
	out, err := c.client.ModifyReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySecurityGroupRules(input *ec2.ModifySecurityGroupRulesInput) (*ec2.ModifySecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySecurityGroupRulesOutput), nil
	}
	out, err := c.client.ModifySecurityGroupRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySecurityGroupRulesWithContext(ctx context.Context, input *ec2.ModifySecurityGroupRulesInput, opts ...request.Option) (*ec2.ModifySecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySecurityGroupRulesOutput), nil
	}
	out, err := c.client.ModifySecurityGroupRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySnapshotAttribute(input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotAttributeOutput), nil
	}
	out, err := c.client.ModifySnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySnapshotAttributeWithContext(ctx context.Context, input *ec2.ModifySnapshotAttributeInput, opts ...request.Option) (*ec2.ModifySnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotAttributeOutput), nil
	}
	out, err := c.client.ModifySnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySnapshotTier(input *ec2.ModifySnapshotTierInput) (*ec2.ModifySnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotTierOutput), nil
	}
	out, err := c.client.ModifySnapshotTier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySnapshotTierWithContext(ctx context.Context, input *ec2.ModifySnapshotTierInput, opts ...request.Option) (*ec2.ModifySnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySnapshotTierOutput), nil
	}
	out, err := c.client.ModifySnapshotTierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySpotFleetRequestWithContext(ctx context.Context, input *ec2.ModifySpotFleetRequestInput, opts ...request.Option) (*ec2.ModifySpotFleetRequestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySpotFleetRequestOutput), nil
	}
	out, err := c.client.ModifySpotFleetRequestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySubnetAttribute(input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySubnetAttributeOutput), nil
	}
	out, err := c.client.ModifySubnetAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifySubnetAttributeWithContext(ctx context.Context, input *ec2.ModifySubnetAttributeInput, opts ...request.Option) (*ec2.ModifySubnetAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifySubnetAttributeOutput), nil
	}
	out, err := c.client.ModifySubnetAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorFilterNetworkServices(input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorFilterNetworkServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorFilterRule(input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorFilterRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorFilterRuleWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorFilterRuleOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorFilterRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorSession(input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTrafficMirrorSessionWithContext(ctx context.Context, input *ec2.ModifyTrafficMirrorSessionInput, opts ...request.Option) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTrafficMirrorSessionOutput), nil
	}
	out, err := c.client.ModifyTrafficMirrorSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGateway(input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayOutput), nil
	}
	out, err := c.client.ModifyTransitGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGatewayPrefixListReference(input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.ModifyTransitGatewayPrefixListReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGatewayPrefixListReferenceWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput, opts ...request.Option) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayPrefixListReferenceOutput), nil
	}
	out, err := c.client.ModifyTransitGatewayPrefixListReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGatewayVpcAttachment(input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.ModifyTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.ModifyTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyTransitGatewayWithContext(ctx context.Context, input *ec2.ModifyTransitGatewayInput, opts ...request.Option) (*ec2.ModifyTransitGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyTransitGatewayOutput), nil
	}
	out, err := c.client.ModifyTransitGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessEndpoint(input *ec2.ModifyVerifiedAccessEndpointInput) (*ec2.ModifyVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessEndpointPolicy(input *ec2.ModifyVerifiedAccessEndpointPolicyInput) (*ec2.ModifyVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessEndpointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessEndpointPolicyWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessEndpointPolicyInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessEndpointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessEndpointWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessEndpointInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessEndpointOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessGroup(input *ec2.ModifyVerifiedAccessGroupInput) (*ec2.ModifyVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessGroupPolicy(input *ec2.ModifyVerifiedAccessGroupPolicyInput) (*ec2.ModifyVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessGroupPolicyWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessGroupPolicyInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessGroupWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessGroupInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessGroupOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessInstance(input *ec2.ModifyVerifiedAccessInstanceInput) (*ec2.ModifyVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessInstanceLoggingConfiguration(input *ec2.ModifyVerifiedAccessInstanceLoggingConfigurationInput) (*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessInstanceLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessInstanceLoggingConfigurationWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessInstanceLoggingConfigurationInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceLoggingConfigurationOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessInstanceLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessInstanceWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessInstanceInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessInstanceOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessTrustProvider(input *ec2.ModifyVerifiedAccessTrustProviderInput) (*ec2.ModifyVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessTrustProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVerifiedAccessTrustProviderWithContext(ctx context.Context, input *ec2.ModifyVerifiedAccessTrustProviderInput, opts ...request.Option) (*ec2.ModifyVerifiedAccessTrustProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVerifiedAccessTrustProviderOutput), nil
	}
	out, err := c.client.ModifyVerifiedAccessTrustProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVolume(input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeOutput), nil
	}
	out, err := c.client.ModifyVolume(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVolumeAttribute(input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeAttributeOutput), nil
	}
	out, err := c.client.ModifyVolumeAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVolumeAttributeWithContext(ctx context.Context, input *ec2.ModifyVolumeAttributeInput, opts ...request.Option) (*ec2.ModifyVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeAttributeOutput), nil
	}
	out, err := c.client.ModifyVolumeAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVolumeWithContext(ctx context.Context, input *ec2.ModifyVolumeInput, opts ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVolumeOutput), nil
	}
	out, err := c.client.ModifyVolumeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcAttribute(input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcAttributeOutput), nil
	}
	out, err := c.client.ModifyVpcAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcAttributeWithContext(ctx context.Context, input *ec2.ModifyVpcAttributeInput, opts ...request.Option) (*ec2.ModifyVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcAttributeOutput), nil
	}
	out, err := c.client.ModifyVpcAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpoint(input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointOutput), nil
	}
	out, err := c.client.ModifyVpcEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointConnectionNotification(input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointConnectionNotification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointConnectionNotificationWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput, opts ...request.Option) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointConnectionNotificationOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointConnectionNotificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServiceConfiguration(input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServiceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServiceConfigurationWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServiceConfigurationOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServiceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServicePayerResponsibility(input *ec2.ModifyVpcEndpointServicePayerResponsibilityInput) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServicePayerResponsibility(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServicePayerResponsibilityWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServicePayerResponsibilityInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServicePayerResponsibilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServicePermissions(input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServicePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointServicePermissionsWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput, opts ...request.Option) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointServicePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcEndpointWithContext(ctx context.Context, input *ec2.ModifyVpcEndpointInput, opts ...request.Option) (*ec2.ModifyVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcEndpointOutput), nil
	}
	out, err := c.client.ModifyVpcEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcPeeringConnectionOptions(input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcPeeringConnectionOptionsOutput), nil
	}
	out, err := c.client.ModifyVpcPeeringConnectionOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcPeeringConnectionOptionsWithContext(ctx context.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput, opts ...request.Option) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcPeeringConnectionOptionsOutput), nil
	}
	out, err := c.client.ModifyVpcPeeringConnectionOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcTenancy(input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcTenancyOutput), nil
	}
	out, err := c.client.ModifyVpcTenancy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpcTenancyWithContext(ctx context.Context, input *ec2.ModifyVpcTenancyInput, opts ...request.Option) (*ec2.ModifyVpcTenancyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpcTenancyOutput), nil
	}
	out, err := c.client.ModifyVpcTenancyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnConnection(input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOutput), nil
	}
	out, err := c.client.ModifyVpnConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnConnectionOptions(input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOptionsOutput), nil
	}
	out, err := c.client.ModifyVpnConnectionOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnConnectionOptionsWithContext(ctx context.Context, input *ec2.ModifyVpnConnectionOptionsInput, opts ...request.Option) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOptionsOutput), nil
	}
	out, err := c.client.ModifyVpnConnectionOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnConnectionWithContext(ctx context.Context, input *ec2.ModifyVpnConnectionInput, opts ...request.Option) (*ec2.ModifyVpnConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnConnectionOutput), nil
	}
	out, err := c.client.ModifyVpnConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnTunnelCertificate(input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelCertificateOutput), nil
	}
	out, err := c.client.ModifyVpnTunnelCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnTunnelCertificateWithContext(ctx context.Context, input *ec2.ModifyVpnTunnelCertificateInput, opts ...request.Option) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelCertificateOutput), nil
	}
	out, err := c.client.ModifyVpnTunnelCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnTunnelOptions(input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelOptionsOutput), nil
	}
	out, err := c.client.ModifyVpnTunnelOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ModifyVpnTunnelOptionsWithContext(ctx context.Context, input *ec2.ModifyVpnTunnelOptionsInput, opts ...request.Option) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ModifyVpnTunnelOptionsOutput), nil
	}
	out, err := c.client.ModifyVpnTunnelOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MonitorInstances(input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MonitorInstancesOutput), nil
	}
	out, err := c.client.MonitorInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MonitorInstancesWithContext(ctx context.Context, input *ec2.MonitorInstancesInput, opts ...request.Option) (*ec2.MonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MonitorInstancesOutput), nil
	}
	out, err := c.client.MonitorInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MoveAddressToVpc(input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveAddressToVpcOutput), nil
	}
	out, err := c.client.MoveAddressToVpc(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MoveAddressToVpcWithContext(ctx context.Context, input *ec2.MoveAddressToVpcInput, opts ...request.Option) (*ec2.MoveAddressToVpcOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveAddressToVpcOutput), nil
	}
	out, err := c.client.MoveAddressToVpcWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MoveByoipCidrToIpam(input *ec2.MoveByoipCidrToIpamInput) (*ec2.MoveByoipCidrToIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveByoipCidrToIpamOutput), nil
	}
	out, err := c.client.MoveByoipCidrToIpam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) MoveByoipCidrToIpamWithContext(ctx context.Context, input *ec2.MoveByoipCidrToIpamInput, opts ...request.Option) (*ec2.MoveByoipCidrToIpamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.MoveByoipCidrToIpamOutput), nil
	}
	out, err := c.client.MoveByoipCidrToIpamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionByoipCidr(input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionByoipCidrOutput), nil
	}
	out, err := c.client.ProvisionByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionByoipCidrWithContext(ctx context.Context, input *ec2.ProvisionByoipCidrInput, opts ...request.Option) (*ec2.ProvisionByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionByoipCidrOutput), nil
	}
	out, err := c.client.ProvisionByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionIpamPoolCidr(input *ec2.ProvisionIpamPoolCidrInput) (*ec2.ProvisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionIpamPoolCidrOutput), nil
	}
	out, err := c.client.ProvisionIpamPoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionIpamPoolCidrWithContext(ctx context.Context, input *ec2.ProvisionIpamPoolCidrInput, opts ...request.Option) (*ec2.ProvisionIpamPoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionIpamPoolCidrOutput), nil
	}
	out, err := c.client.ProvisionIpamPoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionPublicIpv4PoolCidr(input *ec2.ProvisionPublicIpv4PoolCidrInput) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.client.ProvisionPublicIpv4PoolCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ProvisionPublicIpv4PoolCidrWithContext(ctx context.Context, input *ec2.ProvisionPublicIpv4PoolCidrInput, opts ...request.Option) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ProvisionPublicIpv4PoolCidrOutput), nil
	}
	out, err := c.client.ProvisionPublicIpv4PoolCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseHostReservation(input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseHostReservationOutput), nil
	}
	out, err := c.client.PurchaseHostReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseHostReservationWithContext(ctx context.Context, input *ec2.PurchaseHostReservationInput, opts ...request.Option) (*ec2.PurchaseHostReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseHostReservationOutput), nil
	}
	out, err := c.client.PurchaseHostReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseReservedInstancesOffering(input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseReservedInstancesOfferingOutput), nil
	}
	out, err := c.client.PurchaseReservedInstancesOffering(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseReservedInstancesOfferingWithContext(ctx context.Context, input *ec2.PurchaseReservedInstancesOfferingInput, opts ...request.Option) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseReservedInstancesOfferingOutput), nil
	}
	out, err := c.client.PurchaseReservedInstancesOfferingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseScheduledInstances(input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseScheduledInstancesOutput), nil
	}
	out, err := c.client.PurchaseScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) PurchaseScheduledInstancesWithContext(ctx context.Context, input *ec2.PurchaseScheduledInstancesInput, opts ...request.Option) (*ec2.PurchaseScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.PurchaseScheduledInstancesOutput), nil
	}
	out, err := c.client.PurchaseScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RebootInstances(input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RebootInstancesOutput), nil
	}
	out, err := c.client.RebootInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RebootInstancesWithContext(ctx context.Context, input *ec2.RebootInstancesInput, opts ...request.Option) (*ec2.RebootInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RebootInstancesOutput), nil
	}
	out, err := c.client.RebootInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterImage(input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterImageOutput), nil
	}
	out, err := c.client.RegisterImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterImageWithContext(ctx context.Context, input *ec2.RegisterImageInput, opts ...request.Option) (*ec2.RegisterImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterImageOutput), nil
	}
	out, err := c.client.RegisterImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterInstanceEventNotificationAttributes(input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.RegisterInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.client.RegisterInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterTransitGatewayMulticastGroupMembers(input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.client.RegisterTransitGatewayMulticastGroupMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterTransitGatewayMulticastGroupMembersWithContext(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput, opts ...request.Option) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupMembersOutput), nil
	}
	out, err := c.client.RegisterTransitGatewayMulticastGroupMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterTransitGatewayMulticastGroupSources(input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.client.RegisterTransitGatewayMulticastGroupSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RegisterTransitGatewayMulticastGroupSourcesWithContext(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput, opts ...request.Option) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput), nil
	}
	out, err := c.client.RegisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayMulticastDomainAssociations(input *ec2.RejectTransitGatewayMulticastDomainAssociationsInput) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.RejectTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.RejectTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.client.RejectTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayPeeringAttachment(input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.RejectTransitGatewayPeeringAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayPeeringAttachmentWithContext(ctx context.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput, opts ...request.Option) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayPeeringAttachmentOutput), nil
	}
	out, err := c.client.RejectTransitGatewayPeeringAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayVpcAttachment(input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.RejectTransitGatewayVpcAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectTransitGatewayVpcAttachmentWithContext(ctx context.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput, opts ...request.Option) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectTransitGatewayVpcAttachmentOutput), nil
	}
	out, err := c.client.RejectTransitGatewayVpcAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectVpcEndpointConnections(input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.RejectVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.RejectVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcEndpointConnectionsOutput), nil
	}
	out, err := c.client.RejectVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectVpcPeeringConnection(input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.RejectVpcPeeringConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RejectVpcPeeringConnectionWithContext(ctx context.Context, input *ec2.RejectVpcPeeringConnectionInput, opts ...request.Option) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RejectVpcPeeringConnectionOutput), nil
	}
	out, err := c.client.RejectVpcPeeringConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseAddress(input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseAddressOutput), nil
	}
	out, err := c.client.ReleaseAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseAddressWithContext(ctx context.Context, input *ec2.ReleaseAddressInput, opts ...request.Option) (*ec2.ReleaseAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseAddressOutput), nil
	}
	out, err := c.client.ReleaseAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseHosts(input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseHostsOutput), nil
	}
	out, err := c.client.ReleaseHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseHostsWithContext(ctx context.Context, input *ec2.ReleaseHostsInput, opts ...request.Option) (*ec2.ReleaseHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseHostsOutput), nil
	}
	out, err := c.client.ReleaseHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseIpamPoolAllocation(input *ec2.ReleaseIpamPoolAllocationInput) (*ec2.ReleaseIpamPoolAllocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseIpamPoolAllocationOutput), nil
	}
	out, err := c.client.ReleaseIpamPoolAllocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReleaseIpamPoolAllocationWithContext(ctx context.Context, input *ec2.ReleaseIpamPoolAllocationInput, opts ...request.Option) (*ec2.ReleaseIpamPoolAllocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReleaseIpamPoolAllocationOutput), nil
	}
	out, err := c.client.ReleaseIpamPoolAllocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceIamInstanceProfileAssociation(input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceIamInstanceProfileAssociationOutput), nil
	}
	out, err := c.client.ReplaceIamInstanceProfileAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceIamInstanceProfileAssociationWithContext(ctx context.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput, opts ...request.Option) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceIamInstanceProfileAssociationOutput), nil
	}
	out, err := c.client.ReplaceIamInstanceProfileAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceNetworkAclAssociation(input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclAssociationOutput), nil
	}
	out, err := c.client.ReplaceNetworkAclAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceNetworkAclAssociationWithContext(ctx context.Context, input *ec2.ReplaceNetworkAclAssociationInput, opts ...request.Option) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclAssociationOutput), nil
	}
	out, err := c.client.ReplaceNetworkAclAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceNetworkAclEntry(input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclEntryOutput), nil
	}
	out, err := c.client.ReplaceNetworkAclEntry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceNetworkAclEntryWithContext(ctx context.Context, input *ec2.ReplaceNetworkAclEntryInput, opts ...request.Option) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceNetworkAclEntryOutput), nil
	}
	out, err := c.client.ReplaceNetworkAclEntryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceRoute(input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteOutput), nil
	}
	out, err := c.client.ReplaceRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceRouteTableAssociation(input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteTableAssociationOutput), nil
	}
	out, err := c.client.ReplaceRouteTableAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceRouteTableAssociationWithContext(ctx context.Context, input *ec2.ReplaceRouteTableAssociationInput, opts ...request.Option) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteTableAssociationOutput), nil
	}
	out, err := c.client.ReplaceRouteTableAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceRouteWithContext(ctx context.Context, input *ec2.ReplaceRouteInput, opts ...request.Option) (*ec2.ReplaceRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceRouteOutput), nil
	}
	out, err := c.client.ReplaceRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceTransitGatewayRoute(input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceTransitGatewayRouteOutput), nil
	}
	out, err := c.client.ReplaceTransitGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReplaceTransitGatewayRouteWithContext(ctx context.Context, input *ec2.ReplaceTransitGatewayRouteInput, opts ...request.Option) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReplaceTransitGatewayRouteOutput), nil
	}
	out, err := c.client.ReplaceTransitGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReportInstanceStatus(input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReportInstanceStatusOutput), nil
	}
	out, err := c.client.ReportInstanceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ReportInstanceStatusWithContext(ctx context.Context, input *ec2.ReportInstanceStatusInput, opts ...request.Option) (*ec2.ReportInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ReportInstanceStatusOutput), nil
	}
	out, err := c.client.ReportInstanceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RequestSpotFleet(input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotFleetOutput), nil
	}
	out, err := c.client.RequestSpotFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RequestSpotFleetWithContext(ctx context.Context, input *ec2.RequestSpotFleetInput, opts ...request.Option) (*ec2.RequestSpotFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotFleetOutput), nil
	}
	out, err := c.client.RequestSpotFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RequestSpotInstances(input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotInstancesOutput), nil
	}
	out, err := c.client.RequestSpotInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RequestSpotInstancesWithContext(ctx context.Context, input *ec2.RequestSpotInstancesInput, opts ...request.Option) (*ec2.RequestSpotInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RequestSpotInstancesOutput), nil
	}
	out, err := c.client.RequestSpotInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetAddressAttribute(input *ec2.ResetAddressAttributeInput) (*ec2.ResetAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetAddressAttributeOutput), nil
	}
	out, err := c.client.ResetAddressAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetAddressAttributeWithContext(ctx context.Context, input *ec2.ResetAddressAttributeInput, opts ...request.Option) (*ec2.ResetAddressAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetAddressAttributeOutput), nil
	}
	out, err := c.client.ResetAddressAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetEbsDefaultKmsKeyId(input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.ResetEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.client.ResetEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetFpgaImageAttribute(input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetFpgaImageAttributeOutput), nil
	}
	out, err := c.client.ResetFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetFpgaImageAttributeWithContext(ctx context.Context, input *ec2.ResetFpgaImageAttributeInput, opts ...request.Option) (*ec2.ResetFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetFpgaImageAttributeOutput), nil
	}
	out, err := c.client.ResetFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetImageAttribute(input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetImageAttributeOutput), nil
	}
	out, err := c.client.ResetImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetImageAttributeWithContext(ctx context.Context, input *ec2.ResetImageAttributeInput, opts ...request.Option) (*ec2.ResetImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetImageAttributeOutput), nil
	}
	out, err := c.client.ResetImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetInstanceAttribute(input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetInstanceAttributeOutput), nil
	}
	out, err := c.client.ResetInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetInstanceAttributeWithContext(ctx context.Context, input *ec2.ResetInstanceAttributeInput, opts ...request.Option) (*ec2.ResetInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetInstanceAttributeOutput), nil
	}
	out, err := c.client.ResetInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetNetworkInterfaceAttribute(input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.ResetNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.ResetNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.client.ResetNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetSnapshotAttribute(input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetSnapshotAttributeOutput), nil
	}
	out, err := c.client.ResetSnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) ResetSnapshotAttributeWithContext(ctx context.Context, input *ec2.ResetSnapshotAttributeInput, opts ...request.Option) (*ec2.ResetSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ResetSnapshotAttributeOutput), nil
	}
	out, err := c.client.ResetSnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreAddressToClassic(input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreAddressToClassicOutput), nil
	}
	out, err := c.client.RestoreAddressToClassic(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreAddressToClassicWithContext(ctx context.Context, input *ec2.RestoreAddressToClassicInput, opts ...request.Option) (*ec2.RestoreAddressToClassicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreAddressToClassicOutput), nil
	}
	out, err := c.client.RestoreAddressToClassicWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreImageFromRecycleBin(input *ec2.RestoreImageFromRecycleBinInput) (*ec2.RestoreImageFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreImageFromRecycleBinOutput), nil
	}
	out, err := c.client.RestoreImageFromRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreImageFromRecycleBinWithContext(ctx context.Context, input *ec2.RestoreImageFromRecycleBinInput, opts ...request.Option) (*ec2.RestoreImageFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreImageFromRecycleBinOutput), nil
	}
	out, err := c.client.RestoreImageFromRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreManagedPrefixListVersion(input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreManagedPrefixListVersionOutput), nil
	}
	out, err := c.client.RestoreManagedPrefixListVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreManagedPrefixListVersionWithContext(ctx context.Context, input *ec2.RestoreManagedPrefixListVersionInput, opts ...request.Option) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreManagedPrefixListVersionOutput), nil
	}
	out, err := c.client.RestoreManagedPrefixListVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreSnapshotFromRecycleBin(input *ec2.RestoreSnapshotFromRecycleBinInput) (*ec2.RestoreSnapshotFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotFromRecycleBinOutput), nil
	}
	out, err := c.client.RestoreSnapshotFromRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreSnapshotFromRecycleBinWithContext(ctx context.Context, input *ec2.RestoreSnapshotFromRecycleBinInput, opts ...request.Option) (*ec2.RestoreSnapshotFromRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotFromRecycleBinOutput), nil
	}
	out, err := c.client.RestoreSnapshotFromRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreSnapshotTier(input *ec2.RestoreSnapshotTierInput) (*ec2.RestoreSnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotTierOutput), nil
	}
	out, err := c.client.RestoreSnapshotTier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RestoreSnapshotTierWithContext(ctx context.Context, input *ec2.RestoreSnapshotTierInput, opts ...request.Option) (*ec2.RestoreSnapshotTierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RestoreSnapshotTierOutput), nil
	}
	out, err := c.client.RestoreSnapshotTierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeClientVpnIngress(input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeClientVpnIngressOutput), nil
	}
	out, err := c.client.RevokeClientVpnIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeClientVpnIngressWithContext(ctx context.Context, input *ec2.RevokeClientVpnIngressInput, opts ...request.Option) (*ec2.RevokeClientVpnIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeClientVpnIngressOutput), nil
	}
	out, err := c.client.RevokeClientVpnIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeSecurityGroupEgress(input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupEgressOutput), nil
	}
	out, err := c.client.RevokeSecurityGroupEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeSecurityGroupEgressWithContext(ctx context.Context, input *ec2.RevokeSecurityGroupEgressInput, opts ...request.Option) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupEgressOutput), nil
	}
	out, err := c.client.RevokeSecurityGroupEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupIngressOutput), nil
	}
	out, err := c.client.RevokeSecurityGroupIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RevokeSecurityGroupIngressWithContext(ctx context.Context, input *ec2.RevokeSecurityGroupIngressInput, opts ...request.Option) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RevokeSecurityGroupIngressOutput), nil
	}
	out, err := c.client.RevokeSecurityGroupIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RunScheduledInstances(input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RunScheduledInstancesOutput), nil
	}
	out, err := c.client.RunScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) RunScheduledInstancesWithContext(ctx context.Context, input *ec2.RunScheduledInstancesInput, opts ...request.Option) (*ec2.RunScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.RunScheduledInstancesOutput), nil
	}
	out, err := c.client.RunScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchLocalGatewayRoutes(input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.client.SearchLocalGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchLocalGatewayRoutesPages(input *ec2.SearchLocalGatewayRoutesInput, fn func(*ec2.SearchLocalGatewayRoutesOutput, bool) bool) error {
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
	if err := c.client.SearchLocalGatewayRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) SearchLocalGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput, opts ...request.Option) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.client.SearchLocalGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchTransitGatewayMulticastGroups(input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.client.SearchTransitGatewayMulticastGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchTransitGatewayMulticastGroupsPages(input *ec2.SearchTransitGatewayMulticastGroupsInput, fn func(*ec2.SearchTransitGatewayMulticastGroupsOutput, bool) bool) error {
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
	if err := c.client.SearchTransitGatewayMulticastGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Client) SearchTransitGatewayMulticastGroupsWithContext(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput, opts ...request.Option) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.client.SearchTransitGatewayMulticastGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchTransitGatewayRoutes(input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.client.SearchTransitGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SearchTransitGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchTransitGatewayRoutesInput, opts ...request.Option) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.client.SearchTransitGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SendDiagnosticInterrupt(input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SendDiagnosticInterruptOutput), nil
	}
	out, err := c.client.SendDiagnosticInterrupt(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) SendDiagnosticInterruptWithContext(ctx context.Context, input *ec2.SendDiagnosticInterruptInput, opts ...request.Option) (*ec2.SendDiagnosticInterruptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SendDiagnosticInterruptOutput), nil
	}
	out, err := c.client.SendDiagnosticInterruptWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartInstances(input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartInstancesOutput), nil
	}
	out, err := c.client.StartInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartInstancesWithContext(ctx context.Context, input *ec2.StartInstancesInput, opts ...request.Option) (*ec2.StartInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartInstancesOutput), nil
	}
	out, err := c.client.StartInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartNetworkInsightsAccessScopeAnalysis(input *ec2.StartNetworkInsightsAccessScopeAnalysisInput) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.client.StartNetworkInsightsAccessScopeAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartNetworkInsightsAccessScopeAnalysisWithContext(ctx context.Context, input *ec2.StartNetworkInsightsAccessScopeAnalysisInput, opts ...request.Option) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAccessScopeAnalysisOutput), nil
	}
	out, err := c.client.StartNetworkInsightsAccessScopeAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartNetworkInsightsAnalysis(input *ec2.StartNetworkInsightsAnalysisInput) (*ec2.StartNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.client.StartNetworkInsightsAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartNetworkInsightsAnalysisWithContext(ctx context.Context, input *ec2.StartNetworkInsightsAnalysisInput, opts ...request.Option) (*ec2.StartNetworkInsightsAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartNetworkInsightsAnalysisOutput), nil
	}
	out, err := c.client.StartNetworkInsightsAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartVpcEndpointServicePrivateDnsVerification(input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput), nil
	}
	out, err := c.client.StartVpcEndpointServicePrivateDnsVerification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StartVpcEndpointServicePrivateDnsVerificationWithContext(ctx context.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput, opts ...request.Option) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput), nil
	}
	out, err := c.client.StartVpcEndpointServicePrivateDnsVerificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StopInstances(input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StopInstancesOutput), nil
	}
	out, err := c.client.StopInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) StopInstancesWithContext(ctx context.Context, input *ec2.StopInstancesInput, opts ...request.Option) (*ec2.StopInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.StopInstancesOutput), nil
	}
	out, err := c.client.StopInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) TerminateClientVpnConnections(input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateClientVpnConnectionsOutput), nil
	}
	out, err := c.client.TerminateClientVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) TerminateClientVpnConnectionsWithContext(ctx context.Context, input *ec2.TerminateClientVpnConnectionsInput, opts ...request.Option) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateClientVpnConnectionsOutput), nil
	}
	out, err := c.client.TerminateClientVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateInstancesOutput), nil
	}
	out, err := c.client.TerminateInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) TerminateInstancesWithContext(ctx context.Context, input *ec2.TerminateInstancesInput, opts ...request.Option) (*ec2.TerminateInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.TerminateInstancesOutput), nil
	}
	out, err := c.client.TerminateInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnassignIpv6Addresses(input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignIpv6AddressesOutput), nil
	}
	out, err := c.client.UnassignIpv6Addresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnassignIpv6AddressesWithContext(ctx context.Context, input *ec2.UnassignIpv6AddressesInput, opts ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignIpv6AddressesOutput), nil
	}
	out, err := c.client.UnassignIpv6AddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnassignPrivateIpAddresses(input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignPrivateIpAddressesOutput), nil
	}
	out, err := c.client.UnassignPrivateIpAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnassignPrivateIpAddressesWithContext(ctx context.Context, input *ec2.UnassignPrivateIpAddressesInput, opts ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnassignPrivateIpAddressesOutput), nil
	}
	out, err := c.client.UnassignPrivateIpAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnmonitorInstances(input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnmonitorInstancesOutput), nil
	}
	out, err := c.client.UnmonitorInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UnmonitorInstancesWithContext(ctx context.Context, input *ec2.UnmonitorInstancesInput, opts ...request.Option) (*ec2.UnmonitorInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UnmonitorInstancesOutput), nil
	}
	out, err := c.client.UnmonitorInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UpdateSecurityGroupRuleDescriptionsEgress(input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput), nil
	}
	out, err := c.client.UpdateSecurityGroupRuleDescriptionsEgress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UpdateSecurityGroupRuleDescriptionsEgressWithContext(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, opts ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput), nil
	}
	out, err := c.client.UpdateSecurityGroupRuleDescriptionsEgressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UpdateSecurityGroupRuleDescriptionsIngress(input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput), nil
	}
	out, err := c.client.UpdateSecurityGroupRuleDescriptionsIngress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) UpdateSecurityGroupRuleDescriptionsIngressWithContext(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, opts ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput), nil
	}
	out, err := c.client.UpdateSecurityGroupRuleDescriptionsIngressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) WithdrawByoipCidr(input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.WithdrawByoipCidrOutput), nil
	}
	out, err := c.client.WithdrawByoipCidr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Client) WithdrawByoipCidrWithContext(ctx context.Context, input *ec2.WithdrawByoipCidrInput, opts ...request.Option) (*ec2.WithdrawByoipCidrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.WithdrawByoipCidrOutput), nil
	}
	out, err := c.client.WithdrawByoipCidrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}


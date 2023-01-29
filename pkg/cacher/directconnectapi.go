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
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DirectConnect struct {
	directconnectiface.DirectConnectAPI
	cache *cache.Cache
}

func NewDirectConnect(directconnectapi directconnectiface.DirectConnectAPI) *DirectConnect {
	return &DirectConnect{
		DirectConnectAPI: directconnectapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DirectConnect) DescribeConnectionLoa(input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeConnectionLoaOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeConnectionLoa(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeConnectionLoaWithContext(ctx context.Context, input *directconnect.DescribeConnectionLoaInput, opts ...request.Option) (*directconnect.DescribeConnectionLoaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeConnectionLoaOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeConnectionLoaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeCustomerMetadata(input *directconnect.DescribeCustomerMetadataInput) (*directconnect.DescribeCustomerMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeCustomerMetadataOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeCustomerMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeCustomerMetadataWithContext(ctx context.Context, input *directconnect.DescribeCustomerMetadataInput, opts ...request.Option) (*directconnect.DescribeCustomerMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeCustomerMetadataOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeCustomerMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAssociationProposals(input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAssociationProposals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, opts ...request.Option) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAssociations(input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAssociationsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAssociationsWithContext(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput, opts ...request.Option) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAssociationsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAttachments(input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAttachmentsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewayAttachmentsWithContext(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput, opts ...request.Option) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewayAttachmentsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewayAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGateways(input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewaysOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeDirectConnectGatewaysWithContext(ctx context.Context, input *directconnect.DescribeDirectConnectGatewaysInput, opts ...request.Option) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeDirectConnectGatewaysOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeDirectConnectGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeInterconnectLoa(input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeInterconnectLoaOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeInterconnectLoa(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeInterconnectLoaWithContext(ctx context.Context, input *directconnect.DescribeInterconnectLoaInput, opts ...request.Option) (*directconnect.DescribeInterconnectLoaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeInterconnectLoaOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeInterconnectLoaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeInterconnects(input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeInterconnectsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeInterconnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeInterconnectsWithContext(ctx context.Context, input *directconnect.DescribeInterconnectsInput, opts ...request.Option) (*directconnect.DescribeInterconnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeInterconnectsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeInterconnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeLags(input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeLagsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeLags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeLagsWithContext(ctx context.Context, input *directconnect.DescribeLagsInput, opts ...request.Option) (*directconnect.DescribeLagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeLagsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeLagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeLocations(input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeLocationsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeLocationsWithContext(ctx context.Context, input *directconnect.DescribeLocationsInput, opts ...request.Option) (*directconnect.DescribeLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeLocationsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeRouterConfiguration(input *directconnect.DescribeRouterConfigurationInput) (*directconnect.DescribeRouterConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeRouterConfigurationOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeRouterConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeRouterConfigurationWithContext(ctx context.Context, input *directconnect.DescribeRouterConfigurationInput, opts ...request.Option) (*directconnect.DescribeRouterConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeRouterConfigurationOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeRouterConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeTags(input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeTagsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeTagsWithContext(ctx context.Context, input *directconnect.DescribeTagsInput, opts ...request.Option) (*directconnect.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeTagsOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeVirtualGateways(input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeVirtualGatewaysOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeVirtualGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeVirtualGatewaysWithContext(ctx context.Context, input *directconnect.DescribeVirtualGatewaysInput, opts ...request.Option) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeVirtualGatewaysOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeVirtualGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeVirtualInterfaces(input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeVirtualInterfacesOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeVirtualInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) DescribeVirtualInterfacesWithContext(ctx context.Context, input *directconnect.DescribeVirtualInterfacesInput, opts ...request.Option) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.DescribeVirtualInterfacesOutput), nil
	}
	out, err := c.DirectConnectAPI.DescribeVirtualInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) ListVirtualInterfaceTestHistory(input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.ListVirtualInterfaceTestHistoryOutput), nil
	}
	out, err := c.DirectConnectAPI.ListVirtualInterfaceTestHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectConnect) ListVirtualInterfaceTestHistoryWithContext(ctx context.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput, opts ...request.Option) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directconnect.ListVirtualInterfaceTestHistoryOutput), nil
	}
	out, err := c.DirectConnectAPI.ListVirtualInterfaceTestHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

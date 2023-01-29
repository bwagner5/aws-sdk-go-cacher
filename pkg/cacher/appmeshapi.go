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
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppMesh struct {
	appmeshiface.AppMeshAPI
	cache *cache.Cache
}

func NewAppMesh(appmeshapi appmeshiface.AppMeshAPI) *AppMesh {
	return &AppMesh{
		AppMeshAPI: appmeshapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppMesh) DescribeGatewayRoute(input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeGatewayRouteOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeGatewayRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeGatewayRouteWithContext(ctx context.Context, input *appmesh.DescribeGatewayRouteInput, opts ...request.Option) (*appmesh.DescribeGatewayRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeGatewayRouteOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeGatewayRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeMesh(input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeMeshOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeMesh(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeMeshWithContext(ctx context.Context, input *appmesh.DescribeMeshInput, opts ...request.Option) (*appmesh.DescribeMeshOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeMeshOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeMeshWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeRoute(input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeRouteOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeRouteWithContext(ctx context.Context, input *appmesh.DescribeRouteInput, opts ...request.Option) (*appmesh.DescribeRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeRouteOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualGateway(input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualGatewayOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualGatewayWithContext(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput, opts ...request.Option) (*appmesh.DescribeVirtualGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualGatewayOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualNode(input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualNodeOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualNode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualNodeWithContext(ctx context.Context, input *appmesh.DescribeVirtualNodeInput, opts ...request.Option) (*appmesh.DescribeVirtualNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualNodeOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualNodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualRouter(input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualRouterOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualRouter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualRouterWithContext(ctx context.Context, input *appmesh.DescribeVirtualRouterInput, opts ...request.Option) (*appmesh.DescribeVirtualRouterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualRouterOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualRouterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualService(input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualServiceOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualService(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) DescribeVirtualServiceWithContext(ctx context.Context, input *appmesh.DescribeVirtualServiceInput, opts ...request.Option) (*appmesh.DescribeVirtualServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.DescribeVirtualServiceOutput), nil
	}
	out, err := c.AppMeshAPI.DescribeVirtualServiceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListGatewayRoutes(input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListGatewayRoutesOutput), nil
	}
	out, err := c.AppMeshAPI.ListGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListGatewayRoutesPages(input *appmesh.ListGatewayRoutesInput, fn func(*appmesh.ListGatewayRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListGatewayRoutesOutput{}
	fnCacher := func(out *appmesh.ListGatewayRoutesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListGatewayRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListGatewayRoutesPagesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, fn func(*appmesh.ListGatewayRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListGatewayRoutesOutput{}
	fnCacher := func(out *appmesh.ListGatewayRoutesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListGatewayRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListGatewayRoutesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, opts ...request.Option) (*appmesh.ListGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListGatewayRoutesOutput), nil
	}
	out, err := c.AppMeshAPI.ListGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListMeshes(input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListMeshesOutput), nil
	}
	out, err := c.AppMeshAPI.ListMeshes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListMeshesPages(input *appmesh.ListMeshesInput, fn func(*appmesh.ListMeshesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListMeshesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListMeshesOutput{}
	fnCacher := func(out *appmesh.ListMeshesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListMeshesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListMeshesPagesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, fn func(*appmesh.ListMeshesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListMeshesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListMeshesOutput{}
	fnCacher := func(out *appmesh.ListMeshesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListMeshesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListMeshesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, opts ...request.Option) (*appmesh.ListMeshesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListMeshesOutput), nil
	}
	out, err := c.AppMeshAPI.ListMeshesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListRoutes(input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListRoutesOutput), nil
	}
	out, err := c.AppMeshAPI.ListRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListRoutesPages(input *appmesh.ListRoutesInput, fn func(*appmesh.ListRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListRoutesOutput{}
	fnCacher := func(out *appmesh.ListRoutesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListRoutesPagesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, fn func(*appmesh.ListRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListRoutesOutput{}
	fnCacher := func(out *appmesh.ListRoutesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListRoutesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, opts ...request.Option) (*appmesh.ListRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListRoutesOutput), nil
	}
	out, err := c.AppMeshAPI.ListRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListTagsForResource(input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListTagsForResourceOutput), nil
	}
	out, err := c.AppMeshAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListTagsForResourcePages(input *appmesh.ListTagsForResourceInput, fn func(*appmesh.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListTagsForResourceOutput{}
	fnCacher := func(out *appmesh.ListTagsForResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListTagsForResourcePagesWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, fn func(*appmesh.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListTagsForResourceOutput{}
	fnCacher := func(out *appmesh.ListTagsForResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListTagsForResourceWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, opts ...request.Option) (*appmesh.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListTagsForResourceOutput), nil
	}
	out, err := c.AppMeshAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualGateways(input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualGatewaysOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualGatewaysPages(input *appmesh.ListVirtualGatewaysInput, fn func(*appmesh.ListVirtualGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualGatewaysOutput{}
	fnCacher := func(out *appmesh.ListVirtualGatewaysOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualGatewaysPagesWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, fn func(*appmesh.ListVirtualGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualGatewaysOutput{}
	fnCacher := func(out *appmesh.ListVirtualGatewaysOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualGatewaysWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, opts ...request.Option) (*appmesh.ListVirtualGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualGatewaysOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualNodes(input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualNodesOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualNodesPages(input *appmesh.ListVirtualNodesInput, fn func(*appmesh.ListVirtualNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualNodesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualNodesOutput{}
	fnCacher := func(out *appmesh.ListVirtualNodesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualNodesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, fn func(*appmesh.ListVirtualNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualNodesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualNodesOutput{}
	fnCacher := func(out *appmesh.ListVirtualNodesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualNodesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, opts ...request.Option) (*appmesh.ListVirtualNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualNodesOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualRouters(input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualRoutersOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualRouters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualRoutersPages(input *appmesh.ListVirtualRoutersInput, fn func(*appmesh.ListVirtualRoutersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualRoutersOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualRoutersOutput{}
	fnCacher := func(out *appmesh.ListVirtualRoutersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualRoutersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualRoutersPagesWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, fn func(*appmesh.ListVirtualRoutersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualRoutersOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualRoutersOutput{}
	fnCacher := func(out *appmesh.ListVirtualRoutersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualRoutersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualRoutersWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, opts ...request.Option) (*appmesh.ListVirtualRoutersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualRoutersOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualRoutersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualServices(input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualServicesOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppMesh) ListVirtualServicesPages(input *appmesh.ListVirtualServicesInput, fn func(*appmesh.ListVirtualServicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualServicesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualServicesOutput{}
	fnCacher := func(out *appmesh.ListVirtualServicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualServicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualServicesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, fn func(*appmesh.ListVirtualServicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appmesh.ListVirtualServicesOutput), false)
		return nil
	}
	cachable := true
	output := &appmesh.ListVirtualServicesOutput{}
	fnCacher := func(out *appmesh.ListVirtualServicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppMeshAPI.ListVirtualServicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppMesh) ListVirtualServicesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, opts ...request.Option) (*appmesh.ListVirtualServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appmesh.ListVirtualServicesOutput), nil
	}
	out, err := c.AppMeshAPI.ListVirtualServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

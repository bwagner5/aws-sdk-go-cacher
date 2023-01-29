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
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ServiceDiscovery struct {
	servicediscoveryiface.ServiceDiscoveryAPI
	cache *cache.Cache
}

func NewServiceDiscovery(servicediscoveryapi servicediscoveryiface.ServiceDiscoveryAPI) *ServiceDiscovery {
	return &ServiceDiscovery{
		ServiceDiscoveryAPI: servicediscoveryapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ServiceDiscovery) GetInstance(input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetInstanceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetInstanceWithContext(ctx context.Context, input *servicediscovery.GetInstanceInput, opts ...request.Option) (*servicediscovery.GetInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetInstanceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetInstancesHealthStatus(input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetInstancesHealthStatusOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetInstancesHealthStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetInstancesHealthStatusPages(input *servicediscovery.GetInstancesHealthStatusInput, fn func(*servicediscovery.GetInstancesHealthStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.GetInstancesHealthStatusOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.GetInstancesHealthStatusOutput{}
	fnCacher := func(out *servicediscovery.GetInstancesHealthStatusOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.GetInstancesHealthStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) GetInstancesHealthStatusPagesWithContext(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput, fn func(*servicediscovery.GetInstancesHealthStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.GetInstancesHealthStatusOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.GetInstancesHealthStatusOutput{}
	fnCacher := func(out *servicediscovery.GetInstancesHealthStatusOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.GetInstancesHealthStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) GetInstancesHealthStatusWithContext(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput, opts ...request.Option) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetInstancesHealthStatusOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetInstancesHealthStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetNamespace(input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetNamespaceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetNamespace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetNamespaceWithContext(ctx context.Context, input *servicediscovery.GetNamespaceInput, opts ...request.Option) (*servicediscovery.GetNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetNamespaceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetNamespaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetOperation(input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetOperationOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetOperationWithContext(ctx context.Context, input *servicediscovery.GetOperationInput, opts ...request.Option) (*servicediscovery.GetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetOperationOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetService(input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetServiceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetService(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) GetServiceWithContext(ctx context.Context, input *servicediscovery.GetServiceInput, opts ...request.Option) (*servicediscovery.GetServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.GetServiceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.GetServiceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListInstances(input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListInstancesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListInstancesPages(input *servicediscovery.ListInstancesInput, fn func(*servicediscovery.ListInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListInstancesOutput{}
	fnCacher := func(out *servicediscovery.ListInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListInstancesPagesWithContext(ctx context.Context, input *servicediscovery.ListInstancesInput, fn func(*servicediscovery.ListInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListInstancesOutput{}
	fnCacher := func(out *servicediscovery.ListInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListInstancesWithContext(ctx context.Context, input *servicediscovery.ListInstancesInput, opts ...request.Option) (*servicediscovery.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListInstancesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListNamespaces(input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListNamespacesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListNamespaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListNamespacesPages(input *servicediscovery.ListNamespacesInput, fn func(*servicediscovery.ListNamespacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListNamespacesOutput{}
	fnCacher := func(out *servicediscovery.ListNamespacesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListNamespacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListNamespacesPagesWithContext(ctx context.Context, input *servicediscovery.ListNamespacesInput, fn func(*servicediscovery.ListNamespacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListNamespacesOutput{}
	fnCacher := func(out *servicediscovery.ListNamespacesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListNamespacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListNamespacesWithContext(ctx context.Context, input *servicediscovery.ListNamespacesInput, opts ...request.Option) (*servicediscovery.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListNamespacesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListNamespacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListOperations(input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListOperationsOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListOperationsPages(input *servicediscovery.ListOperationsInput, fn func(*servicediscovery.ListOperationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListOperationsOutput{}
	fnCacher := func(out *servicediscovery.ListOperationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListOperationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListOperationsPagesWithContext(ctx context.Context, input *servicediscovery.ListOperationsInput, fn func(*servicediscovery.ListOperationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListOperationsOutput{}
	fnCacher := func(out *servicediscovery.ListOperationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListOperationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListOperationsWithContext(ctx context.Context, input *servicediscovery.ListOperationsInput, opts ...request.Option) (*servicediscovery.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListOperationsOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListServices(input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListServicesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListServicesPages(input *servicediscovery.ListServicesInput, fn func(*servicediscovery.ListServicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListServicesOutput{}
	fnCacher := func(out *servicediscovery.ListServicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListServicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListServicesPagesWithContext(ctx context.Context, input *servicediscovery.ListServicesInput, fn func(*servicediscovery.ListServicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicediscovery.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &servicediscovery.ListServicesOutput{}
	fnCacher := func(out *servicediscovery.ListServicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ServiceDiscoveryAPI.ListServicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceDiscovery) ListServicesWithContext(ctx context.Context, input *servicediscovery.ListServicesInput, opts ...request.Option) (*servicediscovery.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListServicesOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListTagsForResource(input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListTagsForResourceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceDiscovery) ListTagsForResourceWithContext(ctx context.Context, input *servicediscovery.ListTagsForResourceInput, opts ...request.Option) (*servicediscovery.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicediscovery.ListTagsForResourceOutput), nil
	}
	out, err := c.ServiceDiscoveryAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

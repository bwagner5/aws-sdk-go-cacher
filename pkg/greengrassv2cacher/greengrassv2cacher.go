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
package greengrassv2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/greengrassv2"
	"github.com/aws/aws-sdk-go/service/greengrassv2/greengrassv2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GreengrassV2 struct {
	greengrassv2iface.GreengrassV2API
	cache *cache.Cache
}

func New(greengrassv2api greengrassv2iface.GreengrassV2API) *GreengrassV2 {
	return &GreengrassV2{
		GreengrassV2API: greengrassv2api,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GreengrassV2) BatchAssociateClientDeviceWithCoreDevice(input *greengrassv2.BatchAssociateClientDeviceWithCoreDeviceInput) (*greengrassv2.BatchAssociateClientDeviceWithCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.BatchAssociateClientDeviceWithCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.BatchAssociateClientDeviceWithCoreDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) BatchAssociateClientDeviceWithCoreDeviceWithContext(ctx context.Context, input *greengrassv2.BatchAssociateClientDeviceWithCoreDeviceInput, opts ...request.Option) (*greengrassv2.BatchAssociateClientDeviceWithCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.BatchAssociateClientDeviceWithCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.BatchAssociateClientDeviceWithCoreDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) BatchDisassociateClientDeviceFromCoreDevice(input *greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceInput) (*greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.BatchDisassociateClientDeviceFromCoreDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) BatchDisassociateClientDeviceFromCoreDeviceWithContext(ctx context.Context, input *greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceInput, opts ...request.Option) (*greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.BatchDisassociateClientDeviceFromCoreDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) DescribeComponent(input *greengrassv2.DescribeComponentInput) (*greengrassv2.DescribeComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.DescribeComponentOutput), nil
	}
	out, err := c.GreengrassV2API.DescribeComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) DescribeComponentWithContext(ctx context.Context, input *greengrassv2.DescribeComponentInput, opts ...request.Option) (*greengrassv2.DescribeComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.DescribeComponentOutput), nil
	}
	out, err := c.GreengrassV2API.DescribeComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetComponent(input *greengrassv2.GetComponentInput) (*greengrassv2.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetComponentOutput), nil
	}
	out, err := c.GreengrassV2API.GetComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetComponentVersionArtifact(input *greengrassv2.GetComponentVersionArtifactInput) (*greengrassv2.GetComponentVersionArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetComponentVersionArtifactOutput), nil
	}
	out, err := c.GreengrassV2API.GetComponentVersionArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetComponentVersionArtifactWithContext(ctx context.Context, input *greengrassv2.GetComponentVersionArtifactInput, opts ...request.Option) (*greengrassv2.GetComponentVersionArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetComponentVersionArtifactOutput), nil
	}
	out, err := c.GreengrassV2API.GetComponentVersionArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetComponentWithContext(ctx context.Context, input *greengrassv2.GetComponentInput, opts ...request.Option) (*greengrassv2.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetComponentOutput), nil
	}
	out, err := c.GreengrassV2API.GetComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetConnectivityInfo(input *greengrassv2.GetConnectivityInfoInput) (*greengrassv2.GetConnectivityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetConnectivityInfoOutput), nil
	}
	out, err := c.GreengrassV2API.GetConnectivityInfo(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetConnectivityInfoWithContext(ctx context.Context, input *greengrassv2.GetConnectivityInfoInput, opts ...request.Option) (*greengrassv2.GetConnectivityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetConnectivityInfoOutput), nil
	}
	out, err := c.GreengrassV2API.GetConnectivityInfoWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetCoreDevice(input *greengrassv2.GetCoreDeviceInput) (*greengrassv2.GetCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.GetCoreDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetCoreDeviceWithContext(ctx context.Context, input *greengrassv2.GetCoreDeviceInput, opts ...request.Option) (*greengrassv2.GetCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.GetCoreDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetDeployment(input *greengrassv2.GetDeploymentInput) (*greengrassv2.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetDeploymentOutput), nil
	}
	out, err := c.GreengrassV2API.GetDeployment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetDeploymentWithContext(ctx context.Context, input *greengrassv2.GetDeploymentInput, opts ...request.Option) (*greengrassv2.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetDeploymentOutput), nil
	}
	out, err := c.GreengrassV2API.GetDeploymentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetServiceRoleForAccount(input *greengrassv2.GetServiceRoleForAccountInput) (*greengrassv2.GetServiceRoleForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetServiceRoleForAccountOutput), nil
	}
	out, err := c.GreengrassV2API.GetServiceRoleForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) GetServiceRoleForAccountWithContext(ctx context.Context, input *greengrassv2.GetServiceRoleForAccountInput, opts ...request.Option) (*greengrassv2.GetServiceRoleForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.GetServiceRoleForAccountOutput), nil
	}
	out, err := c.GreengrassV2API.GetServiceRoleForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListClientDevicesAssociatedWithCoreDevice(input *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput) (*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.ListClientDevicesAssociatedWithCoreDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListClientDevicesAssociatedWithCoreDevicePages(input *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput, fn func(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput{}
	fnCacher := func(out *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListClientDevicesAssociatedWithCoreDevicePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListClientDevicesAssociatedWithCoreDevicePagesWithContext(ctx context.Context, input *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput, fn func(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput{}
	fnCacher := func(out *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListClientDevicesAssociatedWithCoreDevicePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListClientDevicesAssociatedWithCoreDeviceWithContext(ctx context.Context, input *greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput, opts ...request.Option) (*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput), nil
	}
	out, err := c.GreengrassV2API.ListClientDevicesAssociatedWithCoreDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListComponentVersions(input *greengrassv2.ListComponentVersionsInput) (*greengrassv2.ListComponentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListComponentVersionsOutput), nil
	}
	out, err := c.GreengrassV2API.ListComponentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListComponentVersionsPages(input *greengrassv2.ListComponentVersionsInput, fn func(*greengrassv2.ListComponentVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListComponentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListComponentVersionsOutput{}
	fnCacher := func(out *greengrassv2.ListComponentVersionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListComponentVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListComponentVersionsPagesWithContext(ctx context.Context, input *greengrassv2.ListComponentVersionsInput, fn func(*greengrassv2.ListComponentVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListComponentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListComponentVersionsOutput{}
	fnCacher := func(out *greengrassv2.ListComponentVersionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListComponentVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListComponentVersionsWithContext(ctx context.Context, input *greengrassv2.ListComponentVersionsInput, opts ...request.Option) (*greengrassv2.ListComponentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListComponentVersionsOutput), nil
	}
	out, err := c.GreengrassV2API.ListComponentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListComponents(input *greengrassv2.ListComponentsInput) (*greengrassv2.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListComponentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListComponentsPages(input *greengrassv2.ListComponentsInput, fn func(*greengrassv2.ListComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListComponentsOutput{}
	fnCacher := func(out *greengrassv2.ListComponentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListComponentsPagesWithContext(ctx context.Context, input *greengrassv2.ListComponentsInput, fn func(*greengrassv2.ListComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListComponentsOutput{}
	fnCacher := func(out *greengrassv2.ListComponentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListComponentsWithContext(ctx context.Context, input *greengrassv2.ListComponentsInput, opts ...request.Option) (*greengrassv2.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListComponentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListCoreDevices(input *greengrassv2.ListCoreDevicesInput) (*greengrassv2.ListCoreDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListCoreDevicesOutput), nil
	}
	out, err := c.GreengrassV2API.ListCoreDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListCoreDevicesPages(input *greengrassv2.ListCoreDevicesInput, fn func(*greengrassv2.ListCoreDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListCoreDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListCoreDevicesOutput{}
	fnCacher := func(out *greengrassv2.ListCoreDevicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListCoreDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListCoreDevicesPagesWithContext(ctx context.Context, input *greengrassv2.ListCoreDevicesInput, fn func(*greengrassv2.ListCoreDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListCoreDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListCoreDevicesOutput{}
	fnCacher := func(out *greengrassv2.ListCoreDevicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListCoreDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListCoreDevicesWithContext(ctx context.Context, input *greengrassv2.ListCoreDevicesInput, opts ...request.Option) (*greengrassv2.ListCoreDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListCoreDevicesOutput), nil
	}
	out, err := c.GreengrassV2API.ListCoreDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListDeployments(input *greengrassv2.ListDeploymentsInput) (*greengrassv2.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListDeploymentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListDeploymentsPages(input *greengrassv2.ListDeploymentsInput, fn func(*greengrassv2.ListDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListDeploymentsOutput{}
	fnCacher := func(out *greengrassv2.ListDeploymentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListDeploymentsPagesWithContext(ctx context.Context, input *greengrassv2.ListDeploymentsInput, fn func(*greengrassv2.ListDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListDeploymentsOutput{}
	fnCacher := func(out *greengrassv2.ListDeploymentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListDeploymentsWithContext(ctx context.Context, input *greengrassv2.ListDeploymentsInput, opts ...request.Option) (*greengrassv2.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListDeploymentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListEffectiveDeployments(input *greengrassv2.ListEffectiveDeploymentsInput) (*greengrassv2.ListEffectiveDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListEffectiveDeploymentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListEffectiveDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListEffectiveDeploymentsPages(input *greengrassv2.ListEffectiveDeploymentsInput, fn func(*greengrassv2.ListEffectiveDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListEffectiveDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListEffectiveDeploymentsOutput{}
	fnCacher := func(out *greengrassv2.ListEffectiveDeploymentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListEffectiveDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListEffectiveDeploymentsPagesWithContext(ctx context.Context, input *greengrassv2.ListEffectiveDeploymentsInput, fn func(*greengrassv2.ListEffectiveDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListEffectiveDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListEffectiveDeploymentsOutput{}
	fnCacher := func(out *greengrassv2.ListEffectiveDeploymentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListEffectiveDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListEffectiveDeploymentsWithContext(ctx context.Context, input *greengrassv2.ListEffectiveDeploymentsInput, opts ...request.Option) (*greengrassv2.ListEffectiveDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListEffectiveDeploymentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListEffectiveDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListInstalledComponents(input *greengrassv2.ListInstalledComponentsInput) (*greengrassv2.ListInstalledComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListInstalledComponentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListInstalledComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListInstalledComponentsPages(input *greengrassv2.ListInstalledComponentsInput, fn func(*greengrassv2.ListInstalledComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListInstalledComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListInstalledComponentsOutput{}
	fnCacher := func(out *greengrassv2.ListInstalledComponentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListInstalledComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListInstalledComponentsPagesWithContext(ctx context.Context, input *greengrassv2.ListInstalledComponentsInput, fn func(*greengrassv2.ListInstalledComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*greengrassv2.ListInstalledComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &greengrassv2.ListInstalledComponentsOutput{}
	fnCacher := func(out *greengrassv2.ListInstalledComponentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GreengrassV2API.ListInstalledComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GreengrassV2) ListInstalledComponentsWithContext(ctx context.Context, input *greengrassv2.ListInstalledComponentsInput, opts ...request.Option) (*greengrassv2.ListInstalledComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListInstalledComponentsOutput), nil
	}
	out, err := c.GreengrassV2API.ListInstalledComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListTagsForResource(input *greengrassv2.ListTagsForResourceInput) (*greengrassv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListTagsForResourceOutput), nil
	}
	out, err := c.GreengrassV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GreengrassV2) ListTagsForResourceWithContext(ctx context.Context, input *greengrassv2.ListTagsForResourceInput, opts ...request.Option) (*greengrassv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrassv2.ListTagsForResourceOutput), nil
	}
	out, err := c.GreengrassV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

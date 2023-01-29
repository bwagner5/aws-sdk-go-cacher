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
package panoramacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/panorama"
	"github.com/aws/aws-sdk-go/service/panorama/panoramaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Panorama struct {
	panoramaiface.PanoramaAPI
	cache *cache.Cache
}

func New(panoramaapi panoramaiface.PanoramaAPI) *Panorama {
	return &Panorama{
		PanoramaAPI: panoramaapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Panorama) DescribeApplicationInstance(input *panorama.DescribeApplicationInstanceInput) (*panorama.DescribeApplicationInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeApplicationInstanceOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeApplicationInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeApplicationInstanceDetails(input *panorama.DescribeApplicationInstanceDetailsInput) (*panorama.DescribeApplicationInstanceDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeApplicationInstanceDetailsOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeApplicationInstanceDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeApplicationInstanceDetailsWithContext(ctx context.Context, input *panorama.DescribeApplicationInstanceDetailsInput, opts ...request.Option) (*panorama.DescribeApplicationInstanceDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeApplicationInstanceDetailsOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeApplicationInstanceDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeApplicationInstanceWithContext(ctx context.Context, input *panorama.DescribeApplicationInstanceInput, opts ...request.Option) (*panorama.DescribeApplicationInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeApplicationInstanceOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeApplicationInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeDevice(input *panorama.DescribeDeviceInput) (*panorama.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeDeviceOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeDeviceJob(input *panorama.DescribeDeviceJobInput) (*panorama.DescribeDeviceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeDeviceJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeDeviceJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeDeviceJobWithContext(ctx context.Context, input *panorama.DescribeDeviceJobInput, opts ...request.Option) (*panorama.DescribeDeviceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeDeviceJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeDeviceJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeDeviceWithContext(ctx context.Context, input *panorama.DescribeDeviceInput, opts ...request.Option) (*panorama.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeDeviceOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeNode(input *panorama.DescribeNodeInput) (*panorama.DescribeNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeNodeOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeNode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeNodeFromTemplateJob(input *panorama.DescribeNodeFromTemplateJobInput) (*panorama.DescribeNodeFromTemplateJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeNodeFromTemplateJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeNodeFromTemplateJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeNodeFromTemplateJobWithContext(ctx context.Context, input *panorama.DescribeNodeFromTemplateJobInput, opts ...request.Option) (*panorama.DescribeNodeFromTemplateJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeNodeFromTemplateJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeNodeFromTemplateJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribeNodeWithContext(ctx context.Context, input *panorama.DescribeNodeInput, opts ...request.Option) (*panorama.DescribeNodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribeNodeOutput), nil
	}
	out, err := c.PanoramaAPI.DescribeNodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackage(input *panorama.DescribePackageInput) (*panorama.DescribePackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackageImportJob(input *panorama.DescribePackageImportJobInput) (*panorama.DescribePackageImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageImportJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackageImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackageImportJobWithContext(ctx context.Context, input *panorama.DescribePackageImportJobInput, opts ...request.Option) (*panorama.DescribePackageImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageImportJobOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackageImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackageVersion(input *panorama.DescribePackageVersionInput) (*panorama.DescribePackageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageVersionOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackageVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackageVersionWithContext(ctx context.Context, input *panorama.DescribePackageVersionInput, opts ...request.Option) (*panorama.DescribePackageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageVersionOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackageVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) DescribePackageWithContext(ctx context.Context, input *panorama.DescribePackageInput, opts ...request.Option) (*panorama.DescribePackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.DescribePackageOutput), nil
	}
	out, err := c.PanoramaAPI.DescribePackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstanceDependencies(input *panorama.ListApplicationInstanceDependenciesInput) (*panorama.ListApplicationInstanceDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstanceDependenciesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstanceDependencies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstanceDependenciesPages(input *panorama.ListApplicationInstanceDependenciesInput, fn func(*panorama.ListApplicationInstanceDependenciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstanceDependenciesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstanceDependenciesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstanceDependenciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstanceDependenciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstanceDependenciesPagesWithContext(ctx context.Context, input *panorama.ListApplicationInstanceDependenciesInput, fn func(*panorama.ListApplicationInstanceDependenciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstanceDependenciesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstanceDependenciesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstanceDependenciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstanceDependenciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstanceDependenciesWithContext(ctx context.Context, input *panorama.ListApplicationInstanceDependenciesInput, opts ...request.Option) (*panorama.ListApplicationInstanceDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstanceDependenciesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstanceDependenciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstanceNodeInstances(input *panorama.ListApplicationInstanceNodeInstancesInput) (*panorama.ListApplicationInstanceNodeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstanceNodeInstancesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstanceNodeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstanceNodeInstancesPages(input *panorama.ListApplicationInstanceNodeInstancesInput, fn func(*panorama.ListApplicationInstanceNodeInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstanceNodeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstanceNodeInstancesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstanceNodeInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstanceNodeInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstanceNodeInstancesPagesWithContext(ctx context.Context, input *panorama.ListApplicationInstanceNodeInstancesInput, fn func(*panorama.ListApplicationInstanceNodeInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstanceNodeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstanceNodeInstancesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstanceNodeInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstanceNodeInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstanceNodeInstancesWithContext(ctx context.Context, input *panorama.ListApplicationInstanceNodeInstancesInput, opts ...request.Option) (*panorama.ListApplicationInstanceNodeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstanceNodeInstancesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstanceNodeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstances(input *panorama.ListApplicationInstancesInput) (*panorama.ListApplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstancesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListApplicationInstancesPages(input *panorama.ListApplicationInstancesInput, fn func(*panorama.ListApplicationInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstancesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstancesPagesWithContext(ctx context.Context, input *panorama.ListApplicationInstancesInput, fn func(*panorama.ListApplicationInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListApplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListApplicationInstancesOutput{}
	fnCacher := func(out *panorama.ListApplicationInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListApplicationInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListApplicationInstancesWithContext(ctx context.Context, input *panorama.ListApplicationInstancesInput, opts ...request.Option) (*panorama.ListApplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListApplicationInstancesOutput), nil
	}
	out, err := c.PanoramaAPI.ListApplicationInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListDevices(input *panorama.ListDevicesInput) (*panorama.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListDevicesOutput), nil
	}
	out, err := c.PanoramaAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListDevicesJobs(input *panorama.ListDevicesJobsInput) (*panorama.ListDevicesJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListDevicesJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListDevicesJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListDevicesJobsPages(input *panorama.ListDevicesJobsInput, fn func(*panorama.ListDevicesJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListDevicesJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListDevicesJobsOutput{}
	fnCacher := func(out *panorama.ListDevicesJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListDevicesJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListDevicesJobsPagesWithContext(ctx context.Context, input *panorama.ListDevicesJobsInput, fn func(*panorama.ListDevicesJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListDevicesJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListDevicesJobsOutput{}
	fnCacher := func(out *panorama.ListDevicesJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListDevicesJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListDevicesJobsWithContext(ctx context.Context, input *panorama.ListDevicesJobsInput, opts ...request.Option) (*panorama.ListDevicesJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListDevicesJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListDevicesJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListDevicesPages(input *panorama.ListDevicesInput, fn func(*panorama.ListDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListDevicesOutput{}
	fnCacher := func(out *panorama.ListDevicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListDevicesPagesWithContext(ctx context.Context, input *panorama.ListDevicesInput, fn func(*panorama.ListDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListDevicesOutput{}
	fnCacher := func(out *panorama.ListDevicesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListDevicesWithContext(ctx context.Context, input *panorama.ListDevicesInput, opts ...request.Option) (*panorama.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListDevicesOutput), nil
	}
	out, err := c.PanoramaAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListNodeFromTemplateJobs(input *panorama.ListNodeFromTemplateJobsInput) (*panorama.ListNodeFromTemplateJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListNodeFromTemplateJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListNodeFromTemplateJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListNodeFromTemplateJobsPages(input *panorama.ListNodeFromTemplateJobsInput, fn func(*panorama.ListNodeFromTemplateJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListNodeFromTemplateJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListNodeFromTemplateJobsOutput{}
	fnCacher := func(out *panorama.ListNodeFromTemplateJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListNodeFromTemplateJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListNodeFromTemplateJobsPagesWithContext(ctx context.Context, input *panorama.ListNodeFromTemplateJobsInput, fn func(*panorama.ListNodeFromTemplateJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListNodeFromTemplateJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListNodeFromTemplateJobsOutput{}
	fnCacher := func(out *panorama.ListNodeFromTemplateJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListNodeFromTemplateJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListNodeFromTemplateJobsWithContext(ctx context.Context, input *panorama.ListNodeFromTemplateJobsInput, opts ...request.Option) (*panorama.ListNodeFromTemplateJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListNodeFromTemplateJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListNodeFromTemplateJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListNodes(input *panorama.ListNodesInput) (*panorama.ListNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListNodesOutput), nil
	}
	out, err := c.PanoramaAPI.ListNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListNodesPages(input *panorama.ListNodesInput, fn func(*panorama.ListNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListNodesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListNodesOutput{}
	fnCacher := func(out *panorama.ListNodesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListNodesPagesWithContext(ctx context.Context, input *panorama.ListNodesInput, fn func(*panorama.ListNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListNodesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListNodesOutput{}
	fnCacher := func(out *panorama.ListNodesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListNodesWithContext(ctx context.Context, input *panorama.ListNodesInput, opts ...request.Option) (*panorama.ListNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListNodesOutput), nil
	}
	out, err := c.PanoramaAPI.ListNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListPackageImportJobs(input *panorama.ListPackageImportJobsInput) (*panorama.ListPackageImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListPackageImportJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListPackageImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListPackageImportJobsPages(input *panorama.ListPackageImportJobsInput, fn func(*panorama.ListPackageImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListPackageImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListPackageImportJobsOutput{}
	fnCacher := func(out *panorama.ListPackageImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListPackageImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListPackageImportJobsPagesWithContext(ctx context.Context, input *panorama.ListPackageImportJobsInput, fn func(*panorama.ListPackageImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListPackageImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListPackageImportJobsOutput{}
	fnCacher := func(out *panorama.ListPackageImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListPackageImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListPackageImportJobsWithContext(ctx context.Context, input *panorama.ListPackageImportJobsInput, opts ...request.Option) (*panorama.ListPackageImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListPackageImportJobsOutput), nil
	}
	out, err := c.PanoramaAPI.ListPackageImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListPackages(input *panorama.ListPackagesInput) (*panorama.ListPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListPackagesOutput), nil
	}
	out, err := c.PanoramaAPI.ListPackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListPackagesPages(input *panorama.ListPackagesInput, fn func(*panorama.ListPackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListPackagesOutput{}
	fnCacher := func(out *panorama.ListPackagesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListPackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListPackagesPagesWithContext(ctx context.Context, input *panorama.ListPackagesInput, fn func(*panorama.ListPackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*panorama.ListPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &panorama.ListPackagesOutput{}
	fnCacher := func(out *panorama.ListPackagesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.PanoramaAPI.ListPackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Panorama) ListPackagesWithContext(ctx context.Context, input *panorama.ListPackagesInput, opts ...request.Option) (*panorama.ListPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListPackagesOutput), nil
	}
	out, err := c.PanoramaAPI.ListPackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListTagsForResource(input *panorama.ListTagsForResourceInput) (*panorama.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListTagsForResourceOutput), nil
	}
	out, err := c.PanoramaAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Panorama) ListTagsForResourceWithContext(ctx context.Context, input *panorama.ListTagsForResourceInput, opts ...request.Option) (*panorama.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*panorama.ListTagsForResourceOutput), nil
	}
	out, err := c.PanoramaAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

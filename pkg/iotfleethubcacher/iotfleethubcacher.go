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
package iotfleethubcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotfleethub"
	"github.com/aws/aws-sdk-go/service/iotfleethub/iotfleethubiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTFleetHub struct {
	iotfleethubiface.IoTFleetHubAPI
	cache *cache.Cache
}

func New(iotfleethubapi iotfleethubiface.IoTFleetHubAPI) *IoTFleetHub {
	return &IoTFleetHub{
		IoTFleetHubAPI: iotfleethubapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTFleetHub) DescribeApplication(input *iotfleethub.DescribeApplicationInput) (*iotfleethub.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.DescribeApplicationOutput), nil
	}
	out, err := c.IoTFleetHubAPI.DescribeApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetHub) DescribeApplicationWithContext(ctx context.Context, input *iotfleethub.DescribeApplicationInput, opts ...request.Option) (*iotfleethub.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.DescribeApplicationOutput), nil
	}
	out, err := c.IoTFleetHubAPI.DescribeApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetHub) ListApplications(input *iotfleethub.ListApplicationsInput) (*iotfleethub.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.ListApplicationsOutput), nil
	}
	out, err := c.IoTFleetHubAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetHub) ListApplicationsPages(input *iotfleethub.ListApplicationsInput, fn func(*iotfleethub.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleethub.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleethub.ListApplicationsOutput{}
	fnCacher := func(out *iotfleethub.ListApplicationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTFleetHubAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetHub) ListApplicationsPagesWithContext(ctx context.Context, input *iotfleethub.ListApplicationsInput, fn func(*iotfleethub.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleethub.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleethub.ListApplicationsOutput{}
	fnCacher := func(out *iotfleethub.ListApplicationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTFleetHubAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetHub) ListApplicationsWithContext(ctx context.Context, input *iotfleethub.ListApplicationsInput, opts ...request.Option) (*iotfleethub.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.ListApplicationsOutput), nil
	}
	out, err := c.IoTFleetHubAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetHub) ListTagsForResource(input *iotfleethub.ListTagsForResourceInput) (*iotfleethub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTFleetHubAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetHub) ListTagsForResourceWithContext(ctx context.Context, input *iotfleethub.ListTagsForResourceInput, opts ...request.Option) (*iotfleethub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleethub.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTFleetHubAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

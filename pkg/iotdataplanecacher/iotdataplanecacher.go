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
package iotdataplanecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTDataPlane struct {
	iotdataplaneiface.IoTDataPlaneAPI
	cache *cache.Cache
}

func New(iotdataplaneapi iotdataplaneiface.IoTDataPlaneAPI) *IoTDataPlane {
	return &IoTDataPlane{
		IoTDataPlaneAPI: iotdataplaneapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTDataPlane) GetRetainedMessage(input *iotdataplane.GetRetainedMessageInput) (*iotdataplane.GetRetainedMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.GetRetainedMessageOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.GetRetainedMessage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) GetRetainedMessageWithContext(ctx context.Context, input *iotdataplane.GetRetainedMessageInput, opts ...request.Option) (*iotdataplane.GetRetainedMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.GetRetainedMessageOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.GetRetainedMessageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) GetThingShadow(input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.GetThingShadowOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.GetThingShadow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) GetThingShadowWithContext(ctx context.Context, input *iotdataplane.GetThingShadowInput, opts ...request.Option) (*iotdataplane.GetThingShadowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.GetThingShadowOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.GetThingShadowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) ListNamedShadowsForThing(input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.ListNamedShadowsForThingOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.ListNamedShadowsForThing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) ListNamedShadowsForThingWithContext(ctx context.Context, input *iotdataplane.ListNamedShadowsForThingInput, opts ...request.Option) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.ListNamedShadowsForThingOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.ListNamedShadowsForThingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) ListRetainedMessages(input *iotdataplane.ListRetainedMessagesInput) (*iotdataplane.ListRetainedMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.ListRetainedMessagesOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.ListRetainedMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDataPlane) ListRetainedMessagesPages(input *iotdataplane.ListRetainedMessagesInput, fn func(*iotdataplane.ListRetainedMessagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdataplane.ListRetainedMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotdataplane.ListRetainedMessagesOutput{}
	fnCacher := func(out *iotdataplane.ListRetainedMessagesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTDataPlaneAPI.ListRetainedMessagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDataPlane) ListRetainedMessagesPagesWithContext(ctx context.Context, input *iotdataplane.ListRetainedMessagesInput, fn func(*iotdataplane.ListRetainedMessagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdataplane.ListRetainedMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotdataplane.ListRetainedMessagesOutput{}
	fnCacher := func(out *iotdataplane.ListRetainedMessagesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTDataPlaneAPI.ListRetainedMessagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDataPlane) ListRetainedMessagesWithContext(ctx context.Context, input *iotdataplane.ListRetainedMessagesInput, opts ...request.Option) (*iotdataplane.ListRetainedMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdataplane.ListRetainedMessagesOutput), nil
	}
	out, err := c.IoTDataPlaneAPI.ListRetainedMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

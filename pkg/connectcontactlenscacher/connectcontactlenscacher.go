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
package connectcontactlenscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/connectcontactlens"
	"github.com/aws/aws-sdk-go/service/connectcontactlens/connectcontactlensiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConnectContactLens struct {
	connectcontactlensiface.ConnectContactLensAPI
	cache *cache.Cache
}

func New(connectcontactlensapi connectcontactlensiface.ConnectContactLensAPI) *ConnectContactLens {
	return &ConnectContactLens{
		ConnectContactLensAPI: connectcontactlensapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegments(input *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput), nil
	}
	out, err := c.ConnectContactLensAPI.ListRealtimeContactAnalysisSegments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsPages(input *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, fn func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput{}
	fnCacher := func(out *connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ConnectContactLensAPI.ListRealtimeContactAnalysisSegmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsPagesWithContext(ctx context.Context, input *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, fn func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput{}
	fnCacher := func(out *connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.ConnectContactLensAPI.ListRealtimeContactAnalysisSegmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsWithContext(ctx context.Context, input *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, opts ...request.Option) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput), nil
	}
	out, err := c.ConnectContactLensAPI.ListRealtimeContactAnalysisSegmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package route53recoveryclustercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53recoverycluster"
	"github.com/aws/aws-sdk-go/service/route53recoverycluster/route53recoveryclusteriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53RecoveryCluster struct {
	route53recoveryclusteriface.Route53RecoveryClusterAPI
	cache *cache.Cache
}

func New(route53recoveryclusterapi route53recoveryclusteriface.Route53RecoveryClusterAPI) *Route53RecoveryCluster {
	return &Route53RecoveryCluster{
		Route53RecoveryClusterAPI: route53recoveryclusterapi,
		cache:                     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53RecoveryCluster) GetRoutingControlState(input *route53recoverycluster.GetRoutingControlStateInput) (*route53recoverycluster.GetRoutingControlStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycluster.GetRoutingControlStateOutput), nil
	}
	out, err := c.Route53RecoveryClusterAPI.GetRoutingControlState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryCluster) GetRoutingControlStateWithContext(ctx context.Context, input *route53recoverycluster.GetRoutingControlStateInput, opts ...request.Option) (*route53recoverycluster.GetRoutingControlStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycluster.GetRoutingControlStateOutput), nil
	}
	out, err := c.Route53RecoveryClusterAPI.GetRoutingControlStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryCluster) ListRoutingControls(input *route53recoverycluster.ListRoutingControlsInput) (*route53recoverycluster.ListRoutingControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycluster.ListRoutingControlsOutput), nil
	}
	out, err := c.Route53RecoveryClusterAPI.ListRoutingControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53RecoveryCluster) ListRoutingControlsPages(input *route53recoverycluster.ListRoutingControlsInput, fn func(*route53recoverycluster.ListRoutingControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycluster.ListRoutingControlsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycluster.ListRoutingControlsOutput{}
	fnCacher := func(out *route53recoverycluster.ListRoutingControlsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.Route53RecoveryClusterAPI.ListRoutingControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryCluster) ListRoutingControlsPagesWithContext(ctx context.Context, input *route53recoverycluster.ListRoutingControlsInput, fn func(*route53recoverycluster.ListRoutingControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53recoverycluster.ListRoutingControlsOutput), false)
		return nil
	}
	cachable := true
	output := &route53recoverycluster.ListRoutingControlsOutput{}
	fnCacher := func(out *route53recoverycluster.ListRoutingControlsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.Route53RecoveryClusterAPI.ListRoutingControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53RecoveryCluster) ListRoutingControlsWithContext(ctx context.Context, input *route53recoverycluster.ListRoutingControlsInput, opts ...request.Option) (*route53recoverycluster.ListRoutingControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53recoverycluster.ListRoutingControlsOutput), nil
	}
	out, err := c.Route53RecoveryClusterAPI.ListRoutingControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

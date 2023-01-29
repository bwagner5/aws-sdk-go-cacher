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
	"github.com/aws/aws-sdk-go/service/iotroborunner"
	"github.com/aws/aws-sdk-go/service/iotroborunner/iotroborunneriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTRoboRunner struct {
	iotroborunneriface.IoTRoboRunnerAPI
	cache *cache.Cache
}

func NewIoTRoboRunner(iotroborunnerapi iotroborunneriface.IoTRoboRunnerAPI) *IoTRoboRunner {
	return &IoTRoboRunner{
		IoTRoboRunnerAPI: iotroborunnerapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTRoboRunner) GetDestination(input *iotroborunner.GetDestinationInput) (*iotroborunner.GetDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetDestinationOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetDestinationWithContext(ctx context.Context, input *iotroborunner.GetDestinationInput, opts ...request.Option) (*iotroborunner.GetDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetDestinationOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetSite(input *iotroborunner.GetSiteInput) (*iotroborunner.GetSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetSiteOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetSite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetSiteWithContext(ctx context.Context, input *iotroborunner.GetSiteInput, opts ...request.Option) (*iotroborunner.GetSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetSiteOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetSiteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetWorker(input *iotroborunner.GetWorkerInput) (*iotroborunner.GetWorkerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetWorkerOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetWorker(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetWorkerFleet(input *iotroborunner.GetWorkerFleetInput) (*iotroborunner.GetWorkerFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetWorkerFleetOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetWorkerFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetWorkerFleetWithContext(ctx context.Context, input *iotroborunner.GetWorkerFleetInput, opts ...request.Option) (*iotroborunner.GetWorkerFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetWorkerFleetOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetWorkerFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) GetWorkerWithContext(ctx context.Context, input *iotroborunner.GetWorkerInput, opts ...request.Option) (*iotroborunner.GetWorkerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.GetWorkerOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.GetWorkerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListDestinations(input *iotroborunner.ListDestinationsInput) (*iotroborunner.ListDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListDestinationsOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListDestinationsPages(input *iotroborunner.ListDestinationsInput, fn func(*iotroborunner.ListDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListDestinationsOutput{}
	fnCacher := func(out *iotroborunner.ListDestinationsOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListDestinationsPagesWithContext(ctx context.Context, input *iotroborunner.ListDestinationsInput, fn func(*iotroborunner.ListDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListDestinationsOutput{}
	fnCacher := func(out *iotroborunner.ListDestinationsOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListDestinationsWithContext(ctx context.Context, input *iotroborunner.ListDestinationsInput, opts ...request.Option) (*iotroborunner.ListDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListDestinationsOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListSites(input *iotroborunner.ListSitesInput) (*iotroborunner.ListSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListSitesOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListSites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListSitesPages(input *iotroborunner.ListSitesInput, fn func(*iotroborunner.ListSitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListSitesOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListSitesOutput{}
	fnCacher := func(out *iotroborunner.ListSitesOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListSitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListSitesPagesWithContext(ctx context.Context, input *iotroborunner.ListSitesInput, fn func(*iotroborunner.ListSitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListSitesOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListSitesOutput{}
	fnCacher := func(out *iotroborunner.ListSitesOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListSitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListSitesWithContext(ctx context.Context, input *iotroborunner.ListSitesInput, opts ...request.Option) (*iotroborunner.ListSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListSitesOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListSitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListWorkerFleets(input *iotroborunner.ListWorkerFleetsInput) (*iotroborunner.ListWorkerFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListWorkerFleetsOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListWorkerFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListWorkerFleetsPages(input *iotroborunner.ListWorkerFleetsInput, fn func(*iotroborunner.ListWorkerFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListWorkerFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListWorkerFleetsOutput{}
	fnCacher := func(out *iotroborunner.ListWorkerFleetsOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListWorkerFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListWorkerFleetsPagesWithContext(ctx context.Context, input *iotroborunner.ListWorkerFleetsInput, fn func(*iotroborunner.ListWorkerFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListWorkerFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListWorkerFleetsOutput{}
	fnCacher := func(out *iotroborunner.ListWorkerFleetsOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListWorkerFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListWorkerFleetsWithContext(ctx context.Context, input *iotroborunner.ListWorkerFleetsInput, opts ...request.Option) (*iotroborunner.ListWorkerFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListWorkerFleetsOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListWorkerFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListWorkers(input *iotroborunner.ListWorkersInput) (*iotroborunner.ListWorkersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListWorkersOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListWorkers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTRoboRunner) ListWorkersPages(input *iotroborunner.ListWorkersInput, fn func(*iotroborunner.ListWorkersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListWorkersOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListWorkersOutput{}
	fnCacher := func(out *iotroborunner.ListWorkersOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListWorkersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListWorkersPagesWithContext(ctx context.Context, input *iotroborunner.ListWorkersInput, fn func(*iotroborunner.ListWorkersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotroborunner.ListWorkersOutput), false)
		return nil
	}
	cachable := true
	output := &iotroborunner.ListWorkersOutput{}
	fnCacher := func(out *iotroborunner.ListWorkersOutput, more bool) bool {
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
	if err := c.IoTRoboRunnerAPI.ListWorkersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTRoboRunner) ListWorkersWithContext(ctx context.Context, input *iotroborunner.ListWorkersInput, opts ...request.Option) (*iotroborunner.ListWorkersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotroborunner.ListWorkersOutput), nil
	}
	out, err := c.IoTRoboRunnerAPI.ListWorkersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

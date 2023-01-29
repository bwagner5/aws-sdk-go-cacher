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
package applicationcostprofilercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler/applicationcostprofileriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApplicationCostProfiler struct {
	applicationcostprofileriface.ApplicationCostProfilerAPI
	cache *cache.Cache
}

func New(applicationcostprofilerapi applicationcostprofileriface.ApplicationCostProfilerAPI) *ApplicationCostProfiler {
	return &ApplicationCostProfiler{
		ApplicationCostProfilerAPI: applicationcostprofilerapi,
		cache:                      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApplicationCostProfiler) GetReportDefinition(input *applicationcostprofiler.GetReportDefinitionInput) (*applicationcostprofiler.GetReportDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationcostprofiler.GetReportDefinitionOutput), nil
	}
	out, err := c.ApplicationCostProfilerAPI.GetReportDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationCostProfiler) GetReportDefinitionWithContext(ctx context.Context, input *applicationcostprofiler.GetReportDefinitionInput, opts ...request.Option) (*applicationcostprofiler.GetReportDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationcostprofiler.GetReportDefinitionOutput), nil
	}
	out, err := c.ApplicationCostProfilerAPI.GetReportDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationCostProfiler) ListReportDefinitions(input *applicationcostprofiler.ListReportDefinitionsInput) (*applicationcostprofiler.ListReportDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationcostprofiler.ListReportDefinitionsOutput), nil
	}
	out, err := c.ApplicationCostProfilerAPI.ListReportDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationCostProfiler) ListReportDefinitionsPages(input *applicationcostprofiler.ListReportDefinitionsInput, fn func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationcostprofiler.ListReportDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationcostprofiler.ListReportDefinitionsOutput{}
	fnCacher := func(out *applicationcostprofiler.ListReportDefinitionsOutput, more bool) bool {
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
	if err := c.ApplicationCostProfilerAPI.ListReportDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationCostProfiler) ListReportDefinitionsPagesWithContext(ctx context.Context, input *applicationcostprofiler.ListReportDefinitionsInput, fn func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationcostprofiler.ListReportDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationcostprofiler.ListReportDefinitionsOutput{}
	fnCacher := func(out *applicationcostprofiler.ListReportDefinitionsOutput, more bool) bool {
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
	if err := c.ApplicationCostProfilerAPI.ListReportDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationCostProfiler) ListReportDefinitionsWithContext(ctx context.Context, input *applicationcostprofiler.ListReportDefinitionsInput, opts ...request.Option) (*applicationcostprofiler.ListReportDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationcostprofiler.ListReportDefinitionsOutput), nil
	}
	out, err := c.ApplicationCostProfilerAPI.ListReportDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

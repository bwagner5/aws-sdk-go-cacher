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
package cloudwatchevidentlycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchevidently"
	"github.com/aws/aws-sdk-go/service/cloudwatchevidently/cloudwatchevidentlyiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudWatchEvidently struct {
	cloudwatchevidentlyiface.CloudWatchEvidentlyAPI
	cache *cache.Cache
}

func New(cloudwatchevidentlyapi cloudwatchevidentlyiface.CloudWatchEvidentlyAPI) *CloudWatchEvidently {
	return &CloudWatchEvidently{
		CloudWatchEvidentlyAPI: cloudwatchevidentlyapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudWatchEvidently) BatchEvaluateFeature(input *cloudwatchevidently.BatchEvaluateFeatureInput) (*cloudwatchevidently.BatchEvaluateFeatureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.BatchEvaluateFeatureOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.BatchEvaluateFeature(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) BatchEvaluateFeatureWithContext(ctx context.Context, input *cloudwatchevidently.BatchEvaluateFeatureInput, opts ...request.Option) (*cloudwatchevidently.BatchEvaluateFeatureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.BatchEvaluateFeatureOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.BatchEvaluateFeatureWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetExperiment(input *cloudwatchevidently.GetExperimentInput) (*cloudwatchevidently.GetExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetExperimentOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetExperiment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetExperimentResults(input *cloudwatchevidently.GetExperimentResultsInput) (*cloudwatchevidently.GetExperimentResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetExperimentResultsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetExperimentResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetExperimentResultsWithContext(ctx context.Context, input *cloudwatchevidently.GetExperimentResultsInput, opts ...request.Option) (*cloudwatchevidently.GetExperimentResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetExperimentResultsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetExperimentResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetExperimentWithContext(ctx context.Context, input *cloudwatchevidently.GetExperimentInput, opts ...request.Option) (*cloudwatchevidently.GetExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetExperimentOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetExperimentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetFeature(input *cloudwatchevidently.GetFeatureInput) (*cloudwatchevidently.GetFeatureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetFeatureOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetFeature(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetFeatureWithContext(ctx context.Context, input *cloudwatchevidently.GetFeatureInput, opts ...request.Option) (*cloudwatchevidently.GetFeatureOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetFeatureOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetFeatureWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetLaunch(input *cloudwatchevidently.GetLaunchInput) (*cloudwatchevidently.GetLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetLaunchOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetLaunch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetLaunchWithContext(ctx context.Context, input *cloudwatchevidently.GetLaunchInput, opts ...request.Option) (*cloudwatchevidently.GetLaunchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetLaunchOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetLaunchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetProject(input *cloudwatchevidently.GetProjectInput) (*cloudwatchevidently.GetProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetProjectOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetProjectWithContext(ctx context.Context, input *cloudwatchevidently.GetProjectInput, opts ...request.Option) (*cloudwatchevidently.GetProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetProjectOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetSegment(input *cloudwatchevidently.GetSegmentInput) (*cloudwatchevidently.GetSegmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetSegmentOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetSegment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) GetSegmentWithContext(ctx context.Context, input *cloudwatchevidently.GetSegmentInput, opts ...request.Option) (*cloudwatchevidently.GetSegmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.GetSegmentOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.GetSegmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListExperiments(input *cloudwatchevidently.ListExperimentsInput) (*cloudwatchevidently.ListExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListExperimentsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListExperiments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListExperimentsPages(input *cloudwatchevidently.ListExperimentsInput, fn func(*cloudwatchevidently.ListExperimentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListExperimentsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListExperimentsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListExperimentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListExperimentsPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListExperimentsInput, fn func(*cloudwatchevidently.ListExperimentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListExperimentsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListExperimentsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListExperimentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListExperimentsWithContext(ctx context.Context, input *cloudwatchevidently.ListExperimentsInput, opts ...request.Option) (*cloudwatchevidently.ListExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListExperimentsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListExperimentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListFeatures(input *cloudwatchevidently.ListFeaturesInput) (*cloudwatchevidently.ListFeaturesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListFeaturesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListFeatures(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListFeaturesPages(input *cloudwatchevidently.ListFeaturesInput, fn func(*cloudwatchevidently.ListFeaturesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListFeaturesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListFeaturesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListFeaturesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListFeaturesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListFeaturesPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListFeaturesInput, fn func(*cloudwatchevidently.ListFeaturesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListFeaturesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListFeaturesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListFeaturesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListFeaturesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListFeaturesWithContext(ctx context.Context, input *cloudwatchevidently.ListFeaturesInput, opts ...request.Option) (*cloudwatchevidently.ListFeaturesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListFeaturesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListFeaturesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListLaunches(input *cloudwatchevidently.ListLaunchesInput) (*cloudwatchevidently.ListLaunchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListLaunchesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListLaunches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListLaunchesPages(input *cloudwatchevidently.ListLaunchesInput, fn func(*cloudwatchevidently.ListLaunchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListLaunchesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListLaunchesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListLaunchesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListLaunchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListLaunchesPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListLaunchesInput, fn func(*cloudwatchevidently.ListLaunchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListLaunchesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListLaunchesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListLaunchesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListLaunchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListLaunchesWithContext(ctx context.Context, input *cloudwatchevidently.ListLaunchesInput, opts ...request.Option) (*cloudwatchevidently.ListLaunchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListLaunchesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListLaunchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListProjects(input *cloudwatchevidently.ListProjectsInput) (*cloudwatchevidently.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListProjectsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListProjectsPages(input *cloudwatchevidently.ListProjectsInput, fn func(*cloudwatchevidently.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListProjectsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListProjectsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListProjectsPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListProjectsInput, fn func(*cloudwatchevidently.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListProjectsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListProjectsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListProjectsWithContext(ctx context.Context, input *cloudwatchevidently.ListProjectsInput, opts ...request.Option) (*cloudwatchevidently.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListProjectsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListSegmentReferences(input *cloudwatchevidently.ListSegmentReferencesInput) (*cloudwatchevidently.ListSegmentReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListSegmentReferencesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListSegmentReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListSegmentReferencesPages(input *cloudwatchevidently.ListSegmentReferencesInput, fn func(*cloudwatchevidently.ListSegmentReferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListSegmentReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListSegmentReferencesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListSegmentReferencesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListSegmentReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListSegmentReferencesPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListSegmentReferencesInput, fn func(*cloudwatchevidently.ListSegmentReferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListSegmentReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListSegmentReferencesOutput{}
	fnCacher := func(out *cloudwatchevidently.ListSegmentReferencesOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListSegmentReferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListSegmentReferencesWithContext(ctx context.Context, input *cloudwatchevidently.ListSegmentReferencesInput, opts ...request.Option) (*cloudwatchevidently.ListSegmentReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListSegmentReferencesOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListSegmentReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListSegments(input *cloudwatchevidently.ListSegmentsInput) (*cloudwatchevidently.ListSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListSegmentsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListSegments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListSegmentsPages(input *cloudwatchevidently.ListSegmentsInput, fn func(*cloudwatchevidently.ListSegmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListSegmentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListSegmentsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListSegmentsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListSegmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListSegmentsPagesWithContext(ctx context.Context, input *cloudwatchevidently.ListSegmentsInput, fn func(*cloudwatchevidently.ListSegmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchevidently.ListSegmentsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchevidently.ListSegmentsOutput{}
	fnCacher := func(out *cloudwatchevidently.ListSegmentsOutput, more bool) bool {
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
	if err := c.CloudWatchEvidentlyAPI.ListSegmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchEvidently) ListSegmentsWithContext(ctx context.Context, input *cloudwatchevidently.ListSegmentsInput, opts ...request.Option) (*cloudwatchevidently.ListSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListSegmentsOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListSegmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListTagsForResource(input *cloudwatchevidently.ListTagsForResourceInput) (*cloudwatchevidently.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchEvidently) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatchevidently.ListTagsForResourceInput, opts ...request.Option) (*cloudwatchevidently.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchevidently.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchEvidentlyAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

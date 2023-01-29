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
	"github.com/aws/aws-sdk-go/service/lookoutequipment"
	"github.com/aws/aws-sdk-go/service/lookoutequipment/lookoutequipmentiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LookoutEquipment struct {
	lookoutequipmentiface.LookoutEquipmentAPI
	cache *cache.Cache
}

func NewLookoutEquipment(lookoutequipmentapi lookoutequipmentiface.LookoutEquipmentAPI) *LookoutEquipment {
	return &LookoutEquipment{
		LookoutEquipmentAPI: lookoutequipmentapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LookoutEquipment) DescribeDataIngestionJob(input *lookoutequipment.DescribeDataIngestionJobInput) (*lookoutequipment.DescribeDataIngestionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeDataIngestionJobOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeDataIngestionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeDataIngestionJobWithContext(ctx context.Context, input *lookoutequipment.DescribeDataIngestionJobInput, opts ...request.Option) (*lookoutequipment.DescribeDataIngestionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeDataIngestionJobOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeDataIngestionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeDataset(input *lookoutequipment.DescribeDatasetInput) (*lookoutequipment.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeDatasetOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeDatasetWithContext(ctx context.Context, input *lookoutequipment.DescribeDatasetInput, opts ...request.Option) (*lookoutequipment.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeDatasetOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeInferenceScheduler(input *lookoutequipment.DescribeInferenceSchedulerInput) (*lookoutequipment.DescribeInferenceSchedulerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeInferenceSchedulerOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeInferenceScheduler(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeInferenceSchedulerWithContext(ctx context.Context, input *lookoutequipment.DescribeInferenceSchedulerInput, opts ...request.Option) (*lookoutequipment.DescribeInferenceSchedulerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeInferenceSchedulerOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeInferenceSchedulerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeLabel(input *lookoutequipment.DescribeLabelInput) (*lookoutequipment.DescribeLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeLabelOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeLabel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeLabelGroup(input *lookoutequipment.DescribeLabelGroupInput) (*lookoutequipment.DescribeLabelGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeLabelGroupOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeLabelGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeLabelGroupWithContext(ctx context.Context, input *lookoutequipment.DescribeLabelGroupInput, opts ...request.Option) (*lookoutequipment.DescribeLabelGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeLabelGroupOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeLabelGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeLabelWithContext(ctx context.Context, input *lookoutequipment.DescribeLabelInput, opts ...request.Option) (*lookoutequipment.DescribeLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeLabelOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeLabelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeModel(input *lookoutequipment.DescribeModelInput) (*lookoutequipment.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeModelOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) DescribeModelWithContext(ctx context.Context, input *lookoutequipment.DescribeModelInput, opts ...request.Option) (*lookoutequipment.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.DescribeModelOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.DescribeModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListDataIngestionJobs(input *lookoutequipment.ListDataIngestionJobsInput) (*lookoutequipment.ListDataIngestionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListDataIngestionJobsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListDataIngestionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListDataIngestionJobsPages(input *lookoutequipment.ListDataIngestionJobsInput, fn func(*lookoutequipment.ListDataIngestionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListDataIngestionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListDataIngestionJobsOutput{}
	fnCacher := func(out *lookoutequipment.ListDataIngestionJobsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListDataIngestionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListDataIngestionJobsPagesWithContext(ctx context.Context, input *lookoutequipment.ListDataIngestionJobsInput, fn func(*lookoutequipment.ListDataIngestionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListDataIngestionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListDataIngestionJobsOutput{}
	fnCacher := func(out *lookoutequipment.ListDataIngestionJobsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListDataIngestionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListDataIngestionJobsWithContext(ctx context.Context, input *lookoutequipment.ListDataIngestionJobsInput, opts ...request.Option) (*lookoutequipment.ListDataIngestionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListDataIngestionJobsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListDataIngestionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListDatasets(input *lookoutequipment.ListDatasetsInput) (*lookoutequipment.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListDatasetsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListDatasetsPages(input *lookoutequipment.ListDatasetsInput, fn func(*lookoutequipment.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListDatasetsOutput{}
	fnCacher := func(out *lookoutequipment.ListDatasetsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListDatasetsPagesWithContext(ctx context.Context, input *lookoutequipment.ListDatasetsInput, fn func(*lookoutequipment.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListDatasetsOutput{}
	fnCacher := func(out *lookoutequipment.ListDatasetsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListDatasetsWithContext(ctx context.Context, input *lookoutequipment.ListDatasetsInput, opts ...request.Option) (*lookoutequipment.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListDatasetsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceEvents(input *lookoutequipment.ListInferenceEventsInput) (*lookoutequipment.ListInferenceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceEventsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceEventsPages(input *lookoutequipment.ListInferenceEventsInput, fn func(*lookoutequipment.ListInferenceEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceEventsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceEventsOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceEventsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceEventsPagesWithContext(ctx context.Context, input *lookoutequipment.ListInferenceEventsInput, fn func(*lookoutequipment.ListInferenceEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceEventsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceEventsOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceEventsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceEventsWithContext(ctx context.Context, input *lookoutequipment.ListInferenceEventsInput, opts ...request.Option) (*lookoutequipment.ListInferenceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceEventsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceExecutions(input *lookoutequipment.ListInferenceExecutionsInput) (*lookoutequipment.ListInferenceExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceExecutionsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceExecutionsPages(input *lookoutequipment.ListInferenceExecutionsInput, fn func(*lookoutequipment.ListInferenceExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceExecutionsOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceExecutionsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceExecutionsPagesWithContext(ctx context.Context, input *lookoutequipment.ListInferenceExecutionsInput, fn func(*lookoutequipment.ListInferenceExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceExecutionsOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceExecutionsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceExecutionsWithContext(ctx context.Context, input *lookoutequipment.ListInferenceExecutionsInput, opts ...request.Option) (*lookoutequipment.ListInferenceExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceExecutionsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceSchedulers(input *lookoutequipment.ListInferenceSchedulersInput) (*lookoutequipment.ListInferenceSchedulersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceSchedulersOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceSchedulers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListInferenceSchedulersPages(input *lookoutequipment.ListInferenceSchedulersInput, fn func(*lookoutequipment.ListInferenceSchedulersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceSchedulersOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceSchedulersOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceSchedulersOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceSchedulersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceSchedulersPagesWithContext(ctx context.Context, input *lookoutequipment.ListInferenceSchedulersInput, fn func(*lookoutequipment.ListInferenceSchedulersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListInferenceSchedulersOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListInferenceSchedulersOutput{}
	fnCacher := func(out *lookoutequipment.ListInferenceSchedulersOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListInferenceSchedulersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListInferenceSchedulersWithContext(ctx context.Context, input *lookoutequipment.ListInferenceSchedulersInput, opts ...request.Option) (*lookoutequipment.ListInferenceSchedulersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListInferenceSchedulersOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListInferenceSchedulersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListLabelGroups(input *lookoutequipment.ListLabelGroupsInput) (*lookoutequipment.ListLabelGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListLabelGroupsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListLabelGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListLabelGroupsPages(input *lookoutequipment.ListLabelGroupsInput, fn func(*lookoutequipment.ListLabelGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListLabelGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListLabelGroupsOutput{}
	fnCacher := func(out *lookoutequipment.ListLabelGroupsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListLabelGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListLabelGroupsPagesWithContext(ctx context.Context, input *lookoutequipment.ListLabelGroupsInput, fn func(*lookoutequipment.ListLabelGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListLabelGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListLabelGroupsOutput{}
	fnCacher := func(out *lookoutequipment.ListLabelGroupsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListLabelGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListLabelGroupsWithContext(ctx context.Context, input *lookoutequipment.ListLabelGroupsInput, opts ...request.Option) (*lookoutequipment.ListLabelGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListLabelGroupsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListLabelGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListLabels(input *lookoutequipment.ListLabelsInput) (*lookoutequipment.ListLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListLabelsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListLabels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListLabelsPages(input *lookoutequipment.ListLabelsInput, fn func(*lookoutequipment.ListLabelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListLabelsOutput{}
	fnCacher := func(out *lookoutequipment.ListLabelsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListLabelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListLabelsPagesWithContext(ctx context.Context, input *lookoutequipment.ListLabelsInput, fn func(*lookoutequipment.ListLabelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListLabelsOutput{}
	fnCacher := func(out *lookoutequipment.ListLabelsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListLabelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListLabelsWithContext(ctx context.Context, input *lookoutequipment.ListLabelsInput, opts ...request.Option) (*lookoutequipment.ListLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListLabelsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListLabelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListModels(input *lookoutequipment.ListModelsInput) (*lookoutequipment.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListModelsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListModelsPages(input *lookoutequipment.ListModelsInput, fn func(*lookoutequipment.ListModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListModelsOutput{}
	fnCacher := func(out *lookoutequipment.ListModelsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListModelsPagesWithContext(ctx context.Context, input *lookoutequipment.ListModelsInput, fn func(*lookoutequipment.ListModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListModelsOutput{}
	fnCacher := func(out *lookoutequipment.ListModelsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListModelsWithContext(ctx context.Context, input *lookoutequipment.ListModelsInput, opts ...request.Option) (*lookoutequipment.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListModelsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListSensorStatistics(input *lookoutequipment.ListSensorStatisticsInput) (*lookoutequipment.ListSensorStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListSensorStatisticsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListSensorStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListSensorStatisticsPages(input *lookoutequipment.ListSensorStatisticsInput, fn func(*lookoutequipment.ListSensorStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListSensorStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListSensorStatisticsOutput{}
	fnCacher := func(out *lookoutequipment.ListSensorStatisticsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListSensorStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListSensorStatisticsPagesWithContext(ctx context.Context, input *lookoutequipment.ListSensorStatisticsInput, fn func(*lookoutequipment.ListSensorStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutequipment.ListSensorStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutequipment.ListSensorStatisticsOutput{}
	fnCacher := func(out *lookoutequipment.ListSensorStatisticsOutput, more bool) bool {
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
	if err := c.LookoutEquipmentAPI.ListSensorStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutEquipment) ListSensorStatisticsWithContext(ctx context.Context, input *lookoutequipment.ListSensorStatisticsInput, opts ...request.Option) (*lookoutequipment.ListSensorStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListSensorStatisticsOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListSensorStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListTagsForResource(input *lookoutequipment.ListTagsForResourceInput) (*lookoutequipment.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutEquipment) ListTagsForResourceWithContext(ctx context.Context, input *lookoutequipment.ListTagsForResourceInput, opts ...request.Option) (*lookoutequipment.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutequipment.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutEquipmentAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

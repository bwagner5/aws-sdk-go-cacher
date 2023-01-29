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
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type FraudDetector struct {
	frauddetectoriface.FraudDetectorAPI
	cache *cache.Cache
}

func NewFraudDetector(frauddetectorapi frauddetectoriface.FraudDetectorAPI) *FraudDetector {
	return &FraudDetector{
		FraudDetectorAPI: frauddetectorapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *FraudDetector) BatchCreateVariable(input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.BatchCreateVariableOutput), nil
	}
	out, err := c.FraudDetectorAPI.BatchCreateVariable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) BatchCreateVariableWithContext(ctx context.Context, input *frauddetector.BatchCreateVariableInput, opts ...request.Option) (*frauddetector.BatchCreateVariableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.BatchCreateVariableOutput), nil
	}
	out, err := c.FraudDetectorAPI.BatchCreateVariableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) BatchGetVariable(input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.BatchGetVariableOutput), nil
	}
	out, err := c.FraudDetectorAPI.BatchGetVariable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) BatchGetVariableWithContext(ctx context.Context, input *frauddetector.BatchGetVariableInput, opts ...request.Option) (*frauddetector.BatchGetVariableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.BatchGetVariableOutput), nil
	}
	out, err := c.FraudDetectorAPI.BatchGetVariableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) DescribeDetector(input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.DescribeDetectorOutput), nil
	}
	out, err := c.FraudDetectorAPI.DescribeDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) DescribeDetectorWithContext(ctx context.Context, input *frauddetector.DescribeDetectorInput, opts ...request.Option) (*frauddetector.DescribeDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.DescribeDetectorOutput), nil
	}
	out, err := c.FraudDetectorAPI.DescribeDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) DescribeModelVersions(input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.DescribeModelVersionsOutput), nil
	}
	out, err := c.FraudDetectorAPI.DescribeModelVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) DescribeModelVersionsPages(input *frauddetector.DescribeModelVersionsInput, fn func(*frauddetector.DescribeModelVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.DescribeModelVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.DescribeModelVersionsOutput{}
	fnCacher := func(out *frauddetector.DescribeModelVersionsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.DescribeModelVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) DescribeModelVersionsPagesWithContext(ctx context.Context, input *frauddetector.DescribeModelVersionsInput, fn func(*frauddetector.DescribeModelVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.DescribeModelVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.DescribeModelVersionsOutput{}
	fnCacher := func(out *frauddetector.DescribeModelVersionsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.DescribeModelVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) DescribeModelVersionsWithContext(ctx context.Context, input *frauddetector.DescribeModelVersionsInput, opts ...request.Option) (*frauddetector.DescribeModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.DescribeModelVersionsOutput), nil
	}
	out, err := c.FraudDetectorAPI.DescribeModelVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetBatchImportJobs(input *frauddetector.GetBatchImportJobsInput) (*frauddetector.GetBatchImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetBatchImportJobsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetBatchImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetBatchImportJobsPages(input *frauddetector.GetBatchImportJobsInput, fn func(*frauddetector.GetBatchImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetBatchImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetBatchImportJobsOutput{}
	fnCacher := func(out *frauddetector.GetBatchImportJobsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetBatchImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetBatchImportJobsPagesWithContext(ctx context.Context, input *frauddetector.GetBatchImportJobsInput, fn func(*frauddetector.GetBatchImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetBatchImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetBatchImportJobsOutput{}
	fnCacher := func(out *frauddetector.GetBatchImportJobsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetBatchImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetBatchImportJobsWithContext(ctx context.Context, input *frauddetector.GetBatchImportJobsInput, opts ...request.Option) (*frauddetector.GetBatchImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetBatchImportJobsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetBatchImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetBatchPredictionJobs(input *frauddetector.GetBatchPredictionJobsInput) (*frauddetector.GetBatchPredictionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetBatchPredictionJobsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetBatchPredictionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetBatchPredictionJobsPages(input *frauddetector.GetBatchPredictionJobsInput, fn func(*frauddetector.GetBatchPredictionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetBatchPredictionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetBatchPredictionJobsOutput{}
	fnCacher := func(out *frauddetector.GetBatchPredictionJobsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetBatchPredictionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetBatchPredictionJobsPagesWithContext(ctx context.Context, input *frauddetector.GetBatchPredictionJobsInput, fn func(*frauddetector.GetBatchPredictionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetBatchPredictionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetBatchPredictionJobsOutput{}
	fnCacher := func(out *frauddetector.GetBatchPredictionJobsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetBatchPredictionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetBatchPredictionJobsWithContext(ctx context.Context, input *frauddetector.GetBatchPredictionJobsInput, opts ...request.Option) (*frauddetector.GetBatchPredictionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetBatchPredictionJobsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetBatchPredictionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDeleteEventsByEventTypeStatus(input *frauddetector.GetDeleteEventsByEventTypeStatusInput) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDeleteEventsByEventTypeStatusOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDeleteEventsByEventTypeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDeleteEventsByEventTypeStatusWithContext(ctx context.Context, input *frauddetector.GetDeleteEventsByEventTypeStatusInput, opts ...request.Option) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDeleteEventsByEventTypeStatusOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDeleteEventsByEventTypeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDetectorVersion(input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDetectorVersionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDetectorVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDetectorVersionWithContext(ctx context.Context, input *frauddetector.GetDetectorVersionInput, opts ...request.Option) (*frauddetector.GetDetectorVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDetectorVersionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDetectorVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDetectors(input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDetectorsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetDetectorsPages(input *frauddetector.GetDetectorsInput, fn func(*frauddetector.GetDetectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetDetectorsOutput{}
	fnCacher := func(out *frauddetector.GetDetectorsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetDetectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetDetectorsPagesWithContext(ctx context.Context, input *frauddetector.GetDetectorsInput, fn func(*frauddetector.GetDetectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetDetectorsOutput{}
	fnCacher := func(out *frauddetector.GetDetectorsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetDetectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetDetectorsWithContext(ctx context.Context, input *frauddetector.GetDetectorsInput, opts ...request.Option) (*frauddetector.GetDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetDetectorsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEntityTypes(input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEntityTypesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEntityTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEntityTypesPages(input *frauddetector.GetEntityTypesInput, fn func(*frauddetector.GetEntityTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetEntityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetEntityTypesOutput{}
	fnCacher := func(out *frauddetector.GetEntityTypesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetEntityTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetEntityTypesPagesWithContext(ctx context.Context, input *frauddetector.GetEntityTypesInput, fn func(*frauddetector.GetEntityTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetEntityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetEntityTypesOutput{}
	fnCacher := func(out *frauddetector.GetEntityTypesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetEntityTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetEntityTypesWithContext(ctx context.Context, input *frauddetector.GetEntityTypesInput, opts ...request.Option) (*frauddetector.GetEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEntityTypesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEntityTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEvent(input *frauddetector.GetEventInput) (*frauddetector.GetEventOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEvent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventPrediction(input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventPredictionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventPrediction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventPredictionMetadata(input *frauddetector.GetEventPredictionMetadataInput) (*frauddetector.GetEventPredictionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventPredictionMetadataOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventPredictionMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventPredictionMetadataWithContext(ctx context.Context, input *frauddetector.GetEventPredictionMetadataInput, opts ...request.Option) (*frauddetector.GetEventPredictionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventPredictionMetadataOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventPredictionMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventPredictionWithContext(ctx context.Context, input *frauddetector.GetEventPredictionInput, opts ...request.Option) (*frauddetector.GetEventPredictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventPredictionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventPredictionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventTypes(input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventTypesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventTypesPages(input *frauddetector.GetEventTypesInput, fn func(*frauddetector.GetEventTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetEventTypesOutput{}
	fnCacher := func(out *frauddetector.GetEventTypesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetEventTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetEventTypesPagesWithContext(ctx context.Context, input *frauddetector.GetEventTypesInput, fn func(*frauddetector.GetEventTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetEventTypesOutput{}
	fnCacher := func(out *frauddetector.GetEventTypesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetEventTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetEventTypesWithContext(ctx context.Context, input *frauddetector.GetEventTypesInput, opts ...request.Option) (*frauddetector.GetEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventTypesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetEventWithContext(ctx context.Context, input *frauddetector.GetEventInput, opts ...request.Option) (*frauddetector.GetEventOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetEventOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetEventWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetExternalModels(input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetExternalModelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetExternalModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetExternalModelsPages(input *frauddetector.GetExternalModelsInput, fn func(*frauddetector.GetExternalModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetExternalModelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetExternalModelsOutput{}
	fnCacher := func(out *frauddetector.GetExternalModelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetExternalModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetExternalModelsPagesWithContext(ctx context.Context, input *frauddetector.GetExternalModelsInput, fn func(*frauddetector.GetExternalModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetExternalModelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetExternalModelsOutput{}
	fnCacher := func(out *frauddetector.GetExternalModelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetExternalModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetExternalModelsWithContext(ctx context.Context, input *frauddetector.GetExternalModelsInput, opts ...request.Option) (*frauddetector.GetExternalModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetExternalModelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetExternalModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetKMSEncryptionKey(input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetKMSEncryptionKeyOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetKMSEncryptionKey(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetKMSEncryptionKeyWithContext(ctx context.Context, input *frauddetector.GetKMSEncryptionKeyInput, opts ...request.Option) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetKMSEncryptionKeyOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetKMSEncryptionKeyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetLabels(input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetLabelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetLabels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetLabelsPages(input *frauddetector.GetLabelsInput, fn func(*frauddetector.GetLabelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetLabelsOutput{}
	fnCacher := func(out *frauddetector.GetLabelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetLabelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetLabelsPagesWithContext(ctx context.Context, input *frauddetector.GetLabelsInput, fn func(*frauddetector.GetLabelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetLabelsOutput{}
	fnCacher := func(out *frauddetector.GetLabelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetLabelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetLabelsWithContext(ctx context.Context, input *frauddetector.GetLabelsInput, opts ...request.Option) (*frauddetector.GetLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetLabelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetLabelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetModelVersion(input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetModelVersionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetModelVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetModelVersionWithContext(ctx context.Context, input *frauddetector.GetModelVersionInput, opts ...request.Option) (*frauddetector.GetModelVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetModelVersionOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetModelVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetModels(input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetModelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetModelsPages(input *frauddetector.GetModelsInput, fn func(*frauddetector.GetModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetModelsOutput{}
	fnCacher := func(out *frauddetector.GetModelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetModelsPagesWithContext(ctx context.Context, input *frauddetector.GetModelsInput, fn func(*frauddetector.GetModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetModelsOutput{}
	fnCacher := func(out *frauddetector.GetModelsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetModelsWithContext(ctx context.Context, input *frauddetector.GetModelsInput, opts ...request.Option) (*frauddetector.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetModelsOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetOutcomes(input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetOutcomesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetOutcomes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetOutcomesPages(input *frauddetector.GetOutcomesInput, fn func(*frauddetector.GetOutcomesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetOutcomesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetOutcomesOutput{}
	fnCacher := func(out *frauddetector.GetOutcomesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetOutcomesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetOutcomesPagesWithContext(ctx context.Context, input *frauddetector.GetOutcomesInput, fn func(*frauddetector.GetOutcomesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetOutcomesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetOutcomesOutput{}
	fnCacher := func(out *frauddetector.GetOutcomesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetOutcomesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetOutcomesWithContext(ctx context.Context, input *frauddetector.GetOutcomesInput, opts ...request.Option) (*frauddetector.GetOutcomesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetOutcomesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetOutcomesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetRules(input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetRulesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetRulesPages(input *frauddetector.GetRulesInput, fn func(*frauddetector.GetRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetRulesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetRulesOutput{}
	fnCacher := func(out *frauddetector.GetRulesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetRulesPagesWithContext(ctx context.Context, input *frauddetector.GetRulesInput, fn func(*frauddetector.GetRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetRulesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetRulesOutput{}
	fnCacher := func(out *frauddetector.GetRulesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetRulesWithContext(ctx context.Context, input *frauddetector.GetRulesInput, opts ...request.Option) (*frauddetector.GetRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetRulesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetVariables(input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetVariablesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetVariables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) GetVariablesPages(input *frauddetector.GetVariablesInput, fn func(*frauddetector.GetVariablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetVariablesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetVariablesOutput{}
	fnCacher := func(out *frauddetector.GetVariablesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetVariablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetVariablesPagesWithContext(ctx context.Context, input *frauddetector.GetVariablesInput, fn func(*frauddetector.GetVariablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.GetVariablesOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.GetVariablesOutput{}
	fnCacher := func(out *frauddetector.GetVariablesOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.GetVariablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) GetVariablesWithContext(ctx context.Context, input *frauddetector.GetVariablesInput, opts ...request.Option) (*frauddetector.GetVariablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.GetVariablesOutput), nil
	}
	out, err := c.FraudDetectorAPI.GetVariablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) ListEventPredictions(input *frauddetector.ListEventPredictionsInput) (*frauddetector.ListEventPredictionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.ListEventPredictionsOutput), nil
	}
	out, err := c.FraudDetectorAPI.ListEventPredictions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) ListEventPredictionsPages(input *frauddetector.ListEventPredictionsInput, fn func(*frauddetector.ListEventPredictionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.ListEventPredictionsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.ListEventPredictionsOutput{}
	fnCacher := func(out *frauddetector.ListEventPredictionsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.ListEventPredictionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) ListEventPredictionsPagesWithContext(ctx context.Context, input *frauddetector.ListEventPredictionsInput, fn func(*frauddetector.ListEventPredictionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.ListEventPredictionsOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.ListEventPredictionsOutput{}
	fnCacher := func(out *frauddetector.ListEventPredictionsOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.ListEventPredictionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) ListEventPredictionsWithContext(ctx context.Context, input *frauddetector.ListEventPredictionsInput, opts ...request.Option) (*frauddetector.ListEventPredictionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.ListEventPredictionsOutput), nil
	}
	out, err := c.FraudDetectorAPI.ListEventPredictionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) ListTagsForResource(input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.ListTagsForResourceOutput), nil
	}
	out, err := c.FraudDetectorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FraudDetector) ListTagsForResourcePages(input *frauddetector.ListTagsForResourceInput, fn func(*frauddetector.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.ListTagsForResourceOutput{}
	fnCacher := func(out *frauddetector.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) ListTagsForResourcePagesWithContext(ctx context.Context, input *frauddetector.ListTagsForResourceInput, fn func(*frauddetector.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*frauddetector.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &frauddetector.ListTagsForResourceOutput{}
	fnCacher := func(out *frauddetector.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.FraudDetectorAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FraudDetector) ListTagsForResourceWithContext(ctx context.Context, input *frauddetector.ListTagsForResourceInput, opts ...request.Option) (*frauddetector.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*frauddetector.ListTagsForResourceOutput), nil
	}
	out, err := c.FraudDetectorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

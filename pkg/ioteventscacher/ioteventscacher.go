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
package ioteventscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/iotevents/ioteventsiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTEvents struct {
	ioteventsiface.IoTEventsAPI
	cache *cache.Cache
}

func New(ioteventsapi ioteventsiface.IoTEventsAPI) *IoTEvents {
	return &IoTEvents{
		IoTEventsAPI: ioteventsapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTEvents) DescribeAlarmModel(input *iotevents.DescribeAlarmModelInput) (*iotevents.DescribeAlarmModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeAlarmModelOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeAlarmModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeAlarmModelWithContext(ctx context.Context, input *iotevents.DescribeAlarmModelInput, opts ...request.Option) (*iotevents.DescribeAlarmModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeAlarmModelOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeAlarmModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeDetectorModel(input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeDetectorModelOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeDetectorModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeDetectorModelAnalysis(input *iotevents.DescribeDetectorModelAnalysisInput) (*iotevents.DescribeDetectorModelAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeDetectorModelAnalysisOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeDetectorModelAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeDetectorModelAnalysisWithContext(ctx context.Context, input *iotevents.DescribeDetectorModelAnalysisInput, opts ...request.Option) (*iotevents.DescribeDetectorModelAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeDetectorModelAnalysisOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeDetectorModelAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeDetectorModelWithContext(ctx context.Context, input *iotevents.DescribeDetectorModelInput, opts ...request.Option) (*iotevents.DescribeDetectorModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeDetectorModelOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeDetectorModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeInput(input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeInputOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeInput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeInputWithContext(ctx context.Context, input *iotevents.DescribeInputInput, opts ...request.Option) (*iotevents.DescribeInputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeInputOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeInputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeLoggingOptions(input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeLoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) DescribeLoggingOptionsWithContext(ctx context.Context, input *iotevents.DescribeLoggingOptionsInput, opts ...request.Option) (*iotevents.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTEventsAPI.DescribeLoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) GetDetectorModelAnalysisResults(input *iotevents.GetDetectorModelAnalysisResultsInput) (*iotevents.GetDetectorModelAnalysisResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.GetDetectorModelAnalysisResultsOutput), nil
	}
	out, err := c.IoTEventsAPI.GetDetectorModelAnalysisResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) GetDetectorModelAnalysisResultsWithContext(ctx context.Context, input *iotevents.GetDetectorModelAnalysisResultsInput, opts ...request.Option) (*iotevents.GetDetectorModelAnalysisResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.GetDetectorModelAnalysisResultsOutput), nil
	}
	out, err := c.IoTEventsAPI.GetDetectorModelAnalysisResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListAlarmModelVersions(input *iotevents.ListAlarmModelVersionsInput) (*iotevents.ListAlarmModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListAlarmModelVersionsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListAlarmModelVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListAlarmModelVersionsWithContext(ctx context.Context, input *iotevents.ListAlarmModelVersionsInput, opts ...request.Option) (*iotevents.ListAlarmModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListAlarmModelVersionsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListAlarmModelVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListAlarmModels(input *iotevents.ListAlarmModelsInput) (*iotevents.ListAlarmModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListAlarmModelsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListAlarmModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListAlarmModelsWithContext(ctx context.Context, input *iotevents.ListAlarmModelsInput, opts ...request.Option) (*iotevents.ListAlarmModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListAlarmModelsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListAlarmModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListDetectorModelVersions(input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListDetectorModelVersionsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListDetectorModelVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListDetectorModelVersionsWithContext(ctx context.Context, input *iotevents.ListDetectorModelVersionsInput, opts ...request.Option) (*iotevents.ListDetectorModelVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListDetectorModelVersionsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListDetectorModelVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListDetectorModels(input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListDetectorModelsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListDetectorModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListDetectorModelsWithContext(ctx context.Context, input *iotevents.ListDetectorModelsInput, opts ...request.Option) (*iotevents.ListDetectorModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListDetectorModelsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListDetectorModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListInputRoutings(input *iotevents.ListInputRoutingsInput) (*iotevents.ListInputRoutingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListInputRoutingsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListInputRoutings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListInputRoutingsWithContext(ctx context.Context, input *iotevents.ListInputRoutingsInput, opts ...request.Option) (*iotevents.ListInputRoutingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListInputRoutingsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListInputRoutingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListInputs(input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListInputsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListInputs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListInputsWithContext(ctx context.Context, input *iotevents.ListInputsInput, opts ...request.Option) (*iotevents.ListInputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListInputsOutput), nil
	}
	out, err := c.IoTEventsAPI.ListInputsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListTagsForResource(input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTEventsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEvents) ListTagsForResourceWithContext(ctx context.Context, input *iotevents.ListTagsForResourceInput, opts ...request.Option) (*iotevents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotevents.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTEventsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

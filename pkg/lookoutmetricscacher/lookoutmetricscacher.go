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
package lookoutmetricscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lookoutmetrics"
	"github.com/aws/aws-sdk-go/service/lookoutmetrics/lookoutmetricsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LookoutMetrics struct {
	lookoutmetricsiface.LookoutMetricsAPI
	cache *cache.Cache
}

func New(lookoutmetricsapi lookoutmetricsiface.LookoutMetricsAPI) *LookoutMetrics {
	return &LookoutMetrics{
		LookoutMetricsAPI: lookoutmetricsapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LookoutMetrics) DescribeAlert(input *lookoutmetrics.DescribeAlertInput) (*lookoutmetrics.DescribeAlertOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAlertOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAlert(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeAlertWithContext(ctx context.Context, input *lookoutmetrics.DescribeAlertInput, opts ...request.Option) (*lookoutmetrics.DescribeAlertOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAlertOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAlertWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeAnomalyDetectionExecutions(input *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput) (*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAnomalyDetectionExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeAnomalyDetectionExecutionsPages(input *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, fn func(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput{}
	fnCacher := func(out *lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.DescribeAnomalyDetectionExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) DescribeAnomalyDetectionExecutionsPagesWithContext(ctx context.Context, input *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, fn func(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput{}
	fnCacher := func(out *lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.DescribeAnomalyDetectionExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) DescribeAnomalyDetectionExecutionsWithContext(ctx context.Context, input *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, opts ...request.Option) (*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAnomalyDetectionExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeAnomalyDetector(input *lookoutmetrics.DescribeAnomalyDetectorInput) (*lookoutmetrics.DescribeAnomalyDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectorOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAnomalyDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeAnomalyDetectorWithContext(ctx context.Context, input *lookoutmetrics.DescribeAnomalyDetectorInput, opts ...request.Option) (*lookoutmetrics.DescribeAnomalyDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeAnomalyDetectorOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeAnomalyDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeMetricSet(input *lookoutmetrics.DescribeMetricSetInput) (*lookoutmetrics.DescribeMetricSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeMetricSetOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeMetricSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) DescribeMetricSetWithContext(ctx context.Context, input *lookoutmetrics.DescribeMetricSetInput, opts ...request.Option) (*lookoutmetrics.DescribeMetricSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.DescribeMetricSetOutput), nil
	}
	out, err := c.LookoutMetricsAPI.DescribeMetricSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetAnomalyGroup(input *lookoutmetrics.GetAnomalyGroupInput) (*lookoutmetrics.GetAnomalyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetAnomalyGroupOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetAnomalyGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetAnomalyGroupWithContext(ctx context.Context, input *lookoutmetrics.GetAnomalyGroupInput, opts ...request.Option) (*lookoutmetrics.GetAnomalyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetAnomalyGroupOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetAnomalyGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetDataQualityMetrics(input *lookoutmetrics.GetDataQualityMetricsInput) (*lookoutmetrics.GetDataQualityMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetDataQualityMetricsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetDataQualityMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetDataQualityMetricsWithContext(ctx context.Context, input *lookoutmetrics.GetDataQualityMetricsInput, opts ...request.Option) (*lookoutmetrics.GetDataQualityMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetDataQualityMetricsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetDataQualityMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetFeedback(input *lookoutmetrics.GetFeedbackInput) (*lookoutmetrics.GetFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetFeedbackOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetFeedback(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetFeedbackPages(input *lookoutmetrics.GetFeedbackInput, fn func(*lookoutmetrics.GetFeedbackOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.GetFeedbackOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.GetFeedbackOutput{}
	fnCacher := func(out *lookoutmetrics.GetFeedbackOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.GetFeedbackPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) GetFeedbackPagesWithContext(ctx context.Context, input *lookoutmetrics.GetFeedbackInput, fn func(*lookoutmetrics.GetFeedbackOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.GetFeedbackOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.GetFeedbackOutput{}
	fnCacher := func(out *lookoutmetrics.GetFeedbackOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.GetFeedbackPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) GetFeedbackWithContext(ctx context.Context, input *lookoutmetrics.GetFeedbackInput, opts ...request.Option) (*lookoutmetrics.GetFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetFeedbackOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetFeedbackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetSampleData(input *lookoutmetrics.GetSampleDataInput) (*lookoutmetrics.GetSampleDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetSampleDataOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetSampleData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) GetSampleDataWithContext(ctx context.Context, input *lookoutmetrics.GetSampleDataInput, opts ...request.Option) (*lookoutmetrics.GetSampleDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.GetSampleDataOutput), nil
	}
	out, err := c.LookoutMetricsAPI.GetSampleDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAlerts(input *lookoutmetrics.ListAlertsInput) (*lookoutmetrics.ListAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAlertsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAlerts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAlertsPages(input *lookoutmetrics.ListAlertsInput, fn func(*lookoutmetrics.ListAlertsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAlertsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAlertsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAlertsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAlertsPagesWithContext(ctx context.Context, input *lookoutmetrics.ListAlertsInput, fn func(*lookoutmetrics.ListAlertsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAlertsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAlertsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAlertsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAlertsWithContext(ctx context.Context, input *lookoutmetrics.ListAlertsInput, opts ...request.Option) (*lookoutmetrics.ListAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAlertsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAlertsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyDetectors(input *lookoutmetrics.ListAnomalyDetectorsInput) (*lookoutmetrics.ListAnomalyDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyDetectorsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyDetectorsPages(input *lookoutmetrics.ListAnomalyDetectorsInput, fn func(*lookoutmetrics.ListAnomalyDetectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyDetectorsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyDetectorsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyDetectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyDetectorsPagesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyDetectorsInput, fn func(*lookoutmetrics.ListAnomalyDetectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyDetectorsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyDetectorsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyDetectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyDetectorsWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyDetectorsInput, opts ...request.Option) (*lookoutmetrics.ListAnomalyDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyDetectorsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupRelatedMetrics(input *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput) (*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupRelatedMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupRelatedMetricsPages(input *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, fn func(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupRelatedMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupRelatedMetricsPagesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, fn func(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupRelatedMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupRelatedMetricsWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, opts ...request.Option) (*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupRelatedMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupSummaries(input *lookoutmetrics.ListAnomalyGroupSummariesInput) (*lookoutmetrics.ListAnomalyGroupSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupSummariesOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupSummariesPages(input *lookoutmetrics.ListAnomalyGroupSummariesInput, fn func(*lookoutmetrics.ListAnomalyGroupSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupSummariesOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupSummariesOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupSummariesPagesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupSummariesInput, fn func(*lookoutmetrics.ListAnomalyGroupSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupSummariesOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupSummariesOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupSummariesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupSummariesInput, opts ...request.Option) (*lookoutmetrics.ListAnomalyGroupSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupSummariesOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupTimeSeries(input *lookoutmetrics.ListAnomalyGroupTimeSeriesInput) (*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupTimeSeries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListAnomalyGroupTimeSeriesPages(input *lookoutmetrics.ListAnomalyGroupTimeSeriesInput, fn func(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupTimeSeriesOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupTimeSeriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupTimeSeriesPagesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupTimeSeriesInput, fn func(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListAnomalyGroupTimeSeriesOutput{}
	fnCacher := func(out *lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListAnomalyGroupTimeSeriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListAnomalyGroupTimeSeriesWithContext(ctx context.Context, input *lookoutmetrics.ListAnomalyGroupTimeSeriesInput, opts ...request.Option) (*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListAnomalyGroupTimeSeriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListMetricSets(input *lookoutmetrics.ListMetricSetsInput) (*lookoutmetrics.ListMetricSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListMetricSetsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListMetricSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListMetricSetsPages(input *lookoutmetrics.ListMetricSetsInput, fn func(*lookoutmetrics.ListMetricSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListMetricSetsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListMetricSetsOutput{}
	fnCacher := func(out *lookoutmetrics.ListMetricSetsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListMetricSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListMetricSetsPagesWithContext(ctx context.Context, input *lookoutmetrics.ListMetricSetsInput, fn func(*lookoutmetrics.ListMetricSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutmetrics.ListMetricSetsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutmetrics.ListMetricSetsOutput{}
	fnCacher := func(out *lookoutmetrics.ListMetricSetsOutput, more bool) bool {
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
	if err := c.LookoutMetricsAPI.ListMetricSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutMetrics) ListMetricSetsWithContext(ctx context.Context, input *lookoutmetrics.ListMetricSetsInput, opts ...request.Option) (*lookoutmetrics.ListMetricSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListMetricSetsOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListMetricSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListTagsForResource(input *lookoutmetrics.ListTagsForResourceInput) (*lookoutmetrics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutMetrics) ListTagsForResourceWithContext(ctx context.Context, input *lookoutmetrics.ListTagsForResourceInput, opts ...request.Option) (*lookoutmetrics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutmetrics.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutMetricsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

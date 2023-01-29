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
package cloudwatchcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudWatch struct {
	cloudwatchiface.CloudWatchAPI
	cache *cache.Cache
}

func New(cloudwatchapi cloudwatchiface.CloudWatchAPI) *CloudWatch {
	return &CloudWatch{
		CloudWatchAPI: cloudwatchapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudWatch) DescribeAlarmHistory(input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmHistoryOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarmHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAlarmHistoryPages(input *cloudwatch.DescribeAlarmHistoryInput, fn func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAlarmHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAlarmHistoryOutput{}
	fnCacher := func(out *cloudwatch.DescribeAlarmHistoryOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAlarmHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAlarmHistoryPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, fn func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAlarmHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAlarmHistoryOutput{}
	fnCacher := func(out *cloudwatch.DescribeAlarmHistoryOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAlarmHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAlarmHistoryWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, opts ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmHistoryOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarmHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAlarms(input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmsOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAlarmsForMetric(input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmsForMetricOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarmsForMetric(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAlarmsForMetricWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmsForMetricOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarmsForMetricWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAlarmsPages(input *cloudwatch.DescribeAlarmsInput, fn func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAlarmsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAlarmsOutput{}
	fnCacher := func(out *cloudwatch.DescribeAlarmsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAlarmsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAlarmsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, fn func(*cloudwatch.DescribeAlarmsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAlarmsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAlarmsOutput{}
	fnCacher := func(out *cloudwatch.DescribeAlarmsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAlarmsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAlarmsWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAlarmsOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAlarmsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAnomalyDetectors(input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAnomalyDetectorsOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAnomalyDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeAnomalyDetectorsPages(input *cloudwatch.DescribeAnomalyDetectorsInput, fn func(*cloudwatch.DescribeAnomalyDetectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAnomalyDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAnomalyDetectorsOutput{}
	fnCacher := func(out *cloudwatch.DescribeAnomalyDetectorsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAnomalyDetectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAnomalyDetectorsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, fn func(*cloudwatch.DescribeAnomalyDetectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeAnomalyDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeAnomalyDetectorsOutput{}
	fnCacher := func(out *cloudwatch.DescribeAnomalyDetectorsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeAnomalyDetectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeAnomalyDetectorsWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, opts ...request.Option) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeAnomalyDetectorsOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeAnomalyDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeInsightRules(input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeInsightRulesOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeInsightRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) DescribeInsightRulesPages(input *cloudwatch.DescribeInsightRulesInput, fn func(*cloudwatch.DescribeInsightRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeInsightRulesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeInsightRulesOutput{}
	fnCacher := func(out *cloudwatch.DescribeInsightRulesOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeInsightRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, fn func(*cloudwatch.DescribeInsightRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.DescribeInsightRulesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.DescribeInsightRulesOutput{}
	fnCacher := func(out *cloudwatch.DescribeInsightRulesOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.DescribeInsightRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) DescribeInsightRulesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, opts ...request.Option) (*cloudwatch.DescribeInsightRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.DescribeInsightRulesOutput), nil
	}
	out, err := c.CloudWatchAPI.DescribeInsightRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetDashboard(input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetDashboardOutput), nil
	}
	out, err := c.CloudWatchAPI.GetDashboard(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetDashboardWithContext(ctx context.Context, input *cloudwatch.GetDashboardInput, opts ...request.Option) (*cloudwatch.GetDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetDashboardOutput), nil
	}
	out, err := c.CloudWatchAPI.GetDashboardWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetInsightRuleReport(input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetInsightRuleReportOutput), nil
	}
	out, err := c.CloudWatchAPI.GetInsightRuleReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetInsightRuleReportWithContext(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput, opts ...request.Option) (*cloudwatch.GetInsightRuleReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetInsightRuleReportOutput), nil
	}
	out, err := c.CloudWatchAPI.GetInsightRuleReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricData(input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricDataOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricDataPages(input *cloudwatch.GetMetricDataInput, fn func(*cloudwatch.GetMetricDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.GetMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.GetMetricDataOutput{}
	fnCacher := func(out *cloudwatch.GetMetricDataOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.GetMetricDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) GetMetricDataPagesWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, fn func(*cloudwatch.GetMetricDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.GetMetricDataOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.GetMetricDataOutput{}
	fnCacher := func(out *cloudwatch.GetMetricDataOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.GetMetricDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) GetMetricDataWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, opts ...request.Option) (*cloudwatch.GetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricDataOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricStatistics(input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricStatisticsOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricStatisticsWithContext(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput, opts ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricStatisticsOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricStream(input *cloudwatch.GetMetricStreamInput) (*cloudwatch.GetMetricStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricStreamOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricStreamWithContext(ctx context.Context, input *cloudwatch.GetMetricStreamInput, opts ...request.Option) (*cloudwatch.GetMetricStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricStreamOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricWidgetImage(input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricWidgetImageOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricWidgetImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) GetMetricWidgetImageWithContext(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput, opts ...request.Option) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.GetMetricWidgetImageOutput), nil
	}
	out, err := c.CloudWatchAPI.GetMetricWidgetImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListDashboards(input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListDashboardsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListDashboards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListDashboardsPages(input *cloudwatch.ListDashboardsInput, fn func(*cloudwatch.ListDashboardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListDashboardsOutput{}
	fnCacher := func(out *cloudwatch.ListDashboardsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListDashboardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListDashboardsPagesWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, fn func(*cloudwatch.ListDashboardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListDashboardsOutput{}
	fnCacher := func(out *cloudwatch.ListDashboardsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListDashboardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListDashboardsWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, opts ...request.Option) (*cloudwatch.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListDashboardsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListDashboardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListManagedInsightRules(input *cloudwatch.ListManagedInsightRulesInput) (*cloudwatch.ListManagedInsightRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListManagedInsightRulesOutput), nil
	}
	out, err := c.CloudWatchAPI.ListManagedInsightRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListManagedInsightRulesPages(input *cloudwatch.ListManagedInsightRulesInput, fn func(*cloudwatch.ListManagedInsightRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListManagedInsightRulesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListManagedInsightRulesOutput{}
	fnCacher := func(out *cloudwatch.ListManagedInsightRulesOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListManagedInsightRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListManagedInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, fn func(*cloudwatch.ListManagedInsightRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListManagedInsightRulesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListManagedInsightRulesOutput{}
	fnCacher := func(out *cloudwatch.ListManagedInsightRulesOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListManagedInsightRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListManagedInsightRulesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, opts ...request.Option) (*cloudwatch.ListManagedInsightRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListManagedInsightRulesOutput), nil
	}
	out, err := c.CloudWatchAPI.ListManagedInsightRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListMetricStreams(input *cloudwatch.ListMetricStreamsInput) (*cloudwatch.ListMetricStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListMetricStreamsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListMetricStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListMetricStreamsPages(input *cloudwatch.ListMetricStreamsInput, fn func(*cloudwatch.ListMetricStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListMetricStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListMetricStreamsOutput{}
	fnCacher := func(out *cloudwatch.ListMetricStreamsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListMetricStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListMetricStreamsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, fn func(*cloudwatch.ListMetricStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListMetricStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListMetricStreamsOutput{}
	fnCacher := func(out *cloudwatch.ListMetricStreamsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListMetricStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListMetricStreamsWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, opts ...request.Option) (*cloudwatch.ListMetricStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListMetricStreamsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListMetricStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListMetrics(input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListMetricsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListMetricsPages(input *cloudwatch.ListMetricsInput, fn func(*cloudwatch.ListMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListMetricsOutput{}
	fnCacher := func(out *cloudwatch.ListMetricsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListMetricsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, fn func(*cloudwatch.ListMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatch.ListMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatch.ListMetricsOutput{}
	fnCacher := func(out *cloudwatch.ListMetricsOutput, more bool) bool {
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
	if err := c.CloudWatchAPI.ListMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatch) ListMetricsWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, opts ...request.Option) (*cloudwatch.ListMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListMetricsOutput), nil
	}
	out, err := c.CloudWatchAPI.ListMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListTagsForResource(input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatch) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatch.ListTagsForResourceInput, opts ...request.Option) (*cloudwatch.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatch.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

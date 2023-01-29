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
package forecastservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ForecastService struct {
	forecastserviceiface.ForecastServiceAPI
	cache *cache.Cache
}

func New(forecastserviceapi forecastserviceiface.ForecastServiceAPI) *ForecastService {
	return &ForecastService{
		ForecastServiceAPI: forecastserviceapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ForecastService) DescribeAutoPredictor(input *forecastservice.DescribeAutoPredictorInput) (*forecastservice.DescribeAutoPredictorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeAutoPredictorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeAutoPredictor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeAutoPredictorWithContext(ctx context.Context, input *forecastservice.DescribeAutoPredictorInput, opts ...request.Option) (*forecastservice.DescribeAutoPredictorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeAutoPredictorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeAutoPredictorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDataset(input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDatasetGroup(input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetGroupOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDatasetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDatasetGroupWithContext(ctx context.Context, input *forecastservice.DescribeDatasetGroupInput, opts ...request.Option) (*forecastservice.DescribeDatasetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetGroupOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDatasetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDatasetImportJob(input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetImportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDatasetImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDatasetImportJobWithContext(ctx context.Context, input *forecastservice.DescribeDatasetImportJobInput, opts ...request.Option) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetImportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDatasetImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeDatasetWithContext(ctx context.Context, input *forecastservice.DescribeDatasetInput, opts ...request.Option) (*forecastservice.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeDatasetOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeExplainability(input *forecastservice.DescribeExplainabilityInput) (*forecastservice.DescribeExplainabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeExplainabilityOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeExplainability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeExplainabilityExport(input *forecastservice.DescribeExplainabilityExportInput) (*forecastservice.DescribeExplainabilityExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeExplainabilityExportOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeExplainabilityExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeExplainabilityExportWithContext(ctx context.Context, input *forecastservice.DescribeExplainabilityExportInput, opts ...request.Option) (*forecastservice.DescribeExplainabilityExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeExplainabilityExportOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeExplainabilityExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeExplainabilityWithContext(ctx context.Context, input *forecastservice.DescribeExplainabilityInput, opts ...request.Option) (*forecastservice.DescribeExplainabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeExplainabilityOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeExplainabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeForecast(input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeForecastOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeForecastExportJob(input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeForecastExportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeForecastExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeForecastExportJobWithContext(ctx context.Context, input *forecastservice.DescribeForecastExportJobInput, opts ...request.Option) (*forecastservice.DescribeForecastExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeForecastExportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeForecastExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeForecastWithContext(ctx context.Context, input *forecastservice.DescribeForecastInput, opts ...request.Option) (*forecastservice.DescribeForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeForecastOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeMonitor(input *forecastservice.DescribeMonitorInput) (*forecastservice.DescribeMonitorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeMonitorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeMonitor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeMonitorWithContext(ctx context.Context, input *forecastservice.DescribeMonitorInput, opts ...request.Option) (*forecastservice.DescribeMonitorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeMonitorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeMonitorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribePredictor(input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribePredictorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribePredictor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribePredictorBacktestExportJob(input *forecastservice.DescribePredictorBacktestExportJobInput) (*forecastservice.DescribePredictorBacktestExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribePredictorBacktestExportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribePredictorBacktestExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribePredictorBacktestExportJobWithContext(ctx context.Context, input *forecastservice.DescribePredictorBacktestExportJobInput, opts ...request.Option) (*forecastservice.DescribePredictorBacktestExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribePredictorBacktestExportJobOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribePredictorBacktestExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribePredictorWithContext(ctx context.Context, input *forecastservice.DescribePredictorInput, opts ...request.Option) (*forecastservice.DescribePredictorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribePredictorOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribePredictorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfAnalysis(input *forecastservice.DescribeWhatIfAnalysisInput) (*forecastservice.DescribeWhatIfAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfAnalysisOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfAnalysisWithContext(ctx context.Context, input *forecastservice.DescribeWhatIfAnalysisInput, opts ...request.Option) (*forecastservice.DescribeWhatIfAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfAnalysisOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfForecast(input *forecastservice.DescribeWhatIfForecastInput) (*forecastservice.DescribeWhatIfForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfForecastOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfForecastExport(input *forecastservice.DescribeWhatIfForecastExportInput) (*forecastservice.DescribeWhatIfForecastExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfForecastExportOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfForecastExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfForecastExportWithContext(ctx context.Context, input *forecastservice.DescribeWhatIfForecastExportInput, opts ...request.Option) (*forecastservice.DescribeWhatIfForecastExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfForecastExportOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfForecastExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) DescribeWhatIfForecastWithContext(ctx context.Context, input *forecastservice.DescribeWhatIfForecastInput, opts ...request.Option) (*forecastservice.DescribeWhatIfForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.DescribeWhatIfForecastOutput), nil
	}
	out, err := c.ForecastServiceAPI.DescribeWhatIfForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) GetAccuracyMetrics(input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.GetAccuracyMetricsOutput), nil
	}
	out, err := c.ForecastServiceAPI.GetAccuracyMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) GetAccuracyMetricsWithContext(ctx context.Context, input *forecastservice.GetAccuracyMetricsInput, opts ...request.Option) (*forecastservice.GetAccuracyMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.GetAccuracyMetricsOutput), nil
	}
	out, err := c.ForecastServiceAPI.GetAccuracyMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasetGroups(input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetGroupsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasetGroupsPages(input *forecastservice.ListDatasetGroupsInput, fn func(*forecastservice.ListDatasetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetGroupsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetGroupsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetGroupsPagesWithContext(ctx context.Context, input *forecastservice.ListDatasetGroupsInput, fn func(*forecastservice.ListDatasetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetGroupsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetGroupsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetGroupsWithContext(ctx context.Context, input *forecastservice.ListDatasetGroupsInput, opts ...request.Option) (*forecastservice.ListDatasetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetGroupsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasetImportJobs(input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetImportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasetImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasetImportJobsPages(input *forecastservice.ListDatasetImportJobsInput, fn func(*forecastservice.ListDatasetImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetImportJobsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetImportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetImportJobsPagesWithContext(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput, fn func(*forecastservice.ListDatasetImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetImportJobsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetImportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetImportJobsWithContext(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput, opts ...request.Option) (*forecastservice.ListDatasetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetImportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasetImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasets(input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListDatasetsPages(input *forecastservice.ListDatasetsInput, fn func(*forecastservice.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetsPagesWithContext(ctx context.Context, input *forecastservice.ListDatasetsInput, fn func(*forecastservice.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListDatasetsOutput{}
	fnCacher := func(out *forecastservice.ListDatasetsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListDatasetsWithContext(ctx context.Context, input *forecastservice.ListDatasetsInput, opts ...request.Option) (*forecastservice.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListDatasetsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListExplainabilities(input *forecastservice.ListExplainabilitiesInput) (*forecastservice.ListExplainabilitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListExplainabilitiesOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListExplainabilities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListExplainabilitiesPages(input *forecastservice.ListExplainabilitiesInput, fn func(*forecastservice.ListExplainabilitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListExplainabilitiesOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListExplainabilitiesOutput{}
	fnCacher := func(out *forecastservice.ListExplainabilitiesOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListExplainabilitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListExplainabilitiesPagesWithContext(ctx context.Context, input *forecastservice.ListExplainabilitiesInput, fn func(*forecastservice.ListExplainabilitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListExplainabilitiesOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListExplainabilitiesOutput{}
	fnCacher := func(out *forecastservice.ListExplainabilitiesOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListExplainabilitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListExplainabilitiesWithContext(ctx context.Context, input *forecastservice.ListExplainabilitiesInput, opts ...request.Option) (*forecastservice.ListExplainabilitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListExplainabilitiesOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListExplainabilitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListExplainabilityExports(input *forecastservice.ListExplainabilityExportsInput) (*forecastservice.ListExplainabilityExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListExplainabilityExportsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListExplainabilityExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListExplainabilityExportsPages(input *forecastservice.ListExplainabilityExportsInput, fn func(*forecastservice.ListExplainabilityExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListExplainabilityExportsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListExplainabilityExportsOutput{}
	fnCacher := func(out *forecastservice.ListExplainabilityExportsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListExplainabilityExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListExplainabilityExportsPagesWithContext(ctx context.Context, input *forecastservice.ListExplainabilityExportsInput, fn func(*forecastservice.ListExplainabilityExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListExplainabilityExportsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListExplainabilityExportsOutput{}
	fnCacher := func(out *forecastservice.ListExplainabilityExportsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListExplainabilityExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListExplainabilityExportsWithContext(ctx context.Context, input *forecastservice.ListExplainabilityExportsInput, opts ...request.Option) (*forecastservice.ListExplainabilityExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListExplainabilityExportsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListExplainabilityExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListForecastExportJobs(input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListForecastExportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListForecastExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListForecastExportJobsPages(input *forecastservice.ListForecastExportJobsInput, fn func(*forecastservice.ListForecastExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListForecastExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListForecastExportJobsOutput{}
	fnCacher := func(out *forecastservice.ListForecastExportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListForecastExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListForecastExportJobsPagesWithContext(ctx context.Context, input *forecastservice.ListForecastExportJobsInput, fn func(*forecastservice.ListForecastExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListForecastExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListForecastExportJobsOutput{}
	fnCacher := func(out *forecastservice.ListForecastExportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListForecastExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListForecastExportJobsWithContext(ctx context.Context, input *forecastservice.ListForecastExportJobsInput, opts ...request.Option) (*forecastservice.ListForecastExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListForecastExportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListForecastExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListForecasts(input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListForecastsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListForecasts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListForecastsPages(input *forecastservice.ListForecastsInput, fn func(*forecastservice.ListForecastsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListForecastsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListForecastsOutput{}
	fnCacher := func(out *forecastservice.ListForecastsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListForecastsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListForecastsPagesWithContext(ctx context.Context, input *forecastservice.ListForecastsInput, fn func(*forecastservice.ListForecastsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListForecastsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListForecastsOutput{}
	fnCacher := func(out *forecastservice.ListForecastsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListForecastsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListForecastsWithContext(ctx context.Context, input *forecastservice.ListForecastsInput, opts ...request.Option) (*forecastservice.ListForecastsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListForecastsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListForecastsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListMonitorEvaluations(input *forecastservice.ListMonitorEvaluationsInput) (*forecastservice.ListMonitorEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListMonitorEvaluationsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListMonitorEvaluations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListMonitorEvaluationsPages(input *forecastservice.ListMonitorEvaluationsInput, fn func(*forecastservice.ListMonitorEvaluationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListMonitorEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListMonitorEvaluationsOutput{}
	fnCacher := func(out *forecastservice.ListMonitorEvaluationsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListMonitorEvaluationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListMonitorEvaluationsPagesWithContext(ctx context.Context, input *forecastservice.ListMonitorEvaluationsInput, fn func(*forecastservice.ListMonitorEvaluationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListMonitorEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListMonitorEvaluationsOutput{}
	fnCacher := func(out *forecastservice.ListMonitorEvaluationsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListMonitorEvaluationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListMonitorEvaluationsWithContext(ctx context.Context, input *forecastservice.ListMonitorEvaluationsInput, opts ...request.Option) (*forecastservice.ListMonitorEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListMonitorEvaluationsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListMonitorEvaluationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListMonitors(input *forecastservice.ListMonitorsInput) (*forecastservice.ListMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListMonitorsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListMonitors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListMonitorsPages(input *forecastservice.ListMonitorsInput, fn func(*forecastservice.ListMonitorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListMonitorsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListMonitorsOutput{}
	fnCacher := func(out *forecastservice.ListMonitorsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListMonitorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListMonitorsPagesWithContext(ctx context.Context, input *forecastservice.ListMonitorsInput, fn func(*forecastservice.ListMonitorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListMonitorsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListMonitorsOutput{}
	fnCacher := func(out *forecastservice.ListMonitorsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListMonitorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListMonitorsWithContext(ctx context.Context, input *forecastservice.ListMonitorsInput, opts ...request.Option) (*forecastservice.ListMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListMonitorsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListMonitorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListPredictorBacktestExportJobs(input *forecastservice.ListPredictorBacktestExportJobsInput) (*forecastservice.ListPredictorBacktestExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListPredictorBacktestExportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListPredictorBacktestExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListPredictorBacktestExportJobsPages(input *forecastservice.ListPredictorBacktestExportJobsInput, fn func(*forecastservice.ListPredictorBacktestExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListPredictorBacktestExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListPredictorBacktestExportJobsOutput{}
	fnCacher := func(out *forecastservice.ListPredictorBacktestExportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListPredictorBacktestExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListPredictorBacktestExportJobsPagesWithContext(ctx context.Context, input *forecastservice.ListPredictorBacktestExportJobsInput, fn func(*forecastservice.ListPredictorBacktestExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListPredictorBacktestExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListPredictorBacktestExportJobsOutput{}
	fnCacher := func(out *forecastservice.ListPredictorBacktestExportJobsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListPredictorBacktestExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListPredictorBacktestExportJobsWithContext(ctx context.Context, input *forecastservice.ListPredictorBacktestExportJobsInput, opts ...request.Option) (*forecastservice.ListPredictorBacktestExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListPredictorBacktestExportJobsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListPredictorBacktestExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListPredictors(input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListPredictorsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListPredictors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListPredictorsPages(input *forecastservice.ListPredictorsInput, fn func(*forecastservice.ListPredictorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListPredictorsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListPredictorsOutput{}
	fnCacher := func(out *forecastservice.ListPredictorsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListPredictorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListPredictorsPagesWithContext(ctx context.Context, input *forecastservice.ListPredictorsInput, fn func(*forecastservice.ListPredictorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListPredictorsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListPredictorsOutput{}
	fnCacher := func(out *forecastservice.ListPredictorsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListPredictorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListPredictorsWithContext(ctx context.Context, input *forecastservice.ListPredictorsInput, opts ...request.Option) (*forecastservice.ListPredictorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListPredictorsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListPredictorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListTagsForResource(input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListTagsForResourceWithContext(ctx context.Context, input *forecastservice.ListTagsForResourceInput, opts ...request.Option) (*forecastservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfAnalyses(input *forecastservice.ListWhatIfAnalysesInput) (*forecastservice.ListWhatIfAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfAnalysesOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfAnalysesPages(input *forecastservice.ListWhatIfAnalysesInput, fn func(*forecastservice.ListWhatIfAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfAnalysesOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfAnalysesOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfAnalysesPagesWithContext(ctx context.Context, input *forecastservice.ListWhatIfAnalysesInput, fn func(*forecastservice.ListWhatIfAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfAnalysesOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfAnalysesOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfAnalysesWithContext(ctx context.Context, input *forecastservice.ListWhatIfAnalysesInput, opts ...request.Option) (*forecastservice.ListWhatIfAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfAnalysesOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfForecastExports(input *forecastservice.ListWhatIfForecastExportsInput) (*forecastservice.ListWhatIfForecastExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfForecastExportsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfForecastExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfForecastExportsPages(input *forecastservice.ListWhatIfForecastExportsInput, fn func(*forecastservice.ListWhatIfForecastExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfForecastExportsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfForecastExportsOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfForecastExportsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfForecastExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfForecastExportsPagesWithContext(ctx context.Context, input *forecastservice.ListWhatIfForecastExportsInput, fn func(*forecastservice.ListWhatIfForecastExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfForecastExportsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfForecastExportsOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfForecastExportsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfForecastExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfForecastExportsWithContext(ctx context.Context, input *forecastservice.ListWhatIfForecastExportsInput, opts ...request.Option) (*forecastservice.ListWhatIfForecastExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfForecastExportsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfForecastExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfForecasts(input *forecastservice.ListWhatIfForecastsInput) (*forecastservice.ListWhatIfForecastsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfForecastsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfForecasts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastService) ListWhatIfForecastsPages(input *forecastservice.ListWhatIfForecastsInput, fn func(*forecastservice.ListWhatIfForecastsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfForecastsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfForecastsOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfForecastsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfForecastsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfForecastsPagesWithContext(ctx context.Context, input *forecastservice.ListWhatIfForecastsInput, fn func(*forecastservice.ListWhatIfForecastsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*forecastservice.ListWhatIfForecastsOutput), false)
		return nil
	}
	cachable := true
	output := &forecastservice.ListWhatIfForecastsOutput{}
	fnCacher := func(out *forecastservice.ListWhatIfForecastsOutput, more bool) bool {
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
	if err := c.ForecastServiceAPI.ListWhatIfForecastsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ForecastService) ListWhatIfForecastsWithContext(ctx context.Context, input *forecastservice.ListWhatIfForecastsInput, opts ...request.Option) (*forecastservice.ListWhatIfForecastsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastservice.ListWhatIfForecastsOutput), nil
	}
	out, err := c.ForecastServiceAPI.ListWhatIfForecastsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

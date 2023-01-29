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
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ForecastQueryService struct {
	forecastqueryserviceiface.ForecastQueryServiceAPI
	cache *cache.Cache
}

func NewForecastQueryService(forecastqueryserviceapi forecastqueryserviceiface.ForecastQueryServiceAPI) *ForecastQueryService {
	return &ForecastQueryService{
		ForecastQueryServiceAPI: forecastqueryserviceapi,
		cache:                   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ForecastQueryService) QueryForecast(input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastqueryservice.QueryForecastOutput), nil
	}
	out, err := c.ForecastQueryServiceAPI.QueryForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastQueryService) QueryForecastWithContext(ctx context.Context, input *forecastqueryservice.QueryForecastInput, opts ...request.Option) (*forecastqueryservice.QueryForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastqueryservice.QueryForecastOutput), nil
	}
	out, err := c.ForecastQueryServiceAPI.QueryForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastQueryService) QueryWhatIfForecast(input *forecastqueryservice.QueryWhatIfForecastInput) (*forecastqueryservice.QueryWhatIfForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastqueryservice.QueryWhatIfForecastOutput), nil
	}
	out, err := c.ForecastQueryServiceAPI.QueryWhatIfForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ForecastQueryService) QueryWhatIfForecastWithContext(ctx context.Context, input *forecastqueryservice.QueryWhatIfForecastInput, opts ...request.Option) (*forecastqueryservice.QueryWhatIfForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*forecastqueryservice.QueryWhatIfForecastOutput), nil
	}
	out, err := c.ForecastQueryServiceAPI.QueryWhatIfForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

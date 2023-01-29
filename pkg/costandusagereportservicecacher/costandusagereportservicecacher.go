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
package costandusagereportservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CostandUsageReportService struct {
	costandusagereportserviceiface.CostandUsageReportServiceAPI
	cache *cache.Cache
}

func New(costandusagereportserviceapi costandusagereportserviceiface.CostandUsageReportServiceAPI) *CostandUsageReportService {
	return &CostandUsageReportService{
		CostandUsageReportServiceAPI: costandusagereportserviceapi,
		cache:                        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CostandUsageReportService) DescribeReportDefinitions(input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costandusagereportservice.DescribeReportDefinitionsOutput), nil
	}
	out, err := c.CostandUsageReportServiceAPI.DescribeReportDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostandUsageReportService) DescribeReportDefinitionsPages(input *costandusagereportservice.DescribeReportDefinitionsInput, fn func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costandusagereportservice.DescribeReportDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &costandusagereportservice.DescribeReportDefinitionsOutput{}
	fnCacher := func(out *costandusagereportservice.DescribeReportDefinitionsOutput, more bool) bool {
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
	if err := c.CostandUsageReportServiceAPI.DescribeReportDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostandUsageReportService) DescribeReportDefinitionsPagesWithContext(ctx context.Context, input *costandusagereportservice.DescribeReportDefinitionsInput, fn func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costandusagereportservice.DescribeReportDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &costandusagereportservice.DescribeReportDefinitionsOutput{}
	fnCacher := func(out *costandusagereportservice.DescribeReportDefinitionsOutput, more bool) bool {
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
	if err := c.CostandUsageReportServiceAPI.DescribeReportDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostandUsageReportService) DescribeReportDefinitionsWithContext(ctx context.Context, input *costandusagereportservice.DescribeReportDefinitionsInput, opts ...request.Option) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costandusagereportservice.DescribeReportDefinitionsOutput), nil
	}
	out, err := c.CostandUsageReportServiceAPI.DescribeReportDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

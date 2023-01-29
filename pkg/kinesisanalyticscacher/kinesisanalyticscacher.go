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
package kinesisanalyticscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisAnalytics struct {
	kinesisanalyticsiface.KinesisAnalyticsAPI
	cache *cache.Cache
}

func New(kinesisanalyticsapi kinesisanalyticsiface.KinesisAnalyticsAPI) *KinesisAnalytics {
	return &KinesisAnalytics{
		KinesisAnalyticsAPI: kinesisanalyticsapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisAnalytics) DescribeApplication(input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.DescribeApplicationOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.DescribeApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalytics) DescribeApplicationWithContext(ctx context.Context, input *kinesisanalytics.DescribeApplicationInput, opts ...request.Option) (*kinesisanalytics.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.DescribeApplicationOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.DescribeApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalytics) ListApplications(input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.ListApplicationsOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalytics) ListApplicationsWithContext(ctx context.Context, input *kinesisanalytics.ListApplicationsInput, opts ...request.Option) (*kinesisanalytics.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.ListApplicationsOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalytics) ListTagsForResource(input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisAnalytics) ListTagsForResourceWithContext(ctx context.Context, input *kinesisanalytics.ListTagsForResourceInput, opts ...request.Option) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisanalytics.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisAnalyticsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PersonalizeRuntime struct {
	personalizeruntimeiface.PersonalizeRuntimeAPI
	cache *cache.Cache
}

func NewPersonalizeRuntime(personalizeruntimeapi personalizeruntimeiface.PersonalizeRuntimeAPI) *PersonalizeRuntime {
	return &PersonalizeRuntime{
		PersonalizeRuntimeAPI: personalizeruntimeapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PersonalizeRuntime) GetPersonalizedRanking(input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalizeruntime.GetPersonalizedRankingOutput), nil
	}
	out, err := c.PersonalizeRuntimeAPI.GetPersonalizedRanking(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PersonalizeRuntime) GetPersonalizedRankingWithContext(ctx context.Context, input *personalizeruntime.GetPersonalizedRankingInput, opts ...request.Option) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalizeruntime.GetPersonalizedRankingOutput), nil
	}
	out, err := c.PersonalizeRuntimeAPI.GetPersonalizedRankingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PersonalizeRuntime) GetRecommendations(input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalizeruntime.GetRecommendationsOutput), nil
	}
	out, err := c.PersonalizeRuntimeAPI.GetRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PersonalizeRuntime) GetRecommendationsWithContext(ctx context.Context, input *personalizeruntime.GetRecommendationsInput, opts ...request.Option) (*personalizeruntime.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalizeruntime.GetRecommendationsOutput), nil
	}
	out, err := c.PersonalizeRuntimeAPI.GetRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

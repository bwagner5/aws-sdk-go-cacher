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
	"github.com/aws/aws-sdk-go/service/sagemakergeospatial"
	"github.com/aws/aws-sdk-go/service/sagemakergeospatial/sagemakergeospatialiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SageMakerGeospatial struct {
	sagemakergeospatialiface.SageMakerGeospatialAPI
	cache *cache.Cache
}

func NewSageMakerGeospatial(sagemakergeospatialapi sagemakergeospatialiface.SageMakerGeospatialAPI) *SageMakerGeospatial {
	return &SageMakerGeospatial{
		SageMakerGeospatialAPI: sagemakergeospatialapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SageMakerGeospatial) GetEarthObservationJob(input *sagemakergeospatial.GetEarthObservationJobInput) (*sagemakergeospatial.GetEarthObservationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetEarthObservationJobOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetEarthObservationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetEarthObservationJobWithContext(ctx context.Context, input *sagemakergeospatial.GetEarthObservationJobInput, opts ...request.Option) (*sagemakergeospatial.GetEarthObservationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetEarthObservationJobOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetEarthObservationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetRasterDataCollection(input *sagemakergeospatial.GetRasterDataCollectionInput) (*sagemakergeospatial.GetRasterDataCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetRasterDataCollectionOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetRasterDataCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetRasterDataCollectionWithContext(ctx context.Context, input *sagemakergeospatial.GetRasterDataCollectionInput, opts ...request.Option) (*sagemakergeospatial.GetRasterDataCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetRasterDataCollectionOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetRasterDataCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetTile(input *sagemakergeospatial.GetTileInput) (*sagemakergeospatial.GetTileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetTileOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetTile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetTileWithContext(ctx context.Context, input *sagemakergeospatial.GetTileInput, opts ...request.Option) (*sagemakergeospatial.GetTileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetTileOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetTileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetVectorEnrichmentJob(input *sagemakergeospatial.GetVectorEnrichmentJobInput) (*sagemakergeospatial.GetVectorEnrichmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetVectorEnrichmentJobOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetVectorEnrichmentJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) GetVectorEnrichmentJobWithContext(ctx context.Context, input *sagemakergeospatial.GetVectorEnrichmentJobInput, opts ...request.Option) (*sagemakergeospatial.GetVectorEnrichmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.GetVectorEnrichmentJobOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.GetVectorEnrichmentJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListEarthObservationJobs(input *sagemakergeospatial.ListEarthObservationJobsInput) (*sagemakergeospatial.ListEarthObservationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListEarthObservationJobsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListEarthObservationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListEarthObservationJobsPages(input *sagemakergeospatial.ListEarthObservationJobsInput, fn func(*sagemakergeospatial.ListEarthObservationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListEarthObservationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListEarthObservationJobsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListEarthObservationJobsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListEarthObservationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListEarthObservationJobsPagesWithContext(ctx context.Context, input *sagemakergeospatial.ListEarthObservationJobsInput, fn func(*sagemakergeospatial.ListEarthObservationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListEarthObservationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListEarthObservationJobsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListEarthObservationJobsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListEarthObservationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListEarthObservationJobsWithContext(ctx context.Context, input *sagemakergeospatial.ListEarthObservationJobsInput, opts ...request.Option) (*sagemakergeospatial.ListEarthObservationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListEarthObservationJobsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListEarthObservationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListRasterDataCollections(input *sagemakergeospatial.ListRasterDataCollectionsInput) (*sagemakergeospatial.ListRasterDataCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListRasterDataCollectionsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListRasterDataCollections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListRasterDataCollectionsPages(input *sagemakergeospatial.ListRasterDataCollectionsInput, fn func(*sagemakergeospatial.ListRasterDataCollectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListRasterDataCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListRasterDataCollectionsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListRasterDataCollectionsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListRasterDataCollectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListRasterDataCollectionsPagesWithContext(ctx context.Context, input *sagemakergeospatial.ListRasterDataCollectionsInput, fn func(*sagemakergeospatial.ListRasterDataCollectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListRasterDataCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListRasterDataCollectionsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListRasterDataCollectionsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListRasterDataCollectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListRasterDataCollectionsWithContext(ctx context.Context, input *sagemakergeospatial.ListRasterDataCollectionsInput, opts ...request.Option) (*sagemakergeospatial.ListRasterDataCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListRasterDataCollectionsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListRasterDataCollectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListTagsForResource(input *sagemakergeospatial.ListTagsForResourceInput) (*sagemakergeospatial.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListTagsForResourceOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListTagsForResourceWithContext(ctx context.Context, input *sagemakergeospatial.ListTagsForResourceInput, opts ...request.Option) (*sagemakergeospatial.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListTagsForResourceOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListVectorEnrichmentJobs(input *sagemakergeospatial.ListVectorEnrichmentJobsInput) (*sagemakergeospatial.ListVectorEnrichmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListVectorEnrichmentJobsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListVectorEnrichmentJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) ListVectorEnrichmentJobsPages(input *sagemakergeospatial.ListVectorEnrichmentJobsInput, fn func(*sagemakergeospatial.ListVectorEnrichmentJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListVectorEnrichmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListVectorEnrichmentJobsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListVectorEnrichmentJobsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListVectorEnrichmentJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListVectorEnrichmentJobsPagesWithContext(ctx context.Context, input *sagemakergeospatial.ListVectorEnrichmentJobsInput, fn func(*sagemakergeospatial.ListVectorEnrichmentJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.ListVectorEnrichmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.ListVectorEnrichmentJobsOutput{}
	fnCacher := func(out *sagemakergeospatial.ListVectorEnrichmentJobsOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.ListVectorEnrichmentJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) ListVectorEnrichmentJobsWithContext(ctx context.Context, input *sagemakergeospatial.ListVectorEnrichmentJobsInput, opts ...request.Option) (*sagemakergeospatial.ListVectorEnrichmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.ListVectorEnrichmentJobsOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.ListVectorEnrichmentJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) SearchRasterDataCollection(input *sagemakergeospatial.SearchRasterDataCollectionInput) (*sagemakergeospatial.SearchRasterDataCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.SearchRasterDataCollectionOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.SearchRasterDataCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerGeospatial) SearchRasterDataCollectionPages(input *sagemakergeospatial.SearchRasterDataCollectionInput, fn func(*sagemakergeospatial.SearchRasterDataCollectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.SearchRasterDataCollectionOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.SearchRasterDataCollectionOutput{}
	fnCacher := func(out *sagemakergeospatial.SearchRasterDataCollectionOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.SearchRasterDataCollectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) SearchRasterDataCollectionPagesWithContext(ctx context.Context, input *sagemakergeospatial.SearchRasterDataCollectionInput, fn func(*sagemakergeospatial.SearchRasterDataCollectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemakergeospatial.SearchRasterDataCollectionOutput), false)
		return nil
	}
	cachable := true
	output := &sagemakergeospatial.SearchRasterDataCollectionOutput{}
	fnCacher := func(out *sagemakergeospatial.SearchRasterDataCollectionOutput, more bool) bool {
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
	if err := c.SageMakerGeospatialAPI.SearchRasterDataCollectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMakerGeospatial) SearchRasterDataCollectionWithContext(ctx context.Context, input *sagemakergeospatial.SearchRasterDataCollectionInput, opts ...request.Option) (*sagemakergeospatial.SearchRasterDataCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakergeospatial.SearchRasterDataCollectionOutput), nil
	}
	out, err := c.SageMakerGeospatialAPI.SearchRasterDataCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

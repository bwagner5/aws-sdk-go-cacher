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
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Rekognition struct {
	rekognitioniface.RekognitionAPI
	cache *cache.Cache
}

func NewRekognition(rekognitionapi rekognitioniface.RekognitionAPI) *Rekognition {
	return &Rekognition{
		RekognitionAPI: rekognitionapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Rekognition) DescribeCollection(input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeCollectionOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeCollectionWithContext(ctx context.Context, input *rekognition.DescribeCollectionInput, opts ...request.Option) (*rekognition.DescribeCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeCollectionOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeDataset(input *rekognition.DescribeDatasetInput) (*rekognition.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeDatasetOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeDatasetWithContext(ctx context.Context, input *rekognition.DescribeDatasetInput, opts ...request.Option) (*rekognition.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeDatasetOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeProjectVersions(input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeProjectVersionsOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeProjectVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeProjectVersionsPages(input *rekognition.DescribeProjectVersionsInput, fn func(*rekognition.DescribeProjectVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.DescribeProjectVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.DescribeProjectVersionsOutput{}
	fnCacher := func(out *rekognition.DescribeProjectVersionsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.DescribeProjectVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) DescribeProjectVersionsPagesWithContext(ctx context.Context, input *rekognition.DescribeProjectVersionsInput, fn func(*rekognition.DescribeProjectVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.DescribeProjectVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.DescribeProjectVersionsOutput{}
	fnCacher := func(out *rekognition.DescribeProjectVersionsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.DescribeProjectVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) DescribeProjectVersionsWithContext(ctx context.Context, input *rekognition.DescribeProjectVersionsInput, opts ...request.Option) (*rekognition.DescribeProjectVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeProjectVersionsOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeProjectVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeProjects(input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeProjectsOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeProjectsPages(input *rekognition.DescribeProjectsInput, fn func(*rekognition.DescribeProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.DescribeProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.DescribeProjectsOutput{}
	fnCacher := func(out *rekognition.DescribeProjectsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.DescribeProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) DescribeProjectsPagesWithContext(ctx context.Context, input *rekognition.DescribeProjectsInput, fn func(*rekognition.DescribeProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.DescribeProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.DescribeProjectsOutput{}
	fnCacher := func(out *rekognition.DescribeProjectsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.DescribeProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) DescribeProjectsWithContext(ctx context.Context, input *rekognition.DescribeProjectsInput, opts ...request.Option) (*rekognition.DescribeProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeProjectsOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeStreamProcessor(input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeStreamProcessorOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeStreamProcessor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) DescribeStreamProcessorWithContext(ctx context.Context, input *rekognition.DescribeStreamProcessorInput, opts ...request.Option) (*rekognition.DescribeStreamProcessorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.DescribeStreamProcessorOutput), nil
	}
	out, err := c.RekognitionAPI.DescribeStreamProcessorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetCelebrityInfo(input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetCelebrityInfoOutput), nil
	}
	out, err := c.RekognitionAPI.GetCelebrityInfo(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetCelebrityInfoWithContext(ctx context.Context, input *rekognition.GetCelebrityInfoInput, opts ...request.Option) (*rekognition.GetCelebrityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetCelebrityInfoOutput), nil
	}
	out, err := c.RekognitionAPI.GetCelebrityInfoWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetCelebrityRecognition(input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetCelebrityRecognitionOutput), nil
	}
	out, err := c.RekognitionAPI.GetCelebrityRecognition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetCelebrityRecognitionPages(input *rekognition.GetCelebrityRecognitionInput, fn func(*rekognition.GetCelebrityRecognitionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetCelebrityRecognitionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetCelebrityRecognitionOutput{}
	fnCacher := func(out *rekognition.GetCelebrityRecognitionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetCelebrityRecognitionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetCelebrityRecognitionPagesWithContext(ctx context.Context, input *rekognition.GetCelebrityRecognitionInput, fn func(*rekognition.GetCelebrityRecognitionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetCelebrityRecognitionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetCelebrityRecognitionOutput{}
	fnCacher := func(out *rekognition.GetCelebrityRecognitionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetCelebrityRecognitionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetCelebrityRecognitionWithContext(ctx context.Context, input *rekognition.GetCelebrityRecognitionInput, opts ...request.Option) (*rekognition.GetCelebrityRecognitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetCelebrityRecognitionOutput), nil
	}
	out, err := c.RekognitionAPI.GetCelebrityRecognitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetContentModeration(input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetContentModerationOutput), nil
	}
	out, err := c.RekognitionAPI.GetContentModeration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetContentModerationPages(input *rekognition.GetContentModerationInput, fn func(*rekognition.GetContentModerationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetContentModerationOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetContentModerationOutput{}
	fnCacher := func(out *rekognition.GetContentModerationOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetContentModerationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetContentModerationPagesWithContext(ctx context.Context, input *rekognition.GetContentModerationInput, fn func(*rekognition.GetContentModerationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetContentModerationOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetContentModerationOutput{}
	fnCacher := func(out *rekognition.GetContentModerationOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetContentModerationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetContentModerationWithContext(ctx context.Context, input *rekognition.GetContentModerationInput, opts ...request.Option) (*rekognition.GetContentModerationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetContentModerationOutput), nil
	}
	out, err := c.RekognitionAPI.GetContentModerationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetFaceDetection(input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetFaceDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetFaceDetection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetFaceDetectionPages(input *rekognition.GetFaceDetectionInput, fn func(*rekognition.GetFaceDetectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetFaceDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetFaceDetectionOutput{}
	fnCacher := func(out *rekognition.GetFaceDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetFaceDetectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetFaceDetectionPagesWithContext(ctx context.Context, input *rekognition.GetFaceDetectionInput, fn func(*rekognition.GetFaceDetectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetFaceDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetFaceDetectionOutput{}
	fnCacher := func(out *rekognition.GetFaceDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetFaceDetectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetFaceDetectionWithContext(ctx context.Context, input *rekognition.GetFaceDetectionInput, opts ...request.Option) (*rekognition.GetFaceDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetFaceDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetFaceDetectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetFaceSearch(input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetFaceSearchOutput), nil
	}
	out, err := c.RekognitionAPI.GetFaceSearch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetFaceSearchPages(input *rekognition.GetFaceSearchInput, fn func(*rekognition.GetFaceSearchOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetFaceSearchOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetFaceSearchOutput{}
	fnCacher := func(out *rekognition.GetFaceSearchOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetFaceSearchPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetFaceSearchPagesWithContext(ctx context.Context, input *rekognition.GetFaceSearchInput, fn func(*rekognition.GetFaceSearchOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetFaceSearchOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetFaceSearchOutput{}
	fnCacher := func(out *rekognition.GetFaceSearchOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetFaceSearchPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetFaceSearchWithContext(ctx context.Context, input *rekognition.GetFaceSearchInput, opts ...request.Option) (*rekognition.GetFaceSearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetFaceSearchOutput), nil
	}
	out, err := c.RekognitionAPI.GetFaceSearchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetLabelDetection(input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetLabelDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetLabelDetection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetLabelDetectionPages(input *rekognition.GetLabelDetectionInput, fn func(*rekognition.GetLabelDetectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetLabelDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetLabelDetectionOutput{}
	fnCacher := func(out *rekognition.GetLabelDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetLabelDetectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetLabelDetectionPagesWithContext(ctx context.Context, input *rekognition.GetLabelDetectionInput, fn func(*rekognition.GetLabelDetectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetLabelDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetLabelDetectionOutput{}
	fnCacher := func(out *rekognition.GetLabelDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetLabelDetectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetLabelDetectionWithContext(ctx context.Context, input *rekognition.GetLabelDetectionInput, opts ...request.Option) (*rekognition.GetLabelDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetLabelDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetLabelDetectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetPersonTracking(input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetPersonTrackingOutput), nil
	}
	out, err := c.RekognitionAPI.GetPersonTracking(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetPersonTrackingPages(input *rekognition.GetPersonTrackingInput, fn func(*rekognition.GetPersonTrackingOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetPersonTrackingOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetPersonTrackingOutput{}
	fnCacher := func(out *rekognition.GetPersonTrackingOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetPersonTrackingPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetPersonTrackingPagesWithContext(ctx context.Context, input *rekognition.GetPersonTrackingInput, fn func(*rekognition.GetPersonTrackingOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetPersonTrackingOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetPersonTrackingOutput{}
	fnCacher := func(out *rekognition.GetPersonTrackingOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetPersonTrackingPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetPersonTrackingWithContext(ctx context.Context, input *rekognition.GetPersonTrackingInput, opts ...request.Option) (*rekognition.GetPersonTrackingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetPersonTrackingOutput), nil
	}
	out, err := c.RekognitionAPI.GetPersonTrackingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetSegmentDetection(input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetSegmentDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetSegmentDetection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetSegmentDetectionPages(input *rekognition.GetSegmentDetectionInput, fn func(*rekognition.GetSegmentDetectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetSegmentDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetSegmentDetectionOutput{}
	fnCacher := func(out *rekognition.GetSegmentDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetSegmentDetectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetSegmentDetectionPagesWithContext(ctx context.Context, input *rekognition.GetSegmentDetectionInput, fn func(*rekognition.GetSegmentDetectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetSegmentDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetSegmentDetectionOutput{}
	fnCacher := func(out *rekognition.GetSegmentDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetSegmentDetectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetSegmentDetectionWithContext(ctx context.Context, input *rekognition.GetSegmentDetectionInput, opts ...request.Option) (*rekognition.GetSegmentDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetSegmentDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetSegmentDetectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetTextDetection(input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetTextDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetTextDetection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) GetTextDetectionPages(input *rekognition.GetTextDetectionInput, fn func(*rekognition.GetTextDetectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetTextDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetTextDetectionOutput{}
	fnCacher := func(out *rekognition.GetTextDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetTextDetectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetTextDetectionPagesWithContext(ctx context.Context, input *rekognition.GetTextDetectionInput, fn func(*rekognition.GetTextDetectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.GetTextDetectionOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.GetTextDetectionOutput{}
	fnCacher := func(out *rekognition.GetTextDetectionOutput, more bool) bool {
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
	if err := c.RekognitionAPI.GetTextDetectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) GetTextDetectionWithContext(ctx context.Context, input *rekognition.GetTextDetectionInput, opts ...request.Option) (*rekognition.GetTextDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.GetTextDetectionOutput), nil
	}
	out, err := c.RekognitionAPI.GetTextDetectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListCollections(input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListCollectionsOutput), nil
	}
	out, err := c.RekognitionAPI.ListCollections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListCollectionsPages(input *rekognition.ListCollectionsInput, fn func(*rekognition.ListCollectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListCollectionsOutput{}
	fnCacher := func(out *rekognition.ListCollectionsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListCollectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListCollectionsPagesWithContext(ctx context.Context, input *rekognition.ListCollectionsInput, fn func(*rekognition.ListCollectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListCollectionsOutput{}
	fnCacher := func(out *rekognition.ListCollectionsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListCollectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListCollectionsWithContext(ctx context.Context, input *rekognition.ListCollectionsInput, opts ...request.Option) (*rekognition.ListCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListCollectionsOutput), nil
	}
	out, err := c.RekognitionAPI.ListCollectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListDatasetEntries(input *rekognition.ListDatasetEntriesInput) (*rekognition.ListDatasetEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListDatasetEntriesOutput), nil
	}
	out, err := c.RekognitionAPI.ListDatasetEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListDatasetEntriesPages(input *rekognition.ListDatasetEntriesInput, fn func(*rekognition.ListDatasetEntriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListDatasetEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListDatasetEntriesOutput{}
	fnCacher := func(out *rekognition.ListDatasetEntriesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListDatasetEntriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListDatasetEntriesPagesWithContext(ctx context.Context, input *rekognition.ListDatasetEntriesInput, fn func(*rekognition.ListDatasetEntriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListDatasetEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListDatasetEntriesOutput{}
	fnCacher := func(out *rekognition.ListDatasetEntriesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListDatasetEntriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListDatasetEntriesWithContext(ctx context.Context, input *rekognition.ListDatasetEntriesInput, opts ...request.Option) (*rekognition.ListDatasetEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListDatasetEntriesOutput), nil
	}
	out, err := c.RekognitionAPI.ListDatasetEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListDatasetLabels(input *rekognition.ListDatasetLabelsInput) (*rekognition.ListDatasetLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListDatasetLabelsOutput), nil
	}
	out, err := c.RekognitionAPI.ListDatasetLabels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListDatasetLabelsPages(input *rekognition.ListDatasetLabelsInput, fn func(*rekognition.ListDatasetLabelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListDatasetLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListDatasetLabelsOutput{}
	fnCacher := func(out *rekognition.ListDatasetLabelsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListDatasetLabelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListDatasetLabelsPagesWithContext(ctx context.Context, input *rekognition.ListDatasetLabelsInput, fn func(*rekognition.ListDatasetLabelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListDatasetLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListDatasetLabelsOutput{}
	fnCacher := func(out *rekognition.ListDatasetLabelsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListDatasetLabelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListDatasetLabelsWithContext(ctx context.Context, input *rekognition.ListDatasetLabelsInput, opts ...request.Option) (*rekognition.ListDatasetLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListDatasetLabelsOutput), nil
	}
	out, err := c.RekognitionAPI.ListDatasetLabelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListFaces(input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListFacesOutput), nil
	}
	out, err := c.RekognitionAPI.ListFaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListFacesPages(input *rekognition.ListFacesInput, fn func(*rekognition.ListFacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListFacesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListFacesOutput{}
	fnCacher := func(out *rekognition.ListFacesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListFacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListFacesPagesWithContext(ctx context.Context, input *rekognition.ListFacesInput, fn func(*rekognition.ListFacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListFacesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListFacesOutput{}
	fnCacher := func(out *rekognition.ListFacesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListFacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListFacesWithContext(ctx context.Context, input *rekognition.ListFacesInput, opts ...request.Option) (*rekognition.ListFacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListFacesOutput), nil
	}
	out, err := c.RekognitionAPI.ListFacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListProjectPolicies(input *rekognition.ListProjectPoliciesInput) (*rekognition.ListProjectPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListProjectPoliciesOutput), nil
	}
	out, err := c.RekognitionAPI.ListProjectPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListProjectPoliciesPages(input *rekognition.ListProjectPoliciesInput, fn func(*rekognition.ListProjectPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListProjectPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListProjectPoliciesOutput{}
	fnCacher := func(out *rekognition.ListProjectPoliciesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListProjectPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListProjectPoliciesPagesWithContext(ctx context.Context, input *rekognition.ListProjectPoliciesInput, fn func(*rekognition.ListProjectPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListProjectPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListProjectPoliciesOutput{}
	fnCacher := func(out *rekognition.ListProjectPoliciesOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListProjectPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListProjectPoliciesWithContext(ctx context.Context, input *rekognition.ListProjectPoliciesInput, opts ...request.Option) (*rekognition.ListProjectPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListProjectPoliciesOutput), nil
	}
	out, err := c.RekognitionAPI.ListProjectPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListStreamProcessors(input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListStreamProcessorsOutput), nil
	}
	out, err := c.RekognitionAPI.ListStreamProcessors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListStreamProcessorsPages(input *rekognition.ListStreamProcessorsInput, fn func(*rekognition.ListStreamProcessorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListStreamProcessorsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListStreamProcessorsOutput{}
	fnCacher := func(out *rekognition.ListStreamProcessorsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListStreamProcessorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListStreamProcessorsPagesWithContext(ctx context.Context, input *rekognition.ListStreamProcessorsInput, fn func(*rekognition.ListStreamProcessorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rekognition.ListStreamProcessorsOutput), false)
		return nil
	}
	cachable := true
	output := &rekognition.ListStreamProcessorsOutput{}
	fnCacher := func(out *rekognition.ListStreamProcessorsOutput, more bool) bool {
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
	if err := c.RekognitionAPI.ListStreamProcessorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Rekognition) ListStreamProcessorsWithContext(ctx context.Context, input *rekognition.ListStreamProcessorsInput, opts ...request.Option) (*rekognition.ListStreamProcessorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListStreamProcessorsOutput), nil
	}
	out, err := c.RekognitionAPI.ListStreamProcessorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListTagsForResource(input *rekognition.ListTagsForResourceInput) (*rekognition.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListTagsForResourceOutput), nil
	}
	out, err := c.RekognitionAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) ListTagsForResourceWithContext(ctx context.Context, input *rekognition.ListTagsForResourceInput, opts ...request.Option) (*rekognition.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.ListTagsForResourceOutput), nil
	}
	out, err := c.RekognitionAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) SearchFaces(input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.SearchFacesOutput), nil
	}
	out, err := c.RekognitionAPI.SearchFaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) SearchFacesByImage(input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.SearchFacesByImageOutput), nil
	}
	out, err := c.RekognitionAPI.SearchFacesByImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) SearchFacesByImageWithContext(ctx context.Context, input *rekognition.SearchFacesByImageInput, opts ...request.Option) (*rekognition.SearchFacesByImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.SearchFacesByImageOutput), nil
	}
	out, err := c.RekognitionAPI.SearchFacesByImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Rekognition) SearchFacesWithContext(ctx context.Context, input *rekognition.SearchFacesInput, opts ...request.Option) (*rekognition.SearchFacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rekognition.SearchFacesOutput), nil
	}
	out, err := c.RekognitionAPI.SearchFacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

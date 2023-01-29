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
	"github.com/aws/aws-sdk-go/service/healthlake"
	"github.com/aws/aws-sdk-go/service/healthlake/healthlakeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type HealthLake struct {
	healthlakeiface.HealthLakeAPI
	cache *cache.Cache
}

func NewHealthLake(healthlakeapi healthlakeiface.HealthLakeAPI) *HealthLake {
	return &HealthLake{
		HealthLakeAPI: healthlakeapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *HealthLake) DescribeFHIRDatastore(input *healthlake.DescribeFHIRDatastoreInput) (*healthlake.DescribeFHIRDatastoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRDatastoreOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRDatastore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) DescribeFHIRDatastoreWithContext(ctx context.Context, input *healthlake.DescribeFHIRDatastoreInput, opts ...request.Option) (*healthlake.DescribeFHIRDatastoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRDatastoreOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRDatastoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) DescribeFHIRExportJob(input *healthlake.DescribeFHIRExportJobInput) (*healthlake.DescribeFHIRExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRExportJobOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) DescribeFHIRExportJobWithContext(ctx context.Context, input *healthlake.DescribeFHIRExportJobInput, opts ...request.Option) (*healthlake.DescribeFHIRExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRExportJobOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) DescribeFHIRImportJob(input *healthlake.DescribeFHIRImportJobInput) (*healthlake.DescribeFHIRImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRImportJobOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) DescribeFHIRImportJobWithContext(ctx context.Context, input *healthlake.DescribeFHIRImportJobInput, opts ...request.Option) (*healthlake.DescribeFHIRImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.DescribeFHIRImportJobOutput), nil
	}
	out, err := c.HealthLakeAPI.DescribeFHIRImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRDatastores(input *healthlake.ListFHIRDatastoresInput) (*healthlake.ListFHIRDatastoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRDatastoresOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRDatastores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRDatastoresPages(input *healthlake.ListFHIRDatastoresInput, fn func(*healthlake.ListFHIRDatastoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRDatastoresOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRDatastoresOutput{}
	fnCacher := func(out *healthlake.ListFHIRDatastoresOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRDatastoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRDatastoresPagesWithContext(ctx context.Context, input *healthlake.ListFHIRDatastoresInput, fn func(*healthlake.ListFHIRDatastoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRDatastoresOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRDatastoresOutput{}
	fnCacher := func(out *healthlake.ListFHIRDatastoresOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRDatastoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRDatastoresWithContext(ctx context.Context, input *healthlake.ListFHIRDatastoresInput, opts ...request.Option) (*healthlake.ListFHIRDatastoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRDatastoresOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRDatastoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRExportJobs(input *healthlake.ListFHIRExportJobsInput) (*healthlake.ListFHIRExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRExportJobsOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRExportJobsPages(input *healthlake.ListFHIRExportJobsInput, fn func(*healthlake.ListFHIRExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRExportJobsOutput{}
	fnCacher := func(out *healthlake.ListFHIRExportJobsOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRExportJobsPagesWithContext(ctx context.Context, input *healthlake.ListFHIRExportJobsInput, fn func(*healthlake.ListFHIRExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRExportJobsOutput{}
	fnCacher := func(out *healthlake.ListFHIRExportJobsOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRExportJobsWithContext(ctx context.Context, input *healthlake.ListFHIRExportJobsInput, opts ...request.Option) (*healthlake.ListFHIRExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRExportJobsOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRImportJobs(input *healthlake.ListFHIRImportJobsInput) (*healthlake.ListFHIRImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRImportJobsOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListFHIRImportJobsPages(input *healthlake.ListFHIRImportJobsInput, fn func(*healthlake.ListFHIRImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRImportJobsOutput{}
	fnCacher := func(out *healthlake.ListFHIRImportJobsOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRImportJobsPagesWithContext(ctx context.Context, input *healthlake.ListFHIRImportJobsInput, fn func(*healthlake.ListFHIRImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*healthlake.ListFHIRImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &healthlake.ListFHIRImportJobsOutput{}
	fnCacher := func(out *healthlake.ListFHIRImportJobsOutput, more bool) bool {
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
	if err := c.HealthLakeAPI.ListFHIRImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *HealthLake) ListFHIRImportJobsWithContext(ctx context.Context, input *healthlake.ListFHIRImportJobsInput, opts ...request.Option) (*healthlake.ListFHIRImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListFHIRImportJobsOutput), nil
	}
	out, err := c.HealthLakeAPI.ListFHIRImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListTagsForResource(input *healthlake.ListTagsForResourceInput) (*healthlake.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListTagsForResourceOutput), nil
	}
	out, err := c.HealthLakeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *HealthLake) ListTagsForResourceWithContext(ctx context.Context, input *healthlake.ListTagsForResourceInput, opts ...request.Option) (*healthlake.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*healthlake.ListTagsForResourceOutput), nil
	}
	out, err := c.HealthLakeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

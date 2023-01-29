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
package amplifybackendcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/amplifybackend"
	"github.com/aws/aws-sdk-go/service/amplifybackend/amplifybackendiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AmplifyBackend struct {
	amplifybackendiface.AmplifyBackendAPI
	cache *cache.Cache
}

func New(amplifybackendapi amplifybackendiface.AmplifyBackendAPI) *AmplifyBackend {
	return &AmplifyBackend{
		AmplifyBackendAPI: amplifybackendapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AmplifyBackend) GetBackend(input *amplifybackend.GetBackendInput) (*amplifybackend.GetBackendOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackend(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAPI(input *amplifybackend.GetBackendAPIInput) (*amplifybackend.GetBackendAPIOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAPIOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAPI(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAPIModels(input *amplifybackend.GetBackendAPIModelsInput) (*amplifybackend.GetBackendAPIModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAPIModelsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAPIModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAPIModelsWithContext(ctx context.Context, input *amplifybackend.GetBackendAPIModelsInput, opts ...request.Option) (*amplifybackend.GetBackendAPIModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAPIModelsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAPIModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAPIWithContext(ctx context.Context, input *amplifybackend.GetBackendAPIInput, opts ...request.Option) (*amplifybackend.GetBackendAPIOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAPIOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAPIWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAuth(input *amplifybackend.GetBackendAuthInput) (*amplifybackend.GetBackendAuthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAuthOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAuth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendAuthWithContext(ctx context.Context, input *amplifybackend.GetBackendAuthInput, opts ...request.Option) (*amplifybackend.GetBackendAuthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendAuthOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendAuthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendJob(input *amplifybackend.GetBackendJobInput) (*amplifybackend.GetBackendJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendJobOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendJobWithContext(ctx context.Context, input *amplifybackend.GetBackendJobInput, opts ...request.Option) (*amplifybackend.GetBackendJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendJobOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendStorage(input *amplifybackend.GetBackendStorageInput) (*amplifybackend.GetBackendStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendStorageOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendStorage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendStorageWithContext(ctx context.Context, input *amplifybackend.GetBackendStorageInput, opts ...request.Option) (*amplifybackend.GetBackendStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendStorageOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendStorageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetBackendWithContext(ctx context.Context, input *amplifybackend.GetBackendInput, opts ...request.Option) (*amplifybackend.GetBackendOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetBackendOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetBackendWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetToken(input *amplifybackend.GetTokenInput) (*amplifybackend.GetTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetTokenOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) GetTokenWithContext(ctx context.Context, input *amplifybackend.GetTokenInput, opts ...request.Option) (*amplifybackend.GetTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.GetTokenOutput), nil
	}
	out, err := c.AmplifyBackendAPI.GetTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) ListBackendJobs(input *amplifybackend.ListBackendJobsInput) (*amplifybackend.ListBackendJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.ListBackendJobsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.ListBackendJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) ListBackendJobsPages(input *amplifybackend.ListBackendJobsInput, fn func(*amplifybackend.ListBackendJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifybackend.ListBackendJobsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifybackend.ListBackendJobsOutput{}
	fnCacher := func(out *amplifybackend.ListBackendJobsOutput, more bool) bool {
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
	if err := c.AmplifyBackendAPI.ListBackendJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyBackend) ListBackendJobsPagesWithContext(ctx context.Context, input *amplifybackend.ListBackendJobsInput, fn func(*amplifybackend.ListBackendJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifybackend.ListBackendJobsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifybackend.ListBackendJobsOutput{}
	fnCacher := func(out *amplifybackend.ListBackendJobsOutput, more bool) bool {
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
	if err := c.AmplifyBackendAPI.ListBackendJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyBackend) ListBackendJobsWithContext(ctx context.Context, input *amplifybackend.ListBackendJobsInput, opts ...request.Option) (*amplifybackend.ListBackendJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.ListBackendJobsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.ListBackendJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) ListS3Buckets(input *amplifybackend.ListS3BucketsInput) (*amplifybackend.ListS3BucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.ListS3BucketsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.ListS3Buckets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyBackend) ListS3BucketsWithContext(ctx context.Context, input *amplifybackend.ListS3BucketsInput, opts ...request.Option) (*amplifybackend.ListS3BucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifybackend.ListS3BucketsOutput), nil
	}
	out, err := c.AmplifyBackendAPI.ListS3BucketsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package emrserverlesscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/emrserverless"
	"github.com/aws/aws-sdk-go/service/emrserverless/emrserverlessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EMRServerless struct {
	emrserverlessiface.EMRServerlessAPI
	cache *cache.Cache
}

func New(emrserverlessapi emrserverlessiface.EMRServerlessAPI) *EMRServerless {
	return &EMRServerless{
		EMRServerlessAPI: emrserverlessapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EMRServerless) GetApplication(input *emrserverless.GetApplicationInput) (*emrserverless.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetApplicationOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) GetApplicationWithContext(ctx context.Context, input *emrserverless.GetApplicationInput, opts ...request.Option) (*emrserverless.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetApplicationOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) GetDashboardForJobRun(input *emrserverless.GetDashboardForJobRunInput) (*emrserverless.GetDashboardForJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetDashboardForJobRunOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetDashboardForJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) GetDashboardForJobRunWithContext(ctx context.Context, input *emrserverless.GetDashboardForJobRunInput, opts ...request.Option) (*emrserverless.GetDashboardForJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetDashboardForJobRunOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetDashboardForJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) GetJobRun(input *emrserverless.GetJobRunInput) (*emrserverless.GetJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetJobRunOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) GetJobRunWithContext(ctx context.Context, input *emrserverless.GetJobRunInput, opts ...request.Option) (*emrserverless.GetJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.GetJobRunOutput), nil
	}
	out, err := c.EMRServerlessAPI.GetJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListApplications(input *emrserverless.ListApplicationsInput) (*emrserverless.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListApplicationsOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListApplicationsPages(input *emrserverless.ListApplicationsInput, fn func(*emrserverless.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrserverless.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &emrserverless.ListApplicationsOutput{}
	fnCacher := func(out *emrserverless.ListApplicationsOutput, more bool) bool {
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
	if err := c.EMRServerlessAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRServerless) ListApplicationsPagesWithContext(ctx context.Context, input *emrserverless.ListApplicationsInput, fn func(*emrserverless.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrserverless.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &emrserverless.ListApplicationsOutput{}
	fnCacher := func(out *emrserverless.ListApplicationsOutput, more bool) bool {
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
	if err := c.EMRServerlessAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRServerless) ListApplicationsWithContext(ctx context.Context, input *emrserverless.ListApplicationsInput, opts ...request.Option) (*emrserverless.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListApplicationsOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListJobRuns(input *emrserverless.ListJobRunsInput) (*emrserverless.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListJobRunsOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListJobRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListJobRunsPages(input *emrserverless.ListJobRunsInput, fn func(*emrserverless.ListJobRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrserverless.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &emrserverless.ListJobRunsOutput{}
	fnCacher := func(out *emrserverless.ListJobRunsOutput, more bool) bool {
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
	if err := c.EMRServerlessAPI.ListJobRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRServerless) ListJobRunsPagesWithContext(ctx context.Context, input *emrserverless.ListJobRunsInput, fn func(*emrserverless.ListJobRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrserverless.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &emrserverless.ListJobRunsOutput{}
	fnCacher := func(out *emrserverless.ListJobRunsOutput, more bool) bool {
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
	if err := c.EMRServerlessAPI.ListJobRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRServerless) ListJobRunsWithContext(ctx context.Context, input *emrserverless.ListJobRunsInput, opts ...request.Option) (*emrserverless.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListJobRunsOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListJobRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListTagsForResource(input *emrserverless.ListTagsForResourceInput) (*emrserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRServerless) ListTagsForResourceWithContext(ctx context.Context, input *emrserverless.ListTagsForResourceInput, opts ...request.Option) (*emrserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.EMRServerlessAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

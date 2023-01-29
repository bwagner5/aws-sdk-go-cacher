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
package serverlessapplicationrepositorycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ServerlessApplicationRepository struct {
	serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI
	cache *cache.Cache
}

func New(serverlessapplicationrepositoryapi serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI) *ServerlessApplicationRepository {
	return &ServerlessApplicationRepository{
		ServerlessApplicationRepositoryAPI: serverlessapplicationrepositoryapi,
		cache:                              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ServerlessApplicationRepository) GetApplication(input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetApplicationOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) GetApplicationPolicy(input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetApplicationPolicyOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetApplicationPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) GetApplicationPolicyWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetApplicationPolicyOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetApplicationPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) GetApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetApplicationOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) GetCloudFormationTemplate(input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetCloudFormationTemplateOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetCloudFormationTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) GetCloudFormationTemplateWithContext(ctx context.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput, opts ...request.Option) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.GetCloudFormationTemplateOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.GetCloudFormationTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplicationDependencies(input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationDependenciesOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplicationDependencies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplicationDependenciesPages(input *serverlessapplicationrepository.ListApplicationDependenciesInput, fn func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationDependenciesOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationDependenciesOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationDependenciesOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationDependenciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationDependenciesPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, fn func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationDependenciesOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationDependenciesOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationDependenciesOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationDependenciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationDependenciesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationDependenciesOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplicationDependenciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplicationVersions(input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationVersionsOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplicationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplicationVersionsPages(input *serverlessapplicationrepository.ListApplicationVersionsInput, fn func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationVersionsOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationVersionsOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationVersionsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, fn func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationVersionsOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationVersionsOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationVersionsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationVersionsOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplicationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplications(input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationsOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServerlessApplicationRepository) ListApplicationsPages(input *serverlessapplicationrepository.ListApplicationsInput, fn func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationsOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationsOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, fn func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*serverlessapplicationrepository.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &serverlessapplicationrepository.ListApplicationsOutput{}
	fnCacher := func(out *serverlessapplicationrepository.ListApplicationsOutput, more bool) bool {
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
	if err := c.ServerlessApplicationRepositoryAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServerlessApplicationRepository) ListApplicationsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*serverlessapplicationrepository.ListApplicationsOutput), nil
	}
	out, err := c.ServerlessApplicationRepositoryAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

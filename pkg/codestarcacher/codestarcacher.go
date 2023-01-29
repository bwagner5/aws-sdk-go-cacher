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
package codestarcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestar/codestariface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeStar struct {
	codestariface.CodeStarAPI
	cache *cache.Cache
}

func New(codestarapi codestariface.CodeStarAPI) *CodeStar {
	return &CodeStar{
		CodeStarAPI: codestarapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeStar) DescribeProject(input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.DescribeProjectOutput), nil
	}
	out, err := c.CodeStarAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) DescribeProjectWithContext(ctx context.Context, input *codestar.DescribeProjectInput, opts ...request.Option) (*codestar.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.DescribeProjectOutput), nil
	}
	out, err := c.CodeStarAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) DescribeUserProfile(input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.DescribeUserProfileOutput), nil
	}
	out, err := c.CodeStarAPI.DescribeUserProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) DescribeUserProfileWithContext(ctx context.Context, input *codestar.DescribeUserProfileInput, opts ...request.Option) (*codestar.DescribeUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.DescribeUserProfileOutput), nil
	}
	out, err := c.CodeStarAPI.DescribeUserProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListProjects(input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListProjectsOutput), nil
	}
	out, err := c.CodeStarAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListProjectsWithContext(ctx context.Context, input *codestar.ListProjectsInput, opts ...request.Option) (*codestar.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListProjectsOutput), nil
	}
	out, err := c.CodeStarAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListResources(input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListResourcesOutput), nil
	}
	out, err := c.CodeStarAPI.ListResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListResourcesWithContext(ctx context.Context, input *codestar.ListResourcesInput, opts ...request.Option) (*codestar.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListResourcesOutput), nil
	}
	out, err := c.CodeStarAPI.ListResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListTagsForProject(input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListTagsForProjectOutput), nil
	}
	out, err := c.CodeStarAPI.ListTagsForProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListTagsForProjectWithContext(ctx context.Context, input *codestar.ListTagsForProjectInput, opts ...request.Option) (*codestar.ListTagsForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListTagsForProjectOutput), nil
	}
	out, err := c.CodeStarAPI.ListTagsForProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListTeamMembers(input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListTeamMembersOutput), nil
	}
	out, err := c.CodeStarAPI.ListTeamMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListTeamMembersWithContext(ctx context.Context, input *codestar.ListTeamMembersInput, opts ...request.Option) (*codestar.ListTeamMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListTeamMembersOutput), nil
	}
	out, err := c.CodeStarAPI.ListTeamMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListUserProfiles(input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListUserProfilesOutput), nil
	}
	out, err := c.CodeStarAPI.ListUserProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStar) ListUserProfilesWithContext(ctx context.Context, input *codestar.ListUserProfilesInput, opts ...request.Option) (*codestar.ListUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestar.ListUserProfilesOutput), nil
	}
	out, err := c.CodeStarAPI.ListUserProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

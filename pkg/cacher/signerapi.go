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
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Signer struct {
	signeriface.SignerAPI
	cache *cache.Cache
}

func NewSigner(signerapi signeriface.SignerAPI) *Signer {
	return &Signer{
		SignerAPI: signerapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Signer) DescribeSigningJob(input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.DescribeSigningJobOutput), nil
	}
	out, err := c.SignerAPI.DescribeSigningJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) DescribeSigningJobWithContext(ctx context.Context, input *signer.DescribeSigningJobInput, opts ...request.Option) (*signer.DescribeSigningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.DescribeSigningJobOutput), nil
	}
	out, err := c.SignerAPI.DescribeSigningJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) GetSigningPlatform(input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.GetSigningPlatformOutput), nil
	}
	out, err := c.SignerAPI.GetSigningPlatform(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) GetSigningPlatformWithContext(ctx context.Context, input *signer.GetSigningPlatformInput, opts ...request.Option) (*signer.GetSigningPlatformOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.GetSigningPlatformOutput), nil
	}
	out, err := c.SignerAPI.GetSigningPlatformWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) GetSigningProfile(input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.GetSigningProfileOutput), nil
	}
	out, err := c.SignerAPI.GetSigningProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) GetSigningProfileWithContext(ctx context.Context, input *signer.GetSigningProfileInput, opts ...request.Option) (*signer.GetSigningProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.GetSigningProfileOutput), nil
	}
	out, err := c.SignerAPI.GetSigningProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListProfilePermissions(input *signer.ListProfilePermissionsInput) (*signer.ListProfilePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListProfilePermissionsOutput), nil
	}
	out, err := c.SignerAPI.ListProfilePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListProfilePermissionsWithContext(ctx context.Context, input *signer.ListProfilePermissionsInput, opts ...request.Option) (*signer.ListProfilePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListProfilePermissionsOutput), nil
	}
	out, err := c.SignerAPI.ListProfilePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningJobs(input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningJobsOutput), nil
	}
	out, err := c.SignerAPI.ListSigningJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningJobsPages(input *signer.ListSigningJobsInput, fn func(*signer.ListSigningJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningJobsOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningJobsOutput{}
	fnCacher := func(out *signer.ListSigningJobsOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningJobsPagesWithContext(ctx context.Context, input *signer.ListSigningJobsInput, fn func(*signer.ListSigningJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningJobsOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningJobsOutput{}
	fnCacher := func(out *signer.ListSigningJobsOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningJobsWithContext(ctx context.Context, input *signer.ListSigningJobsInput, opts ...request.Option) (*signer.ListSigningJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningJobsOutput), nil
	}
	out, err := c.SignerAPI.ListSigningJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningPlatforms(input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningPlatformsOutput), nil
	}
	out, err := c.SignerAPI.ListSigningPlatforms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningPlatformsPages(input *signer.ListSigningPlatformsInput, fn func(*signer.ListSigningPlatformsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningPlatformsOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningPlatformsOutput{}
	fnCacher := func(out *signer.ListSigningPlatformsOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningPlatformsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningPlatformsPagesWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, fn func(*signer.ListSigningPlatformsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningPlatformsOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningPlatformsOutput{}
	fnCacher := func(out *signer.ListSigningPlatformsOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningPlatformsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningPlatformsWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, opts ...request.Option) (*signer.ListSigningPlatformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningPlatformsOutput), nil
	}
	out, err := c.SignerAPI.ListSigningPlatformsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningProfiles(input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningProfilesOutput), nil
	}
	out, err := c.SignerAPI.ListSigningProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListSigningProfilesPages(input *signer.ListSigningProfilesInput, fn func(*signer.ListSigningProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningProfilesOutput{}
	fnCacher := func(out *signer.ListSigningProfilesOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningProfilesPagesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, fn func(*signer.ListSigningProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*signer.ListSigningProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &signer.ListSigningProfilesOutput{}
	fnCacher := func(out *signer.ListSigningProfilesOutput, more bool) bool {
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
	if err := c.SignerAPI.ListSigningProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Signer) ListSigningProfilesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, opts ...request.Option) (*signer.ListSigningProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListSigningProfilesOutput), nil
	}
	out, err := c.SignerAPI.ListSigningProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListTagsForResource(input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListTagsForResourceOutput), nil
	}
	out, err := c.SignerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Signer) ListTagsForResourceWithContext(ctx context.Context, input *signer.ListTagsForResourceInput, opts ...request.Option) (*signer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*signer.ListTagsForResourceOutput), nil
	}
	out, err := c.SignerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

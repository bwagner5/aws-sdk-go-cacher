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
package rolesanywherecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rolesanywhere"
	"github.com/aws/aws-sdk-go/service/rolesanywhere/rolesanywhereiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RolesAnywhere struct {
	rolesanywhereiface.RolesAnywhereAPI
	cache *cache.Cache
}

func New(rolesanywhereapi rolesanywhereiface.RolesAnywhereAPI) *RolesAnywhere {
	return &RolesAnywhere{
		RolesAnywhereAPI: rolesanywhereapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RolesAnywhere) GetCrl(input *rolesanywhere.GetCrlInput) (*rolesanywhere.GetCrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetCrlOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetCrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetCrlWithContext(ctx context.Context, input *rolesanywhere.GetCrlInput, opts ...request.Option) (*rolesanywhere.GetCrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetCrlOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetCrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetProfile(input *rolesanywhere.GetProfileInput) (*rolesanywhere.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetProfileOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetProfileWithContext(ctx context.Context, input *rolesanywhere.GetProfileInput, opts ...request.Option) (*rolesanywhere.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetProfileOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetSubject(input *rolesanywhere.GetSubjectInput) (*rolesanywhere.GetSubjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetSubjectOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetSubject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetSubjectWithContext(ctx context.Context, input *rolesanywhere.GetSubjectInput, opts ...request.Option) (*rolesanywhere.GetSubjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetSubjectOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetSubjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetTrustAnchor(input *rolesanywhere.GetTrustAnchorInput) (*rolesanywhere.GetTrustAnchorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetTrustAnchorOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetTrustAnchor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) GetTrustAnchorWithContext(ctx context.Context, input *rolesanywhere.GetTrustAnchorInput, opts ...request.Option) (*rolesanywhere.GetTrustAnchorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.GetTrustAnchorOutput), nil
	}
	out, err := c.RolesAnywhereAPI.GetTrustAnchorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListCrls(input *rolesanywhere.ListCrlsInput) (*rolesanywhere.ListCrlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListCrlsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListCrls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListCrlsPages(input *rolesanywhere.ListCrlsInput, fn func(*rolesanywhere.ListCrlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListCrlsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListCrlsOutput{}
	fnCacher := func(out *rolesanywhere.ListCrlsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListCrlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListCrlsPagesWithContext(ctx context.Context, input *rolesanywhere.ListCrlsInput, fn func(*rolesanywhere.ListCrlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListCrlsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListCrlsOutput{}
	fnCacher := func(out *rolesanywhere.ListCrlsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListCrlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListCrlsWithContext(ctx context.Context, input *rolesanywhere.ListCrlsInput, opts ...request.Option) (*rolesanywhere.ListCrlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListCrlsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListCrlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListProfiles(input *rolesanywhere.ListProfilesInput) (*rolesanywhere.ListProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListProfilesOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListProfilesPages(input *rolesanywhere.ListProfilesInput, fn func(*rolesanywhere.ListProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListProfilesOutput{}
	fnCacher := func(out *rolesanywhere.ListProfilesOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListProfilesPagesWithContext(ctx context.Context, input *rolesanywhere.ListProfilesInput, fn func(*rolesanywhere.ListProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListProfilesOutput{}
	fnCacher := func(out *rolesanywhere.ListProfilesOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListProfilesWithContext(ctx context.Context, input *rolesanywhere.ListProfilesInput, opts ...request.Option) (*rolesanywhere.ListProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListProfilesOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListSubjects(input *rolesanywhere.ListSubjectsInput) (*rolesanywhere.ListSubjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListSubjectsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListSubjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListSubjectsPages(input *rolesanywhere.ListSubjectsInput, fn func(*rolesanywhere.ListSubjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListSubjectsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListSubjectsOutput{}
	fnCacher := func(out *rolesanywhere.ListSubjectsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListSubjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListSubjectsPagesWithContext(ctx context.Context, input *rolesanywhere.ListSubjectsInput, fn func(*rolesanywhere.ListSubjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListSubjectsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListSubjectsOutput{}
	fnCacher := func(out *rolesanywhere.ListSubjectsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListSubjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListSubjectsWithContext(ctx context.Context, input *rolesanywhere.ListSubjectsInput, opts ...request.Option) (*rolesanywhere.ListSubjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListSubjectsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListSubjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListTagsForResource(input *rolesanywhere.ListTagsForResourceInput) (*rolesanywhere.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListTagsForResourceOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListTagsForResourceWithContext(ctx context.Context, input *rolesanywhere.ListTagsForResourceInput, opts ...request.Option) (*rolesanywhere.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListTagsForResourceOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListTrustAnchors(input *rolesanywhere.ListTrustAnchorsInput) (*rolesanywhere.ListTrustAnchorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListTrustAnchorsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListTrustAnchors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RolesAnywhere) ListTrustAnchorsPages(input *rolesanywhere.ListTrustAnchorsInput, fn func(*rolesanywhere.ListTrustAnchorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListTrustAnchorsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListTrustAnchorsOutput{}
	fnCacher := func(out *rolesanywhere.ListTrustAnchorsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListTrustAnchorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListTrustAnchorsPagesWithContext(ctx context.Context, input *rolesanywhere.ListTrustAnchorsInput, fn func(*rolesanywhere.ListTrustAnchorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rolesanywhere.ListTrustAnchorsOutput), false)
		return nil
	}
	cachable := true
	output := &rolesanywhere.ListTrustAnchorsOutput{}
	fnCacher := func(out *rolesanywhere.ListTrustAnchorsOutput, more bool) bool {
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
	if err := c.RolesAnywhereAPI.ListTrustAnchorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RolesAnywhere) ListTrustAnchorsWithContext(ctx context.Context, input *rolesanywhere.ListTrustAnchorsInput, opts ...request.Option) (*rolesanywhere.ListTrustAnchorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rolesanywhere.ListTrustAnchorsOutput), nil
	}
	out, err := c.RolesAnywhereAPI.ListTrustAnchorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package redshiftcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshift/redshiftiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Redshift struct {
	redshiftiface.RedshiftAPI
	cache *cache.Cache
}

func New(redshiftapi redshiftiface.RedshiftAPI) *Redshift {
	return &Redshift{
		RedshiftAPI: redshiftapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Redshift) BatchDeleteClusterSnapshots(input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.BatchDeleteClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.BatchDeleteClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) BatchDeleteClusterSnapshotsWithContext(ctx context.Context, input *redshift.BatchDeleteClusterSnapshotsInput, opts ...request.Option) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.BatchDeleteClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.BatchDeleteClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) BatchModifyClusterSnapshots(input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.BatchModifyClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.BatchModifyClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) BatchModifyClusterSnapshotsWithContext(ctx context.Context, input *redshift.BatchModifyClusterSnapshotsInput, opts ...request.Option) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.BatchModifyClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.BatchModifyClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeAccountAttributes(input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeAccountAttributesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeAccountAttributesWithContext(ctx context.Context, input *redshift.DescribeAccountAttributesInput, opts ...request.Option) (*redshift.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeAccountAttributesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeAuthenticationProfiles(input *redshift.DescribeAuthenticationProfilesInput) (*redshift.DescribeAuthenticationProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeAuthenticationProfilesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeAuthenticationProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeAuthenticationProfilesWithContext(ctx context.Context, input *redshift.DescribeAuthenticationProfilesInput, opts ...request.Option) (*redshift.DescribeAuthenticationProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeAuthenticationProfilesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeAuthenticationProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterDbRevisions(input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterDbRevisionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterDbRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterDbRevisionsPages(input *redshift.DescribeClusterDbRevisionsInput, fn func(*redshift.DescribeClusterDbRevisionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterDbRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterDbRevisionsOutput{}
	fnCacher := func(out *redshift.DescribeClusterDbRevisionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterDbRevisionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterDbRevisionsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterDbRevisionsInput, fn func(*redshift.DescribeClusterDbRevisionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterDbRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterDbRevisionsOutput{}
	fnCacher := func(out *redshift.DescribeClusterDbRevisionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterDbRevisionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterDbRevisionsWithContext(ctx context.Context, input *redshift.DescribeClusterDbRevisionsInput, opts ...request.Option) (*redshift.DescribeClusterDbRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterDbRevisionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterDbRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterParameterGroups(input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterParameterGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterParameterGroupsPages(input *redshift.DescribeClusterParameterGroupsInput, fn func(*redshift.DescribeClusterParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterParameterGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterParameterGroupsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterParameterGroupsInput, fn func(*redshift.DescribeClusterParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterParameterGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterParameterGroupsWithContext(ctx context.Context, input *redshift.DescribeClusterParameterGroupsInput, opts ...request.Option) (*redshift.DescribeClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterParameterGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterParameters(input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterParametersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterParametersPages(input *redshift.DescribeClusterParametersInput, fn func(*redshift.DescribeClusterParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterParametersOutput{}
	fnCacher := func(out *redshift.DescribeClusterParametersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterParametersPagesWithContext(ctx context.Context, input *redshift.DescribeClusterParametersInput, fn func(*redshift.DescribeClusterParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterParametersOutput{}
	fnCacher := func(out *redshift.DescribeClusterParametersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterParametersWithContext(ctx context.Context, input *redshift.DescribeClusterParametersInput, opts ...request.Option) (*redshift.DescribeClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterParametersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSecurityGroups(input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSecurityGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSecurityGroupsPages(input *redshift.DescribeClusterSecurityGroupsInput, fn func(*redshift.DescribeClusterSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSecurityGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSecurityGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSecurityGroupsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterSecurityGroupsInput, fn func(*redshift.DescribeClusterSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSecurityGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSecurityGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSecurityGroupsWithContext(ctx context.Context, input *redshift.DescribeClusterSecurityGroupsInput, opts ...request.Option) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSecurityGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSnapshots(input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSnapshotsPages(input *redshift.DescribeClusterSnapshotsInput, fn func(*redshift.DescribeClusterSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSnapshotsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSnapshotsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSnapshotsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterSnapshotsInput, fn func(*redshift.DescribeClusterSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSnapshotsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSnapshotsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSnapshotsWithContext(ctx context.Context, input *redshift.DescribeClusterSnapshotsInput, opts ...request.Option) (*redshift.DescribeClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSnapshotsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSubnetGroups(input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSubnetGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterSubnetGroupsPages(input *redshift.DescribeClusterSubnetGroupsInput, fn func(*redshift.DescribeClusterSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSubnetGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSubnetGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSubnetGroupsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterSubnetGroupsInput, fn func(*redshift.DescribeClusterSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterSubnetGroupsOutput{}
	fnCacher := func(out *redshift.DescribeClusterSubnetGroupsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterSubnetGroupsWithContext(ctx context.Context, input *redshift.DescribeClusterSubnetGroupsInput, opts ...request.Option) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterSubnetGroupsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterTracks(input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterTracksOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterTracks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterTracksPages(input *redshift.DescribeClusterTracksInput, fn func(*redshift.DescribeClusterTracksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterTracksOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterTracksOutput{}
	fnCacher := func(out *redshift.DescribeClusterTracksOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterTracksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterTracksPagesWithContext(ctx context.Context, input *redshift.DescribeClusterTracksInput, fn func(*redshift.DescribeClusterTracksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterTracksOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterTracksOutput{}
	fnCacher := func(out *redshift.DescribeClusterTracksOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterTracksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterTracksWithContext(ctx context.Context, input *redshift.DescribeClusterTracksInput, opts ...request.Option) (*redshift.DescribeClusterTracksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterTracksOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterTracksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterVersions(input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterVersionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusterVersionsPages(input *redshift.DescribeClusterVersionsInput, fn func(*redshift.DescribeClusterVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterVersionsOutput{}
	fnCacher := func(out *redshift.DescribeClusterVersionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterVersionsPagesWithContext(ctx context.Context, input *redshift.DescribeClusterVersionsInput, fn func(*redshift.DescribeClusterVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClusterVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClusterVersionsOutput{}
	fnCacher := func(out *redshift.DescribeClusterVersionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClusterVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClusterVersionsWithContext(ctx context.Context, input *redshift.DescribeClusterVersionsInput, opts ...request.Option) (*redshift.DescribeClusterVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClusterVersionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusterVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClusters(input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClustersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeClustersPages(input *redshift.DescribeClustersInput, fn func(*redshift.DescribeClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClustersOutput{}
	fnCacher := func(out *redshift.DescribeClustersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClustersPagesWithContext(ctx context.Context, input *redshift.DescribeClustersInput, fn func(*redshift.DescribeClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeClustersOutput{}
	fnCacher := func(out *redshift.DescribeClustersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeClustersWithContext(ctx context.Context, input *redshift.DescribeClustersInput, opts ...request.Option) (*redshift.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeClustersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataShares(input *redshift.DescribeDataSharesInput) (*redshift.DescribeDataSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataSharesForConsumer(input *redshift.DescribeDataSharesForConsumerInput) (*redshift.DescribeDataSharesForConsumerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesForConsumerOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataSharesForConsumer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataSharesForConsumerPages(input *redshift.DescribeDataSharesForConsumerInput, fn func(*redshift.DescribeDataSharesForConsumerOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesForConsumerOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesForConsumerOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesForConsumerOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesForConsumerPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesForConsumerPagesWithContext(ctx context.Context, input *redshift.DescribeDataSharesForConsumerInput, fn func(*redshift.DescribeDataSharesForConsumerOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesForConsumerOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesForConsumerOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesForConsumerOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesForConsumerPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesForConsumerWithContext(ctx context.Context, input *redshift.DescribeDataSharesForConsumerInput, opts ...request.Option) (*redshift.DescribeDataSharesForConsumerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesForConsumerOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataSharesForConsumerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataSharesForProducer(input *redshift.DescribeDataSharesForProducerInput) (*redshift.DescribeDataSharesForProducerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesForProducerOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataSharesForProducer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataSharesForProducerPages(input *redshift.DescribeDataSharesForProducerInput, fn func(*redshift.DescribeDataSharesForProducerOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesForProducerOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesForProducerOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesForProducerOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesForProducerPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesForProducerPagesWithContext(ctx context.Context, input *redshift.DescribeDataSharesForProducerInput, fn func(*redshift.DescribeDataSharesForProducerOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesForProducerOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesForProducerOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesForProducerOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesForProducerPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesForProducerWithContext(ctx context.Context, input *redshift.DescribeDataSharesForProducerInput, opts ...request.Option) (*redshift.DescribeDataSharesForProducerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesForProducerOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataSharesForProducerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDataSharesPages(input *redshift.DescribeDataSharesInput, fn func(*redshift.DescribeDataSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesPagesWithContext(ctx context.Context, input *redshift.DescribeDataSharesInput, fn func(*redshift.DescribeDataSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDataSharesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDataSharesOutput{}
	fnCacher := func(out *redshift.DescribeDataSharesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDataSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDataSharesWithContext(ctx context.Context, input *redshift.DescribeDataSharesInput, opts ...request.Option) (*redshift.DescribeDataSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDataSharesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDataSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDefaultClusterParameters(input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDefaultClusterParametersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDefaultClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeDefaultClusterParametersPages(input *redshift.DescribeDefaultClusterParametersInput, fn func(*redshift.DescribeDefaultClusterParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDefaultClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDefaultClusterParametersOutput{}
	fnCacher := func(out *redshift.DescribeDefaultClusterParametersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDefaultClusterParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDefaultClusterParametersPagesWithContext(ctx context.Context, input *redshift.DescribeDefaultClusterParametersInput, fn func(*redshift.DescribeDefaultClusterParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeDefaultClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeDefaultClusterParametersOutput{}
	fnCacher := func(out *redshift.DescribeDefaultClusterParametersOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeDefaultClusterParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeDefaultClusterParametersWithContext(ctx context.Context, input *redshift.DescribeDefaultClusterParametersInput, opts ...request.Option) (*redshift.DescribeDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeDefaultClusterParametersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeDefaultClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEndpointAccess(input *redshift.DescribeEndpointAccessInput) (*redshift.DescribeEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEndpointAccessOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEndpointAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEndpointAccessPages(input *redshift.DescribeEndpointAccessInput, fn func(*redshift.DescribeEndpointAccessOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEndpointAccessOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEndpointAccessOutput{}
	fnCacher := func(out *redshift.DescribeEndpointAccessOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEndpointAccessPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEndpointAccessPagesWithContext(ctx context.Context, input *redshift.DescribeEndpointAccessInput, fn func(*redshift.DescribeEndpointAccessOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEndpointAccessOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEndpointAccessOutput{}
	fnCacher := func(out *redshift.DescribeEndpointAccessOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEndpointAccessPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEndpointAccessWithContext(ctx context.Context, input *redshift.DescribeEndpointAccessInput, opts ...request.Option) (*redshift.DescribeEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEndpointAccessOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEndpointAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEndpointAuthorization(input *redshift.DescribeEndpointAuthorizationInput) (*redshift.DescribeEndpointAuthorizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEndpointAuthorizationOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEndpointAuthorization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEndpointAuthorizationPages(input *redshift.DescribeEndpointAuthorizationInput, fn func(*redshift.DescribeEndpointAuthorizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEndpointAuthorizationOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEndpointAuthorizationOutput{}
	fnCacher := func(out *redshift.DescribeEndpointAuthorizationOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEndpointAuthorizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEndpointAuthorizationPagesWithContext(ctx context.Context, input *redshift.DescribeEndpointAuthorizationInput, fn func(*redshift.DescribeEndpointAuthorizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEndpointAuthorizationOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEndpointAuthorizationOutput{}
	fnCacher := func(out *redshift.DescribeEndpointAuthorizationOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEndpointAuthorizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEndpointAuthorizationWithContext(ctx context.Context, input *redshift.DescribeEndpointAuthorizationInput, opts ...request.Option) (*redshift.DescribeEndpointAuthorizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEndpointAuthorizationOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEndpointAuthorizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEventCategories(input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventCategoriesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEventCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEventCategoriesWithContext(ctx context.Context, input *redshift.DescribeEventCategoriesInput, opts ...request.Option) (*redshift.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventCategoriesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEventCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEventSubscriptions(input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEventSubscriptionsPages(input *redshift.DescribeEventSubscriptionsInput, fn func(*redshift.DescribeEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *redshift.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEventSubscriptionsPagesWithContext(ctx context.Context, input *redshift.DescribeEventSubscriptionsInput, fn func(*redshift.DescribeEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *redshift.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEventSubscriptionsWithContext(ctx context.Context, input *redshift.DescribeEventSubscriptionsInput, opts ...request.Option) (*redshift.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEvents(input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeEventsPages(input *redshift.DescribeEventsInput, fn func(*redshift.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEventsOutput{}
	fnCacher := func(out *redshift.DescribeEventsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEventsPagesWithContext(ctx context.Context, input *redshift.DescribeEventsInput, fn func(*redshift.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeEventsOutput{}
	fnCacher := func(out *redshift.DescribeEventsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeEventsWithContext(ctx context.Context, input *redshift.DescribeEventsInput, opts ...request.Option) (*redshift.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeEventsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeHsmClientCertificates(input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeHsmClientCertificatesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeHsmClientCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeHsmClientCertificatesPages(input *redshift.DescribeHsmClientCertificatesInput, fn func(*redshift.DescribeHsmClientCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeHsmClientCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeHsmClientCertificatesOutput{}
	fnCacher := func(out *redshift.DescribeHsmClientCertificatesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeHsmClientCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeHsmClientCertificatesPagesWithContext(ctx context.Context, input *redshift.DescribeHsmClientCertificatesInput, fn func(*redshift.DescribeHsmClientCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeHsmClientCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeHsmClientCertificatesOutput{}
	fnCacher := func(out *redshift.DescribeHsmClientCertificatesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeHsmClientCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeHsmClientCertificatesWithContext(ctx context.Context, input *redshift.DescribeHsmClientCertificatesInput, opts ...request.Option) (*redshift.DescribeHsmClientCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeHsmClientCertificatesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeHsmClientCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeHsmConfigurations(input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeHsmConfigurationsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeHsmConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeHsmConfigurationsPages(input *redshift.DescribeHsmConfigurationsInput, fn func(*redshift.DescribeHsmConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeHsmConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeHsmConfigurationsOutput{}
	fnCacher := func(out *redshift.DescribeHsmConfigurationsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeHsmConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeHsmConfigurationsPagesWithContext(ctx context.Context, input *redshift.DescribeHsmConfigurationsInput, fn func(*redshift.DescribeHsmConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeHsmConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeHsmConfigurationsOutput{}
	fnCacher := func(out *redshift.DescribeHsmConfigurationsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeHsmConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeHsmConfigurationsWithContext(ctx context.Context, input *redshift.DescribeHsmConfigurationsInput, opts ...request.Option) (*redshift.DescribeHsmConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeHsmConfigurationsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeHsmConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeNodeConfigurationOptions(input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeNodeConfigurationOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeNodeConfigurationOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeNodeConfigurationOptionsPages(input *redshift.DescribeNodeConfigurationOptionsInput, fn func(*redshift.DescribeNodeConfigurationOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeNodeConfigurationOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeNodeConfigurationOptionsOutput{}
	fnCacher := func(out *redshift.DescribeNodeConfigurationOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeNodeConfigurationOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeNodeConfigurationOptionsPagesWithContext(ctx context.Context, input *redshift.DescribeNodeConfigurationOptionsInput, fn func(*redshift.DescribeNodeConfigurationOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeNodeConfigurationOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeNodeConfigurationOptionsOutput{}
	fnCacher := func(out *redshift.DescribeNodeConfigurationOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeNodeConfigurationOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeNodeConfigurationOptionsWithContext(ctx context.Context, input *redshift.DescribeNodeConfigurationOptionsInput, opts ...request.Option) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeNodeConfigurationOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeNodeConfigurationOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeOrderableClusterOptions(input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeOrderableClusterOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeOrderableClusterOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeOrderableClusterOptionsPages(input *redshift.DescribeOrderableClusterOptionsInput, fn func(*redshift.DescribeOrderableClusterOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeOrderableClusterOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeOrderableClusterOptionsOutput{}
	fnCacher := func(out *redshift.DescribeOrderableClusterOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeOrderableClusterOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeOrderableClusterOptionsPagesWithContext(ctx context.Context, input *redshift.DescribeOrderableClusterOptionsInput, fn func(*redshift.DescribeOrderableClusterOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeOrderableClusterOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeOrderableClusterOptionsOutput{}
	fnCacher := func(out *redshift.DescribeOrderableClusterOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeOrderableClusterOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeOrderableClusterOptionsWithContext(ctx context.Context, input *redshift.DescribeOrderableClusterOptionsInput, opts ...request.Option) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeOrderableClusterOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeOrderableClusterOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribePartners(input *redshift.DescribePartnersInput) (*redshift.DescribePartnersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribePartnersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribePartners(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribePartnersWithContext(ctx context.Context, input *redshift.DescribePartnersInput, opts ...request.Option) (*redshift.DescribePartnersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribePartnersOutput), nil
	}
	out, err := c.RedshiftAPI.DescribePartnersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodeExchangeStatus(input *redshift.DescribeReservedNodeExchangeStatusInput) (*redshift.DescribeReservedNodeExchangeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodeExchangeStatusOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodeExchangeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodeExchangeStatusPages(input *redshift.DescribeReservedNodeExchangeStatusInput, fn func(*redshift.DescribeReservedNodeExchangeStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodeExchangeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodeExchangeStatusOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodeExchangeStatusOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodeExchangeStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodeExchangeStatusPagesWithContext(ctx context.Context, input *redshift.DescribeReservedNodeExchangeStatusInput, fn func(*redshift.DescribeReservedNodeExchangeStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodeExchangeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodeExchangeStatusOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodeExchangeStatusOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodeExchangeStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodeExchangeStatusWithContext(ctx context.Context, input *redshift.DescribeReservedNodeExchangeStatusInput, opts ...request.Option) (*redshift.DescribeReservedNodeExchangeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodeExchangeStatusOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodeExchangeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodeOfferings(input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodeOfferingsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodeOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodeOfferingsPages(input *redshift.DescribeReservedNodeOfferingsInput, fn func(*redshift.DescribeReservedNodeOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodeOfferingsOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodeOfferingsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodeOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodeOfferingsPagesWithContext(ctx context.Context, input *redshift.DescribeReservedNodeOfferingsInput, fn func(*redshift.DescribeReservedNodeOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodeOfferingsOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodeOfferingsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodeOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodeOfferingsWithContext(ctx context.Context, input *redshift.DescribeReservedNodeOfferingsInput, opts ...request.Option) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodeOfferingsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodeOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodes(input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeReservedNodesPages(input *redshift.DescribeReservedNodesInput, fn func(*redshift.DescribeReservedNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodesOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodesPagesWithContext(ctx context.Context, input *redshift.DescribeReservedNodesInput, fn func(*redshift.DescribeReservedNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeReservedNodesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeReservedNodesOutput{}
	fnCacher := func(out *redshift.DescribeReservedNodesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeReservedNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeReservedNodesWithContext(ctx context.Context, input *redshift.DescribeReservedNodesInput, opts ...request.Option) (*redshift.DescribeReservedNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeReservedNodesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeReservedNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeResize(input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeResizeOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeResize(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeResizeWithContext(ctx context.Context, input *redshift.DescribeResizeInput, opts ...request.Option) (*redshift.DescribeResizeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeResizeOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeResizeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeScheduledActions(input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeScheduledActionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeScheduledActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeScheduledActionsPages(input *redshift.DescribeScheduledActionsInput, fn func(*redshift.DescribeScheduledActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeScheduledActionsOutput{}
	fnCacher := func(out *redshift.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeScheduledActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeScheduledActionsPagesWithContext(ctx context.Context, input *redshift.DescribeScheduledActionsInput, fn func(*redshift.DescribeScheduledActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeScheduledActionsOutput{}
	fnCacher := func(out *redshift.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeScheduledActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeScheduledActionsWithContext(ctx context.Context, input *redshift.DescribeScheduledActionsInput, opts ...request.Option) (*redshift.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeScheduledActionsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeScheduledActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeSnapshotCopyGrants(input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeSnapshotCopyGrantsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeSnapshotCopyGrants(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeSnapshotCopyGrantsPages(input *redshift.DescribeSnapshotCopyGrantsInput, fn func(*redshift.DescribeSnapshotCopyGrantsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeSnapshotCopyGrantsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeSnapshotCopyGrantsOutput{}
	fnCacher := func(out *redshift.DescribeSnapshotCopyGrantsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeSnapshotCopyGrantsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeSnapshotCopyGrantsPagesWithContext(ctx context.Context, input *redshift.DescribeSnapshotCopyGrantsInput, fn func(*redshift.DescribeSnapshotCopyGrantsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeSnapshotCopyGrantsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeSnapshotCopyGrantsOutput{}
	fnCacher := func(out *redshift.DescribeSnapshotCopyGrantsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeSnapshotCopyGrantsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeSnapshotCopyGrantsWithContext(ctx context.Context, input *redshift.DescribeSnapshotCopyGrantsInput, opts ...request.Option) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeSnapshotCopyGrantsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeSnapshotCopyGrantsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeSnapshotSchedules(input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeSnapshotSchedulesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeSnapshotSchedules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeSnapshotSchedulesPages(input *redshift.DescribeSnapshotSchedulesInput, fn func(*redshift.DescribeSnapshotSchedulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeSnapshotSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeSnapshotSchedulesOutput{}
	fnCacher := func(out *redshift.DescribeSnapshotSchedulesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeSnapshotSchedulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeSnapshotSchedulesPagesWithContext(ctx context.Context, input *redshift.DescribeSnapshotSchedulesInput, fn func(*redshift.DescribeSnapshotSchedulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeSnapshotSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeSnapshotSchedulesOutput{}
	fnCacher := func(out *redshift.DescribeSnapshotSchedulesOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeSnapshotSchedulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeSnapshotSchedulesWithContext(ctx context.Context, input *redshift.DescribeSnapshotSchedulesInput, opts ...request.Option) (*redshift.DescribeSnapshotSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeSnapshotSchedulesOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeSnapshotSchedulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeStorage(input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeStorageOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeStorage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeStorageWithContext(ctx context.Context, input *redshift.DescribeStorageInput, opts ...request.Option) (*redshift.DescribeStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeStorageOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeStorageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeTableRestoreStatus(input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeTableRestoreStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeTableRestoreStatusPages(input *redshift.DescribeTableRestoreStatusInput, fn func(*redshift.DescribeTableRestoreStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeTableRestoreStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeTableRestoreStatusOutput{}
	fnCacher := func(out *redshift.DescribeTableRestoreStatusOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeTableRestoreStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeTableRestoreStatusPagesWithContext(ctx context.Context, input *redshift.DescribeTableRestoreStatusInput, fn func(*redshift.DescribeTableRestoreStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeTableRestoreStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeTableRestoreStatusOutput{}
	fnCacher := func(out *redshift.DescribeTableRestoreStatusOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeTableRestoreStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeTableRestoreStatusWithContext(ctx context.Context, input *redshift.DescribeTableRestoreStatusInput, opts ...request.Option) (*redshift.DescribeTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeTableRestoreStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeTags(input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeTagsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeTagsPages(input *redshift.DescribeTagsInput, fn func(*redshift.DescribeTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeTagsOutput{}
	fnCacher := func(out *redshift.DescribeTagsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeTagsPagesWithContext(ctx context.Context, input *redshift.DescribeTagsInput, fn func(*redshift.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeTagsOutput{}
	fnCacher := func(out *redshift.DescribeTagsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeTagsWithContext(ctx context.Context, input *redshift.DescribeTagsInput, opts ...request.Option) (*redshift.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeTagsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeUsageLimits(input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeUsageLimitsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeUsageLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) DescribeUsageLimitsPages(input *redshift.DescribeUsageLimitsInput, fn func(*redshift.DescribeUsageLimitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeUsageLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeUsageLimitsOutput{}
	fnCacher := func(out *redshift.DescribeUsageLimitsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeUsageLimitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeUsageLimitsPagesWithContext(ctx context.Context, input *redshift.DescribeUsageLimitsInput, fn func(*redshift.DescribeUsageLimitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.DescribeUsageLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.DescribeUsageLimitsOutput{}
	fnCacher := func(out *redshift.DescribeUsageLimitsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.DescribeUsageLimitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) DescribeUsageLimitsWithContext(ctx context.Context, input *redshift.DescribeUsageLimitsInput, opts ...request.Option) (*redshift.DescribeUsageLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.DescribeUsageLimitsOutput), nil
	}
	out, err := c.RedshiftAPI.DescribeUsageLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetClusterCredentials(input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetClusterCredentialsOutput), nil
	}
	out, err := c.RedshiftAPI.GetClusterCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetClusterCredentialsWithContext(ctx context.Context, input *redshift.GetClusterCredentialsInput, opts ...request.Option) (*redshift.GetClusterCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetClusterCredentialsOutput), nil
	}
	out, err := c.RedshiftAPI.GetClusterCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetClusterCredentialsWithIAM(input *redshift.GetClusterCredentialsWithIAMInput) (*redshift.GetClusterCredentialsWithIAMOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetClusterCredentialsWithIAMOutput), nil
	}
	out, err := c.RedshiftAPI.GetClusterCredentialsWithIAM(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetClusterCredentialsWithIAMWithContext(ctx context.Context, input *redshift.GetClusterCredentialsWithIAMInput, opts ...request.Option) (*redshift.GetClusterCredentialsWithIAMOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetClusterCredentialsWithIAMOutput), nil
	}
	out, err := c.RedshiftAPI.GetClusterCredentialsWithIAMWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetReservedNodeExchangeConfigurationOptions(input *redshift.GetReservedNodeExchangeConfigurationOptionsInput) (*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.GetReservedNodeExchangeConfigurationOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetReservedNodeExchangeConfigurationOptionsPages(input *redshift.GetReservedNodeExchangeConfigurationOptionsInput, fn func(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.GetReservedNodeExchangeConfigurationOptionsOutput{}
	fnCacher := func(out *redshift.GetReservedNodeExchangeConfigurationOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.GetReservedNodeExchangeConfigurationOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) GetReservedNodeExchangeConfigurationOptionsPagesWithContext(ctx context.Context, input *redshift.GetReservedNodeExchangeConfigurationOptionsInput, fn func(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.GetReservedNodeExchangeConfigurationOptionsOutput{}
	fnCacher := func(out *redshift.GetReservedNodeExchangeConfigurationOptionsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.GetReservedNodeExchangeConfigurationOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) GetReservedNodeExchangeConfigurationOptionsWithContext(ctx context.Context, input *redshift.GetReservedNodeExchangeConfigurationOptionsInput, opts ...request.Option) (*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput), nil
	}
	out, err := c.RedshiftAPI.GetReservedNodeExchangeConfigurationOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetReservedNodeExchangeOfferings(input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetReservedNodeExchangeOfferingsOutput), nil
	}
	out, err := c.RedshiftAPI.GetReservedNodeExchangeOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Redshift) GetReservedNodeExchangeOfferingsPages(input *redshift.GetReservedNodeExchangeOfferingsInput, fn func(*redshift.GetReservedNodeExchangeOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.GetReservedNodeExchangeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.GetReservedNodeExchangeOfferingsOutput{}
	fnCacher := func(out *redshift.GetReservedNodeExchangeOfferingsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.GetReservedNodeExchangeOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) GetReservedNodeExchangeOfferingsPagesWithContext(ctx context.Context, input *redshift.GetReservedNodeExchangeOfferingsInput, fn func(*redshift.GetReservedNodeExchangeOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshift.GetReservedNodeExchangeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &redshift.GetReservedNodeExchangeOfferingsOutput{}
	fnCacher := func(out *redshift.GetReservedNodeExchangeOfferingsOutput, more bool) bool {
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
	if err := c.RedshiftAPI.GetReservedNodeExchangeOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Redshift) GetReservedNodeExchangeOfferingsWithContext(ctx context.Context, input *redshift.GetReservedNodeExchangeOfferingsInput, opts ...request.Option) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshift.GetReservedNodeExchangeOfferingsOutput), nil
	}
	out, err := c.RedshiftAPI.GetReservedNodeExchangeOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

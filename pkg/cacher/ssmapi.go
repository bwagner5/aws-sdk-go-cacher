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
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SSM struct {
	ssmiface.SSMAPI
	cache *cache.Cache
}

func NewSSM(ssmapi ssmiface.SSMAPI) *SSM {
	return &SSM{
		SSMAPI: ssmapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SSM) DescribeActivations(input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeActivationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeActivations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeActivationsPages(input *ssm.DescribeActivationsInput, fn func(*ssm.DescribeActivationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeActivationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeActivationsOutput{}
	fnCacher := func(out *ssm.DescribeActivationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeActivationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeActivationsPagesWithContext(ctx context.Context, input *ssm.DescribeActivationsInput, fn func(*ssm.DescribeActivationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeActivationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeActivationsOutput{}
	fnCacher := func(out *ssm.DescribeActivationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeActivationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeActivationsWithContext(ctx context.Context, input *ssm.DescribeActivationsInput, opts ...request.Option) (*ssm.DescribeActivationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeActivationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeActivationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociation(input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociationExecutionTargets(input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationExecutionTargetsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociationExecutionTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociationExecutionTargetsPages(input *ssm.DescribeAssociationExecutionTargetsInput, fn func(*ssm.DescribeAssociationExecutionTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAssociationExecutionTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAssociationExecutionTargetsOutput{}
	fnCacher := func(out *ssm.DescribeAssociationExecutionTargetsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAssociationExecutionTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAssociationExecutionTargetsPagesWithContext(ctx context.Context, input *ssm.DescribeAssociationExecutionTargetsInput, fn func(*ssm.DescribeAssociationExecutionTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAssociationExecutionTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAssociationExecutionTargetsOutput{}
	fnCacher := func(out *ssm.DescribeAssociationExecutionTargetsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAssociationExecutionTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAssociationExecutionTargetsWithContext(ctx context.Context, input *ssm.DescribeAssociationExecutionTargetsInput, opts ...request.Option) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationExecutionTargetsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociationExecutionTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociationExecutions(input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociationExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociationExecutionsPages(input *ssm.DescribeAssociationExecutionsInput, fn func(*ssm.DescribeAssociationExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAssociationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAssociationExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAssociationExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAssociationExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAssociationExecutionsPagesWithContext(ctx context.Context, input *ssm.DescribeAssociationExecutionsInput, fn func(*ssm.DescribeAssociationExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAssociationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAssociationExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAssociationExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAssociationExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAssociationExecutionsWithContext(ctx context.Context, input *ssm.DescribeAssociationExecutionsInput, opts ...request.Option) (*ssm.DescribeAssociationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociationExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAssociationWithContext(ctx context.Context, input *ssm.DescribeAssociationInput, opts ...request.Option) (*ssm.DescribeAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAssociationOutput), nil
	}
	out, err := c.SSMAPI.DescribeAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAutomationExecutions(input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAutomationExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAutomationExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAutomationExecutionsPages(input *ssm.DescribeAutomationExecutionsInput, fn func(*ssm.DescribeAutomationExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAutomationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAutomationExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAutomationExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAutomationExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAutomationExecutionsPagesWithContext(ctx context.Context, input *ssm.DescribeAutomationExecutionsInput, fn func(*ssm.DescribeAutomationExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAutomationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAutomationExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAutomationExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAutomationExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAutomationExecutionsWithContext(ctx context.Context, input *ssm.DescribeAutomationExecutionsInput, opts ...request.Option) (*ssm.DescribeAutomationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAutomationExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAutomationExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAutomationStepExecutions(input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAutomationStepExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAutomationStepExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAutomationStepExecutionsPages(input *ssm.DescribeAutomationStepExecutionsInput, fn func(*ssm.DescribeAutomationStepExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAutomationStepExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAutomationStepExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAutomationStepExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAutomationStepExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAutomationStepExecutionsPagesWithContext(ctx context.Context, input *ssm.DescribeAutomationStepExecutionsInput, fn func(*ssm.DescribeAutomationStepExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAutomationStepExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAutomationStepExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeAutomationStepExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAutomationStepExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAutomationStepExecutionsWithContext(ctx context.Context, input *ssm.DescribeAutomationStepExecutionsInput, opts ...request.Option) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAutomationStepExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeAutomationStepExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAvailablePatches(input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAvailablePatchesOutput), nil
	}
	out, err := c.SSMAPI.DescribeAvailablePatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeAvailablePatchesPages(input *ssm.DescribeAvailablePatchesInput, fn func(*ssm.DescribeAvailablePatchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAvailablePatchesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAvailablePatchesOutput{}
	fnCacher := func(out *ssm.DescribeAvailablePatchesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAvailablePatchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAvailablePatchesPagesWithContext(ctx context.Context, input *ssm.DescribeAvailablePatchesInput, fn func(*ssm.DescribeAvailablePatchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeAvailablePatchesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeAvailablePatchesOutput{}
	fnCacher := func(out *ssm.DescribeAvailablePatchesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeAvailablePatchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeAvailablePatchesWithContext(ctx context.Context, input *ssm.DescribeAvailablePatchesInput, opts ...request.Option) (*ssm.DescribeAvailablePatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeAvailablePatchesOutput), nil
	}
	out, err := c.SSMAPI.DescribeAvailablePatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeDocument(input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeDocumentOutput), nil
	}
	out, err := c.SSMAPI.DescribeDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeDocumentPermission(input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeDocumentPermissionOutput), nil
	}
	out, err := c.SSMAPI.DescribeDocumentPermission(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeDocumentPermissionWithContext(ctx context.Context, input *ssm.DescribeDocumentPermissionInput, opts ...request.Option) (*ssm.DescribeDocumentPermissionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeDocumentPermissionOutput), nil
	}
	out, err := c.SSMAPI.DescribeDocumentPermissionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeDocumentWithContext(ctx context.Context, input *ssm.DescribeDocumentInput, opts ...request.Option) (*ssm.DescribeDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeDocumentOutput), nil
	}
	out, err := c.SSMAPI.DescribeDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeEffectiveInstanceAssociations(input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeEffectiveInstanceAssociationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeEffectiveInstanceAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeEffectiveInstanceAssociationsPages(input *ssm.DescribeEffectiveInstanceAssociationsInput, fn func(*ssm.DescribeEffectiveInstanceAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeEffectiveInstanceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeEffectiveInstanceAssociationsOutput{}
	fnCacher := func(out *ssm.DescribeEffectiveInstanceAssociationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeEffectiveInstanceAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeEffectiveInstanceAssociationsPagesWithContext(ctx context.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput, fn func(*ssm.DescribeEffectiveInstanceAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeEffectiveInstanceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeEffectiveInstanceAssociationsOutput{}
	fnCacher := func(out *ssm.DescribeEffectiveInstanceAssociationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeEffectiveInstanceAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeEffectiveInstanceAssociationsWithContext(ctx context.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput, opts ...request.Option) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeEffectiveInstanceAssociationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeEffectiveInstanceAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeEffectivePatchesForPatchBaseline(input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeEffectivePatchesForPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.DescribeEffectivePatchesForPatchBaseline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeEffectivePatchesForPatchBaselinePages(input *ssm.DescribeEffectivePatchesForPatchBaselineInput, fn func(*ssm.DescribeEffectivePatchesForPatchBaselineOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeEffectivePatchesForPatchBaselineOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeEffectivePatchesForPatchBaselineOutput{}
	fnCacher := func(out *ssm.DescribeEffectivePatchesForPatchBaselineOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeEffectivePatchesForPatchBaselinePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeEffectivePatchesForPatchBaselinePagesWithContext(ctx context.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput, fn func(*ssm.DescribeEffectivePatchesForPatchBaselineOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeEffectivePatchesForPatchBaselineOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeEffectivePatchesForPatchBaselineOutput{}
	fnCacher := func(out *ssm.DescribeEffectivePatchesForPatchBaselineOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeEffectivePatchesForPatchBaselinePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeEffectivePatchesForPatchBaselineWithContext(ctx context.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput, opts ...request.Option) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeEffectivePatchesForPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.DescribeEffectivePatchesForPatchBaselineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstanceAssociationsStatus(input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstanceAssociationsStatusOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstanceAssociationsStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstanceAssociationsStatusPages(input *ssm.DescribeInstanceAssociationsStatusInput, fn func(*ssm.DescribeInstanceAssociationsStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstanceAssociationsStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstanceAssociationsStatusOutput{}
	fnCacher := func(out *ssm.DescribeInstanceAssociationsStatusOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstanceAssociationsStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstanceAssociationsStatusPagesWithContext(ctx context.Context, input *ssm.DescribeInstanceAssociationsStatusInput, fn func(*ssm.DescribeInstanceAssociationsStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstanceAssociationsStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstanceAssociationsStatusOutput{}
	fnCacher := func(out *ssm.DescribeInstanceAssociationsStatusOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstanceAssociationsStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstanceAssociationsStatusWithContext(ctx context.Context, input *ssm.DescribeInstanceAssociationsStatusInput, opts ...request.Option) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstanceAssociationsStatusOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstanceAssociationsStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstanceInformation(input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstanceInformationOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstanceInformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstanceInformationPages(input *ssm.DescribeInstanceInformationInput, fn func(*ssm.DescribeInstanceInformationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstanceInformationOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstanceInformationOutput{}
	fnCacher := func(out *ssm.DescribeInstanceInformationOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstanceInformationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstanceInformationPagesWithContext(ctx context.Context, input *ssm.DescribeInstanceInformationInput, fn func(*ssm.DescribeInstanceInformationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstanceInformationOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstanceInformationOutput{}
	fnCacher := func(out *ssm.DescribeInstanceInformationOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstanceInformationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstanceInformationWithContext(ctx context.Context, input *ssm.DescribeInstanceInformationInput, opts ...request.Option) (*ssm.DescribeInstanceInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstanceInformationOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstanceInformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatchStates(input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchStatesOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatchStates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatchStatesForPatchGroup(input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchStatesForPatchGroupOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatchStatesForPatchGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatchStatesForPatchGroupPages(input *ssm.DescribeInstancePatchStatesForPatchGroupInput, fn func(*ssm.DescribeInstancePatchStatesForPatchGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchStatesForPatchGroupOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchStatesForPatchGroupOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchStatesForPatchGroupOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchStatesForPatchGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchStatesForPatchGroupPagesWithContext(ctx context.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput, fn func(*ssm.DescribeInstancePatchStatesForPatchGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchStatesForPatchGroupOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchStatesForPatchGroupOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchStatesForPatchGroupOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchStatesForPatchGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchStatesForPatchGroupWithContext(ctx context.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput, opts ...request.Option) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchStatesForPatchGroupOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatchStatesForPatchGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatchStatesPages(input *ssm.DescribeInstancePatchStatesInput, fn func(*ssm.DescribeInstancePatchStatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchStatesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchStatesOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchStatesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchStatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchStatesPagesWithContext(ctx context.Context, input *ssm.DescribeInstancePatchStatesInput, fn func(*ssm.DescribeInstancePatchStatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchStatesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchStatesOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchStatesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchStatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchStatesWithContext(ctx context.Context, input *ssm.DescribeInstancePatchStatesInput, opts ...request.Option) (*ssm.DescribeInstancePatchStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchStatesOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatchStatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatches(input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchesOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInstancePatchesPages(input *ssm.DescribeInstancePatchesInput, fn func(*ssm.DescribeInstancePatchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchesOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchesPagesWithContext(ctx context.Context, input *ssm.DescribeInstancePatchesInput, fn func(*ssm.DescribeInstancePatchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInstancePatchesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInstancePatchesOutput{}
	fnCacher := func(out *ssm.DescribeInstancePatchesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInstancePatchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInstancePatchesWithContext(ctx context.Context, input *ssm.DescribeInstancePatchesInput, opts ...request.Option) (*ssm.DescribeInstancePatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInstancePatchesOutput), nil
	}
	out, err := c.SSMAPI.DescribeInstancePatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInventoryDeletions(input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInventoryDeletionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeInventoryDeletions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeInventoryDeletionsPages(input *ssm.DescribeInventoryDeletionsInput, fn func(*ssm.DescribeInventoryDeletionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInventoryDeletionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInventoryDeletionsOutput{}
	fnCacher := func(out *ssm.DescribeInventoryDeletionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInventoryDeletionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInventoryDeletionsPagesWithContext(ctx context.Context, input *ssm.DescribeInventoryDeletionsInput, fn func(*ssm.DescribeInventoryDeletionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeInventoryDeletionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeInventoryDeletionsOutput{}
	fnCacher := func(out *ssm.DescribeInventoryDeletionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeInventoryDeletionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeInventoryDeletionsWithContext(ctx context.Context, input *ssm.DescribeInventoryDeletionsInput, opts ...request.Option) (*ssm.DescribeInventoryDeletionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeInventoryDeletionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeInventoryDeletionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutionTaskInvocations(input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutionTaskInvocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutionTaskInvocationsPages(input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, fn func(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionTaskInvocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionTaskInvocationsPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, fn func(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionTaskInvocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionTaskInvocationsWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutionTaskInvocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutionTasks(input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTasksOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutionTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutionTasksPages(input *ssm.DescribeMaintenanceWindowExecutionTasksInput, fn func(*ssm.DescribeMaintenanceWindowExecutionTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionTasksOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionTasksOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionTasksPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput, fn func(*ssm.DescribeMaintenanceWindowExecutionTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionTasksOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionTasksOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionTasksWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionTasksOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutionTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutions(input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowExecutionsPages(input *ssm.DescribeMaintenanceWindowExecutionsInput, fn func(*ssm.DescribeMaintenanceWindowExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionsPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput, fn func(*ssm.DescribeMaintenanceWindowExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowExecutionsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowExecutionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowExecutionsWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowExecutionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowSchedule(input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowScheduleOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowSchedulePages(input *ssm.DescribeMaintenanceWindowScheduleInput, fn func(*ssm.DescribeMaintenanceWindowScheduleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowScheduleOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowScheduleOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowSchedulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowSchedulePagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowScheduleInput, fn func(*ssm.DescribeMaintenanceWindowScheduleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowScheduleOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowScheduleOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowSchedulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowScheduleWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowScheduleInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowScheduleOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowTargets(input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowTargetsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowTargetsPages(input *ssm.DescribeMaintenanceWindowTargetsInput, fn func(*ssm.DescribeMaintenanceWindowTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowTargetsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowTargetsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowTargetsPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowTargetsInput, fn func(*ssm.DescribeMaintenanceWindowTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowTargetsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowTargetsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowTargetsWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowTargetsInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowTargetsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowTasks(input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowTasksOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowTasksPages(input *ssm.DescribeMaintenanceWindowTasksInput, fn func(*ssm.DescribeMaintenanceWindowTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowTasksOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowTasksOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowTasksPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowTasksInput, fn func(*ssm.DescribeMaintenanceWindowTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowTasksOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowTasksOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowTasksWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowTasksInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowTasksOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindows(input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowsForTarget(input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowsForTargetOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowsForTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowsForTargetPages(input *ssm.DescribeMaintenanceWindowsForTargetInput, fn func(*ssm.DescribeMaintenanceWindowsForTargetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowsForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowsForTargetOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowsForTargetOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowsForTargetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowsForTargetPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput, fn func(*ssm.DescribeMaintenanceWindowsForTargetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowsForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowsForTargetOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowsForTargetOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowsForTargetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowsForTargetWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowsForTargetOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowsForTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeMaintenanceWindowsPages(input *ssm.DescribeMaintenanceWindowsInput, fn func(*ssm.DescribeMaintenanceWindowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowsPagesWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowsInput, fn func(*ssm.DescribeMaintenanceWindowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeMaintenanceWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeMaintenanceWindowsOutput{}
	fnCacher := func(out *ssm.DescribeMaintenanceWindowsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeMaintenanceWindowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeMaintenanceWindowsWithContext(ctx context.Context, input *ssm.DescribeMaintenanceWindowsInput, opts ...request.Option) (*ssm.DescribeMaintenanceWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeMaintenanceWindowsOutput), nil
	}
	out, err := c.SSMAPI.DescribeMaintenanceWindowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeOpsItems(input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeOpsItemsOutput), nil
	}
	out, err := c.SSMAPI.DescribeOpsItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeOpsItemsPages(input *ssm.DescribeOpsItemsInput, fn func(*ssm.DescribeOpsItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeOpsItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeOpsItemsOutput{}
	fnCacher := func(out *ssm.DescribeOpsItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeOpsItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeOpsItemsPagesWithContext(ctx context.Context, input *ssm.DescribeOpsItemsInput, fn func(*ssm.DescribeOpsItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeOpsItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeOpsItemsOutput{}
	fnCacher := func(out *ssm.DescribeOpsItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeOpsItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeOpsItemsWithContext(ctx context.Context, input *ssm.DescribeOpsItemsInput, opts ...request.Option) (*ssm.DescribeOpsItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeOpsItemsOutput), nil
	}
	out, err := c.SSMAPI.DescribeOpsItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeParameters(input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeParametersOutput), nil
	}
	out, err := c.SSMAPI.DescribeParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeParametersPages(input *ssm.DescribeParametersInput, fn func(*ssm.DescribeParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeParametersOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeParametersOutput{}
	fnCacher := func(out *ssm.DescribeParametersOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeParametersPagesWithContext(ctx context.Context, input *ssm.DescribeParametersInput, fn func(*ssm.DescribeParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeParametersOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeParametersOutput{}
	fnCacher := func(out *ssm.DescribeParametersOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeParametersWithContext(ctx context.Context, input *ssm.DescribeParametersInput, opts ...request.Option) (*ssm.DescribeParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeParametersOutput), nil
	}
	out, err := c.SSMAPI.DescribeParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchBaselines(input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchBaselinesOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchBaselines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchBaselinesPages(input *ssm.DescribePatchBaselinesInput, fn func(*ssm.DescribePatchBaselinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchBaselinesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchBaselinesOutput{}
	fnCacher := func(out *ssm.DescribePatchBaselinesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchBaselinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchBaselinesPagesWithContext(ctx context.Context, input *ssm.DescribePatchBaselinesInput, fn func(*ssm.DescribePatchBaselinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchBaselinesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchBaselinesOutput{}
	fnCacher := func(out *ssm.DescribePatchBaselinesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchBaselinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchBaselinesWithContext(ctx context.Context, input *ssm.DescribePatchBaselinesInput, opts ...request.Option) (*ssm.DescribePatchBaselinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchBaselinesOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchBaselinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchGroupState(input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchGroupStateOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchGroupState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchGroupStateWithContext(ctx context.Context, input *ssm.DescribePatchGroupStateInput, opts ...request.Option) (*ssm.DescribePatchGroupStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchGroupStateOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchGroupStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchGroups(input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchGroupsOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchGroupsPages(input *ssm.DescribePatchGroupsInput, fn func(*ssm.DescribePatchGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchGroupsOutput{}
	fnCacher := func(out *ssm.DescribePatchGroupsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchGroupsPagesWithContext(ctx context.Context, input *ssm.DescribePatchGroupsInput, fn func(*ssm.DescribePatchGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchGroupsOutput{}
	fnCacher := func(out *ssm.DescribePatchGroupsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchGroupsWithContext(ctx context.Context, input *ssm.DescribePatchGroupsInput, opts ...request.Option) (*ssm.DescribePatchGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchGroupsOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchProperties(input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchPropertiesOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchProperties(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribePatchPropertiesPages(input *ssm.DescribePatchPropertiesInput, fn func(*ssm.DescribePatchPropertiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchPropertiesOutput{}
	fnCacher := func(out *ssm.DescribePatchPropertiesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchPropertiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchPropertiesPagesWithContext(ctx context.Context, input *ssm.DescribePatchPropertiesInput, fn func(*ssm.DescribePatchPropertiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribePatchPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribePatchPropertiesOutput{}
	fnCacher := func(out *ssm.DescribePatchPropertiesOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribePatchPropertiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribePatchPropertiesWithContext(ctx context.Context, input *ssm.DescribePatchPropertiesInput, opts ...request.Option) (*ssm.DescribePatchPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribePatchPropertiesOutput), nil
	}
	out, err := c.SSMAPI.DescribePatchPropertiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeSessions(input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeSessionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) DescribeSessionsPages(input *ssm.DescribeSessionsInput, fn func(*ssm.DescribeSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeSessionsOutput{}
	fnCacher := func(out *ssm.DescribeSessionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeSessionsPagesWithContext(ctx context.Context, input *ssm.DescribeSessionsInput, fn func(*ssm.DescribeSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.DescribeSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.DescribeSessionsOutput{}
	fnCacher := func(out *ssm.DescribeSessionsOutput, more bool) bool {
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
	if err := c.SSMAPI.DescribeSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) DescribeSessionsWithContext(ctx context.Context, input *ssm.DescribeSessionsInput, opts ...request.Option) (*ssm.DescribeSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.DescribeSessionsOutput), nil
	}
	out, err := c.SSMAPI.DescribeSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetAutomationExecution(input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetAutomationExecutionOutput), nil
	}
	out, err := c.SSMAPI.GetAutomationExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetAutomationExecutionWithContext(ctx context.Context, input *ssm.GetAutomationExecutionInput, opts ...request.Option) (*ssm.GetAutomationExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetAutomationExecutionOutput), nil
	}
	out, err := c.SSMAPI.GetAutomationExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetCalendarState(input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetCalendarStateOutput), nil
	}
	out, err := c.SSMAPI.GetCalendarState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetCalendarStateWithContext(ctx context.Context, input *ssm.GetCalendarStateInput, opts ...request.Option) (*ssm.GetCalendarStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetCalendarStateOutput), nil
	}
	out, err := c.SSMAPI.GetCalendarStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetCommandInvocation(input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetCommandInvocationOutput), nil
	}
	out, err := c.SSMAPI.GetCommandInvocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetCommandInvocationWithContext(ctx context.Context, input *ssm.GetCommandInvocationInput, opts ...request.Option) (*ssm.GetCommandInvocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetCommandInvocationOutput), nil
	}
	out, err := c.SSMAPI.GetCommandInvocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetConnectionStatus(input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetConnectionStatusOutput), nil
	}
	out, err := c.SSMAPI.GetConnectionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetConnectionStatusWithContext(ctx context.Context, input *ssm.GetConnectionStatusInput, opts ...request.Option) (*ssm.GetConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetConnectionStatusOutput), nil
	}
	out, err := c.SSMAPI.GetConnectionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDefaultPatchBaseline(input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDefaultPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.GetDefaultPatchBaseline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDefaultPatchBaselineWithContext(ctx context.Context, input *ssm.GetDefaultPatchBaselineInput, opts ...request.Option) (*ssm.GetDefaultPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDefaultPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.GetDefaultPatchBaselineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDeployablePatchSnapshotForInstance(input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDeployablePatchSnapshotForInstanceOutput), nil
	}
	out, err := c.SSMAPI.GetDeployablePatchSnapshotForInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDeployablePatchSnapshotForInstanceWithContext(ctx context.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput, opts ...request.Option) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDeployablePatchSnapshotForInstanceOutput), nil
	}
	out, err := c.SSMAPI.GetDeployablePatchSnapshotForInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDocument(input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDocumentOutput), nil
	}
	out, err := c.SSMAPI.GetDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetDocumentWithContext(ctx context.Context, input *ssm.GetDocumentInput, opts ...request.Option) (*ssm.GetDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetDocumentOutput), nil
	}
	out, err := c.SSMAPI.GetDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetInventory(input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetInventoryOutput), nil
	}
	out, err := c.SSMAPI.GetInventory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetInventoryPages(input *ssm.GetInventoryInput, fn func(*ssm.GetInventoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetInventoryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetInventoryOutput{}
	fnCacher := func(out *ssm.GetInventoryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetInventoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetInventoryPagesWithContext(ctx context.Context, input *ssm.GetInventoryInput, fn func(*ssm.GetInventoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetInventoryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetInventoryOutput{}
	fnCacher := func(out *ssm.GetInventoryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetInventoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetInventorySchema(input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetInventorySchemaOutput), nil
	}
	out, err := c.SSMAPI.GetInventorySchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetInventorySchemaPages(input *ssm.GetInventorySchemaInput, fn func(*ssm.GetInventorySchemaOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetInventorySchemaOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetInventorySchemaOutput{}
	fnCacher := func(out *ssm.GetInventorySchemaOutput, more bool) bool {
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
	if err := c.SSMAPI.GetInventorySchemaPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetInventorySchemaPagesWithContext(ctx context.Context, input *ssm.GetInventorySchemaInput, fn func(*ssm.GetInventorySchemaOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetInventorySchemaOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetInventorySchemaOutput{}
	fnCacher := func(out *ssm.GetInventorySchemaOutput, more bool) bool {
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
	if err := c.SSMAPI.GetInventorySchemaPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetInventorySchemaWithContext(ctx context.Context, input *ssm.GetInventorySchemaInput, opts ...request.Option) (*ssm.GetInventorySchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetInventorySchemaOutput), nil
	}
	out, err := c.SSMAPI.GetInventorySchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetInventoryWithContext(ctx context.Context, input *ssm.GetInventoryInput, opts ...request.Option) (*ssm.GetInventoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetInventoryOutput), nil
	}
	out, err := c.SSMAPI.GetInventoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindow(input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecution(input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecutionTask(input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionTaskOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecutionTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecutionTaskInvocation(input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecutionTaskInvocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecutionTaskInvocationWithContext(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput, opts ...request.Option) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecutionTaskInvocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecutionTaskWithContext(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput, opts ...request.Option) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionTaskOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecutionTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowExecutionWithContext(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionInput, opts ...request.Option) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowExecutionOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowTask(input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowTaskOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowTaskWithContext(ctx context.Context, input *ssm.GetMaintenanceWindowTaskInput, opts ...request.Option) (*ssm.GetMaintenanceWindowTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowTaskOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetMaintenanceWindowWithContext(ctx context.Context, input *ssm.GetMaintenanceWindowInput, opts ...request.Option) (*ssm.GetMaintenanceWindowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetMaintenanceWindowOutput), nil
	}
	out, err := c.SSMAPI.GetMaintenanceWindowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsItem(input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsItemOutput), nil
	}
	out, err := c.SSMAPI.GetOpsItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsItemWithContext(ctx context.Context, input *ssm.GetOpsItemInput, opts ...request.Option) (*ssm.GetOpsItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsItemOutput), nil
	}
	out, err := c.SSMAPI.GetOpsItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsMetadata(input *ssm.GetOpsMetadataInput) (*ssm.GetOpsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsMetadataOutput), nil
	}
	out, err := c.SSMAPI.GetOpsMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsMetadataWithContext(ctx context.Context, input *ssm.GetOpsMetadataInput, opts ...request.Option) (*ssm.GetOpsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsMetadataOutput), nil
	}
	out, err := c.SSMAPI.GetOpsMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsSummary(input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsSummaryOutput), nil
	}
	out, err := c.SSMAPI.GetOpsSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetOpsSummaryPages(input *ssm.GetOpsSummaryInput, fn func(*ssm.GetOpsSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetOpsSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetOpsSummaryOutput{}
	fnCacher := func(out *ssm.GetOpsSummaryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetOpsSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetOpsSummaryPagesWithContext(ctx context.Context, input *ssm.GetOpsSummaryInput, fn func(*ssm.GetOpsSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetOpsSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetOpsSummaryOutput{}
	fnCacher := func(out *ssm.GetOpsSummaryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetOpsSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetOpsSummaryWithContext(ctx context.Context, input *ssm.GetOpsSummaryInput, opts ...request.Option) (*ssm.GetOpsSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetOpsSummaryOutput), nil
	}
	out, err := c.SSMAPI.GetOpsSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParameterOutput), nil
	}
	out, err := c.SSMAPI.GetParameter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParameterHistory(input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParameterHistoryOutput), nil
	}
	out, err := c.SSMAPI.GetParameterHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParameterHistoryPages(input *ssm.GetParameterHistoryInput, fn func(*ssm.GetParameterHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetParameterHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetParameterHistoryOutput{}
	fnCacher := func(out *ssm.GetParameterHistoryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetParameterHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetParameterHistoryPagesWithContext(ctx context.Context, input *ssm.GetParameterHistoryInput, fn func(*ssm.GetParameterHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetParameterHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetParameterHistoryOutput{}
	fnCacher := func(out *ssm.GetParameterHistoryOutput, more bool) bool {
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
	if err := c.SSMAPI.GetParameterHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetParameterHistoryWithContext(ctx context.Context, input *ssm.GetParameterHistoryInput, opts ...request.Option) (*ssm.GetParameterHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParameterHistoryOutput), nil
	}
	out, err := c.SSMAPI.GetParameterHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParameterWithContext(ctx context.Context, input *ssm.GetParameterInput, opts ...request.Option) (*ssm.GetParameterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParameterOutput), nil
	}
	out, err := c.SSMAPI.GetParameterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParameters(input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParametersOutput), nil
	}
	out, err := c.SSMAPI.GetParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParametersByPath(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParametersByPathOutput), nil
	}
	out, err := c.SSMAPI.GetParametersByPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParametersByPathPages(input *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetParametersByPathOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetParametersByPathOutput{}
	fnCacher := func(out *ssm.GetParametersByPathOutput, more bool) bool {
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
	if err := c.SSMAPI.GetParametersByPathPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetParametersByPathPagesWithContext(ctx context.Context, input *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetParametersByPathOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetParametersByPathOutput{}
	fnCacher := func(out *ssm.GetParametersByPathOutput, more bool) bool {
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
	if err := c.SSMAPI.GetParametersByPathPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetParametersByPathWithContext(ctx context.Context, input *ssm.GetParametersByPathInput, opts ...request.Option) (*ssm.GetParametersByPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParametersByPathOutput), nil
	}
	out, err := c.SSMAPI.GetParametersByPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetParametersWithContext(ctx context.Context, input *ssm.GetParametersInput, opts ...request.Option) (*ssm.GetParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetParametersOutput), nil
	}
	out, err := c.SSMAPI.GetParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetPatchBaseline(input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.GetPatchBaseline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetPatchBaselineForPatchGroup(input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetPatchBaselineForPatchGroupOutput), nil
	}
	out, err := c.SSMAPI.GetPatchBaselineForPatchGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetPatchBaselineForPatchGroupWithContext(ctx context.Context, input *ssm.GetPatchBaselineForPatchGroupInput, opts ...request.Option) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetPatchBaselineForPatchGroupOutput), nil
	}
	out, err := c.SSMAPI.GetPatchBaselineForPatchGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetPatchBaselineWithContext(ctx context.Context, input *ssm.GetPatchBaselineInput, opts ...request.Option) (*ssm.GetPatchBaselineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetPatchBaselineOutput), nil
	}
	out, err := c.SSMAPI.GetPatchBaselineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetResourcePolicies(input *ssm.GetResourcePoliciesInput) (*ssm.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetResourcePoliciesOutput), nil
	}
	out, err := c.SSMAPI.GetResourcePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetResourcePoliciesPages(input *ssm.GetResourcePoliciesInput, fn func(*ssm.GetResourcePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetResourcePoliciesOutput{}
	fnCacher := func(out *ssm.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.SSMAPI.GetResourcePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetResourcePoliciesPagesWithContext(ctx context.Context, input *ssm.GetResourcePoliciesInput, fn func(*ssm.GetResourcePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.GetResourcePoliciesOutput{}
	fnCacher := func(out *ssm.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.SSMAPI.GetResourcePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) GetResourcePoliciesWithContext(ctx context.Context, input *ssm.GetResourcePoliciesInput, opts ...request.Option) (*ssm.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetResourcePoliciesOutput), nil
	}
	out, err := c.SSMAPI.GetResourcePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetServiceSetting(input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetServiceSettingOutput), nil
	}
	out, err := c.SSMAPI.GetServiceSetting(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) GetServiceSettingWithContext(ctx context.Context, input *ssm.GetServiceSettingInput, opts ...request.Option) (*ssm.GetServiceSettingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.GetServiceSettingOutput), nil
	}
	out, err := c.SSMAPI.GetServiceSettingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListAssociationVersions(input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListAssociationVersionsOutput), nil
	}
	out, err := c.SSMAPI.ListAssociationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListAssociationVersionsPages(input *ssm.ListAssociationVersionsInput, fn func(*ssm.ListAssociationVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListAssociationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListAssociationVersionsOutput{}
	fnCacher := func(out *ssm.ListAssociationVersionsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListAssociationVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListAssociationVersionsPagesWithContext(ctx context.Context, input *ssm.ListAssociationVersionsInput, fn func(*ssm.ListAssociationVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListAssociationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListAssociationVersionsOutput{}
	fnCacher := func(out *ssm.ListAssociationVersionsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListAssociationVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListAssociationVersionsWithContext(ctx context.Context, input *ssm.ListAssociationVersionsInput, opts ...request.Option) (*ssm.ListAssociationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListAssociationVersionsOutput), nil
	}
	out, err := c.SSMAPI.ListAssociationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListAssociations(input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListAssociationsOutput), nil
	}
	out, err := c.SSMAPI.ListAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListAssociationsPages(input *ssm.ListAssociationsInput, fn func(*ssm.ListAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListAssociationsOutput{}
	fnCacher := func(out *ssm.ListAssociationsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListAssociationsPagesWithContext(ctx context.Context, input *ssm.ListAssociationsInput, fn func(*ssm.ListAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListAssociationsOutput{}
	fnCacher := func(out *ssm.ListAssociationsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListAssociationsWithContext(ctx context.Context, input *ssm.ListAssociationsInput, opts ...request.Option) (*ssm.ListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListAssociationsOutput), nil
	}
	out, err := c.SSMAPI.ListAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListCommandInvocations(input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListCommandInvocationsOutput), nil
	}
	out, err := c.SSMAPI.ListCommandInvocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListCommandInvocationsPages(input *ssm.ListCommandInvocationsInput, fn func(*ssm.ListCommandInvocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListCommandInvocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListCommandInvocationsOutput{}
	fnCacher := func(out *ssm.ListCommandInvocationsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListCommandInvocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListCommandInvocationsPagesWithContext(ctx context.Context, input *ssm.ListCommandInvocationsInput, fn func(*ssm.ListCommandInvocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListCommandInvocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListCommandInvocationsOutput{}
	fnCacher := func(out *ssm.ListCommandInvocationsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListCommandInvocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListCommandInvocationsWithContext(ctx context.Context, input *ssm.ListCommandInvocationsInput, opts ...request.Option) (*ssm.ListCommandInvocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListCommandInvocationsOutput), nil
	}
	out, err := c.SSMAPI.ListCommandInvocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListCommands(input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListCommandsOutput), nil
	}
	out, err := c.SSMAPI.ListCommands(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListCommandsPages(input *ssm.ListCommandsInput, fn func(*ssm.ListCommandsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListCommandsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListCommandsOutput{}
	fnCacher := func(out *ssm.ListCommandsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListCommandsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListCommandsPagesWithContext(ctx context.Context, input *ssm.ListCommandsInput, fn func(*ssm.ListCommandsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListCommandsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListCommandsOutput{}
	fnCacher := func(out *ssm.ListCommandsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListCommandsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListCommandsWithContext(ctx context.Context, input *ssm.ListCommandsInput, opts ...request.Option) (*ssm.ListCommandsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListCommandsOutput), nil
	}
	out, err := c.SSMAPI.ListCommandsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListComplianceItems(input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListComplianceItemsOutput), nil
	}
	out, err := c.SSMAPI.ListComplianceItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListComplianceItemsPages(input *ssm.ListComplianceItemsInput, fn func(*ssm.ListComplianceItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListComplianceItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListComplianceItemsOutput{}
	fnCacher := func(out *ssm.ListComplianceItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListComplianceItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListComplianceItemsPagesWithContext(ctx context.Context, input *ssm.ListComplianceItemsInput, fn func(*ssm.ListComplianceItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListComplianceItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListComplianceItemsOutput{}
	fnCacher := func(out *ssm.ListComplianceItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListComplianceItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListComplianceItemsWithContext(ctx context.Context, input *ssm.ListComplianceItemsInput, opts ...request.Option) (*ssm.ListComplianceItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListComplianceItemsOutput), nil
	}
	out, err := c.SSMAPI.ListComplianceItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListComplianceSummaries(input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListComplianceSummariesOutput), nil
	}
	out, err := c.SSMAPI.ListComplianceSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListComplianceSummariesPages(input *ssm.ListComplianceSummariesInput, fn func(*ssm.ListComplianceSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListComplianceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListComplianceSummariesOutput{}
	fnCacher := func(out *ssm.ListComplianceSummariesOutput, more bool) bool {
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
	if err := c.SSMAPI.ListComplianceSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListComplianceSummariesPagesWithContext(ctx context.Context, input *ssm.ListComplianceSummariesInput, fn func(*ssm.ListComplianceSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListComplianceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListComplianceSummariesOutput{}
	fnCacher := func(out *ssm.ListComplianceSummariesOutput, more bool) bool {
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
	if err := c.SSMAPI.ListComplianceSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListComplianceSummariesWithContext(ctx context.Context, input *ssm.ListComplianceSummariesInput, opts ...request.Option) (*ssm.ListComplianceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListComplianceSummariesOutput), nil
	}
	out, err := c.SSMAPI.ListComplianceSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocumentMetadataHistory(input *ssm.ListDocumentMetadataHistoryInput) (*ssm.ListDocumentMetadataHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentMetadataHistoryOutput), nil
	}
	out, err := c.SSMAPI.ListDocumentMetadataHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocumentMetadataHistoryWithContext(ctx context.Context, input *ssm.ListDocumentMetadataHistoryInput, opts ...request.Option) (*ssm.ListDocumentMetadataHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentMetadataHistoryOutput), nil
	}
	out, err := c.SSMAPI.ListDocumentMetadataHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocumentVersions(input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentVersionsOutput), nil
	}
	out, err := c.SSMAPI.ListDocumentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocumentVersionsPages(input *ssm.ListDocumentVersionsInput, fn func(*ssm.ListDocumentVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListDocumentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListDocumentVersionsOutput{}
	fnCacher := func(out *ssm.ListDocumentVersionsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListDocumentVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListDocumentVersionsPagesWithContext(ctx context.Context, input *ssm.ListDocumentVersionsInput, fn func(*ssm.ListDocumentVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListDocumentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListDocumentVersionsOutput{}
	fnCacher := func(out *ssm.ListDocumentVersionsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListDocumentVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListDocumentVersionsWithContext(ctx context.Context, input *ssm.ListDocumentVersionsInput, opts ...request.Option) (*ssm.ListDocumentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentVersionsOutput), nil
	}
	out, err := c.SSMAPI.ListDocumentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocuments(input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentsOutput), nil
	}
	out, err := c.SSMAPI.ListDocuments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListDocumentsPages(input *ssm.ListDocumentsInput, fn func(*ssm.ListDocumentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListDocumentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListDocumentsOutput{}
	fnCacher := func(out *ssm.ListDocumentsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListDocumentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListDocumentsPagesWithContext(ctx context.Context, input *ssm.ListDocumentsInput, fn func(*ssm.ListDocumentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListDocumentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListDocumentsOutput{}
	fnCacher := func(out *ssm.ListDocumentsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListDocumentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListDocumentsWithContext(ctx context.Context, input *ssm.ListDocumentsInput, opts ...request.Option) (*ssm.ListDocumentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListDocumentsOutput), nil
	}
	out, err := c.SSMAPI.ListDocumentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListInventoryEntries(input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListInventoryEntriesOutput), nil
	}
	out, err := c.SSMAPI.ListInventoryEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListInventoryEntriesWithContext(ctx context.Context, input *ssm.ListInventoryEntriesInput, opts ...request.Option) (*ssm.ListInventoryEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListInventoryEntriesOutput), nil
	}
	out, err := c.SSMAPI.ListInventoryEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsItemEvents(input *ssm.ListOpsItemEventsInput) (*ssm.ListOpsItemEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsItemEventsOutput), nil
	}
	out, err := c.SSMAPI.ListOpsItemEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsItemEventsPages(input *ssm.ListOpsItemEventsInput, fn func(*ssm.ListOpsItemEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsItemEventsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsItemEventsOutput{}
	fnCacher := func(out *ssm.ListOpsItemEventsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsItemEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsItemEventsPagesWithContext(ctx context.Context, input *ssm.ListOpsItemEventsInput, fn func(*ssm.ListOpsItemEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsItemEventsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsItemEventsOutput{}
	fnCacher := func(out *ssm.ListOpsItemEventsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsItemEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsItemEventsWithContext(ctx context.Context, input *ssm.ListOpsItemEventsInput, opts ...request.Option) (*ssm.ListOpsItemEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsItemEventsOutput), nil
	}
	out, err := c.SSMAPI.ListOpsItemEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsItemRelatedItems(input *ssm.ListOpsItemRelatedItemsInput) (*ssm.ListOpsItemRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsItemRelatedItemsOutput), nil
	}
	out, err := c.SSMAPI.ListOpsItemRelatedItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsItemRelatedItemsPages(input *ssm.ListOpsItemRelatedItemsInput, fn func(*ssm.ListOpsItemRelatedItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsItemRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsItemRelatedItemsOutput{}
	fnCacher := func(out *ssm.ListOpsItemRelatedItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsItemRelatedItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsItemRelatedItemsPagesWithContext(ctx context.Context, input *ssm.ListOpsItemRelatedItemsInput, fn func(*ssm.ListOpsItemRelatedItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsItemRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsItemRelatedItemsOutput{}
	fnCacher := func(out *ssm.ListOpsItemRelatedItemsOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsItemRelatedItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsItemRelatedItemsWithContext(ctx context.Context, input *ssm.ListOpsItemRelatedItemsInput, opts ...request.Option) (*ssm.ListOpsItemRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsItemRelatedItemsOutput), nil
	}
	out, err := c.SSMAPI.ListOpsItemRelatedItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsMetadata(input *ssm.ListOpsMetadataInput) (*ssm.ListOpsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsMetadataOutput), nil
	}
	out, err := c.SSMAPI.ListOpsMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListOpsMetadataPages(input *ssm.ListOpsMetadataInput, fn func(*ssm.ListOpsMetadataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsMetadataOutput{}
	fnCacher := func(out *ssm.ListOpsMetadataOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsMetadataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsMetadataPagesWithContext(ctx context.Context, input *ssm.ListOpsMetadataInput, fn func(*ssm.ListOpsMetadataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListOpsMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListOpsMetadataOutput{}
	fnCacher := func(out *ssm.ListOpsMetadataOutput, more bool) bool {
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
	if err := c.SSMAPI.ListOpsMetadataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListOpsMetadataWithContext(ctx context.Context, input *ssm.ListOpsMetadataInput, opts ...request.Option) (*ssm.ListOpsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListOpsMetadataOutput), nil
	}
	out, err := c.SSMAPI.ListOpsMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListResourceComplianceSummaries(input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListResourceComplianceSummariesOutput), nil
	}
	out, err := c.SSMAPI.ListResourceComplianceSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListResourceComplianceSummariesPages(input *ssm.ListResourceComplianceSummariesInput, fn func(*ssm.ListResourceComplianceSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListResourceComplianceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListResourceComplianceSummariesOutput{}
	fnCacher := func(out *ssm.ListResourceComplianceSummariesOutput, more bool) bool {
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
	if err := c.SSMAPI.ListResourceComplianceSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListResourceComplianceSummariesPagesWithContext(ctx context.Context, input *ssm.ListResourceComplianceSummariesInput, fn func(*ssm.ListResourceComplianceSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListResourceComplianceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListResourceComplianceSummariesOutput{}
	fnCacher := func(out *ssm.ListResourceComplianceSummariesOutput, more bool) bool {
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
	if err := c.SSMAPI.ListResourceComplianceSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListResourceComplianceSummariesWithContext(ctx context.Context, input *ssm.ListResourceComplianceSummariesInput, opts ...request.Option) (*ssm.ListResourceComplianceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListResourceComplianceSummariesOutput), nil
	}
	out, err := c.SSMAPI.ListResourceComplianceSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListResourceDataSync(input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListResourceDataSyncOutput), nil
	}
	out, err := c.SSMAPI.ListResourceDataSync(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListResourceDataSyncPages(input *ssm.ListResourceDataSyncInput, fn func(*ssm.ListResourceDataSyncOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListResourceDataSyncOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListResourceDataSyncOutput{}
	fnCacher := func(out *ssm.ListResourceDataSyncOutput, more bool) bool {
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
	if err := c.SSMAPI.ListResourceDataSyncPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListResourceDataSyncPagesWithContext(ctx context.Context, input *ssm.ListResourceDataSyncInput, fn func(*ssm.ListResourceDataSyncOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssm.ListResourceDataSyncOutput), false)
		return nil
	}
	cachable := true
	output := &ssm.ListResourceDataSyncOutput{}
	fnCacher := func(out *ssm.ListResourceDataSyncOutput, more bool) bool {
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
	if err := c.SSMAPI.ListResourceDataSyncPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSM) ListResourceDataSyncWithContext(ctx context.Context, input *ssm.ListResourceDataSyncInput, opts ...request.Option) (*ssm.ListResourceDataSyncOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListResourceDataSyncOutput), nil
	}
	out, err := c.SSMAPI.ListResourceDataSyncWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListTagsForResource(input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSM) ListTagsForResourceWithContext(ctx context.Context, input *ssm.ListTagsForResourceInput, opts ...request.Option) (*ssm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssm.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

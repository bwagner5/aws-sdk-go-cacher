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
	"github.com/aws/aws-sdk-go/service/drs"
	"github.com/aws/aws-sdk-go/service/drs/drsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Drs struct {
	drsiface.DrsAPI
	cache *cache.Cache
}

func NewDrs(drsapi drsiface.DrsAPI) *Drs {
	return &Drs{
		DrsAPI: drsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Drs) DescribeJobLogItems(input *drs.DescribeJobLogItemsInput) (*drs.DescribeJobLogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeJobLogItemsOutput), nil
	}
	out, err := c.DrsAPI.DescribeJobLogItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeJobLogItemsPages(input *drs.DescribeJobLogItemsInput, fn func(*drs.DescribeJobLogItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeJobLogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeJobLogItemsOutput{}
	fnCacher := func(out *drs.DescribeJobLogItemsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeJobLogItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeJobLogItemsPagesWithContext(ctx context.Context, input *drs.DescribeJobLogItemsInput, fn func(*drs.DescribeJobLogItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeJobLogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeJobLogItemsOutput{}
	fnCacher := func(out *drs.DescribeJobLogItemsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeJobLogItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeJobLogItemsWithContext(ctx context.Context, input *drs.DescribeJobLogItemsInput, opts ...request.Option) (*drs.DescribeJobLogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeJobLogItemsOutput), nil
	}
	out, err := c.DrsAPI.DescribeJobLogItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeJobs(input *drs.DescribeJobsInput) (*drs.DescribeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeJobsOutput), nil
	}
	out, err := c.DrsAPI.DescribeJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeJobsPages(input *drs.DescribeJobsInput, fn func(*drs.DescribeJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeJobsOutput{}
	fnCacher := func(out *drs.DescribeJobsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeJobsPagesWithContext(ctx context.Context, input *drs.DescribeJobsInput, fn func(*drs.DescribeJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeJobsOutput{}
	fnCacher := func(out *drs.DescribeJobsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeJobsWithContext(ctx context.Context, input *drs.DescribeJobsInput, opts ...request.Option) (*drs.DescribeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeJobsOutput), nil
	}
	out, err := c.DrsAPI.DescribeJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeRecoveryInstances(input *drs.DescribeRecoveryInstancesInput) (*drs.DescribeRecoveryInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeRecoveryInstancesOutput), nil
	}
	out, err := c.DrsAPI.DescribeRecoveryInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeRecoveryInstancesPages(input *drs.DescribeRecoveryInstancesInput, fn func(*drs.DescribeRecoveryInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeRecoveryInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeRecoveryInstancesOutput{}
	fnCacher := func(out *drs.DescribeRecoveryInstancesOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeRecoveryInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeRecoveryInstancesPagesWithContext(ctx context.Context, input *drs.DescribeRecoveryInstancesInput, fn func(*drs.DescribeRecoveryInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeRecoveryInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeRecoveryInstancesOutput{}
	fnCacher := func(out *drs.DescribeRecoveryInstancesOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeRecoveryInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeRecoveryInstancesWithContext(ctx context.Context, input *drs.DescribeRecoveryInstancesInput, opts ...request.Option) (*drs.DescribeRecoveryInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeRecoveryInstancesOutput), nil
	}
	out, err := c.DrsAPI.DescribeRecoveryInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeRecoverySnapshots(input *drs.DescribeRecoverySnapshotsInput) (*drs.DescribeRecoverySnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeRecoverySnapshotsOutput), nil
	}
	out, err := c.DrsAPI.DescribeRecoverySnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeRecoverySnapshotsPages(input *drs.DescribeRecoverySnapshotsInput, fn func(*drs.DescribeRecoverySnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeRecoverySnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeRecoverySnapshotsOutput{}
	fnCacher := func(out *drs.DescribeRecoverySnapshotsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeRecoverySnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeRecoverySnapshotsPagesWithContext(ctx context.Context, input *drs.DescribeRecoverySnapshotsInput, fn func(*drs.DescribeRecoverySnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeRecoverySnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeRecoverySnapshotsOutput{}
	fnCacher := func(out *drs.DescribeRecoverySnapshotsOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeRecoverySnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeRecoverySnapshotsWithContext(ctx context.Context, input *drs.DescribeRecoverySnapshotsInput, opts ...request.Option) (*drs.DescribeRecoverySnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeRecoverySnapshotsOutput), nil
	}
	out, err := c.DrsAPI.DescribeRecoverySnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeReplicationConfigurationTemplates(input *drs.DescribeReplicationConfigurationTemplatesInput) (*drs.DescribeReplicationConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeReplicationConfigurationTemplatesOutput), nil
	}
	out, err := c.DrsAPI.DescribeReplicationConfigurationTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeReplicationConfigurationTemplatesPages(input *drs.DescribeReplicationConfigurationTemplatesInput, fn func(*drs.DescribeReplicationConfigurationTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeReplicationConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeReplicationConfigurationTemplatesOutput{}
	fnCacher := func(out *drs.DescribeReplicationConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeReplicationConfigurationTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeReplicationConfigurationTemplatesPagesWithContext(ctx context.Context, input *drs.DescribeReplicationConfigurationTemplatesInput, fn func(*drs.DescribeReplicationConfigurationTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeReplicationConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeReplicationConfigurationTemplatesOutput{}
	fnCacher := func(out *drs.DescribeReplicationConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeReplicationConfigurationTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeReplicationConfigurationTemplatesWithContext(ctx context.Context, input *drs.DescribeReplicationConfigurationTemplatesInput, opts ...request.Option) (*drs.DescribeReplicationConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeReplicationConfigurationTemplatesOutput), nil
	}
	out, err := c.DrsAPI.DescribeReplicationConfigurationTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeSourceServers(input *drs.DescribeSourceServersInput) (*drs.DescribeSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeSourceServersOutput), nil
	}
	out, err := c.DrsAPI.DescribeSourceServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) DescribeSourceServersPages(input *drs.DescribeSourceServersInput, fn func(*drs.DescribeSourceServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeSourceServersOutput{}
	fnCacher := func(out *drs.DescribeSourceServersOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeSourceServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeSourceServersPagesWithContext(ctx context.Context, input *drs.DescribeSourceServersInput, fn func(*drs.DescribeSourceServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.DescribeSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &drs.DescribeSourceServersOutput{}
	fnCacher := func(out *drs.DescribeSourceServersOutput, more bool) bool {
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
	if err := c.DrsAPI.DescribeSourceServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) DescribeSourceServersWithContext(ctx context.Context, input *drs.DescribeSourceServersInput, opts ...request.Option) (*drs.DescribeSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.DescribeSourceServersOutput), nil
	}
	out, err := c.DrsAPI.DescribeSourceServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetFailbackReplicationConfiguration(input *drs.GetFailbackReplicationConfigurationInput) (*drs.GetFailbackReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetFailbackReplicationConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetFailbackReplicationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetFailbackReplicationConfigurationWithContext(ctx context.Context, input *drs.GetFailbackReplicationConfigurationInput, opts ...request.Option) (*drs.GetFailbackReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetFailbackReplicationConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetFailbackReplicationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetLaunchConfiguration(input *drs.GetLaunchConfigurationInput) (*drs.GetLaunchConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetLaunchConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetLaunchConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetLaunchConfigurationWithContext(ctx context.Context, input *drs.GetLaunchConfigurationInput, opts ...request.Option) (*drs.GetLaunchConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetLaunchConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetLaunchConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetReplicationConfiguration(input *drs.GetReplicationConfigurationInput) (*drs.GetReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetReplicationConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetReplicationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) GetReplicationConfigurationWithContext(ctx context.Context, input *drs.GetReplicationConfigurationInput, opts ...request.Option) (*drs.GetReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.GetReplicationConfigurationOutput), nil
	}
	out, err := c.DrsAPI.GetReplicationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListExtensibleSourceServers(input *drs.ListExtensibleSourceServersInput) (*drs.ListExtensibleSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListExtensibleSourceServersOutput), nil
	}
	out, err := c.DrsAPI.ListExtensibleSourceServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListExtensibleSourceServersPages(input *drs.ListExtensibleSourceServersInput, fn func(*drs.ListExtensibleSourceServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.ListExtensibleSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &drs.ListExtensibleSourceServersOutput{}
	fnCacher := func(out *drs.ListExtensibleSourceServersOutput, more bool) bool {
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
	if err := c.DrsAPI.ListExtensibleSourceServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) ListExtensibleSourceServersPagesWithContext(ctx context.Context, input *drs.ListExtensibleSourceServersInput, fn func(*drs.ListExtensibleSourceServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.ListExtensibleSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &drs.ListExtensibleSourceServersOutput{}
	fnCacher := func(out *drs.ListExtensibleSourceServersOutput, more bool) bool {
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
	if err := c.DrsAPI.ListExtensibleSourceServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) ListExtensibleSourceServersWithContext(ctx context.Context, input *drs.ListExtensibleSourceServersInput, opts ...request.Option) (*drs.ListExtensibleSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListExtensibleSourceServersOutput), nil
	}
	out, err := c.DrsAPI.ListExtensibleSourceServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListStagingAccounts(input *drs.ListStagingAccountsInput) (*drs.ListStagingAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListStagingAccountsOutput), nil
	}
	out, err := c.DrsAPI.ListStagingAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListStagingAccountsPages(input *drs.ListStagingAccountsInput, fn func(*drs.ListStagingAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.ListStagingAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.ListStagingAccountsOutput{}
	fnCacher := func(out *drs.ListStagingAccountsOutput, more bool) bool {
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
	if err := c.DrsAPI.ListStagingAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) ListStagingAccountsPagesWithContext(ctx context.Context, input *drs.ListStagingAccountsInput, fn func(*drs.ListStagingAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*drs.ListStagingAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &drs.ListStagingAccountsOutput{}
	fnCacher := func(out *drs.ListStagingAccountsOutput, more bool) bool {
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
	if err := c.DrsAPI.ListStagingAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Drs) ListStagingAccountsWithContext(ctx context.Context, input *drs.ListStagingAccountsInput, opts ...request.Option) (*drs.ListStagingAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListStagingAccountsOutput), nil
	}
	out, err := c.DrsAPI.ListStagingAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListTagsForResource(input *drs.ListTagsForResourceInput) (*drs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListTagsForResourceOutput), nil
	}
	out, err := c.DrsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Drs) ListTagsForResourceWithContext(ctx context.Context, input *drs.ListTagsForResourceInput, opts ...request.Option) (*drs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*drs.ListTagsForResourceOutput), nil
	}
	out, err := c.DrsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

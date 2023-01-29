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
package mgncacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mgn"
	"github.com/aws/aws-sdk-go/service/mgn/mgniface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Mgn struct {
	mgniface.MgnAPI
	cache *cache.Cache
}

func New(mgnapi mgniface.MgnAPI) *Mgn {
	return &Mgn{
		MgnAPI: mgnapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Mgn) DescribeJobLogItems(input *mgn.DescribeJobLogItemsInput) (*mgn.DescribeJobLogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeJobLogItemsOutput), nil
	}
	out, err := c.MgnAPI.DescribeJobLogItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeJobLogItemsPages(input *mgn.DescribeJobLogItemsInput, fn func(*mgn.DescribeJobLogItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeJobLogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeJobLogItemsOutput{}
	fnCacher := func(out *mgn.DescribeJobLogItemsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeJobLogItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeJobLogItemsPagesWithContext(ctx context.Context, input *mgn.DescribeJobLogItemsInput, fn func(*mgn.DescribeJobLogItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeJobLogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeJobLogItemsOutput{}
	fnCacher := func(out *mgn.DescribeJobLogItemsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeJobLogItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeJobLogItemsWithContext(ctx context.Context, input *mgn.DescribeJobLogItemsInput, opts ...request.Option) (*mgn.DescribeJobLogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeJobLogItemsOutput), nil
	}
	out, err := c.MgnAPI.DescribeJobLogItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeJobs(input *mgn.DescribeJobsInput) (*mgn.DescribeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeJobsOutput), nil
	}
	out, err := c.MgnAPI.DescribeJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeJobsPages(input *mgn.DescribeJobsInput, fn func(*mgn.DescribeJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeJobsOutput{}
	fnCacher := func(out *mgn.DescribeJobsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeJobsPagesWithContext(ctx context.Context, input *mgn.DescribeJobsInput, fn func(*mgn.DescribeJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeJobsOutput{}
	fnCacher := func(out *mgn.DescribeJobsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeJobsWithContext(ctx context.Context, input *mgn.DescribeJobsInput, opts ...request.Option) (*mgn.DescribeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeJobsOutput), nil
	}
	out, err := c.MgnAPI.DescribeJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeLaunchConfigurationTemplates(input *mgn.DescribeLaunchConfigurationTemplatesInput) (*mgn.DescribeLaunchConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeLaunchConfigurationTemplatesOutput), nil
	}
	out, err := c.MgnAPI.DescribeLaunchConfigurationTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeLaunchConfigurationTemplatesPages(input *mgn.DescribeLaunchConfigurationTemplatesInput, fn func(*mgn.DescribeLaunchConfigurationTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeLaunchConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeLaunchConfigurationTemplatesOutput{}
	fnCacher := func(out *mgn.DescribeLaunchConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeLaunchConfigurationTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeLaunchConfigurationTemplatesPagesWithContext(ctx context.Context, input *mgn.DescribeLaunchConfigurationTemplatesInput, fn func(*mgn.DescribeLaunchConfigurationTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeLaunchConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeLaunchConfigurationTemplatesOutput{}
	fnCacher := func(out *mgn.DescribeLaunchConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeLaunchConfigurationTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeLaunchConfigurationTemplatesWithContext(ctx context.Context, input *mgn.DescribeLaunchConfigurationTemplatesInput, opts ...request.Option) (*mgn.DescribeLaunchConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeLaunchConfigurationTemplatesOutput), nil
	}
	out, err := c.MgnAPI.DescribeLaunchConfigurationTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeReplicationConfigurationTemplates(input *mgn.DescribeReplicationConfigurationTemplatesInput) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeReplicationConfigurationTemplatesOutput), nil
	}
	out, err := c.MgnAPI.DescribeReplicationConfigurationTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeReplicationConfigurationTemplatesPages(input *mgn.DescribeReplicationConfigurationTemplatesInput, fn func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeReplicationConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeReplicationConfigurationTemplatesOutput{}
	fnCacher := func(out *mgn.DescribeReplicationConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeReplicationConfigurationTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeReplicationConfigurationTemplatesPagesWithContext(ctx context.Context, input *mgn.DescribeReplicationConfigurationTemplatesInput, fn func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeReplicationConfigurationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeReplicationConfigurationTemplatesOutput{}
	fnCacher := func(out *mgn.DescribeReplicationConfigurationTemplatesOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeReplicationConfigurationTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeReplicationConfigurationTemplatesWithContext(ctx context.Context, input *mgn.DescribeReplicationConfigurationTemplatesInput, opts ...request.Option) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeReplicationConfigurationTemplatesOutput), nil
	}
	out, err := c.MgnAPI.DescribeReplicationConfigurationTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeSourceServers(input *mgn.DescribeSourceServersInput) (*mgn.DescribeSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeSourceServersOutput), nil
	}
	out, err := c.MgnAPI.DescribeSourceServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeSourceServersPages(input *mgn.DescribeSourceServersInput, fn func(*mgn.DescribeSourceServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeSourceServersOutput{}
	fnCacher := func(out *mgn.DescribeSourceServersOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeSourceServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeSourceServersPagesWithContext(ctx context.Context, input *mgn.DescribeSourceServersInput, fn func(*mgn.DescribeSourceServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeSourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeSourceServersOutput{}
	fnCacher := func(out *mgn.DescribeSourceServersOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeSourceServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeSourceServersWithContext(ctx context.Context, input *mgn.DescribeSourceServersInput, opts ...request.Option) (*mgn.DescribeSourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeSourceServersOutput), nil
	}
	out, err := c.MgnAPI.DescribeSourceServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeVcenterClients(input *mgn.DescribeVcenterClientsInput) (*mgn.DescribeVcenterClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeVcenterClientsOutput), nil
	}
	out, err := c.MgnAPI.DescribeVcenterClients(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) DescribeVcenterClientsPages(input *mgn.DescribeVcenterClientsInput, fn func(*mgn.DescribeVcenterClientsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeVcenterClientsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeVcenterClientsOutput{}
	fnCacher := func(out *mgn.DescribeVcenterClientsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeVcenterClientsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeVcenterClientsPagesWithContext(ctx context.Context, input *mgn.DescribeVcenterClientsInput, fn func(*mgn.DescribeVcenterClientsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.DescribeVcenterClientsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.DescribeVcenterClientsOutput{}
	fnCacher := func(out *mgn.DescribeVcenterClientsOutput, more bool) bool {
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
	if err := c.MgnAPI.DescribeVcenterClientsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) DescribeVcenterClientsWithContext(ctx context.Context, input *mgn.DescribeVcenterClientsInput, opts ...request.Option) (*mgn.DescribeVcenterClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.DescribeVcenterClientsOutput), nil
	}
	out, err := c.MgnAPI.DescribeVcenterClientsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) GetLaunchConfiguration(input *mgn.GetLaunchConfigurationInput) (*mgn.GetLaunchConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.GetLaunchConfigurationOutput), nil
	}
	out, err := c.MgnAPI.GetLaunchConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) GetLaunchConfigurationWithContext(ctx context.Context, input *mgn.GetLaunchConfigurationInput, opts ...request.Option) (*mgn.GetLaunchConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.GetLaunchConfigurationOutput), nil
	}
	out, err := c.MgnAPI.GetLaunchConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) GetReplicationConfiguration(input *mgn.GetReplicationConfigurationInput) (*mgn.GetReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.GetReplicationConfigurationOutput), nil
	}
	out, err := c.MgnAPI.GetReplicationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) GetReplicationConfigurationWithContext(ctx context.Context, input *mgn.GetReplicationConfigurationInput, opts ...request.Option) (*mgn.GetReplicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.GetReplicationConfigurationOutput), nil
	}
	out, err := c.MgnAPI.GetReplicationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListApplications(input *mgn.ListApplicationsInput) (*mgn.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListApplicationsOutput), nil
	}
	out, err := c.MgnAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListApplicationsPages(input *mgn.ListApplicationsInput, fn func(*mgn.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListApplicationsOutput{}
	fnCacher := func(out *mgn.ListApplicationsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListApplicationsPagesWithContext(ctx context.Context, input *mgn.ListApplicationsInput, fn func(*mgn.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListApplicationsOutput{}
	fnCacher := func(out *mgn.ListApplicationsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListApplicationsWithContext(ctx context.Context, input *mgn.ListApplicationsInput, opts ...request.Option) (*mgn.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListApplicationsOutput), nil
	}
	out, err := c.MgnAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListSourceServerActions(input *mgn.ListSourceServerActionsInput) (*mgn.ListSourceServerActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListSourceServerActionsOutput), nil
	}
	out, err := c.MgnAPI.ListSourceServerActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListSourceServerActionsPages(input *mgn.ListSourceServerActionsInput, fn func(*mgn.ListSourceServerActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListSourceServerActionsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListSourceServerActionsOutput{}
	fnCacher := func(out *mgn.ListSourceServerActionsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListSourceServerActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListSourceServerActionsPagesWithContext(ctx context.Context, input *mgn.ListSourceServerActionsInput, fn func(*mgn.ListSourceServerActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListSourceServerActionsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListSourceServerActionsOutput{}
	fnCacher := func(out *mgn.ListSourceServerActionsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListSourceServerActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListSourceServerActionsWithContext(ctx context.Context, input *mgn.ListSourceServerActionsInput, opts ...request.Option) (*mgn.ListSourceServerActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListSourceServerActionsOutput), nil
	}
	out, err := c.MgnAPI.ListSourceServerActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListTagsForResource(input *mgn.ListTagsForResourceInput) (*mgn.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListTagsForResourceOutput), nil
	}
	out, err := c.MgnAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListTagsForResourceWithContext(ctx context.Context, input *mgn.ListTagsForResourceInput, opts ...request.Option) (*mgn.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListTagsForResourceOutput), nil
	}
	out, err := c.MgnAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListTemplateActions(input *mgn.ListTemplateActionsInput) (*mgn.ListTemplateActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListTemplateActionsOutput), nil
	}
	out, err := c.MgnAPI.ListTemplateActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListTemplateActionsPages(input *mgn.ListTemplateActionsInput, fn func(*mgn.ListTemplateActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListTemplateActionsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListTemplateActionsOutput{}
	fnCacher := func(out *mgn.ListTemplateActionsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListTemplateActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListTemplateActionsPagesWithContext(ctx context.Context, input *mgn.ListTemplateActionsInput, fn func(*mgn.ListTemplateActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListTemplateActionsOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListTemplateActionsOutput{}
	fnCacher := func(out *mgn.ListTemplateActionsOutput, more bool) bool {
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
	if err := c.MgnAPI.ListTemplateActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListTemplateActionsWithContext(ctx context.Context, input *mgn.ListTemplateActionsInput, opts ...request.Option) (*mgn.ListTemplateActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListTemplateActionsOutput), nil
	}
	out, err := c.MgnAPI.ListTemplateActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListWaves(input *mgn.ListWavesInput) (*mgn.ListWavesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListWavesOutput), nil
	}
	out, err := c.MgnAPI.ListWaves(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Mgn) ListWavesPages(input *mgn.ListWavesInput, fn func(*mgn.ListWavesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListWavesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListWavesOutput{}
	fnCacher := func(out *mgn.ListWavesOutput, more bool) bool {
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
	if err := c.MgnAPI.ListWavesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListWavesPagesWithContext(ctx context.Context, input *mgn.ListWavesInput, fn func(*mgn.ListWavesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mgn.ListWavesOutput), false)
		return nil
	}
	cachable := true
	output := &mgn.ListWavesOutput{}
	fnCacher := func(out *mgn.ListWavesOutput, more bool) bool {
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
	if err := c.MgnAPI.ListWavesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Mgn) ListWavesWithContext(ctx context.Context, input *mgn.ListWavesInput, opts ...request.Option) (*mgn.ListWavesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mgn.ListWavesOutput), nil
	}
	out, err := c.MgnAPI.ListWavesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

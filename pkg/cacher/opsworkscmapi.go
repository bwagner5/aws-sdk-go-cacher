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
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type OpsWorksCM struct {
	opsworkscmiface.OpsWorksCMAPI
	cache *cache.Cache
}

func NewOpsWorksCM(opsworkscmapi opsworkscmiface.OpsWorksCMAPI) *OpsWorksCM {
	return &OpsWorksCM{
		OpsWorksCMAPI: opsworkscmapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *OpsWorksCM) DescribeAccountAttributes(input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeAccountAttributesOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeAccountAttributesWithContext(ctx context.Context, input *opsworkscm.DescribeAccountAttributesInput, opts ...request.Option) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeAccountAttributesOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeBackups(input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeBackupsOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeBackupsPages(input *opsworkscm.DescribeBackupsInput, fn func(*opsworkscm.DescribeBackupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeBackupsOutput{}
	fnCacher := func(out *opsworkscm.DescribeBackupsOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeBackupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeBackupsPagesWithContext(ctx context.Context, input *opsworkscm.DescribeBackupsInput, fn func(*opsworkscm.DescribeBackupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeBackupsOutput{}
	fnCacher := func(out *opsworkscm.DescribeBackupsOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeBackupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeBackupsWithContext(ctx context.Context, input *opsworkscm.DescribeBackupsInput, opts ...request.Option) (*opsworkscm.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeBackupsOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeEvents(input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeEventsOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeEventsPages(input *opsworkscm.DescribeEventsInput, fn func(*opsworkscm.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeEventsOutput{}
	fnCacher := func(out *opsworkscm.DescribeEventsOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeEventsPagesWithContext(ctx context.Context, input *opsworkscm.DescribeEventsInput, fn func(*opsworkscm.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeEventsOutput{}
	fnCacher := func(out *opsworkscm.DescribeEventsOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeEventsWithContext(ctx context.Context, input *opsworkscm.DescribeEventsInput, opts ...request.Option) (*opsworkscm.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeEventsOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeNodeAssociationStatus(input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeNodeAssociationStatusOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeNodeAssociationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeNodeAssociationStatusWithContext(ctx context.Context, input *opsworkscm.DescribeNodeAssociationStatusInput, opts ...request.Option) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeNodeAssociationStatusOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeNodeAssociationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeServers(input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeServersOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) DescribeServersPages(input *opsworkscm.DescribeServersInput, fn func(*opsworkscm.DescribeServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeServersOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeServersOutput{}
	fnCacher := func(out *opsworkscm.DescribeServersOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeServersPagesWithContext(ctx context.Context, input *opsworkscm.DescribeServersInput, fn func(*opsworkscm.DescribeServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.DescribeServersOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.DescribeServersOutput{}
	fnCacher := func(out *opsworkscm.DescribeServersOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.DescribeServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) DescribeServersWithContext(ctx context.Context, input *opsworkscm.DescribeServersInput, opts ...request.Option) (*opsworkscm.DescribeServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.DescribeServersOutput), nil
	}
	out, err := c.OpsWorksCMAPI.DescribeServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) ListTagsForResource(input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.ListTagsForResourceOutput), nil
	}
	out, err := c.OpsWorksCMAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorksCM) ListTagsForResourcePages(input *opsworkscm.ListTagsForResourceInput, fn func(*opsworkscm.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.ListTagsForResourceOutput{}
	fnCacher := func(out *opsworkscm.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) ListTagsForResourcePagesWithContext(ctx context.Context, input *opsworkscm.ListTagsForResourceInput, fn func(*opsworkscm.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworkscm.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &opsworkscm.ListTagsForResourceOutput{}
	fnCacher := func(out *opsworkscm.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.OpsWorksCMAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorksCM) ListTagsForResourceWithContext(ctx context.Context, input *opsworkscm.ListTagsForResourceInput, opts ...request.Option) (*opsworkscm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworkscm.ListTagsForResourceOutput), nil
	}
	out, err := c.OpsWorksCMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

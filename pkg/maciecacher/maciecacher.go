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
package maciecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/macie/macieiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Macie struct {
	macieiface.MacieAPI
	cache *cache.Cache
}

func New(macieapi macieiface.MacieAPI) *Macie {
	return &Macie{
		MacieAPI: macieapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Macie) ListMemberAccounts(input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie.ListMemberAccountsOutput), nil
	}
	out, err := c.MacieAPI.ListMemberAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie) ListMemberAccountsPages(input *macie.ListMemberAccountsInput, fn func(*macie.ListMemberAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie.ListMemberAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &macie.ListMemberAccountsOutput{}
	fnCacher := func(out *macie.ListMemberAccountsOutput, more bool) bool {
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
	if err := c.MacieAPI.ListMemberAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie) ListMemberAccountsPagesWithContext(ctx context.Context, input *macie.ListMemberAccountsInput, fn func(*macie.ListMemberAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie.ListMemberAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &macie.ListMemberAccountsOutput{}
	fnCacher := func(out *macie.ListMemberAccountsOutput, more bool) bool {
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
	if err := c.MacieAPI.ListMemberAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie) ListMemberAccountsWithContext(ctx context.Context, input *macie.ListMemberAccountsInput, opts ...request.Option) (*macie.ListMemberAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie.ListMemberAccountsOutput), nil
	}
	out, err := c.MacieAPI.ListMemberAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie) ListS3Resources(input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie.ListS3ResourcesOutput), nil
	}
	out, err := c.MacieAPI.ListS3Resources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie) ListS3ResourcesPages(input *macie.ListS3ResourcesInput, fn func(*macie.ListS3ResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie.ListS3ResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &macie.ListS3ResourcesOutput{}
	fnCacher := func(out *macie.ListS3ResourcesOutput, more bool) bool {
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
	if err := c.MacieAPI.ListS3ResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie) ListS3ResourcesPagesWithContext(ctx context.Context, input *macie.ListS3ResourcesInput, fn func(*macie.ListS3ResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie.ListS3ResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &macie.ListS3ResourcesOutput{}
	fnCacher := func(out *macie.ListS3ResourcesOutput, more bool) bool {
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
	if err := c.MacieAPI.ListS3ResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie) ListS3ResourcesWithContext(ctx context.Context, input *macie.ListS3ResourcesInput, opts ...request.Option) (*macie.ListS3ResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie.ListS3ResourcesOutput), nil
	}
	out, err := c.MacieAPI.ListS3ResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

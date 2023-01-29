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
package connectcasescacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/connectcases"
	"github.com/aws/aws-sdk-go/service/connectcases/connectcasesiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConnectCases struct {
	connectcasesiface.ConnectCasesAPI
	cache *cache.Cache
}

func New(connectcasesapi connectcasesiface.ConnectCasesAPI) *ConnectCases {
	return &ConnectCases{
		ConnectCasesAPI: connectcasesapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConnectCases) BatchGetField(input *connectcases.BatchGetFieldInput) (*connectcases.BatchGetFieldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.BatchGetFieldOutput), nil
	}
	out, err := c.ConnectCasesAPI.BatchGetField(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) BatchGetFieldWithContext(ctx context.Context, input *connectcases.BatchGetFieldInput, opts ...request.Option) (*connectcases.BatchGetFieldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.BatchGetFieldOutput), nil
	}
	out, err := c.ConnectCasesAPI.BatchGetFieldWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) BatchPutFieldOptions(input *connectcases.BatchPutFieldOptionsInput) (*connectcases.BatchPutFieldOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.BatchPutFieldOptionsOutput), nil
	}
	out, err := c.ConnectCasesAPI.BatchPutFieldOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) BatchPutFieldOptionsWithContext(ctx context.Context, input *connectcases.BatchPutFieldOptionsInput, opts ...request.Option) (*connectcases.BatchPutFieldOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.BatchPutFieldOptionsOutput), nil
	}
	out, err := c.ConnectCasesAPI.BatchPutFieldOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetCase(input *connectcases.GetCaseInput) (*connectcases.GetCaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetCaseOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetCase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetCaseEventConfiguration(input *connectcases.GetCaseEventConfigurationInput) (*connectcases.GetCaseEventConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetCaseEventConfigurationOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetCaseEventConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetCaseEventConfigurationWithContext(ctx context.Context, input *connectcases.GetCaseEventConfigurationInput, opts ...request.Option) (*connectcases.GetCaseEventConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetCaseEventConfigurationOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetCaseEventConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetCasePages(input *connectcases.GetCaseInput, fn func(*connectcases.GetCaseOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.GetCaseOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.GetCaseOutput{}
	fnCacher := func(out *connectcases.GetCaseOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.GetCasePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) GetCasePagesWithContext(ctx context.Context, input *connectcases.GetCaseInput, fn func(*connectcases.GetCaseOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.GetCaseOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.GetCaseOutput{}
	fnCacher := func(out *connectcases.GetCaseOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.GetCasePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) GetCaseWithContext(ctx context.Context, input *connectcases.GetCaseInput, opts ...request.Option) (*connectcases.GetCaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetCaseOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetCaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetDomain(input *connectcases.GetDomainInput) (*connectcases.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetDomainOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetDomainWithContext(ctx context.Context, input *connectcases.GetDomainInput, opts ...request.Option) (*connectcases.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetDomainOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetLayout(input *connectcases.GetLayoutInput) (*connectcases.GetLayoutOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetLayoutOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetLayout(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetLayoutWithContext(ctx context.Context, input *connectcases.GetLayoutInput, opts ...request.Option) (*connectcases.GetLayoutOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetLayoutOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetLayoutWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetTemplate(input *connectcases.GetTemplateInput) (*connectcases.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetTemplateOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) GetTemplateWithContext(ctx context.Context, input *connectcases.GetTemplateInput, opts ...request.Option) (*connectcases.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.GetTemplateOutput), nil
	}
	out, err := c.ConnectCasesAPI.GetTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListCasesForContact(input *connectcases.ListCasesForContactInput) (*connectcases.ListCasesForContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListCasesForContactOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListCasesForContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListCasesForContactPages(input *connectcases.ListCasesForContactInput, fn func(*connectcases.ListCasesForContactOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListCasesForContactOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListCasesForContactOutput{}
	fnCacher := func(out *connectcases.ListCasesForContactOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListCasesForContactPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListCasesForContactPagesWithContext(ctx context.Context, input *connectcases.ListCasesForContactInput, fn func(*connectcases.ListCasesForContactOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListCasesForContactOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListCasesForContactOutput{}
	fnCacher := func(out *connectcases.ListCasesForContactOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListCasesForContactPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListCasesForContactWithContext(ctx context.Context, input *connectcases.ListCasesForContactInput, opts ...request.Option) (*connectcases.ListCasesForContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListCasesForContactOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListCasesForContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListDomains(input *connectcases.ListDomainsInput) (*connectcases.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListDomainsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListDomainsPages(input *connectcases.ListDomainsInput, fn func(*connectcases.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListDomainsOutput{}
	fnCacher := func(out *connectcases.ListDomainsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListDomainsPagesWithContext(ctx context.Context, input *connectcases.ListDomainsInput, fn func(*connectcases.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListDomainsOutput{}
	fnCacher := func(out *connectcases.ListDomainsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListDomainsWithContext(ctx context.Context, input *connectcases.ListDomainsInput, opts ...request.Option) (*connectcases.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListDomainsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListFieldOptions(input *connectcases.ListFieldOptionsInput) (*connectcases.ListFieldOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListFieldOptionsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListFieldOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListFieldOptionsPages(input *connectcases.ListFieldOptionsInput, fn func(*connectcases.ListFieldOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListFieldOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListFieldOptionsOutput{}
	fnCacher := func(out *connectcases.ListFieldOptionsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListFieldOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListFieldOptionsPagesWithContext(ctx context.Context, input *connectcases.ListFieldOptionsInput, fn func(*connectcases.ListFieldOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListFieldOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListFieldOptionsOutput{}
	fnCacher := func(out *connectcases.ListFieldOptionsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListFieldOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListFieldOptionsWithContext(ctx context.Context, input *connectcases.ListFieldOptionsInput, opts ...request.Option) (*connectcases.ListFieldOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListFieldOptionsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListFieldOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListFields(input *connectcases.ListFieldsInput) (*connectcases.ListFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListFieldsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListFields(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListFieldsPages(input *connectcases.ListFieldsInput, fn func(*connectcases.ListFieldsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListFieldsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListFieldsOutput{}
	fnCacher := func(out *connectcases.ListFieldsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListFieldsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListFieldsPagesWithContext(ctx context.Context, input *connectcases.ListFieldsInput, fn func(*connectcases.ListFieldsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListFieldsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListFieldsOutput{}
	fnCacher := func(out *connectcases.ListFieldsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListFieldsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListFieldsWithContext(ctx context.Context, input *connectcases.ListFieldsInput, opts ...request.Option) (*connectcases.ListFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListFieldsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListFieldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListLayouts(input *connectcases.ListLayoutsInput) (*connectcases.ListLayoutsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListLayoutsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListLayouts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListLayoutsPages(input *connectcases.ListLayoutsInput, fn func(*connectcases.ListLayoutsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListLayoutsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListLayoutsOutput{}
	fnCacher := func(out *connectcases.ListLayoutsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListLayoutsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListLayoutsPagesWithContext(ctx context.Context, input *connectcases.ListLayoutsInput, fn func(*connectcases.ListLayoutsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListLayoutsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListLayoutsOutput{}
	fnCacher := func(out *connectcases.ListLayoutsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListLayoutsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListLayoutsWithContext(ctx context.Context, input *connectcases.ListLayoutsInput, opts ...request.Option) (*connectcases.ListLayoutsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListLayoutsOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListLayoutsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListTagsForResource(input *connectcases.ListTagsForResourceInput) (*connectcases.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListTagsForResourceWithContext(ctx context.Context, input *connectcases.ListTagsForResourceInput, opts ...request.Option) (*connectcases.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListTemplates(input *connectcases.ListTemplatesInput) (*connectcases.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListTemplatesOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) ListTemplatesPages(input *connectcases.ListTemplatesInput, fn func(*connectcases.ListTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListTemplatesOutput{}
	fnCacher := func(out *connectcases.ListTemplatesOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListTemplatesPagesWithContext(ctx context.Context, input *connectcases.ListTemplatesInput, fn func(*connectcases.ListTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.ListTemplatesOutput{}
	fnCacher := func(out *connectcases.ListTemplatesOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.ListTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) ListTemplatesWithContext(ctx context.Context, input *connectcases.ListTemplatesInput, opts ...request.Option) (*connectcases.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.ListTemplatesOutput), nil
	}
	out, err := c.ConnectCasesAPI.ListTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) SearchCases(input *connectcases.SearchCasesInput) (*connectcases.SearchCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.SearchCasesOutput), nil
	}
	out, err := c.ConnectCasesAPI.SearchCases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) SearchCasesPages(input *connectcases.SearchCasesInput, fn func(*connectcases.SearchCasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.SearchCasesOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.SearchCasesOutput{}
	fnCacher := func(out *connectcases.SearchCasesOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.SearchCasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) SearchCasesPagesWithContext(ctx context.Context, input *connectcases.SearchCasesInput, fn func(*connectcases.SearchCasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.SearchCasesOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.SearchCasesOutput{}
	fnCacher := func(out *connectcases.SearchCasesOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.SearchCasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) SearchCasesWithContext(ctx context.Context, input *connectcases.SearchCasesInput, opts ...request.Option) (*connectcases.SearchCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.SearchCasesOutput), nil
	}
	out, err := c.ConnectCasesAPI.SearchCasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) SearchRelatedItems(input *connectcases.SearchRelatedItemsInput) (*connectcases.SearchRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.SearchRelatedItemsOutput), nil
	}
	out, err := c.ConnectCasesAPI.SearchRelatedItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCases) SearchRelatedItemsPages(input *connectcases.SearchRelatedItemsInput, fn func(*connectcases.SearchRelatedItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.SearchRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.SearchRelatedItemsOutput{}
	fnCacher := func(out *connectcases.SearchRelatedItemsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.SearchRelatedItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) SearchRelatedItemsPagesWithContext(ctx context.Context, input *connectcases.SearchRelatedItemsInput, fn func(*connectcases.SearchRelatedItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcases.SearchRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcases.SearchRelatedItemsOutput{}
	fnCacher := func(out *connectcases.SearchRelatedItemsOutput, more bool) bool {
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
	if err := c.ConnectCasesAPI.SearchRelatedItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCases) SearchRelatedItemsWithContext(ctx context.Context, input *connectcases.SearchRelatedItemsInput, opts ...request.Option) (*connectcases.SearchRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcases.SearchRelatedItemsOutput), nil
	}
	out, err := c.ConnectCasesAPI.SearchRelatedItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

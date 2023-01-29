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
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ServiceQuotas struct {
	servicequotasiface.ServiceQuotasAPI
	cache *cache.Cache
}

func NewServiceQuotas(servicequotasapi servicequotasiface.ServiceQuotasAPI) *ServiceQuotas {
	return &ServiceQuotas{
		ServiceQuotasAPI: servicequotasapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ServiceQuotas) GetAWSDefaultServiceQuota(input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetAWSDefaultServiceQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetAWSDefaultServiceQuota(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetAWSDefaultServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput, opts ...request.Option) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetAWSDefaultServiceQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetAWSDefaultServiceQuotaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetAssociationForServiceQuotaTemplate(input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetAssociationForServiceQuotaTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetAssociationForServiceQuotaTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetAssociationForServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetAssociationForServiceQuotaTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetAssociationForServiceQuotaTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetRequestedServiceQuotaChange(input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetRequestedServiceQuotaChangeOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetRequestedServiceQuotaChange(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetRequestedServiceQuotaChangeWithContext(ctx context.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput, opts ...request.Option) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetRequestedServiceQuotaChangeOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetRequestedServiceQuotaChangeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetServiceQuota(input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetServiceQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetServiceQuota(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetServiceQuotaIncreaseRequestFromTemplate(input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetServiceQuotaIncreaseRequestFromTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, opts ...request.Option) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) GetServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaInput, opts ...request.Option) (*servicequotas.GetServiceQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.GetServiceQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.GetServiceQuotaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListAWSDefaultServiceQuotas(input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListAWSDefaultServiceQuotasOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListAWSDefaultServiceQuotas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListAWSDefaultServiceQuotasPages(input *servicequotas.ListAWSDefaultServiceQuotasInput, fn func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListAWSDefaultServiceQuotasOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListAWSDefaultServiceQuotasOutput{}
	fnCacher := func(out *servicequotas.ListAWSDefaultServiceQuotasOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListAWSDefaultServiceQuotasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListAWSDefaultServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, fn func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListAWSDefaultServiceQuotasOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListAWSDefaultServiceQuotasOutput{}
	fnCacher := func(out *servicequotas.ListAWSDefaultServiceQuotasOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListAWSDefaultServiceQuotasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListAWSDefaultServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, opts ...request.Option) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListAWSDefaultServiceQuotasOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListAWSDefaultServiceQuotasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistory(input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryByQuota(input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuota(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryByQuotaPages(input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, fn func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput{}
	fnCacher := func(out *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuotaPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, fn func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput{}
	fnCacher := func(out *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryPages(input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, fn func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListRequestedServiceQuotaChangeHistoryOutput{}
	fnCacher := func(out *servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, fn func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListRequestedServiceQuotaChangeHistoryOutput{}
	fnCacher := func(out *servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListRequestedServiceQuotaChangeHistoryWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServiceQuotaIncreaseRequestsInTemplate(input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServiceQuotaIncreaseRequestsInTemplatePages(input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, fn func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput{}
	fnCacher := func(out *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplatePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, fn func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput{}
	fnCacher := func(out *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, opts ...request.Option) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServiceQuotas(input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServiceQuotasOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServiceQuotas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServiceQuotasPages(input *servicequotas.ListServiceQuotasInput, fn func(*servicequotas.ListServiceQuotasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServiceQuotasOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServiceQuotasOutput{}
	fnCacher := func(out *servicequotas.ListServiceQuotasOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServiceQuotasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, fn func(*servicequotas.ListServiceQuotasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServiceQuotasOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServiceQuotasOutput{}
	fnCacher := func(out *servicequotas.ListServiceQuotasOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServiceQuotasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, opts ...request.Option) (*servicequotas.ListServiceQuotasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServiceQuotasOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServiceQuotasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServices(input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServicesOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListServicesPages(input *servicequotas.ListServicesInput, fn func(*servicequotas.ListServicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServicesOutput{}
	fnCacher := func(out *servicequotas.ListServicesOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServicesPagesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, fn func(*servicequotas.ListServicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicequotas.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &servicequotas.ListServicesOutput{}
	fnCacher := func(out *servicequotas.ListServicesOutput, more bool) bool {
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
	if err := c.ServiceQuotasAPI.ListServicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceQuotas) ListServicesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, opts ...request.Option) (*servicequotas.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListServicesOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListTagsForResource(input *servicequotas.ListTagsForResourceInput) (*servicequotas.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListTagsForResourceOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceQuotas) ListTagsForResourceWithContext(ctx context.Context, input *servicequotas.ListTagsForResourceInput, opts ...request.Option) (*servicequotas.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicequotas.ListTagsForResourceOutput), nil
	}
	out, err := c.ServiceQuotasAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

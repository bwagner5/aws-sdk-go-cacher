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
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ACMPCA struct {
	acmpcaiface.ACMPCAAPI
	cache *cache.Cache
}

func NewACMPCA(acmpcaapi acmpcaiface.ACMPCAAPI) *ACMPCA {
	return &ACMPCA{
		ACMPCAAPI: acmpcaapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ACMPCA) DescribeCertificateAuthority(input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.DescribeCertificateAuthorityOutput), nil
	}
	out, err := c.ACMPCAAPI.DescribeCertificateAuthority(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) DescribeCertificateAuthorityAuditReport(input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.DescribeCertificateAuthorityAuditReportOutput), nil
	}
	out, err := c.ACMPCAAPI.DescribeCertificateAuthorityAuditReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) DescribeCertificateAuthorityAuditReportWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.DescribeCertificateAuthorityAuditReportOutput), nil
	}
	out, err := c.ACMPCAAPI.DescribeCertificateAuthorityAuditReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) DescribeCertificateAuthorityWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.DescribeCertificateAuthorityOutput), nil
	}
	out, err := c.ACMPCAAPI.DescribeCertificateAuthorityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificate(input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificateAuthorityCertificate(input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateAuthorityCertificateOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificateAuthorityCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificateAuthorityCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCertificateInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateAuthorityCertificateOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificateAuthorityCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificateAuthorityCsr(input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateAuthorityCsrOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificateAuthorityCsr(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificateAuthorityCsrWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateAuthorityCsrOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificateAuthorityCsrWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateInput, opts ...request.Option) (*acmpca.GetCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetCertificateOutput), nil
	}
	out, err := c.ACMPCAAPI.GetCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetPolicy(input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetPolicyOutput), nil
	}
	out, err := c.ACMPCAAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) GetPolicyWithContext(ctx context.Context, input *acmpca.GetPolicyInput, opts ...request.Option) (*acmpca.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.GetPolicyOutput), nil
	}
	out, err := c.ACMPCAAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListCertificateAuthorities(input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListCertificateAuthoritiesOutput), nil
	}
	out, err := c.ACMPCAAPI.ListCertificateAuthorities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListCertificateAuthoritiesPages(input *acmpca.ListCertificateAuthoritiesInput, fn func(*acmpca.ListCertificateAuthoritiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListCertificateAuthoritiesOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListCertificateAuthoritiesOutput{}
	fnCacher := func(out *acmpca.ListCertificateAuthoritiesOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListCertificateAuthoritiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListCertificateAuthoritiesPagesWithContext(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput, fn func(*acmpca.ListCertificateAuthoritiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListCertificateAuthoritiesOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListCertificateAuthoritiesOutput{}
	fnCacher := func(out *acmpca.ListCertificateAuthoritiesOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListCertificateAuthoritiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListCertificateAuthoritiesWithContext(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput, opts ...request.Option) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListCertificateAuthoritiesOutput), nil
	}
	out, err := c.ACMPCAAPI.ListCertificateAuthoritiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListPermissions(input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListPermissionsOutput), nil
	}
	out, err := c.ACMPCAAPI.ListPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListPermissionsPages(input *acmpca.ListPermissionsInput, fn func(*acmpca.ListPermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListPermissionsOutput{}
	fnCacher := func(out *acmpca.ListPermissionsOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListPermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListPermissionsPagesWithContext(ctx context.Context, input *acmpca.ListPermissionsInput, fn func(*acmpca.ListPermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListPermissionsOutput{}
	fnCacher := func(out *acmpca.ListPermissionsOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListPermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListPermissionsWithContext(ctx context.Context, input *acmpca.ListPermissionsInput, opts ...request.Option) (*acmpca.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListPermissionsOutput), nil
	}
	out, err := c.ACMPCAAPI.ListPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListTags(input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListTagsOutput), nil
	}
	out, err := c.ACMPCAAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ACMPCA) ListTagsPages(input *acmpca.ListTagsInput, fn func(*acmpca.ListTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListTagsOutput{}
	fnCacher := func(out *acmpca.ListTagsOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListTagsPagesWithContext(ctx context.Context, input *acmpca.ListTagsInput, fn func(*acmpca.ListTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*acmpca.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &acmpca.ListTagsOutput{}
	fnCacher := func(out *acmpca.ListTagsOutput, more bool) bool {
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
	if err := c.ACMPCAAPI.ListTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ACMPCA) ListTagsWithContext(ctx context.Context, input *acmpca.ListTagsInput, opts ...request.Option) (*acmpca.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*acmpca.ListTagsOutput), nil
	}
	out, err := c.ACMPCAAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

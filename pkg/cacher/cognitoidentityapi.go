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
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CognitoIdentity struct {
	cognitoidentityiface.CognitoIdentityAPI
	cache *cache.Cache
}

func NewCognitoIdentity(cognitoidentityapi cognitoidentityiface.CognitoIdentityAPI) *CognitoIdentity {
	return &CognitoIdentity{
		CognitoIdentityAPI: cognitoidentityapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CognitoIdentity) GetCredentialsForIdentity(input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetCredentialsForIdentityOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetCredentialsForIdentity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetCredentialsForIdentityWithContext(ctx context.Context, input *cognitoidentity.GetCredentialsForIdentityInput, opts ...request.Option) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetCredentialsForIdentityOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetCredentialsForIdentityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetId(input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetIdOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetIdWithContext(ctx context.Context, input *cognitoidentity.GetIdInput, opts ...request.Option) (*cognitoidentity.GetIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetIdOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetIdentityPoolRoles(input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetIdentityPoolRolesOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetIdentityPoolRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetIdentityPoolRolesWithContext(ctx context.Context, input *cognitoidentity.GetIdentityPoolRolesInput, opts ...request.Option) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetIdentityPoolRolesOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetIdentityPoolRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetOpenIdToken(input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetOpenIdTokenOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetOpenIdToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetOpenIdTokenForDeveloperIdentity(input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetOpenIdTokenForDeveloperIdentity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetOpenIdTokenForDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetOpenIdTokenForDeveloperIdentityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetOpenIdTokenWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetOpenIdTokenOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetOpenIdTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetPrincipalTagAttributeMap(input *cognitoidentity.GetPrincipalTagAttributeMapInput) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetPrincipalTagAttributeMapOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetPrincipalTagAttributeMap(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) GetPrincipalTagAttributeMapWithContext(ctx context.Context, input *cognitoidentity.GetPrincipalTagAttributeMapInput, opts ...request.Option) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.GetPrincipalTagAttributeMapOutput), nil
	}
	out, err := c.CognitoIdentityAPI.GetPrincipalTagAttributeMapWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListIdentities(input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListIdentitiesOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListIdentitiesWithContext(ctx context.Context, input *cognitoidentity.ListIdentitiesInput, opts ...request.Option) (*cognitoidentity.ListIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListIdentitiesOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListIdentityPools(input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListIdentityPoolsOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListIdentityPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListIdentityPoolsPages(input *cognitoidentity.ListIdentityPoolsInput, fn func(*cognitoidentity.ListIdentityPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentity.ListIdentityPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentity.ListIdentityPoolsOutput{}
	fnCacher := func(out *cognitoidentity.ListIdentityPoolsOutput, more bool) bool {
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
	if err := c.CognitoIdentityAPI.ListIdentityPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentity) ListIdentityPoolsPagesWithContext(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput, fn func(*cognitoidentity.ListIdentityPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentity.ListIdentityPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentity.ListIdentityPoolsOutput{}
	fnCacher := func(out *cognitoidentity.ListIdentityPoolsOutput, more bool) bool {
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
	if err := c.CognitoIdentityAPI.ListIdentityPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentity) ListIdentityPoolsWithContext(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput, opts ...request.Option) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListIdentityPoolsOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListIdentityPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListTagsForResource(input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListTagsForResourceOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentity) ListTagsForResourceWithContext(ctx context.Context, input *cognitoidentity.ListTagsForResourceInput, opts ...request.Option) (*cognitoidentity.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentity.ListTagsForResourceOutput), nil
	}
	out, err := c.CognitoIdentityAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

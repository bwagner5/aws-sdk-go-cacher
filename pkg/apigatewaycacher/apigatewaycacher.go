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
package apigatewaycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type APIGateway struct {
	apigatewayiface.APIGatewayAPI
	cache *cache.Cache
}

func New(apigatewayapi apigatewayiface.APIGatewayAPI) *APIGateway {
	return &APIGateway{
		APIGatewayAPI: apigatewayapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *APIGateway) GetApiKeys(input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetApiKeysOutput), nil
	}
	out, err := c.APIGatewayAPI.GetApiKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetApiKeysPages(input *apigateway.GetApiKeysInput, fn func(*apigateway.GetApiKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetApiKeysOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetApiKeysOutput{}
	fnCacher := func(out *apigateway.GetApiKeysOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetApiKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetApiKeysPagesWithContext(ctx context.Context, input *apigateway.GetApiKeysInput, fn func(*apigateway.GetApiKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetApiKeysOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetApiKeysOutput{}
	fnCacher := func(out *apigateway.GetApiKeysOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetApiKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetApiKeysWithContext(ctx context.Context, input *apigateway.GetApiKeysInput, opts ...request.Option) (*apigateway.GetApiKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetApiKeysOutput), nil
	}
	out, err := c.APIGatewayAPI.GetApiKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetAuthorizers(input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetAuthorizersOutput), nil
	}
	out, err := c.APIGatewayAPI.GetAuthorizers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetAuthorizersWithContext(ctx context.Context, input *apigateway.GetAuthorizersInput, opts ...request.Option) (*apigateway.GetAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetAuthorizersOutput), nil
	}
	out, err := c.APIGatewayAPI.GetAuthorizersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetBasePathMappings(input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetBasePathMappingsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetBasePathMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetBasePathMappingsPages(input *apigateway.GetBasePathMappingsInput, fn func(*apigateway.GetBasePathMappingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetBasePathMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetBasePathMappingsOutput{}
	fnCacher := func(out *apigateway.GetBasePathMappingsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetBasePathMappingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetBasePathMappingsPagesWithContext(ctx context.Context, input *apigateway.GetBasePathMappingsInput, fn func(*apigateway.GetBasePathMappingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetBasePathMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetBasePathMappingsOutput{}
	fnCacher := func(out *apigateway.GetBasePathMappingsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetBasePathMappingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetBasePathMappingsWithContext(ctx context.Context, input *apigateway.GetBasePathMappingsInput, opts ...request.Option) (*apigateway.GetBasePathMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetBasePathMappingsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetBasePathMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetClientCertificates(input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetClientCertificatesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetClientCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetClientCertificatesPages(input *apigateway.GetClientCertificatesInput, fn func(*apigateway.GetClientCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetClientCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetClientCertificatesOutput{}
	fnCacher := func(out *apigateway.GetClientCertificatesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetClientCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetClientCertificatesPagesWithContext(ctx context.Context, input *apigateway.GetClientCertificatesInput, fn func(*apigateway.GetClientCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetClientCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetClientCertificatesOutput{}
	fnCacher := func(out *apigateway.GetClientCertificatesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetClientCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetClientCertificatesWithContext(ctx context.Context, input *apigateway.GetClientCertificatesInput, opts ...request.Option) (*apigateway.GetClientCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetClientCertificatesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetClientCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDeployments(input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDeploymentsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDeploymentsPages(input *apigateway.GetDeploymentsInput, fn func(*apigateway.GetDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetDeploymentsOutput{}
	fnCacher := func(out *apigateway.GetDeploymentsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetDeploymentsPagesWithContext(ctx context.Context, input *apigateway.GetDeploymentsInput, fn func(*apigateway.GetDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetDeploymentsOutput{}
	fnCacher := func(out *apigateway.GetDeploymentsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetDeploymentsWithContext(ctx context.Context, input *apigateway.GetDeploymentsInput, opts ...request.Option) (*apigateway.GetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDeploymentsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDocumentationParts(input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDocumentationPartsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDocumentationParts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDocumentationPartsWithContext(ctx context.Context, input *apigateway.GetDocumentationPartsInput, opts ...request.Option) (*apigateway.GetDocumentationPartsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDocumentationPartsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDocumentationPartsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDocumentationVersions(input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDocumentationVersionsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDocumentationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDocumentationVersionsWithContext(ctx context.Context, input *apigateway.GetDocumentationVersionsInput, opts ...request.Option) (*apigateway.GetDocumentationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDocumentationVersionsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDocumentationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDomainNames(input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDomainNamesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetDomainNamesPages(input *apigateway.GetDomainNamesInput, fn func(*apigateway.GetDomainNamesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetDomainNamesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetDomainNamesOutput{}
	fnCacher := func(out *apigateway.GetDomainNamesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetDomainNamesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetDomainNamesPagesWithContext(ctx context.Context, input *apigateway.GetDomainNamesInput, fn func(*apigateway.GetDomainNamesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetDomainNamesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetDomainNamesOutput{}
	fnCacher := func(out *apigateway.GetDomainNamesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetDomainNamesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetDomainNamesWithContext(ctx context.Context, input *apigateway.GetDomainNamesInput, opts ...request.Option) (*apigateway.GetDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetDomainNamesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetExport(input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetExportOutput), nil
	}
	out, err := c.APIGatewayAPI.GetExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetExportWithContext(ctx context.Context, input *apigateway.GetExportInput, opts ...request.Option) (*apigateway.GetExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetExportOutput), nil
	}
	out, err := c.APIGatewayAPI.GetExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetGatewayResponse(input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateGatewayResponseOutput), nil
	}
	out, err := c.APIGatewayAPI.GetGatewayResponse(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetGatewayResponseWithContext(ctx context.Context, input *apigateway.GetGatewayResponseInput, opts ...request.Option) (*apigateway.UpdateGatewayResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateGatewayResponseOutput), nil
	}
	out, err := c.APIGatewayAPI.GetGatewayResponseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetGatewayResponses(input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetGatewayResponsesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetGatewayResponses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetGatewayResponsesWithContext(ctx context.Context, input *apigateway.GetGatewayResponsesInput, opts ...request.Option) (*apigateway.GetGatewayResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetGatewayResponsesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetGatewayResponsesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetModelTemplate(input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetModelTemplateOutput), nil
	}
	out, err := c.APIGatewayAPI.GetModelTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetModelTemplateWithContext(ctx context.Context, input *apigateway.GetModelTemplateInput, opts ...request.Option) (*apigateway.GetModelTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetModelTemplateOutput), nil
	}
	out, err := c.APIGatewayAPI.GetModelTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetModels(input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetModelsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetModelsPages(input *apigateway.GetModelsInput, fn func(*apigateway.GetModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetModelsOutput{}
	fnCacher := func(out *apigateway.GetModelsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetModelsPagesWithContext(ctx context.Context, input *apigateway.GetModelsInput, fn func(*apigateway.GetModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetModelsOutput{}
	fnCacher := func(out *apigateway.GetModelsOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetModelsWithContext(ctx context.Context, input *apigateway.GetModelsInput, opts ...request.Option) (*apigateway.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetModelsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRequestValidator(input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateRequestValidatorOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRequestValidator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRequestValidatorWithContext(ctx context.Context, input *apigateway.GetRequestValidatorInput, opts ...request.Option) (*apigateway.UpdateRequestValidatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateRequestValidatorOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRequestValidatorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRequestValidators(input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetRequestValidatorsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRequestValidators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRequestValidatorsWithContext(ctx context.Context, input *apigateway.GetRequestValidatorsInput, opts ...request.Option) (*apigateway.GetRequestValidatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetRequestValidatorsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRequestValidatorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetResources(input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetResourcesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetResourcesPages(input *apigateway.GetResourcesInput, fn func(*apigateway.GetResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetResourcesOutput{}
	fnCacher := func(out *apigateway.GetResourcesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetResourcesPagesWithContext(ctx context.Context, input *apigateway.GetResourcesInput, fn func(*apigateway.GetResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetResourcesOutput{}
	fnCacher := func(out *apigateway.GetResourcesOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetResourcesWithContext(ctx context.Context, input *apigateway.GetResourcesInput, opts ...request.Option) (*apigateway.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetResourcesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRestApis(input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetRestApisOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRestApis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetRestApisPages(input *apigateway.GetRestApisInput, fn func(*apigateway.GetRestApisOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetRestApisOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetRestApisOutput{}
	fnCacher := func(out *apigateway.GetRestApisOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetRestApisPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetRestApisPagesWithContext(ctx context.Context, input *apigateway.GetRestApisInput, fn func(*apigateway.GetRestApisOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetRestApisOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetRestApisOutput{}
	fnCacher := func(out *apigateway.GetRestApisOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetRestApisPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetRestApisWithContext(ctx context.Context, input *apigateway.GetRestApisInput, opts ...request.Option) (*apigateway.GetRestApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetRestApisOutput), nil
	}
	out, err := c.APIGatewayAPI.GetRestApisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetSdk(input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetSdkOutput), nil
	}
	out, err := c.APIGatewayAPI.GetSdk(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetSdkTypes(input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetSdkTypesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetSdkTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetSdkTypesWithContext(ctx context.Context, input *apigateway.GetSdkTypesInput, opts ...request.Option) (*apigateway.GetSdkTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetSdkTypesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetSdkTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetSdkWithContext(ctx context.Context, input *apigateway.GetSdkInput, opts ...request.Option) (*apigateway.GetSdkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetSdkOutput), nil
	}
	out, err := c.APIGatewayAPI.GetSdkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetStages(input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetStagesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetStages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetStagesWithContext(ctx context.Context, input *apigateway.GetStagesInput, opts ...request.Option) (*apigateway.GetStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetStagesOutput), nil
	}
	out, err := c.APIGatewayAPI.GetStagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetTags(input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetTagsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetTagsWithContext(ctx context.Context, input *apigateway.GetTagsInput, opts ...request.Option) (*apigateway.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetTagsOutput), nil
	}
	out, err := c.APIGatewayAPI.GetTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetUsagePages(input *apigateway.GetUsageInput, fn func(*apigateway.Usage, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.Usage), false)
		return nil
	}
	cachable := true
	output := &apigateway.Usage{}
	fnCacher := func(out *apigateway.Usage, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePagesWithContext(ctx context.Context, input *apigateway.GetUsageInput, fn func(*apigateway.Usage, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.Usage), false)
		return nil
	}
	cachable := true
	output := &apigateway.Usage{}
	fnCacher := func(out *apigateway.Usage, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePlanKeys(input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetUsagePlanKeysOutput), nil
	}
	out, err := c.APIGatewayAPI.GetUsagePlanKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetUsagePlanKeysPages(input *apigateway.GetUsagePlanKeysInput, fn func(*apigateway.GetUsagePlanKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetUsagePlanKeysOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetUsagePlanKeysOutput{}
	fnCacher := func(out *apigateway.GetUsagePlanKeysOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePlanKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePlanKeysPagesWithContext(ctx context.Context, input *apigateway.GetUsagePlanKeysInput, fn func(*apigateway.GetUsagePlanKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetUsagePlanKeysOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetUsagePlanKeysOutput{}
	fnCacher := func(out *apigateway.GetUsagePlanKeysOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePlanKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePlanKeysWithContext(ctx context.Context, input *apigateway.GetUsagePlanKeysInput, opts ...request.Option) (*apigateway.GetUsagePlanKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetUsagePlanKeysOutput), nil
	}
	out, err := c.APIGatewayAPI.GetUsagePlanKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetUsagePlans(input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetUsagePlansOutput), nil
	}
	out, err := c.APIGatewayAPI.GetUsagePlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetUsagePlansPages(input *apigateway.GetUsagePlansInput, fn func(*apigateway.GetUsagePlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetUsagePlansOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetUsagePlansOutput{}
	fnCacher := func(out *apigateway.GetUsagePlansOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePlansPagesWithContext(ctx context.Context, input *apigateway.GetUsagePlansInput, fn func(*apigateway.GetUsagePlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetUsagePlansOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetUsagePlansOutput{}
	fnCacher := func(out *apigateway.GetUsagePlansOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetUsagePlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetUsagePlansWithContext(ctx context.Context, input *apigateway.GetUsagePlansInput, opts ...request.Option) (*apigateway.GetUsagePlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetUsagePlansOutput), nil
	}
	out, err := c.APIGatewayAPI.GetUsagePlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetVpcLink(input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateVpcLinkOutput), nil
	}
	out, err := c.APIGatewayAPI.GetVpcLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetVpcLinkWithContext(ctx context.Context, input *apigateway.GetVpcLinkInput, opts ...request.Option) (*apigateway.UpdateVpcLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.UpdateVpcLinkOutput), nil
	}
	out, err := c.APIGatewayAPI.GetVpcLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetVpcLinks(input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetVpcLinksOutput), nil
	}
	out, err := c.APIGatewayAPI.GetVpcLinks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *APIGateway) GetVpcLinksPages(input *apigateway.GetVpcLinksInput, fn func(*apigateway.GetVpcLinksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetVpcLinksOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetVpcLinksOutput{}
	fnCacher := func(out *apigateway.GetVpcLinksOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetVpcLinksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetVpcLinksPagesWithContext(ctx context.Context, input *apigateway.GetVpcLinksInput, fn func(*apigateway.GetVpcLinksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apigateway.GetVpcLinksOutput), false)
		return nil
	}
	cachable := true
	output := &apigateway.GetVpcLinksOutput{}
	fnCacher := func(out *apigateway.GetVpcLinksOutput, more bool) bool {
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
	if err := c.APIGatewayAPI.GetVpcLinksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *APIGateway) GetVpcLinksWithContext(ctx context.Context, input *apigateway.GetVpcLinksInput, opts ...request.Option) (*apigateway.GetVpcLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigateway.GetVpcLinksOutput), nil
	}
	out, err := c.APIGatewayAPI.GetVpcLinksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

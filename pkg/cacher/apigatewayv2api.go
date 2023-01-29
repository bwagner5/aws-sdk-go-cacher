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
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApiGatewayV2 struct {
	apigatewayv2iface.ApiGatewayV2API
	cache *cache.Cache
}

func NewApiGatewayV2(apigatewayv2api apigatewayv2iface.ApiGatewayV2API) *ApiGatewayV2 {
	return &ApiGatewayV2{
		ApiGatewayV2API: apigatewayv2api,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApiGatewayV2) GetApi(input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApiMapping(input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiMappingOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApiMapping(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApiMappingWithContext(ctx context.Context, input *apigatewayv2.GetApiMappingInput, opts ...request.Option) (*apigatewayv2.GetApiMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiMappingOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApiMappingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApiMappings(input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiMappingsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApiMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApiMappingsWithContext(ctx context.Context, input *apigatewayv2.GetApiMappingsInput, opts ...request.Option) (*apigatewayv2.GetApiMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiMappingsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApiMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApiWithContext(ctx context.Context, input *apigatewayv2.GetApiInput, opts ...request.Option) (*apigatewayv2.GetApiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApiOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApis(input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApisOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetApisWithContext(ctx context.Context, input *apigatewayv2.GetApisInput, opts ...request.Option) (*apigatewayv2.GetApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetApisOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetApisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetAuthorizer(input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetAuthorizerOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetAuthorizer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetAuthorizerWithContext(ctx context.Context, input *apigatewayv2.GetAuthorizerInput, opts ...request.Option) (*apigatewayv2.GetAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetAuthorizerOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetAuthorizerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetAuthorizers(input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetAuthorizersOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetAuthorizers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetAuthorizersWithContext(ctx context.Context, input *apigatewayv2.GetAuthorizersInput, opts ...request.Option) (*apigatewayv2.GetAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetAuthorizersOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetAuthorizersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDeployment(input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDeploymentOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDeployment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDeploymentWithContext(ctx context.Context, input *apigatewayv2.GetDeploymentInput, opts ...request.Option) (*apigatewayv2.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDeploymentOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDeploymentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDeployments(input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDeploymentsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDeploymentsWithContext(ctx context.Context, input *apigatewayv2.GetDeploymentsInput, opts ...request.Option) (*apigatewayv2.GetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDeploymentsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDomainName(input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDomainNameOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDomainName(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDomainNameWithContext(ctx context.Context, input *apigatewayv2.GetDomainNameInput, opts ...request.Option) (*apigatewayv2.GetDomainNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDomainNameOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDomainNameWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDomainNames(input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDomainNamesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetDomainNamesWithContext(ctx context.Context, input *apigatewayv2.GetDomainNamesInput, opts ...request.Option) (*apigatewayv2.GetDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetDomainNamesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegration(input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationResponse(input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationResponseOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationResponse(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationResponseWithContext(ctx context.Context, input *apigatewayv2.GetIntegrationResponseInput, opts ...request.Option) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationResponseOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationResponseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationResponses(input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationResponsesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationResponses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationResponsesWithContext(ctx context.Context, input *apigatewayv2.GetIntegrationResponsesInput, opts ...request.Option) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationResponsesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationResponsesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationWithContext(ctx context.Context, input *apigatewayv2.GetIntegrationInput, opts ...request.Option) (*apigatewayv2.GetIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrations(input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetIntegrationsWithContext(ctx context.Context, input *apigatewayv2.GetIntegrationsInput, opts ...request.Option) (*apigatewayv2.GetIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetIntegrationsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetIntegrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModel(input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModelTemplate(input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelTemplateOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModelTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModelTemplateWithContext(ctx context.Context, input *apigatewayv2.GetModelTemplateInput, opts ...request.Option) (*apigatewayv2.GetModelTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelTemplateOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModelTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModelWithContext(ctx context.Context, input *apigatewayv2.GetModelInput, opts ...request.Option) (*apigatewayv2.GetModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModels(input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetModelsWithContext(ctx context.Context, input *apigatewayv2.GetModelsInput, opts ...request.Option) (*apigatewayv2.GetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetModelsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRoute(input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRouteResponse(input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteResponseOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRouteResponse(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRouteResponseWithContext(ctx context.Context, input *apigatewayv2.GetRouteResponseInput, opts ...request.Option) (*apigatewayv2.GetRouteResponseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteResponseOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRouteResponseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRouteResponses(input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteResponsesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRouteResponses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRouteResponsesWithContext(ctx context.Context, input *apigatewayv2.GetRouteResponsesInput, opts ...request.Option) (*apigatewayv2.GetRouteResponsesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteResponsesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRouteResponsesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRouteWithContext(ctx context.Context, input *apigatewayv2.GetRouteInput, opts ...request.Option) (*apigatewayv2.GetRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRouteOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRoutes(input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRoutesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetRoutesWithContext(ctx context.Context, input *apigatewayv2.GetRoutesInput, opts ...request.Option) (*apigatewayv2.GetRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetRoutesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetStage(input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetStageOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetStage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetStageWithContext(ctx context.Context, input *apigatewayv2.GetStageInput, opts ...request.Option) (*apigatewayv2.GetStageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetStageOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetStageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetStages(input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetStagesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetStages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetStagesWithContext(ctx context.Context, input *apigatewayv2.GetStagesInput, opts ...request.Option) (*apigatewayv2.GetStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetStagesOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetStagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetTags(input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetTagsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetTagsWithContext(ctx context.Context, input *apigatewayv2.GetTagsInput, opts ...request.Option) (*apigatewayv2.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetTagsOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetVpcLink(input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetVpcLinkOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetVpcLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetVpcLinkWithContext(ctx context.Context, input *apigatewayv2.GetVpcLinkInput, opts ...request.Option) (*apigatewayv2.GetVpcLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetVpcLinkOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetVpcLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetVpcLinks(input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetVpcLinksOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetVpcLinks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayV2) GetVpcLinksWithContext(ctx context.Context, input *apigatewayv2.GetVpcLinksInput, opts ...request.Option) (*apigatewayv2.GetVpcLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewayv2.GetVpcLinksOutput), nil
	}
	out, err := c.ApiGatewayV2API.GetVpcLinksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

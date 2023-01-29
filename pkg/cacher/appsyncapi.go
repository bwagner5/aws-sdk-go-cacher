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
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/appsync/appsynciface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppSync struct {
	appsynciface.AppSyncAPI
	cache *cache.Cache
}

func NewAppSync(appsyncapi appsynciface.AppSyncAPI) *AppSync {
	return &AppSync{
		AppSyncAPI: appsyncapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppSync) GetApiAssociation(input *appsync.GetApiAssociationInput) (*appsync.GetApiAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetApiAssociationOutput), nil
	}
	out, err := c.AppSyncAPI.GetApiAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetApiAssociationWithContext(ctx context.Context, input *appsync.GetApiAssociationInput, opts ...request.Option) (*appsync.GetApiAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetApiAssociationOutput), nil
	}
	out, err := c.AppSyncAPI.GetApiAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetApiCache(input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetApiCacheOutput), nil
	}
	out, err := c.AppSyncAPI.GetApiCache(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetApiCacheWithContext(ctx context.Context, input *appsync.GetApiCacheInput, opts ...request.Option) (*appsync.GetApiCacheOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetApiCacheOutput), nil
	}
	out, err := c.AppSyncAPI.GetApiCacheWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetDataSource(input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetDataSourceOutput), nil
	}
	out, err := c.AppSyncAPI.GetDataSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetDataSourceWithContext(ctx context.Context, input *appsync.GetDataSourceInput, opts ...request.Option) (*appsync.GetDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetDataSourceOutput), nil
	}
	out, err := c.AppSyncAPI.GetDataSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetDomainName(input *appsync.GetDomainNameInput) (*appsync.GetDomainNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetDomainNameOutput), nil
	}
	out, err := c.AppSyncAPI.GetDomainName(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetDomainNameWithContext(ctx context.Context, input *appsync.GetDomainNameInput, opts ...request.Option) (*appsync.GetDomainNameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetDomainNameOutput), nil
	}
	out, err := c.AppSyncAPI.GetDomainNameWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetFunction(input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetFunctionOutput), nil
	}
	out, err := c.AppSyncAPI.GetFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetFunctionWithContext(ctx context.Context, input *appsync.GetFunctionInput, opts ...request.Option) (*appsync.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetFunctionOutput), nil
	}
	out, err := c.AppSyncAPI.GetFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetGraphqlApi(input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetGraphqlApiOutput), nil
	}
	out, err := c.AppSyncAPI.GetGraphqlApi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetGraphqlApiWithContext(ctx context.Context, input *appsync.GetGraphqlApiInput, opts ...request.Option) (*appsync.GetGraphqlApiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetGraphqlApiOutput), nil
	}
	out, err := c.AppSyncAPI.GetGraphqlApiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetIntrospectionSchema(input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetIntrospectionSchemaOutput), nil
	}
	out, err := c.AppSyncAPI.GetIntrospectionSchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetIntrospectionSchemaWithContext(ctx context.Context, input *appsync.GetIntrospectionSchemaInput, opts ...request.Option) (*appsync.GetIntrospectionSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetIntrospectionSchemaOutput), nil
	}
	out, err := c.AppSyncAPI.GetIntrospectionSchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetResolver(input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetResolverOutput), nil
	}
	out, err := c.AppSyncAPI.GetResolver(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetResolverWithContext(ctx context.Context, input *appsync.GetResolverInput, opts ...request.Option) (*appsync.GetResolverOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetResolverOutput), nil
	}
	out, err := c.AppSyncAPI.GetResolverWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetSchemaCreationStatus(input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetSchemaCreationStatusOutput), nil
	}
	out, err := c.AppSyncAPI.GetSchemaCreationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetSchemaCreationStatusWithContext(ctx context.Context, input *appsync.GetSchemaCreationStatusInput, opts ...request.Option) (*appsync.GetSchemaCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetSchemaCreationStatusOutput), nil
	}
	out, err := c.AppSyncAPI.GetSchemaCreationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetType(input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetTypeOutput), nil
	}
	out, err := c.AppSyncAPI.GetType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) GetTypeWithContext(ctx context.Context, input *appsync.GetTypeInput, opts ...request.Option) (*appsync.GetTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.GetTypeOutput), nil
	}
	out, err := c.AppSyncAPI.GetTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListApiKeys(input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListApiKeysOutput), nil
	}
	out, err := c.AppSyncAPI.ListApiKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListApiKeysWithContext(ctx context.Context, input *appsync.ListApiKeysInput, opts ...request.Option) (*appsync.ListApiKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListApiKeysOutput), nil
	}
	out, err := c.AppSyncAPI.ListApiKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListDataSources(input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListDataSourcesOutput), nil
	}
	out, err := c.AppSyncAPI.ListDataSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListDataSourcesWithContext(ctx context.Context, input *appsync.ListDataSourcesInput, opts ...request.Option) (*appsync.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListDataSourcesOutput), nil
	}
	out, err := c.AppSyncAPI.ListDataSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListDomainNames(input *appsync.ListDomainNamesInput) (*appsync.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListDomainNamesOutput), nil
	}
	out, err := c.AppSyncAPI.ListDomainNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListDomainNamesWithContext(ctx context.Context, input *appsync.ListDomainNamesInput, opts ...request.Option) (*appsync.ListDomainNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListDomainNamesOutput), nil
	}
	out, err := c.AppSyncAPI.ListDomainNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListFunctions(input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListFunctionsOutput), nil
	}
	out, err := c.AppSyncAPI.ListFunctions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListFunctionsWithContext(ctx context.Context, input *appsync.ListFunctionsInput, opts ...request.Option) (*appsync.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListFunctionsOutput), nil
	}
	out, err := c.AppSyncAPI.ListFunctionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListGraphqlApis(input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListGraphqlApisOutput), nil
	}
	out, err := c.AppSyncAPI.ListGraphqlApis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListGraphqlApisWithContext(ctx context.Context, input *appsync.ListGraphqlApisInput, opts ...request.Option) (*appsync.ListGraphqlApisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListGraphqlApisOutput), nil
	}
	out, err := c.AppSyncAPI.ListGraphqlApisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListResolvers(input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListResolversOutput), nil
	}
	out, err := c.AppSyncAPI.ListResolvers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListResolversByFunction(input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListResolversByFunctionOutput), nil
	}
	out, err := c.AppSyncAPI.ListResolversByFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListResolversByFunctionWithContext(ctx context.Context, input *appsync.ListResolversByFunctionInput, opts ...request.Option) (*appsync.ListResolversByFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListResolversByFunctionOutput), nil
	}
	out, err := c.AppSyncAPI.ListResolversByFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListResolversWithContext(ctx context.Context, input *appsync.ListResolversInput, opts ...request.Option) (*appsync.ListResolversOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListResolversOutput), nil
	}
	out, err := c.AppSyncAPI.ListResolversWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListTagsForResource(input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListTagsForResourceOutput), nil
	}
	out, err := c.AppSyncAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListTagsForResourceWithContext(ctx context.Context, input *appsync.ListTagsForResourceInput, opts ...request.Option) (*appsync.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListTagsForResourceOutput), nil
	}
	out, err := c.AppSyncAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListTypes(input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListTypesOutput), nil
	}
	out, err := c.AppSyncAPI.ListTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppSync) ListTypesWithContext(ctx context.Context, input *appsync.ListTypesInput, opts ...request.Option) (*appsync.ListTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appsync.ListTypesOutput), nil
	}
	out, err := c.AppSyncAPI.ListTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

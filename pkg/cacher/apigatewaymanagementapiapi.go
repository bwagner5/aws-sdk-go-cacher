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
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi/apigatewaymanagementapiiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApiGatewayManagementApi struct {
	apigatewaymanagementapiiface.ApiGatewayManagementApiAPI
	cache *cache.Cache
}

func NewApiGatewayManagementApi(apigatewaymanagementapiapi apigatewaymanagementapiiface.ApiGatewayManagementApiAPI) *ApiGatewayManagementApi {
	return &ApiGatewayManagementApi{
		ApiGatewayManagementApiAPI: apigatewaymanagementapiapi,
		cache:                      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApiGatewayManagementApi) GetConnection(input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewaymanagementapi.GetConnectionOutput), nil
	}
	out, err := c.ApiGatewayManagementApiAPI.GetConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApiGatewayManagementApi) GetConnectionWithContext(ctx context.Context, input *apigatewaymanagementapi.GetConnectionInput, opts ...request.Option) (*apigatewaymanagementapi.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apigatewaymanagementapi.GetConnectionOutput), nil
	}
	out, err := c.ApiGatewayManagementApiAPI.GetConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

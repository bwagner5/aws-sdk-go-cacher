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
package appintegrationsservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appintegrationsservice"
	"github.com/aws/aws-sdk-go/service/appintegrationsservice/appintegrationsserviceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppIntegrationsService struct {
	appintegrationsserviceiface.AppIntegrationsServiceAPI
	cache *cache.Cache
}

func New(appintegrationsserviceapi appintegrationsserviceiface.AppIntegrationsServiceAPI) *AppIntegrationsService {
	return &AppIntegrationsService{
		AppIntegrationsServiceAPI: appintegrationsserviceapi,
		cache:                     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppIntegrationsService) GetDataIntegration(input *appintegrationsservice.GetDataIntegrationInput) (*appintegrationsservice.GetDataIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.GetDataIntegrationOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.GetDataIntegration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) GetDataIntegrationWithContext(ctx context.Context, input *appintegrationsservice.GetDataIntegrationInput, opts ...request.Option) (*appintegrationsservice.GetDataIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.GetDataIntegrationOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.GetDataIntegrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) GetEventIntegration(input *appintegrationsservice.GetEventIntegrationInput) (*appintegrationsservice.GetEventIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.GetEventIntegrationOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.GetEventIntegration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) GetEventIntegrationWithContext(ctx context.Context, input *appintegrationsservice.GetEventIntegrationInput, opts ...request.Option) (*appintegrationsservice.GetEventIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.GetEventIntegrationOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.GetEventIntegrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListDataIntegrationAssociations(input *appintegrationsservice.ListDataIntegrationAssociationsInput) (*appintegrationsservice.ListDataIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListDataIntegrationAssociationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListDataIntegrationAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListDataIntegrationAssociationsWithContext(ctx context.Context, input *appintegrationsservice.ListDataIntegrationAssociationsInput, opts ...request.Option) (*appintegrationsservice.ListDataIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListDataIntegrationAssociationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListDataIntegrationAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListDataIntegrations(input *appintegrationsservice.ListDataIntegrationsInput) (*appintegrationsservice.ListDataIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListDataIntegrationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListDataIntegrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListDataIntegrationsWithContext(ctx context.Context, input *appintegrationsservice.ListDataIntegrationsInput, opts ...request.Option) (*appintegrationsservice.ListDataIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListDataIntegrationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListDataIntegrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListEventIntegrationAssociations(input *appintegrationsservice.ListEventIntegrationAssociationsInput) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListEventIntegrationAssociationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListEventIntegrationAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListEventIntegrationAssociationsWithContext(ctx context.Context, input *appintegrationsservice.ListEventIntegrationAssociationsInput, opts ...request.Option) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListEventIntegrationAssociationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListEventIntegrationAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListEventIntegrations(input *appintegrationsservice.ListEventIntegrationsInput) (*appintegrationsservice.ListEventIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListEventIntegrationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListEventIntegrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListEventIntegrationsWithContext(ctx context.Context, input *appintegrationsservice.ListEventIntegrationsInput, opts ...request.Option) (*appintegrationsservice.ListEventIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListEventIntegrationsOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListEventIntegrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListTagsForResource(input *appintegrationsservice.ListTagsForResourceInput) (*appintegrationsservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListTagsForResourceOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppIntegrationsService) ListTagsForResourceWithContext(ctx context.Context, input *appintegrationsservice.ListTagsForResourceInput, opts ...request.Option) (*appintegrationsservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appintegrationsservice.ListTagsForResourceOutput), nil
	}
	out, err := c.AppIntegrationsServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

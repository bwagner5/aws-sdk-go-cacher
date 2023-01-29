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
package appconfigcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appconfig/appconfigiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppConfig struct {
	appconfigiface.AppConfigAPI
	cache *cache.Cache
}

func New(appconfigapi appconfigiface.AppConfigAPI) *AppConfig {
	return &AppConfig{
		AppConfigAPI: appconfigapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppConfig) GetApplication(input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetApplicationOutput), nil
	}
	out, err := c.AppConfigAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetApplicationWithContext(ctx context.Context, input *appconfig.GetApplicationInput, opts ...request.Option) (*appconfig.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetApplicationOutput), nil
	}
	out, err := c.AppConfigAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetConfiguration(input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetConfigurationOutput), nil
	}
	out, err := c.AppConfigAPI.GetConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetConfigurationProfile(input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetConfigurationProfileOutput), nil
	}
	out, err := c.AppConfigAPI.GetConfigurationProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetConfigurationProfileWithContext(ctx context.Context, input *appconfig.GetConfigurationProfileInput, opts ...request.Option) (*appconfig.GetConfigurationProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetConfigurationProfileOutput), nil
	}
	out, err := c.AppConfigAPI.GetConfigurationProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetConfigurationWithContext(ctx context.Context, input *appconfig.GetConfigurationInput, opts ...request.Option) (*appconfig.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetConfigurationOutput), nil
	}
	out, err := c.AppConfigAPI.GetConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetDeployment(input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetDeploymentOutput), nil
	}
	out, err := c.AppConfigAPI.GetDeployment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetDeploymentStrategy(input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetDeploymentStrategyOutput), nil
	}
	out, err := c.AppConfigAPI.GetDeploymentStrategy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetDeploymentStrategyWithContext(ctx context.Context, input *appconfig.GetDeploymentStrategyInput, opts ...request.Option) (*appconfig.GetDeploymentStrategyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetDeploymentStrategyOutput), nil
	}
	out, err := c.AppConfigAPI.GetDeploymentStrategyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetDeploymentWithContext(ctx context.Context, input *appconfig.GetDeploymentInput, opts ...request.Option) (*appconfig.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetDeploymentOutput), nil
	}
	out, err := c.AppConfigAPI.GetDeploymentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetEnvironment(input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetEnvironmentOutput), nil
	}
	out, err := c.AppConfigAPI.GetEnvironment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetEnvironmentWithContext(ctx context.Context, input *appconfig.GetEnvironmentInput, opts ...request.Option) (*appconfig.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetEnvironmentOutput), nil
	}
	out, err := c.AppConfigAPI.GetEnvironmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetExtension(input *appconfig.GetExtensionInput) (*appconfig.GetExtensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetExtensionOutput), nil
	}
	out, err := c.AppConfigAPI.GetExtension(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetExtensionAssociation(input *appconfig.GetExtensionAssociationInput) (*appconfig.GetExtensionAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetExtensionAssociationOutput), nil
	}
	out, err := c.AppConfigAPI.GetExtensionAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetExtensionAssociationWithContext(ctx context.Context, input *appconfig.GetExtensionAssociationInput, opts ...request.Option) (*appconfig.GetExtensionAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetExtensionAssociationOutput), nil
	}
	out, err := c.AppConfigAPI.GetExtensionAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetExtensionWithContext(ctx context.Context, input *appconfig.GetExtensionInput, opts ...request.Option) (*appconfig.GetExtensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetExtensionOutput), nil
	}
	out, err := c.AppConfigAPI.GetExtensionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetHostedConfigurationVersion(input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetHostedConfigurationVersionOutput), nil
	}
	out, err := c.AppConfigAPI.GetHostedConfigurationVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) GetHostedConfigurationVersionWithContext(ctx context.Context, input *appconfig.GetHostedConfigurationVersionInput, opts ...request.Option) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.GetHostedConfigurationVersionOutput), nil
	}
	out, err := c.AppConfigAPI.GetHostedConfigurationVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListApplications(input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListApplicationsOutput), nil
	}
	out, err := c.AppConfigAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListApplicationsPages(input *appconfig.ListApplicationsInput, fn func(*appconfig.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListApplicationsOutput{}
	fnCacher := func(out *appconfig.ListApplicationsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListApplicationsPagesWithContext(ctx context.Context, input *appconfig.ListApplicationsInput, fn func(*appconfig.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListApplicationsOutput{}
	fnCacher := func(out *appconfig.ListApplicationsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListApplicationsWithContext(ctx context.Context, input *appconfig.ListApplicationsInput, opts ...request.Option) (*appconfig.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListApplicationsOutput), nil
	}
	out, err := c.AppConfigAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListConfigurationProfiles(input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListConfigurationProfilesOutput), nil
	}
	out, err := c.AppConfigAPI.ListConfigurationProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListConfigurationProfilesPages(input *appconfig.ListConfigurationProfilesInput, fn func(*appconfig.ListConfigurationProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListConfigurationProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListConfigurationProfilesOutput{}
	fnCacher := func(out *appconfig.ListConfigurationProfilesOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListConfigurationProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListConfigurationProfilesPagesWithContext(ctx context.Context, input *appconfig.ListConfigurationProfilesInput, fn func(*appconfig.ListConfigurationProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListConfigurationProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListConfigurationProfilesOutput{}
	fnCacher := func(out *appconfig.ListConfigurationProfilesOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListConfigurationProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListConfigurationProfilesWithContext(ctx context.Context, input *appconfig.ListConfigurationProfilesInput, opts ...request.Option) (*appconfig.ListConfigurationProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListConfigurationProfilesOutput), nil
	}
	out, err := c.AppConfigAPI.ListConfigurationProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListDeploymentStrategies(input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListDeploymentStrategiesOutput), nil
	}
	out, err := c.AppConfigAPI.ListDeploymentStrategies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListDeploymentStrategiesPages(input *appconfig.ListDeploymentStrategiesInput, fn func(*appconfig.ListDeploymentStrategiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListDeploymentStrategiesOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListDeploymentStrategiesOutput{}
	fnCacher := func(out *appconfig.ListDeploymentStrategiesOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListDeploymentStrategiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListDeploymentStrategiesPagesWithContext(ctx context.Context, input *appconfig.ListDeploymentStrategiesInput, fn func(*appconfig.ListDeploymentStrategiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListDeploymentStrategiesOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListDeploymentStrategiesOutput{}
	fnCacher := func(out *appconfig.ListDeploymentStrategiesOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListDeploymentStrategiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListDeploymentStrategiesWithContext(ctx context.Context, input *appconfig.ListDeploymentStrategiesInput, opts ...request.Option) (*appconfig.ListDeploymentStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListDeploymentStrategiesOutput), nil
	}
	out, err := c.AppConfigAPI.ListDeploymentStrategiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListDeployments(input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListDeploymentsOutput), nil
	}
	out, err := c.AppConfigAPI.ListDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListDeploymentsPages(input *appconfig.ListDeploymentsInput, fn func(*appconfig.ListDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListDeploymentsOutput{}
	fnCacher := func(out *appconfig.ListDeploymentsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListDeploymentsPagesWithContext(ctx context.Context, input *appconfig.ListDeploymentsInput, fn func(*appconfig.ListDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListDeploymentsOutput{}
	fnCacher := func(out *appconfig.ListDeploymentsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListDeploymentsWithContext(ctx context.Context, input *appconfig.ListDeploymentsInput, opts ...request.Option) (*appconfig.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListDeploymentsOutput), nil
	}
	out, err := c.AppConfigAPI.ListDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListEnvironments(input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListEnvironmentsOutput), nil
	}
	out, err := c.AppConfigAPI.ListEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListEnvironmentsPages(input *appconfig.ListEnvironmentsInput, fn func(*appconfig.ListEnvironmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListEnvironmentsOutput{}
	fnCacher := func(out *appconfig.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListEnvironmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListEnvironmentsPagesWithContext(ctx context.Context, input *appconfig.ListEnvironmentsInput, fn func(*appconfig.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListEnvironmentsOutput{}
	fnCacher := func(out *appconfig.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListEnvironmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListEnvironmentsWithContext(ctx context.Context, input *appconfig.ListEnvironmentsInput, opts ...request.Option) (*appconfig.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListEnvironmentsOutput), nil
	}
	out, err := c.AppConfigAPI.ListEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListExtensionAssociations(input *appconfig.ListExtensionAssociationsInput) (*appconfig.ListExtensionAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListExtensionAssociationsOutput), nil
	}
	out, err := c.AppConfigAPI.ListExtensionAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListExtensionAssociationsPages(input *appconfig.ListExtensionAssociationsInput, fn func(*appconfig.ListExtensionAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListExtensionAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListExtensionAssociationsOutput{}
	fnCacher := func(out *appconfig.ListExtensionAssociationsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListExtensionAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListExtensionAssociationsPagesWithContext(ctx context.Context, input *appconfig.ListExtensionAssociationsInput, fn func(*appconfig.ListExtensionAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListExtensionAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListExtensionAssociationsOutput{}
	fnCacher := func(out *appconfig.ListExtensionAssociationsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListExtensionAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListExtensionAssociationsWithContext(ctx context.Context, input *appconfig.ListExtensionAssociationsInput, opts ...request.Option) (*appconfig.ListExtensionAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListExtensionAssociationsOutput), nil
	}
	out, err := c.AppConfigAPI.ListExtensionAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListExtensions(input *appconfig.ListExtensionsInput) (*appconfig.ListExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListExtensionsOutput), nil
	}
	out, err := c.AppConfigAPI.ListExtensions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListExtensionsPages(input *appconfig.ListExtensionsInput, fn func(*appconfig.ListExtensionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListExtensionsOutput{}
	fnCacher := func(out *appconfig.ListExtensionsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListExtensionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListExtensionsPagesWithContext(ctx context.Context, input *appconfig.ListExtensionsInput, fn func(*appconfig.ListExtensionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListExtensionsOutput{}
	fnCacher := func(out *appconfig.ListExtensionsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListExtensionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListExtensionsWithContext(ctx context.Context, input *appconfig.ListExtensionsInput, opts ...request.Option) (*appconfig.ListExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListExtensionsOutput), nil
	}
	out, err := c.AppConfigAPI.ListExtensionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListHostedConfigurationVersions(input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListHostedConfigurationVersionsOutput), nil
	}
	out, err := c.AppConfigAPI.ListHostedConfigurationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListHostedConfigurationVersionsPages(input *appconfig.ListHostedConfigurationVersionsInput, fn func(*appconfig.ListHostedConfigurationVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListHostedConfigurationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListHostedConfigurationVersionsOutput{}
	fnCacher := func(out *appconfig.ListHostedConfigurationVersionsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListHostedConfigurationVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListHostedConfigurationVersionsPagesWithContext(ctx context.Context, input *appconfig.ListHostedConfigurationVersionsInput, fn func(*appconfig.ListHostedConfigurationVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appconfig.ListHostedConfigurationVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &appconfig.ListHostedConfigurationVersionsOutput{}
	fnCacher := func(out *appconfig.ListHostedConfigurationVersionsOutput, more bool) bool {
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
	if err := c.AppConfigAPI.ListHostedConfigurationVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppConfig) ListHostedConfigurationVersionsWithContext(ctx context.Context, input *appconfig.ListHostedConfigurationVersionsInput, opts ...request.Option) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListHostedConfigurationVersionsOutput), nil
	}
	out, err := c.AppConfigAPI.ListHostedConfigurationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListTagsForResource(input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListTagsForResourceOutput), nil
	}
	out, err := c.AppConfigAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfig) ListTagsForResourceWithContext(ctx context.Context, input *appconfig.ListTagsForResourceInput, opts ...request.Option) (*appconfig.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfig.ListTagsForResourceOutput), nil
	}
	out, err := c.AppConfigAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

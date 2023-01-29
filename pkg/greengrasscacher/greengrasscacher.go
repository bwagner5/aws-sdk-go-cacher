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
package greengrasscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrass/greengrassiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Greengrass struct {
	greengrassiface.GreengrassAPI
	cache *cache.Cache
}

func New(greengrassapi greengrassiface.GreengrassAPI) *Greengrass {
	return &Greengrass{
		GreengrassAPI: greengrassapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Greengrass) GetAssociatedRole(input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetAssociatedRoleOutput), nil
	}
	out, err := c.GreengrassAPI.GetAssociatedRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetAssociatedRoleWithContext(ctx context.Context, input *greengrass.GetAssociatedRoleInput, opts ...request.Option) (*greengrass.GetAssociatedRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetAssociatedRoleOutput), nil
	}
	out, err := c.GreengrassAPI.GetAssociatedRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetBulkDeploymentStatus(input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetBulkDeploymentStatusOutput), nil
	}
	out, err := c.GreengrassAPI.GetBulkDeploymentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetBulkDeploymentStatusWithContext(ctx context.Context, input *greengrass.GetBulkDeploymentStatusInput, opts ...request.Option) (*greengrass.GetBulkDeploymentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetBulkDeploymentStatusOutput), nil
	}
	out, err := c.GreengrassAPI.GetBulkDeploymentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectivityInfo(input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectivityInfoOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectivityInfo(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectivityInfoWithContext(ctx context.Context, input *greengrass.GetConnectivityInfoInput, opts ...request.Option) (*greengrass.GetConnectivityInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectivityInfoOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectivityInfoWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectorDefinition(input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectorDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectorDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectorDefinitionVersion(input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectorDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectorDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectorDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetConnectorDefinitionVersionInput, opts ...request.Option) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectorDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectorDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetConnectorDefinitionWithContext(ctx context.Context, input *greengrass.GetConnectorDefinitionInput, opts ...request.Option) (*greengrass.GetConnectorDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetConnectorDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetConnectorDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetCoreDefinition(input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetCoreDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetCoreDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetCoreDefinitionVersion(input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetCoreDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetCoreDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetCoreDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetCoreDefinitionVersionInput, opts ...request.Option) (*greengrass.GetCoreDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetCoreDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetCoreDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetCoreDefinitionWithContext(ctx context.Context, input *greengrass.GetCoreDefinitionInput, opts ...request.Option) (*greengrass.GetCoreDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetCoreDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetCoreDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeploymentStatus(input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeploymentStatusOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeploymentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeploymentStatusWithContext(ctx context.Context, input *greengrass.GetDeploymentStatusInput, opts ...request.Option) (*greengrass.GetDeploymentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeploymentStatusOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeploymentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeviceDefinition(input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeviceDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeviceDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeviceDefinitionVersion(input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeviceDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeviceDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeviceDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetDeviceDefinitionVersionInput, opts ...request.Option) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeviceDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeviceDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetDeviceDefinitionWithContext(ctx context.Context, input *greengrass.GetDeviceDefinitionInput, opts ...request.Option) (*greengrass.GetDeviceDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetDeviceDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetDeviceDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetFunctionDefinition(input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetFunctionDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetFunctionDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetFunctionDefinitionVersion(input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetFunctionDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetFunctionDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetFunctionDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetFunctionDefinitionVersionInput, opts ...request.Option) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetFunctionDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetFunctionDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetFunctionDefinitionWithContext(ctx context.Context, input *greengrass.GetFunctionDefinitionInput, opts ...request.Option) (*greengrass.GetFunctionDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetFunctionDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetFunctionDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroup(input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupCertificateAuthority(input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupCertificateAuthorityOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupCertificateAuthority(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupCertificateAuthorityWithContext(ctx context.Context, input *greengrass.GetGroupCertificateAuthorityInput, opts ...request.Option) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupCertificateAuthorityOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupCertificateAuthorityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupCertificateConfiguration(input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupCertificateConfigurationOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupCertificateConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupCertificateConfigurationWithContext(ctx context.Context, input *greengrass.GetGroupCertificateConfigurationInput, opts ...request.Option) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupCertificateConfigurationOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupCertificateConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupVersion(input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupVersionWithContext(ctx context.Context, input *greengrass.GetGroupVersionInput, opts ...request.Option) (*greengrass.GetGroupVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetGroupWithContext(ctx context.Context, input *greengrass.GetGroupInput, opts ...request.Option) (*greengrass.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetGroupOutput), nil
	}
	out, err := c.GreengrassAPI.GetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetLoggerDefinition(input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetLoggerDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetLoggerDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetLoggerDefinitionVersion(input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetLoggerDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetLoggerDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetLoggerDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetLoggerDefinitionVersionInput, opts ...request.Option) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetLoggerDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetLoggerDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetLoggerDefinitionWithContext(ctx context.Context, input *greengrass.GetLoggerDefinitionInput, opts ...request.Option) (*greengrass.GetLoggerDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetLoggerDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetLoggerDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetResourceDefinition(input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetResourceDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetResourceDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetResourceDefinitionVersion(input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetResourceDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetResourceDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetResourceDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetResourceDefinitionVersionInput, opts ...request.Option) (*greengrass.GetResourceDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetResourceDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetResourceDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetResourceDefinitionWithContext(ctx context.Context, input *greengrass.GetResourceDefinitionInput, opts ...request.Option) (*greengrass.GetResourceDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetResourceDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetResourceDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetServiceRoleForAccount(input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetServiceRoleForAccountOutput), nil
	}
	out, err := c.GreengrassAPI.GetServiceRoleForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetServiceRoleForAccountWithContext(ctx context.Context, input *greengrass.GetServiceRoleForAccountInput, opts ...request.Option) (*greengrass.GetServiceRoleForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetServiceRoleForAccountOutput), nil
	}
	out, err := c.GreengrassAPI.GetServiceRoleForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetSubscriptionDefinition(input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetSubscriptionDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetSubscriptionDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetSubscriptionDefinitionVersion(input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetSubscriptionDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetSubscriptionDefinitionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetSubscriptionDefinitionVersionWithContext(ctx context.Context, input *greengrass.GetSubscriptionDefinitionVersionInput, opts ...request.Option) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetSubscriptionDefinitionVersionOutput), nil
	}
	out, err := c.GreengrassAPI.GetSubscriptionDefinitionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetSubscriptionDefinitionWithContext(ctx context.Context, input *greengrass.GetSubscriptionDefinitionInput, opts ...request.Option) (*greengrass.GetSubscriptionDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetSubscriptionDefinitionOutput), nil
	}
	out, err := c.GreengrassAPI.GetSubscriptionDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetThingRuntimeConfiguration(input *greengrass.GetThingRuntimeConfigurationInput) (*greengrass.GetThingRuntimeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetThingRuntimeConfigurationOutput), nil
	}
	out, err := c.GreengrassAPI.GetThingRuntimeConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) GetThingRuntimeConfigurationWithContext(ctx context.Context, input *greengrass.GetThingRuntimeConfigurationInput, opts ...request.Option) (*greengrass.GetThingRuntimeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.GetThingRuntimeConfigurationOutput), nil
	}
	out, err := c.GreengrassAPI.GetThingRuntimeConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListBulkDeploymentDetailedReports(input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListBulkDeploymentDetailedReportsOutput), nil
	}
	out, err := c.GreengrassAPI.ListBulkDeploymentDetailedReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListBulkDeploymentDetailedReportsWithContext(ctx context.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput, opts ...request.Option) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListBulkDeploymentDetailedReportsOutput), nil
	}
	out, err := c.GreengrassAPI.ListBulkDeploymentDetailedReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListBulkDeployments(input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListBulkDeploymentsOutput), nil
	}
	out, err := c.GreengrassAPI.ListBulkDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListBulkDeploymentsWithContext(ctx context.Context, input *greengrass.ListBulkDeploymentsInput, opts ...request.Option) (*greengrass.ListBulkDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListBulkDeploymentsOutput), nil
	}
	out, err := c.GreengrassAPI.ListBulkDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListConnectorDefinitionVersions(input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListConnectorDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListConnectorDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListConnectorDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListConnectorDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListConnectorDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListConnectorDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListConnectorDefinitions(input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListConnectorDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListConnectorDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListConnectorDefinitionsWithContext(ctx context.Context, input *greengrass.ListConnectorDefinitionsInput, opts ...request.Option) (*greengrass.ListConnectorDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListConnectorDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListConnectorDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListCoreDefinitionVersions(input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListCoreDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListCoreDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListCoreDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListCoreDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListCoreDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListCoreDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListCoreDefinitions(input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListCoreDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListCoreDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListCoreDefinitionsWithContext(ctx context.Context, input *greengrass.ListCoreDefinitionsInput, opts ...request.Option) (*greengrass.ListCoreDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListCoreDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListCoreDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeployments(input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeploymentsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeploymentsWithContext(ctx context.Context, input *greengrass.ListDeploymentsInput, opts ...request.Option) (*greengrass.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeploymentsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeviceDefinitionVersions(input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeviceDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeviceDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeviceDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListDeviceDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeviceDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeviceDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeviceDefinitions(input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeviceDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeviceDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListDeviceDefinitionsWithContext(ctx context.Context, input *greengrass.ListDeviceDefinitionsInput, opts ...request.Option) (*greengrass.ListDeviceDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListDeviceDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListDeviceDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListFunctionDefinitionVersions(input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListFunctionDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListFunctionDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListFunctionDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListFunctionDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListFunctionDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListFunctionDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListFunctionDefinitions(input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListFunctionDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListFunctionDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListFunctionDefinitionsWithContext(ctx context.Context, input *greengrass.ListFunctionDefinitionsInput, opts ...request.Option) (*greengrass.ListFunctionDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListFunctionDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListFunctionDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroupCertificateAuthorities(input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupCertificateAuthoritiesOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroupCertificateAuthorities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroupCertificateAuthoritiesWithContext(ctx context.Context, input *greengrass.ListGroupCertificateAuthoritiesInput, opts ...request.Option) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupCertificateAuthoritiesOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroupCertificateAuthoritiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroupVersions(input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroupVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroupVersionsWithContext(ctx context.Context, input *greengrass.ListGroupVersionsInput, opts ...request.Option) (*greengrass.ListGroupVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroupVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroups(input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupsOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListGroupsWithContext(ctx context.Context, input *greengrass.ListGroupsInput, opts ...request.Option) (*greengrass.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListGroupsOutput), nil
	}
	out, err := c.GreengrassAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListLoggerDefinitionVersions(input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListLoggerDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListLoggerDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListLoggerDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListLoggerDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListLoggerDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListLoggerDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListLoggerDefinitions(input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListLoggerDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListLoggerDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListLoggerDefinitionsWithContext(ctx context.Context, input *greengrass.ListLoggerDefinitionsInput, opts ...request.Option) (*greengrass.ListLoggerDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListLoggerDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListLoggerDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListResourceDefinitionVersions(input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListResourceDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListResourceDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListResourceDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListResourceDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListResourceDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListResourceDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListResourceDefinitions(input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListResourceDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListResourceDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListResourceDefinitionsWithContext(ctx context.Context, input *greengrass.ListResourceDefinitionsInput, opts ...request.Option) (*greengrass.ListResourceDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListResourceDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListResourceDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListSubscriptionDefinitionVersions(input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListSubscriptionDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListSubscriptionDefinitionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListSubscriptionDefinitionVersionsWithContext(ctx context.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput, opts ...request.Option) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListSubscriptionDefinitionVersionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListSubscriptionDefinitionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListSubscriptionDefinitions(input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListSubscriptionDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListSubscriptionDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListSubscriptionDefinitionsWithContext(ctx context.Context, input *greengrass.ListSubscriptionDefinitionsInput, opts ...request.Option) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListSubscriptionDefinitionsOutput), nil
	}
	out, err := c.GreengrassAPI.ListSubscriptionDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListTagsForResource(input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListTagsForResourceOutput), nil
	}
	out, err := c.GreengrassAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Greengrass) ListTagsForResourceWithContext(ctx context.Context, input *greengrass.ListTagsForResourceInput, opts ...request.Option) (*greengrass.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*greengrass.ListTagsForResourceOutput), nil
	}
	out, err := c.GreengrassAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

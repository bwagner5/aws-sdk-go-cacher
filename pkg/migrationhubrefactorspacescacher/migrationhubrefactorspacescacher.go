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
package migrationhubrefactorspacescacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhubrefactorspaces"
	"github.com/aws/aws-sdk-go/service/migrationhubrefactorspaces/migrationhubrefactorspacesiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MigrationHubRefactorSpaces struct {
	migrationhubrefactorspacesiface.MigrationHubRefactorSpacesAPI
	cache *cache.Cache
}

func New(migrationhubrefactorspacesapi migrationhubrefactorspacesiface.MigrationHubRefactorSpacesAPI) *MigrationHubRefactorSpaces {
	return &MigrationHubRefactorSpaces{
		MigrationHubRefactorSpacesAPI: migrationhubrefactorspacesapi,
		cache:                         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MigrationHubRefactorSpaces) GetApplication(input *migrationhubrefactorspaces.GetApplicationInput) (*migrationhubrefactorspaces.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetApplicationOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetApplicationWithContext(ctx context.Context, input *migrationhubrefactorspaces.GetApplicationInput, opts ...request.Option) (*migrationhubrefactorspaces.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetApplicationOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetEnvironment(input *migrationhubrefactorspaces.GetEnvironmentInput) (*migrationhubrefactorspaces.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetEnvironmentOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetEnvironment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetEnvironmentWithContext(ctx context.Context, input *migrationhubrefactorspaces.GetEnvironmentInput, opts ...request.Option) (*migrationhubrefactorspaces.GetEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetEnvironmentOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetEnvironmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetResourcePolicy(input *migrationhubrefactorspaces.GetResourcePolicyInput) (*migrationhubrefactorspaces.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetResourcePolicyOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetResourcePolicyWithContext(ctx context.Context, input *migrationhubrefactorspaces.GetResourcePolicyInput, opts ...request.Option) (*migrationhubrefactorspaces.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetResourcePolicyOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetRoute(input *migrationhubrefactorspaces.GetRouteInput) (*migrationhubrefactorspaces.GetRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetRouteOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetRoute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetRouteWithContext(ctx context.Context, input *migrationhubrefactorspaces.GetRouteInput, opts ...request.Option) (*migrationhubrefactorspaces.GetRouteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetRouteOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetRouteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetService(input *migrationhubrefactorspaces.GetServiceInput) (*migrationhubrefactorspaces.GetServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetServiceOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetService(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) GetServiceWithContext(ctx context.Context, input *migrationhubrefactorspaces.GetServiceInput, opts ...request.Option) (*migrationhubrefactorspaces.GetServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.GetServiceOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.GetServiceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListApplications(input *migrationhubrefactorspaces.ListApplicationsInput) (*migrationhubrefactorspaces.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListApplicationsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListApplicationsPages(input *migrationhubrefactorspaces.ListApplicationsInput, fn func(*migrationhubrefactorspaces.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListApplicationsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListApplicationsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListApplicationsPagesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListApplicationsInput, fn func(*migrationhubrefactorspaces.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListApplicationsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListApplicationsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListApplicationsWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListApplicationsInput, opts ...request.Option) (*migrationhubrefactorspaces.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListApplicationsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentVpcs(input *migrationhubrefactorspaces.ListEnvironmentVpcsInput) (*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentVpcs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentVpcsPages(input *migrationhubrefactorspaces.ListEnvironmentVpcsInput, fn func(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListEnvironmentVpcsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListEnvironmentVpcsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentVpcsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentVpcsPagesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListEnvironmentVpcsInput, fn func(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListEnvironmentVpcsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListEnvironmentVpcsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentVpcsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentVpcsWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListEnvironmentVpcsInput, opts ...request.Option) (*migrationhubrefactorspaces.ListEnvironmentVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentVpcsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentVpcsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListEnvironments(input *migrationhubrefactorspaces.ListEnvironmentsInput) (*migrationhubrefactorspaces.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentsPages(input *migrationhubrefactorspaces.ListEnvironmentsInput, fn func(*migrationhubrefactorspaces.ListEnvironmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListEnvironmentsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentsPagesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListEnvironmentsInput, fn func(*migrationhubrefactorspaces.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListEnvironmentsOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListEnvironmentsOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListEnvironmentsWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListEnvironmentsInput, opts ...request.Option) (*migrationhubrefactorspaces.ListEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListEnvironmentsOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListRoutes(input *migrationhubrefactorspaces.ListRoutesInput) (*migrationhubrefactorspaces.ListRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListRoutesOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListRoutesPages(input *migrationhubrefactorspaces.ListRoutesInput, fn func(*migrationhubrefactorspaces.ListRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListRoutesOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListRoutesOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListRoutesPagesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListRoutesInput, fn func(*migrationhubrefactorspaces.ListRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListRoutesOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListRoutesOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListRoutesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListRoutesInput, opts ...request.Option) (*migrationhubrefactorspaces.ListRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListRoutesOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListServices(input *migrationhubrefactorspaces.ListServicesInput) (*migrationhubrefactorspaces.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListServicesOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListServicesPages(input *migrationhubrefactorspaces.ListServicesInput, fn func(*migrationhubrefactorspaces.ListServicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListServicesOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListServicesOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListServicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListServicesPagesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListServicesInput, fn func(*migrationhubrefactorspaces.ListServicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubrefactorspaces.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubrefactorspaces.ListServicesOutput{}
	fnCacher := func(out *migrationhubrefactorspaces.ListServicesOutput, more bool) bool {
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
	if err := c.MigrationHubRefactorSpacesAPI.ListServicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubRefactorSpaces) ListServicesWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListServicesInput, opts ...request.Option) (*migrationhubrefactorspaces.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListServicesOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListTagsForResource(input *migrationhubrefactorspaces.ListTagsForResourceInput) (*migrationhubrefactorspaces.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListTagsForResourceOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubRefactorSpaces) ListTagsForResourceWithContext(ctx context.Context, input *migrationhubrefactorspaces.ListTagsForResourceInput, opts ...request.Option) (*migrationhubrefactorspaces.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubrefactorspaces.ListTagsForResourceOutput), nil
	}
	out, err := c.MigrationHubRefactorSpacesAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

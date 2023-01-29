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
package redshiftserverlesscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/aws/aws-sdk-go/service/redshiftserverless/redshiftserverlessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RedshiftServerless struct {
	redshiftserverlessiface.RedshiftServerlessAPI
	cache *cache.Cache
}

func New(redshiftserverlessapi redshiftserverlessiface.RedshiftServerlessAPI) *RedshiftServerless {
	return &RedshiftServerless{
		RedshiftServerlessAPI: redshiftserverlessapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RedshiftServerless) GetCredentials(input *redshiftserverless.GetCredentialsInput) (*redshiftserverless.GetCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetCredentialsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetCredentialsWithContext(ctx context.Context, input *redshiftserverless.GetCredentialsInput, opts ...request.Option) (*redshiftserverless.GetCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetCredentialsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetEndpointAccess(input *redshiftserverless.GetEndpointAccessInput) (*redshiftserverless.GetEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetEndpointAccessOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetEndpointAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetEndpointAccessWithContext(ctx context.Context, input *redshiftserverless.GetEndpointAccessInput, opts ...request.Option) (*redshiftserverless.GetEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetEndpointAccessOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetEndpointAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetNamespace(input *redshiftserverless.GetNamespaceInput) (*redshiftserverless.GetNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetNamespaceOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetNamespace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetNamespaceWithContext(ctx context.Context, input *redshiftserverless.GetNamespaceInput, opts ...request.Option) (*redshiftserverless.GetNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetNamespaceOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetNamespaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetRecoveryPoint(input *redshiftserverless.GetRecoveryPointInput) (*redshiftserverless.GetRecoveryPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetRecoveryPointOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetRecoveryPoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetRecoveryPointWithContext(ctx context.Context, input *redshiftserverless.GetRecoveryPointInput, opts ...request.Option) (*redshiftserverless.GetRecoveryPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetRecoveryPointOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetRecoveryPointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetResourcePolicy(input *redshiftserverless.GetResourcePolicyInput) (*redshiftserverless.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetResourcePolicyOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetResourcePolicyWithContext(ctx context.Context, input *redshiftserverless.GetResourcePolicyInput, opts ...request.Option) (*redshiftserverless.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetResourcePolicyOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetSnapshot(input *redshiftserverless.GetSnapshotInput) (*redshiftserverless.GetSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetSnapshotOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetSnapshotWithContext(ctx context.Context, input *redshiftserverless.GetSnapshotInput, opts ...request.Option) (*redshiftserverless.GetSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetSnapshotOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetTableRestoreStatus(input *redshiftserverless.GetTableRestoreStatusInput) (*redshiftserverless.GetTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetTableRestoreStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetTableRestoreStatusWithContext(ctx context.Context, input *redshiftserverless.GetTableRestoreStatusInput, opts ...request.Option) (*redshiftserverless.GetTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetTableRestoreStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetUsageLimit(input *redshiftserverless.GetUsageLimitInput) (*redshiftserverless.GetUsageLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetUsageLimitOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetUsageLimit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetUsageLimitWithContext(ctx context.Context, input *redshiftserverless.GetUsageLimitInput, opts ...request.Option) (*redshiftserverless.GetUsageLimitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetUsageLimitOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetUsageLimitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetWorkgroup(input *redshiftserverless.GetWorkgroupInput) (*redshiftserverless.GetWorkgroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetWorkgroupOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetWorkgroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) GetWorkgroupWithContext(ctx context.Context, input *redshiftserverless.GetWorkgroupInput, opts ...request.Option) (*redshiftserverless.GetWorkgroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.GetWorkgroupOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.GetWorkgroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListEndpointAccess(input *redshiftserverless.ListEndpointAccessInput) (*redshiftserverless.ListEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListEndpointAccessOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListEndpointAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListEndpointAccessPages(input *redshiftserverless.ListEndpointAccessInput, fn func(*redshiftserverless.ListEndpointAccessOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListEndpointAccessOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListEndpointAccessOutput{}
	fnCacher := func(out *redshiftserverless.ListEndpointAccessOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListEndpointAccessPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListEndpointAccessPagesWithContext(ctx context.Context, input *redshiftserverless.ListEndpointAccessInput, fn func(*redshiftserverless.ListEndpointAccessOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListEndpointAccessOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListEndpointAccessOutput{}
	fnCacher := func(out *redshiftserverless.ListEndpointAccessOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListEndpointAccessPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListEndpointAccessWithContext(ctx context.Context, input *redshiftserverless.ListEndpointAccessInput, opts ...request.Option) (*redshiftserverless.ListEndpointAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListEndpointAccessOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListEndpointAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListNamespaces(input *redshiftserverless.ListNamespacesInput) (*redshiftserverless.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListNamespacesOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListNamespaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListNamespacesPages(input *redshiftserverless.ListNamespacesInput, fn func(*redshiftserverless.ListNamespacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListNamespacesOutput{}
	fnCacher := func(out *redshiftserverless.ListNamespacesOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListNamespacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListNamespacesPagesWithContext(ctx context.Context, input *redshiftserverless.ListNamespacesInput, fn func(*redshiftserverless.ListNamespacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListNamespacesOutput{}
	fnCacher := func(out *redshiftserverless.ListNamespacesOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListNamespacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListNamespacesWithContext(ctx context.Context, input *redshiftserverless.ListNamespacesInput, opts ...request.Option) (*redshiftserverless.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListNamespacesOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListNamespacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListRecoveryPoints(input *redshiftserverless.ListRecoveryPointsInput) (*redshiftserverless.ListRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListRecoveryPointsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListRecoveryPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListRecoveryPointsPages(input *redshiftserverless.ListRecoveryPointsInput, fn func(*redshiftserverless.ListRecoveryPointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListRecoveryPointsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListRecoveryPointsOutput{}
	fnCacher := func(out *redshiftserverless.ListRecoveryPointsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListRecoveryPointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListRecoveryPointsPagesWithContext(ctx context.Context, input *redshiftserverless.ListRecoveryPointsInput, fn func(*redshiftserverless.ListRecoveryPointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListRecoveryPointsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListRecoveryPointsOutput{}
	fnCacher := func(out *redshiftserverless.ListRecoveryPointsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListRecoveryPointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListRecoveryPointsWithContext(ctx context.Context, input *redshiftserverless.ListRecoveryPointsInput, opts ...request.Option) (*redshiftserverless.ListRecoveryPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListRecoveryPointsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListRecoveryPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListSnapshots(input *redshiftserverless.ListSnapshotsInput) (*redshiftserverless.ListSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListSnapshotsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListSnapshotsPages(input *redshiftserverless.ListSnapshotsInput, fn func(*redshiftserverless.ListSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListSnapshotsOutput{}
	fnCacher := func(out *redshiftserverless.ListSnapshotsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListSnapshotsPagesWithContext(ctx context.Context, input *redshiftserverless.ListSnapshotsInput, fn func(*redshiftserverless.ListSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListSnapshotsOutput{}
	fnCacher := func(out *redshiftserverless.ListSnapshotsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListSnapshotsWithContext(ctx context.Context, input *redshiftserverless.ListSnapshotsInput, opts ...request.Option) (*redshiftserverless.ListSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListSnapshotsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListTableRestoreStatus(input *redshiftserverless.ListTableRestoreStatusInput) (*redshiftserverless.ListTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListTableRestoreStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListTableRestoreStatusPages(input *redshiftserverless.ListTableRestoreStatusInput, fn func(*redshiftserverless.ListTableRestoreStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListTableRestoreStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListTableRestoreStatusOutput{}
	fnCacher := func(out *redshiftserverless.ListTableRestoreStatusOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListTableRestoreStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListTableRestoreStatusPagesWithContext(ctx context.Context, input *redshiftserverless.ListTableRestoreStatusInput, fn func(*redshiftserverless.ListTableRestoreStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListTableRestoreStatusOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListTableRestoreStatusOutput{}
	fnCacher := func(out *redshiftserverless.ListTableRestoreStatusOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListTableRestoreStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListTableRestoreStatusWithContext(ctx context.Context, input *redshiftserverless.ListTableRestoreStatusInput, opts ...request.Option) (*redshiftserverless.ListTableRestoreStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListTableRestoreStatusOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListTableRestoreStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListTagsForResource(input *redshiftserverless.ListTagsForResourceInput) (*redshiftserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListTagsForResourceWithContext(ctx context.Context, input *redshiftserverless.ListTagsForResourceInput, opts ...request.Option) (*redshiftserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListUsageLimits(input *redshiftserverless.ListUsageLimitsInput) (*redshiftserverless.ListUsageLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListUsageLimitsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListUsageLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListUsageLimitsPages(input *redshiftserverless.ListUsageLimitsInput, fn func(*redshiftserverless.ListUsageLimitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListUsageLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListUsageLimitsOutput{}
	fnCacher := func(out *redshiftserverless.ListUsageLimitsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListUsageLimitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListUsageLimitsPagesWithContext(ctx context.Context, input *redshiftserverless.ListUsageLimitsInput, fn func(*redshiftserverless.ListUsageLimitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListUsageLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListUsageLimitsOutput{}
	fnCacher := func(out *redshiftserverless.ListUsageLimitsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListUsageLimitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListUsageLimitsWithContext(ctx context.Context, input *redshiftserverless.ListUsageLimitsInput, opts ...request.Option) (*redshiftserverless.ListUsageLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListUsageLimitsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListUsageLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListWorkgroups(input *redshiftserverless.ListWorkgroupsInput) (*redshiftserverless.ListWorkgroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListWorkgroupsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListWorkgroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftServerless) ListWorkgroupsPages(input *redshiftserverless.ListWorkgroupsInput, fn func(*redshiftserverless.ListWorkgroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListWorkgroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListWorkgroupsOutput{}
	fnCacher := func(out *redshiftserverless.ListWorkgroupsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListWorkgroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListWorkgroupsPagesWithContext(ctx context.Context, input *redshiftserverless.ListWorkgroupsInput, fn func(*redshiftserverless.ListWorkgroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftserverless.ListWorkgroupsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftserverless.ListWorkgroupsOutput{}
	fnCacher := func(out *redshiftserverless.ListWorkgroupsOutput, more bool) bool {
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
	if err := c.RedshiftServerlessAPI.ListWorkgroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftServerless) ListWorkgroupsWithContext(ctx context.Context, input *redshiftserverless.ListWorkgroupsInput, opts ...request.Option) (*redshiftserverless.ListWorkgroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftserverless.ListWorkgroupsOutput), nil
	}
	out, err := c.RedshiftServerlessAPI.ListWorkgroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

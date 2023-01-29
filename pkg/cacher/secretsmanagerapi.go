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
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SecretsManager struct {
	secretsmanageriface.SecretsManagerAPI
	cache *cache.Cache
}

func NewSecretsManager(secretsmanagerapi secretsmanageriface.SecretsManagerAPI) *SecretsManager {
	return &SecretsManager{
		SecretsManagerAPI: secretsmanagerapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SecretsManager) DescribeSecret(input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.DescribeSecretOutput), nil
	}
	out, err := c.SecretsManagerAPI.DescribeSecret(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) DescribeSecretWithContext(ctx context.Context, input *secretsmanager.DescribeSecretInput, opts ...request.Option) (*secretsmanager.DescribeSecretOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.DescribeSecretOutput), nil
	}
	out, err := c.SecretsManagerAPI.DescribeSecretWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetRandomPassword(input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetRandomPasswordOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetRandomPassword(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetRandomPasswordWithContext(ctx context.Context, input *secretsmanager.GetRandomPasswordInput, opts ...request.Option) (*secretsmanager.GetRandomPasswordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetRandomPasswordOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetRandomPasswordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetResourcePolicy(input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetResourcePolicyOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetResourcePolicyWithContext(ctx context.Context, input *secretsmanager.GetResourcePolicyInput, opts ...request.Option) (*secretsmanager.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetResourcePolicyOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetSecretValueOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetSecretValue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) GetSecretValueWithContext(ctx context.Context, input *secretsmanager.GetSecretValueInput, opts ...request.Option) (*secretsmanager.GetSecretValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.GetSecretValueOutput), nil
	}
	out, err := c.SecretsManagerAPI.GetSecretValueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) ListSecretVersionIds(input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.ListSecretVersionIdsOutput), nil
	}
	out, err := c.SecretsManagerAPI.ListSecretVersionIds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) ListSecretVersionIdsPages(input *secretsmanager.ListSecretVersionIdsInput, fn func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*secretsmanager.ListSecretVersionIdsOutput), false)
		return nil
	}
	cachable := true
	output := &secretsmanager.ListSecretVersionIdsOutput{}
	fnCacher := func(out *secretsmanager.ListSecretVersionIdsOutput, more bool) bool {
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
	if err := c.SecretsManagerAPI.ListSecretVersionIdsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecretsManager) ListSecretVersionIdsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, fn func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*secretsmanager.ListSecretVersionIdsOutput), false)
		return nil
	}
	cachable := true
	output := &secretsmanager.ListSecretVersionIdsOutput{}
	fnCacher := func(out *secretsmanager.ListSecretVersionIdsOutput, more bool) bool {
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
	if err := c.SecretsManagerAPI.ListSecretVersionIdsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecretsManager) ListSecretVersionIdsWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, opts ...request.Option) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.ListSecretVersionIdsOutput), nil
	}
	out, err := c.SecretsManagerAPI.ListSecretVersionIdsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) ListSecrets(input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.ListSecretsOutput), nil
	}
	out, err := c.SecretsManagerAPI.ListSecrets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecretsManager) ListSecretsPages(input *secretsmanager.ListSecretsInput, fn func(*secretsmanager.ListSecretsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*secretsmanager.ListSecretsOutput), false)
		return nil
	}
	cachable := true
	output := &secretsmanager.ListSecretsOutput{}
	fnCacher := func(out *secretsmanager.ListSecretsOutput, more bool) bool {
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
	if err := c.SecretsManagerAPI.ListSecretsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecretsManager) ListSecretsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, fn func(*secretsmanager.ListSecretsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*secretsmanager.ListSecretsOutput), false)
		return nil
	}
	cachable := true
	output := &secretsmanager.ListSecretsOutput{}
	fnCacher := func(out *secretsmanager.ListSecretsOutput, more bool) bool {
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
	if err := c.SecretsManagerAPI.ListSecretsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecretsManager) ListSecretsWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, opts ...request.Option) (*secretsmanager.ListSecretsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*secretsmanager.ListSecretsOutput), nil
	}
	out, err := c.SecretsManagerAPI.ListSecretsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/migrationhub/migrationhubiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MigrationHub struct {
	migrationhubiface.MigrationHubAPI
	cache *cache.Cache
}

func NewMigrationHub(migrationhubapi migrationhubiface.MigrationHubAPI) *MigrationHub {
	return &MigrationHub{
		MigrationHubAPI: migrationhubapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MigrationHub) DescribeApplicationState(input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.DescribeApplicationStateOutput), nil
	}
	out, err := c.MigrationHubAPI.DescribeApplicationState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) DescribeApplicationStateWithContext(ctx context.Context, input *migrationhub.DescribeApplicationStateInput, opts ...request.Option) (*migrationhub.DescribeApplicationStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.DescribeApplicationStateOutput), nil
	}
	out, err := c.MigrationHubAPI.DescribeApplicationStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) DescribeMigrationTask(input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.DescribeMigrationTaskOutput), nil
	}
	out, err := c.MigrationHubAPI.DescribeMigrationTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) DescribeMigrationTaskWithContext(ctx context.Context, input *migrationhub.DescribeMigrationTaskInput, opts ...request.Option) (*migrationhub.DescribeMigrationTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.DescribeMigrationTaskOutput), nil
	}
	out, err := c.MigrationHubAPI.DescribeMigrationTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListApplicationStates(input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListApplicationStatesOutput), nil
	}
	out, err := c.MigrationHubAPI.ListApplicationStates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListApplicationStatesPages(input *migrationhub.ListApplicationStatesInput, fn func(*migrationhub.ListApplicationStatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListApplicationStatesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListApplicationStatesOutput{}
	fnCacher := func(out *migrationhub.ListApplicationStatesOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListApplicationStatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListApplicationStatesPagesWithContext(ctx context.Context, input *migrationhub.ListApplicationStatesInput, fn func(*migrationhub.ListApplicationStatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListApplicationStatesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListApplicationStatesOutput{}
	fnCacher := func(out *migrationhub.ListApplicationStatesOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListApplicationStatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListApplicationStatesWithContext(ctx context.Context, input *migrationhub.ListApplicationStatesInput, opts ...request.Option) (*migrationhub.ListApplicationStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListApplicationStatesOutput), nil
	}
	out, err := c.MigrationHubAPI.ListApplicationStatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListCreatedArtifacts(input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListCreatedArtifactsOutput), nil
	}
	out, err := c.MigrationHubAPI.ListCreatedArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListCreatedArtifactsPages(input *migrationhub.ListCreatedArtifactsInput, fn func(*migrationhub.ListCreatedArtifactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListCreatedArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListCreatedArtifactsOutput{}
	fnCacher := func(out *migrationhub.ListCreatedArtifactsOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListCreatedArtifactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListCreatedArtifactsPagesWithContext(ctx context.Context, input *migrationhub.ListCreatedArtifactsInput, fn func(*migrationhub.ListCreatedArtifactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListCreatedArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListCreatedArtifactsOutput{}
	fnCacher := func(out *migrationhub.ListCreatedArtifactsOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListCreatedArtifactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListCreatedArtifactsWithContext(ctx context.Context, input *migrationhub.ListCreatedArtifactsInput, opts ...request.Option) (*migrationhub.ListCreatedArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListCreatedArtifactsOutput), nil
	}
	out, err := c.MigrationHubAPI.ListCreatedArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListDiscoveredResources(input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.MigrationHubAPI.ListDiscoveredResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListDiscoveredResourcesPages(input *migrationhub.ListDiscoveredResourcesInput, fn func(*migrationhub.ListDiscoveredResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListDiscoveredResourcesOutput{}
	fnCacher := func(out *migrationhub.ListDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListDiscoveredResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListDiscoveredResourcesPagesWithContext(ctx context.Context, input *migrationhub.ListDiscoveredResourcesInput, fn func(*migrationhub.ListDiscoveredResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListDiscoveredResourcesOutput{}
	fnCacher := func(out *migrationhub.ListDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListDiscoveredResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListDiscoveredResourcesWithContext(ctx context.Context, input *migrationhub.ListDiscoveredResourcesInput, opts ...request.Option) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.MigrationHubAPI.ListDiscoveredResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListMigrationTasks(input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListMigrationTasksOutput), nil
	}
	out, err := c.MigrationHubAPI.ListMigrationTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListMigrationTasksPages(input *migrationhub.ListMigrationTasksInput, fn func(*migrationhub.ListMigrationTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListMigrationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListMigrationTasksOutput{}
	fnCacher := func(out *migrationhub.ListMigrationTasksOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListMigrationTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListMigrationTasksPagesWithContext(ctx context.Context, input *migrationhub.ListMigrationTasksInput, fn func(*migrationhub.ListMigrationTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListMigrationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListMigrationTasksOutput{}
	fnCacher := func(out *migrationhub.ListMigrationTasksOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListMigrationTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListMigrationTasksWithContext(ctx context.Context, input *migrationhub.ListMigrationTasksInput, opts ...request.Option) (*migrationhub.ListMigrationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListMigrationTasksOutput), nil
	}
	out, err := c.MigrationHubAPI.ListMigrationTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListProgressUpdateStreams(input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListProgressUpdateStreamsOutput), nil
	}
	out, err := c.MigrationHubAPI.ListProgressUpdateStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHub) ListProgressUpdateStreamsPages(input *migrationhub.ListProgressUpdateStreamsInput, fn func(*migrationhub.ListProgressUpdateStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListProgressUpdateStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListProgressUpdateStreamsOutput{}
	fnCacher := func(out *migrationhub.ListProgressUpdateStreamsOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListProgressUpdateStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListProgressUpdateStreamsPagesWithContext(ctx context.Context, input *migrationhub.ListProgressUpdateStreamsInput, fn func(*migrationhub.ListProgressUpdateStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhub.ListProgressUpdateStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhub.ListProgressUpdateStreamsOutput{}
	fnCacher := func(out *migrationhub.ListProgressUpdateStreamsOutput, more bool) bool {
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
	if err := c.MigrationHubAPI.ListProgressUpdateStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHub) ListProgressUpdateStreamsWithContext(ctx context.Context, input *migrationhub.ListProgressUpdateStreamsInput, opts ...request.Option) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhub.ListProgressUpdateStreamsOutput), nil
	}
	out, err := c.MigrationHubAPI.ListProgressUpdateStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

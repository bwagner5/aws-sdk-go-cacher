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
package kafkaconnectcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kafkaconnect"
	"github.com/aws/aws-sdk-go/service/kafkaconnect/kafkaconnectiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KafkaConnect struct {
	kafkaconnectiface.KafkaConnectAPI
	cache *cache.Cache
}

func New(kafkaconnectapi kafkaconnectiface.KafkaConnectAPI) *KafkaConnect {
	return &KafkaConnect{
		KafkaConnectAPI: kafkaconnectapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KafkaConnect) DescribeConnector(input *kafkaconnect.DescribeConnectorInput) (*kafkaconnect.DescribeConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeConnectorOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeConnector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) DescribeConnectorWithContext(ctx context.Context, input *kafkaconnect.DescribeConnectorInput, opts ...request.Option) (*kafkaconnect.DescribeConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeConnectorOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeConnectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) DescribeCustomPlugin(input *kafkaconnect.DescribeCustomPluginInput) (*kafkaconnect.DescribeCustomPluginOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeCustomPluginOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeCustomPlugin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) DescribeCustomPluginWithContext(ctx context.Context, input *kafkaconnect.DescribeCustomPluginInput, opts ...request.Option) (*kafkaconnect.DescribeCustomPluginOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeCustomPluginOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeCustomPluginWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) DescribeWorkerConfiguration(input *kafkaconnect.DescribeWorkerConfigurationInput) (*kafkaconnect.DescribeWorkerConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeWorkerConfigurationOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeWorkerConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) DescribeWorkerConfigurationWithContext(ctx context.Context, input *kafkaconnect.DescribeWorkerConfigurationInput, opts ...request.Option) (*kafkaconnect.DescribeWorkerConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.DescribeWorkerConfigurationOutput), nil
	}
	out, err := c.KafkaConnectAPI.DescribeWorkerConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListConnectors(input *kafkaconnect.ListConnectorsInput) (*kafkaconnect.ListConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListConnectorsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListConnectorsPages(input *kafkaconnect.ListConnectorsInput, fn func(*kafkaconnect.ListConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListConnectorsOutput{}
	fnCacher := func(out *kafkaconnect.ListConnectorsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListConnectorsPagesWithContext(ctx context.Context, input *kafkaconnect.ListConnectorsInput, fn func(*kafkaconnect.ListConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListConnectorsOutput{}
	fnCacher := func(out *kafkaconnect.ListConnectorsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListConnectorsWithContext(ctx context.Context, input *kafkaconnect.ListConnectorsInput, opts ...request.Option) (*kafkaconnect.ListConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListConnectorsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListCustomPlugins(input *kafkaconnect.ListCustomPluginsInput) (*kafkaconnect.ListCustomPluginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListCustomPluginsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListCustomPlugins(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListCustomPluginsPages(input *kafkaconnect.ListCustomPluginsInput, fn func(*kafkaconnect.ListCustomPluginsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListCustomPluginsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListCustomPluginsOutput{}
	fnCacher := func(out *kafkaconnect.ListCustomPluginsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListCustomPluginsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListCustomPluginsPagesWithContext(ctx context.Context, input *kafkaconnect.ListCustomPluginsInput, fn func(*kafkaconnect.ListCustomPluginsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListCustomPluginsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListCustomPluginsOutput{}
	fnCacher := func(out *kafkaconnect.ListCustomPluginsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListCustomPluginsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListCustomPluginsWithContext(ctx context.Context, input *kafkaconnect.ListCustomPluginsInput, opts ...request.Option) (*kafkaconnect.ListCustomPluginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListCustomPluginsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListCustomPluginsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListWorkerConfigurations(input *kafkaconnect.ListWorkerConfigurationsInput) (*kafkaconnect.ListWorkerConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListWorkerConfigurationsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListWorkerConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KafkaConnect) ListWorkerConfigurationsPages(input *kafkaconnect.ListWorkerConfigurationsInput, fn func(*kafkaconnect.ListWorkerConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListWorkerConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListWorkerConfigurationsOutput{}
	fnCacher := func(out *kafkaconnect.ListWorkerConfigurationsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListWorkerConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListWorkerConfigurationsPagesWithContext(ctx context.Context, input *kafkaconnect.ListWorkerConfigurationsInput, fn func(*kafkaconnect.ListWorkerConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kafkaconnect.ListWorkerConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &kafkaconnect.ListWorkerConfigurationsOutput{}
	fnCacher := func(out *kafkaconnect.ListWorkerConfigurationsOutput, more bool) bool {
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
	if err := c.KafkaConnectAPI.ListWorkerConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KafkaConnect) ListWorkerConfigurationsWithContext(ctx context.Context, input *kafkaconnect.ListWorkerConfigurationsInput, opts ...request.Option) (*kafkaconnect.ListWorkerConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kafkaconnect.ListWorkerConfigurationsOutput), nil
	}
	out, err := c.KafkaConnectAPI.ListWorkerConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

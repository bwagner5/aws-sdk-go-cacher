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
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Polly struct {
	pollyiface.PollyAPI
	cache *cache.Cache
}

func NewPolly(pollyapi pollyiface.PollyAPI) *Polly {
	return &Polly{
		PollyAPI: pollyapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Polly) DescribeVoices(input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.DescribeVoicesOutput), nil
	}
	out, err := c.PollyAPI.DescribeVoices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) DescribeVoicesWithContext(ctx context.Context, input *polly.DescribeVoicesInput, opts ...request.Option) (*polly.DescribeVoicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.DescribeVoicesOutput), nil
	}
	out, err := c.PollyAPI.DescribeVoicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) GetLexicon(input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.GetLexiconOutput), nil
	}
	out, err := c.PollyAPI.GetLexicon(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) GetLexiconWithContext(ctx context.Context, input *polly.GetLexiconInput, opts ...request.Option) (*polly.GetLexiconOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.GetLexiconOutput), nil
	}
	out, err := c.PollyAPI.GetLexiconWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) GetSpeechSynthesisTask(input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.GetSpeechSynthesisTaskOutput), nil
	}
	out, err := c.PollyAPI.GetSpeechSynthesisTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) GetSpeechSynthesisTaskWithContext(ctx context.Context, input *polly.GetSpeechSynthesisTaskInput, opts ...request.Option) (*polly.GetSpeechSynthesisTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.GetSpeechSynthesisTaskOutput), nil
	}
	out, err := c.PollyAPI.GetSpeechSynthesisTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) ListLexicons(input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.ListLexiconsOutput), nil
	}
	out, err := c.PollyAPI.ListLexicons(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) ListLexiconsWithContext(ctx context.Context, input *polly.ListLexiconsInput, opts ...request.Option) (*polly.ListLexiconsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.ListLexiconsOutput), nil
	}
	out, err := c.PollyAPI.ListLexiconsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) ListSpeechSynthesisTasks(input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.ListSpeechSynthesisTasksOutput), nil
	}
	out, err := c.PollyAPI.ListSpeechSynthesisTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Polly) ListSpeechSynthesisTasksPages(input *polly.ListSpeechSynthesisTasksInput, fn func(*polly.ListSpeechSynthesisTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*polly.ListSpeechSynthesisTasksOutput), false)
		return nil
	}
	cachable := true
	output := &polly.ListSpeechSynthesisTasksOutput{}
	fnCacher := func(out *polly.ListSpeechSynthesisTasksOutput, more bool) bool {
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
	if err := c.PollyAPI.ListSpeechSynthesisTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Polly) ListSpeechSynthesisTasksPagesWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, fn func(*polly.ListSpeechSynthesisTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*polly.ListSpeechSynthesisTasksOutput), false)
		return nil
	}
	cachable := true
	output := &polly.ListSpeechSynthesisTasksOutput{}
	fnCacher := func(out *polly.ListSpeechSynthesisTasksOutput, more bool) bool {
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
	if err := c.PollyAPI.ListSpeechSynthesisTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Polly) ListSpeechSynthesisTasksWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, opts ...request.Option) (*polly.ListSpeechSynthesisTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*polly.ListSpeechSynthesisTasksOutput), nil
	}
	out, err := c.PollyAPI.ListSpeechSynthesisTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

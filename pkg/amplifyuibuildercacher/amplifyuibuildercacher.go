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
package amplifyuibuildercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/amplifyuibuilder"
	"github.com/aws/aws-sdk-go/service/amplifyuibuilder/amplifyuibuilderiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AmplifyUIBuilder struct {
	amplifyuibuilderiface.AmplifyUIBuilderAPI
	cache *cache.Cache
}

func New(amplifyuibuilderapi amplifyuibuilderiface.AmplifyUIBuilderAPI) *AmplifyUIBuilder {
	return &AmplifyUIBuilder{
		AmplifyUIBuilderAPI: amplifyuibuilderapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AmplifyUIBuilder) GetComponent(input *amplifyuibuilder.GetComponentInput) (*amplifyuibuilder.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetComponentOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetComponentWithContext(ctx context.Context, input *amplifyuibuilder.GetComponentInput, opts ...request.Option) (*amplifyuibuilder.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetComponentOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetForm(input *amplifyuibuilder.GetFormInput) (*amplifyuibuilder.GetFormOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetFormOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetForm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetFormWithContext(ctx context.Context, input *amplifyuibuilder.GetFormInput, opts ...request.Option) (*amplifyuibuilder.GetFormOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetFormOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetFormWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetMetadata(input *amplifyuibuilder.GetMetadataInput) (*amplifyuibuilder.GetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetMetadataOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetMetadataWithContext(ctx context.Context, input *amplifyuibuilder.GetMetadataInput, opts ...request.Option) (*amplifyuibuilder.GetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetMetadataOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetTheme(input *amplifyuibuilder.GetThemeInput) (*amplifyuibuilder.GetThemeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetThemeOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetTheme(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) GetThemeWithContext(ctx context.Context, input *amplifyuibuilder.GetThemeInput, opts ...request.Option) (*amplifyuibuilder.GetThemeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.GetThemeOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.GetThemeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListComponents(input *amplifyuibuilder.ListComponentsInput) (*amplifyuibuilder.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListComponentsOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListComponentsPages(input *amplifyuibuilder.ListComponentsInput, fn func(*amplifyuibuilder.ListComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListComponentsOutput{}
	fnCacher := func(out *amplifyuibuilder.ListComponentsOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListComponentsPagesWithContext(ctx context.Context, input *amplifyuibuilder.ListComponentsInput, fn func(*amplifyuibuilder.ListComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListComponentsOutput{}
	fnCacher := func(out *amplifyuibuilder.ListComponentsOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListComponentsWithContext(ctx context.Context, input *amplifyuibuilder.ListComponentsInput, opts ...request.Option) (*amplifyuibuilder.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListComponentsOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListForms(input *amplifyuibuilder.ListFormsInput) (*amplifyuibuilder.ListFormsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListFormsOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListForms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListFormsPages(input *amplifyuibuilder.ListFormsInput, fn func(*amplifyuibuilder.ListFormsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListFormsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListFormsOutput{}
	fnCacher := func(out *amplifyuibuilder.ListFormsOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListFormsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListFormsPagesWithContext(ctx context.Context, input *amplifyuibuilder.ListFormsInput, fn func(*amplifyuibuilder.ListFormsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListFormsOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListFormsOutput{}
	fnCacher := func(out *amplifyuibuilder.ListFormsOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListFormsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListFormsWithContext(ctx context.Context, input *amplifyuibuilder.ListFormsInput, opts ...request.Option) (*amplifyuibuilder.ListFormsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListFormsOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListFormsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListThemes(input *amplifyuibuilder.ListThemesInput) (*amplifyuibuilder.ListThemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListThemesOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListThemes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AmplifyUIBuilder) ListThemesPages(input *amplifyuibuilder.ListThemesInput, fn func(*amplifyuibuilder.ListThemesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListThemesOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListThemesOutput{}
	fnCacher := func(out *amplifyuibuilder.ListThemesOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListThemesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListThemesPagesWithContext(ctx context.Context, input *amplifyuibuilder.ListThemesInput, fn func(*amplifyuibuilder.ListThemesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*amplifyuibuilder.ListThemesOutput), false)
		return nil
	}
	cachable := true
	output := &amplifyuibuilder.ListThemesOutput{}
	fnCacher := func(out *amplifyuibuilder.ListThemesOutput, more bool) bool {
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
	if err := c.AmplifyUIBuilderAPI.ListThemesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AmplifyUIBuilder) ListThemesWithContext(ctx context.Context, input *amplifyuibuilder.ListThemesInput, opts ...request.Option) (*amplifyuibuilder.ListThemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplifyuibuilder.ListThemesOutput), nil
	}
	out, err := c.AmplifyUIBuilderAPI.ListThemesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

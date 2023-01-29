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
package imagebuildercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/imagebuilder/imagebuilderiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Imagebuilder struct {
	imagebuilderiface.ImagebuilderAPI
	cache *cache.Cache
}

func New(imagebuilderapi imagebuilderiface.ImagebuilderAPI) *Imagebuilder {
	return &Imagebuilder{
		ImagebuilderAPI: imagebuilderapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Imagebuilder) GetComponent(input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetComponentOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetComponentPolicy(input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetComponentPolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetComponentPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetComponentPolicyWithContext(ctx context.Context, input *imagebuilder.GetComponentPolicyInput, opts ...request.Option) (*imagebuilder.GetComponentPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetComponentPolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetComponentPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetComponentWithContext(ctx context.Context, input *imagebuilder.GetComponentInput, opts ...request.Option) (*imagebuilder.GetComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetComponentOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetContainerRecipe(input *imagebuilder.GetContainerRecipeInput) (*imagebuilder.GetContainerRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetContainerRecipeOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetContainerRecipe(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetContainerRecipePolicy(input *imagebuilder.GetContainerRecipePolicyInput) (*imagebuilder.GetContainerRecipePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetContainerRecipePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetContainerRecipePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetContainerRecipePolicyWithContext(ctx context.Context, input *imagebuilder.GetContainerRecipePolicyInput, opts ...request.Option) (*imagebuilder.GetContainerRecipePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetContainerRecipePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetContainerRecipePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetContainerRecipeWithContext(ctx context.Context, input *imagebuilder.GetContainerRecipeInput, opts ...request.Option) (*imagebuilder.GetContainerRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetContainerRecipeOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetContainerRecipeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetDistributionConfiguration(input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetDistributionConfigurationOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetDistributionConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetDistributionConfigurationWithContext(ctx context.Context, input *imagebuilder.GetDistributionConfigurationInput, opts ...request.Option) (*imagebuilder.GetDistributionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetDistributionConfigurationOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetDistributionConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImage(input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImagePipeline(input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImagePipelineOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImagePipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImagePipelineWithContext(ctx context.Context, input *imagebuilder.GetImagePipelineInput, opts ...request.Option) (*imagebuilder.GetImagePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImagePipelineOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImagePipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImagePolicy(input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImagePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImagePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImagePolicyWithContext(ctx context.Context, input *imagebuilder.GetImagePolicyInput, opts ...request.Option) (*imagebuilder.GetImagePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImagePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImagePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImageRecipe(input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageRecipeOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImageRecipe(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImageRecipePolicy(input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageRecipePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImageRecipePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImageRecipePolicyWithContext(ctx context.Context, input *imagebuilder.GetImageRecipePolicyInput, opts ...request.Option) (*imagebuilder.GetImageRecipePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageRecipePolicyOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImageRecipePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImageRecipeWithContext(ctx context.Context, input *imagebuilder.GetImageRecipeInput, opts ...request.Option) (*imagebuilder.GetImageRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageRecipeOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImageRecipeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetImageWithContext(ctx context.Context, input *imagebuilder.GetImageInput, opts ...request.Option) (*imagebuilder.GetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetImageOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetInfrastructureConfiguration(input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetInfrastructureConfigurationOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetInfrastructureConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) GetInfrastructureConfigurationWithContext(ctx context.Context, input *imagebuilder.GetInfrastructureConfigurationInput, opts ...request.Option) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.GetInfrastructureConfigurationOutput), nil
	}
	out, err := c.ImagebuilderAPI.GetInfrastructureConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListComponentBuildVersions(input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListComponentBuildVersionsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListComponentBuildVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListComponentBuildVersionsPages(input *imagebuilder.ListComponentBuildVersionsInput, fn func(*imagebuilder.ListComponentBuildVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListComponentBuildVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListComponentBuildVersionsOutput{}
	fnCacher := func(out *imagebuilder.ListComponentBuildVersionsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListComponentBuildVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListComponentBuildVersionsPagesWithContext(ctx context.Context, input *imagebuilder.ListComponentBuildVersionsInput, fn func(*imagebuilder.ListComponentBuildVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListComponentBuildVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListComponentBuildVersionsOutput{}
	fnCacher := func(out *imagebuilder.ListComponentBuildVersionsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListComponentBuildVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListComponentBuildVersionsWithContext(ctx context.Context, input *imagebuilder.ListComponentBuildVersionsInput, opts ...request.Option) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListComponentBuildVersionsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListComponentBuildVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListComponents(input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListComponentsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListComponentsPages(input *imagebuilder.ListComponentsInput, fn func(*imagebuilder.ListComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListComponentsOutput{}
	fnCacher := func(out *imagebuilder.ListComponentsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListComponentsPagesWithContext(ctx context.Context, input *imagebuilder.ListComponentsInput, fn func(*imagebuilder.ListComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListComponentsOutput{}
	fnCacher := func(out *imagebuilder.ListComponentsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListComponentsWithContext(ctx context.Context, input *imagebuilder.ListComponentsInput, opts ...request.Option) (*imagebuilder.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListComponentsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListContainerRecipes(input *imagebuilder.ListContainerRecipesInput) (*imagebuilder.ListContainerRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListContainerRecipesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListContainerRecipes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListContainerRecipesPages(input *imagebuilder.ListContainerRecipesInput, fn func(*imagebuilder.ListContainerRecipesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListContainerRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListContainerRecipesOutput{}
	fnCacher := func(out *imagebuilder.ListContainerRecipesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListContainerRecipesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListContainerRecipesPagesWithContext(ctx context.Context, input *imagebuilder.ListContainerRecipesInput, fn func(*imagebuilder.ListContainerRecipesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListContainerRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListContainerRecipesOutput{}
	fnCacher := func(out *imagebuilder.ListContainerRecipesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListContainerRecipesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListContainerRecipesWithContext(ctx context.Context, input *imagebuilder.ListContainerRecipesInput, opts ...request.Option) (*imagebuilder.ListContainerRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListContainerRecipesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListContainerRecipesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListDistributionConfigurations(input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListDistributionConfigurationsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListDistributionConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListDistributionConfigurationsPages(input *imagebuilder.ListDistributionConfigurationsInput, fn func(*imagebuilder.ListDistributionConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListDistributionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListDistributionConfigurationsOutput{}
	fnCacher := func(out *imagebuilder.ListDistributionConfigurationsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListDistributionConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListDistributionConfigurationsPagesWithContext(ctx context.Context, input *imagebuilder.ListDistributionConfigurationsInput, fn func(*imagebuilder.ListDistributionConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListDistributionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListDistributionConfigurationsOutput{}
	fnCacher := func(out *imagebuilder.ListDistributionConfigurationsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListDistributionConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListDistributionConfigurationsWithContext(ctx context.Context, input *imagebuilder.ListDistributionConfigurationsInput, opts ...request.Option) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListDistributionConfigurationsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListDistributionConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImageBuildVersions(input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImageBuildVersionsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImageBuildVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImageBuildVersionsPages(input *imagebuilder.ListImageBuildVersionsInput, fn func(*imagebuilder.ListImageBuildVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImageBuildVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImageBuildVersionsOutput{}
	fnCacher := func(out *imagebuilder.ListImageBuildVersionsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImageBuildVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImageBuildVersionsPagesWithContext(ctx context.Context, input *imagebuilder.ListImageBuildVersionsInput, fn func(*imagebuilder.ListImageBuildVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImageBuildVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImageBuildVersionsOutput{}
	fnCacher := func(out *imagebuilder.ListImageBuildVersionsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImageBuildVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImageBuildVersionsWithContext(ctx context.Context, input *imagebuilder.ListImageBuildVersionsInput, opts ...request.Option) (*imagebuilder.ListImageBuildVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImageBuildVersionsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImageBuildVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePackages(input *imagebuilder.ListImagePackagesInput) (*imagebuilder.ListImagePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePackagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePackagesPages(input *imagebuilder.ListImagePackagesInput, fn func(*imagebuilder.ListImagePackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePackagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePackagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePackagesPagesWithContext(ctx context.Context, input *imagebuilder.ListImagePackagesInput, fn func(*imagebuilder.ListImagePackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePackagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePackagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePackagesWithContext(ctx context.Context, input *imagebuilder.ListImagePackagesInput, opts ...request.Option) (*imagebuilder.ListImagePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePackagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePipelineImages(input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePipelineImagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePipelineImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePipelineImagesPages(input *imagebuilder.ListImagePipelineImagesInput, fn func(*imagebuilder.ListImagePipelineImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePipelineImagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePipelineImagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePipelineImagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePipelineImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePipelineImagesPagesWithContext(ctx context.Context, input *imagebuilder.ListImagePipelineImagesInput, fn func(*imagebuilder.ListImagePipelineImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePipelineImagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePipelineImagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePipelineImagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePipelineImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePipelineImagesWithContext(ctx context.Context, input *imagebuilder.ListImagePipelineImagesInput, opts ...request.Option) (*imagebuilder.ListImagePipelineImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePipelineImagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePipelineImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePipelines(input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePipelinesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagePipelinesPages(input *imagebuilder.ListImagePipelinesInput, fn func(*imagebuilder.ListImagePipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePipelinesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePipelinesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePipelinesPagesWithContext(ctx context.Context, input *imagebuilder.ListImagePipelinesInput, fn func(*imagebuilder.ListImagePipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagePipelinesOutput{}
	fnCacher := func(out *imagebuilder.ListImagePipelinesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagePipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagePipelinesWithContext(ctx context.Context, input *imagebuilder.ListImagePipelinesInput, opts ...request.Option) (*imagebuilder.ListImagePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagePipelinesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagePipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImageRecipes(input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImageRecipesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImageRecipes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImageRecipesPages(input *imagebuilder.ListImageRecipesInput, fn func(*imagebuilder.ListImageRecipesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImageRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImageRecipesOutput{}
	fnCacher := func(out *imagebuilder.ListImageRecipesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImageRecipesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImageRecipesPagesWithContext(ctx context.Context, input *imagebuilder.ListImageRecipesInput, fn func(*imagebuilder.ListImageRecipesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImageRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImageRecipesOutput{}
	fnCacher := func(out *imagebuilder.ListImageRecipesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImageRecipesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImageRecipesWithContext(ctx context.Context, input *imagebuilder.ListImageRecipesInput, opts ...request.Option) (*imagebuilder.ListImageRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImageRecipesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImageRecipesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImages(input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListImagesPages(input *imagebuilder.ListImagesInput, fn func(*imagebuilder.ListImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagesPagesWithContext(ctx context.Context, input *imagebuilder.ListImagesInput, fn func(*imagebuilder.ListImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListImagesOutput{}
	fnCacher := func(out *imagebuilder.ListImagesOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListImagesWithContext(ctx context.Context, input *imagebuilder.ListImagesInput, opts ...request.Option) (*imagebuilder.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListImagesOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListInfrastructureConfigurations(input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListInfrastructureConfigurationsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListInfrastructureConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListInfrastructureConfigurationsPages(input *imagebuilder.ListInfrastructureConfigurationsInput, fn func(*imagebuilder.ListInfrastructureConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListInfrastructureConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListInfrastructureConfigurationsOutput{}
	fnCacher := func(out *imagebuilder.ListInfrastructureConfigurationsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListInfrastructureConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListInfrastructureConfigurationsPagesWithContext(ctx context.Context, input *imagebuilder.ListInfrastructureConfigurationsInput, fn func(*imagebuilder.ListInfrastructureConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*imagebuilder.ListInfrastructureConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &imagebuilder.ListInfrastructureConfigurationsOutput{}
	fnCacher := func(out *imagebuilder.ListInfrastructureConfigurationsOutput, more bool) bool {
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
	if err := c.ImagebuilderAPI.ListInfrastructureConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Imagebuilder) ListInfrastructureConfigurationsWithContext(ctx context.Context, input *imagebuilder.ListInfrastructureConfigurationsInput, opts ...request.Option) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListInfrastructureConfigurationsOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListInfrastructureConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListTagsForResource(input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListTagsForResourceOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Imagebuilder) ListTagsForResourceWithContext(ctx context.Context, input *imagebuilder.ListTagsForResourceInput, opts ...request.Option) (*imagebuilder.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*imagebuilder.ListTagsForResourceOutput), nil
	}
	out, err := c.ImagebuilderAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

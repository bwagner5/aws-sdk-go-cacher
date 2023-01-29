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
package clouddirectorycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudDirectory struct {
	clouddirectoryiface.CloudDirectoryAPI
	cache *cache.Cache
}

func New(clouddirectoryapi clouddirectoryiface.CloudDirectoryAPI) *CloudDirectory {
	return &CloudDirectory{
		CloudDirectoryAPI: clouddirectoryapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudDirectory) BatchRead(input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.BatchReadOutput), nil
	}
	out, err := c.CloudDirectoryAPI.BatchRead(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) BatchReadWithContext(ctx context.Context, input *clouddirectory.BatchReadInput, opts ...request.Option) (*clouddirectory.BatchReadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.BatchReadOutput), nil
	}
	out, err := c.CloudDirectoryAPI.BatchReadWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) BatchWrite(input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.BatchWriteOutput), nil
	}
	out, err := c.CloudDirectoryAPI.BatchWrite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) BatchWriteWithContext(ctx context.Context, input *clouddirectory.BatchWriteInput, opts ...request.Option) (*clouddirectory.BatchWriteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.BatchWriteOutput), nil
	}
	out, err := c.CloudDirectoryAPI.BatchWriteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetAppliedSchemaVersion(input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetAppliedSchemaVersionOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetAppliedSchemaVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetAppliedSchemaVersionWithContext(ctx context.Context, input *clouddirectory.GetAppliedSchemaVersionInput, opts ...request.Option) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetAppliedSchemaVersionOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetAppliedSchemaVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetDirectory(input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetDirectoryOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetDirectory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetDirectoryWithContext(ctx context.Context, input *clouddirectory.GetDirectoryInput, opts ...request.Option) (*clouddirectory.GetDirectoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetDirectoryOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetDirectoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetFacet(input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetFacetOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetFacet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetFacetWithContext(ctx context.Context, input *clouddirectory.GetFacetInput, opts ...request.Option) (*clouddirectory.GetFacetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetFacetOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetFacetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetLinkAttributes(input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetLinkAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetLinkAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetLinkAttributesWithContext(ctx context.Context, input *clouddirectory.GetLinkAttributesInput, opts ...request.Option) (*clouddirectory.GetLinkAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetLinkAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetLinkAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetObjectAttributes(input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetObjectAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetObjectAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetObjectAttributesWithContext(ctx context.Context, input *clouddirectory.GetObjectAttributesInput, opts ...request.Option) (*clouddirectory.GetObjectAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetObjectAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetObjectAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetObjectInformation(input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetObjectInformationOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetObjectInformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetObjectInformationWithContext(ctx context.Context, input *clouddirectory.GetObjectInformationInput, opts ...request.Option) (*clouddirectory.GetObjectInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetObjectInformationOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetObjectInformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetSchemaAsJson(input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetSchemaAsJsonOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetSchemaAsJson(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetSchemaAsJsonWithContext(ctx context.Context, input *clouddirectory.GetSchemaAsJsonInput, opts ...request.Option) (*clouddirectory.GetSchemaAsJsonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetSchemaAsJsonOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetSchemaAsJsonWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetTypedLinkFacetInformation(input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetTypedLinkFacetInformationOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetTypedLinkFacetInformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) GetTypedLinkFacetInformationWithContext(ctx context.Context, input *clouddirectory.GetTypedLinkFacetInformationInput, opts ...request.Option) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.GetTypedLinkFacetInformationOutput), nil
	}
	out, err := c.CloudDirectoryAPI.GetTypedLinkFacetInformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListAppliedSchemaArns(input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListAppliedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListAppliedSchemaArns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListAppliedSchemaArnsPages(input *clouddirectory.ListAppliedSchemaArnsInput, fn func(*clouddirectory.ListAppliedSchemaArnsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListAppliedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListAppliedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListAppliedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListAppliedSchemaArnsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListAppliedSchemaArnsPagesWithContext(ctx context.Context, input *clouddirectory.ListAppliedSchemaArnsInput, fn func(*clouddirectory.ListAppliedSchemaArnsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListAppliedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListAppliedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListAppliedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListAppliedSchemaArnsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListAppliedSchemaArnsWithContext(ctx context.Context, input *clouddirectory.ListAppliedSchemaArnsInput, opts ...request.Option) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListAppliedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListAppliedSchemaArnsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListAttachedIndices(input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListAttachedIndicesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListAttachedIndices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListAttachedIndicesPages(input *clouddirectory.ListAttachedIndicesInput, fn func(*clouddirectory.ListAttachedIndicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListAttachedIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListAttachedIndicesOutput{}
	fnCacher := func(out *clouddirectory.ListAttachedIndicesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListAttachedIndicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListAttachedIndicesPagesWithContext(ctx context.Context, input *clouddirectory.ListAttachedIndicesInput, fn func(*clouddirectory.ListAttachedIndicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListAttachedIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListAttachedIndicesOutput{}
	fnCacher := func(out *clouddirectory.ListAttachedIndicesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListAttachedIndicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListAttachedIndicesWithContext(ctx context.Context, input *clouddirectory.ListAttachedIndicesInput, opts ...request.Option) (*clouddirectory.ListAttachedIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListAttachedIndicesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListAttachedIndicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListDevelopmentSchemaArns(input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListDevelopmentSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListDevelopmentSchemaArns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListDevelopmentSchemaArnsPages(input *clouddirectory.ListDevelopmentSchemaArnsInput, fn func(*clouddirectory.ListDevelopmentSchemaArnsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListDevelopmentSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListDevelopmentSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListDevelopmentSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListDevelopmentSchemaArnsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListDevelopmentSchemaArnsPagesWithContext(ctx context.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput, fn func(*clouddirectory.ListDevelopmentSchemaArnsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListDevelopmentSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListDevelopmentSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListDevelopmentSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListDevelopmentSchemaArnsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListDevelopmentSchemaArnsWithContext(ctx context.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput, opts ...request.Option) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListDevelopmentSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListDevelopmentSchemaArnsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListDirectories(input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListDirectoriesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListDirectories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListDirectoriesPages(input *clouddirectory.ListDirectoriesInput, fn func(*clouddirectory.ListDirectoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListDirectoriesOutput{}
	fnCacher := func(out *clouddirectory.ListDirectoriesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListDirectoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListDirectoriesPagesWithContext(ctx context.Context, input *clouddirectory.ListDirectoriesInput, fn func(*clouddirectory.ListDirectoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListDirectoriesOutput{}
	fnCacher := func(out *clouddirectory.ListDirectoriesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListDirectoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListDirectoriesWithContext(ctx context.Context, input *clouddirectory.ListDirectoriesInput, opts ...request.Option) (*clouddirectory.ListDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListDirectoriesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListDirectoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListFacetAttributes(input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListFacetAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListFacetAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListFacetAttributesPages(input *clouddirectory.ListFacetAttributesInput, fn func(*clouddirectory.ListFacetAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListFacetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListFacetAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListFacetAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListFacetAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListFacetAttributesPagesWithContext(ctx context.Context, input *clouddirectory.ListFacetAttributesInput, fn func(*clouddirectory.ListFacetAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListFacetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListFacetAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListFacetAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListFacetAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListFacetAttributesWithContext(ctx context.Context, input *clouddirectory.ListFacetAttributesInput, opts ...request.Option) (*clouddirectory.ListFacetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListFacetAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListFacetAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListFacetNames(input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListFacetNamesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListFacetNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListFacetNamesPages(input *clouddirectory.ListFacetNamesInput, fn func(*clouddirectory.ListFacetNamesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListFacetNamesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListFacetNamesOutput{}
	fnCacher := func(out *clouddirectory.ListFacetNamesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListFacetNamesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListFacetNamesPagesWithContext(ctx context.Context, input *clouddirectory.ListFacetNamesInput, fn func(*clouddirectory.ListFacetNamesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListFacetNamesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListFacetNamesOutput{}
	fnCacher := func(out *clouddirectory.ListFacetNamesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListFacetNamesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListFacetNamesWithContext(ctx context.Context, input *clouddirectory.ListFacetNamesInput, opts ...request.Option) (*clouddirectory.ListFacetNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListFacetNamesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListFacetNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListIncomingTypedLinks(input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListIncomingTypedLinksOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListIncomingTypedLinks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListIncomingTypedLinksWithContext(ctx context.Context, input *clouddirectory.ListIncomingTypedLinksInput, opts ...request.Option) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListIncomingTypedLinksOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListIncomingTypedLinksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListIndex(input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListIndexOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListIndexPages(input *clouddirectory.ListIndexInput, fn func(*clouddirectory.ListIndexOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListIndexOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListIndexOutput{}
	fnCacher := func(out *clouddirectory.ListIndexOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListIndexPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListIndexPagesWithContext(ctx context.Context, input *clouddirectory.ListIndexInput, fn func(*clouddirectory.ListIndexOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListIndexOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListIndexOutput{}
	fnCacher := func(out *clouddirectory.ListIndexOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListIndexPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListIndexWithContext(ctx context.Context, input *clouddirectory.ListIndexInput, opts ...request.Option) (*clouddirectory.ListIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListIndexOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListManagedSchemaArns(input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListManagedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListManagedSchemaArns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListManagedSchemaArnsPages(input *clouddirectory.ListManagedSchemaArnsInput, fn func(*clouddirectory.ListManagedSchemaArnsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListManagedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListManagedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListManagedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListManagedSchemaArnsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListManagedSchemaArnsPagesWithContext(ctx context.Context, input *clouddirectory.ListManagedSchemaArnsInput, fn func(*clouddirectory.ListManagedSchemaArnsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListManagedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListManagedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListManagedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListManagedSchemaArnsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListManagedSchemaArnsWithContext(ctx context.Context, input *clouddirectory.ListManagedSchemaArnsInput, opts ...request.Option) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListManagedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListManagedSchemaArnsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectAttributes(input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectAttributesPages(input *clouddirectory.ListObjectAttributesInput, fn func(*clouddirectory.ListObjectAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListObjectAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectAttributesPagesWithContext(ctx context.Context, input *clouddirectory.ListObjectAttributesInput, fn func(*clouddirectory.ListObjectAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListObjectAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectAttributesWithContext(ctx context.Context, input *clouddirectory.ListObjectAttributesInput, opts ...request.Option) (*clouddirectory.ListObjectAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectChildren(input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectChildrenOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectChildren(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectChildrenPages(input *clouddirectory.ListObjectChildrenInput, fn func(*clouddirectory.ListObjectChildrenOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectChildrenOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectChildrenOutput{}
	fnCacher := func(out *clouddirectory.ListObjectChildrenOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectChildrenPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectChildrenPagesWithContext(ctx context.Context, input *clouddirectory.ListObjectChildrenInput, fn func(*clouddirectory.ListObjectChildrenOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectChildrenOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectChildrenOutput{}
	fnCacher := func(out *clouddirectory.ListObjectChildrenOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectChildrenPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectChildrenWithContext(ctx context.Context, input *clouddirectory.ListObjectChildrenInput, opts ...request.Option) (*clouddirectory.ListObjectChildrenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectChildrenOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectChildrenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectParentPaths(input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectParentPathsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectParentPaths(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectParentPathsPages(input *clouddirectory.ListObjectParentPathsInput, fn func(*clouddirectory.ListObjectParentPathsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectParentPathsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectParentPathsOutput{}
	fnCacher := func(out *clouddirectory.ListObjectParentPathsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectParentPathsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectParentPathsPagesWithContext(ctx context.Context, input *clouddirectory.ListObjectParentPathsInput, fn func(*clouddirectory.ListObjectParentPathsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectParentPathsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectParentPathsOutput{}
	fnCacher := func(out *clouddirectory.ListObjectParentPathsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectParentPathsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectParentPathsWithContext(ctx context.Context, input *clouddirectory.ListObjectParentPathsInput, opts ...request.Option) (*clouddirectory.ListObjectParentPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectParentPathsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectParentPathsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectParents(input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectParentsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectParents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectParentsPages(input *clouddirectory.ListObjectParentsInput, fn func(*clouddirectory.ListObjectParentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectParentsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectParentsOutput{}
	fnCacher := func(out *clouddirectory.ListObjectParentsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectParentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectParentsPagesWithContext(ctx context.Context, input *clouddirectory.ListObjectParentsInput, fn func(*clouddirectory.ListObjectParentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectParentsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectParentsOutput{}
	fnCacher := func(out *clouddirectory.ListObjectParentsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectParentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectParentsWithContext(ctx context.Context, input *clouddirectory.ListObjectParentsInput, opts ...request.Option) (*clouddirectory.ListObjectParentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectParentsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectParentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectPolicies(input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectPoliciesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListObjectPoliciesPages(input *clouddirectory.ListObjectPoliciesInput, fn func(*clouddirectory.ListObjectPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectPoliciesOutput{}
	fnCacher := func(out *clouddirectory.ListObjectPoliciesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectPoliciesPagesWithContext(ctx context.Context, input *clouddirectory.ListObjectPoliciesInput, fn func(*clouddirectory.ListObjectPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListObjectPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListObjectPoliciesOutput{}
	fnCacher := func(out *clouddirectory.ListObjectPoliciesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListObjectPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListObjectPoliciesWithContext(ctx context.Context, input *clouddirectory.ListObjectPoliciesInput, opts ...request.Option) (*clouddirectory.ListObjectPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListObjectPoliciesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListObjectPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListOutgoingTypedLinks(input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListOutgoingTypedLinksOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListOutgoingTypedLinks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListOutgoingTypedLinksWithContext(ctx context.Context, input *clouddirectory.ListOutgoingTypedLinksInput, opts ...request.Option) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListOutgoingTypedLinksOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListOutgoingTypedLinksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListPolicyAttachments(input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListPolicyAttachmentsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListPolicyAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListPolicyAttachmentsPages(input *clouddirectory.ListPolicyAttachmentsInput, fn func(*clouddirectory.ListPolicyAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListPolicyAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListPolicyAttachmentsOutput{}
	fnCacher := func(out *clouddirectory.ListPolicyAttachmentsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListPolicyAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListPolicyAttachmentsPagesWithContext(ctx context.Context, input *clouddirectory.ListPolicyAttachmentsInput, fn func(*clouddirectory.ListPolicyAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListPolicyAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListPolicyAttachmentsOutput{}
	fnCacher := func(out *clouddirectory.ListPolicyAttachmentsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListPolicyAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListPolicyAttachmentsWithContext(ctx context.Context, input *clouddirectory.ListPolicyAttachmentsInput, opts ...request.Option) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListPolicyAttachmentsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListPolicyAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListPublishedSchemaArns(input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListPublishedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListPublishedSchemaArns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListPublishedSchemaArnsPages(input *clouddirectory.ListPublishedSchemaArnsInput, fn func(*clouddirectory.ListPublishedSchemaArnsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListPublishedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListPublishedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListPublishedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListPublishedSchemaArnsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListPublishedSchemaArnsPagesWithContext(ctx context.Context, input *clouddirectory.ListPublishedSchemaArnsInput, fn func(*clouddirectory.ListPublishedSchemaArnsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListPublishedSchemaArnsOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListPublishedSchemaArnsOutput{}
	fnCacher := func(out *clouddirectory.ListPublishedSchemaArnsOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListPublishedSchemaArnsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListPublishedSchemaArnsWithContext(ctx context.Context, input *clouddirectory.ListPublishedSchemaArnsInput, opts ...request.Option) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListPublishedSchemaArnsOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListPublishedSchemaArnsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTagsForResource(input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTagsForResourcePages(input *clouddirectory.ListTagsForResourceInput, fn func(*clouddirectory.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTagsForResourceOutput{}
	fnCacher := func(out *clouddirectory.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTagsForResourcePagesWithContext(ctx context.Context, input *clouddirectory.ListTagsForResourceInput, fn func(*clouddirectory.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTagsForResourceOutput{}
	fnCacher := func(out *clouddirectory.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTagsForResourceWithContext(ctx context.Context, input *clouddirectory.ListTagsForResourceInput, opts ...request.Option) (*clouddirectory.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTypedLinkFacetAttributes(input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTypedLinkFacetAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTypedLinkFacetAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTypedLinkFacetAttributesPages(input *clouddirectory.ListTypedLinkFacetAttributesInput, fn func(*clouddirectory.ListTypedLinkFacetAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTypedLinkFacetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTypedLinkFacetAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListTypedLinkFacetAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTypedLinkFacetAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTypedLinkFacetAttributesPagesWithContext(ctx context.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput, fn func(*clouddirectory.ListTypedLinkFacetAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTypedLinkFacetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTypedLinkFacetAttributesOutput{}
	fnCacher := func(out *clouddirectory.ListTypedLinkFacetAttributesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTypedLinkFacetAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTypedLinkFacetAttributesWithContext(ctx context.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput, opts ...request.Option) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTypedLinkFacetAttributesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTypedLinkFacetAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTypedLinkFacetNames(input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTypedLinkFacetNamesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTypedLinkFacetNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudDirectory) ListTypedLinkFacetNamesPages(input *clouddirectory.ListTypedLinkFacetNamesInput, fn func(*clouddirectory.ListTypedLinkFacetNamesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTypedLinkFacetNamesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTypedLinkFacetNamesOutput{}
	fnCacher := func(out *clouddirectory.ListTypedLinkFacetNamesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTypedLinkFacetNamesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTypedLinkFacetNamesPagesWithContext(ctx context.Context, input *clouddirectory.ListTypedLinkFacetNamesInput, fn func(*clouddirectory.ListTypedLinkFacetNamesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*clouddirectory.ListTypedLinkFacetNamesOutput), false)
		return nil
	}
	cachable := true
	output := &clouddirectory.ListTypedLinkFacetNamesOutput{}
	fnCacher := func(out *clouddirectory.ListTypedLinkFacetNamesOutput, more bool) bool {
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
	if err := c.CloudDirectoryAPI.ListTypedLinkFacetNamesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudDirectory) ListTypedLinkFacetNamesWithContext(ctx context.Context, input *clouddirectory.ListTypedLinkFacetNamesInput, opts ...request.Option) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*clouddirectory.ListTypedLinkFacetNamesOutput), nil
	}
	out, err := c.CloudDirectoryAPI.ListTypedLinkFacetNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

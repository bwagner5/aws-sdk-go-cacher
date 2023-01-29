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
package connectwisdomservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/connectwisdomservice"
	"github.com/aws/aws-sdk-go/service/connectwisdomservice/connectwisdomserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConnectWisdomService struct {
	connectwisdomserviceiface.ConnectWisdomServiceAPI
	cache *cache.Cache
}

func New(connectwisdomserviceapi connectwisdomserviceiface.ConnectWisdomServiceAPI) *ConnectWisdomService {
	return &ConnectWisdomService{
		ConnectWisdomServiceAPI: connectwisdomserviceapi,
		cache:                   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConnectWisdomService) GetAssistant(input *connectwisdomservice.GetAssistantInput) (*connectwisdomservice.GetAssistantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetAssistantOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetAssistant(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetAssistantAssociation(input *connectwisdomservice.GetAssistantAssociationInput) (*connectwisdomservice.GetAssistantAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetAssistantAssociationOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetAssistantAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetAssistantAssociationWithContext(ctx context.Context, input *connectwisdomservice.GetAssistantAssociationInput, opts ...request.Option) (*connectwisdomservice.GetAssistantAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetAssistantAssociationOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetAssistantAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetAssistantWithContext(ctx context.Context, input *connectwisdomservice.GetAssistantInput, opts ...request.Option) (*connectwisdomservice.GetAssistantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetAssistantOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetAssistantWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetContent(input *connectwisdomservice.GetContentInput) (*connectwisdomservice.GetContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetContentOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetContentSummary(input *connectwisdomservice.GetContentSummaryInput) (*connectwisdomservice.GetContentSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetContentSummaryOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetContentSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetContentSummaryWithContext(ctx context.Context, input *connectwisdomservice.GetContentSummaryInput, opts ...request.Option) (*connectwisdomservice.GetContentSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetContentSummaryOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetContentSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetContentWithContext(ctx context.Context, input *connectwisdomservice.GetContentInput, opts ...request.Option) (*connectwisdomservice.GetContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetContentOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetKnowledgeBase(input *connectwisdomservice.GetKnowledgeBaseInput) (*connectwisdomservice.GetKnowledgeBaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetKnowledgeBaseOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetKnowledgeBase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetKnowledgeBaseWithContext(ctx context.Context, input *connectwisdomservice.GetKnowledgeBaseInput, opts ...request.Option) (*connectwisdomservice.GetKnowledgeBaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetKnowledgeBaseOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetKnowledgeBaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetRecommendations(input *connectwisdomservice.GetRecommendationsInput) (*connectwisdomservice.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetRecommendationsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetRecommendationsWithContext(ctx context.Context, input *connectwisdomservice.GetRecommendationsInput, opts ...request.Option) (*connectwisdomservice.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetRecommendationsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetSession(input *connectwisdomservice.GetSessionInput) (*connectwisdomservice.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetSessionOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) GetSessionWithContext(ctx context.Context, input *connectwisdomservice.GetSessionInput, opts ...request.Option) (*connectwisdomservice.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.GetSessionOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.GetSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListAssistantAssociations(input *connectwisdomservice.ListAssistantAssociationsInput) (*connectwisdomservice.ListAssistantAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListAssistantAssociationsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListAssistantAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListAssistantAssociationsPages(input *connectwisdomservice.ListAssistantAssociationsInput, fn func(*connectwisdomservice.ListAssistantAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListAssistantAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListAssistantAssociationsOutput{}
	fnCacher := func(out *connectwisdomservice.ListAssistantAssociationsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListAssistantAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListAssistantAssociationsPagesWithContext(ctx context.Context, input *connectwisdomservice.ListAssistantAssociationsInput, fn func(*connectwisdomservice.ListAssistantAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListAssistantAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListAssistantAssociationsOutput{}
	fnCacher := func(out *connectwisdomservice.ListAssistantAssociationsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListAssistantAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListAssistantAssociationsWithContext(ctx context.Context, input *connectwisdomservice.ListAssistantAssociationsInput, opts ...request.Option) (*connectwisdomservice.ListAssistantAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListAssistantAssociationsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListAssistantAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListAssistants(input *connectwisdomservice.ListAssistantsInput) (*connectwisdomservice.ListAssistantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListAssistantsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListAssistants(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListAssistantsPages(input *connectwisdomservice.ListAssistantsInput, fn func(*connectwisdomservice.ListAssistantsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListAssistantsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListAssistantsOutput{}
	fnCacher := func(out *connectwisdomservice.ListAssistantsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListAssistantsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListAssistantsPagesWithContext(ctx context.Context, input *connectwisdomservice.ListAssistantsInput, fn func(*connectwisdomservice.ListAssistantsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListAssistantsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListAssistantsOutput{}
	fnCacher := func(out *connectwisdomservice.ListAssistantsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListAssistantsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListAssistantsWithContext(ctx context.Context, input *connectwisdomservice.ListAssistantsInput, opts ...request.Option) (*connectwisdomservice.ListAssistantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListAssistantsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListAssistantsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListContents(input *connectwisdomservice.ListContentsInput) (*connectwisdomservice.ListContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListContentsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListContents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListContentsPages(input *connectwisdomservice.ListContentsInput, fn func(*connectwisdomservice.ListContentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListContentsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListContentsOutput{}
	fnCacher := func(out *connectwisdomservice.ListContentsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListContentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListContentsPagesWithContext(ctx context.Context, input *connectwisdomservice.ListContentsInput, fn func(*connectwisdomservice.ListContentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListContentsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListContentsOutput{}
	fnCacher := func(out *connectwisdomservice.ListContentsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListContentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListContentsWithContext(ctx context.Context, input *connectwisdomservice.ListContentsInput, opts ...request.Option) (*connectwisdomservice.ListContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListContentsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListContentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListKnowledgeBases(input *connectwisdomservice.ListKnowledgeBasesInput) (*connectwisdomservice.ListKnowledgeBasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListKnowledgeBasesOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListKnowledgeBases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListKnowledgeBasesPages(input *connectwisdomservice.ListKnowledgeBasesInput, fn func(*connectwisdomservice.ListKnowledgeBasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListKnowledgeBasesOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListKnowledgeBasesOutput{}
	fnCacher := func(out *connectwisdomservice.ListKnowledgeBasesOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListKnowledgeBasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListKnowledgeBasesPagesWithContext(ctx context.Context, input *connectwisdomservice.ListKnowledgeBasesInput, fn func(*connectwisdomservice.ListKnowledgeBasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.ListKnowledgeBasesOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.ListKnowledgeBasesOutput{}
	fnCacher := func(out *connectwisdomservice.ListKnowledgeBasesOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.ListKnowledgeBasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) ListKnowledgeBasesWithContext(ctx context.Context, input *connectwisdomservice.ListKnowledgeBasesInput, opts ...request.Option) (*connectwisdomservice.ListKnowledgeBasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListKnowledgeBasesOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListKnowledgeBasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListTagsForResource(input *connectwisdomservice.ListTagsForResourceInput) (*connectwisdomservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) ListTagsForResourceWithContext(ctx context.Context, input *connectwisdomservice.ListTagsForResourceInput, opts ...request.Option) (*connectwisdomservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) QueryAssistant(input *connectwisdomservice.QueryAssistantInput) (*connectwisdomservice.QueryAssistantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.QueryAssistantOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.QueryAssistant(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) QueryAssistantPages(input *connectwisdomservice.QueryAssistantInput, fn func(*connectwisdomservice.QueryAssistantOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.QueryAssistantOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.QueryAssistantOutput{}
	fnCacher := func(out *connectwisdomservice.QueryAssistantOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.QueryAssistantPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) QueryAssistantPagesWithContext(ctx context.Context, input *connectwisdomservice.QueryAssistantInput, fn func(*connectwisdomservice.QueryAssistantOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.QueryAssistantOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.QueryAssistantOutput{}
	fnCacher := func(out *connectwisdomservice.QueryAssistantOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.QueryAssistantPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) QueryAssistantWithContext(ctx context.Context, input *connectwisdomservice.QueryAssistantInput, opts ...request.Option) (*connectwisdomservice.QueryAssistantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.QueryAssistantOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.QueryAssistantWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) SearchContent(input *connectwisdomservice.SearchContentInput) (*connectwisdomservice.SearchContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.SearchContentOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.SearchContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) SearchContentPages(input *connectwisdomservice.SearchContentInput, fn func(*connectwisdomservice.SearchContentOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.SearchContentOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.SearchContentOutput{}
	fnCacher := func(out *connectwisdomservice.SearchContentOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.SearchContentPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) SearchContentPagesWithContext(ctx context.Context, input *connectwisdomservice.SearchContentInput, fn func(*connectwisdomservice.SearchContentOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.SearchContentOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.SearchContentOutput{}
	fnCacher := func(out *connectwisdomservice.SearchContentOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.SearchContentPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) SearchContentWithContext(ctx context.Context, input *connectwisdomservice.SearchContentInput, opts ...request.Option) (*connectwisdomservice.SearchContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.SearchContentOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.SearchContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) SearchSessions(input *connectwisdomservice.SearchSessionsInput) (*connectwisdomservice.SearchSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.SearchSessionsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.SearchSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectWisdomService) SearchSessionsPages(input *connectwisdomservice.SearchSessionsInput, fn func(*connectwisdomservice.SearchSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.SearchSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.SearchSessionsOutput{}
	fnCacher := func(out *connectwisdomservice.SearchSessionsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.SearchSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) SearchSessionsPagesWithContext(ctx context.Context, input *connectwisdomservice.SearchSessionsInput, fn func(*connectwisdomservice.SearchSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectwisdomservice.SearchSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &connectwisdomservice.SearchSessionsOutput{}
	fnCacher := func(out *connectwisdomservice.SearchSessionsOutput, more bool) bool {
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
	if err := c.ConnectWisdomServiceAPI.SearchSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectWisdomService) SearchSessionsWithContext(ctx context.Context, input *connectwisdomservice.SearchSessionsInput, opts ...request.Option) (*connectwisdomservice.SearchSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectwisdomservice.SearchSessionsOutput), nil
	}
	out, err := c.ConnectWisdomServiceAPI.SearchSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package kinesisvideoarchivedmediacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisVideoArchivedMedia struct {
	kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI
	cache *cache.Cache
}

func New(kinesisvideoarchivedmediaapi kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI) *KinesisVideoArchivedMedia {
	return &KinesisVideoArchivedMedia{
		KinesisVideoArchivedMediaAPI: kinesisvideoarchivedmediaapi,
		cache:                        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisVideoArchivedMedia) GetClip(input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetClipOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetClip(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetClipWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetClipInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetClipOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetClipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetDASHStreamingSessionURL(input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetDASHStreamingSessionURL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetDASHStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetDASHStreamingSessionURLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetHLSStreamingSessionURL(input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetHLSStreamingSessionURL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetHLSStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetHLSStreamingSessionURLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetImages(input *kinesisvideoarchivedmedia.GetImagesInput) (*kinesisvideoarchivedmedia.GetImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetImagesOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetImagesPages(input *kinesisvideoarchivedmedia.GetImagesInput, fn func(*kinesisvideoarchivedmedia.GetImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideoarchivedmedia.GetImagesOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideoarchivedmedia.GetImagesOutput{}
	fnCacher := func(out *kinesisvideoarchivedmedia.GetImagesOutput, more bool) bool {
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
	if err := c.KinesisVideoArchivedMediaAPI.GetImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideoArchivedMedia) GetImagesPagesWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetImagesInput, fn func(*kinesisvideoarchivedmedia.GetImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideoarchivedmedia.GetImagesOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideoarchivedmedia.GetImagesOutput{}
	fnCacher := func(out *kinesisvideoarchivedmedia.GetImagesOutput, more bool) bool {
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
	if err := c.KinesisVideoArchivedMediaAPI.GetImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideoArchivedMedia) GetImagesWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetImagesInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetImagesOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentList(input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetMediaForFragmentList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentListWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.GetMediaForFragmentListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) ListFragments(input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.ListFragmentsOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.ListFragments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoArchivedMedia) ListFragmentsPages(input *kinesisvideoarchivedmedia.ListFragmentsInput, fn func(*kinesisvideoarchivedmedia.ListFragmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideoarchivedmedia.ListFragmentsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideoarchivedmedia.ListFragmentsOutput{}
	fnCacher := func(out *kinesisvideoarchivedmedia.ListFragmentsOutput, more bool) bool {
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
	if err := c.KinesisVideoArchivedMediaAPI.ListFragmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideoArchivedMedia) ListFragmentsPagesWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput, fn func(*kinesisvideoarchivedmedia.ListFragmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideoarchivedmedia.ListFragmentsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideoarchivedmedia.ListFragmentsOutput{}
	fnCacher := func(out *kinesisvideoarchivedmedia.ListFragmentsOutput, more bool) bool {
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
	if err := c.KinesisVideoArchivedMediaAPI.ListFragmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideoArchivedMedia) ListFragmentsWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput, opts ...request.Option) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideoarchivedmedia.ListFragmentsOutput), nil
	}
	out, err := c.KinesisVideoArchivedMediaAPI.ListFragmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

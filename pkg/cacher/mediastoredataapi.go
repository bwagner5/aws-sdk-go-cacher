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
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediastoredata/mediastoredataiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaStoreData struct {
	mediastoredataiface.MediaStoreDataAPI
	cache *cache.Cache
}

func NewMediaStoreData(mediastoredataapi mediastoredataiface.MediaStoreDataAPI) *MediaStoreData {
	return &MediaStoreData{
		MediaStoreDataAPI: mediastoredataapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaStoreData) DescribeObject(input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.DescribeObjectOutput), nil
	}
	out, err := c.MediaStoreDataAPI.DescribeObject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStoreData) DescribeObjectWithContext(ctx context.Context, input *mediastoredata.DescribeObjectInput, opts ...request.Option) (*mediastoredata.DescribeObjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.DescribeObjectOutput), nil
	}
	out, err := c.MediaStoreDataAPI.DescribeObjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStoreData) GetObject(input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.GetObjectOutput), nil
	}
	out, err := c.MediaStoreDataAPI.GetObject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStoreData) GetObjectWithContext(ctx context.Context, input *mediastoredata.GetObjectInput, opts ...request.Option) (*mediastoredata.GetObjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.GetObjectOutput), nil
	}
	out, err := c.MediaStoreDataAPI.GetObjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStoreData) ListItems(input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.ListItemsOutput), nil
	}
	out, err := c.MediaStoreDataAPI.ListItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaStoreData) ListItemsPages(input *mediastoredata.ListItemsInput, fn func(*mediastoredata.ListItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediastoredata.ListItemsOutput), false)
		return nil
	}
	cachable := true
	output := &mediastoredata.ListItemsOutput{}
	fnCacher := func(out *mediastoredata.ListItemsOutput, more bool) bool {
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
	if err := c.MediaStoreDataAPI.ListItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaStoreData) ListItemsPagesWithContext(ctx context.Context, input *mediastoredata.ListItemsInput, fn func(*mediastoredata.ListItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediastoredata.ListItemsOutput), false)
		return nil
	}
	cachable := true
	output := &mediastoredata.ListItemsOutput{}
	fnCacher := func(out *mediastoredata.ListItemsOutput, more bool) bool {
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
	if err := c.MediaStoreDataAPI.ListItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaStoreData) ListItemsWithContext(ctx context.Context, input *mediastoredata.ListItemsInput, opts ...request.Option) (*mediastoredata.ListItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediastoredata.ListItemsOutput), nil
	}
	out, err := c.MediaStoreDataAPI.ListItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

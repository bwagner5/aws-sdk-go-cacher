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
package kinesisvideomediacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisVideoMedia struct {
	kinesisvideomediaiface.KinesisVideoMediaAPI
	cache *cache.Cache
}

func New(kinesisvideomediaapi kinesisvideomediaiface.KinesisVideoMediaAPI) *KinesisVideoMedia {
	return &KinesisVideoMedia{
		KinesisVideoMediaAPI: kinesisvideomediaapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisVideoMedia) GetMedia(input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideomedia.GetMediaOutput), nil
	}
	out, err := c.KinesisVideoMediaAPI.GetMedia(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoMedia) GetMediaWithContext(ctx context.Context, input *kinesisvideomedia.GetMediaInput, opts ...request.Option) (*kinesisvideomedia.GetMediaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideomedia.GetMediaOutput), nil
	}
	out, err := c.KinesisVideoMediaAPI.GetMediaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

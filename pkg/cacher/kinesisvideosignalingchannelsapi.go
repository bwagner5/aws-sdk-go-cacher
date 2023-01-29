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
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels/kinesisvideosignalingchannelsiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisVideoSignalingChannels struct {
	kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI
	cache *cache.Cache
}

func NewKinesisVideoSignalingChannels(kinesisvideosignalingchannelsapi kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI) *KinesisVideoSignalingChannels {
	return &KinesisVideoSignalingChannels{
		KinesisVideoSignalingChannelsAPI: kinesisvideosignalingchannelsapi,
		cache:                            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisVideoSignalingChannels) GetIceServerConfig(input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideosignalingchannels.GetIceServerConfigOutput), nil
	}
	out, err := c.KinesisVideoSignalingChannelsAPI.GetIceServerConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideoSignalingChannels) GetIceServerConfigWithContext(ctx context.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput, opts ...request.Option) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideosignalingchannels.GetIceServerConfigOutput), nil
	}
	out, err := c.KinesisVideoSignalingChannelsAPI.GetIceServerConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConnectParticipant struct {
	connectparticipantiface.ConnectParticipantAPI
	cache *cache.Cache
}

func NewConnectParticipant(connectparticipantapi connectparticipantiface.ConnectParticipantAPI) *ConnectParticipant {
	return &ConnectParticipant{
		ConnectParticipantAPI: connectparticipantapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConnectParticipant) GetAttachment(input *connectparticipant.GetAttachmentInput) (*connectparticipant.GetAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectparticipant.GetAttachmentOutput), nil
	}
	out, err := c.ConnectParticipantAPI.GetAttachment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectParticipant) GetAttachmentWithContext(ctx context.Context, input *connectparticipant.GetAttachmentInput, opts ...request.Option) (*connectparticipant.GetAttachmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectparticipant.GetAttachmentOutput), nil
	}
	out, err := c.ConnectParticipantAPI.GetAttachmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectParticipant) GetTranscript(input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectparticipant.GetTranscriptOutput), nil
	}
	out, err := c.ConnectParticipantAPI.GetTranscript(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectParticipant) GetTranscriptPages(input *connectparticipant.GetTranscriptInput, fn func(*connectparticipant.GetTranscriptOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectparticipant.GetTranscriptOutput), false)
		return nil
	}
	cachable := true
	output := &connectparticipant.GetTranscriptOutput{}
	fnCacher := func(out *connectparticipant.GetTranscriptOutput, more bool) bool {
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
	if err := c.ConnectParticipantAPI.GetTranscriptPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectParticipant) GetTranscriptPagesWithContext(ctx context.Context, input *connectparticipant.GetTranscriptInput, fn func(*connectparticipant.GetTranscriptOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectparticipant.GetTranscriptOutput), false)
		return nil
	}
	cachable := true
	output := &connectparticipant.GetTranscriptOutput{}
	fnCacher := func(out *connectparticipant.GetTranscriptOutput, more bool) bool {
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
	if err := c.ConnectParticipantAPI.GetTranscriptPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectParticipant) GetTranscriptWithContext(ctx context.Context, input *connectparticipant.GetTranscriptInput, opts ...request.Option) (*connectparticipant.GetTranscriptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectparticipant.GetTranscriptOutput), nil
	}
	out, err := c.ConnectParticipantAPI.GetTranscriptWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
	"github.com/aws/aws-sdk-go/service/voiceid"
	"github.com/aws/aws-sdk-go/service/voiceid/voiceidiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type VoiceID struct {
	voiceidiface.VoiceIDAPI
	cache *cache.Cache
}

func NewVoiceID(voiceidapi voiceidiface.VoiceIDAPI) *VoiceID {
	return &VoiceID{
		VoiceIDAPI: voiceidapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *VoiceID) DescribeDomain(input *voiceid.DescribeDomainInput) (*voiceid.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeDomainOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeDomainWithContext(ctx context.Context, input *voiceid.DescribeDomainInput, opts ...request.Option) (*voiceid.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeDomainOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeFraudster(input *voiceid.DescribeFraudsterInput) (*voiceid.DescribeFraudsterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeFraudsterOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeFraudster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeFraudsterRegistrationJob(input *voiceid.DescribeFraudsterRegistrationJobInput) (*voiceid.DescribeFraudsterRegistrationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeFraudsterRegistrationJobOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeFraudsterRegistrationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeFraudsterRegistrationJobWithContext(ctx context.Context, input *voiceid.DescribeFraudsterRegistrationJobInput, opts ...request.Option) (*voiceid.DescribeFraudsterRegistrationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeFraudsterRegistrationJobOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeFraudsterRegistrationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeFraudsterWithContext(ctx context.Context, input *voiceid.DescribeFraudsterInput, opts ...request.Option) (*voiceid.DescribeFraudsterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeFraudsterOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeFraudsterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeSpeaker(input *voiceid.DescribeSpeakerInput) (*voiceid.DescribeSpeakerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeSpeakerOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeSpeaker(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeSpeakerEnrollmentJob(input *voiceid.DescribeSpeakerEnrollmentJobInput) (*voiceid.DescribeSpeakerEnrollmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeSpeakerEnrollmentJobOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeSpeakerEnrollmentJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeSpeakerEnrollmentJobWithContext(ctx context.Context, input *voiceid.DescribeSpeakerEnrollmentJobInput, opts ...request.Option) (*voiceid.DescribeSpeakerEnrollmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeSpeakerEnrollmentJobOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeSpeakerEnrollmentJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) DescribeSpeakerWithContext(ctx context.Context, input *voiceid.DescribeSpeakerInput, opts ...request.Option) (*voiceid.DescribeSpeakerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.DescribeSpeakerOutput), nil
	}
	out, err := c.VoiceIDAPI.DescribeSpeakerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListDomains(input *voiceid.ListDomainsInput) (*voiceid.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListDomainsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListDomainsPages(input *voiceid.ListDomainsInput, fn func(*voiceid.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListDomainsOutput{}
	fnCacher := func(out *voiceid.ListDomainsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListDomainsPagesWithContext(ctx context.Context, input *voiceid.ListDomainsInput, fn func(*voiceid.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListDomainsOutput{}
	fnCacher := func(out *voiceid.ListDomainsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListDomainsWithContext(ctx context.Context, input *voiceid.ListDomainsInput, opts ...request.Option) (*voiceid.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListDomainsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListFraudsterRegistrationJobs(input *voiceid.ListFraudsterRegistrationJobsInput) (*voiceid.ListFraudsterRegistrationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListFraudsterRegistrationJobsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListFraudsterRegistrationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListFraudsterRegistrationJobsPages(input *voiceid.ListFraudsterRegistrationJobsInput, fn func(*voiceid.ListFraudsterRegistrationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListFraudsterRegistrationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListFraudsterRegistrationJobsOutput{}
	fnCacher := func(out *voiceid.ListFraudsterRegistrationJobsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListFraudsterRegistrationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListFraudsterRegistrationJobsPagesWithContext(ctx context.Context, input *voiceid.ListFraudsterRegistrationJobsInput, fn func(*voiceid.ListFraudsterRegistrationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListFraudsterRegistrationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListFraudsterRegistrationJobsOutput{}
	fnCacher := func(out *voiceid.ListFraudsterRegistrationJobsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListFraudsterRegistrationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListFraudsterRegistrationJobsWithContext(ctx context.Context, input *voiceid.ListFraudsterRegistrationJobsInput, opts ...request.Option) (*voiceid.ListFraudsterRegistrationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListFraudsterRegistrationJobsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListFraudsterRegistrationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListSpeakerEnrollmentJobs(input *voiceid.ListSpeakerEnrollmentJobsInput) (*voiceid.ListSpeakerEnrollmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListSpeakerEnrollmentJobsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListSpeakerEnrollmentJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListSpeakerEnrollmentJobsPages(input *voiceid.ListSpeakerEnrollmentJobsInput, fn func(*voiceid.ListSpeakerEnrollmentJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListSpeakerEnrollmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListSpeakerEnrollmentJobsOutput{}
	fnCacher := func(out *voiceid.ListSpeakerEnrollmentJobsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListSpeakerEnrollmentJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListSpeakerEnrollmentJobsPagesWithContext(ctx context.Context, input *voiceid.ListSpeakerEnrollmentJobsInput, fn func(*voiceid.ListSpeakerEnrollmentJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListSpeakerEnrollmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListSpeakerEnrollmentJobsOutput{}
	fnCacher := func(out *voiceid.ListSpeakerEnrollmentJobsOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListSpeakerEnrollmentJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListSpeakerEnrollmentJobsWithContext(ctx context.Context, input *voiceid.ListSpeakerEnrollmentJobsInput, opts ...request.Option) (*voiceid.ListSpeakerEnrollmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListSpeakerEnrollmentJobsOutput), nil
	}
	out, err := c.VoiceIDAPI.ListSpeakerEnrollmentJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListSpeakers(input *voiceid.ListSpeakersInput) (*voiceid.ListSpeakersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListSpeakersOutput), nil
	}
	out, err := c.VoiceIDAPI.ListSpeakers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListSpeakersPages(input *voiceid.ListSpeakersInput, fn func(*voiceid.ListSpeakersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListSpeakersOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListSpeakersOutput{}
	fnCacher := func(out *voiceid.ListSpeakersOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListSpeakersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListSpeakersPagesWithContext(ctx context.Context, input *voiceid.ListSpeakersInput, fn func(*voiceid.ListSpeakersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*voiceid.ListSpeakersOutput), false)
		return nil
	}
	cachable := true
	output := &voiceid.ListSpeakersOutput{}
	fnCacher := func(out *voiceid.ListSpeakersOutput, more bool) bool {
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
	if err := c.VoiceIDAPI.ListSpeakersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *VoiceID) ListSpeakersWithContext(ctx context.Context, input *voiceid.ListSpeakersInput, opts ...request.Option) (*voiceid.ListSpeakersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListSpeakersOutput), nil
	}
	out, err := c.VoiceIDAPI.ListSpeakersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListTagsForResource(input *voiceid.ListTagsForResourceInput) (*voiceid.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListTagsForResourceOutput), nil
	}
	out, err := c.VoiceIDAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *VoiceID) ListTagsForResourceWithContext(ctx context.Context, input *voiceid.ListTagsForResourceInput, opts ...request.Option) (*voiceid.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*voiceid.ListTagsForResourceOutput), nil
	}
	out, err := c.VoiceIDAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

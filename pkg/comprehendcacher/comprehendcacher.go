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
package comprehendcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Comprehend struct {
	comprehendiface.ComprehendAPI
	cache *cache.Cache
}

func New(comprehendapi comprehendiface.ComprehendAPI) *Comprehend {
	return &Comprehend{
		ComprehendAPI: comprehendapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Comprehend) BatchDetectDominantLanguage(input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectDominantLanguageOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectDominantLanguage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectDominantLanguageWithContext(ctx context.Context, input *comprehend.BatchDetectDominantLanguageInput, opts ...request.Option) (*comprehend.BatchDetectDominantLanguageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectDominantLanguageOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectDominantLanguageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectEntities(input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectEntitiesOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectEntitiesWithContext(ctx context.Context, input *comprehend.BatchDetectEntitiesInput, opts ...request.Option) (*comprehend.BatchDetectEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectEntitiesOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectKeyPhrases(input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectKeyPhrasesOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectKeyPhrases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectKeyPhrasesWithContext(ctx context.Context, input *comprehend.BatchDetectKeyPhrasesInput, opts ...request.Option) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectKeyPhrasesOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectKeyPhrasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectSentiment(input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectSentimentOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectSentiment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectSentimentWithContext(ctx context.Context, input *comprehend.BatchDetectSentimentInput, opts ...request.Option) (*comprehend.BatchDetectSentimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectSentimentOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectSentimentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectSyntax(input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectSyntaxOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectSyntax(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectSyntaxWithContext(ctx context.Context, input *comprehend.BatchDetectSyntaxInput, opts ...request.Option) (*comprehend.BatchDetectSyntaxOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectSyntaxOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectSyntaxWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectTargetedSentiment(input *comprehend.BatchDetectTargetedSentimentInput) (*comprehend.BatchDetectTargetedSentimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectTargetedSentimentOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectTargetedSentiment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) BatchDetectTargetedSentimentWithContext(ctx context.Context, input *comprehend.BatchDetectTargetedSentimentInput, opts ...request.Option) (*comprehend.BatchDetectTargetedSentimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.BatchDetectTargetedSentimentOutput), nil
	}
	out, err := c.ComprehendAPI.BatchDetectTargetedSentimentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDocumentClassificationJob(input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDocumentClassificationJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDocumentClassificationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDocumentClassificationJobWithContext(ctx context.Context, input *comprehend.DescribeDocumentClassificationJobInput, opts ...request.Option) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDocumentClassificationJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDocumentClassificationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDocumentClassifier(input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDocumentClassifierOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDocumentClassifier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDocumentClassifierWithContext(ctx context.Context, input *comprehend.DescribeDocumentClassifierInput, opts ...request.Option) (*comprehend.DescribeDocumentClassifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDocumentClassifierOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDocumentClassifierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDominantLanguageDetectionJob(input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDominantLanguageDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDominantLanguageDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeDominantLanguageDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput, opts ...request.Option) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeDominantLanguageDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeDominantLanguageDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEndpoint(input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEndpointOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEndpointWithContext(ctx context.Context, input *comprehend.DescribeEndpointInput, opts ...request.Option) (*comprehend.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEndpointOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEntitiesDetectionJob(input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEntitiesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEntitiesDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEntitiesDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeEntitiesDetectionJobInput, opts ...request.Option) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEntitiesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEntitiesDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEntityRecognizer(input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEntityRecognizerOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEntityRecognizer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEntityRecognizerWithContext(ctx context.Context, input *comprehend.DescribeEntityRecognizerInput, opts ...request.Option) (*comprehend.DescribeEntityRecognizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEntityRecognizerOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEntityRecognizerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEventsDetectionJob(input *comprehend.DescribeEventsDetectionJobInput) (*comprehend.DescribeEventsDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEventsDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEventsDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeEventsDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeEventsDetectionJobInput, opts ...request.Option) (*comprehend.DescribeEventsDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeEventsDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeEventsDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeKeyPhrasesDetectionJob(input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeKeyPhrasesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeKeyPhrasesDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeKeyPhrasesDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput, opts ...request.Option) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeKeyPhrasesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeKeyPhrasesDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribePiiEntitiesDetectionJob(input *comprehend.DescribePiiEntitiesDetectionJobInput) (*comprehend.DescribePiiEntitiesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribePiiEntitiesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribePiiEntitiesDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribePiiEntitiesDetectionJobWithContext(ctx context.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput, opts ...request.Option) (*comprehend.DescribePiiEntitiesDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribePiiEntitiesDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribePiiEntitiesDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeResourcePolicy(input *comprehend.DescribeResourcePolicyInput) (*comprehend.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeResourcePolicyOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeResourcePolicyWithContext(ctx context.Context, input *comprehend.DescribeResourcePolicyInput, opts ...request.Option) (*comprehend.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeResourcePolicyOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeSentimentDetectionJob(input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeSentimentDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeSentimentDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeSentimentDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeSentimentDetectionJobInput, opts ...request.Option) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeSentimentDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeSentimentDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeTargetedSentimentDetectionJob(input *comprehend.DescribeTargetedSentimentDetectionJobInput) (*comprehend.DescribeTargetedSentimentDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeTargetedSentimentDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeTargetedSentimentDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeTargetedSentimentDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeTargetedSentimentDetectionJobInput, opts ...request.Option) (*comprehend.DescribeTargetedSentimentDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeTargetedSentimentDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeTargetedSentimentDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeTopicsDetectionJob(input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeTopicsDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeTopicsDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) DescribeTopicsDetectionJobWithContext(ctx context.Context, input *comprehend.DescribeTopicsDetectionJobInput, opts ...request.Option) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.DescribeTopicsDetectionJobOutput), nil
	}
	out, err := c.ComprehendAPI.DescribeTopicsDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassificationJobs(input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassificationJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassificationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassificationJobsPages(input *comprehend.ListDocumentClassificationJobsInput, fn func(*comprehend.ListDocumentClassificationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassificationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassificationJobsOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassificationJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassificationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassificationJobsPagesWithContext(ctx context.Context, input *comprehend.ListDocumentClassificationJobsInput, fn func(*comprehend.ListDocumentClassificationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassificationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassificationJobsOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassificationJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassificationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassificationJobsWithContext(ctx context.Context, input *comprehend.ListDocumentClassificationJobsInput, opts ...request.Option) (*comprehend.ListDocumentClassificationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassificationJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassificationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassifierSummaries(input *comprehend.ListDocumentClassifierSummariesInput) (*comprehend.ListDocumentClassifierSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassifierSummariesOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassifierSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassifierSummariesPages(input *comprehend.ListDocumentClassifierSummariesInput, fn func(*comprehend.ListDocumentClassifierSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassifierSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassifierSummariesOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassifierSummariesOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassifierSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassifierSummariesPagesWithContext(ctx context.Context, input *comprehend.ListDocumentClassifierSummariesInput, fn func(*comprehend.ListDocumentClassifierSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassifierSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassifierSummariesOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassifierSummariesOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassifierSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassifierSummariesWithContext(ctx context.Context, input *comprehend.ListDocumentClassifierSummariesInput, opts ...request.Option) (*comprehend.ListDocumentClassifierSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassifierSummariesOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassifierSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassifiers(input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassifiersOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDocumentClassifiersPages(input *comprehend.ListDocumentClassifiersInput, fn func(*comprehend.ListDocumentClassifiersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassifiersOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassifiersOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassifiersOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassifiersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassifiersPagesWithContext(ctx context.Context, input *comprehend.ListDocumentClassifiersInput, fn func(*comprehend.ListDocumentClassifiersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDocumentClassifiersOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDocumentClassifiersOutput{}
	fnCacher := func(out *comprehend.ListDocumentClassifiersOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDocumentClassifiersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDocumentClassifiersWithContext(ctx context.Context, input *comprehend.ListDocumentClassifiersInput, opts ...request.Option) (*comprehend.ListDocumentClassifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDocumentClassifiersOutput), nil
	}
	out, err := c.ComprehendAPI.ListDocumentClassifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDominantLanguageDetectionJobs(input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDominantLanguageDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListDominantLanguageDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListDominantLanguageDetectionJobsPages(input *comprehend.ListDominantLanguageDetectionJobsInput, fn func(*comprehend.ListDominantLanguageDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDominantLanguageDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDominantLanguageDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListDominantLanguageDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDominantLanguageDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDominantLanguageDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListDominantLanguageDetectionJobsInput, fn func(*comprehend.ListDominantLanguageDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListDominantLanguageDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListDominantLanguageDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListDominantLanguageDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListDominantLanguageDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListDominantLanguageDetectionJobsWithContext(ctx context.Context, input *comprehend.ListDominantLanguageDetectionJobsInput, opts ...request.Option) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListDominantLanguageDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListDominantLanguageDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEndpoints(input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEndpointsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEndpointsPages(input *comprehend.ListEndpointsInput, fn func(*comprehend.ListEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEndpointsOutput{}
	fnCacher := func(out *comprehend.ListEndpointsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEndpointsPagesWithContext(ctx context.Context, input *comprehend.ListEndpointsInput, fn func(*comprehend.ListEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEndpointsOutput{}
	fnCacher := func(out *comprehend.ListEndpointsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEndpointsWithContext(ctx context.Context, input *comprehend.ListEndpointsInput, opts ...request.Option) (*comprehend.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEndpointsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntitiesDetectionJobs(input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntitiesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntitiesDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntitiesDetectionJobsPages(input *comprehend.ListEntitiesDetectionJobsInput, fn func(*comprehend.ListEntitiesDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntitiesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntitiesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListEntitiesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntitiesDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntitiesDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListEntitiesDetectionJobsInput, fn func(*comprehend.ListEntitiesDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntitiesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntitiesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListEntitiesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntitiesDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntitiesDetectionJobsWithContext(ctx context.Context, input *comprehend.ListEntitiesDetectionJobsInput, opts ...request.Option) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntitiesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntitiesDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntityRecognizerSummaries(input *comprehend.ListEntityRecognizerSummariesInput) (*comprehend.ListEntityRecognizerSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntityRecognizerSummariesOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntityRecognizerSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntityRecognizerSummariesPages(input *comprehend.ListEntityRecognizerSummariesInput, fn func(*comprehend.ListEntityRecognizerSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntityRecognizerSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntityRecognizerSummariesOutput{}
	fnCacher := func(out *comprehend.ListEntityRecognizerSummariesOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntityRecognizerSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntityRecognizerSummariesPagesWithContext(ctx context.Context, input *comprehend.ListEntityRecognizerSummariesInput, fn func(*comprehend.ListEntityRecognizerSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntityRecognizerSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntityRecognizerSummariesOutput{}
	fnCacher := func(out *comprehend.ListEntityRecognizerSummariesOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntityRecognizerSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntityRecognizerSummariesWithContext(ctx context.Context, input *comprehend.ListEntityRecognizerSummariesInput, opts ...request.Option) (*comprehend.ListEntityRecognizerSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntityRecognizerSummariesOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntityRecognizerSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntityRecognizers(input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntityRecognizersOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntityRecognizers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEntityRecognizersPages(input *comprehend.ListEntityRecognizersInput, fn func(*comprehend.ListEntityRecognizersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntityRecognizersOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntityRecognizersOutput{}
	fnCacher := func(out *comprehend.ListEntityRecognizersOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntityRecognizersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntityRecognizersPagesWithContext(ctx context.Context, input *comprehend.ListEntityRecognizersInput, fn func(*comprehend.ListEntityRecognizersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEntityRecognizersOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEntityRecognizersOutput{}
	fnCacher := func(out *comprehend.ListEntityRecognizersOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEntityRecognizersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEntityRecognizersWithContext(ctx context.Context, input *comprehend.ListEntityRecognizersInput, opts ...request.Option) (*comprehend.ListEntityRecognizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEntityRecognizersOutput), nil
	}
	out, err := c.ComprehendAPI.ListEntityRecognizersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEventsDetectionJobs(input *comprehend.ListEventsDetectionJobsInput) (*comprehend.ListEventsDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEventsDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEventsDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListEventsDetectionJobsPages(input *comprehend.ListEventsDetectionJobsInput, fn func(*comprehend.ListEventsDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEventsDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEventsDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListEventsDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEventsDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEventsDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListEventsDetectionJobsInput, fn func(*comprehend.ListEventsDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListEventsDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListEventsDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListEventsDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListEventsDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListEventsDetectionJobsWithContext(ctx context.Context, input *comprehend.ListEventsDetectionJobsInput, opts ...request.Option) (*comprehend.ListEventsDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListEventsDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListEventsDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListKeyPhrasesDetectionJobs(input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListKeyPhrasesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListKeyPhrasesDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListKeyPhrasesDetectionJobsPages(input *comprehend.ListKeyPhrasesDetectionJobsInput, fn func(*comprehend.ListKeyPhrasesDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListKeyPhrasesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListKeyPhrasesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListKeyPhrasesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListKeyPhrasesDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListKeyPhrasesDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput, fn func(*comprehend.ListKeyPhrasesDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListKeyPhrasesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListKeyPhrasesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListKeyPhrasesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListKeyPhrasesDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListKeyPhrasesDetectionJobsWithContext(ctx context.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput, opts ...request.Option) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListKeyPhrasesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListKeyPhrasesDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListPiiEntitiesDetectionJobs(input *comprehend.ListPiiEntitiesDetectionJobsInput) (*comprehend.ListPiiEntitiesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListPiiEntitiesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListPiiEntitiesDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListPiiEntitiesDetectionJobsPages(input *comprehend.ListPiiEntitiesDetectionJobsInput, fn func(*comprehend.ListPiiEntitiesDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListPiiEntitiesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListPiiEntitiesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListPiiEntitiesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListPiiEntitiesDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListPiiEntitiesDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput, fn func(*comprehend.ListPiiEntitiesDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListPiiEntitiesDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListPiiEntitiesDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListPiiEntitiesDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListPiiEntitiesDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListPiiEntitiesDetectionJobsWithContext(ctx context.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput, opts ...request.Option) (*comprehend.ListPiiEntitiesDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListPiiEntitiesDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListPiiEntitiesDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListSentimentDetectionJobs(input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListSentimentDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListSentimentDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListSentimentDetectionJobsPages(input *comprehend.ListSentimentDetectionJobsInput, fn func(*comprehend.ListSentimentDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListSentimentDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListSentimentDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListSentimentDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListSentimentDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListSentimentDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListSentimentDetectionJobsInput, fn func(*comprehend.ListSentimentDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListSentimentDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListSentimentDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListSentimentDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListSentimentDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListSentimentDetectionJobsWithContext(ctx context.Context, input *comprehend.ListSentimentDetectionJobsInput, opts ...request.Option) (*comprehend.ListSentimentDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListSentimentDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListSentimentDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTagsForResource(input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTagsForResourceOutput), nil
	}
	out, err := c.ComprehendAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTagsForResourceWithContext(ctx context.Context, input *comprehend.ListTagsForResourceInput, opts ...request.Option) (*comprehend.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTagsForResourceOutput), nil
	}
	out, err := c.ComprehendAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTargetedSentimentDetectionJobs(input *comprehend.ListTargetedSentimentDetectionJobsInput) (*comprehend.ListTargetedSentimentDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTargetedSentimentDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListTargetedSentimentDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTargetedSentimentDetectionJobsPages(input *comprehend.ListTargetedSentimentDetectionJobsInput, fn func(*comprehend.ListTargetedSentimentDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListTargetedSentimentDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListTargetedSentimentDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListTargetedSentimentDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListTargetedSentimentDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListTargetedSentimentDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListTargetedSentimentDetectionJobsInput, fn func(*comprehend.ListTargetedSentimentDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListTargetedSentimentDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListTargetedSentimentDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListTargetedSentimentDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListTargetedSentimentDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListTargetedSentimentDetectionJobsWithContext(ctx context.Context, input *comprehend.ListTargetedSentimentDetectionJobsInput, opts ...request.Option) (*comprehend.ListTargetedSentimentDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTargetedSentimentDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListTargetedSentimentDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTopicsDetectionJobs(input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTopicsDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListTopicsDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Comprehend) ListTopicsDetectionJobsPages(input *comprehend.ListTopicsDetectionJobsInput, fn func(*comprehend.ListTopicsDetectionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListTopicsDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListTopicsDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListTopicsDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListTopicsDetectionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListTopicsDetectionJobsPagesWithContext(ctx context.Context, input *comprehend.ListTopicsDetectionJobsInput, fn func(*comprehend.ListTopicsDetectionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*comprehend.ListTopicsDetectionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &comprehend.ListTopicsDetectionJobsOutput{}
	fnCacher := func(out *comprehend.ListTopicsDetectionJobsOutput, more bool) bool {
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
	if err := c.ComprehendAPI.ListTopicsDetectionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Comprehend) ListTopicsDetectionJobsWithContext(ctx context.Context, input *comprehend.ListTopicsDetectionJobsInput, opts ...request.Option) (*comprehend.ListTopicsDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehend.ListTopicsDetectionJobsOutput), nil
	}
	out, err := c.ComprehendAPI.ListTopicsDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

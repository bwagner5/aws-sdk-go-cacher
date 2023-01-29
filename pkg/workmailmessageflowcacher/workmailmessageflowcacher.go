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
package workmailmessageflowcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkMailMessageFlow struct {
	workmailmessageflowiface.WorkMailMessageFlowAPI
	cache *cache.Cache
}

func New(workmailmessageflowapi workmailmessageflowiface.WorkMailMessageFlowAPI) *WorkMailMessageFlow {
	return &WorkMailMessageFlow{
		WorkMailMessageFlowAPI: workmailmessageflowapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkMailMessageFlow) GetRawMessageContent(input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmailmessageflow.GetRawMessageContentOutput), nil
	}
	out, err := c.WorkMailMessageFlowAPI.GetRawMessageContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMailMessageFlow) GetRawMessageContentWithContext(ctx context.Context, input *workmailmessageflow.GetRawMessageContentInput, opts ...request.Option) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmailmessageflow.GetRawMessageContentOutput), nil
	}
	out, err := c.WorkMailMessageFlowAPI.GetRawMessageContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

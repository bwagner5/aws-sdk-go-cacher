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
package lexruntimev2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lexruntimev2"
	"github.com/aws/aws-sdk-go/service/lexruntimev2/lexruntimev2iface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LexRuntimeV2 struct {
	lexruntimev2iface.LexRuntimeV2API
	cache *cache.Cache
}

func New(lexruntimev2api lexruntimev2iface.LexRuntimeV2API) *LexRuntimeV2 {
	return &LexRuntimeV2{
		LexRuntimeV2API: lexruntimev2api,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LexRuntimeV2) GetSession(input *lexruntimev2.GetSessionInput) (*lexruntimev2.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexruntimev2.GetSessionOutput), nil
	}
	out, err := c.LexRuntimeV2API.GetSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexRuntimeV2) GetSessionWithContext(ctx context.Context, input *lexruntimev2.GetSessionInput, opts ...request.Option) (*lexruntimev2.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexruntimev2.GetSessionOutput), nil
	}
	out, err := c.LexRuntimeV2API.GetSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

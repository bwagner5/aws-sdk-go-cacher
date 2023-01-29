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
package cloudsearchdomaincacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudSearchDomain struct {
	cloudsearchdomainiface.CloudSearchDomainAPI
	cache *cache.Cache
}

func New(cloudsearchdomainapi cloudsearchdomainiface.CloudSearchDomainAPI) *CloudSearchDomain {
	return &CloudSearchDomain{
		CloudSearchDomainAPI: cloudsearchdomainapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudSearchDomain) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearchdomain.SearchOutput), nil
	}
	out, err := c.CloudSearchDomainAPI.Search(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudSearchDomain) SearchWithContext(ctx context.Context, input *cloudsearchdomain.SearchInput, opts ...request.Option) (*cloudsearchdomain.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudsearchdomain.SearchOutput), nil
	}
	out, err := c.CloudSearchDomainAPI.SearchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

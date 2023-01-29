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
package rdsdataservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RDSDataService struct {
	rdsdataserviceiface.RDSDataServiceAPI
	cache *cache.Cache
}

func New(rdsdataserviceapi rdsdataserviceiface.RDSDataServiceAPI) *RDSDataService {
	return &RDSDataService{
		RDSDataServiceAPI: rdsdataserviceapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RDSDataService) BatchExecuteStatement(input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rdsdataservice.BatchExecuteStatementOutput), nil
	}
	out, err := c.RDSDataServiceAPI.BatchExecuteStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDSDataService) BatchExecuteStatementWithContext(ctx context.Context, input *rdsdataservice.BatchExecuteStatementInput, opts ...request.Option) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rdsdataservice.BatchExecuteStatementOutput), nil
	}
	out, err := c.RDSDataServiceAPI.BatchExecuteStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

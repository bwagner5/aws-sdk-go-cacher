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
package appconfigdatacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appconfigdata"
	"github.com/aws/aws-sdk-go/service/appconfigdata/appconfigdataiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppConfigData struct {
	appconfigdataiface.AppConfigDataAPI
	cache *cache.Cache
}

func New(appconfigdataapi appconfigdataiface.AppConfigDataAPI) *AppConfigData {
	return &AppConfigData{
		AppConfigDataAPI: appconfigdataapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppConfigData) GetLatestConfiguration(input *appconfigdata.GetLatestConfigurationInput) (*appconfigdata.GetLatestConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfigdata.GetLatestConfigurationOutput), nil
	}
	out, err := c.AppConfigDataAPI.GetLatestConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppConfigData) GetLatestConfigurationWithContext(ctx context.Context, input *appconfigdata.GetLatestConfigurationInput, opts ...request.Option) (*appconfigdata.GetLatestConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appconfigdata.GetLatestConfigurationOutput), nil
	}
	out, err := c.AppConfigDataAPI.GetLatestConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

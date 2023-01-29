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
package locationservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/locationservice"
	"github.com/aws/aws-sdk-go/service/locationservice/locationserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LocationService struct {
	locationserviceiface.LocationServiceAPI
	cache *cache.Cache
}

func New(locationserviceapi locationserviceiface.LocationServiceAPI) *LocationService {
	return &LocationService{
		LocationServiceAPI: locationserviceapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LocationService) BatchDeleteDevicePositionHistory(input *locationservice.BatchDeleteDevicePositionHistoryInput) (*locationservice.BatchDeleteDevicePositionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchDeleteDevicePositionHistoryOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchDeleteDevicePositionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchDeleteDevicePositionHistoryWithContext(ctx context.Context, input *locationservice.BatchDeleteDevicePositionHistoryInput, opts ...request.Option) (*locationservice.BatchDeleteDevicePositionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchDeleteDevicePositionHistoryOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchDeleteDevicePositionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchDeleteGeofence(input *locationservice.BatchDeleteGeofenceInput) (*locationservice.BatchDeleteGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchDeleteGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchDeleteGeofence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchDeleteGeofenceWithContext(ctx context.Context, input *locationservice.BatchDeleteGeofenceInput, opts ...request.Option) (*locationservice.BatchDeleteGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchDeleteGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchDeleteGeofenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchEvaluateGeofences(input *locationservice.BatchEvaluateGeofencesInput) (*locationservice.BatchEvaluateGeofencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchEvaluateGeofencesOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchEvaluateGeofences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchEvaluateGeofencesWithContext(ctx context.Context, input *locationservice.BatchEvaluateGeofencesInput, opts ...request.Option) (*locationservice.BatchEvaluateGeofencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchEvaluateGeofencesOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchEvaluateGeofencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchGetDevicePosition(input *locationservice.BatchGetDevicePositionInput) (*locationservice.BatchGetDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchGetDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchGetDevicePosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchGetDevicePositionWithContext(ctx context.Context, input *locationservice.BatchGetDevicePositionInput, opts ...request.Option) (*locationservice.BatchGetDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchGetDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchGetDevicePositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchPutGeofence(input *locationservice.BatchPutGeofenceInput) (*locationservice.BatchPutGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchPutGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchPutGeofence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchPutGeofenceWithContext(ctx context.Context, input *locationservice.BatchPutGeofenceInput, opts ...request.Option) (*locationservice.BatchPutGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchPutGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchPutGeofenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchUpdateDevicePosition(input *locationservice.BatchUpdateDevicePositionInput) (*locationservice.BatchUpdateDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchUpdateDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchUpdateDevicePosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) BatchUpdateDevicePositionWithContext(ctx context.Context, input *locationservice.BatchUpdateDevicePositionInput, opts ...request.Option) (*locationservice.BatchUpdateDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.BatchUpdateDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.BatchUpdateDevicePositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeGeofenceCollection(input *locationservice.DescribeGeofenceCollectionInput) (*locationservice.DescribeGeofenceCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeGeofenceCollectionOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeGeofenceCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeGeofenceCollectionWithContext(ctx context.Context, input *locationservice.DescribeGeofenceCollectionInput, opts ...request.Option) (*locationservice.DescribeGeofenceCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeGeofenceCollectionOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeGeofenceCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeMap(input *locationservice.DescribeMapInput) (*locationservice.DescribeMapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeMapOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeMap(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeMapWithContext(ctx context.Context, input *locationservice.DescribeMapInput, opts ...request.Option) (*locationservice.DescribeMapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeMapOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeMapWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribePlaceIndex(input *locationservice.DescribePlaceIndexInput) (*locationservice.DescribePlaceIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribePlaceIndexOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribePlaceIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribePlaceIndexWithContext(ctx context.Context, input *locationservice.DescribePlaceIndexInput, opts ...request.Option) (*locationservice.DescribePlaceIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribePlaceIndexOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribePlaceIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeRouteCalculator(input *locationservice.DescribeRouteCalculatorInput) (*locationservice.DescribeRouteCalculatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeRouteCalculatorOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeRouteCalculator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeRouteCalculatorWithContext(ctx context.Context, input *locationservice.DescribeRouteCalculatorInput, opts ...request.Option) (*locationservice.DescribeRouteCalculatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeRouteCalculatorOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeRouteCalculatorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeTracker(input *locationservice.DescribeTrackerInput) (*locationservice.DescribeTrackerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeTrackerOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeTracker(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) DescribeTrackerWithContext(ctx context.Context, input *locationservice.DescribeTrackerInput, opts ...request.Option) (*locationservice.DescribeTrackerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.DescribeTrackerOutput), nil
	}
	out, err := c.LocationServiceAPI.DescribeTrackerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetDevicePosition(input *locationservice.GetDevicePositionInput) (*locationservice.GetDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.GetDevicePosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetDevicePositionHistory(input *locationservice.GetDevicePositionHistoryInput) (*locationservice.GetDevicePositionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetDevicePositionHistoryOutput), nil
	}
	out, err := c.LocationServiceAPI.GetDevicePositionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetDevicePositionHistoryPages(input *locationservice.GetDevicePositionHistoryInput, fn func(*locationservice.GetDevicePositionHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.GetDevicePositionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.GetDevicePositionHistoryOutput{}
	fnCacher := func(out *locationservice.GetDevicePositionHistoryOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.GetDevicePositionHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) GetDevicePositionHistoryPagesWithContext(ctx context.Context, input *locationservice.GetDevicePositionHistoryInput, fn func(*locationservice.GetDevicePositionHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.GetDevicePositionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.GetDevicePositionHistoryOutput{}
	fnCacher := func(out *locationservice.GetDevicePositionHistoryOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.GetDevicePositionHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) GetDevicePositionHistoryWithContext(ctx context.Context, input *locationservice.GetDevicePositionHistoryInput, opts ...request.Option) (*locationservice.GetDevicePositionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetDevicePositionHistoryOutput), nil
	}
	out, err := c.LocationServiceAPI.GetDevicePositionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetDevicePositionWithContext(ctx context.Context, input *locationservice.GetDevicePositionInput, opts ...request.Option) (*locationservice.GetDevicePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetDevicePositionOutput), nil
	}
	out, err := c.LocationServiceAPI.GetDevicePositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetGeofence(input *locationservice.GetGeofenceInput) (*locationservice.GetGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.GetGeofence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetGeofenceWithContext(ctx context.Context, input *locationservice.GetGeofenceInput, opts ...request.Option) (*locationservice.GetGeofenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetGeofenceOutput), nil
	}
	out, err := c.LocationServiceAPI.GetGeofenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapGlyphs(input *locationservice.GetMapGlyphsInput) (*locationservice.GetMapGlyphsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapGlyphsOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapGlyphs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapGlyphsWithContext(ctx context.Context, input *locationservice.GetMapGlyphsInput, opts ...request.Option) (*locationservice.GetMapGlyphsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapGlyphsOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapGlyphsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapSprites(input *locationservice.GetMapSpritesInput) (*locationservice.GetMapSpritesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapSpritesOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapSprites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapSpritesWithContext(ctx context.Context, input *locationservice.GetMapSpritesInput, opts ...request.Option) (*locationservice.GetMapSpritesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapSpritesOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapSpritesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapStyleDescriptor(input *locationservice.GetMapStyleDescriptorInput) (*locationservice.GetMapStyleDescriptorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapStyleDescriptorOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapStyleDescriptor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapStyleDescriptorWithContext(ctx context.Context, input *locationservice.GetMapStyleDescriptorInput, opts ...request.Option) (*locationservice.GetMapStyleDescriptorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapStyleDescriptorOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapStyleDescriptorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapTile(input *locationservice.GetMapTileInput) (*locationservice.GetMapTileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapTileOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapTile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetMapTileWithContext(ctx context.Context, input *locationservice.GetMapTileInput, opts ...request.Option) (*locationservice.GetMapTileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetMapTileOutput), nil
	}
	out, err := c.LocationServiceAPI.GetMapTileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetPlace(input *locationservice.GetPlaceInput) (*locationservice.GetPlaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetPlaceOutput), nil
	}
	out, err := c.LocationServiceAPI.GetPlace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) GetPlaceWithContext(ctx context.Context, input *locationservice.GetPlaceInput, opts ...request.Option) (*locationservice.GetPlaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.GetPlaceOutput), nil
	}
	out, err := c.LocationServiceAPI.GetPlaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListDevicePositions(input *locationservice.ListDevicePositionsInput) (*locationservice.ListDevicePositionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListDevicePositionsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListDevicePositions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListDevicePositionsPages(input *locationservice.ListDevicePositionsInput, fn func(*locationservice.ListDevicePositionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListDevicePositionsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListDevicePositionsOutput{}
	fnCacher := func(out *locationservice.ListDevicePositionsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListDevicePositionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListDevicePositionsPagesWithContext(ctx context.Context, input *locationservice.ListDevicePositionsInput, fn func(*locationservice.ListDevicePositionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListDevicePositionsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListDevicePositionsOutput{}
	fnCacher := func(out *locationservice.ListDevicePositionsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListDevicePositionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListDevicePositionsWithContext(ctx context.Context, input *locationservice.ListDevicePositionsInput, opts ...request.Option) (*locationservice.ListDevicePositionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListDevicePositionsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListDevicePositionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListGeofenceCollections(input *locationservice.ListGeofenceCollectionsInput) (*locationservice.ListGeofenceCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListGeofenceCollectionsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListGeofenceCollections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListGeofenceCollectionsPages(input *locationservice.ListGeofenceCollectionsInput, fn func(*locationservice.ListGeofenceCollectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListGeofenceCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListGeofenceCollectionsOutput{}
	fnCacher := func(out *locationservice.ListGeofenceCollectionsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListGeofenceCollectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListGeofenceCollectionsPagesWithContext(ctx context.Context, input *locationservice.ListGeofenceCollectionsInput, fn func(*locationservice.ListGeofenceCollectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListGeofenceCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListGeofenceCollectionsOutput{}
	fnCacher := func(out *locationservice.ListGeofenceCollectionsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListGeofenceCollectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListGeofenceCollectionsWithContext(ctx context.Context, input *locationservice.ListGeofenceCollectionsInput, opts ...request.Option) (*locationservice.ListGeofenceCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListGeofenceCollectionsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListGeofenceCollectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListGeofences(input *locationservice.ListGeofencesInput) (*locationservice.ListGeofencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListGeofencesOutput), nil
	}
	out, err := c.LocationServiceAPI.ListGeofences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListGeofencesPages(input *locationservice.ListGeofencesInput, fn func(*locationservice.ListGeofencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListGeofencesOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListGeofencesOutput{}
	fnCacher := func(out *locationservice.ListGeofencesOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListGeofencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListGeofencesPagesWithContext(ctx context.Context, input *locationservice.ListGeofencesInput, fn func(*locationservice.ListGeofencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListGeofencesOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListGeofencesOutput{}
	fnCacher := func(out *locationservice.ListGeofencesOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListGeofencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListGeofencesWithContext(ctx context.Context, input *locationservice.ListGeofencesInput, opts ...request.Option) (*locationservice.ListGeofencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListGeofencesOutput), nil
	}
	out, err := c.LocationServiceAPI.ListGeofencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListMaps(input *locationservice.ListMapsInput) (*locationservice.ListMapsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListMapsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListMaps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListMapsPages(input *locationservice.ListMapsInput, fn func(*locationservice.ListMapsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListMapsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListMapsOutput{}
	fnCacher := func(out *locationservice.ListMapsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListMapsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListMapsPagesWithContext(ctx context.Context, input *locationservice.ListMapsInput, fn func(*locationservice.ListMapsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListMapsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListMapsOutput{}
	fnCacher := func(out *locationservice.ListMapsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListMapsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListMapsWithContext(ctx context.Context, input *locationservice.ListMapsInput, opts ...request.Option) (*locationservice.ListMapsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListMapsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListMapsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListPlaceIndexes(input *locationservice.ListPlaceIndexesInput) (*locationservice.ListPlaceIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListPlaceIndexesOutput), nil
	}
	out, err := c.LocationServiceAPI.ListPlaceIndexes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListPlaceIndexesPages(input *locationservice.ListPlaceIndexesInput, fn func(*locationservice.ListPlaceIndexesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListPlaceIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListPlaceIndexesOutput{}
	fnCacher := func(out *locationservice.ListPlaceIndexesOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListPlaceIndexesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListPlaceIndexesPagesWithContext(ctx context.Context, input *locationservice.ListPlaceIndexesInput, fn func(*locationservice.ListPlaceIndexesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListPlaceIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListPlaceIndexesOutput{}
	fnCacher := func(out *locationservice.ListPlaceIndexesOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListPlaceIndexesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListPlaceIndexesWithContext(ctx context.Context, input *locationservice.ListPlaceIndexesInput, opts ...request.Option) (*locationservice.ListPlaceIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListPlaceIndexesOutput), nil
	}
	out, err := c.LocationServiceAPI.ListPlaceIndexesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListRouteCalculators(input *locationservice.ListRouteCalculatorsInput) (*locationservice.ListRouteCalculatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListRouteCalculatorsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListRouteCalculators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListRouteCalculatorsPages(input *locationservice.ListRouteCalculatorsInput, fn func(*locationservice.ListRouteCalculatorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListRouteCalculatorsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListRouteCalculatorsOutput{}
	fnCacher := func(out *locationservice.ListRouteCalculatorsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListRouteCalculatorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListRouteCalculatorsPagesWithContext(ctx context.Context, input *locationservice.ListRouteCalculatorsInput, fn func(*locationservice.ListRouteCalculatorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListRouteCalculatorsOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListRouteCalculatorsOutput{}
	fnCacher := func(out *locationservice.ListRouteCalculatorsOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListRouteCalculatorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListRouteCalculatorsWithContext(ctx context.Context, input *locationservice.ListRouteCalculatorsInput, opts ...request.Option) (*locationservice.ListRouteCalculatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListRouteCalculatorsOutput), nil
	}
	out, err := c.LocationServiceAPI.ListRouteCalculatorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTagsForResource(input *locationservice.ListTagsForResourceInput) (*locationservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTagsForResourceOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTagsForResourceWithContext(ctx context.Context, input *locationservice.ListTagsForResourceInput, opts ...request.Option) (*locationservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTagsForResourceOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTrackerConsumers(input *locationservice.ListTrackerConsumersInput) (*locationservice.ListTrackerConsumersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTrackerConsumersOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTrackerConsumers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTrackerConsumersPages(input *locationservice.ListTrackerConsumersInput, fn func(*locationservice.ListTrackerConsumersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListTrackerConsumersOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListTrackerConsumersOutput{}
	fnCacher := func(out *locationservice.ListTrackerConsumersOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListTrackerConsumersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListTrackerConsumersPagesWithContext(ctx context.Context, input *locationservice.ListTrackerConsumersInput, fn func(*locationservice.ListTrackerConsumersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListTrackerConsumersOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListTrackerConsumersOutput{}
	fnCacher := func(out *locationservice.ListTrackerConsumersOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListTrackerConsumersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListTrackerConsumersWithContext(ctx context.Context, input *locationservice.ListTrackerConsumersInput, opts ...request.Option) (*locationservice.ListTrackerConsumersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTrackerConsumersOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTrackerConsumersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTrackers(input *locationservice.ListTrackersInput) (*locationservice.ListTrackersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTrackersOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTrackers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) ListTrackersPages(input *locationservice.ListTrackersInput, fn func(*locationservice.ListTrackersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListTrackersOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListTrackersOutput{}
	fnCacher := func(out *locationservice.ListTrackersOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListTrackersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListTrackersPagesWithContext(ctx context.Context, input *locationservice.ListTrackersInput, fn func(*locationservice.ListTrackersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*locationservice.ListTrackersOutput), false)
		return nil
	}
	cachable := true
	output := &locationservice.ListTrackersOutput{}
	fnCacher := func(out *locationservice.ListTrackersOutput, more bool) bool {
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
	if err := c.LocationServiceAPI.ListTrackersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LocationService) ListTrackersWithContext(ctx context.Context, input *locationservice.ListTrackersInput, opts ...request.Option) (*locationservice.ListTrackersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.ListTrackersOutput), nil
	}
	out, err := c.LocationServiceAPI.ListTrackersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForPosition(input *locationservice.SearchPlaceIndexForPositionInput) (*locationservice.SearchPlaceIndexForPositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForPositionOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForPosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForPositionWithContext(ctx context.Context, input *locationservice.SearchPlaceIndexForPositionInput, opts ...request.Option) (*locationservice.SearchPlaceIndexForPositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForPositionOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForPositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForSuggestions(input *locationservice.SearchPlaceIndexForSuggestionsInput) (*locationservice.SearchPlaceIndexForSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForSuggestionsOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForSuggestions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForSuggestionsWithContext(ctx context.Context, input *locationservice.SearchPlaceIndexForSuggestionsInput, opts ...request.Option) (*locationservice.SearchPlaceIndexForSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForSuggestionsOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForSuggestionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForText(input *locationservice.SearchPlaceIndexForTextInput) (*locationservice.SearchPlaceIndexForTextOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForTextOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForText(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LocationService) SearchPlaceIndexForTextWithContext(ctx context.Context, input *locationservice.SearchPlaceIndexForTextInput, opts ...request.Option) (*locationservice.SearchPlaceIndexForTextOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*locationservice.SearchPlaceIndexForTextOutput), nil
	}
	out, err := c.LocationServiceAPI.SearchPlaceIndexForTextWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package ioteventsdatacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTEventsData struct {
	ioteventsdataiface.IoTEventsDataAPI
	cache *cache.Cache
}

func New(ioteventsdataapi ioteventsdataiface.IoTEventsDataAPI) *IoTEventsData {
	return &IoTEventsData{
		IoTEventsDataAPI: ioteventsdataapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTEventsData) BatchAcknowledgeAlarm(input *ioteventsdata.BatchAcknowledgeAlarmInput) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchAcknowledgeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchAcknowledgeAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchAcknowledgeAlarmWithContext(ctx context.Context, input *ioteventsdata.BatchAcknowledgeAlarmInput, opts ...request.Option) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchAcknowledgeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchAcknowledgeAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchDeleteDetector(input *ioteventsdata.BatchDeleteDetectorInput) (*ioteventsdata.BatchDeleteDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchDeleteDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchDeleteDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchDeleteDetectorWithContext(ctx context.Context, input *ioteventsdata.BatchDeleteDetectorInput, opts ...request.Option) (*ioteventsdata.BatchDeleteDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchDeleteDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchDeleteDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchDisableAlarm(input *ioteventsdata.BatchDisableAlarmInput) (*ioteventsdata.BatchDisableAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchDisableAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchDisableAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchDisableAlarmWithContext(ctx context.Context, input *ioteventsdata.BatchDisableAlarmInput, opts ...request.Option) (*ioteventsdata.BatchDisableAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchDisableAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchDisableAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchEnableAlarm(input *ioteventsdata.BatchEnableAlarmInput) (*ioteventsdata.BatchEnableAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchEnableAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchEnableAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchEnableAlarmWithContext(ctx context.Context, input *ioteventsdata.BatchEnableAlarmInput, opts ...request.Option) (*ioteventsdata.BatchEnableAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchEnableAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchEnableAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchPutMessage(input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchPutMessageOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchPutMessage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchPutMessageWithContext(ctx context.Context, input *ioteventsdata.BatchPutMessageInput, opts ...request.Option) (*ioteventsdata.BatchPutMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchPutMessageOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchPutMessageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchResetAlarm(input *ioteventsdata.BatchResetAlarmInput) (*ioteventsdata.BatchResetAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchResetAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchResetAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchResetAlarmWithContext(ctx context.Context, input *ioteventsdata.BatchResetAlarmInput, opts ...request.Option) (*ioteventsdata.BatchResetAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchResetAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchResetAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchSnoozeAlarm(input *ioteventsdata.BatchSnoozeAlarmInput) (*ioteventsdata.BatchSnoozeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchSnoozeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchSnoozeAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchSnoozeAlarmWithContext(ctx context.Context, input *ioteventsdata.BatchSnoozeAlarmInput, opts ...request.Option) (*ioteventsdata.BatchSnoozeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchSnoozeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchSnoozeAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchUpdateDetector(input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchUpdateDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchUpdateDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) BatchUpdateDetectorWithContext(ctx context.Context, input *ioteventsdata.BatchUpdateDetectorInput, opts ...request.Option) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.BatchUpdateDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.BatchUpdateDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) DescribeAlarm(input *ioteventsdata.DescribeAlarmInput) (*ioteventsdata.DescribeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.DescribeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.DescribeAlarm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) DescribeAlarmWithContext(ctx context.Context, input *ioteventsdata.DescribeAlarmInput, opts ...request.Option) (*ioteventsdata.DescribeAlarmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.DescribeAlarmOutput), nil
	}
	out, err := c.IoTEventsDataAPI.DescribeAlarmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) DescribeDetector(input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.DescribeDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.DescribeDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) DescribeDetectorWithContext(ctx context.Context, input *ioteventsdata.DescribeDetectorInput, opts ...request.Option) (*ioteventsdata.DescribeDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.DescribeDetectorOutput), nil
	}
	out, err := c.IoTEventsDataAPI.DescribeDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) ListAlarms(input *ioteventsdata.ListAlarmsInput) (*ioteventsdata.ListAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.ListAlarmsOutput), nil
	}
	out, err := c.IoTEventsDataAPI.ListAlarms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) ListAlarmsWithContext(ctx context.Context, input *ioteventsdata.ListAlarmsInput, opts ...request.Option) (*ioteventsdata.ListAlarmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.ListAlarmsOutput), nil
	}
	out, err := c.IoTEventsDataAPI.ListAlarmsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) ListDetectors(input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.ListDetectorsOutput), nil
	}
	out, err := c.IoTEventsDataAPI.ListDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTEventsData) ListDetectorsWithContext(ctx context.Context, input *ioteventsdata.ListDetectorsInput, opts ...request.Option) (*ioteventsdata.ListDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ioteventsdata.ListDetectorsOutput), nil
	}
	out, err := c.IoTEventsDataAPI.ListDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

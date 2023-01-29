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
package iotcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot/iotiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoT struct {
	iotiface.IoTAPI
	cache *cache.Cache
}

func New(iotapi iotiface.IoTAPI) *IoT {
	return &IoT{
		IoTAPI: iotapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoT) DescribeAccountAuditConfiguration(input *iot.DescribeAccountAuditConfigurationInput) (*iot.DescribeAccountAuditConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAccountAuditConfigurationOutput), nil
	}
	out, err := c.IoTAPI.DescribeAccountAuditConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAccountAuditConfigurationWithContext(ctx context.Context, input *iot.DescribeAccountAuditConfigurationInput, opts ...request.Option) (*iot.DescribeAccountAuditConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAccountAuditConfigurationOutput), nil
	}
	out, err := c.IoTAPI.DescribeAccountAuditConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditFinding(input *iot.DescribeAuditFindingInput) (*iot.DescribeAuditFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditFindingOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditFinding(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditFindingWithContext(ctx context.Context, input *iot.DescribeAuditFindingInput, opts ...request.Option) (*iot.DescribeAuditFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditFindingOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditFindingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditMitigationActionsTask(input *iot.DescribeAuditMitigationActionsTaskInput) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditMitigationActionsTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditMitigationActionsTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditMitigationActionsTaskWithContext(ctx context.Context, input *iot.DescribeAuditMitigationActionsTaskInput, opts ...request.Option) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditMitigationActionsTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditMitigationActionsTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditSuppression(input *iot.DescribeAuditSuppressionInput) (*iot.DescribeAuditSuppressionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditSuppressionOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditSuppression(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditSuppressionWithContext(ctx context.Context, input *iot.DescribeAuditSuppressionInput, opts ...request.Option) (*iot.DescribeAuditSuppressionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditSuppressionOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditSuppressionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditTask(input *iot.DescribeAuditTaskInput) (*iot.DescribeAuditTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuditTaskWithContext(ctx context.Context, input *iot.DescribeAuditTaskInput, opts ...request.Option) (*iot.DescribeAuditTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuditTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuditTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuthorizer(input *iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuthorizerOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuthorizer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeAuthorizerWithContext(ctx context.Context, input *iot.DescribeAuthorizerInput, opts ...request.Option) (*iot.DescribeAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeAuthorizerOutput), nil
	}
	out, err := c.IoTAPI.DescribeAuthorizerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeBillingGroup(input *iot.DescribeBillingGroupInput) (*iot.DescribeBillingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeBillingGroupOutput), nil
	}
	out, err := c.IoTAPI.DescribeBillingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeBillingGroupWithContext(ctx context.Context, input *iot.DescribeBillingGroupInput, opts ...request.Option) (*iot.DescribeBillingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeBillingGroupOutput), nil
	}
	out, err := c.IoTAPI.DescribeBillingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCACertificate(input *iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCACertificateOutput), nil
	}
	out, err := c.IoTAPI.DescribeCACertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCACertificateWithContext(ctx context.Context, input *iot.DescribeCACertificateInput, opts ...request.Option) (*iot.DescribeCACertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCACertificateOutput), nil
	}
	out, err := c.IoTAPI.DescribeCACertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCertificate(input *iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCertificateOutput), nil
	}
	out, err := c.IoTAPI.DescribeCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCertificateWithContext(ctx context.Context, input *iot.DescribeCertificateInput, opts ...request.Option) (*iot.DescribeCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCertificateOutput), nil
	}
	out, err := c.IoTAPI.DescribeCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCustomMetric(input *iot.DescribeCustomMetricInput) (*iot.DescribeCustomMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCustomMetricOutput), nil
	}
	out, err := c.IoTAPI.DescribeCustomMetric(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeCustomMetricWithContext(ctx context.Context, input *iot.DescribeCustomMetricInput, opts ...request.Option) (*iot.DescribeCustomMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeCustomMetricOutput), nil
	}
	out, err := c.IoTAPI.DescribeCustomMetricWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDefaultAuthorizer(input *iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDefaultAuthorizerOutput), nil
	}
	out, err := c.IoTAPI.DescribeDefaultAuthorizer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDefaultAuthorizerWithContext(ctx context.Context, input *iot.DescribeDefaultAuthorizerInput, opts ...request.Option) (*iot.DescribeDefaultAuthorizerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDefaultAuthorizerOutput), nil
	}
	out, err := c.IoTAPI.DescribeDefaultAuthorizerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDetectMitigationActionsTask(input *iot.DescribeDetectMitigationActionsTaskInput) (*iot.DescribeDetectMitigationActionsTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDetectMitigationActionsTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeDetectMitigationActionsTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDetectMitigationActionsTaskWithContext(ctx context.Context, input *iot.DescribeDetectMitigationActionsTaskInput, opts ...request.Option) (*iot.DescribeDetectMitigationActionsTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDetectMitigationActionsTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeDetectMitigationActionsTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDimension(input *iot.DescribeDimensionInput) (*iot.DescribeDimensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDimensionOutput), nil
	}
	out, err := c.IoTAPI.DescribeDimension(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDimensionWithContext(ctx context.Context, input *iot.DescribeDimensionInput, opts ...request.Option) (*iot.DescribeDimensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDimensionOutput), nil
	}
	out, err := c.IoTAPI.DescribeDimensionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDomainConfiguration(input *iot.DescribeDomainConfigurationInput) (*iot.DescribeDomainConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDomainConfigurationOutput), nil
	}
	out, err := c.IoTAPI.DescribeDomainConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeDomainConfigurationWithContext(ctx context.Context, input *iot.DescribeDomainConfigurationInput, opts ...request.Option) (*iot.DescribeDomainConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeDomainConfigurationOutput), nil
	}
	out, err := c.IoTAPI.DescribeDomainConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeEndpoint(input *iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeEndpointOutput), nil
	}
	out, err := c.IoTAPI.DescribeEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeEndpointWithContext(ctx context.Context, input *iot.DescribeEndpointInput, opts ...request.Option) (*iot.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeEndpointOutput), nil
	}
	out, err := c.IoTAPI.DescribeEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeEventConfigurations(input *iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeEventConfigurationsOutput), nil
	}
	out, err := c.IoTAPI.DescribeEventConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeEventConfigurationsWithContext(ctx context.Context, input *iot.DescribeEventConfigurationsInput, opts ...request.Option) (*iot.DescribeEventConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeEventConfigurationsOutput), nil
	}
	out, err := c.IoTAPI.DescribeEventConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeFleetMetric(input *iot.DescribeFleetMetricInput) (*iot.DescribeFleetMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeFleetMetricOutput), nil
	}
	out, err := c.IoTAPI.DescribeFleetMetric(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeFleetMetricWithContext(ctx context.Context, input *iot.DescribeFleetMetricInput, opts ...request.Option) (*iot.DescribeFleetMetricOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeFleetMetricOutput), nil
	}
	out, err := c.IoTAPI.DescribeFleetMetricWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeIndex(input *iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeIndexOutput), nil
	}
	out, err := c.IoTAPI.DescribeIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeIndexWithContext(ctx context.Context, input *iot.DescribeIndexInput, opts ...request.Option) (*iot.DescribeIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeIndexOutput), nil
	}
	out, err := c.IoTAPI.DescribeIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJob(input *iot.DescribeJobInput) (*iot.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobOutput), nil
	}
	out, err := c.IoTAPI.DescribeJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJobExecution(input *iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobExecutionOutput), nil
	}
	out, err := c.IoTAPI.DescribeJobExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJobExecutionWithContext(ctx context.Context, input *iot.DescribeJobExecutionInput, opts ...request.Option) (*iot.DescribeJobExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobExecutionOutput), nil
	}
	out, err := c.IoTAPI.DescribeJobExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJobTemplate(input *iot.DescribeJobTemplateInput) (*iot.DescribeJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeJobTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJobTemplateWithContext(ctx context.Context, input *iot.DescribeJobTemplateInput, opts ...request.Option) (*iot.DescribeJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeJobTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeJobWithContext(ctx context.Context, input *iot.DescribeJobInput, opts ...request.Option) (*iot.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeJobOutput), nil
	}
	out, err := c.IoTAPI.DescribeJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeManagedJobTemplate(input *iot.DescribeManagedJobTemplateInput) (*iot.DescribeManagedJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeManagedJobTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeManagedJobTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeManagedJobTemplateWithContext(ctx context.Context, input *iot.DescribeManagedJobTemplateInput, opts ...request.Option) (*iot.DescribeManagedJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeManagedJobTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeManagedJobTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeMitigationAction(input *iot.DescribeMitigationActionInput) (*iot.DescribeMitigationActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeMitigationActionOutput), nil
	}
	out, err := c.IoTAPI.DescribeMitigationAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeMitigationActionWithContext(ctx context.Context, input *iot.DescribeMitigationActionInput, opts ...request.Option) (*iot.DescribeMitigationActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeMitigationActionOutput), nil
	}
	out, err := c.IoTAPI.DescribeMitigationActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeProvisioningTemplate(input *iot.DescribeProvisioningTemplateInput) (*iot.DescribeProvisioningTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeProvisioningTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeProvisioningTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeProvisioningTemplateVersion(input *iot.DescribeProvisioningTemplateVersionInput) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeProvisioningTemplateVersionOutput), nil
	}
	out, err := c.IoTAPI.DescribeProvisioningTemplateVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeProvisioningTemplateVersionWithContext(ctx context.Context, input *iot.DescribeProvisioningTemplateVersionInput, opts ...request.Option) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeProvisioningTemplateVersionOutput), nil
	}
	out, err := c.IoTAPI.DescribeProvisioningTemplateVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeProvisioningTemplateWithContext(ctx context.Context, input *iot.DescribeProvisioningTemplateInput, opts ...request.Option) (*iot.DescribeProvisioningTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeProvisioningTemplateOutput), nil
	}
	out, err := c.IoTAPI.DescribeProvisioningTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeRoleAlias(input *iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeRoleAliasOutput), nil
	}
	out, err := c.IoTAPI.DescribeRoleAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeRoleAliasWithContext(ctx context.Context, input *iot.DescribeRoleAliasInput, opts ...request.Option) (*iot.DescribeRoleAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeRoleAliasOutput), nil
	}
	out, err := c.IoTAPI.DescribeRoleAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeScheduledAudit(input *iot.DescribeScheduledAuditInput) (*iot.DescribeScheduledAuditOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeScheduledAuditOutput), nil
	}
	out, err := c.IoTAPI.DescribeScheduledAudit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeScheduledAuditWithContext(ctx context.Context, input *iot.DescribeScheduledAuditInput, opts ...request.Option) (*iot.DescribeScheduledAuditOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeScheduledAuditOutput), nil
	}
	out, err := c.IoTAPI.DescribeScheduledAuditWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeSecurityProfile(input *iot.DescribeSecurityProfileInput) (*iot.DescribeSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeSecurityProfileOutput), nil
	}
	out, err := c.IoTAPI.DescribeSecurityProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeSecurityProfileWithContext(ctx context.Context, input *iot.DescribeSecurityProfileInput, opts ...request.Option) (*iot.DescribeSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeSecurityProfileOutput), nil
	}
	out, err := c.IoTAPI.DescribeSecurityProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeStream(input *iot.DescribeStreamInput) (*iot.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeStreamOutput), nil
	}
	out, err := c.IoTAPI.DescribeStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeStreamWithContext(ctx context.Context, input *iot.DescribeStreamInput, opts ...request.Option) (*iot.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeStreamOutput), nil
	}
	out, err := c.IoTAPI.DescribeStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThing(input *iot.DescribeThingInput) (*iot.DescribeThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingOutput), nil
	}
	out, err := c.IoTAPI.DescribeThing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingGroup(input *iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingGroupOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingGroupWithContext(ctx context.Context, input *iot.DescribeThingGroupInput, opts ...request.Option) (*iot.DescribeThingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingGroupOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingRegistrationTask(input *iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingRegistrationTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingRegistrationTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingRegistrationTaskWithContext(ctx context.Context, input *iot.DescribeThingRegistrationTaskInput, opts ...request.Option) (*iot.DescribeThingRegistrationTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingRegistrationTaskOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingRegistrationTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingType(input *iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingTypeOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingTypeWithContext(ctx context.Context, input *iot.DescribeThingTypeInput, opts ...request.Option) (*iot.DescribeThingTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingTypeOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) DescribeThingWithContext(ctx context.Context, input *iot.DescribeThingInput, opts ...request.Option) (*iot.DescribeThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.DescribeThingOutput), nil
	}
	out, err := c.IoTAPI.DescribeThingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetBehaviorModelTrainingSummaries(input *iot.GetBehaviorModelTrainingSummariesInput) (*iot.GetBehaviorModelTrainingSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetBehaviorModelTrainingSummariesOutput), nil
	}
	out, err := c.IoTAPI.GetBehaviorModelTrainingSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetBehaviorModelTrainingSummariesPages(input *iot.GetBehaviorModelTrainingSummariesInput, fn func(*iot.GetBehaviorModelTrainingSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.GetBehaviorModelTrainingSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.GetBehaviorModelTrainingSummariesOutput{}
	fnCacher := func(out *iot.GetBehaviorModelTrainingSummariesOutput, more bool) bool {
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
	if err := c.IoTAPI.GetBehaviorModelTrainingSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) GetBehaviorModelTrainingSummariesPagesWithContext(ctx context.Context, input *iot.GetBehaviorModelTrainingSummariesInput, fn func(*iot.GetBehaviorModelTrainingSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.GetBehaviorModelTrainingSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.GetBehaviorModelTrainingSummariesOutput{}
	fnCacher := func(out *iot.GetBehaviorModelTrainingSummariesOutput, more bool) bool {
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
	if err := c.IoTAPI.GetBehaviorModelTrainingSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) GetBehaviorModelTrainingSummariesWithContext(ctx context.Context, input *iot.GetBehaviorModelTrainingSummariesInput, opts ...request.Option) (*iot.GetBehaviorModelTrainingSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetBehaviorModelTrainingSummariesOutput), nil
	}
	out, err := c.IoTAPI.GetBehaviorModelTrainingSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetBucketsAggregation(input *iot.GetBucketsAggregationInput) (*iot.GetBucketsAggregationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetBucketsAggregationOutput), nil
	}
	out, err := c.IoTAPI.GetBucketsAggregation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetBucketsAggregationWithContext(ctx context.Context, input *iot.GetBucketsAggregationInput, opts ...request.Option) (*iot.GetBucketsAggregationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetBucketsAggregationOutput), nil
	}
	out, err := c.IoTAPI.GetBucketsAggregationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetCardinality(input *iot.GetCardinalityInput) (*iot.GetCardinalityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetCardinalityOutput), nil
	}
	out, err := c.IoTAPI.GetCardinality(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetCardinalityWithContext(ctx context.Context, input *iot.GetCardinalityInput, opts ...request.Option) (*iot.GetCardinalityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetCardinalityOutput), nil
	}
	out, err := c.IoTAPI.GetCardinalityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetEffectivePolicies(input *iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetEffectivePoliciesOutput), nil
	}
	out, err := c.IoTAPI.GetEffectivePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetEffectivePoliciesWithContext(ctx context.Context, input *iot.GetEffectivePoliciesInput, opts ...request.Option) (*iot.GetEffectivePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetEffectivePoliciesOutput), nil
	}
	out, err := c.IoTAPI.GetEffectivePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetIndexingConfiguration(input *iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetIndexingConfigurationOutput), nil
	}
	out, err := c.IoTAPI.GetIndexingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetIndexingConfigurationWithContext(ctx context.Context, input *iot.GetIndexingConfigurationInput, opts ...request.Option) (*iot.GetIndexingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetIndexingConfigurationOutput), nil
	}
	out, err := c.IoTAPI.GetIndexingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetJobDocument(input *iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetJobDocumentOutput), nil
	}
	out, err := c.IoTAPI.GetJobDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetJobDocumentWithContext(ctx context.Context, input *iot.GetJobDocumentInput, opts ...request.Option) (*iot.GetJobDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetJobDocumentOutput), nil
	}
	out, err := c.IoTAPI.GetJobDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetLoggingOptions(input *iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetLoggingOptionsOutput), nil
	}
	out, err := c.IoTAPI.GetLoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetLoggingOptionsWithContext(ctx context.Context, input *iot.GetLoggingOptionsInput, opts ...request.Option) (*iot.GetLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetLoggingOptionsOutput), nil
	}
	out, err := c.IoTAPI.GetLoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetOTAUpdate(input *iot.GetOTAUpdateInput) (*iot.GetOTAUpdateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetOTAUpdateOutput), nil
	}
	out, err := c.IoTAPI.GetOTAUpdate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetOTAUpdateWithContext(ctx context.Context, input *iot.GetOTAUpdateInput, opts ...request.Option) (*iot.GetOTAUpdateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetOTAUpdateOutput), nil
	}
	out, err := c.IoTAPI.GetOTAUpdateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPercentiles(input *iot.GetPercentilesInput) (*iot.GetPercentilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPercentilesOutput), nil
	}
	out, err := c.IoTAPI.GetPercentiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPercentilesWithContext(ctx context.Context, input *iot.GetPercentilesInput, opts ...request.Option) (*iot.GetPercentilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPercentilesOutput), nil
	}
	out, err := c.IoTAPI.GetPercentilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPolicy(input *iot.GetPolicyInput) (*iot.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPolicyOutput), nil
	}
	out, err := c.IoTAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPolicyVersion(input *iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPolicyVersionOutput), nil
	}
	out, err := c.IoTAPI.GetPolicyVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPolicyVersionWithContext(ctx context.Context, input *iot.GetPolicyVersionInput, opts ...request.Option) (*iot.GetPolicyVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPolicyVersionOutput), nil
	}
	out, err := c.IoTAPI.GetPolicyVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetPolicyWithContext(ctx context.Context, input *iot.GetPolicyInput, opts ...request.Option) (*iot.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetPolicyOutput), nil
	}
	out, err := c.IoTAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetRegistrationCode(input *iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetRegistrationCodeOutput), nil
	}
	out, err := c.IoTAPI.GetRegistrationCode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetRegistrationCodeWithContext(ctx context.Context, input *iot.GetRegistrationCodeInput, opts ...request.Option) (*iot.GetRegistrationCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetRegistrationCodeOutput), nil
	}
	out, err := c.IoTAPI.GetRegistrationCodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetStatistics(input *iot.GetStatisticsInput) (*iot.GetStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetStatisticsOutput), nil
	}
	out, err := c.IoTAPI.GetStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetStatisticsWithContext(ctx context.Context, input *iot.GetStatisticsInput, opts ...request.Option) (*iot.GetStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetStatisticsOutput), nil
	}
	out, err := c.IoTAPI.GetStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetTopicRule(input *iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetTopicRuleOutput), nil
	}
	out, err := c.IoTAPI.GetTopicRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetTopicRuleDestination(input *iot.GetTopicRuleDestinationInput) (*iot.GetTopicRuleDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetTopicRuleDestinationOutput), nil
	}
	out, err := c.IoTAPI.GetTopicRuleDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetTopicRuleDestinationWithContext(ctx context.Context, input *iot.GetTopicRuleDestinationInput, opts ...request.Option) (*iot.GetTopicRuleDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetTopicRuleDestinationOutput), nil
	}
	out, err := c.IoTAPI.GetTopicRuleDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetTopicRuleWithContext(ctx context.Context, input *iot.GetTopicRuleInput, opts ...request.Option) (*iot.GetTopicRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetTopicRuleOutput), nil
	}
	out, err := c.IoTAPI.GetTopicRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetV2LoggingOptions(input *iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetV2LoggingOptionsOutput), nil
	}
	out, err := c.IoTAPI.GetV2LoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) GetV2LoggingOptionsWithContext(ctx context.Context, input *iot.GetV2LoggingOptionsInput, opts ...request.Option) (*iot.GetV2LoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.GetV2LoggingOptionsOutput), nil
	}
	out, err := c.IoTAPI.GetV2LoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListActiveViolations(input *iot.ListActiveViolationsInput) (*iot.ListActiveViolationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListActiveViolationsOutput), nil
	}
	out, err := c.IoTAPI.ListActiveViolations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListActiveViolationsPages(input *iot.ListActiveViolationsInput, fn func(*iot.ListActiveViolationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListActiveViolationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListActiveViolationsOutput{}
	fnCacher := func(out *iot.ListActiveViolationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListActiveViolationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListActiveViolationsPagesWithContext(ctx context.Context, input *iot.ListActiveViolationsInput, fn func(*iot.ListActiveViolationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListActiveViolationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListActiveViolationsOutput{}
	fnCacher := func(out *iot.ListActiveViolationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListActiveViolationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListActiveViolationsWithContext(ctx context.Context, input *iot.ListActiveViolationsInput, opts ...request.Option) (*iot.ListActiveViolationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListActiveViolationsOutput), nil
	}
	out, err := c.IoTAPI.ListActiveViolationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAttachedPolicies(input *iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAttachedPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListAttachedPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAttachedPoliciesPages(input *iot.ListAttachedPoliciesInput, fn func(*iot.ListAttachedPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAttachedPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAttachedPoliciesOutput{}
	fnCacher := func(out *iot.ListAttachedPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAttachedPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAttachedPoliciesPagesWithContext(ctx context.Context, input *iot.ListAttachedPoliciesInput, fn func(*iot.ListAttachedPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAttachedPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAttachedPoliciesOutput{}
	fnCacher := func(out *iot.ListAttachedPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAttachedPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAttachedPoliciesWithContext(ctx context.Context, input *iot.ListAttachedPoliciesInput, opts ...request.Option) (*iot.ListAttachedPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAttachedPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListAttachedPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditFindings(input *iot.ListAuditFindingsInput) (*iot.ListAuditFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditFindingsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditFindingsPages(input *iot.ListAuditFindingsInput, fn func(*iot.ListAuditFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditFindingsOutput{}
	fnCacher := func(out *iot.ListAuditFindingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditFindingsPagesWithContext(ctx context.Context, input *iot.ListAuditFindingsInput, fn func(*iot.ListAuditFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditFindingsOutput{}
	fnCacher := func(out *iot.ListAuditFindingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditFindingsWithContext(ctx context.Context, input *iot.ListAuditFindingsInput, opts ...request.Option) (*iot.ListAuditFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditFindingsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditMitigationActionsExecutions(input *iot.ListAuditMitigationActionsExecutionsInput) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditMitigationActionsExecutionsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditMitigationActionsExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditMitigationActionsExecutionsPages(input *iot.ListAuditMitigationActionsExecutionsInput, fn func(*iot.ListAuditMitigationActionsExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditMitigationActionsExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditMitigationActionsExecutionsOutput{}
	fnCacher := func(out *iot.ListAuditMitigationActionsExecutionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditMitigationActionsExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditMitigationActionsExecutionsPagesWithContext(ctx context.Context, input *iot.ListAuditMitigationActionsExecutionsInput, fn func(*iot.ListAuditMitigationActionsExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditMitigationActionsExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditMitigationActionsExecutionsOutput{}
	fnCacher := func(out *iot.ListAuditMitigationActionsExecutionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditMitigationActionsExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditMitigationActionsExecutionsWithContext(ctx context.Context, input *iot.ListAuditMitigationActionsExecutionsInput, opts ...request.Option) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditMitigationActionsExecutionsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditMitigationActionsExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditMitigationActionsTasks(input *iot.ListAuditMitigationActionsTasksInput) (*iot.ListAuditMitigationActionsTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditMitigationActionsTasksOutput), nil
	}
	out, err := c.IoTAPI.ListAuditMitigationActionsTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditMitigationActionsTasksPages(input *iot.ListAuditMitigationActionsTasksInput, fn func(*iot.ListAuditMitigationActionsTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditMitigationActionsTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditMitigationActionsTasksOutput{}
	fnCacher := func(out *iot.ListAuditMitigationActionsTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditMitigationActionsTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditMitigationActionsTasksPagesWithContext(ctx context.Context, input *iot.ListAuditMitigationActionsTasksInput, fn func(*iot.ListAuditMitigationActionsTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditMitigationActionsTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditMitigationActionsTasksOutput{}
	fnCacher := func(out *iot.ListAuditMitigationActionsTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditMitigationActionsTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditMitigationActionsTasksWithContext(ctx context.Context, input *iot.ListAuditMitigationActionsTasksInput, opts ...request.Option) (*iot.ListAuditMitigationActionsTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditMitigationActionsTasksOutput), nil
	}
	out, err := c.IoTAPI.ListAuditMitigationActionsTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditSuppressions(input *iot.ListAuditSuppressionsInput) (*iot.ListAuditSuppressionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditSuppressionsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditSuppressions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditSuppressionsPages(input *iot.ListAuditSuppressionsInput, fn func(*iot.ListAuditSuppressionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditSuppressionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditSuppressionsOutput{}
	fnCacher := func(out *iot.ListAuditSuppressionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditSuppressionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditSuppressionsPagesWithContext(ctx context.Context, input *iot.ListAuditSuppressionsInput, fn func(*iot.ListAuditSuppressionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditSuppressionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditSuppressionsOutput{}
	fnCacher := func(out *iot.ListAuditSuppressionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditSuppressionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditSuppressionsWithContext(ctx context.Context, input *iot.ListAuditSuppressionsInput, opts ...request.Option) (*iot.ListAuditSuppressionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditSuppressionsOutput), nil
	}
	out, err := c.IoTAPI.ListAuditSuppressionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditTasks(input *iot.ListAuditTasksInput) (*iot.ListAuditTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditTasksOutput), nil
	}
	out, err := c.IoTAPI.ListAuditTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuditTasksPages(input *iot.ListAuditTasksInput, fn func(*iot.ListAuditTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditTasksOutput{}
	fnCacher := func(out *iot.ListAuditTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditTasksPagesWithContext(ctx context.Context, input *iot.ListAuditTasksInput, fn func(*iot.ListAuditTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuditTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuditTasksOutput{}
	fnCacher := func(out *iot.ListAuditTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuditTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuditTasksWithContext(ctx context.Context, input *iot.ListAuditTasksInput, opts ...request.Option) (*iot.ListAuditTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuditTasksOutput), nil
	}
	out, err := c.IoTAPI.ListAuditTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuthorizers(input *iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuthorizersOutput), nil
	}
	out, err := c.IoTAPI.ListAuthorizers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListAuthorizersPages(input *iot.ListAuthorizersInput, fn func(*iot.ListAuthorizersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuthorizersOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuthorizersOutput{}
	fnCacher := func(out *iot.ListAuthorizersOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuthorizersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuthorizersPagesWithContext(ctx context.Context, input *iot.ListAuthorizersInput, fn func(*iot.ListAuthorizersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListAuthorizersOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListAuthorizersOutput{}
	fnCacher := func(out *iot.ListAuthorizersOutput, more bool) bool {
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
	if err := c.IoTAPI.ListAuthorizersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListAuthorizersWithContext(ctx context.Context, input *iot.ListAuthorizersInput, opts ...request.Option) (*iot.ListAuthorizersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListAuthorizersOutput), nil
	}
	out, err := c.IoTAPI.ListAuthorizersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListBillingGroups(input *iot.ListBillingGroupsInput) (*iot.ListBillingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListBillingGroupsOutput), nil
	}
	out, err := c.IoTAPI.ListBillingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListBillingGroupsPages(input *iot.ListBillingGroupsInput, fn func(*iot.ListBillingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListBillingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListBillingGroupsOutput{}
	fnCacher := func(out *iot.ListBillingGroupsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListBillingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListBillingGroupsPagesWithContext(ctx context.Context, input *iot.ListBillingGroupsInput, fn func(*iot.ListBillingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListBillingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListBillingGroupsOutput{}
	fnCacher := func(out *iot.ListBillingGroupsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListBillingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListBillingGroupsWithContext(ctx context.Context, input *iot.ListBillingGroupsInput, opts ...request.Option) (*iot.ListBillingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListBillingGroupsOutput), nil
	}
	out, err := c.IoTAPI.ListBillingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCACertificates(input *iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCACertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListCACertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCACertificatesPages(input *iot.ListCACertificatesInput, fn func(*iot.ListCACertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCACertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCACertificatesOutput{}
	fnCacher := func(out *iot.ListCACertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCACertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCACertificatesPagesWithContext(ctx context.Context, input *iot.ListCACertificatesInput, fn func(*iot.ListCACertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCACertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCACertificatesOutput{}
	fnCacher := func(out *iot.ListCACertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCACertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCACertificatesWithContext(ctx context.Context, input *iot.ListCACertificatesInput, opts ...request.Option) (*iot.ListCACertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCACertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListCACertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCertificates(input *iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCertificatesByCA(input *iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCertificatesByCAOutput), nil
	}
	out, err := c.IoTAPI.ListCertificatesByCA(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCertificatesByCAPages(input *iot.ListCertificatesByCAInput, fn func(*iot.ListCertificatesByCAOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCertificatesByCAOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCertificatesByCAOutput{}
	fnCacher := func(out *iot.ListCertificatesByCAOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCertificatesByCAPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCertificatesByCAPagesWithContext(ctx context.Context, input *iot.ListCertificatesByCAInput, fn func(*iot.ListCertificatesByCAOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCertificatesByCAOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCertificatesByCAOutput{}
	fnCacher := func(out *iot.ListCertificatesByCAOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCertificatesByCAPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCertificatesByCAWithContext(ctx context.Context, input *iot.ListCertificatesByCAInput, opts ...request.Option) (*iot.ListCertificatesByCAOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCertificatesByCAOutput), nil
	}
	out, err := c.IoTAPI.ListCertificatesByCAWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCertificatesPages(input *iot.ListCertificatesInput, fn func(*iot.ListCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCertificatesOutput{}
	fnCacher := func(out *iot.ListCertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCertificatesPagesWithContext(ctx context.Context, input *iot.ListCertificatesInput, fn func(*iot.ListCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCertificatesOutput{}
	fnCacher := func(out *iot.ListCertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCertificatesWithContext(ctx context.Context, input *iot.ListCertificatesInput, opts ...request.Option) (*iot.ListCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCustomMetrics(input *iot.ListCustomMetricsInput) (*iot.ListCustomMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCustomMetricsOutput), nil
	}
	out, err := c.IoTAPI.ListCustomMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListCustomMetricsPages(input *iot.ListCustomMetricsInput, fn func(*iot.ListCustomMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCustomMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCustomMetricsOutput{}
	fnCacher := func(out *iot.ListCustomMetricsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCustomMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCustomMetricsPagesWithContext(ctx context.Context, input *iot.ListCustomMetricsInput, fn func(*iot.ListCustomMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListCustomMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListCustomMetricsOutput{}
	fnCacher := func(out *iot.ListCustomMetricsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListCustomMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListCustomMetricsWithContext(ctx context.Context, input *iot.ListCustomMetricsInput, opts ...request.Option) (*iot.ListCustomMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListCustomMetricsOutput), nil
	}
	out, err := c.IoTAPI.ListCustomMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDetectMitigationActionsExecutions(input *iot.ListDetectMitigationActionsExecutionsInput) (*iot.ListDetectMitigationActionsExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDetectMitigationActionsExecutionsOutput), nil
	}
	out, err := c.IoTAPI.ListDetectMitigationActionsExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDetectMitigationActionsExecutionsPages(input *iot.ListDetectMitigationActionsExecutionsInput, fn func(*iot.ListDetectMitigationActionsExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDetectMitigationActionsExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDetectMitigationActionsExecutionsOutput{}
	fnCacher := func(out *iot.ListDetectMitigationActionsExecutionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDetectMitigationActionsExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDetectMitigationActionsExecutionsPagesWithContext(ctx context.Context, input *iot.ListDetectMitigationActionsExecutionsInput, fn func(*iot.ListDetectMitigationActionsExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDetectMitigationActionsExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDetectMitigationActionsExecutionsOutput{}
	fnCacher := func(out *iot.ListDetectMitigationActionsExecutionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDetectMitigationActionsExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDetectMitigationActionsExecutionsWithContext(ctx context.Context, input *iot.ListDetectMitigationActionsExecutionsInput, opts ...request.Option) (*iot.ListDetectMitigationActionsExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDetectMitigationActionsExecutionsOutput), nil
	}
	out, err := c.IoTAPI.ListDetectMitigationActionsExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDetectMitigationActionsTasks(input *iot.ListDetectMitigationActionsTasksInput) (*iot.ListDetectMitigationActionsTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDetectMitigationActionsTasksOutput), nil
	}
	out, err := c.IoTAPI.ListDetectMitigationActionsTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDetectMitigationActionsTasksPages(input *iot.ListDetectMitigationActionsTasksInput, fn func(*iot.ListDetectMitigationActionsTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDetectMitigationActionsTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDetectMitigationActionsTasksOutput{}
	fnCacher := func(out *iot.ListDetectMitigationActionsTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDetectMitigationActionsTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDetectMitigationActionsTasksPagesWithContext(ctx context.Context, input *iot.ListDetectMitigationActionsTasksInput, fn func(*iot.ListDetectMitigationActionsTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDetectMitigationActionsTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDetectMitigationActionsTasksOutput{}
	fnCacher := func(out *iot.ListDetectMitigationActionsTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDetectMitigationActionsTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDetectMitigationActionsTasksWithContext(ctx context.Context, input *iot.ListDetectMitigationActionsTasksInput, opts ...request.Option) (*iot.ListDetectMitigationActionsTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDetectMitigationActionsTasksOutput), nil
	}
	out, err := c.IoTAPI.ListDetectMitigationActionsTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDimensions(input *iot.ListDimensionsInput) (*iot.ListDimensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDimensionsOutput), nil
	}
	out, err := c.IoTAPI.ListDimensions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDimensionsPages(input *iot.ListDimensionsInput, fn func(*iot.ListDimensionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDimensionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDimensionsOutput{}
	fnCacher := func(out *iot.ListDimensionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDimensionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDimensionsPagesWithContext(ctx context.Context, input *iot.ListDimensionsInput, fn func(*iot.ListDimensionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDimensionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDimensionsOutput{}
	fnCacher := func(out *iot.ListDimensionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDimensionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDimensionsWithContext(ctx context.Context, input *iot.ListDimensionsInput, opts ...request.Option) (*iot.ListDimensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDimensionsOutput), nil
	}
	out, err := c.IoTAPI.ListDimensionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDomainConfigurations(input *iot.ListDomainConfigurationsInput) (*iot.ListDomainConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDomainConfigurationsOutput), nil
	}
	out, err := c.IoTAPI.ListDomainConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListDomainConfigurationsPages(input *iot.ListDomainConfigurationsInput, fn func(*iot.ListDomainConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDomainConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDomainConfigurationsOutput{}
	fnCacher := func(out *iot.ListDomainConfigurationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDomainConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDomainConfigurationsPagesWithContext(ctx context.Context, input *iot.ListDomainConfigurationsInput, fn func(*iot.ListDomainConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListDomainConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListDomainConfigurationsOutput{}
	fnCacher := func(out *iot.ListDomainConfigurationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListDomainConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListDomainConfigurationsWithContext(ctx context.Context, input *iot.ListDomainConfigurationsInput, opts ...request.Option) (*iot.ListDomainConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListDomainConfigurationsOutput), nil
	}
	out, err := c.IoTAPI.ListDomainConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListFleetMetrics(input *iot.ListFleetMetricsInput) (*iot.ListFleetMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListFleetMetricsOutput), nil
	}
	out, err := c.IoTAPI.ListFleetMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListFleetMetricsPages(input *iot.ListFleetMetricsInput, fn func(*iot.ListFleetMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListFleetMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListFleetMetricsOutput{}
	fnCacher := func(out *iot.ListFleetMetricsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListFleetMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListFleetMetricsPagesWithContext(ctx context.Context, input *iot.ListFleetMetricsInput, fn func(*iot.ListFleetMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListFleetMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListFleetMetricsOutput{}
	fnCacher := func(out *iot.ListFleetMetricsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListFleetMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListFleetMetricsWithContext(ctx context.Context, input *iot.ListFleetMetricsInput, opts ...request.Option) (*iot.ListFleetMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListFleetMetricsOutput), nil
	}
	out, err := c.IoTAPI.ListFleetMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListIndices(input *iot.ListIndicesInput) (*iot.ListIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListIndicesOutput), nil
	}
	out, err := c.IoTAPI.ListIndices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListIndicesPages(input *iot.ListIndicesInput, fn func(*iot.ListIndicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListIndicesOutput{}
	fnCacher := func(out *iot.ListIndicesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListIndicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListIndicesPagesWithContext(ctx context.Context, input *iot.ListIndicesInput, fn func(*iot.ListIndicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListIndicesOutput{}
	fnCacher := func(out *iot.ListIndicesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListIndicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListIndicesWithContext(ctx context.Context, input *iot.ListIndicesInput, opts ...request.Option) (*iot.ListIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListIndicesOutput), nil
	}
	out, err := c.IoTAPI.ListIndicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobExecutionsForJob(input *iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobExecutionsForJobOutput), nil
	}
	out, err := c.IoTAPI.ListJobExecutionsForJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobExecutionsForJobPages(input *iot.ListJobExecutionsForJobInput, fn func(*iot.ListJobExecutionsForJobOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobExecutionsForJobOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobExecutionsForJobOutput{}
	fnCacher := func(out *iot.ListJobExecutionsForJobOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobExecutionsForJobPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobExecutionsForJobPagesWithContext(ctx context.Context, input *iot.ListJobExecutionsForJobInput, fn func(*iot.ListJobExecutionsForJobOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobExecutionsForJobOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobExecutionsForJobOutput{}
	fnCacher := func(out *iot.ListJobExecutionsForJobOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobExecutionsForJobPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobExecutionsForJobWithContext(ctx context.Context, input *iot.ListJobExecutionsForJobInput, opts ...request.Option) (*iot.ListJobExecutionsForJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobExecutionsForJobOutput), nil
	}
	out, err := c.IoTAPI.ListJobExecutionsForJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobExecutionsForThing(input *iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobExecutionsForThingOutput), nil
	}
	out, err := c.IoTAPI.ListJobExecutionsForThing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobExecutionsForThingPages(input *iot.ListJobExecutionsForThingInput, fn func(*iot.ListJobExecutionsForThingOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobExecutionsForThingOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobExecutionsForThingOutput{}
	fnCacher := func(out *iot.ListJobExecutionsForThingOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobExecutionsForThingPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobExecutionsForThingPagesWithContext(ctx context.Context, input *iot.ListJobExecutionsForThingInput, fn func(*iot.ListJobExecutionsForThingOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobExecutionsForThingOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobExecutionsForThingOutput{}
	fnCacher := func(out *iot.ListJobExecutionsForThingOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobExecutionsForThingPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobExecutionsForThingWithContext(ctx context.Context, input *iot.ListJobExecutionsForThingInput, opts ...request.Option) (*iot.ListJobExecutionsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobExecutionsForThingOutput), nil
	}
	out, err := c.IoTAPI.ListJobExecutionsForThingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobTemplates(input *iot.ListJobTemplatesInput) (*iot.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListJobTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobTemplatesPages(input *iot.ListJobTemplatesInput, fn func(*iot.ListJobTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobTemplatesOutput{}
	fnCacher := func(out *iot.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobTemplatesPagesWithContext(ctx context.Context, input *iot.ListJobTemplatesInput, fn func(*iot.ListJobTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobTemplatesOutput{}
	fnCacher := func(out *iot.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobTemplatesWithContext(ctx context.Context, input *iot.ListJobTemplatesInput, opts ...request.Option) (*iot.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListJobTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobs(input *iot.ListJobsInput) (*iot.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobsOutput), nil
	}
	out, err := c.IoTAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListJobsPages(input *iot.ListJobsInput, fn func(*iot.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobsOutput{}
	fnCacher := func(out *iot.ListJobsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobsPagesWithContext(ctx context.Context, input *iot.ListJobsInput, fn func(*iot.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListJobsOutput{}
	fnCacher := func(out *iot.ListJobsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListJobsWithContext(ctx context.Context, input *iot.ListJobsInput, opts ...request.Option) (*iot.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListJobsOutput), nil
	}
	out, err := c.IoTAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListManagedJobTemplates(input *iot.ListManagedJobTemplatesInput) (*iot.ListManagedJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListManagedJobTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListManagedJobTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListManagedJobTemplatesWithContext(ctx context.Context, input *iot.ListManagedJobTemplatesInput, opts ...request.Option) (*iot.ListManagedJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListManagedJobTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListManagedJobTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListMetricValues(input *iot.ListMetricValuesInput) (*iot.ListMetricValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListMetricValuesOutput), nil
	}
	out, err := c.IoTAPI.ListMetricValues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListMetricValuesPages(input *iot.ListMetricValuesInput, fn func(*iot.ListMetricValuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListMetricValuesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListMetricValuesOutput{}
	fnCacher := func(out *iot.ListMetricValuesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListMetricValuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListMetricValuesPagesWithContext(ctx context.Context, input *iot.ListMetricValuesInput, fn func(*iot.ListMetricValuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListMetricValuesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListMetricValuesOutput{}
	fnCacher := func(out *iot.ListMetricValuesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListMetricValuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListMetricValuesWithContext(ctx context.Context, input *iot.ListMetricValuesInput, opts ...request.Option) (*iot.ListMetricValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListMetricValuesOutput), nil
	}
	out, err := c.IoTAPI.ListMetricValuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListMitigationActions(input *iot.ListMitigationActionsInput) (*iot.ListMitigationActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListMitigationActionsOutput), nil
	}
	out, err := c.IoTAPI.ListMitigationActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListMitigationActionsPages(input *iot.ListMitigationActionsInput, fn func(*iot.ListMitigationActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListMitigationActionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListMitigationActionsOutput{}
	fnCacher := func(out *iot.ListMitigationActionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListMitigationActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListMitigationActionsPagesWithContext(ctx context.Context, input *iot.ListMitigationActionsInput, fn func(*iot.ListMitigationActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListMitigationActionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListMitigationActionsOutput{}
	fnCacher := func(out *iot.ListMitigationActionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListMitigationActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListMitigationActionsWithContext(ctx context.Context, input *iot.ListMitigationActionsInput, opts ...request.Option) (*iot.ListMitigationActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListMitigationActionsOutput), nil
	}
	out, err := c.IoTAPI.ListMitigationActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListOTAUpdates(input *iot.ListOTAUpdatesInput) (*iot.ListOTAUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListOTAUpdatesOutput), nil
	}
	out, err := c.IoTAPI.ListOTAUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListOTAUpdatesPages(input *iot.ListOTAUpdatesInput, fn func(*iot.ListOTAUpdatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListOTAUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListOTAUpdatesOutput{}
	fnCacher := func(out *iot.ListOTAUpdatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListOTAUpdatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListOTAUpdatesPagesWithContext(ctx context.Context, input *iot.ListOTAUpdatesInput, fn func(*iot.ListOTAUpdatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListOTAUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListOTAUpdatesOutput{}
	fnCacher := func(out *iot.ListOTAUpdatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListOTAUpdatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListOTAUpdatesWithContext(ctx context.Context, input *iot.ListOTAUpdatesInput, opts ...request.Option) (*iot.ListOTAUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListOTAUpdatesOutput), nil
	}
	out, err := c.IoTAPI.ListOTAUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListOutgoingCertificates(input *iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListOutgoingCertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListOutgoingCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListOutgoingCertificatesPages(input *iot.ListOutgoingCertificatesInput, fn func(*iot.ListOutgoingCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListOutgoingCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListOutgoingCertificatesOutput{}
	fnCacher := func(out *iot.ListOutgoingCertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListOutgoingCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListOutgoingCertificatesPagesWithContext(ctx context.Context, input *iot.ListOutgoingCertificatesInput, fn func(*iot.ListOutgoingCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListOutgoingCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListOutgoingCertificatesOutput{}
	fnCacher := func(out *iot.ListOutgoingCertificatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListOutgoingCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListOutgoingCertificatesWithContext(ctx context.Context, input *iot.ListOutgoingCertificatesInput, opts ...request.Option) (*iot.ListOutgoingCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListOutgoingCertificatesOutput), nil
	}
	out, err := c.IoTAPI.ListOutgoingCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPolicies(input *iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPoliciesPages(input *iot.ListPoliciesInput, fn func(*iot.ListPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPoliciesOutput{}
	fnCacher := func(out *iot.ListPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPoliciesPagesWithContext(ctx context.Context, input *iot.ListPoliciesInput, fn func(*iot.ListPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPoliciesOutput{}
	fnCacher := func(out *iot.ListPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPoliciesWithContext(ctx context.Context, input *iot.ListPoliciesInput, opts ...request.Option) (*iot.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPolicyPrincipals(input *iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPolicyPrincipalsOutput), nil
	}
	out, err := c.IoTAPI.ListPolicyPrincipals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPolicyPrincipalsPages(input *iot.ListPolicyPrincipalsInput, fn func(*iot.ListPolicyPrincipalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPolicyPrincipalsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPolicyPrincipalsOutput{}
	fnCacher := func(out *iot.ListPolicyPrincipalsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPolicyPrincipalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPolicyPrincipalsPagesWithContext(ctx context.Context, input *iot.ListPolicyPrincipalsInput, fn func(*iot.ListPolicyPrincipalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPolicyPrincipalsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPolicyPrincipalsOutput{}
	fnCacher := func(out *iot.ListPolicyPrincipalsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPolicyPrincipalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPolicyPrincipalsWithContext(ctx context.Context, input *iot.ListPolicyPrincipalsInput, opts ...request.Option) (*iot.ListPolicyPrincipalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPolicyPrincipalsOutput), nil
	}
	out, err := c.IoTAPI.ListPolicyPrincipalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPolicyVersions(input *iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPolicyVersionsOutput), nil
	}
	out, err := c.IoTAPI.ListPolicyVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPolicyVersionsWithContext(ctx context.Context, input *iot.ListPolicyVersionsInput, opts ...request.Option) (*iot.ListPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPolicyVersionsOutput), nil
	}
	out, err := c.IoTAPI.ListPolicyVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPrincipalPolicies(input *iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPrincipalPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListPrincipalPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPrincipalPoliciesPages(input *iot.ListPrincipalPoliciesInput, fn func(*iot.ListPrincipalPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPrincipalPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPrincipalPoliciesOutput{}
	fnCacher := func(out *iot.ListPrincipalPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPrincipalPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPrincipalPoliciesPagesWithContext(ctx context.Context, input *iot.ListPrincipalPoliciesInput, fn func(*iot.ListPrincipalPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPrincipalPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPrincipalPoliciesOutput{}
	fnCacher := func(out *iot.ListPrincipalPoliciesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPrincipalPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPrincipalPoliciesWithContext(ctx context.Context, input *iot.ListPrincipalPoliciesInput, opts ...request.Option) (*iot.ListPrincipalPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPrincipalPoliciesOutput), nil
	}
	out, err := c.IoTAPI.ListPrincipalPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPrincipalThings(input *iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPrincipalThingsOutput), nil
	}
	out, err := c.IoTAPI.ListPrincipalThings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListPrincipalThingsPages(input *iot.ListPrincipalThingsInput, fn func(*iot.ListPrincipalThingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPrincipalThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPrincipalThingsOutput{}
	fnCacher := func(out *iot.ListPrincipalThingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPrincipalThingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPrincipalThingsPagesWithContext(ctx context.Context, input *iot.ListPrincipalThingsInput, fn func(*iot.ListPrincipalThingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListPrincipalThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListPrincipalThingsOutput{}
	fnCacher := func(out *iot.ListPrincipalThingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListPrincipalThingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListPrincipalThingsWithContext(ctx context.Context, input *iot.ListPrincipalThingsInput, opts ...request.Option) (*iot.ListPrincipalThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListPrincipalThingsOutput), nil
	}
	out, err := c.IoTAPI.ListPrincipalThingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListProvisioningTemplateVersions(input *iot.ListProvisioningTemplateVersionsInput) (*iot.ListProvisioningTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListProvisioningTemplateVersionsOutput), nil
	}
	out, err := c.IoTAPI.ListProvisioningTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListProvisioningTemplateVersionsPages(input *iot.ListProvisioningTemplateVersionsInput, fn func(*iot.ListProvisioningTemplateVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListProvisioningTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListProvisioningTemplateVersionsOutput{}
	fnCacher := func(out *iot.ListProvisioningTemplateVersionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListProvisioningTemplateVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListProvisioningTemplateVersionsPagesWithContext(ctx context.Context, input *iot.ListProvisioningTemplateVersionsInput, fn func(*iot.ListProvisioningTemplateVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListProvisioningTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListProvisioningTemplateVersionsOutput{}
	fnCacher := func(out *iot.ListProvisioningTemplateVersionsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListProvisioningTemplateVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListProvisioningTemplateVersionsWithContext(ctx context.Context, input *iot.ListProvisioningTemplateVersionsInput, opts ...request.Option) (*iot.ListProvisioningTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListProvisioningTemplateVersionsOutput), nil
	}
	out, err := c.IoTAPI.ListProvisioningTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListProvisioningTemplates(input *iot.ListProvisioningTemplatesInput) (*iot.ListProvisioningTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListProvisioningTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListProvisioningTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListProvisioningTemplatesPages(input *iot.ListProvisioningTemplatesInput, fn func(*iot.ListProvisioningTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListProvisioningTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListProvisioningTemplatesOutput{}
	fnCacher := func(out *iot.ListProvisioningTemplatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListProvisioningTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListProvisioningTemplatesPagesWithContext(ctx context.Context, input *iot.ListProvisioningTemplatesInput, fn func(*iot.ListProvisioningTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListProvisioningTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListProvisioningTemplatesOutput{}
	fnCacher := func(out *iot.ListProvisioningTemplatesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListProvisioningTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListProvisioningTemplatesWithContext(ctx context.Context, input *iot.ListProvisioningTemplatesInput, opts ...request.Option) (*iot.ListProvisioningTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListProvisioningTemplatesOutput), nil
	}
	out, err := c.IoTAPI.ListProvisioningTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListRelatedResourcesForAuditFinding(input *iot.ListRelatedResourcesForAuditFindingInput) (*iot.ListRelatedResourcesForAuditFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListRelatedResourcesForAuditFindingOutput), nil
	}
	out, err := c.IoTAPI.ListRelatedResourcesForAuditFinding(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListRelatedResourcesForAuditFindingWithContext(ctx context.Context, input *iot.ListRelatedResourcesForAuditFindingInput, opts ...request.Option) (*iot.ListRelatedResourcesForAuditFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListRelatedResourcesForAuditFindingOutput), nil
	}
	out, err := c.IoTAPI.ListRelatedResourcesForAuditFindingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListRoleAliases(input *iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListRoleAliasesOutput), nil
	}
	out, err := c.IoTAPI.ListRoleAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListRoleAliasesPages(input *iot.ListRoleAliasesInput, fn func(*iot.ListRoleAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListRoleAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListRoleAliasesOutput{}
	fnCacher := func(out *iot.ListRoleAliasesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListRoleAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListRoleAliasesPagesWithContext(ctx context.Context, input *iot.ListRoleAliasesInput, fn func(*iot.ListRoleAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListRoleAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListRoleAliasesOutput{}
	fnCacher := func(out *iot.ListRoleAliasesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListRoleAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListRoleAliasesWithContext(ctx context.Context, input *iot.ListRoleAliasesInput, opts ...request.Option) (*iot.ListRoleAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListRoleAliasesOutput), nil
	}
	out, err := c.IoTAPI.ListRoleAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListScheduledAudits(input *iot.ListScheduledAuditsInput) (*iot.ListScheduledAuditsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListScheduledAuditsOutput), nil
	}
	out, err := c.IoTAPI.ListScheduledAudits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListScheduledAuditsPages(input *iot.ListScheduledAuditsInput, fn func(*iot.ListScheduledAuditsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListScheduledAuditsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListScheduledAuditsOutput{}
	fnCacher := func(out *iot.ListScheduledAuditsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListScheduledAuditsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListScheduledAuditsPagesWithContext(ctx context.Context, input *iot.ListScheduledAuditsInput, fn func(*iot.ListScheduledAuditsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListScheduledAuditsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListScheduledAuditsOutput{}
	fnCacher := func(out *iot.ListScheduledAuditsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListScheduledAuditsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListScheduledAuditsWithContext(ctx context.Context, input *iot.ListScheduledAuditsInput, opts ...request.Option) (*iot.ListScheduledAuditsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListScheduledAuditsOutput), nil
	}
	out, err := c.IoTAPI.ListScheduledAuditsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListSecurityProfiles(input *iot.ListSecurityProfilesInput) (*iot.ListSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListSecurityProfilesOutput), nil
	}
	out, err := c.IoTAPI.ListSecurityProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListSecurityProfilesForTarget(input *iot.ListSecurityProfilesForTargetInput) (*iot.ListSecurityProfilesForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListSecurityProfilesForTargetOutput), nil
	}
	out, err := c.IoTAPI.ListSecurityProfilesForTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListSecurityProfilesForTargetPages(input *iot.ListSecurityProfilesForTargetInput, fn func(*iot.ListSecurityProfilesForTargetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListSecurityProfilesForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListSecurityProfilesForTargetOutput{}
	fnCacher := func(out *iot.ListSecurityProfilesForTargetOutput, more bool) bool {
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
	if err := c.IoTAPI.ListSecurityProfilesForTargetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListSecurityProfilesForTargetPagesWithContext(ctx context.Context, input *iot.ListSecurityProfilesForTargetInput, fn func(*iot.ListSecurityProfilesForTargetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListSecurityProfilesForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListSecurityProfilesForTargetOutput{}
	fnCacher := func(out *iot.ListSecurityProfilesForTargetOutput, more bool) bool {
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
	if err := c.IoTAPI.ListSecurityProfilesForTargetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListSecurityProfilesForTargetWithContext(ctx context.Context, input *iot.ListSecurityProfilesForTargetInput, opts ...request.Option) (*iot.ListSecurityProfilesForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListSecurityProfilesForTargetOutput), nil
	}
	out, err := c.IoTAPI.ListSecurityProfilesForTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListSecurityProfilesPages(input *iot.ListSecurityProfilesInput, fn func(*iot.ListSecurityProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListSecurityProfilesOutput{}
	fnCacher := func(out *iot.ListSecurityProfilesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListSecurityProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListSecurityProfilesPagesWithContext(ctx context.Context, input *iot.ListSecurityProfilesInput, fn func(*iot.ListSecurityProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListSecurityProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListSecurityProfilesOutput{}
	fnCacher := func(out *iot.ListSecurityProfilesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListSecurityProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListSecurityProfilesWithContext(ctx context.Context, input *iot.ListSecurityProfilesInput, opts ...request.Option) (*iot.ListSecurityProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListSecurityProfilesOutput), nil
	}
	out, err := c.IoTAPI.ListSecurityProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListStreams(input *iot.ListStreamsInput) (*iot.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListStreamsOutput), nil
	}
	out, err := c.IoTAPI.ListStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListStreamsPages(input *iot.ListStreamsInput, fn func(*iot.ListStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListStreamsOutput{}
	fnCacher := func(out *iot.ListStreamsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListStreamsPagesWithContext(ctx context.Context, input *iot.ListStreamsInput, fn func(*iot.ListStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListStreamsOutput{}
	fnCacher := func(out *iot.ListStreamsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListStreamsWithContext(ctx context.Context, input *iot.ListStreamsInput, opts ...request.Option) (*iot.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListStreamsOutput), nil
	}
	out, err := c.IoTAPI.ListStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTagsForResource(input *iot.ListTagsForResourceInput) (*iot.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTagsForResourcePages(input *iot.ListTagsForResourceInput, fn func(*iot.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTagsForResourceOutput{}
	fnCacher := func(out *iot.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTagsForResourcePagesWithContext(ctx context.Context, input *iot.ListTagsForResourceInput, fn func(*iot.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTagsForResourceOutput{}
	fnCacher := func(out *iot.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTagsForResourceWithContext(ctx context.Context, input *iot.ListTagsForResourceInput, opts ...request.Option) (*iot.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTargetsForPolicy(input *iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTargetsForPolicyOutput), nil
	}
	out, err := c.IoTAPI.ListTargetsForPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTargetsForPolicyPages(input *iot.ListTargetsForPolicyInput, fn func(*iot.ListTargetsForPolicyOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTargetsForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTargetsForPolicyOutput{}
	fnCacher := func(out *iot.ListTargetsForPolicyOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTargetsForPolicyPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTargetsForPolicyPagesWithContext(ctx context.Context, input *iot.ListTargetsForPolicyInput, fn func(*iot.ListTargetsForPolicyOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTargetsForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTargetsForPolicyOutput{}
	fnCacher := func(out *iot.ListTargetsForPolicyOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTargetsForPolicyPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTargetsForPolicyWithContext(ctx context.Context, input *iot.ListTargetsForPolicyInput, opts ...request.Option) (*iot.ListTargetsForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTargetsForPolicyOutput), nil
	}
	out, err := c.IoTAPI.ListTargetsForPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTargetsForSecurityProfile(input *iot.ListTargetsForSecurityProfileInput) (*iot.ListTargetsForSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTargetsForSecurityProfileOutput), nil
	}
	out, err := c.IoTAPI.ListTargetsForSecurityProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTargetsForSecurityProfilePages(input *iot.ListTargetsForSecurityProfileInput, fn func(*iot.ListTargetsForSecurityProfileOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTargetsForSecurityProfileOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTargetsForSecurityProfileOutput{}
	fnCacher := func(out *iot.ListTargetsForSecurityProfileOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTargetsForSecurityProfilePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTargetsForSecurityProfilePagesWithContext(ctx context.Context, input *iot.ListTargetsForSecurityProfileInput, fn func(*iot.ListTargetsForSecurityProfileOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTargetsForSecurityProfileOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTargetsForSecurityProfileOutput{}
	fnCacher := func(out *iot.ListTargetsForSecurityProfileOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTargetsForSecurityProfilePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTargetsForSecurityProfileWithContext(ctx context.Context, input *iot.ListTargetsForSecurityProfileInput, opts ...request.Option) (*iot.ListTargetsForSecurityProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTargetsForSecurityProfileOutput), nil
	}
	out, err := c.IoTAPI.ListTargetsForSecurityProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingGroups(input *iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingGroupsOutput), nil
	}
	out, err := c.IoTAPI.ListThingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingGroupsForThing(input *iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingGroupsForThingOutput), nil
	}
	out, err := c.IoTAPI.ListThingGroupsForThing(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingGroupsForThingPages(input *iot.ListThingGroupsForThingInput, fn func(*iot.ListThingGroupsForThingOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingGroupsForThingOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingGroupsForThingOutput{}
	fnCacher := func(out *iot.ListThingGroupsForThingOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingGroupsForThingPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingGroupsForThingPagesWithContext(ctx context.Context, input *iot.ListThingGroupsForThingInput, fn func(*iot.ListThingGroupsForThingOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingGroupsForThingOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingGroupsForThingOutput{}
	fnCacher := func(out *iot.ListThingGroupsForThingOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingGroupsForThingPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingGroupsForThingWithContext(ctx context.Context, input *iot.ListThingGroupsForThingInput, opts ...request.Option) (*iot.ListThingGroupsForThingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingGroupsForThingOutput), nil
	}
	out, err := c.IoTAPI.ListThingGroupsForThingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingGroupsPages(input *iot.ListThingGroupsInput, fn func(*iot.ListThingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingGroupsOutput{}
	fnCacher := func(out *iot.ListThingGroupsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingGroupsPagesWithContext(ctx context.Context, input *iot.ListThingGroupsInput, fn func(*iot.ListThingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingGroupsOutput{}
	fnCacher := func(out *iot.ListThingGroupsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingGroupsWithContext(ctx context.Context, input *iot.ListThingGroupsInput, opts ...request.Option) (*iot.ListThingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingGroupsOutput), nil
	}
	out, err := c.IoTAPI.ListThingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingPrincipals(input *iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingPrincipalsOutput), nil
	}
	out, err := c.IoTAPI.ListThingPrincipals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingPrincipalsPages(input *iot.ListThingPrincipalsInput, fn func(*iot.ListThingPrincipalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingPrincipalsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingPrincipalsOutput{}
	fnCacher := func(out *iot.ListThingPrincipalsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingPrincipalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingPrincipalsPagesWithContext(ctx context.Context, input *iot.ListThingPrincipalsInput, fn func(*iot.ListThingPrincipalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingPrincipalsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingPrincipalsOutput{}
	fnCacher := func(out *iot.ListThingPrincipalsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingPrincipalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingPrincipalsWithContext(ctx context.Context, input *iot.ListThingPrincipalsInput, opts ...request.Option) (*iot.ListThingPrincipalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingPrincipalsOutput), nil
	}
	out, err := c.IoTAPI.ListThingPrincipalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingRegistrationTaskReports(input *iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingRegistrationTaskReportsOutput), nil
	}
	out, err := c.IoTAPI.ListThingRegistrationTaskReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingRegistrationTaskReportsPages(input *iot.ListThingRegistrationTaskReportsInput, fn func(*iot.ListThingRegistrationTaskReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingRegistrationTaskReportsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingRegistrationTaskReportsOutput{}
	fnCacher := func(out *iot.ListThingRegistrationTaskReportsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingRegistrationTaskReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingRegistrationTaskReportsPagesWithContext(ctx context.Context, input *iot.ListThingRegistrationTaskReportsInput, fn func(*iot.ListThingRegistrationTaskReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingRegistrationTaskReportsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingRegistrationTaskReportsOutput{}
	fnCacher := func(out *iot.ListThingRegistrationTaskReportsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingRegistrationTaskReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingRegistrationTaskReportsWithContext(ctx context.Context, input *iot.ListThingRegistrationTaskReportsInput, opts ...request.Option) (*iot.ListThingRegistrationTaskReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingRegistrationTaskReportsOutput), nil
	}
	out, err := c.IoTAPI.ListThingRegistrationTaskReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingRegistrationTasks(input *iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingRegistrationTasksOutput), nil
	}
	out, err := c.IoTAPI.ListThingRegistrationTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingRegistrationTasksPages(input *iot.ListThingRegistrationTasksInput, fn func(*iot.ListThingRegistrationTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingRegistrationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingRegistrationTasksOutput{}
	fnCacher := func(out *iot.ListThingRegistrationTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingRegistrationTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingRegistrationTasksPagesWithContext(ctx context.Context, input *iot.ListThingRegistrationTasksInput, fn func(*iot.ListThingRegistrationTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingRegistrationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingRegistrationTasksOutput{}
	fnCacher := func(out *iot.ListThingRegistrationTasksOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingRegistrationTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingRegistrationTasksWithContext(ctx context.Context, input *iot.ListThingRegistrationTasksInput, opts ...request.Option) (*iot.ListThingRegistrationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingRegistrationTasksOutput), nil
	}
	out, err := c.IoTAPI.ListThingRegistrationTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingTypes(input *iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingTypesOutput), nil
	}
	out, err := c.IoTAPI.ListThingTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingTypesPages(input *iot.ListThingTypesInput, fn func(*iot.ListThingTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingTypesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingTypesOutput{}
	fnCacher := func(out *iot.ListThingTypesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingTypesPagesWithContext(ctx context.Context, input *iot.ListThingTypesInput, fn func(*iot.ListThingTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingTypesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingTypesOutput{}
	fnCacher := func(out *iot.ListThingTypesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingTypesWithContext(ctx context.Context, input *iot.ListThingTypesInput, opts ...request.Option) (*iot.ListThingTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingTypesOutput), nil
	}
	out, err := c.IoTAPI.ListThingTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThings(input *iot.ListThingsInput) (*iot.ListThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsOutput), nil
	}
	out, err := c.IoTAPI.ListThings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingsInBillingGroup(input *iot.ListThingsInBillingGroupInput) (*iot.ListThingsInBillingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsInBillingGroupOutput), nil
	}
	out, err := c.IoTAPI.ListThingsInBillingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingsInBillingGroupPages(input *iot.ListThingsInBillingGroupInput, fn func(*iot.ListThingsInBillingGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsInBillingGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsInBillingGroupOutput{}
	fnCacher := func(out *iot.ListThingsInBillingGroupOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsInBillingGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsInBillingGroupPagesWithContext(ctx context.Context, input *iot.ListThingsInBillingGroupInput, fn func(*iot.ListThingsInBillingGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsInBillingGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsInBillingGroupOutput{}
	fnCacher := func(out *iot.ListThingsInBillingGroupOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsInBillingGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsInBillingGroupWithContext(ctx context.Context, input *iot.ListThingsInBillingGroupInput, opts ...request.Option) (*iot.ListThingsInBillingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsInBillingGroupOutput), nil
	}
	out, err := c.IoTAPI.ListThingsInBillingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingsInThingGroup(input *iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsInThingGroupOutput), nil
	}
	out, err := c.IoTAPI.ListThingsInThingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingsInThingGroupPages(input *iot.ListThingsInThingGroupInput, fn func(*iot.ListThingsInThingGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsInThingGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsInThingGroupOutput{}
	fnCacher := func(out *iot.ListThingsInThingGroupOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsInThingGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsInThingGroupPagesWithContext(ctx context.Context, input *iot.ListThingsInThingGroupInput, fn func(*iot.ListThingsInThingGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsInThingGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsInThingGroupOutput{}
	fnCacher := func(out *iot.ListThingsInThingGroupOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsInThingGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsInThingGroupWithContext(ctx context.Context, input *iot.ListThingsInThingGroupInput, opts ...request.Option) (*iot.ListThingsInThingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsInThingGroupOutput), nil
	}
	out, err := c.IoTAPI.ListThingsInThingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListThingsPages(input *iot.ListThingsInput, fn func(*iot.ListThingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsOutput{}
	fnCacher := func(out *iot.ListThingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsPagesWithContext(ctx context.Context, input *iot.ListThingsInput, fn func(*iot.ListThingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListThingsOutput{}
	fnCacher := func(out *iot.ListThingsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListThingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListThingsWithContext(ctx context.Context, input *iot.ListThingsInput, opts ...request.Option) (*iot.ListThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListThingsOutput), nil
	}
	out, err := c.IoTAPI.ListThingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTopicRuleDestinations(input *iot.ListTopicRuleDestinationsInput) (*iot.ListTopicRuleDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTopicRuleDestinationsOutput), nil
	}
	out, err := c.IoTAPI.ListTopicRuleDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTopicRuleDestinationsPages(input *iot.ListTopicRuleDestinationsInput, fn func(*iot.ListTopicRuleDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTopicRuleDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTopicRuleDestinationsOutput{}
	fnCacher := func(out *iot.ListTopicRuleDestinationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTopicRuleDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTopicRuleDestinationsPagesWithContext(ctx context.Context, input *iot.ListTopicRuleDestinationsInput, fn func(*iot.ListTopicRuleDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTopicRuleDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTopicRuleDestinationsOutput{}
	fnCacher := func(out *iot.ListTopicRuleDestinationsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTopicRuleDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTopicRuleDestinationsWithContext(ctx context.Context, input *iot.ListTopicRuleDestinationsInput, opts ...request.Option) (*iot.ListTopicRuleDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTopicRuleDestinationsOutput), nil
	}
	out, err := c.IoTAPI.ListTopicRuleDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTopicRules(input *iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTopicRulesOutput), nil
	}
	out, err := c.IoTAPI.ListTopicRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListTopicRulesPages(input *iot.ListTopicRulesInput, fn func(*iot.ListTopicRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTopicRulesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTopicRulesOutput{}
	fnCacher := func(out *iot.ListTopicRulesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTopicRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTopicRulesPagesWithContext(ctx context.Context, input *iot.ListTopicRulesInput, fn func(*iot.ListTopicRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListTopicRulesOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListTopicRulesOutput{}
	fnCacher := func(out *iot.ListTopicRulesOutput, more bool) bool {
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
	if err := c.IoTAPI.ListTopicRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListTopicRulesWithContext(ctx context.Context, input *iot.ListTopicRulesInput, opts ...request.Option) (*iot.ListTopicRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListTopicRulesOutput), nil
	}
	out, err := c.IoTAPI.ListTopicRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListV2LoggingLevels(input *iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListV2LoggingLevelsOutput), nil
	}
	out, err := c.IoTAPI.ListV2LoggingLevels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListV2LoggingLevelsPages(input *iot.ListV2LoggingLevelsInput, fn func(*iot.ListV2LoggingLevelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListV2LoggingLevelsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListV2LoggingLevelsOutput{}
	fnCacher := func(out *iot.ListV2LoggingLevelsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListV2LoggingLevelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListV2LoggingLevelsPagesWithContext(ctx context.Context, input *iot.ListV2LoggingLevelsInput, fn func(*iot.ListV2LoggingLevelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListV2LoggingLevelsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListV2LoggingLevelsOutput{}
	fnCacher := func(out *iot.ListV2LoggingLevelsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListV2LoggingLevelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListV2LoggingLevelsWithContext(ctx context.Context, input *iot.ListV2LoggingLevelsInput, opts ...request.Option) (*iot.ListV2LoggingLevelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListV2LoggingLevelsOutput), nil
	}
	out, err := c.IoTAPI.ListV2LoggingLevelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListViolationEvents(input *iot.ListViolationEventsInput) (*iot.ListViolationEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListViolationEventsOutput), nil
	}
	out, err := c.IoTAPI.ListViolationEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) ListViolationEventsPages(input *iot.ListViolationEventsInput, fn func(*iot.ListViolationEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListViolationEventsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListViolationEventsOutput{}
	fnCacher := func(out *iot.ListViolationEventsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListViolationEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListViolationEventsPagesWithContext(ctx context.Context, input *iot.ListViolationEventsInput, fn func(*iot.ListViolationEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iot.ListViolationEventsOutput), false)
		return nil
	}
	cachable := true
	output := &iot.ListViolationEventsOutput{}
	fnCacher := func(out *iot.ListViolationEventsOutput, more bool) bool {
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
	if err := c.IoTAPI.ListViolationEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoT) ListViolationEventsWithContext(ctx context.Context, input *iot.ListViolationEventsInput, opts ...request.Option) (*iot.ListViolationEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.ListViolationEventsOutput), nil
	}
	out, err := c.IoTAPI.ListViolationEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) SearchIndex(input *iot.SearchIndexInput) (*iot.SearchIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.SearchIndexOutput), nil
	}
	out, err := c.IoTAPI.SearchIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT) SearchIndexWithContext(ctx context.Context, input *iot.SearchIndexInput, opts ...request.Option) (*iot.SearchIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot.SearchIndexOutput), nil
	}
	out, err := c.IoTAPI.SearchIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

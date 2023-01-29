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
package databasemigrationservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DatabaseMigrationService struct {
	databasemigrationserviceiface.DatabaseMigrationServiceAPI
	cache *cache.Cache
}

func New(databasemigrationserviceapi databasemigrationserviceiface.DatabaseMigrationServiceAPI) *DatabaseMigrationService {
	return &DatabaseMigrationService{
		DatabaseMigrationServiceAPI: databasemigrationserviceapi,
		cache:                       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DatabaseMigrationService) DescribeAccountAttributes(input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeAccountAttributesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeAccountAttributesWithContext(ctx context.Context, input *databasemigrationservice.DescribeAccountAttributesInput, opts ...request.Option) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeAccountAttributesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeApplicableIndividualAssessments(input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeApplicableIndividualAssessments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeApplicableIndividualAssessmentsPages(input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, fn func(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeApplicableIndividualAssessmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeApplicableIndividualAssessmentsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, fn func(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeApplicableIndividualAssessmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeApplicableIndividualAssessmentsWithContext(ctx context.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, opts ...request.Option) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeApplicableIndividualAssessmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeCertificates(input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeCertificatesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeCertificatesPages(input *databasemigrationservice.DescribeCertificatesInput, fn func(*databasemigrationservice.DescribeCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeCertificatesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeCertificatesPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeCertificatesInput, fn func(*databasemigrationservice.DescribeCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeCertificatesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeCertificatesWithContext(ctx context.Context, input *databasemigrationservice.DescribeCertificatesInput, opts ...request.Option) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeCertificatesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeConnections(input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeConnectionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeConnectionsPages(input *databasemigrationservice.DescribeConnectionsInput, fn func(*databasemigrationservice.DescribeConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeConnectionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeConnectionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeConnectionsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput, fn func(*databasemigrationservice.DescribeConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeConnectionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeConnectionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeConnectionsWithContext(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput, opts ...request.Option) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeConnectionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpointSettings(input *databasemigrationservice.DescribeEndpointSettingsInput) (*databasemigrationservice.DescribeEndpointSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointSettingsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpointSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpointSettingsPages(input *databasemigrationservice.DescribeEndpointSettingsInput, fn func(*databasemigrationservice.DescribeEndpointSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointSettingsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointSettingsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointSettingsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointSettingsInput, fn func(*databasemigrationservice.DescribeEndpointSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointSettingsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointSettingsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointSettingsWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointSettingsInput, opts ...request.Option) (*databasemigrationservice.DescribeEndpointSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointSettingsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpointSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpointTypes(input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointTypesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpointTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpointTypesPages(input *databasemigrationservice.DescribeEndpointTypesInput, fn func(*databasemigrationservice.DescribeEndpointTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointTypesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointTypesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointTypesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointTypesPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointTypesInput, fn func(*databasemigrationservice.DescribeEndpointTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointTypesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointTypesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointTypesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointTypesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointTypesInput, opts ...request.Option) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointTypesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpointTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpoints(input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEndpointsPages(input *databasemigrationservice.DescribeEndpointsInput, fn func(*databasemigrationservice.DescribeEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput, fn func(*databasemigrationservice.DescribeEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEndpointsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEndpointsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEndpointsWithContext(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput, opts ...request.Option) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEndpointsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEventCategories(input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventCategoriesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEventCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEventCategoriesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEventCategoriesInput, opts ...request.Option) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventCategoriesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEventCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEventSubscriptions(input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEventSubscriptionsPages(input *databasemigrationservice.DescribeEventSubscriptionsInput, fn func(*databasemigrationservice.DescribeEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEventSubscriptionsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput, fn func(*databasemigrationservice.DescribeEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEventSubscriptionsWithContext(ctx context.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput, opts ...request.Option) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEvents(input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeEventsPages(input *databasemigrationservice.DescribeEventsInput, fn func(*databasemigrationservice.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEventsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEventsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEventsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeEventsInput, fn func(*databasemigrationservice.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeEventsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeEventsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeEventsWithContext(ctx context.Context, input *databasemigrationservice.DescribeEventsInput, opts ...request.Option) (*databasemigrationservice.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeEventsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorCollectors(input *databasemigrationservice.DescribeFleetAdvisorCollectorsInput) (*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorCollectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorCollectorsPages(input *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, fn func(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorCollectorsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorCollectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorCollectorsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, fn func(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorCollectorsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorCollectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorCollectorsWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, opts ...request.Option) (*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorCollectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorDatabases(input *databasemigrationservice.DescribeFleetAdvisorDatabasesInput) (*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorDatabasesPages(input *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, fn func(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorDatabasesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorDatabasesPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, fn func(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorDatabasesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorDatabasesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, opts ...request.Option) (*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorLsaAnalysis(input *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput) (*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorLsaAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorLsaAnalysisPages(input *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, fn func(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorLsaAnalysisPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorLsaAnalysisPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, fn func(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorLsaAnalysisPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorLsaAnalysisWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, opts ...request.Option) (*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorLsaAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemaObjectSummary(input *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput) (*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemaObjectSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemaObjectSummaryPages(input *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, fn func(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemaObjectSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemaObjectSummaryPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, fn func(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemaObjectSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemaObjectSummaryWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, opts ...request.Option) (*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemaObjectSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemas(input *databasemigrationservice.DescribeFleetAdvisorSchemasInput) (*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemasPages(input *databasemigrationservice.DescribeFleetAdvisorSchemasInput, fn func(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorSchemasOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorSchemasOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemasPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorSchemasInput, fn func(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeFleetAdvisorSchemasOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeFleetAdvisorSchemasOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeFleetAdvisorSchemasWithContext(ctx context.Context, input *databasemigrationservice.DescribeFleetAdvisorSchemasInput, opts ...request.Option) (*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeFleetAdvisorSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeOrderableReplicationInstances(input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeOrderableReplicationInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeOrderableReplicationInstancesPages(input *databasemigrationservice.DescribeOrderableReplicationInstancesInput, fn func(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeOrderableReplicationInstancesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeOrderableReplicationInstancesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeOrderableReplicationInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeOrderableReplicationInstancesPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput, fn func(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeOrderableReplicationInstancesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeOrderableReplicationInstancesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeOrderableReplicationInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeOrderableReplicationInstancesWithContext(ctx context.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput, opts ...request.Option) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeOrderableReplicationInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribePendingMaintenanceActions(input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribePendingMaintenanceActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribePendingMaintenanceActionsPages(input *databasemigrationservice.DescribePendingMaintenanceActionsInput, fn func(*databasemigrationservice.DescribePendingMaintenanceActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribePendingMaintenanceActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribePendingMaintenanceActionsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput, fn func(*databasemigrationservice.DescribePendingMaintenanceActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribePendingMaintenanceActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribePendingMaintenanceActionsWithContext(ctx context.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput, opts ...request.Option) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribePendingMaintenanceActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeRefreshSchemasStatus(input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeRefreshSchemasStatusOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeRefreshSchemasStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeRefreshSchemasStatusWithContext(ctx context.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput, opts ...request.Option) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeRefreshSchemasStatusOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeRefreshSchemasStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationInstanceTaskLogs(input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstanceTaskLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationInstanceTaskLogsPages(input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, fn func(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstanceTaskLogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationInstanceTaskLogsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, fn func(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstanceTaskLogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationInstanceTaskLogsWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstanceTaskLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationInstances(input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationInstancesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationInstancesPages(input *databasemigrationservice.DescribeReplicationInstancesInput, fn func(*databasemigrationservice.DescribeReplicationInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationInstancesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationInstancesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationInstancesPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput, fn func(*databasemigrationservice.DescribeReplicationInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationInstancesOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationInstancesOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationInstancesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationInstancesOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationSubnetGroups(input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationSubnetGroupsPages(input *databasemigrationservice.DescribeReplicationSubnetGroupsInput, fn func(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationSubnetGroupsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationSubnetGroupsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationSubnetGroupsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput, fn func(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationSubnetGroupsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationSubnetGroupsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationSubnetGroupsWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentResults(input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentResultsPages(input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, fn func(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentResultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentResultsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, fn func(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentResultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentResultsWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentRuns(input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentRunsPages(input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, fn func(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentRunsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, fn func(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskAssessmentRunsWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskAssessmentRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskIndividualAssessments(input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskIndividualAssessments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTaskIndividualAssessmentsPages(input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, fn func(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskIndividualAssessmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskIndividualAssessmentsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, fn func(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskIndividualAssessmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTaskIndividualAssessmentsWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTaskIndividualAssessmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTasks(input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTasksOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeReplicationTasksPages(input *databasemigrationservice.DescribeReplicationTasksInput, fn func(*databasemigrationservice.DescribeReplicationTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTasksOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTasksOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTasksPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput, fn func(*databasemigrationservice.DescribeReplicationTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeReplicationTasksOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeReplicationTasksOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeReplicationTasksOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeReplicationTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeReplicationTasksWithContext(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput, opts ...request.Option) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeReplicationTasksOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeReplicationTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeSchemas(input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeSchemasOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeSchemasPages(input *databasemigrationservice.DescribeSchemasInput, fn func(*databasemigrationservice.DescribeSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeSchemasOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeSchemasOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeSchemasPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeSchemasInput, fn func(*databasemigrationservice.DescribeSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeSchemasOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeSchemasOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeSchemasWithContext(ctx context.Context, input *databasemigrationservice.DescribeSchemasInput, opts ...request.Option) (*databasemigrationservice.DescribeSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeSchemasOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeTableStatistics(input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeTableStatisticsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeTableStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) DescribeTableStatisticsPages(input *databasemigrationservice.DescribeTableStatisticsInput, fn func(*databasemigrationservice.DescribeTableStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeTableStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeTableStatisticsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeTableStatisticsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeTableStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeTableStatisticsPagesWithContext(ctx context.Context, input *databasemigrationservice.DescribeTableStatisticsInput, fn func(*databasemigrationservice.DescribeTableStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*databasemigrationservice.DescribeTableStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &databasemigrationservice.DescribeTableStatisticsOutput{}
	fnCacher := func(out *databasemigrationservice.DescribeTableStatisticsOutput, more bool) bool {
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
	if err := c.DatabaseMigrationServiceAPI.DescribeTableStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DatabaseMigrationService) DescribeTableStatisticsWithContext(ctx context.Context, input *databasemigrationservice.DescribeTableStatisticsInput, opts ...request.Option) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.DescribeTableStatisticsOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.DescribeTableStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) ListTagsForResource(input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.ListTagsForResourceOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DatabaseMigrationService) ListTagsForResourceWithContext(ctx context.Context, input *databasemigrationservice.ListTagsForResourceInput, opts ...request.Option) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*databasemigrationservice.ListTagsForResourceOutput), nil
	}
	out, err := c.DatabaseMigrationServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RDS struct {
	rdsiface.RDSAPI
	cache *cache.Cache
}

func NewRDS(rdsapi rdsiface.RDSAPI) *RDS {
	return &RDS{
		RDSAPI: rdsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RDS) DescribeAccountAttributes(input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeAccountAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeAccountAttributesWithContext(ctx context.Context, input *rds.DescribeAccountAttributesInput, opts ...request.Option) (*rds.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeAccountAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeBlueGreenDeployments(input *rds.DescribeBlueGreenDeploymentsInput) (*rds.DescribeBlueGreenDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeBlueGreenDeploymentsOutput), nil
	}
	out, err := c.RDSAPI.DescribeBlueGreenDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeBlueGreenDeploymentsPages(input *rds.DescribeBlueGreenDeploymentsInput, fn func(*rds.DescribeBlueGreenDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeBlueGreenDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeBlueGreenDeploymentsOutput{}
	fnCacher := func(out *rds.DescribeBlueGreenDeploymentsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeBlueGreenDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeBlueGreenDeploymentsPagesWithContext(ctx context.Context, input *rds.DescribeBlueGreenDeploymentsInput, fn func(*rds.DescribeBlueGreenDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeBlueGreenDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeBlueGreenDeploymentsOutput{}
	fnCacher := func(out *rds.DescribeBlueGreenDeploymentsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeBlueGreenDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeBlueGreenDeploymentsWithContext(ctx context.Context, input *rds.DescribeBlueGreenDeploymentsInput, opts ...request.Option) (*rds.DescribeBlueGreenDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeBlueGreenDeploymentsOutput), nil
	}
	out, err := c.RDSAPI.DescribeBlueGreenDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeCertificates(input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeCertificatesOutput), nil
	}
	out, err := c.RDSAPI.DescribeCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeCertificatesPages(input *rds.DescribeCertificatesInput, fn func(*rds.DescribeCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeCertificatesOutput{}
	fnCacher := func(out *rds.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeCertificatesPagesWithContext(ctx context.Context, input *rds.DescribeCertificatesInput, fn func(*rds.DescribeCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeCertificatesOutput{}
	fnCacher := func(out *rds.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeCertificatesWithContext(ctx context.Context, input *rds.DescribeCertificatesInput, opts ...request.Option) (*rds.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeCertificatesOutput), nil
	}
	out, err := c.RDSAPI.DescribeCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterBacktracks(input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterBacktracksOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterBacktracks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterBacktracksPages(input *rds.DescribeDBClusterBacktracksInput, fn func(*rds.DescribeDBClusterBacktracksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterBacktracksOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterBacktracksOutput{}
	fnCacher := func(out *rds.DescribeDBClusterBacktracksOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterBacktracksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterBacktracksPagesWithContext(ctx context.Context, input *rds.DescribeDBClusterBacktracksInput, fn func(*rds.DescribeDBClusterBacktracksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterBacktracksOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterBacktracksOutput{}
	fnCacher := func(out *rds.DescribeDBClusterBacktracksOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterBacktracksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterBacktracksWithContext(ctx context.Context, input *rds.DescribeDBClusterBacktracksInput, opts ...request.Option) (*rds.DescribeDBClusterBacktracksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterBacktracksOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterBacktracksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterEndpoints(input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterEndpointsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterEndpointsPages(input *rds.DescribeDBClusterEndpointsInput, fn func(*rds.DescribeDBClusterEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterEndpointsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterEndpointsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterEndpointsPagesWithContext(ctx context.Context, input *rds.DescribeDBClusterEndpointsInput, fn func(*rds.DescribeDBClusterEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterEndpointsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterEndpointsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterEndpointsWithContext(ctx context.Context, input *rds.DescribeDBClusterEndpointsInput, opts ...request.Option) (*rds.DescribeDBClusterEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterEndpointsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterParameterGroups(input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterParameterGroupsPages(input *rds.DescribeDBClusterParameterGroupsInput, fn func(*rds.DescribeDBClusterParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterParameterGroupsPagesWithContext(ctx context.Context, input *rds.DescribeDBClusterParameterGroupsInput, fn func(*rds.DescribeDBClusterParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterParameterGroupsWithContext(ctx context.Context, input *rds.DescribeDBClusterParameterGroupsInput, opts ...request.Option) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterParameters(input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterParametersPages(input *rds.DescribeDBClusterParametersInput, fn func(*rds.DescribeDBClusterParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *rds.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterParametersPagesWithContext(ctx context.Context, input *rds.DescribeDBClusterParametersInput, fn func(*rds.DescribeDBClusterParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *rds.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterParametersWithContext(ctx context.Context, input *rds.DescribeDBClusterParametersInput, opts ...request.Option) (*rds.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterSnapshotAttributes(input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterSnapshotAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterSnapshotAttributesWithContext(ctx context.Context, input *rds.DescribeDBClusterSnapshotAttributesInput, opts ...request.Option) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterSnapshotAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterSnapshots(input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusterSnapshotsPages(input *rds.DescribeDBClusterSnapshotsInput, fn func(*rds.DescribeDBClusterSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterSnapshotsPagesWithContext(ctx context.Context, input *rds.DescribeDBClusterSnapshotsInput, fn func(*rds.DescribeDBClusterSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *rds.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClusterSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClusterSnapshotsWithContext(ctx context.Context, input *rds.DescribeDBClusterSnapshotsInput, opts ...request.Option) (*rds.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClusters(input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClustersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBClustersPages(input *rds.DescribeDBClustersInput, fn func(*rds.DescribeDBClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClustersOutput{}
	fnCacher := func(out *rds.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClustersPagesWithContext(ctx context.Context, input *rds.DescribeDBClustersInput, fn func(*rds.DescribeDBClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBClustersOutput{}
	fnCacher := func(out *rds.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBClustersWithContext(ctx context.Context, input *rds.DescribeDBClustersInput, opts ...request.Option) (*rds.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBClustersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBEngineVersions(input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBEngineVersionsPages(input *rds.DescribeDBEngineVersionsInput, fn func(*rds.DescribeDBEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *rds.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBEngineVersionsPagesWithContext(ctx context.Context, input *rds.DescribeDBEngineVersionsInput, fn func(*rds.DescribeDBEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *rds.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBEngineVersionsWithContext(ctx context.Context, input *rds.DescribeDBEngineVersionsInput, opts ...request.Option) (*rds.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBInstanceAutomatedBackups(input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBInstanceAutomatedBackupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBInstanceAutomatedBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBInstanceAutomatedBackupsPages(input *rds.DescribeDBInstanceAutomatedBackupsInput, fn func(*rds.DescribeDBInstanceAutomatedBackupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBInstanceAutomatedBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBInstanceAutomatedBackupsOutput{}
	fnCacher := func(out *rds.DescribeDBInstanceAutomatedBackupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBInstanceAutomatedBackupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBInstanceAutomatedBackupsPagesWithContext(ctx context.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput, fn func(*rds.DescribeDBInstanceAutomatedBackupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBInstanceAutomatedBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBInstanceAutomatedBackupsOutput{}
	fnCacher := func(out *rds.DescribeDBInstanceAutomatedBackupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBInstanceAutomatedBackupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBInstanceAutomatedBackupsWithContext(ctx context.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput, opts ...request.Option) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBInstanceAutomatedBackupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBInstanceAutomatedBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBInstances(input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBInstancesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBInstancesPages(input *rds.DescribeDBInstancesInput, fn func(*rds.DescribeDBInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBInstancesOutput{}
	fnCacher := func(out *rds.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBInstancesPagesWithContext(ctx context.Context, input *rds.DescribeDBInstancesInput, fn func(*rds.DescribeDBInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBInstancesOutput{}
	fnCacher := func(out *rds.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBInstancesWithContext(ctx context.Context, input *rds.DescribeDBInstancesInput, opts ...request.Option) (*rds.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBInstancesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBLogFiles(input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBLogFilesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBLogFiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBLogFilesPages(input *rds.DescribeDBLogFilesInput, fn func(*rds.DescribeDBLogFilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBLogFilesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBLogFilesOutput{}
	fnCacher := func(out *rds.DescribeDBLogFilesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBLogFilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBLogFilesPagesWithContext(ctx context.Context, input *rds.DescribeDBLogFilesInput, fn func(*rds.DescribeDBLogFilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBLogFilesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBLogFilesOutput{}
	fnCacher := func(out *rds.DescribeDBLogFilesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBLogFilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBLogFilesWithContext(ctx context.Context, input *rds.DescribeDBLogFilesInput, opts ...request.Option) (*rds.DescribeDBLogFilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBLogFilesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBLogFilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBParameterGroups(input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBParameterGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBParameterGroupsPages(input *rds.DescribeDBParameterGroupsInput, fn func(*rds.DescribeDBParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBParameterGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBParameterGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBParameterGroupsPagesWithContext(ctx context.Context, input *rds.DescribeDBParameterGroupsInput, fn func(*rds.DescribeDBParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBParameterGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBParameterGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBParameterGroupsWithContext(ctx context.Context, input *rds.DescribeDBParameterGroupsInput, opts ...request.Option) (*rds.DescribeDBParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBParameterGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBParameters(input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBParametersPages(input *rds.DescribeDBParametersInput, fn func(*rds.DescribeDBParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBParametersOutput{}
	fnCacher := func(out *rds.DescribeDBParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBParametersPagesWithContext(ctx context.Context, input *rds.DescribeDBParametersInput, fn func(*rds.DescribeDBParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBParametersOutput{}
	fnCacher := func(out *rds.DescribeDBParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBParametersWithContext(ctx context.Context, input *rds.DescribeDBParametersInput, opts ...request.Option) (*rds.DescribeDBParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxies(input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxiesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxiesPages(input *rds.DescribeDBProxiesInput, fn func(*rds.DescribeDBProxiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxiesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxiesOutput{}
	fnCacher := func(out *rds.DescribeDBProxiesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxiesPagesWithContext(ctx context.Context, input *rds.DescribeDBProxiesInput, fn func(*rds.DescribeDBProxiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxiesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxiesOutput{}
	fnCacher := func(out *rds.DescribeDBProxiesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxiesWithContext(ctx context.Context, input *rds.DescribeDBProxiesInput, opts ...request.Option) (*rds.DescribeDBProxiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxiesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyEndpoints(input *rds.DescribeDBProxyEndpointsInput) (*rds.DescribeDBProxyEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyEndpointsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyEndpointsPages(input *rds.DescribeDBProxyEndpointsInput, fn func(*rds.DescribeDBProxyEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyEndpointsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyEndpointsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyEndpointsPagesWithContext(ctx context.Context, input *rds.DescribeDBProxyEndpointsInput, fn func(*rds.DescribeDBProxyEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyEndpointsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyEndpointsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyEndpointsWithContext(ctx context.Context, input *rds.DescribeDBProxyEndpointsInput, opts ...request.Option) (*rds.DescribeDBProxyEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyEndpointsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyTargetGroups(input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyTargetGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyTargetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyTargetGroupsPages(input *rds.DescribeDBProxyTargetGroupsInput, fn func(*rds.DescribeDBProxyTargetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyTargetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyTargetGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyTargetGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyTargetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyTargetGroupsPagesWithContext(ctx context.Context, input *rds.DescribeDBProxyTargetGroupsInput, fn func(*rds.DescribeDBProxyTargetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyTargetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyTargetGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyTargetGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyTargetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyTargetGroupsWithContext(ctx context.Context, input *rds.DescribeDBProxyTargetGroupsInput, opts ...request.Option) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyTargetGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyTargetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyTargets(input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyTargetsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBProxyTargetsPages(input *rds.DescribeDBProxyTargetsInput, fn func(*rds.DescribeDBProxyTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyTargetsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyTargetsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyTargetsPagesWithContext(ctx context.Context, input *rds.DescribeDBProxyTargetsInput, fn func(*rds.DescribeDBProxyTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBProxyTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBProxyTargetsOutput{}
	fnCacher := func(out *rds.DescribeDBProxyTargetsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBProxyTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBProxyTargetsWithContext(ctx context.Context, input *rds.DescribeDBProxyTargetsInput, opts ...request.Option) (*rds.DescribeDBProxyTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBProxyTargetsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBProxyTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSecurityGroups(input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSecurityGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSecurityGroupsPages(input *rds.DescribeDBSecurityGroupsInput, fn func(*rds.DescribeDBSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSecurityGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBSecurityGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSecurityGroupsPagesWithContext(ctx context.Context, input *rds.DescribeDBSecurityGroupsInput, fn func(*rds.DescribeDBSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSecurityGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBSecurityGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSecurityGroupsWithContext(ctx context.Context, input *rds.DescribeDBSecurityGroupsInput, opts ...request.Option) (*rds.DescribeDBSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSecurityGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSnapshotAttributes(input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSnapshotAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSnapshotAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSnapshotAttributesWithContext(ctx context.Context, input *rds.DescribeDBSnapshotAttributesInput, opts ...request.Option) (*rds.DescribeDBSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSnapshotAttributesOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSnapshotAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSnapshots(input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSnapshotsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSnapshotsPages(input *rds.DescribeDBSnapshotsInput, fn func(*rds.DescribeDBSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSnapshotsOutput{}
	fnCacher := func(out *rds.DescribeDBSnapshotsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSnapshotsPagesWithContext(ctx context.Context, input *rds.DescribeDBSnapshotsInput, fn func(*rds.DescribeDBSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSnapshotsOutput{}
	fnCacher := func(out *rds.DescribeDBSnapshotsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSnapshotsWithContext(ctx context.Context, input *rds.DescribeDBSnapshotsInput, opts ...request.Option) (*rds.DescribeDBSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSnapshotsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSubnetGroups(input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeDBSubnetGroupsPages(input *rds.DescribeDBSubnetGroupsInput, fn func(*rds.DescribeDBSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSubnetGroupsPagesWithContext(ctx context.Context, input *rds.DescribeDBSubnetGroupsInput, fn func(*rds.DescribeDBSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *rds.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeDBSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeDBSubnetGroupsWithContext(ctx context.Context, input *rds.DescribeDBSubnetGroupsInput, opts ...request.Option) (*rds.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeDBSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEngineDefaultClusterParameters(input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeEngineDefaultClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEngineDefaultClusterParametersWithContext(ctx context.Context, input *rds.DescribeEngineDefaultClusterParametersInput, opts ...request.Option) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeEngineDefaultClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEngineDefaultParameters(input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeEngineDefaultParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEngineDefaultParametersPages(input *rds.DescribeEngineDefaultParametersInput, fn func(*rds.DescribeEngineDefaultParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *rds.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEngineDefaultParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEngineDefaultParametersPagesWithContext(ctx context.Context, input *rds.DescribeEngineDefaultParametersInput, fn func(*rds.DescribeEngineDefaultParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *rds.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEngineDefaultParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEngineDefaultParametersWithContext(ctx context.Context, input *rds.DescribeEngineDefaultParametersInput, opts ...request.Option) (*rds.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.RDSAPI.DescribeEngineDefaultParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEventCategories(input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventCategoriesOutput), nil
	}
	out, err := c.RDSAPI.DescribeEventCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEventCategoriesWithContext(ctx context.Context, input *rds.DescribeEventCategoriesInput, opts ...request.Option) (*rds.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventCategoriesOutput), nil
	}
	out, err := c.RDSAPI.DescribeEventCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEventSubscriptions(input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEventSubscriptionsPages(input *rds.DescribeEventSubscriptionsInput, fn func(*rds.DescribeEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *rds.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEventSubscriptionsPagesWithContext(ctx context.Context, input *rds.DescribeEventSubscriptionsInput, fn func(*rds.DescribeEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *rds.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEventSubscriptionsWithContext(ctx context.Context, input *rds.DescribeEventSubscriptionsInput, opts ...request.Option) (*rds.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEvents(input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventsOutput), nil
	}
	out, err := c.RDSAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeEventsPages(input *rds.DescribeEventsInput, fn func(*rds.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEventsOutput{}
	fnCacher := func(out *rds.DescribeEventsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEventsPagesWithContext(ctx context.Context, input *rds.DescribeEventsInput, fn func(*rds.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeEventsOutput{}
	fnCacher := func(out *rds.DescribeEventsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeEventsWithContext(ctx context.Context, input *rds.DescribeEventsInput, opts ...request.Option) (*rds.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeEventsOutput), nil
	}
	out, err := c.RDSAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeExportTasks(input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeExportTasksOutput), nil
	}
	out, err := c.RDSAPI.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeExportTasksPages(input *rds.DescribeExportTasksInput, fn func(*rds.DescribeExportTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeExportTasksOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeExportTasksOutput{}
	fnCacher := func(out *rds.DescribeExportTasksOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeExportTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeExportTasksPagesWithContext(ctx context.Context, input *rds.DescribeExportTasksInput, fn func(*rds.DescribeExportTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeExportTasksOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeExportTasksOutput{}
	fnCacher := func(out *rds.DescribeExportTasksOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeExportTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeExportTasksWithContext(ctx context.Context, input *rds.DescribeExportTasksInput, opts ...request.Option) (*rds.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeExportTasksOutput), nil
	}
	out, err := c.RDSAPI.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeGlobalClusters(input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeGlobalClustersOutput), nil
	}
	out, err := c.RDSAPI.DescribeGlobalClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeGlobalClustersPages(input *rds.DescribeGlobalClustersInput, fn func(*rds.DescribeGlobalClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeGlobalClustersOutput{}
	fnCacher := func(out *rds.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeGlobalClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeGlobalClustersPagesWithContext(ctx context.Context, input *rds.DescribeGlobalClustersInput, fn func(*rds.DescribeGlobalClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeGlobalClustersOutput{}
	fnCacher := func(out *rds.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeGlobalClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeGlobalClustersWithContext(ctx context.Context, input *rds.DescribeGlobalClustersInput, opts ...request.Option) (*rds.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeGlobalClustersOutput), nil
	}
	out, err := c.RDSAPI.DescribeGlobalClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOptionGroupOptions(input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOptionGroupOptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOptionGroupOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOptionGroupOptionsPages(input *rds.DescribeOptionGroupOptionsInput, fn func(*rds.DescribeOptionGroupOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOptionGroupOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOptionGroupOptionsOutput{}
	fnCacher := func(out *rds.DescribeOptionGroupOptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOptionGroupOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOptionGroupOptionsPagesWithContext(ctx context.Context, input *rds.DescribeOptionGroupOptionsInput, fn func(*rds.DescribeOptionGroupOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOptionGroupOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOptionGroupOptionsOutput{}
	fnCacher := func(out *rds.DescribeOptionGroupOptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOptionGroupOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOptionGroupOptionsWithContext(ctx context.Context, input *rds.DescribeOptionGroupOptionsInput, opts ...request.Option) (*rds.DescribeOptionGroupOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOptionGroupOptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOptionGroupOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOptionGroups(input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOptionGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOptionGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOptionGroupsPages(input *rds.DescribeOptionGroupsInput, fn func(*rds.DescribeOptionGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOptionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOptionGroupsOutput{}
	fnCacher := func(out *rds.DescribeOptionGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOptionGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOptionGroupsPagesWithContext(ctx context.Context, input *rds.DescribeOptionGroupsInput, fn func(*rds.DescribeOptionGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOptionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOptionGroupsOutput{}
	fnCacher := func(out *rds.DescribeOptionGroupsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOptionGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOptionGroupsWithContext(ctx context.Context, input *rds.DescribeOptionGroupsInput, opts ...request.Option) (*rds.DescribeOptionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOptionGroupsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOptionGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOrderableDBInstanceOptions(input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOrderableDBInstanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeOrderableDBInstanceOptionsPages(input *rds.DescribeOrderableDBInstanceOptionsInput, fn func(*rds.DescribeOrderableDBInstanceOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *rds.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOrderableDBInstanceOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOrderableDBInstanceOptionsPagesWithContext(ctx context.Context, input *rds.DescribeOrderableDBInstanceOptionsInput, fn func(*rds.DescribeOrderableDBInstanceOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *rds.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeOrderableDBInstanceOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeOrderableDBInstanceOptionsWithContext(ctx context.Context, input *rds.DescribeOrderableDBInstanceOptionsInput, opts ...request.Option) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeOrderableDBInstanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribePendingMaintenanceActions(input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.RDSAPI.DescribePendingMaintenanceActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribePendingMaintenanceActionsPages(input *rds.DescribePendingMaintenanceActionsInput, fn func(*rds.DescribePendingMaintenanceActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *rds.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribePendingMaintenanceActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribePendingMaintenanceActionsPagesWithContext(ctx context.Context, input *rds.DescribePendingMaintenanceActionsInput, fn func(*rds.DescribePendingMaintenanceActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *rds.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribePendingMaintenanceActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribePendingMaintenanceActionsWithContext(ctx context.Context, input *rds.DescribePendingMaintenanceActionsInput, opts ...request.Option) (*rds.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.RDSAPI.DescribePendingMaintenanceActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeReservedDBInstances(input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeReservedDBInstancesOutput), nil
	}
	out, err := c.RDSAPI.DescribeReservedDBInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeReservedDBInstancesOfferings(input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeReservedDBInstancesOfferingsOutput), nil
	}
	out, err := c.RDSAPI.DescribeReservedDBInstancesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeReservedDBInstancesOfferingsPages(input *rds.DescribeReservedDBInstancesOfferingsInput, fn func(*rds.DescribeReservedDBInstancesOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeReservedDBInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeReservedDBInstancesOfferingsOutput{}
	fnCacher := func(out *rds.DescribeReservedDBInstancesOfferingsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeReservedDBInstancesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeReservedDBInstancesOfferingsPagesWithContext(ctx context.Context, input *rds.DescribeReservedDBInstancesOfferingsInput, fn func(*rds.DescribeReservedDBInstancesOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeReservedDBInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeReservedDBInstancesOfferingsOutput{}
	fnCacher := func(out *rds.DescribeReservedDBInstancesOfferingsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeReservedDBInstancesOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeReservedDBInstancesOfferingsWithContext(ctx context.Context, input *rds.DescribeReservedDBInstancesOfferingsInput, opts ...request.Option) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeReservedDBInstancesOfferingsOutput), nil
	}
	out, err := c.RDSAPI.DescribeReservedDBInstancesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeReservedDBInstancesPages(input *rds.DescribeReservedDBInstancesInput, fn func(*rds.DescribeReservedDBInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeReservedDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeReservedDBInstancesOutput{}
	fnCacher := func(out *rds.DescribeReservedDBInstancesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeReservedDBInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeReservedDBInstancesPagesWithContext(ctx context.Context, input *rds.DescribeReservedDBInstancesInput, fn func(*rds.DescribeReservedDBInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeReservedDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeReservedDBInstancesOutput{}
	fnCacher := func(out *rds.DescribeReservedDBInstancesOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeReservedDBInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeReservedDBInstancesWithContext(ctx context.Context, input *rds.DescribeReservedDBInstancesInput, opts ...request.Option) (*rds.DescribeReservedDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeReservedDBInstancesOutput), nil
	}
	out, err := c.RDSAPI.DescribeReservedDBInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeSourceRegions(input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeSourceRegionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeSourceRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeSourceRegionsPages(input *rds.DescribeSourceRegionsInput, fn func(*rds.DescribeSourceRegionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeSourceRegionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeSourceRegionsOutput{}
	fnCacher := func(out *rds.DescribeSourceRegionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeSourceRegionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeSourceRegionsPagesWithContext(ctx context.Context, input *rds.DescribeSourceRegionsInput, fn func(*rds.DescribeSourceRegionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*rds.DescribeSourceRegionsOutput), false)
		return nil
	}
	cachable := true
	output := &rds.DescribeSourceRegionsOutput{}
	fnCacher := func(out *rds.DescribeSourceRegionsOutput, more bool) bool {
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
	if err := c.RDSAPI.DescribeSourceRegionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RDS) DescribeSourceRegionsWithContext(ctx context.Context, input *rds.DescribeSourceRegionsInput, opts ...request.Option) (*rds.DescribeSourceRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeSourceRegionsOutput), nil
	}
	out, err := c.RDSAPI.DescribeSourceRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeValidDBInstanceModifications(input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeValidDBInstanceModificationsOutput), nil
	}
	out, err := c.RDSAPI.DescribeValidDBInstanceModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) DescribeValidDBInstanceModificationsWithContext(ctx context.Context, input *rds.DescribeValidDBInstanceModificationsInput, opts ...request.Option) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.DescribeValidDBInstanceModificationsOutput), nil
	}
	out, err := c.RDSAPI.DescribeValidDBInstanceModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) ListTagsForResource(input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.ListTagsForResourceOutput), nil
	}
	out, err := c.RDSAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RDS) ListTagsForResourceWithContext(ctx context.Context, input *rds.ListTagsForResourceInput, opts ...request.Option) (*rds.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*rds.ListTagsForResourceOutput), nil
	}
	out, err := c.RDSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

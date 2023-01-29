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
package docdbcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DocDB struct {
	docdbiface.DocDBAPI
	cache *cache.Cache
}

func New(docdbapi docdbiface.DocDBAPI) *DocDB {
	return &DocDB{
		DocDBAPI: docdbapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DocDB) DescribeCertificates(input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeCertificatesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeCertificatesPages(input *docdb.DescribeCertificatesInput, fn func(*docdb.DescribeCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeCertificatesOutput{}
	fnCacher := func(out *docdb.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeCertificatesPagesWithContext(ctx context.Context, input *docdb.DescribeCertificatesInput, fn func(*docdb.DescribeCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeCertificatesOutput{}
	fnCacher := func(out *docdb.DescribeCertificatesOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeCertificatesWithContext(ctx context.Context, input *docdb.DescribeCertificatesInput, opts ...request.Option) (*docdb.DescribeCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeCertificatesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterParameterGroups(input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterParameterGroupsPages(input *docdb.DescribeDBClusterParameterGroupsInput, fn func(*docdb.DescribeDBClusterParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterParameterGroupsPagesWithContext(ctx context.Context, input *docdb.DescribeDBClusterParameterGroupsInput, fn func(*docdb.DescribeDBClusterParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterParameterGroupsWithContext(ctx context.Context, input *docdb.DescribeDBClusterParameterGroupsInput, opts ...request.Option) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterParameters(input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterParametersPages(input *docdb.DescribeDBClusterParametersInput, fn func(*docdb.DescribeDBClusterParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterParametersPagesWithContext(ctx context.Context, input *docdb.DescribeDBClusterParametersInput, fn func(*docdb.DescribeDBClusterParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterParametersWithContext(ctx context.Context, input *docdb.DescribeDBClusterParametersInput, opts ...request.Option) (*docdb.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterSnapshotAttributes(input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterSnapshotAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterSnapshotAttributesWithContext(ctx context.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput, opts ...request.Option) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterSnapshotAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterSnapshots(input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusterSnapshotsPages(input *docdb.DescribeDBClusterSnapshotsInput, fn func(*docdb.DescribeDBClusterSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterSnapshotsPagesWithContext(ctx context.Context, input *docdb.DescribeDBClusterSnapshotsInput, fn func(*docdb.DescribeDBClusterSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *docdb.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClusterSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClusterSnapshotsWithContext(ctx context.Context, input *docdb.DescribeDBClusterSnapshotsInput, opts ...request.Option) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClusters(input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClustersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBClustersPages(input *docdb.DescribeDBClustersInput, fn func(*docdb.DescribeDBClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClustersOutput{}
	fnCacher := func(out *docdb.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClustersPagesWithContext(ctx context.Context, input *docdb.DescribeDBClustersInput, fn func(*docdb.DescribeDBClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBClustersOutput{}
	fnCacher := func(out *docdb.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBClustersWithContext(ctx context.Context, input *docdb.DescribeDBClustersInput, opts ...request.Option) (*docdb.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBClustersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBEngineVersions(input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBEngineVersionsPages(input *docdb.DescribeDBEngineVersionsInput, fn func(*docdb.DescribeDBEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *docdb.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBEngineVersionsPagesWithContext(ctx context.Context, input *docdb.DescribeDBEngineVersionsInput, fn func(*docdb.DescribeDBEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *docdb.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBEngineVersionsWithContext(ctx context.Context, input *docdb.DescribeDBEngineVersionsInput, opts ...request.Option) (*docdb.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBInstances(input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBInstancesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBInstancesPages(input *docdb.DescribeDBInstancesInput, fn func(*docdb.DescribeDBInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBInstancesOutput{}
	fnCacher := func(out *docdb.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBInstancesPagesWithContext(ctx context.Context, input *docdb.DescribeDBInstancesInput, fn func(*docdb.DescribeDBInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBInstancesOutput{}
	fnCacher := func(out *docdb.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBInstancesWithContext(ctx context.Context, input *docdb.DescribeDBInstancesInput, opts ...request.Option) (*docdb.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBInstancesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBSubnetGroups(input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeDBSubnetGroupsPages(input *docdb.DescribeDBSubnetGroupsInput, fn func(*docdb.DescribeDBSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *docdb.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBSubnetGroupsPagesWithContext(ctx context.Context, input *docdb.DescribeDBSubnetGroupsInput, fn func(*docdb.DescribeDBSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *docdb.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeDBSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeDBSubnetGroupsWithContext(ctx context.Context, input *docdb.DescribeDBSubnetGroupsInput, opts ...request.Option) (*docdb.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeDBSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEngineDefaultClusterParameters(input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEngineDefaultClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEngineDefaultClusterParametersWithContext(ctx context.Context, input *docdb.DescribeEngineDefaultClusterParametersInput, opts ...request.Option) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEngineDefaultClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEventCategories(input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventCategoriesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEventCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEventCategoriesWithContext(ctx context.Context, input *docdb.DescribeEventCategoriesInput, opts ...request.Option) (*docdb.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventCategoriesOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEventCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEventSubscriptions(input *docdb.DescribeEventSubscriptionsInput) (*docdb.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEventSubscriptionsPages(input *docdb.DescribeEventSubscriptionsInput, fn func(*docdb.DescribeEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *docdb.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeEventSubscriptionsPagesWithContext(ctx context.Context, input *docdb.DescribeEventSubscriptionsInput, fn func(*docdb.DescribeEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *docdb.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeEventSubscriptionsWithContext(ctx context.Context, input *docdb.DescribeEventSubscriptionsInput, opts ...request.Option) (*docdb.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEvents(input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeEventsPages(input *docdb.DescribeEventsInput, fn func(*docdb.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeEventsOutput{}
	fnCacher := func(out *docdb.DescribeEventsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeEventsPagesWithContext(ctx context.Context, input *docdb.DescribeEventsInput, fn func(*docdb.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeEventsOutput{}
	fnCacher := func(out *docdb.DescribeEventsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeEventsWithContext(ctx context.Context, input *docdb.DescribeEventsInput, opts ...request.Option) (*docdb.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeEventsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeGlobalClusters(input *docdb.DescribeGlobalClustersInput) (*docdb.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeGlobalClustersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeGlobalClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeGlobalClustersPages(input *docdb.DescribeGlobalClustersInput, fn func(*docdb.DescribeGlobalClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeGlobalClustersOutput{}
	fnCacher := func(out *docdb.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeGlobalClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeGlobalClustersPagesWithContext(ctx context.Context, input *docdb.DescribeGlobalClustersInput, fn func(*docdb.DescribeGlobalClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeGlobalClustersOutput{}
	fnCacher := func(out *docdb.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeGlobalClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeGlobalClustersWithContext(ctx context.Context, input *docdb.DescribeGlobalClustersInput, opts ...request.Option) (*docdb.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeGlobalClustersOutput), nil
	}
	out, err := c.DocDBAPI.DescribeGlobalClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeOrderableDBInstanceOptions(input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeOrderableDBInstanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribeOrderableDBInstanceOptionsPages(input *docdb.DescribeOrderableDBInstanceOptionsInput, fn func(*docdb.DescribeOrderableDBInstanceOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *docdb.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeOrderableDBInstanceOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeOrderableDBInstanceOptionsPagesWithContext(ctx context.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput, fn func(*docdb.DescribeOrderableDBInstanceOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *docdb.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribeOrderableDBInstanceOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribeOrderableDBInstanceOptionsWithContext(ctx context.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput, opts ...request.Option) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribeOrderableDBInstanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribePendingMaintenanceActions(input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribePendingMaintenanceActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) DescribePendingMaintenanceActionsPages(input *docdb.DescribePendingMaintenanceActionsInput, fn func(*docdb.DescribePendingMaintenanceActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *docdb.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribePendingMaintenanceActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribePendingMaintenanceActionsPagesWithContext(ctx context.Context, input *docdb.DescribePendingMaintenanceActionsInput, fn func(*docdb.DescribePendingMaintenanceActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdb.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &docdb.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *docdb.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.DocDBAPI.DescribePendingMaintenanceActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDB) DescribePendingMaintenanceActionsWithContext(ctx context.Context, input *docdb.DescribePendingMaintenanceActionsInput, opts ...request.Option) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.DocDBAPI.DescribePendingMaintenanceActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) ListTagsForResource(input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.ListTagsForResourceOutput), nil
	}
	out, err := c.DocDBAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDB) ListTagsForResourceWithContext(ctx context.Context, input *docdb.ListTagsForResourceInput, opts ...request.Option) (*docdb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdb.ListTagsForResourceOutput), nil
	}
	out, err := c.DocDBAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

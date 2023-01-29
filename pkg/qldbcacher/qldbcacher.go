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
package qldbcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type QLDB struct {
	qldbiface.QLDBAPI
	cache *cache.Cache
}

func New(qldbapi qldbiface.QLDBAPI) *QLDB {
	return &QLDB{
		QLDBAPI: qldbapi,
		cache:   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *QLDB) DescribeJournalKinesisStream(input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeJournalKinesisStreamOutput), nil
	}
	out, err := c.QLDBAPI.DescribeJournalKinesisStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) DescribeJournalKinesisStreamWithContext(ctx context.Context, input *qldb.DescribeJournalKinesisStreamInput, opts ...request.Option) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeJournalKinesisStreamOutput), nil
	}
	out, err := c.QLDBAPI.DescribeJournalKinesisStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) DescribeJournalS3Export(input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeJournalS3ExportOutput), nil
	}
	out, err := c.QLDBAPI.DescribeJournalS3Export(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) DescribeJournalS3ExportWithContext(ctx context.Context, input *qldb.DescribeJournalS3ExportInput, opts ...request.Option) (*qldb.DescribeJournalS3ExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeJournalS3ExportOutput), nil
	}
	out, err := c.QLDBAPI.DescribeJournalS3ExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) DescribeLedger(input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeLedgerOutput), nil
	}
	out, err := c.QLDBAPI.DescribeLedger(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) DescribeLedgerWithContext(ctx context.Context, input *qldb.DescribeLedgerInput, opts ...request.Option) (*qldb.DescribeLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.DescribeLedgerOutput), nil
	}
	out, err := c.QLDBAPI.DescribeLedgerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetBlock(input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetBlockOutput), nil
	}
	out, err := c.QLDBAPI.GetBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetBlockWithContext(ctx context.Context, input *qldb.GetBlockInput, opts ...request.Option) (*qldb.GetBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetBlockOutput), nil
	}
	out, err := c.QLDBAPI.GetBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetDigest(input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetDigestOutput), nil
	}
	out, err := c.QLDBAPI.GetDigest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetDigestWithContext(ctx context.Context, input *qldb.GetDigestInput, opts ...request.Option) (*qldb.GetDigestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetDigestOutput), nil
	}
	out, err := c.QLDBAPI.GetDigestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetRevision(input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetRevisionOutput), nil
	}
	out, err := c.QLDBAPI.GetRevision(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) GetRevisionWithContext(ctx context.Context, input *qldb.GetRevisionInput, opts ...request.Option) (*qldb.GetRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.GetRevisionOutput), nil
	}
	out, err := c.QLDBAPI.GetRevisionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalKinesisStreamsForLedger(input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalKinesisStreamsForLedgerOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalKinesisStreamsForLedger(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalKinesisStreamsForLedgerPages(input *qldb.ListJournalKinesisStreamsForLedgerInput, fn func(*qldb.ListJournalKinesisStreamsForLedgerOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalKinesisStreamsForLedgerOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalKinesisStreamsForLedgerOutput{}
	fnCacher := func(out *qldb.ListJournalKinesisStreamsForLedgerOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalKinesisStreamsForLedgerPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalKinesisStreamsForLedgerPagesWithContext(ctx context.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput, fn func(*qldb.ListJournalKinesisStreamsForLedgerOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalKinesisStreamsForLedgerOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalKinesisStreamsForLedgerOutput{}
	fnCacher := func(out *qldb.ListJournalKinesisStreamsForLedgerOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalKinesisStreamsForLedgerPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalKinesisStreamsForLedgerWithContext(ctx context.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput, opts ...request.Option) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalKinesisStreamsForLedgerOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalKinesisStreamsForLedgerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalS3Exports(input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalS3ExportsOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalS3Exports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalS3ExportsForLedger(input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalS3ExportsForLedgerOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalS3ExportsForLedger(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalS3ExportsForLedgerPages(input *qldb.ListJournalS3ExportsForLedgerInput, fn func(*qldb.ListJournalS3ExportsForLedgerOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalS3ExportsForLedgerOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalS3ExportsForLedgerOutput{}
	fnCacher := func(out *qldb.ListJournalS3ExportsForLedgerOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalS3ExportsForLedgerPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalS3ExportsForLedgerPagesWithContext(ctx context.Context, input *qldb.ListJournalS3ExportsForLedgerInput, fn func(*qldb.ListJournalS3ExportsForLedgerOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalS3ExportsForLedgerOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalS3ExportsForLedgerOutput{}
	fnCacher := func(out *qldb.ListJournalS3ExportsForLedgerOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalS3ExportsForLedgerPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalS3ExportsForLedgerWithContext(ctx context.Context, input *qldb.ListJournalS3ExportsForLedgerInput, opts ...request.Option) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalS3ExportsForLedgerOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalS3ExportsForLedgerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListJournalS3ExportsPages(input *qldb.ListJournalS3ExportsInput, fn func(*qldb.ListJournalS3ExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalS3ExportsOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalS3ExportsOutput{}
	fnCacher := func(out *qldb.ListJournalS3ExportsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalS3ExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalS3ExportsPagesWithContext(ctx context.Context, input *qldb.ListJournalS3ExportsInput, fn func(*qldb.ListJournalS3ExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListJournalS3ExportsOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListJournalS3ExportsOutput{}
	fnCacher := func(out *qldb.ListJournalS3ExportsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListJournalS3ExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListJournalS3ExportsWithContext(ctx context.Context, input *qldb.ListJournalS3ExportsInput, opts ...request.Option) (*qldb.ListJournalS3ExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListJournalS3ExportsOutput), nil
	}
	out, err := c.QLDBAPI.ListJournalS3ExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListLedgers(input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListLedgersOutput), nil
	}
	out, err := c.QLDBAPI.ListLedgers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListLedgersPages(input *qldb.ListLedgersInput, fn func(*qldb.ListLedgersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListLedgersOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListLedgersOutput{}
	fnCacher := func(out *qldb.ListLedgersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListLedgersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListLedgersPagesWithContext(ctx context.Context, input *qldb.ListLedgersInput, fn func(*qldb.ListLedgersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*qldb.ListLedgersOutput), false)
		return nil
	}
	cachable := true
	output := &qldb.ListLedgersOutput{}
	fnCacher := func(out *qldb.ListLedgersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.QLDBAPI.ListLedgersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QLDB) ListLedgersWithContext(ctx context.Context, input *qldb.ListLedgersInput, opts ...request.Option) (*qldb.ListLedgersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListLedgersOutput), nil
	}
	out, err := c.QLDBAPI.ListLedgersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListTagsForResource(input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListTagsForResourceOutput), nil
	}
	out, err := c.QLDBAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QLDB) ListTagsForResourceWithContext(ctx context.Context, input *qldb.ListTagsForResourceInput, opts ...request.Option) (*qldb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*qldb.ListTagsForResourceOutput), nil
	}
	out, err := c.QLDBAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

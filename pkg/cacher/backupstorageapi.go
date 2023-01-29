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
	"github.com/aws/aws-sdk-go/service/backupstorage"
	"github.com/aws/aws-sdk-go/service/backupstorage/backupstorageiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type BackupStorage struct {
	backupstorageiface.BackupStorageAPI
	cache *cache.Cache
}

func NewBackupStorage(backupstorageapi backupstorageiface.BackupStorageAPI) *BackupStorage {
	return &BackupStorage{
		BackupStorageAPI: backupstorageapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *BackupStorage) GetChunk(input *backupstorage.GetChunkInput) (*backupstorage.GetChunkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.GetChunkOutput), nil
	}
	out, err := c.BackupStorageAPI.GetChunk(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) GetChunkWithContext(ctx context.Context, input *backupstorage.GetChunkInput, opts ...request.Option) (*backupstorage.GetChunkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.GetChunkOutput), nil
	}
	out, err := c.BackupStorageAPI.GetChunkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) GetObjectMetadata(input *backupstorage.GetObjectMetadataInput) (*backupstorage.GetObjectMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.GetObjectMetadataOutput), nil
	}
	out, err := c.BackupStorageAPI.GetObjectMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) GetObjectMetadataWithContext(ctx context.Context, input *backupstorage.GetObjectMetadataInput, opts ...request.Option) (*backupstorage.GetObjectMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.GetObjectMetadataOutput), nil
	}
	out, err := c.BackupStorageAPI.GetObjectMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) ListChunks(input *backupstorage.ListChunksInput) (*backupstorage.ListChunksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.ListChunksOutput), nil
	}
	out, err := c.BackupStorageAPI.ListChunks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) ListChunksPages(input *backupstorage.ListChunksInput, fn func(*backupstorage.ListChunksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupstorage.ListChunksOutput), false)
		return nil
	}
	cachable := true
	output := &backupstorage.ListChunksOutput{}
	fnCacher := func(out *backupstorage.ListChunksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupStorageAPI.ListChunksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupStorage) ListChunksPagesWithContext(ctx context.Context, input *backupstorage.ListChunksInput, fn func(*backupstorage.ListChunksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupstorage.ListChunksOutput), false)
		return nil
	}
	cachable := true
	output := &backupstorage.ListChunksOutput{}
	fnCacher := func(out *backupstorage.ListChunksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupStorageAPI.ListChunksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupStorage) ListChunksWithContext(ctx context.Context, input *backupstorage.ListChunksInput, opts ...request.Option) (*backupstorage.ListChunksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.ListChunksOutput), nil
	}
	out, err := c.BackupStorageAPI.ListChunksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) ListObjects(input *backupstorage.ListObjectsInput) (*backupstorage.ListObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.ListObjectsOutput), nil
	}
	out, err := c.BackupStorageAPI.ListObjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupStorage) ListObjectsPages(input *backupstorage.ListObjectsInput, fn func(*backupstorage.ListObjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupstorage.ListObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &backupstorage.ListObjectsOutput{}
	fnCacher := func(out *backupstorage.ListObjectsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupStorageAPI.ListObjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupStorage) ListObjectsPagesWithContext(ctx context.Context, input *backupstorage.ListObjectsInput, fn func(*backupstorage.ListObjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupstorage.ListObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &backupstorage.ListObjectsOutput{}
	fnCacher := func(out *backupstorage.ListObjectsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupStorageAPI.ListObjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupStorage) ListObjectsWithContext(ctx context.Context, input *backupstorage.ListObjectsInput, opts ...request.Option) (*backupstorage.ListObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupstorage.ListObjectsOutput), nil
	}
	out, err := c.BackupStorageAPI.ListObjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

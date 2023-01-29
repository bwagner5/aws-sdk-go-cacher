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
package omicscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/omics"
	"github.com/aws/aws-sdk-go/service/omics/omicsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Omics struct {
	omicsiface.OmicsAPI
	cache *cache.Cache
}

func New(omicsapi omicsiface.OmicsAPI) *Omics {
	return &Omics{
		OmicsAPI: omicsapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Omics) BatchDeleteReadSet(input *omics.BatchDeleteReadSetInput) (*omics.BatchDeleteReadSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.BatchDeleteReadSetOutput), nil
	}
	out, err := c.OmicsAPI.BatchDeleteReadSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) BatchDeleteReadSetWithContext(ctx context.Context, input *omics.BatchDeleteReadSetInput, opts ...request.Option) (*omics.BatchDeleteReadSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.BatchDeleteReadSetOutput), nil
	}
	out, err := c.OmicsAPI.BatchDeleteReadSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetAnnotationImportJob(input *omics.GetAnnotationImportJobInput) (*omics.GetAnnotationImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetAnnotationImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetAnnotationImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetAnnotationImportJobWithContext(ctx context.Context, input *omics.GetAnnotationImportJobInput, opts ...request.Option) (*omics.GetAnnotationImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetAnnotationImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetAnnotationImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetAnnotationStore(input *omics.GetAnnotationStoreInput) (*omics.GetAnnotationStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetAnnotationStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetAnnotationStore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetAnnotationStoreWithContext(ctx context.Context, input *omics.GetAnnotationStoreInput, opts ...request.Option) (*omics.GetAnnotationStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetAnnotationStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetAnnotationStoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSet(input *omics.GetReadSetInput) (*omics.GetReadSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetActivationJob(input *omics.GetReadSetActivationJobInput) (*omics.GetReadSetActivationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetActivationJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetActivationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetActivationJobWithContext(ctx context.Context, input *omics.GetReadSetActivationJobInput, opts ...request.Option) (*omics.GetReadSetActivationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetActivationJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetActivationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetExportJob(input *omics.GetReadSetExportJobInput) (*omics.GetReadSetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetExportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetExportJobWithContext(ctx context.Context, input *omics.GetReadSetExportJobInput, opts ...request.Option) (*omics.GetReadSetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetExportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetImportJob(input *omics.GetReadSetImportJobInput) (*omics.GetReadSetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetImportJobWithContext(ctx context.Context, input *omics.GetReadSetImportJobInput, opts ...request.Option) (*omics.GetReadSetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetMetadata(input *omics.GetReadSetMetadataInput) (*omics.GetReadSetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetMetadataOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetMetadataWithContext(ctx context.Context, input *omics.GetReadSetMetadataInput, opts ...request.Option) (*omics.GetReadSetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetMetadataOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReadSetWithContext(ctx context.Context, input *omics.GetReadSetInput, opts ...request.Option) (*omics.GetReadSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReadSetOutput), nil
	}
	out, err := c.OmicsAPI.GetReadSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReference(input *omics.GetReferenceInput) (*omics.GetReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceOutput), nil
	}
	out, err := c.OmicsAPI.GetReference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceImportJob(input *omics.GetReferenceImportJobInput) (*omics.GetReferenceImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceImportJobWithContext(ctx context.Context, input *omics.GetReferenceImportJobInput, opts ...request.Option) (*omics.GetReferenceImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceMetadata(input *omics.GetReferenceMetadataInput) (*omics.GetReferenceMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceMetadataOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceMetadataWithContext(ctx context.Context, input *omics.GetReferenceMetadataInput, opts ...request.Option) (*omics.GetReferenceMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceMetadataOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceStore(input *omics.GetReferenceStoreInput) (*omics.GetReferenceStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceStore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceStoreWithContext(ctx context.Context, input *omics.GetReferenceStoreInput, opts ...request.Option) (*omics.GetReferenceStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceStoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetReferenceWithContext(ctx context.Context, input *omics.GetReferenceInput, opts ...request.Option) (*omics.GetReferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetReferenceOutput), nil
	}
	out, err := c.OmicsAPI.GetReferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRun(input *omics.GetRunInput) (*omics.GetRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunOutput), nil
	}
	out, err := c.OmicsAPI.GetRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRunGroup(input *omics.GetRunGroupInput) (*omics.GetRunGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunGroupOutput), nil
	}
	out, err := c.OmicsAPI.GetRunGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRunGroupWithContext(ctx context.Context, input *omics.GetRunGroupInput, opts ...request.Option) (*omics.GetRunGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunGroupOutput), nil
	}
	out, err := c.OmicsAPI.GetRunGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRunTask(input *omics.GetRunTaskInput) (*omics.GetRunTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunTaskOutput), nil
	}
	out, err := c.OmicsAPI.GetRunTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRunTaskWithContext(ctx context.Context, input *omics.GetRunTaskInput, opts ...request.Option) (*omics.GetRunTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunTaskOutput), nil
	}
	out, err := c.OmicsAPI.GetRunTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetRunWithContext(ctx context.Context, input *omics.GetRunInput, opts ...request.Option) (*omics.GetRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetRunOutput), nil
	}
	out, err := c.OmicsAPI.GetRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetSequenceStore(input *omics.GetSequenceStoreInput) (*omics.GetSequenceStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetSequenceStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetSequenceStore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetSequenceStoreWithContext(ctx context.Context, input *omics.GetSequenceStoreInput, opts ...request.Option) (*omics.GetSequenceStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetSequenceStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetSequenceStoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetVariantImportJob(input *omics.GetVariantImportJobInput) (*omics.GetVariantImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetVariantImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetVariantImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetVariantImportJobWithContext(ctx context.Context, input *omics.GetVariantImportJobInput, opts ...request.Option) (*omics.GetVariantImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetVariantImportJobOutput), nil
	}
	out, err := c.OmicsAPI.GetVariantImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetVariantStore(input *omics.GetVariantStoreInput) (*omics.GetVariantStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetVariantStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetVariantStore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetVariantStoreWithContext(ctx context.Context, input *omics.GetVariantStoreInput, opts ...request.Option) (*omics.GetVariantStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetVariantStoreOutput), nil
	}
	out, err := c.OmicsAPI.GetVariantStoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetWorkflow(input *omics.GetWorkflowInput) (*omics.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetWorkflowOutput), nil
	}
	out, err := c.OmicsAPI.GetWorkflow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) GetWorkflowWithContext(ctx context.Context, input *omics.GetWorkflowInput, opts ...request.Option) (*omics.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.GetWorkflowOutput), nil
	}
	out, err := c.OmicsAPI.GetWorkflowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListAnnotationImportJobs(input *omics.ListAnnotationImportJobsInput) (*omics.ListAnnotationImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListAnnotationImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListAnnotationImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListAnnotationImportJobsPages(input *omics.ListAnnotationImportJobsInput, fn func(*omics.ListAnnotationImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListAnnotationImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListAnnotationImportJobsOutput{}
	fnCacher := func(out *omics.ListAnnotationImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListAnnotationImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListAnnotationImportJobsPagesWithContext(ctx context.Context, input *omics.ListAnnotationImportJobsInput, fn func(*omics.ListAnnotationImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListAnnotationImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListAnnotationImportJobsOutput{}
	fnCacher := func(out *omics.ListAnnotationImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListAnnotationImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListAnnotationImportJobsWithContext(ctx context.Context, input *omics.ListAnnotationImportJobsInput, opts ...request.Option) (*omics.ListAnnotationImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListAnnotationImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListAnnotationImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListAnnotationStores(input *omics.ListAnnotationStoresInput) (*omics.ListAnnotationStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListAnnotationStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListAnnotationStores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListAnnotationStoresPages(input *omics.ListAnnotationStoresInput, fn func(*omics.ListAnnotationStoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListAnnotationStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListAnnotationStoresOutput{}
	fnCacher := func(out *omics.ListAnnotationStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListAnnotationStoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListAnnotationStoresPagesWithContext(ctx context.Context, input *omics.ListAnnotationStoresInput, fn func(*omics.ListAnnotationStoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListAnnotationStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListAnnotationStoresOutput{}
	fnCacher := func(out *omics.ListAnnotationStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListAnnotationStoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListAnnotationStoresWithContext(ctx context.Context, input *omics.ListAnnotationStoresInput, opts ...request.Option) (*omics.ListAnnotationStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListAnnotationStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListAnnotationStoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetActivationJobs(input *omics.ListReadSetActivationJobsInput) (*omics.ListReadSetActivationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetActivationJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetActivationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetActivationJobsPages(input *omics.ListReadSetActivationJobsInput, fn func(*omics.ListReadSetActivationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetActivationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetActivationJobsOutput{}
	fnCacher := func(out *omics.ListReadSetActivationJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetActivationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetActivationJobsPagesWithContext(ctx context.Context, input *omics.ListReadSetActivationJobsInput, fn func(*omics.ListReadSetActivationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetActivationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetActivationJobsOutput{}
	fnCacher := func(out *omics.ListReadSetActivationJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetActivationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetActivationJobsWithContext(ctx context.Context, input *omics.ListReadSetActivationJobsInput, opts ...request.Option) (*omics.ListReadSetActivationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetActivationJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetActivationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetExportJobs(input *omics.ListReadSetExportJobsInput) (*omics.ListReadSetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetExportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetExportJobsPages(input *omics.ListReadSetExportJobsInput, fn func(*omics.ListReadSetExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetExportJobsOutput{}
	fnCacher := func(out *omics.ListReadSetExportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetExportJobsPagesWithContext(ctx context.Context, input *omics.ListReadSetExportJobsInput, fn func(*omics.ListReadSetExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetExportJobsOutput{}
	fnCacher := func(out *omics.ListReadSetExportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetExportJobsWithContext(ctx context.Context, input *omics.ListReadSetExportJobsInput, opts ...request.Option) (*omics.ListReadSetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetExportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetImportJobs(input *omics.ListReadSetImportJobsInput) (*omics.ListReadSetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetImportJobsPages(input *omics.ListReadSetImportJobsInput, fn func(*omics.ListReadSetImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetImportJobsOutput{}
	fnCacher := func(out *omics.ListReadSetImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetImportJobsPagesWithContext(ctx context.Context, input *omics.ListReadSetImportJobsInput, fn func(*omics.ListReadSetImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetImportJobsOutput{}
	fnCacher := func(out *omics.ListReadSetImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetImportJobsWithContext(ctx context.Context, input *omics.ListReadSetImportJobsInput, opts ...request.Option) (*omics.ListReadSetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSets(input *omics.ListReadSetsInput) (*omics.ListReadSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReadSetsPages(input *omics.ListReadSetsInput, fn func(*omics.ListReadSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetsOutput{}
	fnCacher := func(out *omics.ListReadSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetsPagesWithContext(ctx context.Context, input *omics.ListReadSetsInput, fn func(*omics.ListReadSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReadSetsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReadSetsOutput{}
	fnCacher := func(out *omics.ListReadSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReadSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReadSetsWithContext(ctx context.Context, input *omics.ListReadSetsInput, opts ...request.Option) (*omics.ListReadSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReadSetsOutput), nil
	}
	out, err := c.OmicsAPI.ListReadSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferenceImportJobs(input *omics.ListReferenceImportJobsInput) (*omics.ListReferenceImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferenceImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReferenceImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferenceImportJobsPages(input *omics.ListReferenceImportJobsInput, fn func(*omics.ListReferenceImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferenceImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferenceImportJobsOutput{}
	fnCacher := func(out *omics.ListReferenceImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferenceImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferenceImportJobsPagesWithContext(ctx context.Context, input *omics.ListReferenceImportJobsInput, fn func(*omics.ListReferenceImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferenceImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferenceImportJobsOutput{}
	fnCacher := func(out *omics.ListReferenceImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferenceImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferenceImportJobsWithContext(ctx context.Context, input *omics.ListReferenceImportJobsInput, opts ...request.Option) (*omics.ListReferenceImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferenceImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListReferenceImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferenceStores(input *omics.ListReferenceStoresInput) (*omics.ListReferenceStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferenceStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListReferenceStores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferenceStoresPages(input *omics.ListReferenceStoresInput, fn func(*omics.ListReferenceStoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferenceStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferenceStoresOutput{}
	fnCacher := func(out *omics.ListReferenceStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferenceStoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferenceStoresPagesWithContext(ctx context.Context, input *omics.ListReferenceStoresInput, fn func(*omics.ListReferenceStoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferenceStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferenceStoresOutput{}
	fnCacher := func(out *omics.ListReferenceStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferenceStoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferenceStoresWithContext(ctx context.Context, input *omics.ListReferenceStoresInput, opts ...request.Option) (*omics.ListReferenceStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferenceStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListReferenceStoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferences(input *omics.ListReferencesInput) (*omics.ListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferencesOutput), nil
	}
	out, err := c.OmicsAPI.ListReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListReferencesPages(input *omics.ListReferencesInput, fn func(*omics.ListReferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferencesOutput{}
	fnCacher := func(out *omics.ListReferencesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferencesPagesWithContext(ctx context.Context, input *omics.ListReferencesInput, fn func(*omics.ListReferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListReferencesOutput{}
	fnCacher := func(out *omics.ListReferencesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListReferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListReferencesWithContext(ctx context.Context, input *omics.ListReferencesInput, opts ...request.Option) (*omics.ListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListReferencesOutput), nil
	}
	out, err := c.OmicsAPI.ListReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRunGroups(input *omics.ListRunGroupsInput) (*omics.ListRunGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunGroupsOutput), nil
	}
	out, err := c.OmicsAPI.ListRunGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRunGroupsPages(input *omics.ListRunGroupsInput, fn func(*omics.ListRunGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunGroupsOutput{}
	fnCacher := func(out *omics.ListRunGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunGroupsPagesWithContext(ctx context.Context, input *omics.ListRunGroupsInput, fn func(*omics.ListRunGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunGroupsOutput{}
	fnCacher := func(out *omics.ListRunGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunGroupsWithContext(ctx context.Context, input *omics.ListRunGroupsInput, opts ...request.Option) (*omics.ListRunGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunGroupsOutput), nil
	}
	out, err := c.OmicsAPI.ListRunGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRunTasks(input *omics.ListRunTasksInput) (*omics.ListRunTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunTasksOutput), nil
	}
	out, err := c.OmicsAPI.ListRunTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRunTasksPages(input *omics.ListRunTasksInput, fn func(*omics.ListRunTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunTasksOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunTasksOutput{}
	fnCacher := func(out *omics.ListRunTasksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunTasksPagesWithContext(ctx context.Context, input *omics.ListRunTasksInput, fn func(*omics.ListRunTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunTasksOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunTasksOutput{}
	fnCacher := func(out *omics.ListRunTasksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunTasksWithContext(ctx context.Context, input *omics.ListRunTasksInput, opts ...request.Option) (*omics.ListRunTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunTasksOutput), nil
	}
	out, err := c.OmicsAPI.ListRunTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRuns(input *omics.ListRunsInput) (*omics.ListRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunsOutput), nil
	}
	out, err := c.OmicsAPI.ListRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListRunsPages(input *omics.ListRunsInput, fn func(*omics.ListRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunsOutput{}
	fnCacher := func(out *omics.ListRunsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunsPagesWithContext(ctx context.Context, input *omics.ListRunsInput, fn func(*omics.ListRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListRunsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListRunsOutput{}
	fnCacher := func(out *omics.ListRunsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListRunsWithContext(ctx context.Context, input *omics.ListRunsInput, opts ...request.Option) (*omics.ListRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListRunsOutput), nil
	}
	out, err := c.OmicsAPI.ListRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListSequenceStores(input *omics.ListSequenceStoresInput) (*omics.ListSequenceStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListSequenceStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListSequenceStores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListSequenceStoresPages(input *omics.ListSequenceStoresInput, fn func(*omics.ListSequenceStoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListSequenceStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListSequenceStoresOutput{}
	fnCacher := func(out *omics.ListSequenceStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListSequenceStoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListSequenceStoresPagesWithContext(ctx context.Context, input *omics.ListSequenceStoresInput, fn func(*omics.ListSequenceStoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListSequenceStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListSequenceStoresOutput{}
	fnCacher := func(out *omics.ListSequenceStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListSequenceStoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListSequenceStoresWithContext(ctx context.Context, input *omics.ListSequenceStoresInput, opts ...request.Option) (*omics.ListSequenceStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListSequenceStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListSequenceStoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListTagsForResource(input *omics.ListTagsForResourceInput) (*omics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListTagsForResourceOutput), nil
	}
	out, err := c.OmicsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListTagsForResourceWithContext(ctx context.Context, input *omics.ListTagsForResourceInput, opts ...request.Option) (*omics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListTagsForResourceOutput), nil
	}
	out, err := c.OmicsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListVariantImportJobs(input *omics.ListVariantImportJobsInput) (*omics.ListVariantImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListVariantImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListVariantImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListVariantImportJobsPages(input *omics.ListVariantImportJobsInput, fn func(*omics.ListVariantImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListVariantImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListVariantImportJobsOutput{}
	fnCacher := func(out *omics.ListVariantImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListVariantImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListVariantImportJobsPagesWithContext(ctx context.Context, input *omics.ListVariantImportJobsInput, fn func(*omics.ListVariantImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListVariantImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListVariantImportJobsOutput{}
	fnCacher := func(out *omics.ListVariantImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListVariantImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListVariantImportJobsWithContext(ctx context.Context, input *omics.ListVariantImportJobsInput, opts ...request.Option) (*omics.ListVariantImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListVariantImportJobsOutput), nil
	}
	out, err := c.OmicsAPI.ListVariantImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListVariantStores(input *omics.ListVariantStoresInput) (*omics.ListVariantStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListVariantStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListVariantStores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListVariantStoresPages(input *omics.ListVariantStoresInput, fn func(*omics.ListVariantStoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListVariantStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListVariantStoresOutput{}
	fnCacher := func(out *omics.ListVariantStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListVariantStoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListVariantStoresPagesWithContext(ctx context.Context, input *omics.ListVariantStoresInput, fn func(*omics.ListVariantStoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListVariantStoresOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListVariantStoresOutput{}
	fnCacher := func(out *omics.ListVariantStoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListVariantStoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListVariantStoresWithContext(ctx context.Context, input *omics.ListVariantStoresInput, opts ...request.Option) (*omics.ListVariantStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListVariantStoresOutput), nil
	}
	out, err := c.OmicsAPI.ListVariantStoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListWorkflows(input *omics.ListWorkflowsInput) (*omics.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListWorkflowsOutput), nil
	}
	out, err := c.OmicsAPI.ListWorkflows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Omics) ListWorkflowsPages(input *omics.ListWorkflowsInput, fn func(*omics.ListWorkflowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListWorkflowsOutput{}
	fnCacher := func(out *omics.ListWorkflowsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListWorkflowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListWorkflowsPagesWithContext(ctx context.Context, input *omics.ListWorkflowsInput, fn func(*omics.ListWorkflowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*omics.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &omics.ListWorkflowsOutput{}
	fnCacher := func(out *omics.ListWorkflowsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.OmicsAPI.ListWorkflowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Omics) ListWorkflowsWithContext(ctx context.Context, input *omics.ListWorkflowsInput, opts ...request.Option) (*omics.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*omics.ListWorkflowsOutput), nil
	}
	out, err := c.OmicsAPI.ListWorkflowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
